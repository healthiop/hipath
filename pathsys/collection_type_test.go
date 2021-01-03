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

package pathsys

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCollectionSource(t *testing.T) {
	ctx := newTestContext(t)
	o := NewCollectionWithSource(ctx.ModelAdapter(), "abc")
	assert.Equal(t, "abc", o.Source())
}

func TestCollectionDataType(t *testing.T) {
	ctx := newTestContext(t)
	c := ctx.NewCollection()
	assert.Equal(t, CollectionDataType, c.DataType())
}

func TestCollectionTypeSpec(t *testing.T) {
	ctx := newTestContext(t)
	c := ctx.NewCollection()
	ti := c.TypeSpec()
	if assert.NotNil(t, ti, "type info expected") {
		assert.Equal(t, "System.Collection", ti.String())
		if assert.NotNil(t, ti, "base type info expected") {
			assert.Equal(t, "System.Any", ti.Base().String())
		}
	}
}

func TestNewCollectionNilCtx(t *testing.T) {
	assert.Panics(t, func() { NewCollection(nil) })
}

func TestNewCollection(t *testing.T) {
	ctx := newTestContext(t)
	c := ctx.NewCollection()
	assert.True(t, c.Empty(), "new collection must be empty")
	assert.Equal(t, 0, c.Count())
	assert.Same(t, UndefinedTypeSpec, c.ItemTypeSpec())
}

func TestNewCollectionGetEmpty(t *testing.T) {
	ctx := newTestContext(t)
	c := ctx.NewCollection()
	assert.Panics(t, func() { c.Get(0) })
}

func TestCollectionAddGet(t *testing.T) {
	ctx := newTestContext(t)
	item1 := NewString("test1")
	item2 := NewString("test2")
	c := ctx.NewCollection()
	c.Add(item1)
	c.Add(item2)
	assert.False(t, c.Empty(), "collection contains elements")
	assert.Equal(t, 2, c.Count())
	assert.Same(t, item1, c.Get(0))
	assert.Same(t, item2, c.Get(1))
	assert.Same(t, item1.TypeSpec(), c.ItemTypeSpec())
}

func TestNewCollectionWithItem(t *testing.T) {
	ctx := newTestContext(t)
	item1 := NewString("test1")
	c := ctx.NewCollectionWithItem(item1)
	assert.False(t, c.Empty(), "collection contains elements")
	assert.Equal(t, 1, c.Count())
	assert.Same(t, item1, c.Get(0))
	assert.Same(t, item1.TypeSpec(), c.ItemTypeSpec())
}

func TestCollectionAddBaseType(t *testing.T) {
	ctx := newTestContext(t)
	item1 := NewString("test1")
	item2 := NewDecimalInt(10)
	c := ctx.NewCollection()
	c.Add(item1)
	c.Add(item2)
	assert.False(t, c.Empty(), "collection contains elements")
	assert.Equal(t, 2, c.Count())
	assert.Same(t, item1, c.Get(0))
	assert.Same(t, item2, c.Get(1))
	assert.Same(t, anyTypeSpec, c.ItemTypeSpec())
}

func TestCollectionAddGetModel(t *testing.T) {
	ctx := newTestContext(t)
	item1 := newTestModelNode(10, false, testTypeSpec)
	item2 := newTestModelNode(11, false, testTypeSpec)
	c := ctx.NewCollection()
	c.Add(item1)
	c.Add(item2)
	assert.False(t, c.Empty(), "collection contains elements")
	assert.Equal(t, 2, c.Count())
	assert.Same(t, item1, c.Get(0))
	assert.Same(t, item2, c.Get(1))
	assert.Same(t, testTypeSpec, c.ItemTypeSpec())
}

