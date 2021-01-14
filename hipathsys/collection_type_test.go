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

func TestColSource(t *testing.T) {
	ctx := newTestContext(t)
	o := NewColWithSource(ctx.ModelAdapter(), "abc")
	assert.Equal(t, "abc", o.Source())
}

func TestColDataType(t *testing.T) {
	ctx := newTestContext(t)
	c := ctx.NewCol()
	assert.Equal(t, ColDataType, c.DataType())
}

func TestColTypeSpec(t *testing.T) {
	ctx := newTestContext(t)
	c := ctx.NewCol()
	ti := c.TypeSpec()
	if assert.NotNil(t, ti, "type info expected") {
		assert.Equal(t, "", ti.String())
		assert.Nil(t, ti.Base(), "no base type expected")
		assert.True(t, ti.Anonymous())
	}
}

func TestColTypeInfo(t *testing.T) {
	ctx := newTestContext(t)
	c := ctx.NewCol()
	i := c.TypeInfo()
	if assert.Implements(t, (*ListTypeInfoAccessor)(nil), i) {
		a := i.(ListTypeInfoAccessor)
		assert.Equal(t, NewString("System"), a.Namespace())
		assert.Nil(t, a.ElementType(), "item type spec has not been set")
	}
}

func TestColTypeInfoWithItemTypeSpec(t *testing.T) {
	ctx := newTestContext(t)
	c := NewColWithSpec(ctx.ModelAdapter(), DecimalTypeSpec)
	i := c.TypeInfo()
	if assert.Implements(t, (*ListTypeInfoAccessor)(nil), i) {
		a := i.(ListTypeInfoAccessor)
		assert.Equal(t, NewString("System"), a.Namespace())
		assert.Equal(t, NewString("System.Decimal"), a.ElementType())
	}
}

func TestNewColNilCtx(t *testing.T) {
	assert.Panics(t, func() { NewCol(nil) })
}

func TestNewCol(t *testing.T) {
	ctx := newTestContext(t)
	c := ctx.NewCol()
	assert.True(t, c.Empty(), "new col must be empty")
	assert.Equal(t, 0, c.Count())
	assert.Nil(t, c.ItemTypeSpec())
}

func TestNewColWithItemTypeSpec(t *testing.T) {
	ctx := newTestContext(t)
	c := NewColWithSpec(ctx.ModelAdapter(), StringTypeSpec)
	assert.True(t, c.Empty(), "new col must be empty")
	assert.Equal(t, 0, c.Count())
	assert.Same(t, StringTypeSpec, c.ItemTypeSpec())
}

func TestColWithItemTypeSpec(t *testing.T) {
	ctx := newTestContext(t)
	c := NewColWithSpec(ctx.ModelAdapter(), StringTypeSpec)
	c.Add(NewInteger(10))
	assert.False(t, c.Empty(), "col is not empty")
	assert.Equal(t, 1, c.Count())
	assert.Same(t, StringTypeSpec, c.ItemTypeSpec())
}

func TestNewColGetEmpty(t *testing.T) {
	ctx := newTestContext(t)
	c := ctx.NewCol()
	assert.Panics(t, func() { c.Get(0) })
}

func TestColAddGet(t *testing.T) {
	ctx := newTestContext(t)
	item1 := NewString("test1")
	item2 := NewString("test2")
	c := ctx.NewCol()
	c.Add(item1)
	c.Add(item2)
	assert.False(t, c.Empty(), "col contains elements")
	assert.Equal(t, 2, c.Count())
	assert.Same(t, item1, c.Get(0))
	assert.Same(t, item2, c.Get(1))
	assert.Nil(t, c.ItemTypeSpec())
}

func TestColAddGetModel(t *testing.T) {
	ctx := newTestContext(t)
	item1 := newTestModelNode(10, false, testTypeSpec)
	item2 := newTestModelNode(11, false, testTypeSpec)
	c := ctx.NewCol()
	c.Add(item1)
	c.Add(item2)
	assert.False(t, c.Empty(), "col contains elements")
	assert.Equal(t, 2, c.Count())
	assert.Same(t, item1, c.Get(0))
	assert.Same(t, item2, c.Get(1))
	assert.Nil(t, c.ItemTypeSpec())
}

func TestNewColWithModelItem(t *testing.T) {
	ctx := newTestContext(t)
	item1 := newTestModelNode(10, false, testTypeSpec)
	c := ctx.NewColWithItem(item1)
	c.Add(item1)
	assert.False(t, c.Empty(), "col contains elements")
	assert.Equal(t, 2, c.Count())
	assert.Same(t, item1, c.Get(0))
	assert.Nil(t, c.ItemTypeSpec())
}

