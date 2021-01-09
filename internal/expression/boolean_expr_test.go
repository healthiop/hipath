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
	"github.com/healthiop/hipath/internal/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

var booleanTests = []struct {
	name      string
	op        BooleanOp
	leftEval  hipathsys.Evaluator
	rightEval hipathsys.Evaluator
	result    hipathsys.AnyAccessor
	error     bool
}{
	{"andFalseFalse", AndOp, NewBooleanLiteral(false), NewBooleanLiteral(false), hipathsys.False, false},
	{"andFalseTrue", AndOp, NewBooleanLiteral(false), NewBooleanLiteral(true), hipathsys.False, false},
	{"andTrueFalse", AndOp, NewBooleanLiteral(true), NewBooleanLiteral(false), hipathsys.False, false},
	{"andTrueTrue", AndOp, NewBooleanLiteral(true), NewBooleanLiteral(true), hipathsys.True, false},
	{"andEmptyEmpty", AndOp, NewEmptyLiteral(), NewEmptyLiteral(), nil, false},
	{"andEmptyTrue", AndOp, NewEmptyLiteral(), NewBooleanLiteral(true), nil, false},
	{"andTrueEmpty", AndOp, NewBooleanLiteral(true), NewEmptyLiteral(), nil, false},

	{"orFalseFalse", OrOp, NewBooleanLiteral(false), NewBooleanLiteral(false), hipathsys.False, false},
	{"orFalseTrue", OrOp, NewBooleanLiteral(false), NewBooleanLiteral(true), hipathsys.True, false},
	{"orTrueFalse", OrOp, NewBooleanLiteral(true), NewBooleanLiteral(false), hipathsys.True, false},
	{"orTrueTrue", OrOp, NewBooleanLiteral(true), NewBooleanLiteral(true), hipathsys.True, false},
	{"orEmptyEmpty", OrOp, NewEmptyLiteral(), NewEmptyLiteral(), nil, false},
	{"orEmptyTrue", OrOp, NewEmptyLiteral(), NewBooleanLiteral(true), nil, false},
	{"orTrueEmpty", OrOp, NewBooleanLiteral(true), NewEmptyLiteral(), nil, false},

	{"xorFalseFalse", XOrOp, NewBooleanLiteral(false), NewBooleanLiteral(false), hipathsys.False, false},
	{"xorFalseTrue", XOrOp, NewBooleanLiteral(false), NewBooleanLiteral(true), hipathsys.True, false},
	{"xorTrueFalse", XOrOp, NewBooleanLiteral(true), NewBooleanLiteral(false), hipathsys.True, false},
	{"xorTrueTrue", XOrOp, NewBooleanLiteral(true), NewBooleanLiteral(true), hipathsys.False, false},
	{"xorEmptyEmpty", XOrOp, NewEmptyLiteral(), NewEmptyLiteral(), nil, false},
	{"xorEmptyTrue", XOrOp, NewEmptyLiteral(), NewBooleanLiteral(true), nil, false},
	{"xorTrueEmpty", XOrOp, NewBooleanLiteral(true), NewEmptyLiteral(), nil, false},

	{"impliesFalseFalse", ImpliesOp, NewBooleanLiteral(false), NewBooleanLiteral(false), hipathsys.True, false},
	{"impliesFalseTrue", ImpliesOp, NewBooleanLiteral(false), NewBooleanLiteral(true), hipathsys.True, false},
	{"impliesTrueFalse", ImpliesOp, NewBooleanLiteral(true), NewBooleanLiteral(false), hipathsys.False, false},
	{"impliesTrueTrue", ImpliesOp, NewBooleanLiteral(true), NewBooleanLiteral(true), hipathsys.True, false},
	{"impliesEmptyEmpty", ImpliesOp, NewEmptyLiteral(), NewEmptyLiteral(), nil, false},
	{"impliesEmptyTrue", ImpliesOp, NewEmptyLiteral(), NewBooleanLiteral(true), hipathsys.True, false},
	{"impliesEmptyFalse", ImpliesOp, NewEmptyLiteral(), NewBooleanLiteral(false), nil, false},
	{"impliesTrueEmpty", ImpliesOp, NewBooleanLiteral(true), NewEmptyLiteral(), nil, false},
	{"impliesFalseEmpty", ImpliesOp, NewBooleanLiteral(false), NewEmptyLiteral(), hipathsys.True, false},

	{"leftError", AndOp, newTestErrorExpression(), NewBooleanLiteral(false), nil, true},
	{"rightError", AndOp, NewBooleanLiteral(false), newTestErrorExpression(), nil, true},
}

