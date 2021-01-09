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
	"github.com/healthiop/hipath/hipathsys"
	"github.com/healthiop/hipath/internal/expression"
	"github.com/healthiop/hipath/internal/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseParenthesizedBooleanLiteral(t *testing.T) {
	res, errorItemCollection := testParse("(false)")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.BooleanLiteral)(nil), res) {
		b, _ := res.(hipathsys.Evaluator).Evaluate(nil, nil, nil)
		assert.Equal(t, false, b.(hipathsys.BooleanAccessor).Bool())
	}
}

func TestParseExtConstant(t *testing.T) {
	res, errorItemCollection := testParse("%ucum")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.ExtConstantTerm)(nil), res) {
		ctx := test.NewTestContext(t)
		s, err := res.(hipathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		assert.Equal(t, hipathsys.UCUMSystemURI, s)
	}
}

func TestParseExtConstantDelimited(t *testing.T) {
	res, errorItemCollection := testParse("%`ucum`")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.ExtConstantTerm)(nil), res) {
		ctx := test.NewTestContext(t)
		s, err := res.(hipathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		assert.Equal(t, hipathsys.UCUMSystemURI, s)
	}
}

func TestParseExtConstantNotDefined(t *testing.T) {
	res, errorItemCollection := testParse("%xxx")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.ExtConstantTerm)(nil), res) {
		ctx := test.NewTestContext(t)
		s, err := res.(hipathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.Error(t, err, "evaluation error expected")
		assert.Nil(t, s, "no res expected due to error")
	}
}

func TestParseInvocationTermEmptyCollection(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("empty()")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.InvocationTerm)(nil), res) {
		col := ctx.NewCollection()
		col.MustAdd(hipathsys.NewString("test"))

		b, err := res.(hipathsys.Evaluator).Evaluate(ctx, col, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*hipathsys.BooleanAccessor)(nil), b) {
			assert.Equal(t, hipathsys.False, b)
		}
	}
}

func TestParseInvocationTermEmptyCollectionEmpty(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("empty()")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.InvocationTerm)(nil), res) {
		ctx = test.NewTestContextWithNode(t, ctx.NewCollection())

		b, err := res.(hipathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*hipathsys.BooleanAccessor)(nil), b) {
			assert.Equal(t, hipathsys.True, b)
		}
	}
}

func TestParseInvocationTermUnion(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, errorItemCollection := testParse("union(12 | 14)")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.InvocationTerm)(nil), res) {
		col := ctx.NewCollection()
		col.MustAdd(hipathsys.NewInteger(18))
		col.MustAdd(hipathsys.NewInteger(19))

		e, err := res.(hipathsys.Evaluator).Evaluate(ctx, col, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*hipathsys.CollectionAccessor)(nil), e) {
			c := e.(hipathsys.CollectionAccessor)
			if assert.Equal(t, 4, c.Count()) {
				assert.Equal(t, hipathsys.NewInteger(18), c.Get(0))
				assert.Equal(t, hipathsys.NewInteger(19), c.Get(1))
				assert.Equal(t, hipathsys.NewInteger(12), c.Get(2))
				assert.Equal(t, hipathsys.NewInteger(14), c.Get(3))
			}
		}
	}
}
