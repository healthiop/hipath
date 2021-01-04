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

func TestEmptyPathFuncNil(t *testing.T) {
	f := newEmptyFunction()
	res, err := f.Execute(nil, nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.True, res)
}

func TestEmptyPathFuncValue(t *testing.T) {
	f := newEmptyFunction()
	res, err := f.Execute(nil, hipathsys.NewString("test"), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.False, res)
}

func TestEmptyPathFuncEmptyCollection(t *testing.T) {
	ctx := test.NewTestContext(t)
	f := newEmptyFunction()
	res, err := f.Execute(nil, ctx.NewCollection(), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.True, res)
}

func TestEmptyPathFuncCollection(t *testing.T) {
	ctx := test.NewTestContext(t)
	c := ctx.NewCollection()
	c.Add(hipathsys.NewString(""))
	f := newEmptyFunction()
	res, err := f.Execute(nil, c, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.False, res)
}

func TestExistsPathFuncNil(t *testing.T) {
	f := newExistsFunction()
	res, err := f.Execute(nil, nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.False, res)
}

func TestExistsPathFuncValueNoFilter(t *testing.T) {
	f := newExistsFunction()
	res, err := f.Execute(nil, hipathsys.NewString("test"), nil, hipathsys.NewLoop(nil))
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.True, res)
}

func TestExistsPathFuncColNoFilter(t *testing.T) {
	ctx := test.NewTestContext(t)
	c := ctx.NewCollection()
	c.Add(hipathsys.NewString("test"))
	f := newExistsFunction()
	res, err := f.Execute(nil, c, nil, hipathsys.NewLoop(nil))
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.True, res)
}

func TestExistsPathFuncEmptyColNoFilter(t *testing.T) {
	ctx := test.NewTestContext(t)
	c := ctx.NewCollection()
	f := newExistsFunction()
	res, err := f.Execute(nil, c, nil, hipathsys.NewLoop(nil))
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.False, res)
}

func TestExistsPathFuncValueNonMatchingFilter(t *testing.T) {
	ctx := test.NewTestContext(t)
	loopEvaluator := NewEqualityExpression(false, false, NewThisInvocation(), NewRawStringLiteral("Test"))

	f := newExistsFunction()
	res, err := f.Execute(ctx, hipathsys.NewString("test"), []interface{}{}, hipathsys.NewLoop(loopEvaluator))
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.False, res)
}

func TestExistsPathFuncValueMatchingFilter(t *testing.T) {
	ctx := test.NewTestContext(t)
	loopEvaluator := NewEqualityExpression(false, false, NewThisInvocation(), NewRawStringLiteral("test"))

	f := newExistsFunction()
	res, err := f.Execute(ctx, hipathsys.NewString("test"), []interface{}{}, hipathsys.NewLoop(loopEvaluator))
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.True, res)
}

func TestExistsPathFuncColNonMatchingFilter(t *testing.T) {
	ctx := test.NewTestContext(t)
	c := ctx.NewCollection()
	c.Add(hipathsys.NewString("test"))
	c.Add(hipathsys.NewString("tesT"))
	loopEvaluator := NewEqualityExpression(false, false, NewThisInvocation(), NewRawStringLiteral("Test"))

	f := newExistsFunction()
	res, err := f.Execute(ctx, c, []interface{}{}, hipathsys.NewLoop(loopEvaluator))
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.False, res)
}

func TestExistsPathFuncColMatchingFilter(t *testing.T) {
	ctx := test.NewTestContext(t)
	c := ctx.NewCollection()
	c.Add(hipathsys.NewString("test"))
	c.Add(hipathsys.NewString("Test"))
	loopEvaluator := NewEqualityExpression(false, false, NewThisInvocation(), NewRawStringLiteral("Test"))

	f := newExistsFunction()
	res, err := f.Execute(ctx, c, []interface{}{}, hipathsys.NewLoop(loopEvaluator))
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.True, res)
}

func TestExistsPathFuncValueNilFilter(t *testing.T) {
	ctx := test.NewTestContext(t)
	loopEvaluator := NewEmptyLiteral()

	f := newExistsFunction()
	res, err := f.Execute(ctx, hipathsys.NewString("test"), []interface{}{}, hipathsys.NewLoop(loopEvaluator))
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.False, res)
}

