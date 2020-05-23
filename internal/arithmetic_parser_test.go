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
	"github.com/volsch/gohipath/internal/expression"
	"github.com/volsch/gohipath/internal/test"
	"github.com/volsch/gohipath/pathsys"
	"testing"
)

func TestParseAdditionExpression(t *testing.T) {
	res, errorItemCollection := testParse("10 + 14")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.ArithmeticExpression)(nil), res) {
		ctx := test.NewTestContext(t)
		a, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.IntegerAccessor)(nil), a) {
			assert.Equal(t, int32(24), a.(pathsys.IntegerAccessor).Int())
		}
	}
}

func TestParseSubtractionExpression(t *testing.T) {
	res, errorItemCollection := testParse("14 - 8")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.ArithmeticExpression)(nil), res) {
		ctx := test.NewTestContext(t)
		a, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.IntegerAccessor)(nil), a) {
			assert.Equal(t, int32(6), a.(pathsys.IntegerAccessor).Int())
		}
	}
}

func TestParseMultiplicationExpression(t *testing.T) {
	res, errorItemCollection := testParse("14 * 8")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.ArithmeticExpression)(nil), res) {
		ctx := test.NewTestContext(t)
		a, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.IntegerAccessor)(nil), a) {
			assert.Equal(t, int32(112), a.(pathsys.IntegerAccessor).Int())
		}
	}
}

func TestParseDivisionExpression(t *testing.T) {
	res, errorItemCollection := testParse("14 / 8")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.ArithmeticExpression)(nil), res) {
		ctx := test.NewTestContext(t)
		a, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.DecimalAccessor)(nil), a) {
			assert.Equal(t, 1.75, a.(pathsys.DecimalAccessor).Float64())
		}
	}
}

func TestParseDivExpression(t *testing.T) {
	res, errorItemCollection := testParse("18 div 8")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.ArithmeticExpression)(nil), res) {
		ctx := test.NewTestContext(t)
		a, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.IntegerAccessor)(nil), a) {
			assert.Equal(t, int32(2), a.(pathsys.IntegerAccessor).Int())
		}
	}
}

func TestParseModExpression(t *testing.T) {
	res, errorItemCollection := testParse("19 mod 8")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.ArithmeticExpression)(nil), res) {
		ctx := test.NewTestContext(t)
		a, err := res.(pathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*pathsys.IntegerAccessor)(nil), a) {
			assert.Equal(t, 3.0, a.(pathsys.IntegerAccessor).Float64())
		}
	}
}
