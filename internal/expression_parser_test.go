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

package internal

import (
	"github.com/stretchr/testify/assert"
	"github.com/volsch/gohipath/internal/expression"
	"github.com/volsch/gohipath/internal/test"
	"github.com/volsch/gohipath/pathsys"
	"testing"
)

func TestParseNegatorExpression(t *testing.T) {
	res, errorItemCollection := testParse("-123.45")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.NegatorExpression)(nil), res) {
		ctx := test.NewTestContext(t)
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.DecimalAccessor)(nil), res) {
			assert.Equal(t, -123.45, res.(pathsys.DecimalAccessor).Float64())
		}
	}
}

func TestParseNegatorExpressionPos(t *testing.T) {
	res, errorItemCollection := testParse("+123.45")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.NumberLiteral)(nil), res) {
		ctx := test.NewTestContext(t)
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.DecimalAccessor)(nil), res) {
			assert.Equal(t, 123.45, res.(pathsys.DecimalAccessor).Float64())
		}
	}
}

func TestParseNegatorExpressionError(t *testing.T) {
	res, errorItemCollection := testParse("-'abc'")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.NegatorExpression)(nil), res) {
		ctx := test.NewTestContext(t)
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.Error(t, err, "evaluation error expected")
		assert.Nil(t, res, "no res expected")
	}
}

func TestParseEqualityExpressionEqual(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("123.45=123.45")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.EqualityExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, true, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}

func TestParseEqualityExpressionEqualNot(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("123.45=123.4")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.EqualityExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, false, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}

func TestParseEqualityExpressionNotEqual(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("123.45!=123.4")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.EqualityExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, true, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}

func TestParseEqualityExpressionNotEqualNot(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("123.45!=123.45")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.EqualityExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, false, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}

func TestParseEqualityExpressionEquivalent(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("123.45~123.4")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.EqualityExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, true, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}

func TestParseEqualityExpressionEquivalentNot(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("123.45~123.46")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.EqualityExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, false, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}

func TestParseEqualityExpressionNotEquivalent(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("123.45!~123.46")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.EqualityExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, true, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}

func TestParseEqualityExpressionNotEquivalentNot(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("123.45!~123.4")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.EqualityExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, false, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}

func TestParseUnionExpression(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("10 | 12 | 11 | 10")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.UnionExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.CollectionAccessor)(nil), res) {
			c := res.(pathsys.CollectionAccessor)
			if assert.Equal(t, 3, c.Count()) {
				assert.Equal(t, pathsys.NewInteger(10), c.Get(0))
				assert.Equal(t, pathsys.NewInteger(12), c.Get(1))
				assert.Equal(t, pathsys.NewInteger(11), c.Get(2))
			}
		}
	}
}

func TestParseIndexerExpression(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("(10 | 17 | 11 | 14)[1]")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.IndexerExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.NumberAccessor)(nil), res) {
			assert.Equal(t, int32(17), res.(pathsys.NumberAccessor).Int())
		}
	}
}

func TestParseInvocationExpressionUnion(t *testing.T) {
	res, errorItemCollection := testParse("(18 | 19).union(12 | 14)")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.InvocationExpression)(nil), res) {
		ctx := test.NewTestContextWithNode(t, pathsys.NewString("test"))
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.CollectionAccessor)(nil), res) {
			c := res.(pathsys.CollectionAccessor)
			if assert.Equal(t, 4, c.Count()) {
				assert.Equal(t, pathsys.NewInteger(18), c.Get(0))
				assert.Equal(t, pathsys.NewInteger(19), c.Get(1))
				assert.Equal(t, pathsys.NewInteger(12), c.Get(2))
				assert.Equal(t, pathsys.NewInteger(14), c.Get(3))
			}
		}
	}
}

func TestParseComparisonExpressionLessOrEqualLess(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("10<=10.5")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.ComparisonExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, true, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}

func TestParseComparisonExpressionLessOrEqualEqual(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("10<=10")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.ComparisonExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, true, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}

func TestParseComparisonExpressionLessOrEqualNot(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("10<=9.9")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.ComparisonExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, false, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}

func TestParseComparisonExpressionLessLess(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("10<10.5")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.ComparisonExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, true, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}

func TestParseComparisonExpressionLessEqualNot(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("10<10")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.ComparisonExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, false, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}

func TestParseComparisonExpressionLessNot(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("10<9.9")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.ComparisonExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, false, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}

func TestParseComparisonExpressionGreaterGreater(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("'test9'>'test1'")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.ComparisonExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, true, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}

func TestParseComparisonExpressionGreaterEqualNot(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("'test10'>'test10'")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.ComparisonExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, false, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}

func TestParseComparisonExpressionGreaterNot(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("'test1'>'test9'")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.ComparisonExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, false, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}

func TestParseComparisonExpressionGreaterPrEqualGreater(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("'test9'>='test1'")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.ComparisonExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, true, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}

func TestParseComparisonExpressionGreaterOrEqualEqual(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("'test10'>='test10'")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.ComparisonExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, true, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}

func TestParseComparisonExpressionGreaterOrEqualNot(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("'test1'>='test9'")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.ComparisonExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, false, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}

func TestParseComparisonExpressionPrecision(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("@2018-10-01>=@2018-09")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.ComparisonExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		assert.Nil(t, res, "empty result expected")
	}
}

func TestParseComparisonExpressionInconvertible(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("@2018-10-01>=10")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.ComparisonExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.Error(t, err, "evaluation error expected")
		assert.Nil(t, res, "no result expected")
	}
}

func TestParseMembershipExpressionContains(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("(10|12) contains 12")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.ContainsExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, true, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}

func TestParseMembershipExpressionContainsNot(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("(10|12) contains 14")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.ContainsExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, false, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}

func TestParseMembershipExpressionIn(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("12 in (10|12)")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.ContainsExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, true, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}

func TestParseMembershipExpressionInNot(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("14 in (10|12)")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.ContainsExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, false, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}

func TestParseBooleanExpressionAnd(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("true and true")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.BooleanExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, true, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}

func TestParseBooleanExpressionAndNot(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("true and false")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.BooleanExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, false, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}

func TestParseBooleanExpressionOr(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("true or false")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.BooleanExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, true, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}

func TestParseBooleanExpressionOrNot(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("false or false")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.BooleanExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, false, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}

func TestParseBooleanExpressionXOr(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("true xor false")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.BooleanExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, true, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}

func TestParseBooleanExpressionXOrNot(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("true xor true")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.BooleanExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, false, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}

func TestParseBooleanExpressionImplies(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("{} implies true")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.BooleanExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, true, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}

func TestParseBooleanExpressionImpliesNot(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("true implies false")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.BooleanExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, false, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}

func TestParseAsExpression(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("'test2' as System.String")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.AsTypeExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.StringAccessor)(nil), res) {
			assert.Equal(t, "test2", res.(pathsys.StringAccessor).String())
		}
	}
}

func TestParseAsExpressionNot(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("123 as System.String")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.AsTypeExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		assert.Nil(t, res, "empty result expected")
	}
}

func TestParseIsExpression(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("'test2' is System.String")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.IsTypeExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, true, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}

func TestParseIsExpressionNot(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("123 is System.String")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.IsTypeExpression)(nil), res) {
		res, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, false, res.(pathsys.BooleanAccessor).Bool())
		}
	}
}
