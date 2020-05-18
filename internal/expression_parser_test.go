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

package internal

import (
	"github.com/stretchr/testify/assert"
	"github.com/volsch/gohimodel/datatype"
	"github.com/volsch/gohipath/context"
	"github.com/volsch/gohipath/internal/expression"
	"testing"
)

func TestParseNegatorExpression(t *testing.T) {
	result, errorItemCollection := testParse("-123.45")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.NegatorExpression)(nil), result) {
		ctx := expression.NewEvalContext(datatype.NewString("root"), context.NewContext())
		a, err := result.(expression.Evaluator).Evaluate(ctx, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*datatype.DecimalAccessor)(nil), a) {
			assert.Equal(t, -123.45, a.(datatype.DecimalAccessor).Float64())
		}
	}
}

func TestParseNegatorExpressionPos(t *testing.T) {
	result, errorItemCollection := testParse("+123.45")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.NumberLiteral)(nil), result) {
		ctx := expression.NewEvalContext(datatype.NewString("root"), context.NewContext())
		a, err := result.(expression.Evaluator).Evaluate(ctx, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*datatype.DecimalAccessor)(nil), a) {
			assert.Equal(t, 123.45, a.(datatype.DecimalAccessor).Float64())
		}
	}
}

func TestParseNegatorExpressionError(t *testing.T) {
	result, errorItemCollection := testParse("-'abc'")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.NegatorExpression)(nil), result) {
		ctx := expression.NewEvalContext(datatype.NewString("root"), context.NewContext())
		a, err := result.(expression.Evaluator).Evaluate(ctx, nil)
		assert.Error(t, err, "evaluation error expected")
		assert.Nil(t, a, "no accessor expected")
	}
}

func TestParseEqualityExpressionEqual(t *testing.T) {
	result, errorItemCollection := testParse("123.45=123.45")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.EqualityExpression)(nil), result) {
		a, err := result.(expression.Evaluator).Evaluate(nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*datatype.BooleanAccessor)(nil), a) {
			assert.Equal(t, true, a.(datatype.BooleanAccessor).Bool())
		}
	}
}

func TestParseEqualityExpressionEqualNot(t *testing.T) {
	result, errorItemCollection := testParse("123.45=123.4")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.EqualityExpression)(nil), result) {
		a, err := result.(expression.Evaluator).Evaluate(nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*datatype.BooleanAccessor)(nil), a) {
			assert.Equal(t, false, a.(datatype.BooleanAccessor).Bool())
		}
	}
}

func TestParseEqualityExpressionNotEqual(t *testing.T) {
	result, errorItemCollection := testParse("123.45!=123.4")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.EqualityExpression)(nil), result) {
		a, err := result.(expression.Evaluator).Evaluate(nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*datatype.BooleanAccessor)(nil), a) {
			assert.Equal(t, true, a.(datatype.BooleanAccessor).Bool())
		}
	}
}

func TestParseEqualityExpressionNotEqualNot(t *testing.T) {
	result, errorItemCollection := testParse("123.45!=123.45")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.EqualityExpression)(nil), result) {
		a, err := result.(expression.Evaluator).Evaluate(nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*datatype.BooleanAccessor)(nil), a) {
			assert.Equal(t, false, a.(datatype.BooleanAccessor).Bool())
		}
	}
}

func TestParseEqualityExpressionEquivalent(t *testing.T) {
	result, errorItemCollection := testParse("123.45~123.4")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.EqualityExpression)(nil), result) {
		a, err := result.(expression.Evaluator).Evaluate(nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*datatype.BooleanAccessor)(nil), a) {
			assert.Equal(t, true, a.(datatype.BooleanAccessor).Bool())
		}
	}
}

func TestParseEqualityExpressionEquivalentNot(t *testing.T) {
	result, errorItemCollection := testParse("123.45~123.46")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.EqualityExpression)(nil), result) {
		a, err := result.(expression.Evaluator).Evaluate(nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*datatype.BooleanAccessor)(nil), a) {
			assert.Equal(t, false, a.(datatype.BooleanAccessor).Bool())
		}
	}
}

func TestParseEqualityExpressionNotEquivalent(t *testing.T) {
	result, errorItemCollection := testParse("123.45!~123.46")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.EqualityExpression)(nil), result) {
		a, err := result.(expression.Evaluator).Evaluate(nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*datatype.BooleanAccessor)(nil), a) {
			assert.Equal(t, true, a.(datatype.BooleanAccessor).Bool())
		}
	}
}

func TestParseEqualityExpressionNotEquivalentNot(t *testing.T) {
	result, errorItemCollection := testParse("123.45!~123.4")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.EqualityExpression)(nil), result) {
		a, err := result.(expression.Evaluator).Evaluate(nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*datatype.BooleanAccessor)(nil), a) {
			assert.Equal(t, false, a.(datatype.BooleanAccessor).Bool())
		}
	}
}

func TestParseUnionExpression(t *testing.T) {
	result, errorItemCollection := testParse("10 | 12 | 11 | 10")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.UnionExpression)(nil), result) {
		a, err := result.(expression.Evaluator).Evaluate(nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*datatype.CollectionAccessor)(nil), a) {
			c := a.(datatype.CollectionAccessor)
			if assert.Equal(t, 3, c.Count()) {
				assert.Equal(t, datatype.NewInteger(10), c.Get(0))
				assert.Equal(t, datatype.NewInteger(12), c.Get(1))
				assert.Equal(t, datatype.NewInteger(11), c.Get(2))
			}
			assert.Equal(t, "FHIR.integer", c.ItemTypeInfo().String())
		}
	}
}
