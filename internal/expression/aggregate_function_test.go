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
	"github.com/healthiop/hipath/hipathsys"
	"github.com/healthiop/hipath/internal/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAggregatePathFuncEvaluatorNil(t *testing.T) {
	ctx := test.NewTestContext(t)
	col := ctx.NewCol()
	col.Add(hipathsys.NewInteger(10))
	col.Add(hipathsys.NewInteger(11))
	col.Add(hipathsys.NewInteger(14))

	f := newAggregateFunction()
	res, err := f.Execute(ctx, col, []interface{}{nil}, hipathsys.NewLoop(nil))
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, err, res, "empty res expected")
}

func TestAggregatePathFuncNodeNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newAggregateFunction()
	res, err := f.Execute(ctx, nil, []interface{}{nil}, hipathsys.NewLoop(nil))
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, err, res, "empty res expected")
}

func TestAggregatePathFuncNodeError(t *testing.T) {
	ctx := test.NewTestContext(t)

	loopEvaluator := newTestErrorExpression()
	f := newAggregateFunction()
	res, err := f.Execute(ctx, test.NewTestModelErrorNode(), []interface{}{nil}, hipathsys.NewLoop(loopEvaluator))
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty result expected")
}

func TestAggregatePathFuncEvaluatorErr(t *testing.T) {
	ctx := test.NewTestContext(t)
	col := ctx.NewCol()
	col.Add(hipathsys.NewInteger(10))

	loopEvaluator := newTestErrorExpression()

	f := newAggregateFunction()
	res, err := f.Execute(ctx, col, []interface{}{nil, hipathsys.NewInteger(7)}, hipathsys.NewLoop(loopEvaluator))
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no res expected")
}

func TestAggregatePathFuncTotal(t *testing.T) {
	ctx := test.NewTestContext(t)
	col := ctx.NewCol()
	col.Add(hipathsys.NewInteger(10))
	col.Add(hipathsys.NewInteger(11))
	col.Add(hipathsys.NewInteger(14))

	loopEvaluator := NewArithmeticExpression(
		NewTotalInvocation(), hipathsys.AdditionOp, NewThisInvocation())

	f := newAggregateFunction()
	res, err := f.Execute(ctx, col, []interface{}{nil, hipathsys.NewInteger(7)}, hipathsys.NewLoop(loopEvaluator))
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.IntegerAccessor)(nil), res) {
		assert.Equal(t, hipathsys.NewInteger(42), res)
	}
}

func TestAggregatePathFuncIndex(t *testing.T) {
	ctx := test.NewTestContext(t)
	col := ctx.NewCol()
	col.Add(hipathsys.NewInteger(10))
	col.Add(hipathsys.NewInteger(11))
	col.Add(hipathsys.NewInteger(14))

	loopEvaluator := NewArithmeticExpression(
		NewIndexInvocation(), hipathsys.AdditionOp, NewThisInvocation())

	f := newAggregateFunction()
	res, err := f.Execute(ctx, col, []interface{}{nil}, hipathsys.NewLoop(loopEvaluator))
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.IntegerAccessor)(nil), res) {
		assert.Equal(t, hipathsys.NewInteger(16), res)
	}
}

func TestAggregatePathFuncColEmpty(t *testing.T) {
	ctx := test.NewTestContext(t)
	col := ctx.NewCol()

	loopEvaluator := NewArithmeticExpression(
		NewTotalInvocation(), hipathsys.AdditionOp, NewThisInvocation())

	f := newAggregateFunction()
	res, err := f.Execute(ctx, col, []interface{}{nil}, hipathsys.NewLoop(loopEvaluator))
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, err, res, "empty res expected")
}

func TestAggregatePathFuncColEmptyTotal(t *testing.T) {
	ctx := test.NewTestContext(t)
	col := ctx.NewCol()

	loopEvaluator := NewArithmeticExpression(
		NewTotalInvocation(), hipathsys.AdditionOp, NewThisInvocation())

	f := newAggregateFunction()
	res, err := f.Execute(ctx, col, []interface{}{nil, hipathsys.NewInteger(17)}, hipathsys.NewLoop(loopEvaluator))
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.IntegerAccessor)(nil), res) {
		assert.Equal(t, hipathsys.NewInteger(17), res)
	}
}
