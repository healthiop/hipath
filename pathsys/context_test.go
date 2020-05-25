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

package pathsys

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModelTypeInfoNil(t *testing.T) {
	ctx := newTestContext(t)
	assert.Nil(t, ModelTypeInfo(ctx.ModelAdapter(), nil))
}

func TestModelTypeInfoSystem(t *testing.T) {
	ctx := newTestContext(t)
	assert.Same(t, StringTypeInfo, ModelTypeInfo(ctx.ModelAdapter(), NewString("test")))
}

func TestModelTypeInfoModel(t *testing.T) {
	ctx := newTestContext(t)
	assert.Same(t, testTypeInfo, ModelTypeInfo(ctx.ModelAdapter(),
		newTestModelNode(10, false, testTypeInfo)))
}

func TestModelEqualBothNil(t *testing.T) {
	ctx := newTestContext(t)
	assert.True(t, ModelEqual(ctx.ModelAdapter(), nil, nil))
}

func TestModelEqualLeftNil(t *testing.T) {
	ctx := newTestContext(t)
	assert.False(t, ModelEqual(ctx.ModelAdapter(), nil, NewString("")))
}

func TestModelEqualRightNil(t *testing.T) {
	ctx := newTestContext(t)
	assert.False(t, ModelEqual(ctx.ModelAdapter(), NewString(""), nil))
}

func TestModelEqualLeftSystem(t *testing.T) {
	ctx := newTestContext(t)
	assert.False(t, ModelEqual(ctx.ModelAdapter(), NewString(""), newTestModelNode(10, false, testTypeInfo)))
}

func TestModelEqualRightSystem(t *testing.T) {
	ctx := newTestContext(t)
	assert.False(t, ModelEqual(ctx.ModelAdapter(), newTestModelNode(10, false, testTypeInfo), NewString("")))
}

func TestModelEqualSystem(t *testing.T) {
	ctx := newTestContext(t)
	assert.True(t, ModelEqual(ctx.ModelAdapter(), NewString("test1"), NewString("test1")))
}

func TestModelEqualSystemNot(t *testing.T) {
	ctx := newTestContext(t)
	assert.False(t, ModelEqual(ctx.ModelAdapter(), NewString("test1"), NewString("test2")))
}

func TestModelEqualModel(t *testing.T) {
	ctx := newTestContext(t)
	assert.True(t, ModelEqual(ctx.ModelAdapter(),
		newTestModelNode(10, false, testTypeInfo), newTestModelNode(10, false, testTypeInfo)))
}

func TestModelEqualModelNot(t *testing.T) {
	ctx := newTestContext(t)
	assert.False(t, ModelEqual(ctx.ModelAdapter(),
		newTestModelNode(10, false, testTypeInfo), newTestModelNode(11, false, testTypeInfo)))
}

func TestModelEquivalentBothNil(t *testing.T) {
	ctx := newTestContext(t)
	assert.True(t, ModelEquivalent(ctx.ModelAdapter(), nil, nil))
}

func TestModelEquivalentLeftNil(t *testing.T) {
	ctx := newTestContext(t)
	assert.False(t, ModelEquivalent(ctx.ModelAdapter(), nil, NewString("")))
}

func TestModelEquivalentRightNil(t *testing.T) {
	ctx := newTestContext(t)
	assert.False(t, ModelEquivalent(ctx.ModelAdapter(), NewString(""), nil))
}

func TestModelEquivalentLeftSystem(t *testing.T) {
	ctx := newTestContext(t)
	assert.False(t, ModelEquivalent(ctx.ModelAdapter(), NewString(""), newTestModelNode(10, false, testTypeInfo)))
}

func TestModelEquivalentRightSystem(t *testing.T) {
	ctx := newTestContext(t)
	assert.False(t, ModelEquivalent(ctx.ModelAdapter(), newTestModelNode(10, false, testTypeInfo), NewString("")))
}

func TestModelEquivalentSystem(t *testing.T) {
	ctx := newTestContext(t)
	assert.True(t, ModelEquivalent(ctx.ModelAdapter(), NewString("test1"), NewString("Test1")))
}

func TestModelEquivalentSystemNot(t *testing.T) {
	ctx := newTestContext(t)
	assert.False(t, ModelEquivalent(ctx.ModelAdapter(), NewString("test1"), NewString("test2")))
}

func TestModelEquivalentModel(t *testing.T) {
	ctx := newTestContext(t)
	assert.True(t, ModelEquivalent(ctx.ModelAdapter(),
		newTestModelNode(10, false, testTypeInfo), newTestModelNode(10.1, false, testTypeInfo)))
}

func TestModelEquivalentModelNot(t *testing.T) {
	ctx := newTestContext(t)
	assert.False(t, ModelEquivalent(ctx.ModelAdapter(),
		newTestModelNode(10, false, testTypeInfo), newTestModelNode(11, false, testTypeInfo)))
}

func TestSystemAnyTypeEqualBothNil(t *testing.T) {
	assert.False(t, SystemAnyTypeEqual(nil, nil))
}

func TestSystemAnyTypeEqualLeftNil(t *testing.T) {
	assert.False(t, SystemAnyTypeEqual(nil, NewString("")))
}

func TestSystemAnyTypeEqualRightNil(t *testing.T) {
	assert.False(t, SystemAnyTypeEqual(NewString(""), nil))
}

func TestSystemAnyTypeEqualSystem(t *testing.T) {
	assert.True(t, SystemAnyTypeEqual(NewString(""), NewString("")))
}

func TestSystemAnyTypeEqualSystemNot(t *testing.T) {
	assert.False(t, SystemAnyTypeEqual(NewString(""), False))
}

func TestSystemAnyTypeEqualModel(t *testing.T) {
	assert.False(t, SystemAnyTypeEqual(NewString(""), newTestModelNode(0, false, testTypeInfo)))
}

func TestSystemAnyEqualBothNil(t *testing.T) {
	assert.True(t, SystemAnyEqual(nil, nil))
}

func TestSystemAnyEqualLeftNil(t *testing.T) {
	assert.False(t, SystemAnyEqual(nil, NewString("")))
}

func TestSystemAnyEqualRightNil(t *testing.T) {
	assert.False(t, SystemAnyEqual(NewString(""), nil))
}

func TestSystemAnyEqualSystem(t *testing.T) {
	assert.True(t, SystemAnyEqual(NewString("test1"), NewString("test1")))
}

func TestSystemAnyEqualSystemNot(t *testing.T) {
	assert.False(t, SystemAnyEqual(NewString("test1"), NewString("test2")))
}

func TestSystemAnyEqualModel(t *testing.T) {
	assert.False(t, SystemAnyEqual(NewString("test1"), newTestModelNode(10, false, testTypeInfo)))
}
