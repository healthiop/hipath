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
	"github.com/volsch/gohipath/pathsys"
)

type BooleanOp int

const (
	AndOp BooleanOp = iota + 1
	OrOp
	XOrOp
	ImpliesOp
)

type BooleanExpression struct {
	evalLeft  pathsys.Evaluator
	op        BooleanOp
	evalRight pathsys.Evaluator
}

func NewBooleanExpression(evalLeft pathsys.Evaluator, op BooleanOp, evalRight pathsys.Evaluator) *BooleanExpression {
	return &BooleanExpression{evalLeft, op, evalRight}
}

func (e *BooleanExpression) Evaluate(ctx pathsys.ContextAccessor, node interface{}, loop pathsys.Looper) (interface{}, error) {
	left, err := e.evalLeft.Evaluate(ctx, node, loop)
	if err != nil {
		return nil, err
	}
	right, err := e.evalRight.Evaluate(ctx, node, loop)
	if err != nil {
		return nil, err
	}

	leftBool, err := unwrapBooleanCollection(left)
	if err != nil {
		return nil, err
	}
	rightBool, err := unwrapBooleanCollection(right)
	if err != nil {
		return nil, err
	}

	if leftBool == nil && rightBool == nil {
		return nil, nil
	}

	if e.op == ImpliesOp {
		if leftBool == nil {
			if rightBool.Bool() {
				return pathsys.True, nil
			}
			return nil, nil
		}
		if leftBool.Bool() {
			return rightBool, nil
		}
		return pathsys.True, nil
	} else {
		if leftBool == nil || rightBool == nil {
			return nil, nil
		}

		switch e.op {
		case AndOp:
			return pathsys.BooleanOf(leftBool.Bool() && rightBool.Bool()), nil
		case OrOp:
			return pathsys.BooleanOf(leftBool.Bool() || rightBool.Bool()), nil
		case XOrOp:
			return pathsys.BooleanOf(leftBool.Bool() != rightBool.Bool()), nil
		default:
			panic(fmt.Sprintf("unhandled boolean operator: %d", e.op))
		}
	}
}

func unwrapBooleanCollection(node interface{}) (pathsys.BooleanAccessor, error) {
	if node == nil {
		return nil, nil
	}

	if col, ok := node.(pathsys.CollectionAccessor); ok {
		if col.Empty() {
			return nil, nil
		}
		if col.Count() > 1 {
			return nil, fmt.Errorf("multi-valued collection cannot be converted to a boolean")
		}

		v := col.Get(0)
		if b, ok := v.(pathsys.BooleanAccessor); ok {
			return b, nil
		}
		if v == nil {
			return nil, nil
		}
		return pathsys.True, nil
	}

	if b, ok := node.(pathsys.BooleanAccessor); ok {
		return b, nil
	}

	return nil, fmt.Errorf("value cannot be converted to a boolean: %T", node)
}
