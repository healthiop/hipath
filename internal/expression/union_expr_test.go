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

func TestUnionExpressionLiteral(t *testing.T) {
	e := NewUnionExpression(ParseStringLiteral("test1"), ParseStringLiteral("test2"))
	accessor, err := e.Evaluate(nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*datatype.CollectionAccessor)(nil), accessor) {
		result := accessor.(datatype.CollectionAccessor)
		if assert.Equal(t, 2, result.Count()) {
			assert.Equal(t, datatype.NewString("test1"), result.Get(0))
			assert.Equal(t, datatype.NewString("test2"), result.Get(1))
		}
	}
}

func TestUnionExpressionCollection(t *testing.T) {
	c1 := datatype.NewCollectionUndefined()
	c1.Add(datatype.NewPositiveInt(10))
	c1.Add(datatype.NewPositiveInt(11))
	c1.Add(datatype.NewPositiveInt(14))

	c2 := datatype.NewCollectionUndefined()
	c2.Add(datatype.NewUnsignedInt(11))
	c2.Add(datatype.NewUnsignedInt(12))

	e := NewUnionExpression(newTestExpression(c1), newTestExpression(c2))
	accessor, err := e.Evaluate(nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*datatype.CollectionAccessor)(nil), accessor) {
		result := accessor.(datatype.CollectionAccessor)
		if assert.Equal(t, 4, result.Count()) {
			assert.Equal(t, datatype.NewPositiveInt(10), result.Get(0))
			assert.Equal(t, datatype.NewPositiveInt(11), result.Get(1))
			assert.Equal(t, datatype.NewPositiveInt(14), result.Get(2))
			assert.Equal(t, datatype.NewUnsignedInt(12), result.Get(3))
		}
		assert.Equal(t, "FHIR.integer", result.ItemTypeInfo().String())
	}
}

func TestUnionExpressionCollectionEmpty(t *testing.T) {
	c1 := datatype.NewCollectionUndefined()

	e := NewUnionExpression(newTestExpression(c1), newTestExpression(nil))
	accessor, err := e.Evaluate(nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, accessor, "empty result expected")
}

func TestUnionExpressionLeftNil(t *testing.T) {
	e := NewUnionExpression(NewEmptyLiteral(), ParseStringLiteral("test"))
	accessor, err := e.Evaluate(nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*datatype.CollectionAccessor)(nil), accessor) {
		result := accessor.(datatype.CollectionAccessor)
		if assert.Equal(t, 1, result.Count()) {
			assert.Equal(t, datatype.NewString("test"), result.Get(0))
		}
	}
}

func TestUnionExpressionRightNil(t *testing.T) {
	e := NewUnionExpression(ParseStringLiteral("test"), NewEmptyLiteral())
	accessor, err := e.Evaluate(nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*datatype.CollectionAccessor)(nil), accessor) {
		result := accessor.(datatype.CollectionAccessor)
		if assert.Equal(t, 1, result.Count()) {
			assert.Equal(t, datatype.NewString("test"), result.Get(0))
		}
	}
}

func TestUnionExpressionBothNil(t *testing.T) {
	e := NewUnionExpression(NewEmptyLiteral(), NewEmptyLiteral())
	accessor, err := e.Evaluate(nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, accessor, "empty collection expected")
}

func TestUnionExpressionLeftError(t *testing.T) {
	ctx := NewEvalContext(datatype.NewString("rootObj"), context.NewContext())
	e := NewUnionExpression(ParseExtConstantTerm("test"), ParseStringLiteral("test"))
	accessor, err := e.Evaluate(ctx, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, accessor, "empty collection expected")
}

func TestUnionExpressionRightError(t *testing.T) {
	ctx := NewEvalContext(datatype.NewString("rootObj"), context.NewContext())
	e := NewUnionExpression(ParseStringLiteral("test"), ParseExtConstantTerm("test"))
	accessor, err := e.Evaluate(ctx, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, accessor, "empty collection expected")
}
