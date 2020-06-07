// Copyright (c) 2020, Volker Schmidt (volker@volsch.eu)
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its
//    contributors may be used to endorse or promote products derived from
//    this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package expression

import (
	"fmt"
	"github.com/volsch/gohipath/pathsys"
	"regexp"
	"strings"
)

type iifFunction struct {
	pathsys.BaseFunction
}

func newIIfFunction() *iifFunction {
	return &iifFunction{
		BaseFunction: pathsys.NewBaseFunction("iif", -1, 2, 3),
	}
}

func (f *iifFunction) Execute(ctx pathsys.ContextAccessor, node interface{}, args []interface{}, loop pathsys.Looper) (interface{}, error) {
	criterion := unwrapCollection(args[0])
	var criterionValue bool
	if criterion != nil {
		if b, ok := criterion.(pathsys.BooleanAccessor); !ok {
			return nil, fmt.Errorf("criterion must be a boolean: %T", criterion)
		} else {
			criterionValue = b.Bool()
		}
	}

	var res interface{}
	if criterionValue {
		res = args[1]
	} else if len(args) > 2 {
		res = args[2]
	}

	return res, nil
}

type toBooleanFunction struct {
	pathsys.BaseFunction
}

var toBooleanFunc = &toBooleanFunction{
	BaseFunction: pathsys.NewBaseFunction("toBoolean", -1, 0, 0),
}

func (f *toBooleanFunction) Execute(ctx pathsys.ContextAccessor, node interface{}, args []interface{}, loop pathsys.Looper) (interface{}, error) {
	any, err := convertibleAny(node)
	if any == nil || err != nil {
		return nil, err
	}

	if b, ok := any.(pathsys.BooleanAccessor); ok {
		return b, nil
	}
	if n, ok := any.(pathsys.NumberAccessor); ok {
		f := n.Float64()
		if f == 1.0 {
			return pathsys.True, nil
		}
		if f == 0.0 {
			return pathsys.False, nil
		}
	} else if s, ok := any.(pathsys.StringAccessor); ok {
		s := strings.ToLower(s.String())
		switch s {
		case "true", "t", "yes", "y", "1", "1.0":
			return pathsys.True, nil
		case "false", "f", "no", "n", "0", "0.0":
			return pathsys.False, nil
		}
	}

	return nil, nil
}

type convertsToBooleanFunction struct {
	convertsToFunction
}

func newConvertsToBooleanFunction() *convertsToBooleanFunction {
	return &convertsToBooleanFunction{
		convertsToFunction: convertsToFunction{
			BaseFunction: pathsys.NewBaseFunction("convertsToBoolean", -1, 0, 0),
			converter:    toBooleanFunc,
		},
	}
}

type toIntegerFunction struct {
	pathsys.BaseFunction
}

var toIntegerFunc = &toIntegerFunction{
	BaseFunction: pathsys.NewBaseFunction("toInteger", -1, 0, 0),
}

func (f *toIntegerFunction) Execute(ctx pathsys.ContextAccessor, node interface{}, args []interface{}, loop pathsys.Looper) (interface{}, error) {
	any, err := convertibleAny(node)
	if any == nil || err != nil {
		return nil, err
	}

	switch any.DataType() {
	case pathsys.IntegerDataType:
		return any, nil
	case pathsys.BooleanDataType:
		if any.(pathsys.BooleanAccessor).Bool() {
			return pathsys.NewInteger(1), nil
		}
		return pathsys.NewInteger(0), nil
	case pathsys.StringDataType:
		i, err := pathsys.ParseInteger(any.(pathsys.StringAccessor).String())
		if err == nil {
			return i, nil
		}
	}

	return nil, nil
}

type convertsToIntegerFunction struct {
	convertsToFunction
}

func newConvertsToIntegerFunction() *convertsToIntegerFunction {
	return &convertsToIntegerFunction{
		convertsToFunction: convertsToFunction{
			BaseFunction: pathsys.NewBaseFunction("convertsToInteger", -1, 0, 0),
			converter:    toIntegerFunc,
		},
	}
}

type toDecimalFunction struct {
	pathsys.BaseFunction
}

var toDecimalFunc = &toDecimalFunction{
	BaseFunction: pathsys.NewBaseFunction("toDecimal", -1, 0, 0),
}

