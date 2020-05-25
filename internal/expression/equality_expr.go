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
	"github.com/volsch/gohipath/pathsys"
)

type EqualityExpression struct {
	not        bool
	equivalent bool
	evalLeft   pathsys.Evaluator
	evalRight  pathsys.Evaluator
}

func NewEqualityExpression(not bool, equivalent bool, evalLeft pathsys.Evaluator, evalRight pathsys.Evaluator) *EqualityExpression {
	return &EqualityExpression{not, equivalent, evalLeft, evalRight}
}

func (e *EqualityExpression) Evaluate(ctx pathsys.ContextAccessor, node interface{}, loop pathsys.Looper) (interface{}, error) {
	res, err := e.evaluateInternally(ctx, node, loop)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}
	if e.not {
		return res.(pathsys.BooleanAccessor).Negate(), nil
	}
	return res, nil
}

func (e *EqualityExpression) evaluateInternally(ctx pathsys.ContextAccessor, node interface{}, loop pathsys.Looper) (interface{}, error) {
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
			return pathsys.BooleanOf(pathsys.ModelEquivalent(
				ctx.ModelAdapter(), left, right)), nil
		} else {
			return nil, nil
		}
	}

	r, match := e.stringsEqual(left, right)
	if match {
		return pathsys.BooleanOf(r), nil
	}

	if e.equivalent {
		r = pathsys.ModelEquivalent(ctx.ModelAdapter(), left, right)
	} else {
		r = pathsys.ModelEqual(ctx.ModelAdapter(), left, right)
		if !r && temporalPrecisionNotEqual(left, right) {
			return nil, nil
		}
	}
	return pathsys.BooleanOf(r), nil
}

func (e *EqualityExpression) stringsEqual(n1 interface{}, n2 interface{}) (equal bool, match bool) {
	var ok bool
	var s1, s2 pathsys.Stringifier
	if s1, ok = n1.(pathsys.Stringifier); !ok {
		return
	}
	if s2, ok = n2.(pathsys.Stringifier); !ok {
		return
	}

	if s1.DataType() != pathsys.StringDataType && s2.DataType() != pathsys.StringDataType {
		return
	}

	match = true
	if e.equivalent {
		equal = pathsys.NormalizedStringEqual(s1.String(), s2.String())
	} else {
		equal = s1.String() == s2.String()
	}
	return
}

func temporalPrecisionNotEqual(n1 interface{}, n2 interface{}) bool {
	var ok bool
	var t1, t2 pathsys.TemporalAccessor
	if t1, ok = n1.(pathsys.TemporalAccessor); !ok {
		return false
	}
	if t2, ok = n2.(pathsys.TemporalAccessor); !ok {
		return false
	}
	return !pathsys.TemporalPrecisionEqual(t1, t2)
}
