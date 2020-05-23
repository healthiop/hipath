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
	"github.com/volsch/gohipath/internal/test"
	"github.com/volsch/gohipath/pathsys"
	"testing"
)

func TestEqualityExpressionEqual(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewEqualityExpression(false, false,
		ParseStringLiteral("test"), ParseStringLiteral("test"))
	accessor, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), accessor) {
		assert.Equal(t, true, accessor.(pathsys.BooleanAccessor).Bool())
	}
}

func TestEqualityExpressionEqualNot(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewEqualityExpression(false, false,
		ParseStringLiteral("test"), ParseStringLiteral("TEST"))
	accessor, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), accessor) {
		assert.Equal(t, false, accessor.(pathsys.BooleanAccessor).Bool())
	}
}

func TestEqualityExpressionLeftNoStringifier(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewEqualityExpression(false, false,
		newTestExpression(test.NewTestModelNode(10, false)),
		ParseStringLiteral("test"))
	accessor, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), accessor) {
		assert.Equal(t, false, accessor.(pathsys.BooleanAccessor).Bool())
	}
}

func TestEqualityExpressionRightStringifier(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewEqualityExpression(false, false,
		ParseStringLiteral("test"),
		newTestExpression(test.NewTestModelNode(10, false)))
	accessor, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), accessor) {
		assert.Equal(t, false, accessor.(pathsys.BooleanAccessor).Bool())
	}
}

func TestEqualityExpressionEqualModel(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewEqualityExpression(false, false,
		newTestExpression(test.NewTestModelNode(10, false)),
		newTestExpression(test.NewTestModelNode(10, false)))
	accessor, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), accessor) {
		assert.Equal(t, true, accessor.(pathsys.BooleanAccessor).Bool())
	}
}

func TestEqualityExpressionEqualModelNot(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewEqualityExpression(false, false,
		newTestExpression(test.NewTestModelNode(10, false)),
		newTestExpression(test.NewTestModelNode(11, false)))
	accessor, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), accessor) {
		assert.Equal(t, false, accessor.(pathsys.BooleanAccessor).Bool())
	}
}

func TestEqualityExpressionEquality(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewEqualityExpression(false, true,
		ParseStringLiteral("test VALUE"), ParseStringLiteral("TEST\nvalue"))
	accessor, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), accessor) {
		assert.Equal(t, true, accessor.(pathsys.BooleanAccessor).Bool())
	}
}

func TestEqualityExpressionEqualityNot(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewEqualityExpression(false, true,
		ParseStringLiteral("test VALUE"), ParseStringLiteral("TEST\nvTEST"))
	accessor, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), accessor) {
		assert.Equal(t, false, accessor.(pathsys.BooleanAccessor).Bool())
	}
}

func TestEqualityExpressionNotEqual(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewEqualityExpression(true, false,
		ParseStringLiteral("test"), ParseStringLiteral("test"))
	accessor, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), accessor) {
		assert.Equal(t, false, accessor.(pathsys.BooleanAccessor).Bool())
	}
}

func TestEqualityExpressionNotEqualNot(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewEqualityExpression(true, false,
		ParseStringLiteral("test"), ParseStringLiteral("TEST"))
	accessor, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), accessor) {
		assert.Equal(t, true, accessor.(pathsys.BooleanAccessor).Bool())
	}
}

func TestEqualityExpressionEqualStringLeft(t *testing.T) {
	n, err := ParseNumberLiteral("2020")
	if err != nil {
		t.Fatal(err)
	}
	ctx := test.NewTestContext(t)
	e := NewEqualityExpression(false, false,
		ParseStringLiteral("2020"), n)
	accessor, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), accessor) {
		assert.Equal(t, true, accessor.(pathsys.BooleanAccessor).Bool())
	}
}

func TestEqualityExpressionEqualStringRight(t *testing.T) {
	n, err := ParseNumberLiteral("2020")
	if err != nil {
		t.Fatal(err)
	}
	ctx := test.NewTestContext(t)
	e := NewEqualityExpression(false, false,
		n, ParseStringLiteral("2020"))
	accessor, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), accessor) {
		assert.Equal(t, true, accessor.(pathsys.BooleanAccessor).Bool())
	}
}

func TestEqualityExpressionEqualStringDiffers(t *testing.T) {
	n, err := ParseNumberLiteral("2021")
	if err != nil {
		t.Fatal(err)
	}
	ctx := test.NewTestContext(t)
	e := NewEqualityExpression(false, false,
		ParseStringLiteral("2020"), n)
	accessor, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), accessor) {
		assert.Equal(t, false, accessor.(pathsys.BooleanAccessor).Bool())
	}
}

func TestEqualityExpressionEqualNumber(t *testing.T) {
	n1, err := ParseNumberLiteral("1278.12")
	if err != nil {
		t.Fatal(err)
	}
	n2, err := ParseNumberLiteral("1278.12")
	if err != nil {
		t.Fatal(err)
	}
	ctx := test.NewTestContext(t)
	e := NewEqualityExpression(false, false, n1, n2)
	accessor, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), accessor) {
		assert.Equal(t, true, accessor.(pathsys.BooleanAccessor).Bool())
	}
}