func (f *toDecimalFunction) Execute(ctx pathsys.ContextAccessor, node interface{}, args []interface{}, loop pathsys.Looper) (interface{}, error) {
	any, err := convertibleAny(node)
	if any == nil || err != nil {
		return nil, err
	}

	if n, ok := any.(pathsys.NumberAccessor); ok {
		return n.Value(), nil
	}

	switch any.DataType() {
	case pathsys.BooleanDataType:
		if any.(pathsys.BooleanAccessor).Bool() {
			return pathsys.NewDecimalInt(1), nil
		}
		return pathsys.NewDecimalInt(0), nil
	case pathsys.StringDataType:
		d, err := pathsys.ParseDecimal(any.(pathsys.StringAccessor).String())
		if err == nil {
			return d, nil
		}
	}

	return nil, nil
}

type convertsToDecimalFunction struct {
	convertsToFunction
}

func newConvertsToDecimalFunction() *convertsToDecimalFunction {
	return &convertsToDecimalFunction{
		convertsToFunction: convertsToFunction{
			BaseFunction: pathsys.NewBaseFunction("convertsToDecimal", -1, 0, 0),
			converter:    toDecimalFunc,
		},
	}
}

type toDateFunction struct {
	pathsys.BaseFunction
}

var toDateFunc = &toDateFunction{
	BaseFunction: pathsys.NewBaseFunction("toDate", -1, 0, 0),
}

func (f *toDateFunction) Execute(ctx pathsys.ContextAccessor, node interface{}, args []interface{}, loop pathsys.Looper) (interface{}, error) {
	any, err := convertibleAny(node)
	if any == nil || err != nil {
		return nil, err
	}

	var d pathsys.DateAccessor
	if t, ok := any.(pathsys.DateTemporalAccessor); ok {
		d = t.Date()
	} else if s, ok := any.(pathsys.StringAccessor); ok {
		var err error
		d, err = pathsys.ParseDate(s.String())
		if err != nil {
			return nil, nil
		}
	}

	return d, nil
}

type convertsToDateFunction struct {
	convertsToFunction
}

func newConvertsToDateFunction() *convertsToDateFunction {
	return &convertsToDateFunction{
		convertsToFunction: convertsToFunction{
			BaseFunction: pathsys.NewBaseFunction("convertsToDate", -1, 0, 0),
			converter:    toDateFunc,
		},
	}
}

type toDateTimeFunction struct {
	pathsys.BaseFunction
}

var toDateTimeFunc = &toDateTimeFunction{
	BaseFunction: pathsys.NewBaseFunction("toDateTime", -1, 0, 0),
}

func (f *toDateTimeFunction) Execute(ctx pathsys.ContextAccessor, node interface{}, args []interface{}, loop pathsys.Looper) (interface{}, error) {
	any, err := convertibleAny(node)
	if any == nil || err != nil {
		return nil, err
	}

	var d pathsys.DateTimeAccessor
	if t, ok := any.(pathsys.DateTemporalAccessor); ok {
		d = t.DateTime()
	} else if s, ok := any.(pathsys.StringAccessor); ok {
		var err error
		d, err = pathsys.ParseDateTime(s.String())
		if err != nil {
			return nil, nil
		}
	}

	return d, nil
}

type convertsToDateTimeFunction struct {
	convertsToFunction
}

func newConvertsToDateTimeFunction() *convertsToDateTimeFunction {
	return &convertsToDateTimeFunction{
		convertsToFunction: convertsToFunction{
			BaseFunction: pathsys.NewBaseFunction("convertsToDateTime", -1, 0, 0),
			converter:    toDateTimeFunc,
		},
	}
}

type toQuantityFunction struct {
	pathsys.BaseFunction
}

var toQuantityFunc = &toQuantityFunction{
	BaseFunction: pathsys.NewBaseFunction("toQuantity", -1, 0, 1),
}

var quantityStringPattern = regexp.MustCompile("^([+\\-]?\\d+(?:\\.\\d+)?)\\s*('(?:[^']+)'|(?:[a-zA-Z]+))?$")

