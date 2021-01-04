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
)

type ComparisonOp int

const (
	LessThanOp ComparisonOp = iota + 1
	LessOrEqualThanOp
	GreaterOrEqualThanOp
	GreaterThanOp
)

type ComparisonExpression struct {
	evalLeft  hipathsys.Evaluator
	op        ComparisonOp
	evalRight hipathsys.Evaluator
}

func NewComparisonExpression(evalLeft hipathsys.Evaluator, op ComparisonOp, evalRight hipathsys.Evaluator) *ComparisonExpression {
	return &ComparisonExpression{evalLeft, op, evalRight}
}

func (e *ComparisonExpression) Evaluate(ctx hipathsys.ContextAccessor, node interface{}, loop hipathsys.Looper) (interface{}, error) {
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

	var ok bool
	var leftCmp, rightCmp hipathsys.Comparator
	if leftCmp, ok = left.(hipathsys.Comparator); !ok {
		return nil, fmt.Errorf("operand cannot be used for comparison: %T", left)
	}
	if rightCmp, ok = right.(hipathsys.Comparator); !ok {
		return nil, fmt.Errorf("operand cannot be used for comparison: %T", right)
	}

	res, status := leftCmp.Compare(rightCmp)
	if status == hipathsys.Empty {
		return nil, nil
	}
	if status != hipathsys.Evaluated {
		return nil, fmt.Errorf("operands cannot be compared: %T <> %T", leftCmp, rightCmp)
	}

	var b bool
	switch e.op {
	case LessOrEqualThanOp:
		b = res <= 0
	case LessThanOp:
		b = res < 0
	case GreaterThanOp:
		b = res > 0
	case GreaterOrEqualThanOp:
		b = res >= 0
	default:
		panic(fmt.Sprintf("unhandled comparison operator: %d", e.op))
	}
	return hipathsys.BooleanOf(b), nil
}
