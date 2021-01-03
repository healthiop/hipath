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

type ArithmeticExpression struct {
	evalLeft  pathsys.Evaluator
	op        pathsys.ArithmeticOps
	evalRight pathsys.Evaluator
}

func NewArithmeticExpression(evalLeft pathsys.Evaluator, op pathsys.ArithmeticOps, evalRight pathsys.Evaluator) *ArithmeticExpression {
	return &ArithmeticExpression{evalLeft, op, evalRight}
}

func (e *ArithmeticExpression) Evaluate(ctx pathsys.ContextAccessor, node interface{}, loop pathsys.Looper) (interface{}, error) {
	left, err := e.evalLeft.Evaluate(ctx, node, loop)
	if err != nil {
		return nil, err
	}
	right, err := e.evalRight.Evaluate(ctx, node, loop)
	if err != nil {
		return nil, err
	}

	left, right = unwrapCollection(left), unwrapCollection(right)
	if left == nil || right == nil {
		return nil, nil
	}

	leftOperand, ok := left.(pathsys.ArithmeticApplier)
	if !ok {
		return applyNonNumberArithmetic(left, e.op, right)
	}
	rightOperand, ok := right.(pathsys.DecimalValueAccessor)
	if !ok {
		return applyNonNumberArithmetic(left, e.op, right)
	}

	return leftOperand.Calc(rightOperand, e.op)
}

func applyNonNumberArithmetic(left interface{}, op pathsys.ArithmeticOps, right interface{}) (pathsys.AnyAccessor, error) {
	if op == pathsys.AdditionOp || op == pathsys.SubtractionOp {
		t, err := applyTemporalArithmetic(left, right,
			op == pathsys.SubtractionOp)
		if err != nil {
			return nil, err
		}
		if t != nil {
			return t, nil
		}
	}

	if op == pathsys.AdditionOp {
		if s := applyStringArithmetic(left, right); s != nil {
			return s, nil
		}
	}

	return nil, fmt.Errorf("operands %T and %T do not support arithmetic operation %c", left, op, right)
}

func applyStringArithmetic(left, right interface{}) pathsys.StringAccessor {
	var ok bool
	var leftString, rightString pathsys.Stringifier
	if leftString, ok = left.(pathsys.Stringifier); !ok {
		return nil
	}
	if rightString, ok = right.(pathsys.Stringifier); !ok {
		return nil
	}

	return pathsys.NewString(leftString.String() + rightString.String())
}

func applyTemporalArithmetic(left, right interface{}, negate bool) (pathsys.TemporalAccessor, error) {
	var ok bool
	var temporal pathsys.TemporalAccessor
	var quantity pathsys.QuantityAccessor
	if temporal, ok = left.(pathsys.TemporalAccessor); !ok {
		return nil, nil
	}
	if quantity, ok = right.(pathsys.QuantityAccessor); !ok {
		return nil, fmt.Errorf("only a quantity may be added to a temporal value: %T", right)
	}

	if negate {
		quantity = quantity.Negate().(pathsys.QuantityAccessor)
	}
	return temporal.Add(quantity)
}