func TestNewCollectionWithModemItem(t *testing.T) {
	ctx := newTestContext(t)
	item1 := newTestModelNode(10, false, testTypeSpec)
	c := ctx.NewCollectionWithItem(item1)
	c.Add(item1)
	assert.False(t, c.Empty(), "collection contains elements")
	assert.Equal(t, 2, c.Count())
	assert.Same(t, item1, c.Get(0))
	assert.Same(t, testTypeSpec, c.ItemTypeSpec())
}

func TestCollectionAddGetConvertedModel(t *testing.T) {
	ctx := newTestContext(t)
	item1 := newTestModelNode(10.1, true, testTypeSpec)
	item2 := newTestModelNode(12.1, true, testTypeSpec)
	c := ctx.NewCollection()
	c.Add(item1)
	c.Add(item2)
	assert.False(t, c.Empty(), "collection contains elements")
	assert.Equal(t, 2, c.Count())
	if assert.Implements(t, (*DecimalAccessor)(nil), c.Get(0)) {
		assert.Equal(t, 10.1, c.Get(0).(DecimalAccessor).Float64())
	}
	if assert.Implements(t, (*DecimalAccessor)(nil), c.Get(1)) {
		assert.Equal(t, 12.1, c.Get(1).(DecimalAccessor).Float64())
	}
	assert.Same(t, DecimalTypeSpec, c.ItemTypeSpec())
}

func TestCollectionAddNil(t *testing.T) {
	ctx := newTestContext(t)
	item1 := NewString("test1")
	c := ctx.NewCollection()
	c.Add(item1)
	c.Add(nil)
	assert.False(t, c.Empty(), "collection contains elements")
	assert.Equal(t, 2, c.Count())
	assert.Same(t, item1, c.Get(0))
	assert.Nil(t, c.Get(1))
	assert.Same(t, item1.TypeSpec(), c.ItemTypeSpec())
}

func TestCollectionAddUniqueAll(t *testing.T) {
	ctx := newTestContext(t)
	item1 := NewString("test1")
	item2 := NewString("test2")
	c := ctx.NewCollection()
	assert.Equal(t, true, c.AddUnique(item1))
	assert.Equal(t, true, c.AddUnique(item2))
	assert.False(t, c.Empty(), "collection contains elements")
	assert.Equal(t, 2, c.Count())
	assert.Same(t, item1, c.Get(0))
	assert.Same(t, item2, c.Get(1))
	assert.Same(t, item1.TypeSpec(), c.ItemTypeSpec())
}

func TestCollectionAddUniqueDupNil(t *testing.T) {
	ctx := newTestContext(t)
	c := ctx.NewCollection()
	assert.Equal(t, true, c.AddUnique(nil))
	assert.Equal(t, false, c.AddUnique(nil))
	assert.False(t, c.Empty(), "collection contains elements")
	assert.Equal(t, 1, c.Count())
	assert.Nil(t, c.Get(0))
	assert.Same(t, UndefinedTypeSpec, c.ItemTypeSpec())
}

func TestCollectionAddUniqueModel(t *testing.T) {
	ctx := newTestContext(t)
	item1 := NewString("test1")
	item2 := newTestModelNode(10, false, testTypeSpec)
	item3 := NewString("test1")
	item4 := newTestModelNode(10, false, testTypeSpec)
	c := ctx.NewCollection()
	assert.Equal(t, true, c.AddUnique(item1))
	assert.Equal(t, true, c.AddUnique(item2))
	assert.Equal(t, false, c.AddUnique(item3))
	assert.Equal(t, false, c.AddUnique(item4))
	assert.False(t, c.Empty(), "collection contains elements")
	assert.Equal(t, 2, c.Count())
	assert.Same(t, item1, c.Get(0))
	assert.Same(t, item2, c.Get(1))
	assert.Same(t, UndefinedTypeSpec, c.ItemTypeSpec())
}

