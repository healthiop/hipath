// Copyright (c) 2020-2021, Volker Schmidt (volker@volsch.eu)
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
	"github.com/healthiop/hipath/hipathsys"
	"regexp"
	"strings"
)

type iifFunction struct {
	hipathsys.BaseFunction
}

func newIIfFunction() *iifFunction {
	return &iifFunction{
		BaseFunction: hipathsys.NewBaseFunction("iif", -1, 2, 3),
	}
}

func (f *iifFunction) Execute(_ hipathsys.ContextAccessor, _ interface{}, args []interface{}, _ hipathsys.Looper) (interface{}, error) {
	criterion := unwrapCollection(args[0])
	var criterionValue bool
	if criterion != nil {
		if b, ok := criterion.(hipathsys.BooleanAccessor); !ok {
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
	hipathsys.BaseFunction
}

var toBooleanFunc = &toBooleanFunction{
	BaseFunction: hipathsys.NewBaseFunction("toBoolean", -1, 0, 0),
}

func (f *toBooleanFunction) Execute(_ hipathsys.ContextAccessor, node interface{}, _ []interface{}, _ hipathsys.Looper) (interface{}, error) {
	any, err := convertibleAny(node)
	if any == nil || err != nil {
		return nil, err
	}

	if b, ok := any.(hipathsys.BooleanAccessor); ok {
		return b, nil
	}
	if n, ok := any.(hipathsys.NumberAccessor); ok {
		f := n.Float64()
		if f == 1.0 {
			return hipathsys.True, nil
		}
		if f == 0.0 {
			return hipathsys.False, nil
		}
	} else if s, ok := any.(hipathsys.StringAccessor); ok {
		s := strings.ToLower(s.String())
		switch s {
		case "true", "t", "yes", "y", "1", "1.0":
			return hipathsys.True, nil
		case "false", "f", "no", "n", "0", "0.0":
			return hipathsys.False, nil
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
			BaseFunction: hipathsys.NewBaseFunction("convertsToBoolean", -1, 0, 0),
			converter:    toBooleanFunc,
		},
	}
}

type toIntegerFunction struct {
	hipathsys.BaseFunction
}

var toIntegerFunc = &toIntegerFunction{
	BaseFunction: hipathsys.NewBaseFunction("toInteger", -1, 0, 0),
}

func (f *toIntegerFunction) Execute(_ hipathsys.ContextAccessor, node interface{}, _ []interface{}, _ hipathsys.Looper) (interface{}, error) {
	any, err := convertibleAny(node)
	if any == nil || err != nil {
		return nil, err
	}

	switch any.DataType() {
	case hipathsys.IntegerDataType:
		return any, nil
	case hipathsys.BooleanDataType:
		if any.(hipathsys.BooleanAccessor).Bool() {
			return hipathsys.NewInteger(1), nil
		}
		return hipathsys.NewInteger(0), nil
	case hipathsys.StringDataType:
		i, err := hipathsys.ParseInteger(any.(hipathsys.StringAccessor).String())
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
			BaseFunction: hipathsys.NewBaseFunction("convertsToInteger", -1, 0, 0),
			converter:    toIntegerFunc,
		},
	}
}

type toDecimalFunction struct {
	hipathsys.BaseFunction
}

var toDecimalFunc = &toDecimalFunction{
	BaseFunction: hipathsys.NewBaseFunction("toDecimal", -1, 0, 0),
}

func (f *toDecimalFunction) Execute(_ hipathsys.ContextAccessor, node interface{}, _ []interface{}, _ hipathsys.Looper) (interface{}, error) {
	any, err := convertibleAny(node)
	if any == nil || err != nil {
		return nil, err
	}

	if n, ok := any.(hipathsys.NumberAccessor); ok {
		return n.Value(), nil
	}

	switch any.DataType() {
	case hipathsys.BooleanDataType:
		if any.(hipathsys.BooleanAccessor).Bool() {
			return hipathsys.NewDecimalInt(1), nil
		}
		return hipathsys.NewDecimalInt(0), nil
	case hipathsys.StringDataType:
		d, err := hipathsys.ParseDecimal(any.(hipathsys.StringAccessor).String())
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
			BaseFunction: hipathsys.NewBaseFunction("convertsToDecimal", -1, 0, 0),
			converter:    toDecimalFunc,
		},
	}
}

type toDateFunction struct {
	hipathsys.BaseFunction
}

var toDateFunc = &toDateFunction{
	BaseFunction: hipathsys.NewBaseFunction("toDate", -1, 0, 0),
}

func (f *toDateFunction) Execute(_ hipathsys.ContextAccessor, node interface{}, _ []interface{}, _ hipathsys.Looper) (interface{}, error) {
	any, err := convertibleAny(node)
	if any == nil || err != nil {
		return nil, err
	}

	var d hipathsys.DateAccessor
	if t, ok := any.(hipathsys.DateTemporalAccessor); ok {
		d = t.Date()
	} else if s, ok := any.(hipathsys.StringAccessor); ok {
		var err error
		d, err = hipathsys.ParseDate(s.String())
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
			BaseFunction: hipathsys.NewBaseFunction("convertsToDate", -1, 0, 0),
			converter:    toDateFunc,
		},
	}
}

type toDateTimeFunction struct {
	hipathsys.BaseFunction
}

var toDateTimeFunc = &toDateTimeFunction{
	BaseFunction: hipathsys.NewBaseFunction("toDateTime", -1, 0, 0),
}

