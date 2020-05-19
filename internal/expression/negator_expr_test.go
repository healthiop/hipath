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
	"github.com/stretchr/testify/assert"
	"github.com/volsch/gohimodel/datatype"
	"github.com/volsch/gohimodel/resource"
	"github.com/volsch/gohipath/context"
	"testing"
)

func TestNegatorExpressionEvaluate(t *testing.T) {
	i, _ := ParseNumberLiteral("123.45")
	evaluator := NewNegatorExpression(i)

	accessor, err := evaluator.Evaluate(nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.NotNil(t, accessor, "accessor expected")
	if assert.Implements(t, (*datatype.DecimalAccessor)(nil), accessor) {
		assert.Equal(t, float64(-123.45), accessor.(datatype.DecimalAccessor).Float64())
	}
}

func TestNegatorExpressionEvaluateNonNegator(t *testing.T) {
	s := ParseStringLiteral("'Test'")
	evaluator := NewNegatorExpression(s)

	accessor, err := evaluator.Evaluate(nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, accessor, "no accessor expected")
}

func TestNegatorExpressionEvaluateEmpty(t *testing.T) {
	empty := NewEmptyLiteral()
	evaluator := NewNegatorExpression(empty)

	accessor, err := evaluator.Evaluate(nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, accessor, "no accessor expected")
}

func TestNegatorExpressionEvaluateError(t *testing.T) {
	ctx := NewEvalContext(resource.NewDynamicResource("Patient"), context.NewContext())
	extConstant := ParseExtConstantTerm("xxx")
	evaluator := NewNegatorExpression(extConstant)

	accessor, err := evaluator.Evaluate(ctx, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, accessor, "no accessor expected")
}
