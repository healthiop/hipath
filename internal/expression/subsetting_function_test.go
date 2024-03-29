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

func TestSinglePathFuncEmpty(t *testing.T) {
	ctx := test.NewTestContext(t)
	f := newSingleFunction()
	res, err := f.Execute(ctx, nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty result expected")
}

func TestSinglePathFunc(t *testing.T) {
	ctx := test.NewTestContext(t)
	col := ctx.NewCol()

	col.Add(hipathsys.NewString("test"))
	f := newSingleFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.NewString("test"), res)
}

func TestSinglePathFuncMulti(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCol()
	col.Add(hipathsys.NewString("test1"))
	col.Add(hipathsys.NewString("test2"))

	f := newSingleFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestFirstPathFuncEmpty(t *testing.T) {
	ctx := test.NewTestContext(t)
	f := newFirstFunction()
	res, err := f.Execute(ctx, nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty result expected")
}

func TestFirstPathFunc(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCol()
	col.Add(hipathsys.NewString("test1"))
	col.Add(hipathsys.NewString("test2"))

	f := newFirstFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.NewString("test1"), res)
}

func TestLastPathFuncEmpty(t *testing.T) {
	ctx := test.NewTestContext(t)
	f := newLastFunction()
	res, err := f.Execute(ctx, nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty result expected")
}

func TestLastPathFunc(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCol()
	col.Add(hipathsys.NewString("test1"))
	col.Add(hipathsys.NewString("test2"))

	f := newLastFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.NewString("test2"), res)
}

func TestTailPathFuncEmpty(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCol()
	col.Add(hipathsys.NewString("test1"))

	f := newTailFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestTailPathFunc(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCol()
	col.Add(hipathsys.NewString("test1"))
	col.Add(hipathsys.NewString("test2"))
	col.Add(hipathsys.NewString("test3"))

	f := newTailFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.ColAccessor)(nil), res) {
		col := res.(hipathsys.ColAccessor)
		if assert.Equal(t, 2, col.Count()) {
			assert.Equal(t, hipathsys.NewString("test2"), col.Get(0))
			assert.Equal(t, hipathsys.NewString("test3"), col.Get(1))
		}
	}
}

