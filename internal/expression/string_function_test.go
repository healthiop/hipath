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

func TestIndexOfFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newIndexOfFunction()
	res, err := f.Execute(ctx, nil, []interface{}{pathsys.NewString("test")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestIndexOfFuncSubstringNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newIndexOfFunction()
	res, err := f.Execute(ctx, pathsys.NewString("test"), []interface{}{nil}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestIndexOfFuncOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newIndexOfFunction()
	res, err := f.Execute(ctx, "test", []interface{}{pathsys.NewString("test")}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected expected")
}

func TestIndexOfFuncSubstringOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newIndexOfFunction()
	res, err := f.Execute(ctx, pathsys.NewString("test"), []interface{}{"test"}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected expected")
}

func TestIndexOfFuncMultiCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewString("test"))
	col.Add(pathsys.NewString("test"))

	f := newIndexOfFunction()
	res, err := f.Execute(ctx, col, []interface{}{pathsys.NewString("test")}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected expected")
}

func TestIndexOfFuncSubstringMultiCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewString("test"))
	col.Add(pathsys.NewString("test"))

	f := newIndexOfFunction()
	res, err := f.Execute(ctx, pathsys.NewString("test"), []interface{}{col}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected expected")
}

func TestIndexOfFunc(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newIndexOfFunction()
	res, err := f.Execute(ctx, pathsys.NewString("This is a test. xy ABC xy"),
		[]interface{}{pathsys.NewString("xy")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewInteger(16), res)
}

func TestIndexOfFuncNotFound(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newIndexOfFunction()
	res, err := f.Execute(ctx, pathsys.NewString("This is a test. xy ABC xy"),
		[]interface{}{pathsys.NewString("xyz")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewInteger(-1), res)
}

func TestSubstringFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newSubstringFunction()
	res, err := f.Execute(ctx, nil, []interface{}{pathsys.NewInteger(3)}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestSubstringFuncNoString(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newSubstringFunction()
	res, err := f.Execute(ctx, pathsys.NewInteger(10), []interface{}{pathsys.NewInteger(3)}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestSubstringFuncValidStart(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newSubstringFunction()
	res, err := f.Execute(ctx, pathsys.NewString("abcdefg"), []interface{}{pathsys.NewInteger(3)}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.StringAccessor)(nil), res) {
		assert.Equal(t, "defg", res.(pathsys.StringAccessor).String())
	}
}

func TestSubstringFuncColValidStart(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewString("abcdefg"))

	f := newSubstringFunction()
	res, err := f.Execute(ctx, col, []interface{}{pathsys.NewInteger(3)}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.StringAccessor)(nil), res) {
		assert.Equal(t, "defg", res.(pathsys.StringAccessor).String())
	}
}

func TestSubstringFuncValidStartCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewInteger(3))

	f := newSubstringFunction()
	res, err := f.Execute(ctx, pathsys.NewString("abcdefg"), []interface{}{col}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.StringAccessor)(nil), res) {
		assert.Equal(t, "defg", res.(pathsys.StringAccessor).String())
	}
}

func TestSubstringFuncValidStartLenNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newSubstringFunction()
	res, err := f.Execute(ctx, pathsys.NewString("abcdefg"),
		[]interface{}{pathsys.NewInteger(3), nil}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.StringAccessor)(nil), res) {
		assert.Equal(t, "defg", res.(pathsys.StringAccessor).String())
	}
}

