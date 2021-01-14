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

var EmptyCol = NewEmptyCol()

var colTypeSpec = UndefinedTypeSpec
var colTypeInfo = NewListTypeInfo(nil)

type ColAccessor interface {
	AnyAccessor
	ItemTypeSpec() TypeSpecAccessor
	Empty() bool
	Count() int
	Get(i int) interface{}
	Contains(item interface{}) bool
}

type ColModifier interface {
	ColAccessor
	Add(item interface{})
	AddUnique(item interface{}) bool
	AddAll(collection ColAccessor) int
	AddAllUnique(collection ColAccessor) int
}

type baseColType struct {
	baseAnyType
	adapter      ModelAdapter
	itemTypeSpec TypeSpecAccessor
	items        []interface{}
}

type colType struct {
	baseColType
}

type sysArrayCol struct {
	baseColType
}

type emptyCol struct {
	baseAnyType
}

type colDelegate struct {
	delegate ColAccessor
}

func IsCol(node interface{}) bool {
	if _, ok := node.(ColAccessor); ok {
		return true
	}
	return false
}

func NewCol(adapter ModelAdapter) ColModifier {
	return NewColWithSource(adapter, nil)
}

func NewColWithItem(adapter ModelAdapter, item interface{}) ColModifier {
	c := newCollection(adapter, nil, nil)
	c.items = make([]interface{}, 1)
	c.items[0] = item
	return c
}

func NewColWithSource(adapter ModelAdapter, source interface{}) ColModifier {
	return newCollection(adapter, nil, source)
}

func NewColWithSpec(adapter ModelAdapter, itemTypeSpec TypeSpecAccessor) ColModifier {
	return newCollection(adapter, itemTypeSpec, nil)
}

func NewEmptyCol() ColAccessor {
	return NewEmptyColWithSource(nil)
}

func NewEmptyColWithSource(source interface{}) ColAccessor {
	return &emptyCol{
		baseAnyType: baseAnyType{
			source: source,
		},
	}
}

func NewColDelegate(col ColAccessor) ColAccessor {
	return &colDelegate{
		col,
	}
}

func newCollection(adapter ModelAdapter, itemTypeSpec TypeSpecAccessor, source interface{}) *colType {
	if adapter == nil {
		panic("no adapter has been specified")
	}
	return &colType{
		baseColType: baseColType{
			baseAnyType: baseAnyType{
				source: source,
			},
			adapter:      adapter,
			itemTypeSpec: itemTypeSpec,
		},
	}
}

func NewSysArrayCol(itemTypeSpec TypeSpecAccessor, items []interface{}) ColAccessor {
	return &sysArrayCol{
		baseColType{
			itemTypeSpec: itemTypeSpec,
			items:        items,
		},
	}
}

func (c *baseColType) DataType() DataTypes {
	return ColDataType
}

func (c *baseColType) TypeSpec() TypeSpecAccessor {
	return colTypeSpec
}

func (c *baseColType) TypeInfo() TypeInfoAccessor {
	if c.itemTypeSpec == nil {
		return colTypeInfo
	}
	return NewListTypeInfo(NewString(c.itemTypeSpec.String()))
}

func (c *baseColType) ItemTypeSpec() TypeSpecAccessor {
	return c.itemTypeSpec
}

func (c *baseColType) Empty() bool {
	return c.Count() == 0
}

func (c *baseColType) Count() int {
	return len(c.items)
}

func (c *baseColType) Get(i int) interface{} {
	if c.items == nil {
		panic("collection is empty")
	}
	return c.items[i]
}

func (c *baseColType) Contains(item interface{}) bool {
	if c.items == nil {
		return false
	}

	for _, o := range c.items {
		if ModelEqual(c.adapter, item, o) {
			return true
		}
	}
	return false
}

func (c *baseColType) Equal(item interface{}) bool {
	if o, ok := item.(ColAccessor); !ok {
		return false
	} else {
		return c.Count() == o.Count() &&
			colDeepEqual(c.adapter, c, o)
	}
}