func (f *toDateTimeFunction) Execute(_ hipathsys.ContextAccessor, node interface{}, _ []interface{}, _ hipathsys.Looper) (interface{}, error) {
	any, err := convertibleAny(node)
	if any == nil || err != nil {
		return nil, err
	}

	var d hipathsys.DateTimeAccessor
	if t, ok := any.(hipathsys.DateTemporalAccessor); ok {
		d = t.DateTime()
	} else if s, ok := any.(hipathsys.StringAccessor); ok {
		var err error
		d, err = hipathsys.ParseDateTime(s.String())
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
			BaseFunction: hipathsys.NewBaseFunction("convertsToDateTime", -1, 0, 0),
			converter:    toDateTimeFunc,
		},
	}
}

type toQuantityFunction struct {
	hipathsys.BaseFunction
}

var toQuantityFunc = &toQuantityFunction{
	BaseFunction: hipathsys.NewBaseFunction("toQuantity", -1, 0, 1),
}

var quantityStringPattern = regexp.MustCompile("^([+\\-]?\\d+(?:\\.\\d+)?)\\s*('(?:[^']+)'|(?:[a-zA-Z]+))?$")

func (f *toQuantityFunction) Execute(_ hipathsys.ContextAccessor, node interface{}, args []interface{}, _ hipathsys.Looper) (interface{}, error) {
	any, err := convertibleAny(node)
	if any == nil || err != nil {
		return nil, err
	}

	var q hipathsys.QuantityAccessor
	if o, ok := any.(hipathsys.QuantityAccessor); ok {
		q = o
	} else if n, ok := any.(hipathsys.NumberAccessor); ok {
		v := n.Value()
		q = hipathsys.NewQuantity(v, hipathsys.DefaultQuantityUnit.Name(v))
	} else if b, ok := any.(hipathsys.BooleanAccessor); ok {
		var v hipathsys.DecimalAccessor
		if b.Bool() {
			v = hipathsys.DecimalOne
		} else {
			v = hipathsys.DecimalZero
		}
		q = hipathsys.NewQuantity(v, hipathsys.DefaultQuantityUnit.Name(v))
	} else if s, ok := any.(hipathsys.StringAccessor); ok {
		m := quantityStringPattern.FindStringSubmatch(s.String())
		if m == nil {
			return nil, nil
		}
		v, _ := hipathsys.ParseDecimal(m[1])
		u, exp := hipathsys.QuantityUnitWithName(parseStringLiteral(m[2], stringDelimiterChar))
		if u == nil {
			q = hipathsys.NewQuantity(v, nil)
		} else {
			q = hipathsys.NewQuantity(v, u.NameWithExp(v, exp))
		}
	}

	if q != nil && len(args) > 0 {
		if s, ok := args[0].(hipathsys.StringAccessor); !ok {
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
			BaseFunction: hipathsys.NewBaseFunction("convertsToQuantity", -1, 0, 1),
			converter:    toQuantityFunc,
		},
	}
}

type toStringFunction struct {
	hipathsys.BaseFunction
}

var toStringFunc = &toStringFunction{
	BaseFunction: hipathsys.NewBaseFunction("toString", -1, 0, 0),
}

func (f *toStringFunction) Execute(_ hipathsys.ContextAccessor, node interface{}, _ []interface{}, _ hipathsys.Looper) (interface{}, error) {
	any, err := convertibleAny(node)
	if any == nil || err != nil {
		return nil, err
	}

	return hipathsys.NewString(any.(hipathsys.Stringifier).String()), nil
}

type convertsToStringFunction struct {
	convertsToFunction
}

func newConvertsToStringFunction() *convertsToStringFunction {
	return &convertsToStringFunction{
		convertsToFunction: convertsToFunction{
			BaseFunction: hipathsys.NewBaseFunction("convertsToString", -1, 0, 0),
			converter:    toStringFunc,
		},
	}
}

type toTimeFunction struct {
	hipathsys.BaseFunction
}

var toTimeFunc = &toTimeFunction{
	BaseFunction: hipathsys.NewBaseFunction("toTime", -1, 0, 0),
}

func (f *toTimeFunction) Execute(_ hipathsys.ContextAccessor, node interface{}, _ []interface{}, _ hipathsys.Looper) (interface{}, error) {
	any, err := convertibleAny(node)
	if any == nil || err != nil {
		return nil, err
	}

	var d hipathsys.TimeAccessor
	if t, ok := any.(hipathsys.TimeAccessor); ok {
		d = t
	} else if s, ok := any.(hipathsys.StringAccessor); ok {
		var err error
		d, err = hipathsys.ParseTime(s.String())
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
			BaseFunction: hipathsys.NewBaseFunction("convertsToTime", -1, 0, 0),
			converter:    toTimeFunc,
		},
	}
}

type convertsToFunction struct {
	hipathsys.BaseFunction
	converter hipathsys.FunctionExecutor
}

func (f *convertsToFunction) Execute(ctx hipathsys.ContextAccessor, node interface{}, args []interface{}, loop hipathsys.Looper) (interface{}, error) {
	if emptyCollection(node) {
		return hipathsys.False, nil
	}

	res, err := f.converter.Execute(ctx, node, args, loop)
	if err != nil {
		return nil, err
	}

	if res != nil {
		return hipathsys.True, nil
	}
	return hipathsys.False, nil
}

func convertibleAny(node interface{}) (hipathsys.AnyAccessor, error) {
	value := unwrapCollection(node)
	if value == nil {
		return nil, nil
	}

	if any, ok := value.(hipathsys.AnyAccessor); !ok {
		return nil, nil
	} else {
		if _, ok := any.(hipathsys.CollectionAccessor); ok {
			return nil, fmt.Errorf("collection with multiple items cannot be converted")
		}
		return any, nil
	}
}