func TestColAddNil(t *testing.T) {
	ctx := newTestContext(t)
	item1 := NewString("test1")
	c := ctx.NewCol()
	c.Add(item1)
	c.Add(nil)
	assert.False(t, c.Empty(), "col contains elements")
	assert.Equal(t, 2, c.Count())
	assert.Same(t, item1, c.Get(0))
	assert.Nil(t, c.Get(1))
	assert.Nil(t, c.ItemTypeSpec())
}

func TestColAddUniqueAll(t *testing.T) {
	ctx := newTestContext(t)
	item1 := NewString("test1")
	item2 := NewString("test2")
	c := ctx.NewCol()
	a := c.AddUnique(item1)
	assert.Equal(t, true, a)
	a = c.AddUnique(item2)
	assert.Equal(t, true, a)
	assert.False(t, c.Empty(), "col contains elements")
	assert.Equal(t, 2, c.Count())
	assert.Same(t, item1, c.Get(0))
	assert.Same(t, item2, c.Get(1))
	assert.Nil(t, c.ItemTypeSpec())
}

func TestColAddUniqueDupNil(t *testing.T) {
	ctx := newTestContext(t)
	c := ctx.NewCol()
	a := c.AddUnique(nil)
	assert.Equal(t, true, a)
	a = c.AddUnique(nil)
	assert.Equal(t, false, a)
	assert.False(t, c.Empty(), "col contains elements")
	assert.Equal(t, 1, c.Count())
	assert.Nil(t, c.Get(0))
	assert.Nil(t, c.ItemTypeSpec())
}

func TestColAddUniqueModel(t *testing.T) {
	ctx := newTestContext(t)
	item1 := NewString("test1")
	item2 := newTestModelNode(10, false, testTypeSpec)
	item3 := NewString("test1")
	item4 := newTestModelNode(10, false, testTypeSpec)
	c := ctx.NewCol()
	a := c.AddUnique(item1)
	assert.Equal(t, true, a)
	a = c.AddUnique(item2)
	assert.Equal(t, true, a)
	a = c.AddUnique(item3)
	assert.Equal(t, false, a)
	a = c.AddUnique(item4)
	assert.Equal(t, false, a)
	assert.False(t, c.Empty(), "col contains elements")
	assert.Equal(t, 2, c.Count())
	assert.Same(t, item1, c.Get(0))
	assert.Same(t, item2, c.Get(1))
	assert.Nil(t, c.ItemTypeSpec())
}

func TestColAddUniqueDiscard(t *testing.T) {
	ctx := newTestContext(t)
	item1 := NewString("test1")
	item2 := NewString("test1")
	c := ctx.NewCol()
	a := c.AddUnique(item1)
	assert.Equal(t, true, a)
	a = c.AddUnique(item2)
	assert.Equal(t, false, a)
	assert.False(t, c.Empty(), "col contains elements")
	assert.Equal(t, 1, c.Count())
	assert.Same(t, item1, c.Get(0))
	assert.Nil(t, c.ItemTypeSpec())
}

func TestColAddUniqueExistingNil(t *testing.T) {
	ctx := newTestContext(t)
	item1 := NewString("test1")
	c := ctx.NewCol()
	a := c.AddUnique(nil)
	assert.Equal(t, true, a)
	a = c.AddUnique(item1)
	assert.Equal(t, true, a)
	assert.False(t, c.Empty(), "col contains elements")
	assert.Equal(t, 2, c.Count())
	assert.Nil(t, c.Get(0))
	assert.Same(t, item1, c.Get(1))
	assert.Nil(t, c.ItemTypeSpec())
}

func TestColAddAllUnique(t *testing.T) {
	ctx := newTestContext(t)
	item1 := NewString("test1")
	item2 := NewString("test2")
	item3 := NewString("test3")
	item4 := NewString("test2")
	item5 := NewString("test4")

	c1 := ctx.NewCol()
	a := c1.AddUnique(item1)
	assert.Equal(t, true, a)
	a = c1.AddUnique(item2)
	assert.Equal(t, true, a)
	a = c1.AddUnique(item3)
	assert.Equal(t, true, a)

	c2 := ctx.NewCol()
	a = c2.AddUnique(item4)
	assert.Equal(t, true, a)
	a = c2.AddUnique(item5)
	assert.Equal(t, true, a)

	count := c2.AddAllUnique(c1)
	assert.Equal(t, 2, count)
	if assert.Equal(t, 4, c2.Count()) {
		assert.Same(t, item4, c2.Get(0))
		assert.Same(t, item5, c2.Get(1))
		assert.Same(t, item1, c2.Get(2))
		assert.Same(t, item3, c2.Get(3))
	}
}