func (c *baseColType) Equivalent(item interface{}) bool {
	if o, ok := item.(ColAccessor); !ok {
		return false
	} else {
		return c.Count() == o.Count() &&
			colDeepEquivalent(c.adapter, c, o)
	}
}

func colDeepEqual(adapter ModelAdapter, c1 ColAccessor, c2 ColAccessor) bool {
	count := c1.Count()
	for i := 0; i < count; i++ {
		if !ModelEqual(adapter, c1.Get(i), c2.Get(i)) {
			return false
		}
	}
	return true
}

func colDeepEquivalent(adapter ModelAdapter, c1 ColAccessor, c2 ColAccessor) bool {
	count := c1.Count()
	for i := 0; i < count; i++ {
		if !ModelEquivalent(adapter, c1.Get(i), c2.Get(i)) {
			return false
		}
	}
	return true
}

func (c *colType) Add(item interface{}) {
	c.add(item)
}

func (c *colType) add(item interface{}) {
	if c.items == nil {
		c.items = make([]interface{}, 0)
	}
	c.items = append(c.items, item)
}

func (c *colType) AddUnique(item interface{}) bool {
	if c.items == nil {
		c.add(item)
		return true
	}

	if item == nil {
		for _, o := range c.items {
			if o == nil {
				return false
			}
		}
	} else {
		adapter := c.adapter
		for _, o := range c.items {
			if o != nil && ModelEqual(adapter, item, o) {
				return false
			}
		}
	}

	c.items = append(c.items, item)
	return true
}

func (c *colType) AddAll(collection ColAccessor) int {
	count := collection.Count()
	for i := 0; i < count; i++ {
		c.add(collection.Get(i))
	}
	return count
}

func (c *colType) AddAllUnique(collection ColAccessor) int {
	added := 0
	count := collection.Count()
	for i := 0; i < count; i++ {
		a := c.AddUnique(collection.Get(i))
		if a {
			added = added + 1
		}
	}
	return added
}

func (c *emptyCol) DataType() DataTypes {
	return ColDataType
}

func (c *emptyCol) TypeSpec() TypeSpecAccessor {
	return colTypeSpec
}

func (c *emptyCol) TypeInfo() TypeInfoAccessor {
	return colTypeInfo
}

func (c *emptyCol) Equal(item interface{}) bool {
	if o, ok := item.(ColAccessor); !ok {
		return false
	} else {
		return o.Empty()
	}
}

func (c *emptyCol) Equivalent(item interface{}) bool {
	return c.Equal(item)
}

func (c *emptyCol) ItemTypeSpec() TypeSpecAccessor {
	return UndefinedTypeSpec
}

func (c *emptyCol) Empty() bool {
	return true
}

func (c *emptyCol) Count() int {
	return 0
}

func (c *emptyCol) Get(int) interface{} {
	panic("cannot get an item from an empty collection")
}

func (c *emptyCol) Contains(interface{}) bool {
	return false
}

func (c *colDelegate) DataType() DataTypes {
	return c.delegate.DataType()
}

func (c *colDelegate) TypeSpec() TypeSpecAccessor {
	return c.delegate.TypeSpec()
}

func (c *colDelegate) TypeInfo() TypeInfoAccessor {
	return c.delegate.TypeInfo()
}

func (c *colDelegate) Source() interface{} {
	return c.delegate.Source()
}

func (c *colDelegate) Equal(node interface{}) bool {
	return c.delegate.Equal(node)
}

func (c *colDelegate) Equivalent(node interface{}) bool {
	return c.delegate.Equivalent(node)
}

func (c *colDelegate) ItemTypeSpec() TypeSpecAccessor {
	return c.delegate.ItemTypeSpec()
}

func (c *colDelegate) Empty() bool {
	return c.delegate.Empty()
}

func (c *colDelegate) Count() int {
	return c.delegate.Count()
}

func (c *colDelegate) Get(i int) interface{} {
	return c.delegate.Get(i)
}

func (c *colDelegate) Contains(item interface{}) bool {
	return c.delegate.Contains(item)
}
