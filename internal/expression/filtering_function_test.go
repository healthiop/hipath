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

func TestWherePathFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)
	f := newWhereFunction()
	res, err := f.Execute(ctx, nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty result expected")
}

func TestWherePathFunc(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCol()
	col.Add(hipathsys.NewInteger(0))
	col.Add(hipathsys.NewInteger(7))
	col.Add(hipathsys.NewInteger(2))

	loopEval := NewEqualityExpression(false, false,
		NewThisInvocation(), NewIndexInvocation())

	f := newWhereFunction()
	res, err := f.Execute(ctx, col, nil, hipathsys.NewLoop(loopEval))
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.ColAccessor)(nil), res) {
		col := res.(hipathsys.ColAccessor)
		if assert.Equal(t, 2, col.Count()) {
			assert.Equal(t, hipathsys.NewInteger(0), col.Get(0))
			assert.Equal(t, hipathsys.NewInteger(2), col.Get(1))
		}
	}
}

func TestWherePathFuncNoMatch(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCol()
	col.Add(hipathsys.NewInteger(2))
	col.Add(hipathsys.NewInteger(7))
	col.Add(hipathsys.NewInteger(0))

	loopEval := NewEqualityExpression(false, false,
		NewThisInvocation(), NewIndexInvocation())

	f := newWhereFunction()
	res, err := f.Execute(ctx, col, nil, hipathsys.NewLoop(loopEval))
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty result expected")
}

func TestWherePathFuncEvalNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCol()
	col.Add(hipathsys.NewInteger(2))
	col.Add(hipathsys.NewInteger(7))
	col.Add(hipathsys.NewInteger(0))

	f := newWhereFunction()
	res, err := f.Execute(ctx, col, nil, hipathsys.NewLoop(newTestExpression(nil)))
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty result expected")
}

func TestWherePathFuncError(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCol()
	col.Add(hipathsys.NewInteger(0))

	f := newWhereFunction()
	res, err := f.Execute(ctx, col, nil, hipathsys.NewLoop(newTestErrorExpression()))
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestWherePathFuncTypeError(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCol()
	col.Add(hipathsys.NewInteger(0))

	f := newWhereFunction()
	res, err := f.Execute(ctx, col, nil, hipathsys.NewLoop(NewRawStringLiteral("")))
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestSelectPathFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)
	f := newSelectFunction()
	res, err := f.Execute(ctx, nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty result expected")
}

func TestSelectPathFunc(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCol()
	col.Add(hipathsys.NewInteger(5))
	col.Add(hipathsys.NewInteger(7))

	loopEval := NewUnionExpression(NewThisInvocation(), NewIndexInvocation())

	f := newSelectFunction()
	res, err := f.Execute(ctx, col, nil, hipathsys.NewLoop(loopEval))
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.ColAccessor)(nil), res) {
		col := res.(hipathsys.ColAccessor)
		if assert.Equal(t, 4, col.Count()) {
			assert.Equal(t, hipathsys.NewInteger(5), col.Get(0))
			assert.Equal(t, hipathsys.NewInteger(0), col.Get(1))
			assert.Equal(t, hipathsys.NewInteger(7), col.Get(2))
			assert.Equal(t, hipathsys.NewInteger(1), col.Get(3))
		}
	}
}

func TestSelectPathFuncSingleItems(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCol()
	col.Add(hipathsys.NewInteger(5))
	col.Add(hipathsys.NewInteger(7))

	f := newSelectFunction()
	res, err := f.Execute(ctx, col, nil, hipathsys.NewLoop(NewIndexInvocation()))
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.ColAccessor)(nil), res) {
		col := res.(hipathsys.ColAccessor)
		if assert.Equal(t, 2, col.Count()) {
			assert.Equal(t, hipathsys.NewInteger(0), col.Get(0))
			assert.Equal(t, hipathsys.NewInteger(1), col.Get(1))
		}
	}
}

func TestSelectPathFuncError(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCol()
	col.Add(hipathsys.NewInteger(5))

	f := newSelectFunction()
	res, err := f.Execute(ctx, col, nil, hipathsys.NewLoop(newTestErrorExpression()))
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestSelectPathFuncEvalNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCol()
	col.Add(hipathsys.NewInteger(5))

	f := newSelectFunction()
	res, err := f.Execute(ctx, col, nil, hipathsys.NewLoop(newTestExpression(nil)))
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty result expected")
}

func TestRepeatPathFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)
	f := repeatFunc
	res, err := f.Execute(ctx, nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty result expected")
}

func TestRepeatPathFunc(t *testing.T) {
	ctx := test.NewTestContext(t)

	node111 := make(map[string]interface{})
	node111["id"] = "111"
	node111["item"] = nil

	node11 := make(map[string]interface{})
	node11["id"] = "11"
	node11["item"] = node111

	node1 := make(map[string]interface{})
	node1["id"] = "1"
	node1["item"] = node11

	loopEval := NewMemberInvocation("item")

	f := repeatFunc
	res, err := f.Execute(ctx, node1, nil, hipathsys.NewLoop(loopEval))
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.ColAccessor)(nil), res) {
		col := res.(hipathsys.ColAccessor)
		if assert.Equal(t, 2, col.Count()) {
			assert.Equal(t, node11, col.Get(0))
			assert.Equal(t, node111, col.Get(1))
		}
	}
}

