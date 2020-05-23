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

func TestUnionExpressionLiteral(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewUnionExpression(ParseStringLiteral("test1"), ParseStringLiteral("test2"))
	res, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.CollectionAccessor)(nil), res) {
		col := res.(pathsys.CollectionAccessor)
		if assert.Equal(t, 2, col.Count()) {
			assert.Equal(t, pathsys.NewString("test1"), col.Get(0))
			assert.Equal(t, pathsys.NewString("test2"), col.Get(1))
		}
	}
}

func TestUnionExpressionCollection(t *testing.T) {
	ctx := test.NewTestContext(t)
	c1 := ctx.NewCollection()
	c1.Add(pathsys.NewInteger(10))
	c1.Add(pathsys.NewInteger(11))
	c1.Add(pathsys.NewInteger(14))

	c2 := ctx.NewCollection()
	c2.Add(pathsys.NewDecimalInt(11))
	c2.Add(pathsys.NewDecimalInt(12))

	e := NewUnionExpression(newTestExpression(c1), newTestExpression(c2))
	res, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.CollectionAccessor)(nil), res) {
		col := res.(pathsys.CollectionAccessor)
		if assert.Equal(t, 4, col.Count()) {
			assert.Equal(t, pathsys.NewInteger(10), col.Get(0))
			assert.Equal(t, pathsys.NewInteger(11), col.Get(1))
			assert.Equal(t, pathsys.NewInteger(14), col.Get(2))
			assert.Condition(t, func() bool {
				return pathsys.NewDecimalInt(12).Equal(col.Get(3))
			})
		}
		assert.Equal(t, "System.Any", col.ItemTypeInfo().String())
	}
}

func TestUnionExpressionCollectionEmpty(t *testing.T) {
	ctx := test.NewTestContext(t)
	c1 := ctx.NewCollection()

	e := NewUnionExpression(newTestExpression(c1), newTestExpression(nil))
	res, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty res expected")
}

func TestUnionExpressionLeftNil(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewUnionExpression(NewEmptyLiteral(), ParseStringLiteral("test"))
	res, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.CollectionAccessor)(nil), res) {
		col := res.(pathsys.CollectionAccessor)
		if assert.Equal(t, 1, col.Count()) {
			assert.Equal(t, pathsys.NewString("test"), col.Get(0))
		}
	}
}

func TestUnionExpressionRightNil(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewUnionExpression(ParseStringLiteral("test"), NewEmptyLiteral())
	res, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.CollectionAccessor)(nil), res) {
		col := res.(pathsys.CollectionAccessor)
		if assert.Equal(t, 1, col.Count()) {
			assert.Equal(t, pathsys.NewString("test"), col.Get(0))
		}
	}
}

func TestUnionExpressionBothNil(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewUnionExpression(NewEmptyLiteral(), NewEmptyLiteral())
	res, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestUnionExpressionLeftError(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewUnionExpression(newTestErrorExpression(), ParseStringLiteral("test"))
	res, err := e.Evaluate(ctx, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestUnionExpressionRightError(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewUnionExpression(ParseStringLiteral("test"), newTestErrorExpression())
	res, err := e.Evaluate(ctx, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}
