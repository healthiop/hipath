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

func TestEmptyPathFuncNil(t *testing.T) {
	f := newEmptyFunction()
	res, err := f.Execute(nil, nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewBoolean(true), res)
}

func TestEmptyPathFuncValue(t *testing.T) {
	f := newEmptyFunction()
	res, err := f.Execute(nil, pathsys.NewString("test"), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewBoolean(false), res)
}

func TestEmptyPathFuncEmptyCollection(t *testing.T) {
	ctx := test.NewTestContext(t)
	f := newEmptyFunction()
	res, err := f.Execute(nil, ctx.NewCollection(), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewBoolean(true), res)
}

func TestEmptyPathFuncCollection(t *testing.T) {
	ctx := test.NewTestContext(t)
	c := ctx.NewCollection()
	c.Add(pathsys.NewString(""))
	f := newEmptyFunction()
	res, err := f.Execute(nil, c, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewBoolean(false), res)
}

func TestExistsPathFuncNil(t *testing.T) {
	f := newExistsFunction()
	res, err := f.Execute(nil, nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewBoolean(false), res)
}

func TestExistsPathFuncValueNoFilter(t *testing.T) {
	f := newExistsFunction()
	res, err := f.Execute(nil, pathsys.NewString("test"), nil, pathsys.NewLoop(nil))
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewBoolean(true), res)
}

func TestExistsPathFuncColNoFilter(t *testing.T) {
	ctx := test.NewTestContext(t)
	c := ctx.NewCollection()
	c.Add(pathsys.NewString("test"))
	f := newExistsFunction()
	res, err := f.Execute(nil, c, nil, pathsys.NewLoop(nil))
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewBoolean(true), res)
}

func TestExistsPathFuncEmptyColNoFilter(t *testing.T) {
	ctx := test.NewTestContext(t)
	c := ctx.NewCollection()
	f := newExistsFunction()
	res, err := f.Execute(nil, c, nil, pathsys.NewLoop(nil))
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewBoolean(false), res)
}

func TestExistsPathFuncValueNonMatchingFilter(t *testing.T) {
	ctx := test.NewTestContext(t)
	loopEvaluator := NewEqualityExpression(false, false, NewThisInvocation(), NewRawStringLiteral("Test"))

	f := newExistsFunction()
	res, err := f.Execute(ctx, pathsys.NewString("test"), []interface{}{}, pathsys.NewLoop(loopEvaluator))
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewBoolean(false), res)
}

func TestExistsPathFuncValueMatchingFilter(t *testing.T) {
	ctx := test.NewTestContext(t)
	loopEvaluator := NewEqualityExpression(false, false, NewThisInvocation(), NewRawStringLiteral("test"))

	f := newExistsFunction()
	res, err := f.Execute(ctx, pathsys.NewString("test"), []interface{}{}, pathsys.NewLoop(loopEvaluator))
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewBoolean(true), res)
}

func TestExistsPathFuncColNonMatchingFilter(t *testing.T) {
	ctx := test.NewTestContext(t)
	c := ctx.NewCollection()
	c.Add(pathsys.NewString("test"))
	c.Add(pathsys.NewString("tesT"))
	loopEvaluator := NewEqualityExpression(false, false, NewThisInvocation(), NewRawStringLiteral("Test"))

	f := newExistsFunction()
	res, err := f.Execute(ctx, c, []interface{}{}, pathsys.NewLoop(loopEvaluator))
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewBoolean(false), res)
}

func TestExistsPathFuncColMatchingFilter(t *testing.T) {
	ctx := test.NewTestContext(t)
	c := ctx.NewCollection()
	c.Add(pathsys.NewString("test"))
	c.Add(pathsys.NewString("Test"))
	loopEvaluator := NewEqualityExpression(false, false, NewThisInvocation(), NewRawStringLiteral("Test"))

	f := newExistsFunction()
	res, err := f.Execute(ctx, c, []interface{}{}, pathsys.NewLoop(loopEvaluator))
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewBoolean(true), res)
}

func TestExistsPathFuncValueNilFilter(t *testing.T) {
	ctx := test.NewTestContext(t)
	loopEvaluator := NewEmptyLiteral()

	f := newExistsFunction()
	res, err := f.Execute(ctx, pathsys.NewString("test"), []interface{}{}, pathsys.NewLoop(loopEvaluator))
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewBoolean(false), res)
}

func TestExistsPathFuncValueInvalidFilter(t *testing.T) {
	ctx := test.NewTestContext(t)
	loopEvaluator := NewRawStringLiteral("test")

	f := newExistsFunction()
	res, err := f.Execute(ctx, pathsys.NewString("test"), []interface{}{}, pathsys.NewLoop(loopEvaluator))
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no res expected")
}

func TestExistsPathFuncValueErrorFilter(t *testing.T) {
	ctx := test.NewTestContext(t)
	loopEvaluator := newTestErrorExpression()

	f := newExistsFunction()
	res, err := f.Execute(ctx, pathsys.NewString("test"), []interface{}{}, pathsys.NewLoop(loopEvaluator))
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no res expected")
}
