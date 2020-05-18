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
	"github.com/volsch/gohimodel/datatype"
)

type EqualityExpression struct {
	not        bool
	equivalent bool
	evalLeft   Evaluator
	evalRight  Evaluator
}

func NewEqualityExpression(not bool, equivalent bool, evalLeft Evaluator, evalRight Evaluator) *EqualityExpression {
	return &EqualityExpression{not, equivalent, evalLeft, evalRight}
}

func (e *EqualityExpression) Evaluate(ctx *EvalContext, curObj datatype.Accessor) (datatype.Accessor, error) {
	accessor, err := e.evaluateInternally(ctx, curObj)
	if err != nil {
		return nil, err
	}

	if accessor == nil {
		return nil, nil
	}
	if e.not {
		return accessor.(datatype.BooleanAccessor).Negate(), nil
	}
	return accessor, nil
}

func (e *EqualityExpression) evaluateInternally(ctx *EvalContext, curObj datatype.Accessor) (datatype.Accessor, error) {
	a1, err := e.evalLeft.Evaluate(ctx, curObj)
	if err != nil {
		return nil, err
	}
	a2, err := e.evalRight.Evaluate(ctx, curObj)
	if err != nil {
		return nil, err
	}

	a1, a2 = unwrapCollection(a1), unwrapCollection(a2)
	if datatype.ValueEmpty(a1) || datatype.ValueEmpty(a2) {
		if e.equivalent {
			return datatype.NewBoolean(datatype.ValueEquivalent(a1, a2)), nil
		} else {
			return nil, nil
		}
	}

	if isStringEqualityCheck(a1, a2) {
		return datatype.NewBoolean(e.primitiveStringEqual(a1, a2)), nil
	}

	var r bool
	if e.equivalent {
		r = datatype.ValueEquivalent(a1, a2)
	} else {
		r = datatype.ValueEqual(a1, a2)
		if !r && temporalTypeEqual(a1, a2) && !temporalPrecisionEqual(a1, a2) {
			return nil, nil
		}
	}
	return datatype.NewBoolean(r), nil
}

func (e *EqualityExpression) primitiveStringEqual(a1 datatype.Accessor, a2 datatype.Accessor) bool {
	p1 := a1.(datatype.PrimitiveAccessor)
	p2 := a2.(datatype.PrimitiveAccessor)
	if e.equivalent {
		return datatype.NormalizedStringEqual(p1.String(), p2.String())
	}
	return p1.String() == p2.String()
}

func isStringEqualityCheck(a1 datatype.Accessor, a2 datatype.Accessor) bool {
	return (datatype.IsString(a1) && datatype.IsPrimitive(a2)) ||
		(datatype.IsPrimitive(a1) && datatype.IsString(a2))
}

func temporalTypeEqual(a1 datatype.Accessor, a2 datatype.Accessor) bool {
	return datatype.IsTemporal(a1) && datatype.TypeEqual(a1, a2)
}

func temporalPrecisionEqual(a1 datatype.Accessor, a2 datatype.Accessor) bool {
	return a1.(datatype.TemporalAccessor).Precision() == a2.(datatype.TemporalAccessor).Precision()
}
