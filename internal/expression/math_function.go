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
)

type absFunction struct {
	pathsys.BaseFunction
}

func newAbsFunction() *absFunction {
	return &absFunction{
		BaseFunction: pathsys.NewBaseFunction("abs", -1, 0, 0),
	}
}

func (f *absFunction) Execute(_ pathsys.ContextAccessor, node interface{}, _ []interface{}, _ pathsys.Looper) (interface{}, error) {
	a, err := arithmeticNode(node)
	if a == nil || err != nil {
		return nil, err
	}

	return a.Abs(), nil
}

type ceilingFunction struct {
	pathsys.BaseFunction
}

func newCeilingFunction() *ceilingFunction {
	return &ceilingFunction{
		BaseFunction: pathsys.NewBaseFunction("ceiling", -1, 0, 0),
	}
}

func (f *ceilingFunction) Execute(_ pathsys.ContextAccessor, node interface{}, _ []interface{}, _ pathsys.Looper) (interface{}, error) {
	n, err := numberNode(node)
	if n == nil || err != nil {
		return nil, err
	}

	return n.Ceiling(), nil
}

type expFunction struct {
	pathsys.BaseFunction
}

func newExpFunction() *expFunction {
	return &expFunction{
		BaseFunction: pathsys.NewBaseFunction("exp", -1, 0, 0),
	}
}

func (f *expFunction) Execute(_ pathsys.ContextAccessor, node interface{}, _ []interface{}, _ pathsys.Looper) (interface{}, error) {
	n, err := numberNode(node)
	if n == nil || err != nil {
		return nil, err
	}

	return n.Exp(), nil
}

type floorFunction struct {
	pathsys.BaseFunction
}

func newFloorFunction() *floorFunction {
	return &floorFunction{
		BaseFunction: pathsys.NewBaseFunction("floor", -1, 0, 0),
	}
}

func (f *floorFunction) Execute(_ pathsys.ContextAccessor, node interface{}, _ []interface{}, _ pathsys.Looper) (interface{}, error) {
	n, err := numberNode(node)
	if n == nil || err != nil {
		return nil, err
	}

	return n.Floor(), nil
}

type lnFunction struct {
	pathsys.BaseFunction
}

func newLnFunction() *lnFunction {
	return &lnFunction{
		BaseFunction: pathsys.NewBaseFunction("ln", -1, 0, 0),
	}
}

func (f *lnFunction) Execute(_ pathsys.ContextAccessor, node interface{}, _ []interface{}, _ pathsys.Looper) (interface{}, error) {
	n, err := numberNode(node)
	if n == nil || err != nil {
		return nil, err
	}

	return n.Ln()
}

type logFunction struct {
	pathsys.BaseFunction
}

func newLogFunction() *logFunction {
	return &logFunction{
		BaseFunction: pathsys.NewBaseFunction("log", -1, 1, 1),
	}
}

func (f *logFunction) Execute(_ pathsys.ContextAccessor, node interface{}, args []interface{}, _ pathsys.Looper) (interface{}, error) {
	n, err := numberNode(node)
	if n == nil || err != nil {
		return nil, err
	}

	base, err := numberNode(args[0])
	if base == nil || err != nil {
		return nil, err
	}

	return n.Log(base)
}

type powerFunction struct {
	pathsys.BaseFunction
}

func newPowerFunction() *powerFunction {
	return &powerFunction{
		BaseFunction: pathsys.NewBaseFunction("power", -1, 1, 1),
	}
}

func (f *powerFunction) Execute(_ pathsys.ContextAccessor, node interface{}, args []interface{}, _ pathsys.Looper) (interface{}, error) {
	n, err := numberNode(node)
	if n == nil || err != nil {
		return nil, err
	}

	exponent, err := numberNode(args[0])
	if exponent == nil || err != nil {
		return nil, err
	}

	r, ok := n.Power(exponent)
	if !ok {
		return nil, nil
	}
	return r, nil
}

type roundFunction struct {
	pathsys.BaseFunction
}

func newRoundFunction() *roundFunction {
	return &roundFunction{
		BaseFunction: pathsys.NewBaseFunction("round", -1, 0, 1),
	}
}

func (f *roundFunction) Execute(_ pathsys.ContextAccessor, node interface{}, args []interface{}, _ pathsys.Looper) (interface{}, error) {
	n, err := numberNode(node)
	if n == nil || err != nil {
		return nil, err
	}

	precision := int32(0)
	if len(args) > 0 {
		p, err := integerNode(args[0])
		if p == nil || err != nil {
			return nil, err
		}
		precision = p.Int()
	}

	r, err := n.Round(precision)
	if err != nil {
		return nil, err
	}

	if i, ok := r.(pathsys.IntegerAccessor); ok {
		return pathsys.NewDecimal(i.Decimal()), nil
	}
	return r, nil
}

type sqrtFunction struct {
	pathsys.BaseFunction
}

func newSqrtFunction() *sqrtFunction {
	return &sqrtFunction{
		BaseFunction: pathsys.NewBaseFunction("sqrt", -1, 0, 0),
	}
}

func (f *sqrtFunction) Execute(_ pathsys.ContextAccessor, node interface{}, _ []interface{}, _ pathsys.Looper) (interface{}, error) {
	n, err := numberNode(node)
	if n == nil || err != nil {
		return nil, err
	}

	r, ok := n.Sqrt()
	if !ok {
		return nil, nil
	}
	return r, nil
}

type truncateFunction struct {
	pathsys.BaseFunction
}

func newTruncateFunction() *truncateFunction {
	return &truncateFunction{
		BaseFunction: pathsys.NewBaseFunction("truncate", -1, 0, 0),
	}
}

func (f *truncateFunction) Execute(_ pathsys.ContextAccessor, node interface{}, _ []interface{}, _ pathsys.Looper) (interface{}, error) {
	n, err := numberNode(node)
	if n == nil || err != nil {
		return nil, err
	}

	t := n.Truncate(0)
	if i, ok := t.(pathsys.IntegerAccessor); ok {
		return i, nil
	}
	return pathsys.NewInteger(t.Int()), nil
}

func arithmeticNode(node interface{}) (pathsys.ArithmeticApplier, error) {
	value := unwrapCollection(node)
	if value == nil {
		return nil, nil
	}

	if a, ok := value.(pathsys.ArithmeticApplier); !ok {
		return nil, fmt.Errorf("arithmetic cannot be applied: %T", value)
	} else {
		return a, nil
	}
}

func numberNode(node interface{}) (pathsys.NumberAccessor, error) {
	value := unwrapCollection(node)
	if value == nil {
		return nil, nil
	}

	if a, ok := value.(pathsys.NumberAccessor); !ok {
		return nil, fmt.Errorf("not a number: %T", value)
	} else {
		return a, nil
	}
}

func integerNode(node interface{}) (pathsys.IntegerAccessor, error) {
	value := unwrapCollection(node)
	if value == nil {
		return nil, nil
	}

	if a, ok := value.(pathsys.IntegerAccessor); !ok {
		return nil, fmt.Errorf("not an integer: %T", value)
	} else {
		return a, nil
	}
}