func TestColAddAll(t *testing.T) {
	ctx := newTestContext(t)
	item1 := NewString("test1")
	item2 := NewString("test2")
	item3 := NewString("test3")
	item4 := NewString("test2")
	item5 := NewString("test4")

	c1 := ctx.NewCol()
	a := c1.AddUnique(item1)
	assert.Equal(t, true, a)
	a = c1.AddUnique(item2)
	assert.Equal(t, true, a)
	a = c1.AddUnique(item3)
	assert.Equal(t, true, a)

	c2 := ctx.NewCol()
	a = c2.AddUnique(item4)
	assert.Equal(t, true, a)
	a = c2.AddUnique(item5)
	assert.Equal(t, true, a)

	count := c2.AddAll(c1)
	assert.Equal(t, 3, count)
	if assert.Equal(t, 5, c2.Count()) {
		assert.Same(t, item4, c2.Get(0))
		assert.Same(t, item5, c2.Get(1))
		assert.Same(t, item1, c2.Get(2))
		assert.Same(t, item2, c2.Get(3))
		assert.Same(t, item3, c2.Get(4))
	}
}

func TestColEqualTypeDiffers(t *testing.T) {
	ctx := newTestContext(t)
	assert.Equal(t, false, ctx.NewCol().Equal(NewString("")))
	assert.Equal(t, false, ctx.NewCol().Equivalent(NewString("")))
}

func TestColEqualNil(t *testing.T) {
	ctx := newTestContext(t)
	assert.Equal(t, false, ctx.NewCol().Equal(nil))
	assert.Equal(t, false, ctx.NewCol().Equivalent(nil))
}

func TestColEqualEmpty(t *testing.T) {
	ctx := newTestContext(t)
	assert.Equal(t, true, ctx.NewCol().Equal(ctx.NewCol()))
	assert.Equal(t, true, ctx.NewCol().Equivalent(ctx.NewCol()))
}

func TestColEqual(t *testing.T) {
	ctx := newTestContext(t)
	c1 := ctx.NewCol()
	c1.Add(NewString("test1"))
	c1.Add(NewString("test2"))
	c2 := ctx.NewCol()
	c2.Add(NewString("test1"))
	c2.Add(NewString("test2"))
	assert.Equal(t, true, c1.Equal(c2))
	assert.Equal(t, true, c1.Equivalent(c2))
}

func TestColEqualModel(t *testing.T) {
	ctx := newTestContext(t)
	c1 := ctx.NewCol()
	c1.Add(newTestModelNode(10, false, testTypeSpec))
	c1.Add(newTestModelNode(12, false, testTypeSpec))
	c2 := ctx.NewCol()
	c2.Add(newTestModelNode(10, false, testTypeSpec))
	c2.Add(newTestModelNode(12, false, testTypeSpec))
	assert.Equal(t, true, c1.Equal(c2))
	assert.Equal(t, true, c1.Equivalent(c2))
}

func TestColEqualModelDiffers(t *testing.T) {
	ctx := newTestContext(t)
	c1 := ctx.NewCol()
	c1.Add(newTestModelNode(10, false, testTypeSpec))
	c1.Add(newTestModelNode(12, false, testTypeSpec))
	c2 := ctx.NewCol()
	c2.Add(newTestModelNode(10, false, testTypeSpec))
	c2.Add(newTestModelNode(14, false, testTypeSpec))
	assert.Equal(t, false, c1.Equal(c2))
	assert.Equal(t, false, c1.Equivalent(c2))
}

func TestColEqualOrderDiffers(t *testing.T) {
	ctx := newTestContext(t)
	c1 := ctx.NewCol()
	c1.Add(NewString("test1"))
	c1.Add(NewString("test2"))
	c2 := ctx.NewCol()
	c2.Add(NewString("test2"))
	c2.Add(NewString("test1"))
	assert.Equal(t, false, c1.Equal(c2))
	assert.Equal(t, false, c1.Equivalent(c2))
}

func TestColEqualCountDiffers(t *testing.T) {
	ctx := newTestContext(t)
	c1 := ctx.NewCol()
	c1.Add(NewString("test1"))
	c2 := ctx.NewCol()
	c2.Add(NewString("test1"))
	c2.Add(NewString("test1"))
	assert.Equal(t, false, c1.Equal(c2))
	assert.Equal(t, false, c1.Equivalent(c2))
}