func TestCollectionAddUniqueConvertedModel(t *testing.T) {
	ctx := newTestContext(t)
	item1 := NewDecimalFloat64(10)
	item2 := newTestModelNode(10, true, testTypeSpec)
	c := ctx.NewCollection()
	assert.Equal(t, true, c.AddUnique(item1))
	assert.Equal(t, false, c.AddUnique(item2))
	assert.False(t, c.Empty(), "collection contains elements")
	assert.Equal(t, 1, c.Count())
	assert.Same(t, item1, c.Get(0))
	assert.Same(t, item1.TypeSpec(), c.ItemTypeSpec())
}

func TestCollectionAddUniqueDiscard(t *testing.T) {
	ctx := newTestContext(t)
	item1 := NewString("test1")
	item2 := NewString("test1")
	c := ctx.NewCollection()
	assert.Equal(t, true, c.AddUnique(item1))
	assert.Equal(t, false, c.AddUnique(item2))
	assert.False(t, c.Empty(), "collection contains elements")
	assert.Equal(t, 1, c.Count())
	assert.Same(t, item1, c.Get(0))
	assert.Same(t, item1.TypeSpec(), c.ItemTypeSpec())
}

func TestCollectionAddUniqueExistingNil(t *testing.T) {
	ctx := newTestContext(t)
	item1 := NewString("test1")
	c := ctx.NewCollection()
	assert.Equal(t, true, c.AddUnique(nil))
	assert.Equal(t, true, c.AddUnique(item1))
	assert.False(t, c.Empty(), "collection contains elements")
	assert.Equal(t, 2, c.Count())
	assert.Nil(t, c.Get(0))
	assert.Same(t, item1, c.Get(1))
	assert.Same(t, item1.TypeSpec(), c.ItemTypeSpec())
}

func TestCollectionAddAllUnique(t *testing.T) {
	ctx := newTestContext(t)
	item1 := NewString("test1")
	item2 := NewString("test2")
	item3 := NewString("test3")
	item4 := NewString("test2")
	item5 := NewString("test4")

	c1 := ctx.NewCollection()
	assert.Equal(t, true, c1.AddUnique(item1))
	assert.Equal(t, true, c1.AddUnique(item2))
	assert.Equal(t, true, c1.AddUnique(item3))

	c2 := ctx.NewCollection()
	assert.Equal(t, true, c2.AddUnique(item4))
	assert.Equal(t, true, c2.AddUnique(item5))

	c2.AddAllUnique(c1)
	if assert.Equal(t, 4, c2.Count()) {
		assert.Same(t, item4, c2.Get(0))
		assert.Same(t, item5, c2.Get(1))
		assert.Same(t, item1, c2.Get(2))
		assert.Same(t, item3, c2.Get(3))
		assert.Same(t, item1.TypeSpec(), c2.ItemTypeSpec())
	}
}

func TestCollectionAddAll(t *testing.T) {
	ctx := newTestContext(t)
	item1 := NewString("test1")
	item2 := NewString("test2")
	item3 := NewString("test3")
	item4 := NewString("test2")
	item5 := NewString("test4")

	c1 := ctx.NewCollection()
	assert.Equal(t, true, c1.AddUnique(item1))
	assert.Equal(t, true, c1.AddUnique(item2))
	assert.Equal(t, true, c1.AddUnique(item3))

	c2 := ctx.NewCollection()
	assert.Equal(t, true, c2.AddUnique(item4))
	assert.Equal(t, true, c2.AddUnique(item5))

	c2.AddAll(c1)
	if assert.Equal(t, 5, c2.Count()) {
		assert.Same(t, item4, c2.Get(0))
		assert.Same(t, item5, c2.Get(1))
		assert.Same(t, item1, c2.Get(2))
		assert.Same(t, item2, c2.Get(3))
		assert.Same(t, item3, c2.Get(4))
		assert.Same(t, item1.TypeSpec(), c2.ItemTypeSpec())
	}
}