func TestBooleanExpression(t *testing.T) {
	for _, tt := range booleanTests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := test.NewTestContext(t)

			eval := NewBooleanExpression(tt.leftEval, tt.op, tt.rightEval)
			res, err := eval.Evaluate(ctx, nil, nil)

			if tt.error {
				assert.Error(t, err, "error expected")
			} else {
				assert.NoError(t, err, "no error expected")
			}

			if tt.result == nil {
				assert.Nil(t, res, "empty result expected")
			} else {
				assert.Equal(t, tt.result, res)
			}
		})
	}
}

func TestBooleanExpressionInvalidOp(t *testing.T) {
	ctx := test.NewTestContext(t)
	evaluator := NewBooleanExpression(NewBooleanLiteral(true), 0, NewBooleanLiteral(true))

	if assert.NotNil(t, evaluator, "evaluator expected") {
		assert.Panics(t, func() { _, _ = evaluator.Evaluate(ctx, nil, nil) })
	}
}

func TestBooleanExpressionMultiColLeft(t *testing.T) {
	ctx := test.NewTestContext(t)
	operand := ctx.NewCollection()
	operand.MustAdd(hipathsys.True)
	operand.MustAdd(hipathsys.True)

	evaluator := NewBooleanExpression(newTestExpression(operand), AndOp, NewBooleanLiteral(true))
	if assert.NotNil(t, evaluator, "evaluator expected") {
		res, err := evaluator.Evaluate(ctx, nil, nil)
		assert.Error(t, err, "error expected")
		assert.Nil(t, res, "no result expected")
	}
}

func TestBooleanExpressionMultiColRight(t *testing.T) {
	ctx := test.NewTestContext(t)
	operand := ctx.NewCollection()
	operand.MustAdd(hipathsys.True)
	operand.MustAdd(hipathsys.True)

	evaluator := NewBooleanExpression(NewBooleanLiteral(true), AndOp, newTestExpression(operand))
	if assert.NotNil(t, evaluator, "evaluator expected") {
		res, err := evaluator.Evaluate(ctx, nil, nil)
		assert.Error(t, err, "error expected")
		assert.Nil(t, res, "no result expected")
	}
}

func TestBooleanExpressionNoBoolean(t *testing.T) {
	ctx := test.NewTestContext(t)

	evaluator := NewBooleanExpression(NewBooleanLiteral(true), AndOp, NewRawStringLiteral("test"))
	if assert.NotNil(t, evaluator, "evaluator expected") {
		res, err := evaluator.Evaluate(ctx, nil, nil)
		assert.Error(t, err, "error expected")
		assert.Nil(t, res, "no result expected")
	}
}

func TestBooleanExpressionSingleColString(t *testing.T) {
	ctx := test.NewTestContext(t)
	operand := ctx.NewCollection()
	operand.MustAdd(hipathsys.NewString("test"))

	evaluator := NewBooleanExpression(newTestExpression(operand), AndOp, NewBooleanLiteral(true))
	if assert.NotNil(t, evaluator, "evaluator expected") {
		res, err := evaluator.Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no error expected")
		assert.Equal(t, hipathsys.True, res)
	}
}

func TestBooleanExpressionSingleColBoolean(t *testing.T) {
	ctx := test.NewTestContext(t)
	operand := ctx.NewCollection()
	operand.MustAdd(hipathsys.False)

	evaluator := NewBooleanExpression(newTestExpression(operand), AndOp, NewBooleanLiteral(true))
	if assert.NotNil(t, evaluator, "evaluator expected") {
		res, err := evaluator.Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no error expected")
		assert.Equal(t, hipathsys.False, res)
	}
}

func TestBooleanExpressionEmptyCol(t *testing.T) {
	ctx := test.NewTestContext(t)
	operand := ctx.NewCollection()

	evaluator := NewBooleanExpression(newTestExpression(operand), AndOp, NewBooleanLiteral(true))
	if assert.NotNil(t, evaluator, "evaluator expected") {
		res, err := evaluator.Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no error expected")
		assert.Nil(t, res, "empty result expected")
	}
}

func TestBooleanExpressionColSingleNil(t *testing.T) {
	ctx := test.NewTestContext(t)
	operand := ctx.NewCollection()
	operand.MustAdd(nil)

	evaluator := NewBooleanExpression(newTestExpression(operand), AndOp, NewBooleanLiteral(true))
	if assert.NotNil(t, evaluator, "evaluator expected") {
		res, err := evaluator.Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no error expected")
		assert.Nil(t, res, "empty result expected")
	}
}
