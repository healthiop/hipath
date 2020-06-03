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

var EmptyCollection = NewEmptyCollection()

var collectionTypeInfo = newAnyTypeInfo("Collection")

type collectionType struct {
	baseAnyType
	adapter      ModelAdapter
	itemTypeInfo TypeInfoAccessor
	items        []interface{}
}

type CollectionAccessor interface {
	AnyAccessor
	ItemTypeInfo() TypeInfoAccessor
	Empty() bool
	Count() int
	Get(i int) interface{}
	Contains(item interface{}) bool
}

func IsCollection(node interface{}) bool {
	if _, ok := node.(CollectionAccessor); ok {
		return true
	}
	return false
}

type CollectionModifier interface {
	CollectionAccessor
	Add(item interface{})
	AddUnique(item interface{}) bool
	AddAll(collection CollectionAccessor) int
	AddAllUnique(collection CollectionAccessor) int
}

func NewCollection(adapter ModelAdapter) CollectionModifier {
	return NewCollectionWithSource(adapter, nil)
}

func NewCollectionWithItem(adapter ModelAdapter, item interface{}) CollectionModifier {
	c := newCollection(adapter, nil)
	c.items = make([]interface{}, 1)
	c.items[0] = c.prepareItem(item, true)
	return c
}

func NewCollectionWithSource(adapter ModelAdapter, source interface{}) CollectionModifier {
	return newCollection(adapter, source)
}

func newCollection(adapter ModelAdapter, source interface{}) *collectionType {
	if adapter == nil {
		panic("no adapter has been specified")
	}
	return &collectionType{
		baseAnyType: baseAnyType{
			source: source,
		},
		adapter: adapter,
	}
}

func (c *collectionType) DataType() DataTypes {
	return CollectionDataType
}

func (c *collectionType) TypeInfo() TypeInfoAccessor {
	return collectionTypeInfo
}

func (c *collectionType) ItemTypeInfo() TypeInfoAccessor {
	if c.itemTypeInfo == nil {
		return UndefinedTypeInfo
	}
	return c.itemTypeInfo
}

func (c *collectionType) Empty() bool {
	return c.Count() == 0
}

func (c *collectionType) Count() int {
	return len(c.items)
}

func (c *collectionType) Get(i int) interface{} {
	if c.items == nil {
		panic("collection is empty")
	}
	return c.items[i]
}

func (c *collectionType) Contains(item interface{}) bool {
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

func (c *collectionType) Add(item interface{}) {
	c.add(item, true)
}

func (c *collectionType) add(item interface{}, convert bool) {
	if c.items == nil {
		c.items = make([]interface{}, 0)
	}
	c.items = append(c.items, c.prepareItem(item, convert))
}

func (c *collectionType) prepareItem(item interface{}, convert bool) interface{} {
	if item == nil {
		return nil
	}

	if convert {
		if _, ok := item.(AnyAccessor); !ok {
			item = c.adapter.ConvertToSystem(item)
		}
	}

	typeInfo := ModelTypeInfo(c.adapter, item)
	if c.itemTypeInfo == nil {
		c.itemTypeInfo = typeInfo
	} else {
		typeInfo = CommonBaseType(c.itemTypeInfo, typeInfo)
		if typeInfo != nil {
			c.itemTypeInfo = typeInfo
		} else {
			c.itemTypeInfo = UndefinedTypeInfo
		}
	}

	return item
}

func (c *collectionType) AddUnique(item interface{}) bool {
	if c.items == nil {
		c.Add(item)
		return true
	}

	if item == nil {
		for _, o := range c.items {
			if o == nil {
				return false
			}
		}
	} else {
		if sysNode, ok := item.(AnyAccessor); ok {
			for _, o := range c.items {
				if o != nil && SystemAnyEqual(sysNode, o) {
					return false
				}
			}
		} else {
			item = c.adapter.ConvertToSystem(item)
			for _, o := range c.items {
				if o != nil && ModelEqual(c.adapter, item, o) {
					return false
				}
			}
		}
	}

	c.add(item, false)
	return true
}

func (c *collectionType) AddAll(collection CollectionAccessor) int {
	count := collection.Count()
	for i := 0; i < count; i++ {
		c.add(collection.Get(i), false)
	}
	return count
}

func (c *collectionType) AddAllUnique(collection CollectionAccessor) int {
	added := 0
	count := collection.Count()
	for i := 0; i < count; i++ {
		if c.AddUnique(collection.Get(i)) {
			added = added + 1
		}
	}
	return added
}

func (c *collectionType) Equal(item interface{}) bool {
	if o, ok := item.(CollectionAccessor); !ok {
		return false
	} else {
		return c.Count() == o.Count() &&
			collectionDeepEqual(c.adapter, c, o)
	}
}

func (c *collectionType) Equivalent(item interface{}) bool {
	if o, ok := item.(CollectionAccessor); !ok {
		return false
	} else {
		return c.Count() == o.Count() &&
			collectionDeepEquivalent(c.adapter, c, o)
	}
}

func collectionDeepEqual(adapter ModelAdapter, c1 CollectionAccessor, c2 CollectionAccessor) bool {
	count := c1.Count()
	for i := 0; i < count; i++ {
		if !ModelEqual(adapter, c1.Get(i), c2.Get(i)) {
			return false
		}
	}
	return true
}

func collectionDeepEquivalent(adapter ModelAdapter, c1 CollectionAccessor, c2 CollectionAccessor) bool {
	count := c1.Count()
	for i := 0; i < count; i++ {
		if !ModelEquivalent(adapter, c1.Get(i), c2.Get(i)) {
			return false
		}
	}
	return true
}

type emptyCollectionType struct {
	baseAnyType
}

func NewEmptyCollection() CollectionAccessor {
	return NewEmptyCollectionWithSource(nil)
}

func NewEmptyCollectionWithSource(source interface{}) CollectionAccessor {
	return &emptyCollectionType{
		baseAnyType: baseAnyType{
			source: source,
		},
	}
}

func (c *emptyCollectionType) DataType() DataTypes {
	return CollectionDataType
}

func (c *emptyCollectionType) TypeInfo() TypeInfoAccessor {
	return collectionTypeInfo
}

func (c *emptyCollectionType) Equal(item interface{}) bool {
	if o, ok := item.(CollectionAccessor); !ok {
		return false
	} else {
		return o.Empty()
	}
}

func (c *emptyCollectionType) Equivalent(item interface{}) bool {
	return c.Equal(item)
}

func (c *emptyCollectionType) ItemTypeInfo() TypeInfoAccessor {
	return UndefinedTypeInfo
}

func (c *emptyCollectionType) Empty() bool {
	return true
}

func (c *emptyCollectionType) Count() int {
	return 0
}

func (c *emptyCollectionType) Get(int) interface{} {
	panic("cannot get an item from an empty collection")
}

func (c *emptyCollectionType) Contains(interface{}) bool {
	return false
}
