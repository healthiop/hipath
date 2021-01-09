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

func TestAsTypeExpression(t *testing.T) {
	ctx := test.NewTestContext(t)
	expr, err := NewAsTypeExpression(NewRawStringLiteral("test1"), "System.String")
	if assert.NoError(t, err, "no error expected") {
		res, err := expr.Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no error expected")
		if assert.Implements(t, (*hipathsys.StringAccessor)(nil), res) {
			s := res.(hipathsys.StringAccessor)
			assert.Equal(t, "test1", s.String())
		}
	}
}

func TestAsTypeExpressionInvalid(t *testing.T) {
	expr, err := NewAsTypeExpression(NewRawStringLiteral("test1"), "System.")
	assert.Error(t, err, "error expected")
	assert.Nil(t, expr, "no result expected")
}

func TestAsTypeExpressionColSingleItem(t *testing.T) {
	ctx := test.NewTestContext(t)

	node := hipathsys.NewString("test1")
	col := ctx.NewCollection()
	col.MustAdd(node)

	expr, err := NewAsTypeExpression(newTestExpression(col), "System.String")
	if assert.NoError(t, err, "no error expected") {
		res, err := expr.Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no error expected")
		assert.Same(t, node, res)
	}
}

func TestAsTypeExpressionTypeDiffers(t *testing.T) {
	ctx := test.NewTestContext(t)
	expr, err := NewAsTypeExpression(NewRawStringLiteral("test1"), "XXX.String")
	if assert.NoError(t, err, "no error expected") {
		res, err := expr.Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no error expected")
		assert.Nil(t, res, "empty result expected")
	}
}

func TestAsTypeExpressionEmpty(t *testing.T) {
	ctx := test.NewTestContext(t)
	expr, err := NewAsTypeExpression(newTestExpression(nil), "XXX.String")
	if assert.NoError(t, err, "no error expected") {
		res, err := expr.Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no error expected")
		assert.Nil(t, res, "empty result expected")
	}
}

func TestAsTypeExpressionEvalError(t *testing.T) {
	ctx := test.NewTestContext(t)
	expr, err := NewAsTypeExpression(newTestErrorExpression(), "XXX.String")
	if assert.NoError(t, err, "no error expected") {
		res, err := expr.Evaluate(ctx, nil, nil)
		assert.Error(t, err, "error expected")
		assert.Nil(t, res, "no result expected")
	}
}

func TestAsTypeExpressionColMultipleItems(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.MustAdd(hipathsys.NewString("test1"))
	col.MustAdd(hipathsys.NewString("test2"))

	expr, err := NewAsTypeExpression(newTestExpression(col), "System.String")
	if assert.NoError(t, err, "no error expected") {
		res, err := expr.Evaluate(ctx, nil, nil)
		assert.Error(t, err, "error expected")
		assert.Nil(t, res, "no result expected")
	}
}

func TestAsTypeExpressionModelUnchanged(t *testing.T) {
	ctx := test.NewTestContext(t)

	n := test.NewTestModelNode(16.1, false)
	expr, err := NewAsTypeExpression(newTestExpression(n), "Test.decimal")
	if assert.NoError(t, err, "no error expected") {
		res, err := expr.Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no error expected")
		assert.Same(t, n, res)
	}
}

func TestAsTypeExpressionModelConverted(t *testing.T) {
	ctx := test.NewTestContext(t)

	n := test.NewTestModelNode(16.1, false)
	expr, err := NewAsTypeExpression(newTestExpression(n), "System.Number")
	if assert.NoError(t, err, "no error expected") {
		res, err := expr.Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no error expected")
		if assert.Implements(t, (*hipathsys.NumberAccessor)(nil), res) {
			assert.Equal(t, 16.1, res.(hipathsys.NumberAccessor).Float64())
		}
	}
}

func TestAsTypeExpressionModelError(t *testing.T) {
	ctx := test.NewTestContext(t)

	n := test.NewTestModelNode(16.1, false)
	expr, err := NewAsTypeExpression(newTestExpression(n), "Other.Number")
	if assert.NoError(t, err, "no error expected") {
		res, err := expr.Evaluate(ctx, nil, nil)
		assert.Error(t, err, "error expected")
		assert.Nil(t, res, "empty result expected")
	}
}

func TestIsTypeExpression(t *testing.T) {
	ctx := test.NewTestContext(t)
	expr, err := NewIsTypeExpression(NewRawStringLiteral("test1"), "System.String")
	if assert.NoError(t, err, "no error expected") {
		res, err := expr.Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no error expected")
		if assert.Implements(t, (*hipathsys.BooleanAccessor)(nil), res) {
			b := res.(hipathsys.BooleanAccessor)
			assert.Equal(t, true, b.Bool())
		}
	}
}

func TestIsTypeExpressionInvalid(t *testing.T) {
	expr, err := NewIsTypeExpression(NewRawStringLiteral("test1"), "System.")
	assert.Error(t, err, "error expected")
	assert.Nil(t, expr, "no result expected")
}

func TestIsTypeExpressionColSingleItem(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.MustAdd(hipathsys.NewString("test1"))

	expr, err := NewIsTypeExpression(newTestExpression(col), "System.String")
	if assert.NoError(t, err, "no error expected") {
		res, err := expr.Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no error expected")
		if assert.Implements(t, (*hipathsys.BooleanAccessor)(nil), res) {
			b := res.(hipathsys.BooleanAccessor)
			assert.Equal(t, true, b.Bool())
		}
	}
}

func TestIsTypeExpressionTypeDiffers(t *testing.T) {
	ctx := test.NewTestContext(t)
	expr, err := NewIsTypeExpression(NewRawStringLiteral("test1"), "XXX.String")
	if assert.NoError(t, err, "no error expected") {
		res, err := expr.Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no error expected")
		if assert.Implements(t, (*hipathsys.BooleanAccessor)(nil), res) {
			b := res.(hipathsys.BooleanAccessor)
			assert.Equal(t, false, b.Bool())
		}
	}
}

func TestIsTypeExpressionEmpty(t *testing.T) {
	ctx := test.NewTestContext(t)
	expr, err := NewIsTypeExpression(newTestExpression(nil), "XXX.String")
	if assert.NoError(t, err, "no error expected") {
		res, err := expr.Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no error expected")
		assert.Nil(t, res, "empty result expected")
	}
}

func TestIsTypeExpressionEvalError(t *testing.T) {
	ctx := test.NewTestContext(t)
	expr, err := NewIsTypeExpression(newTestErrorExpression(), "XXX.String")
	if assert.NoError(t, err, "no error expected") {
		res, err := expr.Evaluate(ctx, nil, nil)
		assert.Error(t, err, "error expected")
		assert.Nil(t, res, "no result expected")
	}
}

func TestIsTypeExpressionColMultipleItems(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.MustAdd(hipathsys.NewString("test1"))
	col.MustAdd(hipathsys.NewString("test2"))

	expr, err := NewIsTypeExpression(newTestExpression(col), "System.String")
	if assert.NoError(t, err, "no error expected") {
		res, err := expr.Evaluate(ctx, nil, nil)
		assert.Error(t, err, "error expected")
		assert.Nil(t, res, "no result expected")
	}
}