func (f *toQuantityFunction) Execute(ctx pathsys.ContextAccessor, node interface{}, args []interface{}, loop pathsys.Looper) (interface{}, error) {
	any, err := convertibleAny(node)
	if any == nil || err != nil {
		return nil, err
	}

	var q pathsys.QuantityAccessor
	if o, ok := any.(pathsys.QuantityAccessor); ok {
		q = o
	} else if n, ok := any.(pathsys.NumberAccessor); ok {
		v := n.Value()
		q = pathsys.NewQuantity(v, pathsys.DefaultQuantityUnit.Name(v))
	} else if b, ok := any.(pathsys.BooleanAccessor); ok {
		var v pathsys.DecimalAccessor
		if b.Bool() {
			v = pathsys.DecimalOne
		} else {
			v = pathsys.DecimalZero
		}
		q = pathsys.NewQuantity(v, pathsys.DefaultQuantityUnit.Name(v))
	} else if s, ok := any.(pathsys.StringAccessor); ok {
		m := quantityStringPattern.FindStringSubmatch(s.String())
		if m == nil {
			return nil, nil
		}
		v, _ := pathsys.ParseDecimal(m[1])
		u, exp := pathsys.QuantityUnitWithName(parseStringLiteral(m[2], stringDelimiterChar))
		if u == nil {
			q = pathsys.NewQuantity(v, nil)
		} else {
			q = pathsys.NewQuantity(v, u.NameWithExp(v, exp))
		}
	}

	if q != nil && len(args) > 0 {
		if s, ok := args[0].(pathsys.StringAccessor); !ok {
			return nil, fmt.Errorf("conversion unit is no string: %T", args[0])
		} else {
			q = q.ToUnit(s)
		}
	}

	return q, nil
}

type convertsToQuantityFunction struct {
	convertsToFunction
}

func newConvertsToQuantityFunction() *convertsToQuantityFunction {
	return &convertsToQuantityFunction{
		convertsToFunction: convertsToFunction{
			BaseFunction: pathsys.NewBaseFunction("convertsToQuantity", -1, 0, 1),
			converter:    toQuantityFunc,
		},
	}
}

type toStringFunction struct {
	pathsys.BaseFunction
}

var toStringFunc = &toStringFunction{
	BaseFunction: pathsys.NewBaseFunction("toString", -1, 0, 0),
}

func (f *toStringFunction) Execute(ctx pathsys.ContextAccessor, node interface{}, args []interface{}, loop pathsys.Looper) (interface{}, error) {
	any, err := convertibleAny(node)
	if any == nil || err != nil {
		return nil, err
	}

	return pathsys.NewString(any.(pathsys.Stringifier).String()), nil
}

type convertsToStringFunction struct {
	convertsToFunction
}

func newConvertsToStringFunction() *convertsToStringFunction {
	return &convertsToStringFunction{
		convertsToFunction: convertsToFunction{
			BaseFunction: pathsys.NewBaseFunction("convertsToString", -1, 0, 0),
			converter:    toStringFunc,
		},
	}
}

type toTimeFunction struct {
	pathsys.BaseFunction
}

var toTimeFunc = &toTimeFunction{
	BaseFunction: pathsys.NewBaseFunction("toTime", -1, 0, 0),
}

func (f *toTimeFunction) Execute(ctx pathsys.ContextAccessor, node interface{}, args []interface{}, loop pathsys.Looper) (interface{}, error) {
	any, err := convertibleAny(node)
	if any == nil || err != nil {
		return nil, err
	}

	var d pathsys.TimeAccessor
	if t, ok := any.(pathsys.TimeAccessor); ok {
		d = t
	} else if s, ok := any.(pathsys.StringAccessor); ok {
		var err error
		d, err = pathsys.ParseTime(s.String())
		if err != nil {
			return nil, nil
		}
	}

	return d, nil
}

type convertsToTimeFunction struct {
	convertsToFunction
}

func newConvertsToTimeFunction() *convertsToTimeFunction {
	return &convertsToTimeFunction{
		convertsToFunction: convertsToFunction{
			BaseFunction: pathsys.NewBaseFunction("convertsToTime", -1, 0, 0),
			converter:    toTimeFunc,
		},
	}
}

type convertsToFunction struct {
	pathsys.BaseFunction
	converter pathsys.FunctionExecutor
}

func (f *convertsToFunction) Execute(ctx pathsys.ContextAccessor, node interface{}, args []interface{}, loop pathsys.Looper) (interface{}, error) {
	if emptyCollection(node) {
		return pathsys.False, nil
	}

	res, err := f.converter.Execute(ctx, node, args, loop)
	if err != nil {
		return nil, err
	}

	if res != nil {
		return pathsys.True, nil
	}
	return pathsys.False, nil
}

func convertibleAny(node interface{}) (pathsys.AnyAccessor, error) {
	value := unwrapCollection(node)
	if value == nil {
		return nil, nil
	}

	if any, ok := value.(pathsys.AnyAccessor); !ok {
		return nil, nil
	} else {
		if _, ok := any.(pathsys.CollectionAccessor); ok {
			return nil, fmt.Errorf("collection with multiple items cannot be converted")
		}
		return any, nil
	}
}