func TestExistsPathFuncValueInvalidFilter(t *testing.T) {
	ctx := test.NewTestContext(t)
	loopEvaluator := NewRawStringLiteral("test")

	f := newExistsFunction()
	res, err := f.Execute(ctx, hipathsys.NewString("test"), []interface{}{}, hipathsys.NewLoop(loopEvaluator))
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no res expected")
}

func TestExistsPathFuncValueErrorFilter(t *testing.T) {
	ctx := test.NewTestContext(t)
	loopEvaluator := newTestErrorExpression()

	f := newExistsFunction()
	res, err := f.Execute(ctx, hipathsys.NewString("test"), []interface{}{}, hipathsys.NewLoop(loopEvaluator))
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no res expected")
}

func TestAllPathFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	loopEval := NewEqualityExpression(false, false,
		NewThisInvocation(), NewIndexInvocation())

	f := newAllFunction()
	res, err := f.Execute(ctx, nil, nil, hipathsys.NewLoop(loopEval))
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.True, res)
}

func TestAllPathFuncEmpty(t *testing.T) {
	ctx := test.NewTestContext(t)

	loopEval := NewEqualityExpression(false, false,
		NewThisInvocation(), NewIndexInvocation())

	f := newAllFunction()
	res, err := f.Execute(ctx, nil, nil, hipathsys.NewLoop(loopEval))
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.True, res)
}

func TestAllPathFuncSingleValue(t *testing.T) {
	ctx := test.NewTestContext(t)

	loopEval := NewEqualityExpression(false, false,
		NewThisInvocation(), NewNumberLiteralInt(10))

	f := newAllFunction()
	res, err := f.Execute(ctx, hipathsys.NewInteger(10), nil, hipathsys.NewLoop(loopEval))
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.True, res)
}

func TestAllPathFuncTrue(t *testing.T) {
	ctx := test.NewTestContext(t)

	node := ctx.NewCollection()
	node.Add(hipathsys.NewInteger(0))
	node.Add(hipathsys.NewInteger(1))
	node.Add(hipathsys.NewInteger(2))

	loopEval := NewEqualityExpression(false, false,
		NewThisInvocation(), NewIndexInvocation())

	f := newAllFunction()
	res, err := f.Execute(ctx, node, nil, hipathsys.NewLoop(loopEval))
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.True, res)
}

func TestAllPathFuncFalse(t *testing.T) {
	ctx := test.NewTestContext(t)

	node := ctx.NewCollection()
	node.Add(hipathsys.NewInteger(0))
	node.Add(hipathsys.NewInteger(2))
	node.Add(hipathsys.NewInteger(1))

	loopEval := NewEqualityExpression(false, false,
		NewThisInvocation(), NewIndexInvocation())

	f := newAllFunction()
	res, err := f.Execute(ctx, node, nil, hipathsys.NewLoop(loopEval))
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.False, res)
}

func TestAllPathFuncError(t *testing.T) {
	ctx := test.NewTestContext(t)
	f := newAllFunction()
	res, err := f.Execute(ctx, hipathsys.NewString("test"), nil, hipathsys.NewLoop(newTestErrorExpression()))
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestAllPathFuncInvalidType(t *testing.T) {
	ctx := test.NewTestContext(t)
	f := newAllFunction()
	res, err := f.Execute(ctx, hipathsys.NewString("test"), nil, hipathsys.NewLoop(NewRawStringLiteral("")))
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestAllTruePathFuncTypeError(t *testing.T) {
	ctx := test.NewTestContext(t)
	col := ctx.NewCollection()
	col.Add(hipathsys.True)
	col.Add(hipathsys.NewString("test"))

	f := newAllTrueFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestAllTruePathFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newAllTrueFunction()
	res, err := f.Execute(ctx, nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.True, res)
}

func TestAllTruePathFuncEmpty(t *testing.T) {
	ctx := test.NewTestContext(t)
	col := ctx.NewCollection()

	f := newAllTrueFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.True, res)
}

func TestAllTruePathFuncTrue(t *testing.T) {
	ctx := test.NewTestContext(t)
	col := ctx.NewCollection()
	col.Add(hipathsys.True)
	col.Add(hipathsys.True)
	col.Add(hipathsys.True)

	f := newAllTrueFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.True, res)
}

