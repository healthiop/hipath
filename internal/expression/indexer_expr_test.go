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
	"github.com/volsch/gohipath/context"
	"testing"
)

func TestIndexerExpressionCollection(t *testing.T) {
	c := datatype.NewCollectionUndefined()
	c.Add(datatype.NewString("test1"))
	c.Add(datatype.NewString("test2"))
	c.Add(datatype.NewString("test3"))

	i, err := ParseNumberLiteral("1")
	if err != nil {
		t.Fatal(err)
	}

	e := NewIndexerExpression(newTestExpression(c), i)
	accessor, err := e.Evaluate(nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*datatype.StringAccessor)(nil), accessor) {
		assert.Equal(t, datatype.NewString("test2"), accessor)
	}
}

func TestIndexerExpressionCollectionIndexNeg(t *testing.T) {
	c := datatype.NewCollectionUndefined()
	c.Add(datatype.NewString("test1"))

	i, err := ParseNumberLiteral("-1")
	if err != nil {
		t.Fatal(err)
	}

	e := NewIndexerExpression(newTestExpression(c), i)
	accessor, err := e.Evaluate(nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, accessor, "empty result expected")
}

func TestIndexerExpressionCollectionInvalidIndexType(t *testing.T) {
	c := datatype.NewCollectionUndefined()
	c.Add(datatype.NewString("test1"))

	e := NewIndexerExpression(newTestExpression(c), ParseStringLiteral("0"))
	accessor, err := e.Evaluate(nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, accessor, "no result expected")
}

func TestIndexerExpressionExpressionNil(t *testing.T) {
	i, err := ParseNumberLiteral("0")
	if err != nil {
		t.Fatal(err)
	}

	e := NewIndexerExpression(newTestExpression(nil), i)
	accessor, err := e.Evaluate(nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, accessor, "empty result expected")
}

func TestIndexerExpressionIndexNil(t *testing.T) {
	c := datatype.NewCollectionUndefined()
	c.Add(datatype.NewString("test1"))

	e := NewIndexerExpression(newTestExpression(c), newTestExpression(nil))
	accessor, err := e.Evaluate(nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, accessor, "empty result expected")
}

func TestIndexerExpressionCollectionCountExceeded(t *testing.T) {
	c := datatype.NewCollectionUndefined()
	c.Add(datatype.NewString("test1"))
	c.Add(datatype.NewString("test2"))
	c.Add(datatype.NewString("test3"))

	i, err := ParseNumberLiteral("3")
	if err != nil {
		t.Fatal(err)
	}

	e := NewIndexerExpression(newTestExpression(c), i)
	accessor, err := e.Evaluate(nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, accessor, "empty result expected")
}

func TestIndexerExpressionCollectionNoCol(t *testing.T) {
	i, err := ParseNumberLiteral("0")
	if err != nil {
		t.Fatal(err)
	}

	e := NewIndexerExpression(ParseStringLiteral("test"), i)
	accessor, err := e.Evaluate(nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*datatype.StringAccessor)(nil), accessor) {
		assert.Equal(t, datatype.NewString("test"), accessor)
	}
}

func TestIndexerExpressionCollectionNoColIndexExceeded(t *testing.T) {
	i, err := ParseNumberLiteral("1")
	if err != nil {
		t.Fatal(err)
	}

	e := NewIndexerExpression(ParseStringLiteral("test"), i)
	accessor, err := e.Evaluate(nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, accessor, "empty result expected")
}

func TestIndexerExpressionExpressionError(t *testing.T) {
	i, err := ParseNumberLiteral("0")
	if err != nil {
		t.Fatal(err)
	}

	ctx := NewEvalContext(datatype.NewString("rootObj"), context.NewContext())
	e := NewIndexerExpression(ParseExtConstantTerm("test"), i)
	accessor, err := e.Evaluate(ctx, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, accessor, "empty collection expected")
}

func TestIndexerExpressionIndexError(t *testing.T) {
	ctx := NewEvalContext(datatype.NewString("rootObj"), context.NewContext())
	e := NewIndexerExpression(ParseStringLiteral("test"), ParseExtConstantTerm("test"))
	accessor, err := e.Evaluate(ctx, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, accessor, "empty collection expected")
}
