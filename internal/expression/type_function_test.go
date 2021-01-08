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

func TestAsTypeFunc(t *testing.T) {
	ctx := test.NewTestContext(t)
	f := newAsFunction()
	res, err := f.Execute(ctx, hipathsys.NewString("test1"),
		[]interface{}{hipathsys.NewString("System.String")}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.StringAccessor)(nil), res) {
		s := res.(hipathsys.StringAccessor)
		assert.Equal(t, "test1", s.String())
	}
}

func TestAsTypeFuncInvalid(t *testing.T) {
	ctx := test.NewTestContext(t)
	f := newAsFunction()
	res, err := f.Execute(ctx, hipathsys.NewString("test1"),
		[]interface{}{hipathsys.NewString("System.")}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestAsTypeFuncColSingleItem(t *testing.T) {
	ctx := test.NewTestContext(t)

	node := hipathsys.NewString("test1")
	col := ctx.NewCollection()
	col.Add(node)

	f := newAsFunction()
	res, err := f.Execute(ctx, col,
		[]interface{}{hipathsys.NewString("System.String")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Same(t, node, res)
}

func TestAsTypeFuncTypeDiffers(t *testing.T) {
	ctx := test.NewTestContext(t)
	f := newAsFunction()
	res, err := f.Execute(ctx, hipathsys.NewString("test1"),
		[]interface{}{hipathsys.NewString("XXX.String")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty result expected")
}

func TestAsTypeFuncEmpty(t *testing.T) {
	ctx := test.NewTestContext(t)
	f := newAsFunction()
	res, err := f.Execute(ctx, nil,
		[]interface{}{hipathsys.NewString("XXX.String")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty result expected")
}

func TestAsTypeFuncEmptyType(t *testing.T) {
	ctx := test.NewTestContext(t)
	f := newAsFunction()
	res, err := f.Execute(ctx, hipathsys.NewString("XXX.String"),
		[]interface{}{nil}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty result expected")
}

func TestAsTypeFuncColMultipleItems(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(hipathsys.NewString("test1"))
	col.Add(hipathsys.NewString("test2"))

	f := newAsFunction()
	res, err := f.Execute(ctx, col,
		[]interface{}{hipathsys.NewString("System.String")}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestAsTypeFuncModelUnchanged(t *testing.T) {
	ctx := test.NewTestContext(t)

	n := test.NewTestModelNode(16.1, false)
	f := newAsFunction()
	res, err := f.Execute(ctx, n,
		[]interface{}{hipathsys.NewString("Test.decimal")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Same(t, n, res)
}

func TestAsTypeFuncModelConverted(t *testing.T) {
	ctx := test.NewTestContext(t)

	n := test.NewTestModelNode(16.1, false)
	f := newAsFunction()
	res, err := f.Execute(ctx, n,
		[]interface{}{hipathsys.NewString("System.Number")}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.NumberAccessor)(nil), res) {
		assert.Equal(t, 16.1, res.(hipathsys.NumberAccessor).Float64())
	}
}

func TestAsTypeFuncModelError(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newAsFunction()
	n := test.NewTestModelNode(16.1, false)
	res, err := f.Execute(ctx, n,
		[]interface{}{hipathsys.NewString("Other.Number")}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty result expected")
}

func TestIsTypeFunc(t *testing.T) {
	ctx := test.NewTestContext(t)
	f := newIsFunction()
	res, err := f.Execute(ctx, hipathsys.NewString("test1"),
		[]interface{}{hipathsys.NewString("System.String")}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.BooleanAccessor)(nil), res) {
		b := res.(hipathsys.BooleanAccessor)
		assert.Equal(t, true, b.Bool())
	}
}

func TestIsTypeFuncInvalid(t *testing.T) {
	ctx := test.NewTestContext(t)
	f := newIsFunction()
	res, err := f.Execute(ctx, hipathsys.NewString("test1"),
		[]interface{}{hipathsys.NewString("System.")}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestIsTypeFuncColSingleItem(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(hipathsys.NewString("test1"))

	f := newIsFunction()
	res, err := f.Execute(ctx, col,
		[]interface{}{hipathsys.NewString("System.String")}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.BooleanAccessor)(nil), res) {
		b := res.(hipathsys.BooleanAccessor)
		assert.Equal(t, true, b.Bool())
	}
}

func TestIsTypeFuncTypeDiffers(t *testing.T) {
	ctx := test.NewTestContext(t)
	f := newIsFunction()
	res, err := f.Execute(ctx, hipathsys.NewString("test1"),
		[]interface{}{hipathsys.NewString("XXX.String")}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.BooleanAccessor)(nil), res) {
		b := res.(hipathsys.BooleanAccessor)
		assert.Equal(t, false, b.Bool())
	}
}

func TestIsTypeFuncEmpty(t *testing.T) {
	ctx := test.NewTestContext(t)
	f := newIsFunction()
	res, err := f.Execute(ctx, nil,
		[]interface{}{hipathsys.NewString("XXX.String")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.False, res)
}

func TestIsTypeFuncEmptyType(t *testing.T) {
	ctx := test.NewTestContext(t)
	f := newIsFunction()
	res, err := f.Execute(ctx, hipathsys.NewString("XXX.String"),
		[]interface{}{nil}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty result expected")
}

func TestIsTypeFuncColMultipleItems(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(hipathsys.NewString("test1"))
	col.Add(hipathsys.NewString("test2"))

	f := newIsFunction()
	res, err := f.Execute(ctx, col,
		[]interface{}{hipathsys.NewString("System.String")}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}