func TestAllTruePathFuncFalse(t *testing.T) {
	ctx := test.NewTestContext(t)
	col := ctx.NewCollection()
	col.Add(hipathsys.True)
	col.Add(hipathsys.False)
	col.Add(hipathsys.True)

	f := newAllTrueFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.False, res)
}

func TestAnyTruePathFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newAnyTrueFunction()
	res, err := f.Execute(ctx, nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.False, res)
}

func TestAnyTruePathFuncEmpty(t *testing.T) {
	ctx := test.NewTestContext(t)
	col := ctx.NewCollection()

	f := newAnyTrueFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.False, res)
}

func TestAnyTruePathFuncTrue(t *testing.T) {
	ctx := test.NewTestContext(t)
	col := ctx.NewCollection()
	col.Add(hipathsys.False)
	col.Add(hipathsys.True)
	col.Add(hipathsys.False)

	f := newAnyTrueFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.True, res)
}

func TestAnyTruePathFuncFalse(t *testing.T) {
	ctx := test.NewTestContext(t)
	col := ctx.NewCollection()
	col.Add(hipathsys.False)
	col.Add(hipathsys.False)
	col.Add(hipathsys.False)

	f := newAnyTrueFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.False, res)
}

func TestAllFalsePathFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newAllFalseFunction()
	res, err := f.Execute(ctx, nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.True, res)
}

func TestAllFalsePathFuncEmpty(t *testing.T) {
	ctx := test.NewTestContext(t)
	col := ctx.NewCollection()

	f := newAllFalseFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.True, res)
}

func TestAllFalsePathFuncTrue(t *testing.T) {
	ctx := test.NewTestContext(t)
	col := ctx.NewCollection()
	col.Add(hipathsys.False)
	col.Add(hipathsys.False)
	col.Add(hipathsys.False)

	f := newAllFalseFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.True, res)
}

func TestAllFalsePathFuncFalse(t *testing.T) {
	ctx := test.NewTestContext(t)
	col := ctx.NewCollection()
	col.Add(hipathsys.False)
	col.Add(hipathsys.True)
	col.Add(hipathsys.False)

	f := newAllFalseFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.False, res)
}

func TestAnyFalsePathFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newAnyFalseFunction()
	res, err := f.Execute(ctx, nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.False, res)
}

func TestAnyFalsePathFuncEmpty(t *testing.T) {
	ctx := test.NewTestContext(t)
	col := ctx.NewCollection()

	f := newAnyFalseFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.False, res)
}

func TestAnyFalsePathFuncTrue(t *testing.T) {
	ctx := test.NewTestContext(t)
	col := ctx.NewCollection()
	col.Add(hipathsys.True)
	col.Add(hipathsys.False)
	col.Add(hipathsys.True)

	f := newAnyFalseFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.True, res)
}

func TestAnyFalsePathFuncFalse(t *testing.T) {
	ctx := test.NewTestContext(t)
	col := ctx.NewCollection()
	col.Add(hipathsys.True)
	col.Add(hipathsys.True)
	col.Add(hipathsys.True)

	f := newAnyFalseFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.False, res)
}

func TestSubsetOfPathFuncTrue(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(hipathsys.NewString("test3"))
	col.Add(hipathsys.NewString("test1"))
	col.Add(hipathsys.NewString("test4"))

	other := ctx.NewCollection()
	other.Add(hipathsys.NewString("test4"))
	other.Add(hipathsys.NewString("test1"))
	other.Add(hipathsys.NewString("test2"))
	other.Add(hipathsys.NewString("test3"))

	f := newSubsetOfFunction()
	res, err := f.Execute(ctx, col, []interface{}{other}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.True, res)
}

