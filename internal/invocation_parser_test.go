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

func TestParseAggregateTotal(t *testing.T) {
	res, errorItemCollection := testParse("(10 | 14 | 3).aggregate($total + $this, 5)")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.InvocationExpression)(nil), res) {
		ctx := test.NewTestContext(t)
		res, err := res.(hipathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*hipathsys.NumberAccessor)(nil), res) {
			assert.Equal(t, 32.0, res.(hipathsys.NumberAccessor).Float64())
		}
	}
}

func TestParseAggregateIndex(t *testing.T) {
	res, errorItemCollection := testParse("(10 | 14 | 3).aggregate($index + $this + $total, 1)")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.InvocationExpression)(nil), res) {
		ctx := test.NewTestContext(t)
		res, err := res.(hipathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*hipathsys.NumberAccessor)(nil), res) {
			assert.Equal(t, 31.0, res.(hipathsys.NumberAccessor).Float64())
		}
	}
}

func TestParseMemberInvocation(t *testing.T) {
	model := make(map[string]interface{})
	model["x1"] = hipathsys.NewString("test")

	res, errorItemCollection := testParse("x1")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.InvocationTerm)(nil), res) {
		ctx := test.NewTestContext(t)
		res, err := res.(hipathsys.Evaluator).Evaluate(ctx, model, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*hipathsys.StringAccessor)(nil), res) {
			assert.Equal(t, "test", res.(hipathsys.StringAccessor).String())
		}
	}
}

func TestParseAsInvocation(t *testing.T) {
	res, errorItemCollection := testParse("'my test'.as(System.String)")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.InvocationExpression)(nil), res) {
		ctx := test.NewTestContextWithNode(t, hipathsys.NewString("test"))
		res, err := res.(hipathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*hipathsys.StringAccessor)(nil), res) {
			assert.Equal(t, hipathsys.NewString("my test"), res)
		}
	}
}

func TestParseIsInvocation(t *testing.T) {
	res, errorItemCollection := testParse("'my test'.is(System.String)")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.InvocationExpression)(nil), res) {
		ctx := test.NewTestContextWithNode(t, hipathsys.NewString("test"))
		res, err := res.(hipathsys.Evaluator).Evaluate(ctx, nil, nil)
		assert.NoError(t, err, "no evaluation error expected")
		if assert.Implements(t, (*hipathsys.BooleanAccessor)(nil), res) {
			assert.Equal(t, hipathsys.True, res)
		}
	}
}
