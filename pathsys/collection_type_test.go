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

func TestCollectionDataType(t *testing.T) {
	ctx := newTestContext(t)
	c := ctx.NewCollection()
	assert.Equal(t, CollectionDataType, c.DataType())
}

func TestCollectionTypeInfo(t *testing.T) {
	ctx := newTestContext(t)
	c := ctx.NewCollection()
	ti := c.TypeInfo()
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
	assert.Same(t, UndefinedTypeInfo, c.ItemTypeInfo())
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
	assert.Same(t, item1.TypeInfo(), c.ItemTypeInfo())
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
	assert.Same(t, anyTypeInfo, c.ItemTypeInfo())
}

func TestCollectionAddGetModel(t *testing.T) {
	ctx := newTestContext(t)
	item1 := newTestModelNode(10, false, testTypeInfo)
	item2 := newTestModelNode(11, false, testTypeInfo)
	c := ctx.NewCollection()
	c.Add(item1)
	c.Add(item2)
	assert.False(t, c.Empty(), "collection contains elements")
	assert.Equal(t, 2, c.Count())
	assert.Same(t, item1, c.Get(0))
	assert.Same(t, item2, c.Get(1))
	assert.Same(t, testTypeInfo, c.ItemTypeInfo())
}

func TestCollectionAddGetConvertedModel(t *testing.T) {
	ctx := newTestContext(t)
	item1 := newTestModelNode(10.1, true, testTypeInfo)
	item2 := newTestModelNode(12.1, true, testTypeInfo)
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
	assert.Same(t, DecimalTypeInfo, c.ItemTypeInfo())
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
	assert.Same(t, item1.TypeInfo(), c.ItemTypeInfo())
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
	assert.Same(t, item1.TypeInfo(), c.ItemTypeInfo())
}

func TestCollectionAddUniqueDupNil(t *testing.T) {
	ctx := newTestContext(t)
	c := ctx.NewCollection()
	assert.Equal(t, true, c.AddUnique(nil))
	assert.Equal(t, false, c.AddUnique(nil))
	assert.False(t, c.Empty(), "collection contains elements")
	assert.Equal(t, 1, c.Count())
	assert.Nil(t, c.Get(0))
	assert.Same(t, UndefinedTypeInfo, c.ItemTypeInfo())
}

func TestCollectionAddUniqueModel(t *testing.T) {
	ctx := newTestContext(t)
	item1 := NewString("test1")
	item2 := newTestModelNode(10, false, testTypeInfo)
	item3 := NewString("test1")
	item4 := newTestModelNode(10, false, testTypeInfo)
	c := ctx.NewCollection()
	assert.Equal(t, true, c.AddUnique(item1))
	assert.Equal(t, true, c.AddUnique(item2))
	assert.Equal(t, false, c.AddUnique(item3))
	assert.Equal(t, false, c.AddUnique(item4))
	assert.False(t, c.Empty(), "collection contains elements")
	assert.Equal(t, 2, c.Count())
	assert.Same(t, item1, c.Get(0))
	assert.Same(t, item2, c.Get(1))
	assert.Same(t, UndefinedTypeInfo, c.ItemTypeInfo())
}

func TestCollectionAddUniqueConvertedModel(t *testing.T) {
	ctx := newTestContext(t)
	item1 := NewDecimalFloat64(10)
	item2 := newTestModelNode(10, true, testTypeInfo)
	c := ctx.NewCollection()
	assert.Equal(t, true, c.AddUnique(item1))
	assert.Equal(t, false, c.AddUnique(item2))
	assert.False(t, c.Empty(), "collection contains elements")
	assert.Equal(t, 1, c.Count())
	assert.Same(t, item1, c.Get(0))
	assert.Same(t, item1.TypeInfo(), c.ItemTypeInfo())
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
	assert.Same(t, item1.TypeInfo(), c.ItemTypeInfo())
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
	assert.Same(t, item1.TypeInfo(), c.ItemTypeInfo())
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
		assert.Same(t, item1.TypeInfo(), c2.ItemTypeInfo())
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
		assert.Same(t, item1.TypeInfo(), c2.ItemTypeInfo())
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
	c1.Add(newTestModelNode(10, false, testTypeInfo))
	c1.Add(newTestModelNode(12, false, testTypeInfo))
	c2 := ctx.NewCollection()
	c2.Add(newTestModelNode(10, false, testTypeInfo))
	c2.Add(newTestModelNode(12, false, testTypeInfo))
	assert.Equal(t, true, c1.Equal(c2))
	assert.Equal(t, true, c1.Equivalent(c2))
}

func TestCollectionEqualModelDiffers(t *testing.T) {
	ctx := newTestContext(t)
	c1 := ctx.NewCollection()
	c1.Add(newTestModelNode(10, false, testTypeInfo))
	c1.Add(newTestModelNode(12, false, testTypeInfo))
	c2 := ctx.NewCollection()
	c2.Add(newTestModelNode(10, false, testTypeInfo))
	c2.Add(newTestModelNode(14, false, testTypeInfo))
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
	c1.Add(newTestModelNode(10, false, testTypeInfo))
	c1.Add(newTestModelNode(12.1, false, testTypeInfo))
	c2 := ctx.NewCollection()
	c2.Add(newTestModelNode(10, false, testTypeInfo))
	c2.Add(newTestModelNode(12.2, false, testTypeInfo))
	assert.Equal(t, false, c1.Equal(c2))
	assert.Equal(t, true, c1.Equivalent(c2))
}