func TestSubsetOfPathFuncFalse(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(hipathsys.NewString("test3"))
	col.Add(hipathsys.NewString("test1"))
	col.Add(hipathsys.NewString("test4"))

	other := ctx.NewCollection()
	other.Add(hipathsys.NewString("test4"))
	other.Add(hipathsys.NewString("test5"))
	other.Add(hipathsys.NewString("test2"))
	other.Add(hipathsys.NewString("test3"))

	f := newSubsetOfFunction()
	res, err := f.Execute(ctx, col, []interface{}{other}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.False, res)
}

func TestSubsetOfPathFuncFalseLessItems(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(hipathsys.NewString("test3"))
	col.Add(hipathsys.NewString("test1"))
	col.Add(hipathsys.NewString("test4"))

	other := ctx.NewCollection()
	other.Add(hipathsys.NewString("test3"))
	other.Add(hipathsys.NewString("test1"))

	f := newSubsetOfFunction()
	res, err := f.Execute(ctx, col, []interface{}{other}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.False, res)
}

func TestSupersetOfPathFuncTrue(t *testing.T) {
	ctx := test.NewTestContext(t)

	other := ctx.NewCollection()
	other.Add(hipathsys.NewString("test3"))
	other.Add(hipathsys.NewString("test1"))
	other.Add(hipathsys.NewString("test4"))

	col := ctx.NewCollection()
	col.Add(hipathsys.NewString("test4"))
	col.Add(hipathsys.NewString("test1"))
	col.Add(hipathsys.NewString("test2"))
	col.Add(hipathsys.NewString("test3"))

	f := newSupersetOfFunction()
	res, err := f.Execute(ctx, col, []interface{}{other}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.True, res)
}

func TestSupersetOfPathFuncFalse(t *testing.T) {
	ctx := test.NewTestContext(t)

	other := ctx.NewCollection()
	other.Add(hipathsys.NewString("test3"))
	other.Add(hipathsys.NewString("test1"))
	other.Add(hipathsys.NewString("test4"))

	col := ctx.NewCollection()
	col.Add(hipathsys.NewString("test4"))
	col.Add(hipathsys.NewString("test5"))
	col.Add(hipathsys.NewString("test2"))
	col.Add(hipathsys.NewString("test3"))

	f := newSupersetOfFunction()
	res, err := f.Execute(ctx, col, []interface{}{other}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.False, res)
}

func TestCountPathFunc(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(hipathsys.NewString("test3"))
	col.Add(hipathsys.NewString("test1"))

	f := newCountFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.IntegerAccessor)(nil), res) {
		assert.Equal(t, int32(2), res.(hipathsys.IntegerAccessor).Int())
	}
}

func TestDistinctPathFuncOneItem(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(hipathsys.NewString("test3"))

	f := newDistinctFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.CollectionAccessor)(nil), res) {
		col := res.(hipathsys.CollectionAccessor)
		if assert.Equal(t, 1, col.Count()) {
			assert.Equal(t, hipathsys.NewString("test3"), col.Get(0))
		}
	}
}

func TestDistinctPathFunc(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(hipathsys.NewString("test3"))
	col.Add(hipathsys.NewString("test4"))
	col.Add(hipathsys.NewString("test3"))
	col.Add(hipathsys.NewString("test1"))

	f := newDistinctFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.CollectionAccessor)(nil), res) {
		col := res.(hipathsys.CollectionAccessor)
		if assert.Equal(t, 3, col.Count()) {
			assert.Equal(t, hipathsys.NewString("test3"), col.Get(0))
			assert.Equal(t, hipathsys.NewString("test4"), col.Get(1))
			assert.Equal(t, hipathsys.NewString("test1"), col.Get(2))
		}
	}
}

func TestIsDistinctPathFuncEmpty(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()

	f := newIsDistinctFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestIsDistinctPathFuncOneItem(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(hipathsys.NewString("test3"))

	f := newIsDistinctFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.True, res)
}

func TestIsDistinctPathFunc(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(hipathsys.NewString("test3"))
	col.Add(hipathsys.NewString("test5"))

	f := newIsDistinctFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.True, res)
}

func TestIsDistinctPathFuncNot(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(hipathsys.NewString("test3"))
	col.Add(hipathsys.NewString("test5"))
	col.Add(hipathsys.NewString("test3"))

	f := newIsDistinctFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.False, res)
}
