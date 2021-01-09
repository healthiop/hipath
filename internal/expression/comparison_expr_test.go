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

func TestComparisonExpressionLessOrEqualCol(t *testing.T) {
	ctx := test.NewTestContext(t)
	c1 := ctx.NewCollection()
	c1.MustAdd(hipathsys.NewString("test1"))
	c2 := ctx.NewCollection()
	c2.MustAdd(hipathsys.NewString("test7"))

	e := NewComparisonExpression(newTestExpression(c1), LessOrEqualThanOp, newTestExpression(c2))
	node, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.BooleanAccessor)(nil), node) {
		assert.Equal(t, true, node.(hipathsys.BooleanAccessor).Bool())
	}
}

func TestComparisonExpressionLessOrEqualLess(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewComparisonExpression(NewRawStringLiteral("test1"),
		LessOrEqualThanOp, NewRawStringLiteral("test7"))
	node, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.BooleanAccessor)(nil), node) {
		assert.Equal(t, true, node.(hipathsys.BooleanAccessor).Bool())
	}
}

func TestComparisonExpressionLessOrEqualEqual(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewComparisonExpression(NewRawStringLiteral("test7"),
		LessOrEqualThanOp, NewRawStringLiteral("test7"))
	node, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.BooleanAccessor)(nil), node) {
		assert.Equal(t, true, node.(hipathsys.BooleanAccessor).Bool())
	}
}

func TestComparisonExpressionLessOrEqualNot(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewComparisonExpression(NewRawStringLiteral("test7"),
		LessOrEqualThanOp, NewRawStringLiteral("test1"))
	node, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.BooleanAccessor)(nil), node) {
		assert.Equal(t, false, node.(hipathsys.BooleanAccessor).Bool())
	}
}

func TestComparisonExpressionLess(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewComparisonExpression(NewRawStringLiteral("test1"),
		LessThanOp, NewRawStringLiteral("test7"))
	node, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.BooleanAccessor)(nil), node) {
		assert.Equal(t, true, node.(hipathsys.BooleanAccessor).Bool())
	}
}

func TestComparisonExpressionLessNot(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewComparisonExpression(NewRawStringLiteral("test7"),
		LessThanOp, NewRawStringLiteral("test7"))
	node, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.BooleanAccessor)(nil), node) {
		assert.Equal(t, false, node.(hipathsys.BooleanAccessor).Bool())
	}
}

func TestComparisonExpressionGreater(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewComparisonExpression(NewRawStringLiteral("test7"),
		GreaterThanOp, NewRawStringLiteral("test1"))
	node, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.BooleanAccessor)(nil), node) {
		assert.Equal(t, true, node.(hipathsys.BooleanAccessor).Bool())
	}
}

func TestComparisonExpressionGreaterNot(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewComparisonExpression(NewRawStringLiteral("test7"),
		GreaterThanOp, NewRawStringLiteral("test7"))
	node, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.BooleanAccessor)(nil), node) {
		assert.Equal(t, false, node.(hipathsys.BooleanAccessor).Bool())
	}
}

func TestComparisonExpressionGreaterOrEqualLess(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewComparisonExpression(NewRawStringLiteral("test7"),
		GreaterOrEqualThanOp, NewRawStringLiteral("test1"))
	node, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.BooleanAccessor)(nil), node) {
		assert.Equal(t, true, node.(hipathsys.BooleanAccessor).Bool())
	}
}

func TestComparisonExpressionGreaterOrEqualEqual(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewComparisonExpression(NewRawStringLiteral("test7"),
		GreaterOrEqualThanOp, NewRawStringLiteral("test7"))
	node, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.BooleanAccessor)(nil), node) {
		assert.Equal(t, true, node.(hipathsys.BooleanAccessor).Bool())
	}
}

func TestComparisonExpressionGreaterOrEqualNot(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewComparisonExpression(NewRawStringLiteral("test1"),
		GreaterOrEqualThanOp, NewRawStringLiteral("test7"))
	node, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.BooleanAccessor)(nil), node) {
		assert.Equal(t, false, node.(hipathsys.BooleanAccessor).Bool())
	}
}

func TestComparisonExpressionLeftError(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewComparisonExpression(newTestErrorExpression(),
		GreaterOrEqualThanOp, NewRawStringLiteral("test7"))
	node, err := e.Evaluate(ctx, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, node, "no result expected")
}

func TestComparisonExpressionRightError(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewComparisonExpression(NewRawStringLiteral("test7"),
		GreaterOrEqualThanOp, newTestErrorExpression())
	node, err := e.Evaluate(ctx, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, node, "no result expected")
}

func TestComparisonExpressionLeftNil(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewComparisonExpression(NewEmptyLiteral(),
		GreaterOrEqualThanOp, NewRawStringLiteral("test7"))
	node, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, node, "empty result expected")
}

func TestComparisonExpressionRightNil(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewComparisonExpression(NewRawStringLiteral("test7"),
		GreaterOrEqualThanOp, NewEmptyLiteral())
	node, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, node, "empty result expected")
}

func TestComparisonExpressionLeftNonCmp(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewComparisonExpression(newTestExpression(test.NewTestModelNode(10, false)),
		GreaterOrEqualThanOp, NewRawStringLiteral("test7"))
	node, err := e.Evaluate(ctx, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, node, "no result expected")
}

func TestComparisonExpressionRightNonCmp(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewComparisonExpression(NewRawStringLiteral("test7"),
		GreaterOrEqualThanOp, newTestExpression(test.NewTestModelNode(10, false)))
	node, err := e.Evaluate(ctx, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, node, "no result expected")
}

func TestComparisonExpressionInconvertible(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewComparisonExpression(NewRawStringLiteral("test7"),
		GreaterOrEqualThanOp, NewNumberLiteralInt(10))
	node, err := e.Evaluate(ctx, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, node, "no result expected")
}

func TestComparisonExpressionEmpty(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewComparisonExpression(
		newTestExpression(hipathsys.NewDateYMDWithPrecision(2020, 7, 1, hipathsys.DayDatePrecision)),
		GreaterOrEqualThanOp, newTestExpression(hipathsys.NewDateYMDWithPrecision(2020, 7, 1, hipathsys.MonthDatePrecision)))
	node, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, node, "empty result expected")
}

func TestComparisonExpressionInvalidOp(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewComparisonExpression(NewRawStringLiteral("test1"),
		0, NewRawStringLiteral("test7"))
	assert.Panics(t, func() { _, _ = e.Evaluate(ctx, nil, nil) })
}