func TestSkipPathFuncEmpty(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCol()
	col.Add(hipathsys.NewString("test1"))
	col.Add(hipathsys.NewString("test2"))

	f := newSkipFunction()
	res, err := f.Execute(ctx, col, []interface{}{hipathsys.NewInteger(2)}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestSkipPathFuncError(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCol()
	col.Add(hipathsys.NewString("test1"))

	f := newSkipFunction()
	res, err := f.Execute(ctx, col, []interface{}{hipathsys.NewString("")}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestSkipPathFunc(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCol()
	col.Add(hipathsys.NewString("test1"))
	col.Add(hipathsys.NewString("test2"))
	col.Add(hipathsys.NewString("test3"))
	col.Add(hipathsys.NewString("test4"))
	col.Add(hipathsys.NewString("test5"))

	f := newSkipFunction()
	res, err := f.Execute(ctx, col, []interface{}{hipathsys.NewInteger(2)}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.ColAccessor)(nil), res) {
		col := res.(hipathsys.ColAccessor)
		if assert.Equal(t, 3, col.Count()) {
			assert.Equal(t, hipathsys.NewString("test3"), col.Get(0))
			assert.Equal(t, hipathsys.NewString("test4"), col.Get(1))
			assert.Equal(t, hipathsys.NewString("test5"), col.Get(2))
		}
	}
}

func TestSkipPathFuncNeg(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCol()
	col.Add(hipathsys.NewString("test1"))
	col.Add(hipathsys.NewString("test2"))
	col.Add(hipathsys.NewString("test3"))

	f := newSkipFunction()
	res, err := f.Execute(ctx, col, []interface{}{hipathsys.NewInteger(-10)}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.ColAccessor)(nil), res) {
		col := res.(hipathsys.ColAccessor)
		if assert.Equal(t, 3, col.Count()) {
			assert.Equal(t, hipathsys.NewString("test1"), col.Get(0))
			assert.Equal(t, hipathsys.NewString("test2"), col.Get(1))
			assert.Equal(t, hipathsys.NewString("test3"), col.Get(2))
		}
	}
}

func TestTakePathFuncEmpty(t *testing.T) {
	ctx := test.NewTestContext(t)
	f := newTakeFunction()
	res, err := f.Execute(ctx, nil, []interface{}{hipathsys.NewInteger(1)}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestTakePathFuncZero(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCol()
	col.Add(hipathsys.NewString("test1"))
	col.Add(hipathsys.NewString("test2"))

	f := newTakeFunction()
	res, err := f.Execute(ctx, col, []interface{}{hipathsys.NewInteger(0)}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestTakePathFuncError(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCol()
	col.Add(hipathsys.NewString("test1"))

	f := newTakeFunction()
	res, err := f.Execute(ctx, col, []interface{}{hipathsys.NewString("")}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestTakePathFunc(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCol()
	col.Add(hipathsys.NewString("test1"))
	col.Add(hipathsys.NewString("test2"))
	col.Add(hipathsys.NewString("test3"))
	col.Add(hipathsys.NewString("test4"))
	col.Add(hipathsys.NewString("test5"))

	f := newTakeFunction()
	res, err := f.Execute(ctx, col, []interface{}{hipathsys.NewInteger(3)}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.ColAccessor)(nil), res) {
		col := res.(hipathsys.ColAccessor)
		if assert.Equal(t, 3, col.Count()) {
			assert.Equal(t, hipathsys.NewString("test1"), col.Get(0))
			assert.Equal(t, hipathsys.NewString("test2"), col.Get(1))
			assert.Equal(t, hipathsys.NewString("test3"), col.Get(2))
		}
	}
}

func TestTakePathFuncMore(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCol()
	col.Add(hipathsys.NewString("test1"))
	col.Add(hipathsys.NewString("test2"))
	col.Add(hipathsys.NewString("test3"))

	f := newTakeFunction()
	res, err := f.Execute(ctx, col, []interface{}{hipathsys.NewInteger(5)}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.ColAccessor)(nil), res) {
		col := res.(hipathsys.ColAccessor)
		if assert.Equal(t, 3, col.Count()) {
			assert.Equal(t, hipathsys.NewString("test1"), col.Get(0))
			assert.Equal(t, hipathsys.NewString("test2"), col.Get(1))
			assert.Equal(t, hipathsys.NewString("test3"), col.Get(2))
		}
	}
}

func TestTakePathFuncNeg(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCol()
	col.Add(hipathsys.NewString("test1"))
	col.Add(hipathsys.NewString("test2"))
	col.Add(hipathsys.NewString("test3"))

	f := newTakeFunction()
	res, err := f.Execute(ctx, col, []interface{}{hipathsys.NewInteger(-10)}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestIntersectPathFuncLeftEmpty(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCol()
	other := ctx.NewCol()
	other.Add(hipathsys.NewString("test1"))

	f := newIntersectFunction()
	res, err := f.Execute(ctx, col, []interface{}{other}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestIntersectPathFuncRightEmpty(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCol()
	col.Add(hipathsys.NewString("test1"))
	other := ctx.NewCol()

	f := newIntersectFunction()
	res, err := f.Execute(ctx, col, []interface{}{other}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestIntersectPathFuncLeftBigger(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCol()
	col.Add(hipathsys.NewString("test2"))
	col.Add(hipathsys.NewString("test3"))

	f := newIntersectFunction()
	res, err := f.Execute(ctx, col, []interface{}{hipathsys.NewString("test3")}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.ColAccessor)(nil), res) {
		col := res.(hipathsys.ColAccessor)
		if assert.Equal(t, 1, col.Count()) {
			assert.Equal(t, hipathsys.NewString("test3"), col.Get(0))
		}
	}
}

func TestIntersectPathFuncRightBigger(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCol()
	col.Add(hipathsys.NewString("test2"))
	col.Add(hipathsys.NewString("test3"))

	f := newIntersectFunction()
	res, err := f.Execute(ctx, hipathsys.NewString("test3"), []interface{}{col}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.ColAccessor)(nil), res) {
		col := res.(hipathsys.ColAccessor)
		if assert.Equal(t, 1, col.Count()) {
			assert.Equal(t, hipathsys.NewString("test3"), col.Get(0))
		}
	}
}

func TestIntersectPathFuncUnique(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCol()
	col.Add(hipathsys.NewString("test2"))
	col.Add(hipathsys.NewString("test3"))
	col.Add(hipathsys.NewString("test5"))
	col.Add(hipathsys.NewString("test2"))
	col.Add(hipathsys.NewString("test1"))

	other := ctx.NewCol()
	other.Add(hipathsys.NewString("test3"))
	other.Add(hipathsys.NewString("test7"))
	other.Add(hipathsys.NewString("test2"))
	other.Add(hipathsys.NewString("test1"))

	f := newIntersectFunction()
	res, err := f.Execute(ctx, col, []interface{}{other}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.ColAccessor)(nil), res) {
		col := res.(hipathsys.ColAccessor)
		if assert.Equal(t, 3, col.Count()) {
			assert.Equal(t, hipathsys.NewString("test3"), col.Get(0))
			assert.Equal(t, hipathsys.NewString("test2"), col.Get(1))
			assert.Equal(t, hipathsys.NewString("test1"), col.Get(2))
		}
	}
}

func TestExcludePathFuncLeftEmpty(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCol()
	other := ctx.NewCol()
	other.Add(hipathsys.NewString("test1"))

	f := newExcludeFunction()
	res, err := f.Execute(ctx, col, []interface{}{other}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestExcludePathFuncRightEmpty(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCol()
	col.Add(hipathsys.NewString("test1"))
	col.Add(hipathsys.NewString("test3"))
	other := ctx.NewCol()

	f := newExcludeFunction()
	res, err := f.Execute(ctx, col, []interface{}{other}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.ColAccessor)(nil), res) {
		col := res.(hipathsys.ColAccessor)
		if assert.Equal(t, 2, col.Count()) {
			assert.Equal(t, hipathsys.NewString("test1"), col.Get(0))
			assert.Equal(t, hipathsys.NewString("test3"), col.Get(1))
		}
	}
}

func TestExcludePathFunc(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCol()
	col.Add(hipathsys.NewString("test1"))
	col.Add(hipathsys.NewString("test3"))
	col.Add(hipathsys.NewString("test7"))
	col.Add(hipathsys.NewString("test3"))
	col.Add(hipathsys.NewString("test8"))
	col.Add(hipathsys.NewString("test9"))
	other := ctx.NewCol()
	other.Add(hipathsys.NewString("test7"))
	other.Add(hipathsys.NewString("test9"))

	f := newExcludeFunction()
	res, err := f.Execute(ctx, col, []interface{}{other}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.ColAccessor)(nil), res) {
		col := res.(hipathsys.ColAccessor)
		if assert.Equal(t, 4, col.Count()) {
			assert.Equal(t, hipathsys.NewString("test1"), col.Get(0))
			assert.Equal(t, hipathsys.NewString("test3"), col.Get(1))
			assert.Equal(t, hipathsys.NewString("test3"), col.Get(2))
			assert.Equal(t, hipathsys.NewString("test8"), col.Get(3))
		}
	}
}
