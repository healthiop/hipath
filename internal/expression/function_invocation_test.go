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
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/volsch/gohipath/internal/test"
	"github.com/volsch/gohipath/pathsys"
	"testing"
)

var testLoop = pathsys.NewLoop(nil)

func TestFunctionInvocationNoArgs(t *testing.T) {
	function := &testInvocationArgsFunction{
		t:            t,
		BaseFunction: pathsys.NewBaseFunction("test", -1, 0, 0),
	}

	ctx := test.NewTestContext(t)
	e := newFunctionInvocation(function, []pathsys.Evaluator{})

	tt := newTestingType(t)
	res, err := e.Evaluate(ctx, tt, testLoop)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.CollectionAccessor)(nil), res) {
		c := res.(pathsys.CollectionAccessor)
		if assert.Equal(t, 1, c.Count()) {
			assert.Equal(t, pathsys.NewInteger(0), c.Get(0))
		}
	}
}

func TestFunctionInvocationArgs(t *testing.T) {
	function := &testInvocationArgsFunction{
		t:            t,
		BaseFunction: pathsys.NewBaseFunction("test", -1, 0, 100),
	}

	testExpression := newTestExpression(pathsys.NewString("test1"))
	ctx := test.NewTestContext(t)
	e := newFunctionInvocation(function, []pathsys.Evaluator{
		testExpression, nil, ParseStringLiteral("test2")})

	tt := newTestingType(t)
	res, err := e.Evaluate(ctx, tt, testLoop)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.CollectionAccessor)(nil), res) {
		c := res.(pathsys.CollectionAccessor)
		if assert.Equal(t, 4, c.Count()) {
			assert.Equal(t, pathsys.NewInteger(3), c.Get(0))
			assert.Equal(t, pathsys.NewString("test1"), c.Get(1))
			assert.Nil(t, c.Get(2))
			assert.Equal(t, pathsys.NewString("test2"), c.Get(3))
		}
	}
	assert.Equal(t, 1, testExpression.invocationCount)
	assert.Same(t, tt, testExpression.node)
	assert.Same(t, testLoop, testExpression.loop)
}

func TestFunctionInvocationLoop(t *testing.T) {
	function := &testInvocationLoopFunction{
		t:            t,
		BaseFunction: pathsys.NewBaseFunction("test", 0, 0, 100),
	}

	loopExpression := newTestExpression(pathsys.NewString("testLoop"))
	ctx := test.NewTestContext(t)
	e := newFunctionInvocation(function, []pathsys.Evaluator{loopExpression})

	tt := newTestingType(t)
	res, err := e.Evaluate(ctx, tt, testLoop)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewString("testLoop"), res)

	assert.Equal(t, 1, loopExpression.invocationCount)
	assert.NotSame(t, testLoop, loopExpression.loop)
}

func TestFunctionInvocationArgsError(t *testing.T) {
	function := &testInvocationArgsFunction{
		t:            t,
		BaseFunction: pathsys.NewBaseFunction("test", -1, 0, 100),
	}

	ctx := test.NewTestContext(t)
	e := newFunctionInvocation(function, []pathsys.Evaluator{
		ParseStringLiteral("test1"), ParseExtConstantTerm("xxx"), ParseStringLiteral("test2")})

	res, err := e.Evaluate(ctx, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no res expected")
}

func TestFunctionInvocationError(t *testing.T) {
	function := &testInvocationErrFunction{
		BaseFunction: pathsys.NewBaseFunction("test", -1, 0, 100),
	}

	ctx := test.NewTestContext(t)
	e := newFunctionInvocation(function, []pathsys.Evaluator{})

	res, err := e.Evaluate(ctx, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no res expected")
}

func TestLookupFunctionInvocationNotFound(t *testing.T) {
	fi, err := LookupFunctionInvocation("test", make([]pathsys.Evaluator, 0))
	assert.EqualError(t, err, "executor has not been defined: test", "error expected")
	assert.Nil(t, fi, "no executor invocation expected")
}

func TestLookupFunctionInvocationTooLessArgs(t *testing.T) {
	fi, err := LookupFunctionInvocation("union", make([]pathsys.Evaluator, 0))
	assert.EqualError(t, err, "executor union requires at least 1 parameters", "error expected")
	assert.Nil(t, fi, "no executor invocation expected")
}

func TestLookupFunctionInvocationTooManyArgs(t *testing.T) {
	fi, err := LookupFunctionInvocation("union", make([]pathsys.Evaluator, 2))
	assert.EqualError(t, err, "executor union accepts at most 1 parameters", "error expected")
	assert.Nil(t, fi, "no executor invocation expected")
}

type testInvocationArgsFunction struct {
	pathsys.BaseFunction
	t *testing.T
}

func (f *testInvocationArgsFunction) Execute(ctx pathsys.ContextAccessor, _ interface{}, args []interface{}, loop pathsys.Looper) (interface{}, error) {
	t := f.t
	assert.Same(t, testLoop, loop)

	c := ctx.NewCollection()
	c.Add(pathsys.NewInteger(int32(len(args))))
	for _, a := range args {
		c.Add(a)
	}
	return c, nil
}

type testInvocationLoopFunction struct {
	pathsys.BaseFunction
	t *testing.T
}

func (f *testInvocationLoopFunction) Execute(_ pathsys.ContextAccessor, _ interface{}, _ []interface{}, loop pathsys.Looper) (interface{}, error) {
	t := f.t
	if assert.NotNil(t, loop) && assert.NotNil(t, loop.Evaluator()) {
		res, err := loop.Evaluator().Evaluate(nil, nil, nil)
		return res, err
	}
	return nil, nil
}

type testInvocationErrFunction struct {
	pathsys.BaseFunction
}

func (f *testInvocationErrFunction) Execute(pathsys.ContextAccessor, interface{}, []interface{}, pathsys.Looper) (interface{}, error) {
	return nil, fmt.Errorf("an error occurred")
}
