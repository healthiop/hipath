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

package gohipath

import (
	"github.com/healthiop/hipath/hipathsys"
	"github.com/healthiop/hipath/internal/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompileLiteral(t *testing.T) {
	path, err := Compile("true")

	assert.Nil(t, err, "no error expected")
	if assert.NotNil(t, path, "path expected") {
		assert.NotNil(t, path.evaluator, "evaluator expected")
	}
}

func TestCompileEmpty(t *testing.T) {
	path, err := Compile("")

	assert.Nil(t, path, "no path expected")
	if assert.NotNil(t, err, "error expected") {
		if assert.NotNil(t, err.Items(), "items expected") {
			assert.Len(t, err.Items(), 1)
		}
	}
}

func TestCompileInvalid(t *testing.T) {
	path, err := Compile("xxx$#@yyy")

	assert.Nil(t, path, "no path expected")
	if assert.NotNil(t, err, "error expected") {
		if assert.NotNil(t, err.Items(), "items expected") {
			assert.Len(t, err.Items(), 2)
		}
	}
}

func TestExecuteEmpty(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, err := Execute(ctx, "{}", nil)
	assert.Nil(t, err, "no error expected")
	if assert.NotNil(t, res, "result expected") {
		assert.Equal(t, 0, res.Count())
	}
}

func TestExecuteSingle(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, err := Execute(ctx, "length()", hipathsys.NewString("This is a test!"))
	assert.Nil(t, err, "no error expected")
	if assert.NotNil(t, res, "result expected") {
		assert.Equal(t, 1, res.Count())
		assert.Equal(t, hipathsys.NewInteger(15), res.Get(0))
	}
}

func TestExecuteMulti(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, err := Execute(ctx, "union('value2' | 'value8')", hipathsys.NewString("value5"))
	assert.Nil(t, err, "no error expected")
	if assert.NotNil(t, res, "result expected") {
		assert.Equal(t, 3, res.Count())
		assert.Equal(t, hipathsys.NewString("value5"), res.Get(0))
		assert.Equal(t, hipathsys.NewString("value2"), res.Get(1))
		assert.Equal(t, hipathsys.NewString("value8"), res.Get(2))
	}
}

func TestExecuteErrorCompile(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, err := Execute(ctx, "xxx$#@yyy", nil)
	if assert.NotNil(t, err, "error expected") {
		if assert.NotNil(t, err.Items(), "items expected") {
			assert.Len(t, err.Items(), 2)
		}
	}
	assert.Nil(t, res, "no result expected")
}

func TestExecuteErrorExecute(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, err := Execute(ctx, "union('value2' | 'value8').single()", nil)
	if assert.NotNil(t, err, "error expected") {
		assert.Nil(t, err.Items(), "no items expected")
	}
	assert.Nil(t, res, "no result expected")
}