func TestColEquivalent(t *testing.T) {
	ctx := newTestContext(t)
	c1 := ctx.NewCol()
	c1.Add(NewString("Test Value"))
	c2 := ctx.NewCol()
	c2.Add(NewString("test\nvalue"))
	assert.Equal(t, false, c1.Equal(c2))
	assert.Equal(t, true, c1.Equivalent(c2))
}

func TestColEquivalentModel(t *testing.T) {
	ctx := newTestContext(t)
	c1 := ctx.NewCol()
	c1.Add(newTestModelNode(10, false, testTypeSpec))
	c1.Add(newTestModelNode(12.1, false, testTypeSpec))
	c2 := ctx.NewCol()
	c2.Add(newTestModelNode(10, false, testTypeSpec))
	c2.Add(newTestModelNode(12.2, false, testTypeSpec))
	assert.Equal(t, false, c1.Equal(c2))
	assert.Equal(t, true, c1.Equivalent(c2))
}

func TestIsCol(t *testing.T) {
	ctx := newTestContext(t)
	col := ctx.NewCol()
	assert.True(t, IsCol(col))
}

func TestIsColNot(t *testing.T) {
	assert.False(t, IsCol(True))
}

func TestColContains(t *testing.T) {
	ctx := newTestContext(t)
	col := ctx.NewCol()
	col.Add(NewInteger(10))
	col.Add(NewInteger(12))
	assert.True(t, col.Contains(NewInteger(12)))
}

func TestColEmptyContains(t *testing.T) {
	ctx := newTestContext(t)
	col := ctx.NewCol()
	assert.False(t, col.Contains(NewInteger(12)))
}

func TestColContainsNot(t *testing.T) {
	ctx := newTestContext(t)
	col := ctx.NewCol()
	col.Add(NewInteger(10))
	col.Add(NewInteger(12))
	assert.False(t, col.Contains(NewInteger(14)))
}

func TestColContainsModel(t *testing.T) {
	ctx := newTestContext(t)
	col := ctx.NewCol()
	col.Add(newTestModelNode(10.0, false, testTypeSpec))
	col.Add(newTestModelNode(12.1, false, testTypeSpec))
	assert.True(t, col.Contains(newTestModelNode(12.1, false, testTypeSpec)))
}

func TestEmptyColSource(t *testing.T) {
	o := NewEmptyColWithSource("abc")
	assert.Equal(t, "abc", o.Source())
}

func TestEmptyColDataType(t *testing.T) {
	c := NewEmptyCol()
	assert.Equal(t, ColDataType, c.DataType())
}

func TestEmptyColTypeSpec(t *testing.T) {
	c := NewEmptyCol()
	ti := c.TypeSpec()
	if assert.NotNil(t, ti, "type info expected") {
		assert.Equal(t, "", ti.String())
		assert.Nil(t, ti.Base(), "no base type expected")
		assert.True(t, ti.Anonymous())
	}
}

func TestEmptyColTypeInfo(t *testing.T) {
	c := NewEmptyCol()
	i := c.TypeInfo()
	if assert.Implements(t, (*ListTypeInfoAccessor)(nil), i) {
		a := i.(ListTypeInfoAccessor)
		assert.Equal(t, NewString("System"), a.Namespace())
		assert.Nil(t, a.ElementType(), "item type spec has not been set")
	}
}

func TestEmptyColEmpty(t *testing.T) {
	c := NewEmptyCol()
	assert.True(t, c.Empty(), "new col must be empty")
	assert.Equal(t, 0, c.Count())
	assert.Same(t, UndefinedTypeSpec, c.ItemTypeSpec())
}

func TestEmptyColGet(t *testing.T) {
	c := NewEmptyCol()
	assert.Panics(t, func() { c.Get(0) })
}

func TestEmptyColEqual(t *testing.T) {
	ctx := newTestContext(t)
	c1 := NewEmptyCol()
	c2 := NewCol(ctx.ModelAdapter())
	assert.True(t, c1.Equal(c2))
	assert.True(t, c1.Equivalent(c2))
}

func TestEmptyColEqualNot(t *testing.T) {
	ctx := newTestContext(t)
	c1 := NewEmptyCol()
	c2 := NewCol(ctx.ModelAdapter())
	c2.Add(NewString("test"))
	assert.False(t, c1.Equal(c2))
	assert.False(t, c1.Equivalent(c2))
}

func TestEmptyColNoCol(t *testing.T) {
	c1 := NewEmptyCol()
	assert.False(t, c1.Equal(NewString("")))
	assert.False(t, c1.Equivalent(NewString("")))
}