func TestSubstringFuncStartNoInteger(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newSubstringFunction()
	res, err := f.Execute(ctx, pathsys.NewString("abcdefg"),
		[]interface{}{pathsys.NewDecimalInt(3)}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestSubstringFuncStartNeg(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newSubstringFunction()
	res, err := f.Execute(ctx, pathsys.NewString("abcdefg"),
		[]interface{}{pathsys.NewInteger(-1), pathsys.NewInteger(3)}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.StringAccessor)(nil), res) {
		assert.Equal(t, "abc", res.(pathsys.StringAccessor).String())
	}
}

func TestSubstringFuncValidStartLenNeg(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newSubstringFunction()
	res, err := f.Execute(ctx, pathsys.NewString("abcdefg"),
		[]interface{}{pathsys.NewInteger(2), pathsys.NewInteger(-1)}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestSubstringFuncValidStartLenNoInteger(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newSubstringFunction()
	res, err := f.Execute(ctx, pathsys.NewString("abcdefg"),
		[]interface{}{pathsys.NewInteger(2), pathsys.NewDecimalInt(-1)}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestSubstringFuncSpecialChars(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newSubstringFunction()
	res, err := f.Execute(ctx, pathsys.NewString("áóúñçÁÓÚ"),
		[]interface{}{pathsys.NewInteger(2), pathsys.NewInteger(4)}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.StringAccessor)(nil), res) {
		assert.Equal(t, "úñçÁ", res.(pathsys.StringAccessor).String())
	}
}

func TestSubstringFuncValidStartLen(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newSubstringFunction()
	res, err := f.Execute(ctx, pathsys.NewString("abcdefg"), []interface{}{
		pathsys.NewInteger(1), pathsys.NewInteger(2)}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.StringAccessor)(nil), res) {
		assert.Equal(t, "bc", res.(pathsys.StringAccessor).String())
	}
}

func TestSubstringFuncValidStartColLen(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewInteger(2))

	f := newSubstringFunction()
	res, err := f.Execute(ctx, pathsys.NewString("abcdefg"), []interface{}{
		pathsys.NewInteger(1), col}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.StringAccessor)(nil), res) {
		assert.Equal(t, "bc", res.(pathsys.StringAccessor).String())
	}
}

func TestSubstringFuncValidStartExceededLen(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newSubstringFunction()
	res, err := f.Execute(ctx, pathsys.NewString("abcdefg"), []interface{}{
		pathsys.NewInteger(6), pathsys.NewInteger(2)}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.StringAccessor)(nil), res) {
		assert.Equal(t, "g", res.(pathsys.StringAccessor).String())
	}
}