func TestEqualityExpressionEqualNumberNot(t *testing.T) {
	n1, err := ParseNumberLiteral("1278.12")
	if err != nil {
		t.Fatal(err)
	}
	n2, err := ParseNumberLiteral("1278.1")
	if err != nil {
		t.Fatal(err)
	}
	ctx := test.NewTestContext(t)
	e := NewEqualityExpression(false, false, n1, n2)
	accessor, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), accessor) {
		assert.Equal(t, false, accessor.(pathsys.BooleanAccessor).Bool())
	}
}

func TestEqualityExpressionEqualTimeDiffers(t *testing.T) {
	n1, err := ParseTimeLiteral("@T12:20")
	if err != nil {
		t.Fatal(err)
	}
	n2, err := ParseTimeLiteral("@T12:21")
	if err != nil {
		t.Fatal(err)
	}
	ctx := test.NewTestContext(t)
	e := NewEqualityExpression(false, false, n1, n2)
	accessor, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), accessor) {
		assert.Equal(t, false, accessor.(pathsys.BooleanAccessor).Bool())
	}
}

func TestEqualityExpressionEqualRightNoTemporal(t *testing.T) {
	n1, err := ParseTimeLiteral("@T12:20")
	if err != nil {
		t.Fatal(err)
	}
	ctx := test.NewTestContext(t)
	e := NewEqualityExpression(false, false, n1,
		newTestExpression(test.NewTestModelNode(10, false)))
	accessor, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), accessor) {
		assert.Equal(t, false, accessor.(pathsys.BooleanAccessor).Bool())
	}
}

func TestEqualityExpressionEqualTimePrecisionDiffers(t *testing.T) {
	n1, err := ParseTimeLiteral("@T12:20:00")
	if err != nil {
		t.Fatal(err)
	}
	n2, err := ParseTimeLiteral("@T12:20")
	if err != nil {
		t.Fatal(err)
	}
	ctx := test.NewTestContext(t)
	e := NewEqualityExpression(false, false, n1, n2)
	accessor, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, accessor, "empty collection expected")
}

func TestEqualityExpressionEquivalent(t *testing.T) {
	n1, err := ParseNumberLiteral("1010.12")
	if err != nil {
		t.Fatal(err)
	}
	n2, err := ParseNumberLiteral("1010.1")
	if err != nil {
		t.Fatal(err)
	}
	ctx := test.NewTestContext(t)
	e := NewEqualityExpression(false, true, n1, n2)
	accessor, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), accessor) {
		assert.Equal(t, true, accessor.(pathsys.BooleanAccessor).Bool())
	}
}

func TestEqualityExpressionEquivalentTimePrecisionDiffers(t *testing.T) {
	n1, err := ParseTimeLiteral("@T12:21:00")
	if err != nil {
		t.Fatal(err)
	}
	n2, err := ParseTimeLiteral("@T12:20")
	if err != nil {
		t.Fatal(err)
	}
	ctx := test.NewTestContext(t)
	e := NewEqualityExpression(false, true, n1, n2)
	accessor, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), accessor) {
		assert.Equal(t, false, accessor.(pathsys.BooleanAccessor).Bool())
	}
}

func TestEqualityExpressionEqualBothNil(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewEqualityExpression(false, false, NewEmptyLiteral(), NewEmptyLiteral())
	accessor, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, accessor, "empty collection expected")
}

func TestEqualityExpressionEqualLeftNil(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewEqualityExpression(false, false, NewEmptyLiteral(), ParseStringLiteral("test"))
	accessor, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, accessor, "empty collection expected")
}

func TestEqualityExpressionEqualRightNil(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewEqualityExpression(false, false, ParseStringLiteral("test"), NewEmptyLiteral())
	accessor, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, accessor, "empty collection expected")
}

func TestEqualityExpressionEquivalentBothNil(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewEqualityExpression(false, true, NewEmptyLiteral(), NewEmptyLiteral())
	accessor, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), accessor) {
		assert.Equal(t, true, accessor.(pathsys.BooleanAccessor).Bool())
	}
}

func TestEqualityExpressionEquivalentLeftNil(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewEqualityExpression(false, true, ParseStringLiteral("test"), NewEmptyLiteral())
	accessor, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), accessor) {
		assert.Equal(t, false, accessor.(pathsys.BooleanAccessor).Bool())
	}
}

func TestEqualityExpressionEqualLeftError(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewEqualityExpression(false, false,
		newTestErrorExpression(), ParseStringLiteral("test"))
	accessor, err := e.Evaluate(ctx, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, accessor, "empty collection expected")
}

func TestEqualityExpressionEqualRightError(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewEqualityExpression(false, false,
		ParseStringLiteral("test"), newTestErrorExpression())
	accessor, err := e.Evaluate(ctx, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, accessor, "empty collection expected")
}