func TestRepeatPathFuncErr(t *testing.T) {
	ctx := test.NewTestContext(t)

	node11 := make(map[string]interface{})
	node11["id"] = "11"

	node1 := make(map[string]interface{})
	node1["id"] = "1"
	node1["item"] = node11

	loopEval := NewMemberInvocation("item")

	f := repeatFunc
	res, err := f.Execute(ctx, node1, nil, hipathsys.NewLoop(loopEval))
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestRepeatPathFuncColInCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	node11 := createRepeatTestColData("11")
	col := ctx.NewCol()
	node111 := createRepeatTestColData("111")
	col.Add(node111)
	node112 := createRepeatTestColData("112")
	col.Add(node112)
	node11["items"] = col

	node12 := createRepeatTestColData("12")
	col = ctx.NewCol()
	node121 := createRepeatTestColData("121")
	col.Add(node121)
	node122 := createRepeatTestColData("122")
	col.Add(node122)
	col.Add(nil)
	node12["items"] = col

	node1 := createRepeatTestColData("1")
	col = ctx.NewCol()
	col.Add(node11)
	col.Add(node12)
	col.Add(node11)
	node1["items"] = col

	loopEval := NewMemberInvocation("items")

	f := repeatFunc
	res, err := f.Execute(ctx, node1, nil, hipathsys.NewLoop(loopEval))
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.ColAccessor)(nil), res) {
		col := res.(hipathsys.ColAccessor)
		if assert.Equal(t, 6, col.Count()) {
			assert.Equal(t, node11, col.Get(0))
			assert.Equal(t, node111, col.Get(1))
			assert.Equal(t, node112, col.Get(2))
			assert.Equal(t, node12, col.Get(3))
			assert.Equal(t, node121, col.Get(4))
			assert.Equal(t, node122, col.Get(5))
		}
	}
}

func TestRepeatPathFuncColInColError(t *testing.T) {
	ctx := test.NewTestContext(t)

	node11 := make(map[string]interface{})
	node11["id"] = "11"

	node1 := createRepeatTestColData("1")
	col := ctx.NewCol()
	col.Add(node11)
	node1["items"] = col

	loopEval := NewMemberInvocation("items")

	f := repeatFunc
	res, err := f.Execute(ctx, node1, nil, hipathsys.NewLoop(loopEval))
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func createRepeatTestColData(id string) map[string]interface{} {
	node := make(map[string]interface{})
	node["id"] = id
	node["items"] = hipathsys.EmptyCol
	return node
}

func TestOfTypePathFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)
	f := newOfTypeFunction()
	res, err := f.Execute(ctx, nil, []interface{}{hipathsys.NewString("System.String")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty result expected")
}

func TestOfTypePathFuncInvalidTypeSpecType(t *testing.T) {
	ctx := test.NewTestContext(t)
	f := newOfTypeFunction()
	res, err := f.Execute(ctx, nil, []interface{}{"test"}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestOfTypePathFuncInvalidTypeName(t *testing.T) {
	ctx := test.NewTestContext(t)
	f := newOfTypeFunction()
	res, err := f.Execute(ctx, nil, []interface{}{hipathsys.NewString("System.String.Any")}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestOfTypePathFunc(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCol()
	col.Add(hipathsys.NewString("test1"))
	col.Add(hipathsys.NewInteger(10))
	col.Add(hipathsys.NewString("test2"))

	f := newOfTypeFunction()
	res, err := f.Execute(ctx, col, []interface{}{hipathsys.NewString("System.String")}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.ColAccessor)(nil), res) {
		col := res.(hipathsys.ColAccessor)
		if assert.Equal(t, 2, col.Count()) {
			assert.Equal(t, hipathsys.NewString("test1"), col.Get(0))
			assert.Equal(t, hipathsys.NewString("test2"), col.Get(1))
		}
	}
}

func TestOfTypePathFuncBaseType(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCol()
	col.Add(hipathsys.NewString("test1"))
	col.Add(hipathsys.NewInteger(10))
	col.Add(hipathsys.NewString("test2"))

	f := newOfTypeFunction()
	res, err := f.Execute(ctx, col, []interface{}{hipathsys.NewString("System.Any")}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.ColAccessor)(nil), res) {
		col := res.(hipathsys.ColAccessor)
		if assert.Equal(t, 3, col.Count()) {
			assert.Equal(t, hipathsys.NewString("test1"), col.Get(0))
			assert.Equal(t, hipathsys.NewInteger(10), col.Get(1))
			assert.Equal(t, hipathsys.NewString("test2"), col.Get(2))
		}
	}
}