func TestCollectionEqualTypeDiffers(t *testing.T) {
	ctx := newTestContext(t)
	assert.Equal(t, false, ctx.NewCollection().Equal(NewString("")))
	assert.Equal(t, false, ctx.NewCollection().Equivalent(NewString("")))
}

func TestCollectionEqualNil(t *testing.T) {
	ctx := newTestContext(t)
	assert.Equal(t, false, ctx.NewCollection().Equal(nil))
	assert.Equal(t, false, ctx.NewCollection().Equivalent(nil))
}

func TestCollectionEqualEmpty(t *testing.T) {
	ctx := newTestContext(t)
	assert.Equal(t, true, ctx.NewCollection().Equal(ctx.NewCollection()))
	assert.Equal(t, true, ctx.NewCollection().Equivalent(ctx.NewCollection()))
}

func TestCollectionEqual(t *testing.T) {
	ctx := newTestContext(t)
	c1 := ctx.NewCollection()
	c1.Add(NewString("test1"))
	c1.Add(NewString("test2"))
	c2 := ctx.NewCollection()
	c2.Add(NewString("test1"))
	c2.Add(NewString("test2"))
	assert.Equal(t, true, c1.Equal(c2))
	assert.Equal(t, true, c1.Equivalent(c2))
}

func TestCollectionEqualModel(t *testing.T) {
	ctx := newTestContext(t)
	c1 := ctx.NewCollection()
	c1.Add(newTestModelNode(10, false, testTypeSpec))
	c1.Add(newTestModelNode(12, false, testTypeSpec))
	c2 := ctx.NewCollection()
	c2.Add(newTestModelNode(10, false, testTypeSpec))
	c2.Add(newTestModelNode(12, false, testTypeSpec))
	assert.Equal(t, true, c1.Equal(c2))
	assert.Equal(t, true, c1.Equivalent(c2))
}

func TestCollectionEqualModelDiffers(t *testing.T) {
	ctx := newTestContext(t)
	c1 := ctx.NewCollection()
	c1.Add(newTestModelNode(10, false, testTypeSpec))
	c1.Add(newTestModelNode(12, false, testTypeSpec))
	c2 := ctx.NewCollection()
	c2.Add(newTestModelNode(10, false, testTypeSpec))
	c2.Add(newTestModelNode(14, false, testTypeSpec))
	assert.Equal(t, false, c1.Equal(c2))
	assert.Equal(t, false, c1.Equivalent(c2))
}

func TestCollectionEqualOrderDiffers(t *testing.T) {
	ctx := newTestContext(t)
	c1 := ctx.NewCollection()
	c1.Add(NewString("test1"))
	c1.Add(NewString("test2"))
	c2 := ctx.NewCollection()
	c2.Add(NewString("test2"))
	c2.Add(NewString("test1"))
	assert.Equal(t, false, c1.Equal(c2))
	assert.Equal(t, false, c1.Equivalent(c2))
}

func TestCollectionEqualCountDiffers(t *testing.T) {
	ctx := newTestContext(t)
	c1 := ctx.NewCollection()
	c1.Add(NewString("test1"))
	c2 := ctx.NewCollection()
	c2.Add(NewString("test1"))
	c2.Add(NewString("test1"))
	assert.Equal(t, false, c1.Equal(c2))
	assert.Equal(t, false, c1.Equivalent(c2))
}

func TestCollectionEquivalent(t *testing.T) {
	ctx := newTestContext(t)
	c1 := ctx.NewCollection()
	c1.Add(NewString("Test Value"))
	c2 := ctx.NewCollection()
	c2.Add(NewString("test\nvalue"))
	assert.Equal(t, false, c1.Equal(c2))
	assert.Equal(t, true, c1.Equivalent(c2))
}

