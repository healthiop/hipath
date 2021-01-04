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

type StringConcatExpression struct {
	evalLeft  hipathsys.Evaluator
	evalRight hipathsys.Evaluator
}

func NewStringConcatExpression(evalLeft hipathsys.Evaluator, evalRight hipathsys.Evaluator) *StringConcatExpression {
	return &StringConcatExpression{evalLeft, evalRight}
}

func (e *StringConcatExpression) Evaluate(ctx hipathsys.ContextAccessor, node interface{}, loop hipathsys.Looper) (interface{}, error) {
	left, err := e.evalLeft.Evaluate(ctx, node, loop)
	if err != nil {
		return nil, err
	}
	right, err := e.evalRight.Evaluate(ctx, node, loop)
	if err != nil {
		return nil, err
	}

	left, right = unwrapCollection(left), unwrapCollection(right)
	if left == nil && right == nil {
		return hipathsys.EmptyString, nil
	}

	var ok bool
	var leftString, rightString hipathsys.Stringifier
	if left != nil {
		if leftString, ok = left.(hipathsys.Stringifier); !ok {
			return nil, fmt.Errorf("left operand is not string: %T", left)
		}
	}
	if right != nil {
		if rightString, ok = right.(hipathsys.Stringifier); !ok {
			return nil, fmt.Errorf("right operand is not string: %T", right)
		}
	}

	if leftString == nil {
		return rightString, nil
	}
	if rightString == nil {
		return leftString, nil
	}
	return hipathsys.NewString(leftString.String() + rightString.String()), nil
}
