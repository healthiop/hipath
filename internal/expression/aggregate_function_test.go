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

func TestAggregatePathFuncEvaluatorNil(t *testing.T) {
	ctx := test.NewTestContext(t)
	col := ctx.NewCollection()
	col.Add(pathsys.NewInteger(10))
	col.Add(pathsys.NewInteger(11))
	col.Add(pathsys.NewInteger(14))

	f := newAggregateFunction()
	res, err := f.Execute(ctx, col, []interface{}{nil}, pathsys.NewLoop(nil))
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, err, res, "empty res expected")
}

func TestAggregatePathFuncNodeNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newAggregateFunction()
	res, err := f.Execute(ctx, nil, []interface{}{nil}, pathsys.NewLoop(nil))
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, err, res, "empty res expected")
}

func TestAggregatePathFuncEvaluatorErr(t *testing.T) {
	ctx := test.NewTestContext(t)
	col := ctx.NewCollection()
	col.Add(pathsys.NewInteger(10))

	loopEvaluator := newTestErrorExpression()

	f := newAggregateFunction()
	res, err := f.Execute(ctx, col, []interface{}{nil, pathsys.NewInteger(7)}, pathsys.NewLoop(loopEvaluator))
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no res expected")
}

func TestAggregatePathFuncTotal(t *testing.T) {
	ctx := test.NewTestContext(t)
	col := ctx.NewCollection()
	col.Add(pathsys.NewInteger(10))
	col.Add(pathsys.NewInteger(11))
	col.Add(pathsys.NewInteger(14))

	loopEvaluator := NewArithmeticExpression(
		NewTotalInvocation(), pathsys.AdditionOp, NewThisInvocation())

	f := newAggregateFunction()
	res, err := f.Execute(ctx, col, []interface{}{nil, pathsys.NewInteger(7)}, pathsys.NewLoop(loopEvaluator))
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.IntegerAccessor)(nil), res) {
		assert.Equal(t, pathsys.NewInteger(42), res)
	}
}

func TestAggregatePathFuncIndex(t *testing.T) {
	ctx := test.NewTestContext(t)
	col := ctx.NewCollection()
	col.Add(pathsys.NewInteger(10))
	col.Add(pathsys.NewInteger(11))
	col.Add(pathsys.NewInteger(14))

	loopEvaluator := NewArithmeticExpression(
		NewIndexInvocation(), pathsys.AdditionOp, NewThisInvocation())

	f := newAggregateFunction()
	res, err := f.Execute(ctx, col, []interface{}{nil}, pathsys.NewLoop(loopEvaluator))
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.IntegerAccessor)(nil), res) {
		assert.Equal(t, pathsys.NewInteger(16), res)
	}
}

func TestAggregatePathFuncColEmpty(t *testing.T) {
	ctx := test.NewTestContext(t)
	col := ctx.NewCollection()

	loopEvaluator := NewArithmeticExpression(
		NewTotalInvocation(), pathsys.AdditionOp, NewThisInvocation())

	f := newAggregateFunction()
	res, err := f.Execute(ctx, col, []interface{}{nil}, pathsys.NewLoop(loopEvaluator))
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, err, res, "empty res expected")
}

func TestAggregatePathFuncColEmptyTotal(t *testing.T) {
	ctx := test.NewTestContext(t)
	col := ctx.NewCollection()

	loopEvaluator := NewArithmeticExpression(
		NewTotalInvocation(), pathsys.AdditionOp, NewThisInvocation())

	f := newAggregateFunction()
	res, err := f.Execute(ctx, col, []interface{}{nil, pathsys.NewInteger(17)}, pathsys.NewLoop(loopEvaluator))
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.IntegerAccessor)(nil), res) {
		assert.Equal(t, pathsys.NewInteger(17), res)
	}
}