func TestCollectionEquivalentModel(t *testing.T) {
	ctx := newTestContext(t)
	c1 := ctx.NewCollection()
	c1.Add(newTestModelNode(10, false, testTypeSpec))
	c1.Add(newTestModelNode(12.1, false, testTypeSpec))
	c2 := ctx.NewCollection()
	c2.Add(newTestModelNode(10, false, testTypeSpec))
	c2.Add(newTestModelNode(12.2, false, testTypeSpec))
	assert.Equal(t, false, c1.Equal(c2))
	assert.Equal(t, true, c1.Equivalent(c2))
}

func TestIsCollection(t *testing.T) {
	ctx := newTestContext(t)
	col := ctx.NewCollection()
	assert.True(t, IsCollection(col))
}

func TestIsCollectionNot(t *testing.T) {
	assert.False(t, IsCollection(True))
}

func TestCollectionContains(t *testing.T) {
	ctx := newTestContext(t)
	col := ctx.NewCollection()
	col.Add(NewInteger(10))
	col.Add(NewInteger(12))
	assert.True(t, col.Contains(NewInteger(12)))
}

func TestCollectionEmptyContains(t *testing.T) {
	ctx := newTestContext(t)
	col := ctx.NewCollection()
	assert.False(t, col.Contains(NewInteger(12)))
}

func TestCollectionContainsNot(t *testing.T) {
	ctx := newTestContext(t)
	col := ctx.NewCollection()
	col.Add(NewInteger(10))
	col.Add(NewInteger(12))
	assert.False(t, col.Contains(NewInteger(14)))
}

func TestCollectionContainsModel(t *testing.T) {
	ctx := newTestContext(t)
	col := ctx.NewCollection()
	col.Add(newTestModelNode(10.0, false, testTypeSpec))
	col.Add(newTestModelNode(12.1, false, testTypeSpec))
	assert.True(t, col.Contains(newTestModelNode(12.1, false, testTypeSpec)))
}

func TestEmptyCollectionSource(t *testing.T) {
	o := NewEmptyCollectionWithSource("abc")
	assert.Equal(t, "abc", o.Source())
}

func TestEmptyCollectionDataType(t *testing.T) {
	c := NewEmptyCollection()
	assert.Equal(t, CollectionDataType, c.DataType())
}

func TestEmptyCollectionTypeSpec(t *testing.T) {
	c := NewEmptyCollection()
	ti := c.TypeSpec()
	if assert.NotNil(t, ti, "type info expected") {
		assert.Equal(t, "System.Collection", ti.String())
		if assert.NotNil(t, ti, "base type info expected") {
			assert.Equal(t, "System.Any", ti.Base().String())
		}
	}
}

func TestEmptyCollectionEmpty(t *testing.T) {
	c := NewEmptyCollection()
	assert.True(t, c.Empty(), "new collection must be empty")
	assert.Equal(t, 0, c.Count())
	assert.Same(t, UndefinedTypeSpec, c.ItemTypeSpec())
}

func TestEmptyCollectionGet(t *testing.T) {
	c := NewEmptyCollection()
	assert.Panics(t, func() { c.Get(0) })
}

func TestEmptyCollectionEqual(t *testing.T) {
	ctx := newTestContext(t)
	c1 := NewEmptyCollection()
	c2 := NewCollection(ctx.ModelAdapter())
	assert.True(t, c1.Equal(c2))
	assert.True(t, c1.Equivalent(c2))
}

func TestEmptyCollectionEqualNot(t *testing.T) {
	ctx := newTestContext(t)
	c1 := NewEmptyCollection()
	c2 := NewCollection(ctx.ModelAdapter())
	c2.Add(NewString("test"))
	assert.False(t, c1.Equal(c2))
	assert.False(t, c1.Equivalent(c2))
}

func TestEmptyCollectionNoCol(t *testing.T) {
	c1 := NewEmptyCollection()
	assert.False(t, c1.Equal(NewString("")))
	assert.False(t, c1.Equivalent(NewString("")))
}

func TestEmptyCollectionContains(t *testing.T) {
	c := NewEmptyCollection()
	assert.False(t, c.Contains(False))
}
