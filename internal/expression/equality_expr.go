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
	"github.com/healthiop/hipath/hipathsys"
)

type EqualityExpression struct {
	not        bool
	equivalent bool
	evalLeft   hipathsys.Evaluator
	evalRight  hipathsys.Evaluator
}

func NewEqualityExpression(not bool, equivalent bool, evalLeft hipathsys.Evaluator, evalRight hipathsys.Evaluator) *EqualityExpression {
	return &EqualityExpression{not, equivalent, evalLeft, evalRight}
}

func (e *EqualityExpression) Evaluate(ctx hipathsys.ContextAccessor, node interface{}, loop hipathsys.Looper) (interface{}, error) {
	res, err := e.evaluateInternally(ctx, node, loop)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}
	if e.not {
		return res.(hipathsys.BooleanAccessor).Negate(), nil
	}
	return res, nil
}

func (e *EqualityExpression) evaluateInternally(ctx hipathsys.ContextAccessor, node interface{}, loop hipathsys.Looper) (interface{}, error) {
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
		if e.equivalent {
			return hipathsys.BooleanOf(hipathsys.ModelEquivalent(
				ctx.ModelAdapter(), left, right)), nil
		} else {
			return nil, nil
		}
	}

	r, match := e.stringsEqual(left, right)
	if match {
		return hipathsys.BooleanOf(r), nil
	}

	if e.equivalent {
		r = hipathsys.ModelEquivalent(ctx.ModelAdapter(), left, right)
	} else {
		r = hipathsys.ModelEqual(ctx.ModelAdapter(), left, right)
		if !r && temporalPrecisionNotEqual(left, right) {
			return nil, nil
		}
	}
	return hipathsys.BooleanOf(r), nil
}

func (e *EqualityExpression) stringsEqual(n1 interface{}, n2 interface{}) (equal bool, match bool) {
	var ok bool
	var s1, s2 hipathsys.Stringifier
	if s1, ok = n1.(hipathsys.Stringifier); !ok {
		return
	}
	if s2, ok = n2.(hipathsys.Stringifier); !ok {
		return
	}

	if s1.DataType() != hipathsys.StringDataType && s2.DataType() != hipathsys.StringDataType {
		return
	}

	match = true
	if e.equivalent {
		equal = hipathsys.NormalizedStringEqual(s1.String(), s2.String())
	} else {
		equal = s1.String() == s2.String()
	}
	return
}

func temporalPrecisionNotEqual(n1 interface{}, n2 interface{}) bool {
	var ok bool
	var t1, t2 hipathsys.TemporalAccessor
	if t1, ok = n1.(hipathsys.TemporalAccessor); !ok {
		return false
	}
	if t2, ok = n2.(hipathsys.TemporalAccessor); !ok {
		return false
	}
	return !hipathsys.TemporalPrecisionEqual(t1, t2)
}
