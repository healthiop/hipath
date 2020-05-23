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

func TestStringConcat(t *testing.T) {
	evaluator := NewStringConcatExpression(
		NewRawStringLiteral("Test"), NewRawStringLiteral(" ABC"))

	res, err := evaluator.Evaluate(nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.NotNil(t, res, "res expected")
	if assert.Implements(t, (*pathsys.StringAccessor)(nil), res) {
		assert.Equal(t, pathsys.NewString("Test ABC"), res)
	}
}

func TestStringConcatBoolean(t *testing.T) {
	numberLiteral, err := ParseNumberLiteral("20")
	if err != nil {
		t.Fatal(err)
	}
	evaluator := NewStringConcatExpression(
		NewRawStringLiteral("Test "), numberLiteral)

	res, err := evaluator.Evaluate(nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.NotNil(t, res, "res expected")
	if assert.Implements(t, (*pathsys.StringAccessor)(nil), res) {
		assert.Equal(t, pathsys.NewString("Test 20"), res)
	}
}

func TestStringConcatBothEmpty(t *testing.T) {
	evaluator := NewStringConcatExpression(
		NewEmptyLiteral(), NewEmptyLiteral())

	res, err := evaluator.Evaluate(nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.NotNil(t, res, "res expected")
	if assert.Implements(t, (*pathsys.StringAccessor)(nil), res) {
		assert.Equal(t, pathsys.NewString(""), res)
	}
}

func TestStringConcatLeftEmpty(t *testing.T) {
	evaluator := NewStringConcatExpression(
		NewEmptyLiteral(), NewRawStringLiteral("Test"))

	res, err := evaluator.Evaluate(nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.NotNil(t, res, "res expected")
	if assert.Implements(t, (*pathsys.StringAccessor)(nil), res) {
		assert.Equal(t, pathsys.NewString("Test"), res)
	}
}

func TestStringConcatLeftError(t *testing.T) {
	evaluator := NewStringConcatExpression(
		newTestErrorExpression(), NewRawStringLiteral("Test"))

	res, err := evaluator.Evaluate(nil, nil, nil)
	assert.Error(t, err, "no error expected")
	assert.Nil(t, res, "no res expected")
}

func TestStringConcatLeftMultiCol(t *testing.T) {
	ctx := test.NewTestContext(t)
	col := ctx.NewCollection()
	col.Add(pathsys.NewString("test1"))
	col.Add(pathsys.NewString("test2"))

	evaluator := NewStringConcatExpression(
		newTestExpression(col), NewRawStringLiteral("Test"))

	res, err := evaluator.Evaluate(nil, nil, nil)
	assert.Error(t, err, "no error expected")
	assert.Nil(t, res, "no res expected")
}

func TestStringConcatLeftCol(t *testing.T) {
	ctx := test.NewTestContext(t)
	col := ctx.NewCollection()
	col.Add(pathsys.NewString("test1"))

	evaluator := NewStringConcatExpression(
		newTestExpression(col), NewRawStringLiteral("Test"))

	res, err := evaluator.Evaluate(nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.NotNil(t, res, "res expected")
	if assert.Implements(t, (*pathsys.StringAccessor)(nil), res) {
		assert.Equal(t, pathsys.NewString("test1Test"), res)
	}
}

func TestStringConcatRightEmpty(t *testing.T) {
	evaluator := NewStringConcatExpression(
		NewRawStringLiteral("Test"), NewEmptyLiteral())

	res, err := evaluator.Evaluate(nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.NotNil(t, res, "res expected")
	if assert.Implements(t, (*pathsys.StringAccessor)(nil), res) {
		assert.Equal(t, pathsys.NewString("Test"), res)
	}
}

func TestStringConcatRightError(t *testing.T) {
	evaluator := NewStringConcatExpression(
		NewRawStringLiteral("Test"), newTestErrorExpression())

	res, err := evaluator.Evaluate(nil, nil, nil)
	assert.Error(t, err, "no error expected")
	assert.Nil(t, res, "no res expected")
}

func TestStringConcatRightMultiCol(t *testing.T) {
	ctx := test.NewTestContext(t)
	col := ctx.NewCollection()
	col.Add(pathsys.NewString("test1"))
	col.Add(pathsys.NewString("test2"))

	evaluator := NewStringConcatExpression(
		NewRawStringLiteral("Test"), newTestExpression(col))

	res, err := evaluator.Evaluate(nil, nil, nil)
	assert.Error(t, err, "no error expected")
	assert.Nil(t, res, "no res expected")
}

func TestStringConcatRightCol(t *testing.T) {
	ctx := test.NewTestContext(t)
	col := ctx.NewCollection()
	col.Add(pathsys.NewString("Test1"))

	evaluator := NewStringConcatExpression(
		NewRawStringLiteral("Test"), newTestExpression(col))

	res, err := evaluator.Evaluate(nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.NotNil(t, res, "res expected")
	if assert.Implements(t, (*pathsys.StringAccessor)(nil), res) {
		assert.Equal(t, pathsys.NewString("TestTest1"), res)
	}
}
