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

func TestContainsExpression(t *testing.T) {
	ctx := test.NewTestContext(t)
	c1 := ctx.NewCol()
	c1.Add(hipathsys.NewInteger(10))
	c1.Add(hipathsys.NewInteger(11))

	e := NewContainsExpression(newTestExpression(c1), NewNumberLiteralInt(11), false)
	res, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.BooleanAccessor)(nil), res) {
		assert.Equal(t, true, res.(hipathsys.BooleanAccessor).Bool())
	}
}

func TestContainsInverseExpression(t *testing.T) {
	ctx := test.NewTestContext(t)
	c1 := ctx.NewCol()
	c1.Add(hipathsys.NewInteger(10))
	c1.Add(hipathsys.NewInteger(11))

	e := NewContainsExpression(NewNumberLiteralInt(11), newTestExpression(c1), true)
	res, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.BooleanAccessor)(nil), res) {
		assert.Equal(t, true, res.(hipathsys.BooleanAccessor).Bool())
	}
}

func TestContainsExpressionNot(t *testing.T) {
	ctx := test.NewTestContext(t)
	c1 := ctx.NewCol()
	c1.Add(hipathsys.NewInteger(10))
	c1.Add(hipathsys.NewInteger(11))

	e := NewContainsExpression(newTestExpression(c1), NewNumberLiteralInt(12), false)
	res, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.BooleanAccessor)(nil), res) {
		assert.Equal(t, false, res.(hipathsys.BooleanAccessor).Bool())
	}
}

func TestContainsExpressionColError(t *testing.T) {
	ctx := test.NewTestContext(t)

	e := NewContainsExpression(newTestErrorExpression(), NewNumberLiteralInt(12), false)
	res, err := e.Evaluate(ctx, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestContainsExpressionValError(t *testing.T) {
	ctx := test.NewTestContext(t)
	c1 := ctx.NewCol()
	c1.Add(hipathsys.NewInteger(10))
	c1.Add(hipathsys.NewInteger(11))

	e := NewContainsExpression(newTestExpression(c1), newTestErrorExpression(), false)
	res, err := e.Evaluate(ctx, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestContainsExpressionNilCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	e := NewContainsExpression(newTestExpression(nil), NewNumberLiteralInt(11), false)
	res, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.BooleanAccessor)(nil), res) {
		assert.Equal(t, false, res.(hipathsys.BooleanAccessor).Bool())
	}
}

func TestContainsExpressionNilVal(t *testing.T) {
	ctx := test.NewTestContext(t)
	c1 := ctx.NewCol()
	c1.Add(hipathsys.NewInteger(10))

	e := NewContainsExpression(newTestExpression(c1), NewEmptyLiteral(), false)
	res, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty result expected")
}

func TestContainsExpressionValCol(t *testing.T) {
	ctx := test.NewTestContext(t)
	c1 := ctx.NewCol()
	c1.Add(hipathsys.NewInteger(10))
	v := ctx.NewCol()
	v.Add(hipathsys.NewInteger(10))
	v.Add(hipathsys.NewInteger(11))

	e := NewContainsExpression(newTestExpression(c1), newTestExpression(v), false)
	res, err := e.Evaluate(ctx, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}
