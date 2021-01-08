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

package hipathsys

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModelTypeSpecNil(t *testing.T) {
	ctx := newTestContext(t)
	assert.Nil(t, ModelTypeSpec(ctx.ModelAdapter(), nil))
}

func TestModelTypeSpecSystem(t *testing.T) {
	ctx := newTestContext(t)
	assert.Same(t, StringTypeSpec, ModelTypeSpec(ctx.ModelAdapter(), NewString("test")))
}

func TestModelTypeSpecModel(t *testing.T) {
	ctx := newTestContext(t)
	assert.Same(t, testTypeSpec, ModelTypeSpec(ctx.ModelAdapter(),
		newTestModelNode(10, false, testTypeSpec)))
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
	assert.False(t, ModelEqual(ctx.ModelAdapter(), NewString(""), newTestModelNode(10, false, testTypeSpec)))
}

func TestModelEqualRightSystem(t *testing.T) {
	ctx := newTestContext(t)
	assert.False(t, ModelEqual(ctx.ModelAdapter(), newTestModelNode(10, false, testTypeSpec), NewString("")))
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
		newTestModelNode(10, false, testTypeSpec), newTestModelNode(10, false, testTypeSpec)))
}

func TestModelEqualModelNot(t *testing.T) {
	ctx := newTestContext(t)
	assert.False(t, ModelEqual(ctx.ModelAdapter(),
		newTestModelNode(10, false, testTypeSpec), newTestModelNode(11, false, testTypeSpec)))
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
	assert.False(t, ModelEquivalent(ctx.ModelAdapter(), NewString(""), newTestModelNode(10, false, testTypeSpec)))
}

func TestModelEquivalentRightSystem(t *testing.T) {
	ctx := newTestContext(t)
	assert.False(t, ModelEquivalent(ctx.ModelAdapter(), newTestModelNode(10, false, testTypeSpec), NewString("")))
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
		newTestModelNode(10, false, testTypeSpec), newTestModelNode(10.1, false, testTypeSpec)))
}

func TestModelEquivalentModelNot(t *testing.T) {
	ctx := newTestContext(t)
	assert.False(t, ModelEquivalent(ctx.ModelAdapter(),
		newTestModelNode(10, false, testTypeSpec), newTestModelNode(11, false, testTypeSpec)))
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
	assert.False(t, SystemAnyTypeEqual(NewString(""), newTestModelNode(0, false, testTypeSpec)))
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
	assert.False(t, SystemAnyEqual(NewString("test1"), newTestModelNode(10, false, testTypeSpec)))
}

func TestHasModelTypeNil(t *testing.T) {
	ctx := newTestContext(t)
	assert.False(t, HasModelType(ctx.ModelAdapter(), nil,
		NewFQTypeName("String", "System")))
}

func TestHasModelTypeSystem(t *testing.T) {
	ctx := newTestContext(t)
	assert.True(t, HasModelType(ctx.ModelAdapter(), NewString("test"),
		NewFQTypeName("String", "System")))
}

func TestHasModelTypeSystemNot(t *testing.T) {
	ctx := newTestContext(t)
	assert.False(t, HasModelType(ctx.ModelAdapter(), NewString("test"),
		NewFQTypeName("Test", "System")))
}

func TestHasModelTypeNonSystem(t *testing.T) {
	ctx := newTestContext(t)
	assert.True(t, HasModelType(ctx.ModelAdapter(), newTestModelNode(10, false, testTypeSpec),
		NewFQTypeName("base", "TEST")))
}

func TestHasModelTypeNonSystemNot(t *testing.T) {
	ctx := newTestContext(t)
	assert.False(t, HasModelType(ctx.ModelAdapter(), newTestModelNode(10, false, testTypeSpec),
		NewFQTypeName("other", "TEST")))
}

func TestHasModelTypeSystemWithModelType(t *testing.T) {
	ctx := newTestContext(t)
	assert.True(t, HasModelType(ctx.ModelAdapter(), NewString("test"),
		NewFQTypeName("string", "TEST")))
}

func TestHasModelTypeSystemWithModelTypeAndWithoutNamespace(t *testing.T) {
	ctx := newTestContext(t)
	assert.True(t, HasModelType(ctx.ModelAdapter(), NewString("test"),
		NewTypeName("string")))
}

func TestCastModelTypeNil(t *testing.T) {
	ctx := newTestContext(t)
	res, err := CastModelType(ctx.ModelAdapter(), nil,
		NewFQTypeName("String", "Other"))
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty result expected")
}

func TestCastModelTypeSystemSelf(t *testing.T) {
	ctx := newTestContext(t)
	n := NewString("Test 123")
	res, err := CastModelType(ctx.ModelAdapter(), n,
		NewFQTypeName("String", "System"))
	assert.NoError(t, err, "no error expected")
	assert.Same(t, n, res)
}

func TestCastModelTypeSystemIncompatible(t *testing.T) {
	ctx := newTestContext(t)
	n := NewString("Test 123")
	res, err := CastModelType(ctx.ModelAdapter(), n,
		NewFQTypeName("Number", "System"))
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty result expected")
}

func TestCastModelTypeModelSelf(t *testing.T) {
	ctx := newTestContext(t)
	n := newTestModelNode(17.4, false, testTypeSpec)
	res, err := CastModelType(ctx.ModelAdapter(), n,
		NewFQTypeName("decimal", "Test"))
	assert.NoError(t, err, "no error expected")
	assert.Same(t, n, res, "empty result expected")
}

func TestCastModelTypeModelSystem(t *testing.T) {
	ctx := newTestContext(t)
	n := newTestModelNode(17.4, false, testTypeSpec)
	res, err := CastModelType(ctx.ModelAdapter(), n,
		NewFQTypeName("Number", "System"))
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*NumberAccessor)(nil), res) {
		assert.Equal(t, 17.4, res.(NumberAccessor).Float64())
	}
}

func TestCastModelTypeModelUnknown(t *testing.T) {
	ctx := newTestContext(t)
	n := newTestModelNode(17.4, false, testTypeSpec)
	res, err := CastModelType(ctx.ModelAdapter(), n,
		NewFQTypeName("Number", "Invalid"))
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty result expected")
}

func TestCastModelTypeModelError(t *testing.T) {
	ctx := newTestContext(t)
	n := newTestModelNode(17.4, false, testTypeSpec)
	res, err := CastModelType(ctx.ModelAdapter(), n,
		NewFQTypeName("number", "Other"))
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty result expected")
}