func TestSubstringFuncExceededStart(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newSubstringFunction()
	res, err := f.Execute(ctx, pathsys.NewString("abcdefg"), []interface{}{
		pathsys.NewInteger(7), pathsys.NewInteger(1)}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestStartsWithFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newStartsWithFunction()
	res, err := f.Execute(ctx, nil, []interface{}{pathsys.NewString("test")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestStartsWithFuncSubstringNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newStartsWithFunction()
	res, err := f.Execute(ctx, pathsys.NewString("test"), []interface{}{nil}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestStartsWithFuncSubstringEmptyString(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newStartsWithFunction()
	res, err := f.Execute(ctx, pathsys.NewString("test"), []interface{}{pathsys.NewString("")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.True, res)
}

func TestStartsWithFuncOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newStartsWithFunction()
	res, err := f.Execute(ctx, "test", []interface{}{pathsys.NewString("test")}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected expected")
}

func TestStartsWithFuncSubstringOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newStartsWithFunction()
	res, err := f.Execute(ctx, pathsys.NewString("test"), []interface{}{"test"}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected expected")
}

func TestStartsWithFuncMultiCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewString("test"))
	col.Add(pathsys.NewString("test"))

	f := newStartsWithFunction()
	res, err := f.Execute(ctx, col, []interface{}{pathsys.NewString("test")}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected expected")
}

func TestStartsWithFuncSubstringMultiCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewString("test"))
	col.Add(pathsys.NewString("test"))

	f := newStartsWithFunction()
	res, err := f.Execute(ctx, pathsys.NewString("test"), []interface{}{col}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected expected")
}

func TestStartsWithFunc(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newStartsWithFunction()
	res, err := f.Execute(ctx, pathsys.NewString("This is a test. xy ABC xy"),
		[]interface{}{pathsys.NewString("This is a")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.True, res)
}

func TestStartsWithFuncNot(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newStartsWithFunction()
	res, err := f.Execute(ctx, pathsys.NewString("This is a test. xy ABC xy"),
		[]interface{}{pathsys.NewString("This is b")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.False, res)
}

func TestStartsWithFuncInside(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newStartsWithFunction()
	res, err := f.Execute(ctx, pathsys.NewString("This is a test. xy ABC xy"),
		[]interface{}{pathsys.NewString("is a test")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.False, res)
}

func TestEndsWithFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newEndsWithFunction()
	res, err := f.Execute(ctx, nil, []interface{}{pathsys.NewString("test")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestEndsWithFuncSubstringNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newEndsWithFunction()
	res, err := f.Execute(ctx, pathsys.NewString("test"), []interface{}{nil}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestEndsWithFuncSubstringEmptyString(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newEndsWithFunction()
	res, err := f.Execute(ctx, pathsys.NewString("test"), []interface{}{pathsys.NewString("")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.True, res)
}

func TestEndsWithFuncOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newEndsWithFunction()
	res, err := f.Execute(ctx, "test", []interface{}{pathsys.NewString("test")}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected expected")
}

func TestEndsWithFuncSubstringOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newEndsWithFunction()
	res, err := f.Execute(ctx, pathsys.NewString("test"), []interface{}{"test"}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected expected")
}

func TestEndsWithFuncMultiCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewString("test"))
	col.Add(pathsys.NewString("test"))

	f := newEndsWithFunction()
	res, err := f.Execute(ctx, col, []interface{}{pathsys.NewString("test")}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected expected")
}

func TestEndsWithFuncSubstringMultiCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewString("test"))
	col.Add(pathsys.NewString("test"))

	f := newEndsWithFunction()
	res, err := f.Execute(ctx, pathsys.NewString("test"), []interface{}{col}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected expected")
}

func TestEndsWithFunc(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newEndsWithFunction()
	res, err := f.Execute(ctx, pathsys.NewString("This is a test. xy ABC xy"),
		[]interface{}{pathsys.NewString("ABC xy")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.True, res)
}

func TestEndsWithFuncNot(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newEndsWithFunction()
	res, err := f.Execute(ctx, pathsys.NewString("This is a test. xy ABC xy"),
		[]interface{}{pathsys.NewString("aBC xy")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.False, res)
}

func TestEndsWithFuncInside(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newEndsWithFunction()
	res, err := f.Execute(ctx, pathsys.NewString("This is a test. xy ABC xy"),
		[]interface{}{pathsys.NewString("is a test")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.False, res)
}

func TestContainsFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newContainsFunction()
	res, err := f.Execute(ctx, nil, []interface{}{pathsys.NewString("test")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestContainsFuncSubstringNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newContainsFunction()
	res, err := f.Execute(ctx, pathsys.NewString("test"), []interface{}{nil}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestContainsFuncSubstringEmptyString(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newContainsFunction()
	res, err := f.Execute(ctx, pathsys.NewString("test"), []interface{}{pathsys.NewString("")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.True, res)
}

func TestContainsFuncOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newContainsFunction()
	res, err := f.Execute(ctx, "test", []interface{}{pathsys.NewString("test")}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected expected")
}

func TestContainsFuncSubstringOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newContainsFunction()
	res, err := f.Execute(ctx, pathsys.NewString("test"), []interface{}{"test"}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected expected")
}

func TestContainsFuncMultiCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewString("test"))
	col.Add(pathsys.NewString("test"))

	f := newContainsFunction()
	res, err := f.Execute(ctx, col, []interface{}{pathsys.NewString("test")}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected expected")
}

func TestContainsFuncSubstringMultiCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewString("test"))
	col.Add(pathsys.NewString("test"))

	f := newContainsFunction()
	res, err := f.Execute(ctx, pathsys.NewString("test"), []interface{}{col}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected expected")
}

func TestContainsFunc(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newContainsFunction()
	res, err := f.Execute(ctx, pathsys.NewString("This is a test. xy ABC xy"), []interface{}{pathsys.NewString("xy ABC")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.True, res)
}

func TestContainsFuncNot(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newContainsFunction()
	res, err := f.Execute(ctx, pathsys.NewString("This is a test. xy ABC xy"), []interface{}{pathsys.NewString("xz ABC")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.False, res)
}

func TestUpperFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newUpperFunction()
	res, err := f.Execute(ctx, nil, []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestUpperFuncOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newUpperFunction()
	res, err := f.Execute(ctx, "test", []interface{}{}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected expected")
}

func TestUpperFuncMultiCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewString("test"))
	col.Add(pathsys.NewString("test"))

	f := newUpperFunction()
	res, err := f.Execute(ctx, col, []interface{}{}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected expected")
}

func TestUpperFunc(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newUpperFunction()
	res, err := f.Execute(ctx, pathsys.NewString("This is a test. xy ABC xy"), []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewString("THIS IS A TEST. XY ABC XY"), res)
}

func TestLowerFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newLowerFunction()
	res, err := f.Execute(ctx, nil, []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestLowerFuncOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newLowerFunction()
	res, err := f.Execute(ctx, "test", []interface{}{}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected expected")
}

func TestLowerFuncMultiCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewString("test"))
	col.Add(pathsys.NewString("test"))

	f := newLowerFunction()
	res, err := f.Execute(ctx, col, []interface{}{}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected expected")
}

func TestLowerFunc(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newLowerFunction()
	res, err := f.Execute(ctx, pathsys.NewString("This is a test. xy ABC xy"), []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewString("this is a test. xy abc xy"), res)
}

func TestReplaceFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newReplaceFunction()
	res, err := f.Execute(ctx, nil, []interface{}{nil, nil}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestReplaceFuncPatternNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newReplaceFunction()
	res, err := f.Execute(ctx, pathsys.NewString("abcdefgcdef"),
		[]interface{}{nil, pathsys.NewString("xy")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestReplaceFuncSubstitutionNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newReplaceFunction()
	res, err := f.Execute(ctx, pathsys.NewString("abcdefgcdef"),
		[]interface{}{pathsys.NewString("cde"), nil}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestReplaceFuncAll(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newReplaceFunction()
	res, err := f.Execute(ctx, pathsys.NewString("abcdefgcdef"),
		[]interface{}{pathsys.NewString("cde"), pathsys.NewString("xy")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewString("abxyfgxyf"), res)
}

func TestReplaceFuncRemove(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newReplaceFunction()
	res, err := f.Execute(ctx, pathsys.NewString("abcdefg"),
		[]interface{}{pathsys.NewString("cde"), pathsys.EmptyString}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewString("abfg"), res)
}

func TestReplaceFuncSurround(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newReplaceFunction()
	res, err := f.Execute(ctx, pathsys.NewString("abc"),
		[]interface{}{pathsys.EmptyString, pathsys.NewString("x")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewString("xaxbxcx"), res)
}

func TestReplaceFuncInputNoString(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newReplaceFunction()
	res, err := f.Execute(ctx, pathsys.True,
		[]interface{}{pathsys.NewString("x"), pathsys.NewString("y")}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestReplaceFuncInputCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewString("xyz"))

	f := newReplaceFunction()
	res, err := f.Execute(ctx, col,
		[]interface{}{pathsys.NewString("x"), pathsys.NewString("y")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewString("yyz"), res)
}

func TestReplaceFuncPatternNoString(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newReplaceFunction()
	res, err := f.Execute(ctx, pathsys.NewString("xyz"),
		[]interface{}{pathsys.True, pathsys.NewString("y")}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestReplaceFuncPatternCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewString("x"))

	f := newReplaceFunction()
	res, err := f.Execute(ctx, pathsys.NewString("xyz"),
		[]interface{}{col, pathsys.NewString("y")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewString("yyz"), res)
}

func TestReplaceFuncSubstitutionNoString(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newReplaceFunction()
	res, err := f.Execute(ctx, pathsys.NewString("xyz"),
		[]interface{}{pathsys.NewString("x"), pathsys.True}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestReplaceFuncSubstitutionCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewString("y"))

	f := newReplaceFunction()
	res, err := f.Execute(ctx, pathsys.NewString("xyz"),
		[]interface{}{pathsys.NewString("x"), col}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewString("yyz"), res)
}

func TestMatchesFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newMatchesFunction()
	res, err := f.Execute(ctx, nil, []interface{}{pathsys.NewString("test")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestMatchesFuncRegexNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newMatchesFunction()
	res, err := f.Execute(ctx, pathsys.NewString("test"), []interface{}{nil}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestMatchesFuncTrue(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newMatchesFunction()
	res, err := f.Execute(ctx, pathsys.NewString("Atest123abcZ"),
		[]interface{}{pathsys.NewString("[a-z]+\\d+[a-z]+")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.True, res)
}

func TestMatchesFuncInputCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewString("Atest123abc"))

	f := newMatchesFunction()
	res, err := f.Execute(ctx, col, []interface{}{pathsys.NewString("[a-z]{3,4}\\d+[a-z]+")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.True, res)
}

func TestMatchesFuncInputInvalid(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newMatchesFunction()
	res, err := f.Execute(ctx, pathsys.True, []interface{}{pathsys.NewString("[a-z]+\\d+[a-z]+")}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestMatchesFuncRegexCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewString("[a-z]+\\d+[a-z]+"))

	f := newMatchesFunction()
	res, err := f.Execute(ctx, pathsys.NewString("Atest123abcZ"), []interface{}{col}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.True, res)
}

func TestMatchesFuncRegexInvalid(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newMatchesFunction()
	res, err := f.Execute(ctx, pathsys.NewString("Atest123abcZ"),
		[]interface{}{pathsys.True}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestMatchesFuncFalse(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newMatchesFunction()
	res, err := f.Execute(ctx, pathsys.NewString("AtEst123abcZ"),
		[]interface{}{pathsys.NewString("[a-z]{3,4}\\d+[a-z]+")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.False, res)
}

func TestMatchesFuncInvalidRegexSyntax(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newMatchesFunction()
	res, err := f.Execute(ctx, pathsys.NewString("Atest123abcZ"),
		[]interface{}{pathsys.NewString("[a-z]+\\d+[a-z+")}, nil)
	if assert.Error(t, err, "error expected") {
		assert.Contains(t, err.Error(), "regexp")
	}
	assert.Nil(t, res, "empty collection expected")
}

func TestReplaceMatchesFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newReplaceMatchesFunction()
	res, err := f.Execute(ctx, nil, []interface{}{nil, nil}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestReplaceMatchesFuncRegexNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newReplaceMatchesFunction()
	res, err := f.Execute(ctx, pathsys.NewString("abcdefgcdef"),
		[]interface{}{nil, pathsys.NewString("xy")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestReplaceMatchesFuncSubstitutionNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newReplaceMatchesFunction()
	res, err := f.Execute(ctx, pathsys.NewString("abcdefgcdef"),
		[]interface{}{pathsys.NewString("cde"), nil}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestReplaceMatchesFuncAllLiteral(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newReplaceMatchesFunction()
	res, err := f.Execute(ctx, pathsys.NewString("abcdefgcdef"),
		[]interface{}{pathsys.NewString("cde"), pathsys.NewString("xy")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewString("abxyfgxyf"), res)
}

func TestReplaceMatchesFuncAll(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newReplaceMatchesFunction()
	res, err := f.Execute(ctx, pathsys.NewString("11/30/1972"),
		[]interface{}{pathsys.NewString("\\b(?P<month>\\d{1,2})/(?P<day>\\d{1,2})/(?P<year>\\d{2,4})\\b"),
			pathsys.NewString("${day}-${month}-${year}")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewString("30-11-1972"), res)
}

func TestReplaceMatchesFuncInputNoString(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newReplaceMatchesFunction()
	res, err := f.Execute(ctx, pathsys.True,
		[]interface{}{pathsys.NewString("x"), pathsys.NewString("y")}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestReplaceMatchesFuncInputCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewString("xyz"))

	f := newReplaceMatchesFunction()
	res, err := f.Execute(ctx, col,
		[]interface{}{pathsys.NewString("x"), pathsys.NewString("y")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewString("yyz"), res)
}

func TestReplaceMatchesFuncRegexNoString(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newReplaceMatchesFunction()
	res, err := f.Execute(ctx, pathsys.NewString("xyz"),
		[]interface{}{pathsys.True, pathsys.NewString("y")}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestReplaceMatchesFuncRegexCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewString("x"))

	f := newReplaceMatchesFunction()
	res, err := f.Execute(ctx, pathsys.NewString("xyz"),
		[]interface{}{col, pathsys.NewString("y")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewString("yyz"), res)
}

func TestReplaceMatchesFuncSubstitutionNoString(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newReplaceMatchesFunction()
	res, err := f.Execute(ctx, pathsys.NewString("xyz"),
		[]interface{}{pathsys.NewString("x"), pathsys.True}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestReplaceMatchesFuncSubstitutionCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewString("y"))

	f := newReplaceMatchesFunction()
	res, err := f.Execute(ctx, pathsys.NewString("xyz"),
		[]interface{}{pathsys.NewString("x"), col}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewString("yyz"), res)
}

func TestReplaceMatchesFuncInvalidRegexNoString(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newReplaceMatchesFunction()
	res, err := f.Execute(ctx, pathsys.NewString("xyz"),
		[]interface{}{pathsys.NewString("x("), pathsys.NewString("y")}, nil)
	if assert.Error(t, err, "error expected") {
		assert.Contains(t, err.Error(), "regexp")
	}
	assert.Nil(t, res, "empty collection expected")
}

func TestLengthFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newLengthFunction()
	res, err := f.Execute(ctx, nil, []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestLengthFuncOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newLengthFunction()
	res, err := f.Execute(ctx, "test", []interface{}{}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected expected")
}

func TestLengthFuncMultiCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewString("test"))
	col.Add(pathsys.NewString("test"))

	f := newLengthFunction()
	res, err := f.Execute(ctx, col, []interface{}{}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected expected")
}

func TestLengthFunc(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newLengthFunction()
	res, err := f.Execute(ctx, pathsys.NewString("This is Á."), []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewInteger(10), res)
}

func TestToCharsFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newToCharsFunction()
	res, err := f.Execute(ctx, nil, []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestToCharsFuncOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newToCharsFunction()
	res, err := f.Execute(ctx, "test", []interface{}{}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected expected")
}

func TestToCharsFuncMultiCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewString("test"))
	col.Add(pathsys.NewString("test"))

	f := newToCharsFunction()
	res, err := f.Execute(ctx, col, []interface{}{}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected expected")
}

func TestToCharsFunc(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newToCharsFunction()
	res, err := f.Execute(ctx, pathsys.NewString("T i Áz"), []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.CollectionAccessor)(nil), res) {
		col := res.(pathsys.CollectionAccessor)
		if assert.Equal(t, 6, col.Count()) {
			assert.Equal(t, pathsys.NewString("T"), col.Get(0))
			assert.Equal(t, pathsys.NewString(" "), col.Get(1))
			assert.Equal(t, pathsys.NewString("i"), col.Get(2))
			assert.Equal(t, pathsys.NewString(" "), col.Get(3))
			assert.Equal(t, pathsys.NewString("Á"), col.Get(4))
			assert.Equal(t, pathsys.NewString("z"), col.Get(5))
		}
	}
}

func TestToCharsFuncEmptyString(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newToCharsFunction()
	res, err := f.Execute(ctx, pathsys.NewString(""), []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.CollectionAccessor)(nil), res) {
		col := res.(pathsys.CollectionAccessor)
		assert.Equal(t, 0, col.Count())
	}
}