func TestEmptyColContains(t *testing.T) {
	c := NewEmptyCol()
	assert.False(t, c.Contains(False))
}

func TestColDelegateSource(t *testing.T) {
	ctx := newTestContext(t)
	o := NewColWithSource(ctx.ModelAdapter(), "abc")
	d := NewColDelegate(o)
	assert.Equal(t, "abc", d.Source())
}

func TestColDelegateDataType(t *testing.T) {
	ctx := newTestContext(t)
	c := ctx.NewCol()
	d := NewColDelegate(c)
	assert.Equal(t, ColDataType, d.DataType())
}

func TestColDelegateTypeSpec(t *testing.T) {
	ctx := newTestContext(t)
	c := ctx.NewCol()
	d := NewColDelegate(c)
	ti := d.TypeSpec()
	if assert.NotNil(t, ti, "type info expected") {
		assert.Equal(t, "", ti.String())
		assert.Nil(t, ti.Base(), "no base type expected")
		assert.True(t, ti.Anonymous())
	}
}

func TestColDelegateTypeInfo(t *testing.T) {
	ctx := newTestContext(t)
	c := ctx.NewCol()
	d := NewColDelegate(c)
	i := d.TypeInfo()
	if assert.Implements(t, (*ListTypeInfoAccessor)(nil), i) {
		a := i.(ListTypeInfoAccessor)
		assert.Equal(t, NewString("System"), a.Namespace())
		assert.Nil(t, a.ElementType(), "item type spec has not been set")
	}
}

func TestColDelegateTypeInfoWithItemTypeSpec(t *testing.T) {
	ctx := newTestContext(t)
	c := NewColWithSpec(ctx.ModelAdapter(), DecimalTypeSpec)
	d := NewColDelegate(c)
	i := d.TypeInfo()
	if assert.Implements(t, (*ListTypeInfoAccessor)(nil), i) {
		a := i.(ListTypeInfoAccessor)
		assert.Equal(t, NewString("System"), a.Namespace())
		assert.Equal(t, NewString("System.Decimal"), a.ElementType())
	}
}

func TestColDelegateGet(t *testing.T) {
	ctx := newTestContext(t)
	item1 := NewString("test1")
	item2 := NewString("test2")
	c := ctx.NewCol()
	c.Add(item1)
	c.Add(item2)
	d := NewColDelegate(c)
	assert.False(t, d.Empty(), "col contains elements")
	assert.Equal(t, 2, d.Count())
	assert.Same(t, item1, d.Get(0))
	assert.Same(t, item2, d.Get(1))
	assert.Nil(t, d.ItemTypeSpec())
}

func TestColDelegateContains(t *testing.T) {
	ctx := newTestContext(t)
	col := ctx.NewCol()
	col.Add(NewInteger(10))
	col.Add(NewInteger(12))
	d := NewColDelegate(col)
	assert.True(t, d.Contains(NewInteger(12)))
}

func TestColDelegateContainsNot(t *testing.T) {
	ctx := newTestContext(t)
	col := ctx.NewCol()
	col.Add(NewInteger(10))
	col.Add(NewInteger(12))
	d := NewColDelegate(col)
	assert.False(t, d.Contains(NewInteger(11)))
}

func TestColDelegateEqual(t *testing.T) {
	ctx := newTestContext(t)
	c1 := ctx.NewCol()
	c1.Add(NewString("test1"))
	c1.Add(NewString("test2"))
	d := NewColDelegate(c1)
	c2 := ctx.NewCol()
	c2.Add(NewString("test1"))
	c2.Add(NewString("test2"))
	assert.Equal(t, true, d.Equal(c2))
	assert.Equal(t, true, d.Equivalent(c2))
}

func TestColDelegateEquivalent(t *testing.T) {
	ctx := newTestContext(t)
	c1 := ctx.NewCol()
	c1.Add(NewString("Test Value"))
	d := NewColDelegate(c1)
	c2 := ctx.NewCol()
	c2.Add(NewString("test\nvalue"))
	assert.Equal(t, false, d.Equal(c2))
	assert.Equal(t, true, d.Equivalent(c2))
}

func TestColDelegateEqualNot(t *testing.T) {
	ctx := newTestContext(t)
	c1 := ctx.NewCol()
	c1.Add(NewString("test1"))
	d := NewColDelegate(c1)
	c2 := ctx.NewCol()
	c2.Add(NewString("test2"))
	assert.Equal(t, false, d.Equal(c2))
	assert.Equal(t, false, d.Equivalent(c2))
}
