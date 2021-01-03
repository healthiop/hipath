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

type StringConcatExpression struct {
	evalLeft  pathsys.Evaluator
	evalRight pathsys.Evaluator
}

func NewStringConcatExpression(evalLeft pathsys.Evaluator, evalRight pathsys.Evaluator) *StringConcatExpression {
	return &StringConcatExpression{evalLeft, evalRight}
}

func (e *StringConcatExpression) Evaluate(ctx pathsys.ContextAccessor, node interface{}, loop pathsys.Looper) (interface{}, error) {
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
		return pathsys.EmptyString, nil
	}

	var ok bool
	var leftString, rightString pathsys.Stringifier
	if left != nil {
		if leftString, ok = left.(pathsys.Stringifier); !ok {
			return nil, fmt.Errorf("left operand is not string: %T", left)
		}
	}
	if right != nil {
		if rightString, ok = right.(pathsys.Stringifier); !ok {
			return nil, fmt.Errorf("right operand is not string: %T", right)
		}
	}

	if leftString == nil {
		return rightString, nil
	}
	if rightString == nil {
		return leftString, nil
	}
	return pathsys.NewString(leftString.String() + rightString.String()), nil
}
