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

import "github.com/volsch/gohimodel/datatype"

func convertContextData(accessor datatype.Accessor) datatype.Accessor {
	if accessor == nil {
		return accessor
	}

	dt := accessor.DataType()
	if dt == datatype.QuantityDataType {
		return convertContextQuantity(accessor.(datatype.QuantityAccessor))
	}
	return accessor
}

func convertContextQuantity(quantity datatype.QuantityAccessor) datatype.QuantityAccessor {
	code := quantity.Code()
	system := quantity.System()
	if code == nil || code.Nil() ||
		!datatype.ValueEqual(system, datatype.UCUMSystemURI) {
		return quantity
	}

	origCodeValue := code.String()
	switch origCodeValue {
	case "a":
		code = YearQuantityCode
		system = nil
	case "mo":
		code = MonthQuantityCode
		system = nil
	case "d":
		code = DayQuantityCode
		system = nil
	case "h":
		code = HourQuantityCode
		system = nil
	case "min":
		code = MinuteQuantityCode
		system = nil
	case "s":
		code = SecondQuantityCode
		system = nil
	}

	if origCodeValue == code.String() {
		return quantity
	}
	return datatype.NewQuantity(quantity.Value(), quantity.Comparator(), quantity.Unit(),
		system, code)
}

func uniteCollections(a1 datatype.Accessor, a2 datatype.Accessor) datatype.CollectionModifier {
	if a1 == nil && a2 == nil {
		return nil
	}

	c := newCollectionWithAccessorTypes([]datatype.Accessor{a1, a2})
	addUniqueCollectionItems(c, a1)
	addUniqueCollectionItems(c, a2)

	if c.Count() == 0 {
		return nil
	}
	return c
}

func addUniqueCollectionItems(collection datatype.CollectionModifier, accessor datatype.Accessor) {
	if accessor == nil {
		return
	}
	if c, ok := accessor.(datatype.CollectionAccessor); ok {
		collection.AddAllUnique(c)
	} else {
		collection.AddUnique(accessor)
	}
}

func combineCollections(a1 datatype.Accessor, a2 datatype.Accessor) datatype.CollectionModifier {
	if a1 == nil && a2 == nil {
		return nil
	}

	c := newCollectionWithAccessorTypes([]datatype.Accessor{a1, a2})
	addCollectionItems(c, a1)
	addCollectionItems(c, a2)

	if c.Count() == 0 {
		return nil
	}
	return c
}

func addCollectionItems(collection datatype.CollectionModifier, accessor datatype.Accessor) {
	if accessor == nil {
		return
	}
	if c, ok := accessor.(datatype.CollectionAccessor); ok {
		collection.AddAll(c)
	} else {
		collection.Add(accessor)
	}
}

func unwrapCollection(accessor datatype.Accessor) datatype.Accessor {
	if accessor == nil {
		return nil
	}
	if c, ok := accessor.(datatype.CollectionAccessor); !ok {
		return accessor
	} else {
		count := c.Count()
		if count == 0 {
			return nil
		}
		if count == 1 {
			return c.Get(0)
		}
		return c
	}
}

func newCollectionWithAccessorTypes(accessors []datatype.Accessor) datatype.CollectionModifier {
	typeInfo := commonAccessorBaseType(accessors)
	if typeInfo == nil {
		return datatype.NewCollectionUndefined()
	}
	return datatype.NewCollection(typeInfo)
}

func commonAccessorBaseType(accessors []datatype.Accessor) datatype.TypeInfoAccessor {
	var typeInfo datatype.TypeInfoAccessor
	for _, accessor := range accessors {
		if accessor != nil {
			if c, ok := accessor.(datatype.CollectionAccessor); ok {
				count := c.Count()
				for i := 0; i < count; i++ {
					a := c.Get(i)
					if a != nil {
						typeInfo = mergeCommonAccessorBaseType(a, typeInfo)
						if typeInfo == nil {
							return nil
						}
					}
				}
			} else {
				typeInfo = mergeCommonAccessorBaseType(accessor, typeInfo)
				if typeInfo == nil {
					return nil
				}
			}
		}
	}
	return typeInfo
}

func mergeCommonAccessorBaseType(accessor datatype.Accessor,
	typeInfo datatype.TypeInfoAccessor) datatype.TypeInfoAccessor {
	if typeInfo == nil {
		return accessor.TypeInfo()
	}
	return datatype.CommonBaseType(typeInfo, accessor.TypeInfo())
}
