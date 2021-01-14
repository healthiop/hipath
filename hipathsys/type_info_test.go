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

func TestNewSimpleTypeInfo(t *testing.T) {
	ti := NewSimpleTypeInfo(NewString("test_ns"), NewString("test_name"), NewString("test.Base"))
	assert.Equal(t, NewString("test_ns"), ti.Namespace())
	assert.Equal(t, NewString("test_name"), ti.Name())
	assert.Equal(t, NewString("test.Base"), ti.BaseType())
}

func TestSimpleTypeInfoDataType(t *testing.T) {
	ti := NewSimpleTypeInfo(NewString("test_ns"), NewString("test_name"), NewString("test.Base"))
	assert.Equal(t, UndefinedDataType, ti.DataType())
}

func TestSimpleTypeInfoSource(t *testing.T) {
	ti := NewSimpleTypeInfo(NewString("test_ns"), NewString("test_name"), NewString("test.Base"))
	assert.Nil(t, ti.Source())
}

func TestSimpleTypeInfoTypeSpec(t *testing.T) {
	ti := NewSimpleTypeInfo(NewString("test_ns"), NewString("test_name"), NewString("test.Base"))
	s := ti.TypeSpec()
	s.Anonymous()
	if assert.NotNil(t, s, "type spec expected") {
		assert.Equal(t, "System.SimpleTypeInfo", s.String())
		assert.Nil(t, s.FQBaseName(), "no base name expected")
	}
}

func TestSimpleTypeInfoTypeInfo(t *testing.T) {
	ti := NewSimpleTypeInfo(NewString("test_ns"), NewString("test_name"), NewString("test.Base"))
	i := ti.TypeInfo()
	if assert.Implements(t, (*ClassInfoAccessor)(nil), i) {
		a := i.(ClassInfoAccessor)
		assert.Equal(t, NewString("System"), a.Namespace())
		assert.Equal(t, NewString("SimpleTypeInfo"), a.Name())
		assert.Nil(t, a.BaseType())

		if assert.Equal(t, 3, a.Element().Count()) {
			e := a.Element().Get(0).(ClassInfoElementAccessor)
			assert.Equal(t, NewString("namespace"), e.Name())
			assert.Equal(t, NewString("System.String"), e.Type())
			assert.Nil(t, e.OneBased())
			e = a.Element().Get(1).(ClassInfoElementAccessor)
			assert.Equal(t, NewString("name"), e.Name())
			assert.Equal(t, NewString("System.String"), e.Type())
			assert.Nil(t, e.OneBased())
			e = a.Element().Get(2).(ClassInfoElementAccessor)
			assert.Equal(t, NewString("baseType"), e.Name())
			assert.Equal(t, NewString("System.String"), e.Type())
			assert.Nil(t, e.OneBased())
		}
	}
}

func TestSimpleTypeInfoEqual(t *testing.T) {
	t1 := NewSimpleTypeInfo(NewString("test_ns"), NewString("test_name"), NewString("test.Base"))
	t2 := NewSimpleTypeInfo(NewString("test_ns"), NewString("test_name"), NewString("test.Base"))
	assert.True(t, t1.Equal(t2))
}

func TestSimpleTypeInfoEqualNotNamespace(t *testing.T) {
	t1 := NewSimpleTypeInfo(NewString("test_ns"), NewString("test_name"), NewString("test.Base"))
	t2 := NewSimpleTypeInfo(NewString("test_ns2"), NewString("test_name"), NewString("test.Base"))
	assert.False(t, t1.Equal(t2))
}

func TestSimpleTypeInfoEqualNotName(t *testing.T) {
	t1 := NewSimpleTypeInfo(NewString("test_ns"), NewString("test_name"), NewString("test.Base"))
	t2 := NewSimpleTypeInfo(NewString("test_ns"), NewString("test_name2"), NewString("test.Base"))
	assert.False(t, t1.Equal(t2))
}

func TestSimpleTypeInfoEqualBaseDiffers(t *testing.T) {
	t1 := NewSimpleTypeInfo(NewString("test_ns"), NewString("test_name"), NewString("test.Base"))
	t2 := NewSimpleTypeInfo(NewString("test_ns"), NewString("test_name"), NewString("test.Base2"))
	assert.True(t, t1.Equal(t2))
}

func TestSimpleTypeInfoEqualNotType(t *testing.T) {
	t1 := NewSimpleTypeInfo(NewString("test_ns"), NewString("test_name"), NewString("test.Base"))
	assert.False(t, t1.Equal(NewString("test")))
}

func TestSimpleTypeInfoEquivalent(t *testing.T) {
	t1 := NewSimpleTypeInfo(NewString("test_ns"), NewString("test_name"), NewString("test.Base"))
	t2 := NewSimpleTypeInfo(NewString("test_ns"), NewString("test_name"), NewString("test.Base"))
	assert.True(t, t1.Equivalent(t2))
}

func TestSimpleTypeInfoEquivalentNot(t *testing.T) {
	t1 := NewSimpleTypeInfo(NewString("test_ns"), NewString("test_name"), NewString("test.Base"))
	t2 := NewSimpleTypeInfo(NewString("test_ns2"), NewString("test_name"), NewString("test.Base"))
	assert.False(t, t1.Equivalent(t2))
}

func TestNewClassInfo(t *testing.T) {
	element := NewSysArrayCol(classInfoElementTypeSpec, []interface{}{
		NewClassInfoElement(NewString("val1"), NewString("System.String"), nil),
		NewClassInfoElement(NewString("val2"), NewString("System.String"), nil),
	})
	ti := NewClassInfo(NewString("test_ns"), NewString("test_name"), NewString("test.Base"), element)
	assert.Equal(t, NewString("test_ns"), ti.Namespace())
	assert.Equal(t, NewString("test_name"), ti.Name())
	assert.Equal(t, NewString("test.Base"), ti.BaseType())
	assert.Same(t, element, ti.Element())
}

func TestClassInfoDataType(t *testing.T) {
	ti := NewClassInfo(NewString("test_ns"), NewString("test_name"), NewString("test.Base"), EmptyCol)
	assert.Equal(t, UndefinedDataType, ti.DataType())
}

func TestClassInfoSource(t *testing.T) {
	ti := NewClassInfo(NewString("test_ns"), NewString("test_name"), NewString("test.Base"), EmptyCol)
	assert.Nil(t, ti.Source())
}

func TestClassInfoTypeSpec(t *testing.T) {
	ti := NewClassInfo(NewString("test_ns"), NewString("test_name"), NewString("test.Base"), EmptyCol)
	s := ti.TypeSpec()
	s.Anonymous()
	if assert.NotNil(t, s, "type spec expected") {
		assert.Equal(t, "System.ClassInfo", s.String())
		assert.Nil(t, s.FQBaseName(), "no base name expected")
	}
}

func TestClassInfoTypeInfo(t *testing.T) {
	ti := NewClassInfo(NewString("test_ns"), NewString("test_name"), NewString("test.Base"), EmptyCol)
	i := ti.TypeInfo()
	if assert.Implements(t, (*ClassInfoAccessor)(nil), i) {
		a := i.(ClassInfoAccessor)
		assert.Equal(t, NewString("System"), a.Namespace())
		assert.Equal(t, NewString("ClassInfo"), a.Name())
		assert.Nil(t, a.BaseType())

		if assert.Equal(t, 4, a.Element().Count()) {
			e := a.Element().Get(0).(ClassInfoElementAccessor)
			assert.Equal(t, NewString("namespace"), e.Name())
			assert.Equal(t, NewString("System.String"), e.Type())
			assert.Nil(t, e.OneBased())
			e = a.Element().Get(1).(ClassInfoElementAccessor)
			assert.Equal(t, NewString("name"), e.Name())
			assert.Equal(t, NewString("System.String"), e.Type())
			assert.Nil(t, e.OneBased())
			e = a.Element().Get(2).(ClassInfoElementAccessor)
			assert.Equal(t, NewString("baseType"), e.Name())
			assert.Equal(t, NewString("System.String"), e.Type())
			assert.Nil(t, e.OneBased())
			e = a.Element().Get(3).(ClassInfoElementAccessor)
			assert.Equal(t, NewString("element"), e.Name())
			assert.Equal(t, NewString("List<System.ClassInfoElement>"), e.Type())
			assert.Equal(t, False, e.OneBased())
		}
	}
}

func TestClassInfoEqual(t *testing.T) {
	t1 := NewClassInfo(NewString("test_ns"), NewString("test_name"), NewString("test.Base"), EmptyCol)
	t2 := NewClassInfo(NewString("test_ns"), NewString("test_name"), NewString("test.Base"), EmptyCol)
	assert.True(t, t1.Equal(t2))
}

func TestClassInfoEqualNotNamespace(t *testing.T) {
	t1 := NewClassInfo(NewString("test_ns"), NewString("test_name"), NewString("test.Base"), EmptyCol)
	t2 := NewClassInfo(NewString("test_ns2"), NewString("test_name"), NewString("test.Base"), EmptyCol)
	assert.False(t, t1.Equal(t2))
}

func TestClassInfoEqualNotName(t *testing.T) {
	t1 := NewClassInfo(NewString("test_ns"), NewString("test_name"), NewString("test.Base"), EmptyCol)
	t2 := NewClassInfo(NewString("test_ns"), NewString("test_name2"), NewString("test.Base"), EmptyCol)
	assert.False(t, t1.Equal(t2))
}

func TestClassInfoEqualBaseDiffers(t *testing.T) {
	t1 := NewClassInfo(NewString("test_ns"), NewString("test_name"), NewString("test.Base"), EmptyCol)
	t2 := NewClassInfo(NewString("test_ns"), NewString("test_name"), NewString("test.Base2"), EmptyCol)
	assert.True(t, t1.Equal(t2))
}

func TestClassInfoEqualNotType(t *testing.T) {
	t1 := NewClassInfo(NewString("test_ns"), NewString("test_name"), NewString("test.Base"), EmptyCol)
	assert.False(t, t1.Equal(NewString("test")))
}

func TestClassInfoEquivalent(t *testing.T) {
	t1 := NewClassInfo(NewString("test_ns"), NewString("test_name"), NewString("test.Base"), EmptyCol)
	t2 := NewClassInfo(NewString("test_ns"), NewString("test_name"), NewString("test.Base"), EmptyCol)
	assert.True(t, t1.Equivalent(t2))
}

func TestClassInfoEquivalentNot(t *testing.T) {
	t1 := NewClassInfo(NewString("test_ns"), NewString("test_name"), NewString("test.Base"), EmptyCol)
	t2 := NewClassInfo(NewString("test_ns2"), NewString("test_name"), NewString("test.Base"), EmptyCol)
	assert.False(t, t1.Equivalent(t2))
}

func TestNewClassInfoElement(t *testing.T) {
	ti := NewClassInfoElement(NewString("valX"), NewString("test.Decimal"), True)
	assert.Equal(t, NewString("valX"), ti.Name())
	assert.Equal(t, NewString("test.Decimal"), ti.Type())
	assert.Equal(t, True, ti.OneBased())
}

func TestClassInfoElementDataType(t *testing.T) {
	ti := NewClassInfoElement(NewString("valX"), NewString("test.Decimal"), True)
	assert.Equal(t, UndefinedDataType, ti.DataType())
}

func TestClassInfoElementSource(t *testing.T) {
	ti := NewClassInfoElement(NewString("valX"), NewString("test.Decimal"), True)
	assert.Nil(t, ti.Source())
}

func TestClassInfoElementTypeSpec(t *testing.T) {
	ti := NewClassInfoElement(NewString("valX"), NewString("test.Decimal"), True)
	s := ti.TypeSpec()
	s.Anonymous()
	if assert.NotNil(t, s, "type spec expected") {
		assert.Equal(t, "System.ClassInfoElement", s.String())
		assert.Nil(t, s.FQBaseName(), "no base name expected")
	}
}

func TestClassInfoElementTypeInfo(t *testing.T) {
	ti := NewClassInfoElement(NewString("valX"), NewString("test.Decimal"), True)
	i := ti.TypeInfo()
	if assert.Implements(t, (*ClassInfoAccessor)(nil), i) {
		a := i.(ClassInfoAccessor)
		assert.Equal(t, NewString("System"), a.Namespace())
		assert.Equal(t, NewString("ClassInfoElement"), a.Name())
		assert.Nil(t, a.BaseType())

		if assert.Equal(t, 3, a.Element().Count()) {
			e := a.Element().Get(0).(ClassInfoElementAccessor)
			assert.Equal(t, NewString("name"), e.Name())
			assert.Equal(t, NewString("System.String"), e.Type())
			assert.Nil(t, e.OneBased())
			e = a.Element().Get(1).(ClassInfoElementAccessor)
			assert.Equal(t, NewString("type"), e.Name())
			assert.Equal(t, NewString("System.String"), e.Type())
			assert.Nil(t, e.OneBased())
			e = a.Element().Get(2).(ClassInfoElementAccessor)
			assert.Equal(t, NewString("isOneBased"), e.Name())
			assert.Equal(t, NewString("System.Boolean"), e.Type())
			assert.Nil(t, e.OneBased())
		}
	}
}

func TestClassInfoElementOther(t *testing.T) {
	t1 := NewClassInfoElement(NewString("valX"), NewString("test.Decimal"), True)
	assert.False(t, t1.Equal(NewString("test")))
}

func TestClassInfoElementEqual(t *testing.T) {
	t1 := NewClassInfoElement(NewString("valX"), NewString("test.Decimal"), True)
	t2 := NewClassInfoElement(NewString("valX"), NewString("test.Decimal"), True)
	assert.True(t, t1.Equal(t2))
}

func TestClassInfoElementEqualNotName(t *testing.T) {
	t1 := NewClassInfoElement(NewString("valX"), NewString("test.Decimal"), True)
	t2 := NewClassInfoElement(NewString("valY"), NewString("test.Decimal"), True)
	assert.False(t, t1.Equal(t2))
}

func TestClassInfoElementEqualNotType(t *testing.T) {
	t1 := NewClassInfoElement(NewString("valX"), NewString("test.Decimal"), True)
	t2 := NewClassInfoElement(NewString("valX"), NewString("test.decimal"), True)
	assert.False(t, t1.Equal(t2))
}

func TestClassInfoElementEqualNotBase(t *testing.T) {
	t1 := NewClassInfoElement(NewString("valX"), NewString("test.Decimal"), True)
	t2 := NewClassInfoElement(NewString("valX"), NewString("test.Decimal"), False)

	assert.False(t, t1.Equal(t2))
}

func TestClassInfoElementEquivalent(t *testing.T) {
	t1 := NewClassInfoElement(NewString("valX"), NewString("test.Decimal"), True)
	t2 := NewClassInfoElement(NewString("valX"), NewString("test.Decimal"), True)
	assert.True(t, t1.Equivalent(t2))
}

func TestClassInfoElementEquivalentNot(t *testing.T) {
	t1 := NewClassInfoElement(NewString("valX"), NewString("test.Decimal"), True)
	t2 := NewClassInfoElement(NewString("valY"), NewString("test.Decimal"), True)
	assert.False(t, t1.Equivalent(t2))
}

func TestNewListTypeInfo(t *testing.T) {
	ti := NewListTypeInfo(NewString("test.Decimal"))
	assert.Equal(t, NewString("test.Decimal"), ti.ElementType())
}

func TestListTypeInfoDataType(t *testing.T) {
	ti := NewListTypeInfo(NewString("test.Decimal"))
	assert.Equal(t, UndefinedDataType, ti.DataType())
}

func TestListTypeInfoSource(t *testing.T) {
	ti := NewListTypeInfo(NewString("test.Decimal"))
	assert.Nil(t, ti.Source())
}

func TestListTypeInfoTypeSpec(t *testing.T) {
	ti := NewListTypeInfo(NewString("test.Decimal"))
	s := ti.TypeSpec()
	s.Anonymous()
	if assert.NotNil(t, s, "type spec expected") {
		assert.Equal(t, "System.ListTypeInfo", s.String())
		assert.Nil(t, s.FQBaseName(), "no base name expected")
	}
}

func TestListTypeInfoTypeInfo(t *testing.T) {
	ti := NewListTypeInfo(NewString("test.Decimal"))
	i := ti.TypeInfo()
	if assert.Implements(t, (*ClassInfoAccessor)(nil), i) {
		a := i.(ClassInfoAccessor)
		assert.Equal(t, NewString("System"), a.Namespace())
		assert.Equal(t, NewString("ListTypeInfo"), a.Name())
		assert.Nil(t, a.BaseType())

		if assert.Equal(t, 1, a.Element().Count()) {
			e := a.Element().Get(0).(ClassInfoElementAccessor)
			assert.Equal(t, NewString("elementType"), e.Name())
			assert.Equal(t, NewString("System.String"), e.Type())
			assert.Nil(t, e.OneBased())
		}
	}
}

func TestListTypeInfoEqualOther(t *testing.T) {
	t1 := NewListTypeInfo(NewString("test.Decimal"))
	assert.False(t, t1.Equal(NewString("test")))
}

func TestListTypeInfoEqual(t *testing.T) {
	t1 := NewListTypeInfo(NewString("test.Decimal"))
	t2 := NewListTypeInfo(NewString("test.Decimal"))
	assert.True(t, t1.Equal(t2))
}

func TestListTypeInfoEqualNotElementType(t *testing.T) {
	t1 := NewListTypeInfo(NewString("test.Decimal"))
	t2 := NewListTypeInfo(NewString("test.decimal"))
	assert.False(t, t1.Equal(t2))
}

func TestListTypeInfoEquivalent(t *testing.T) {
	t1 := NewListTypeInfo(NewString("test.Decimal"))
	t2 := NewListTypeInfo(NewString("test.Decimal"))
	assert.True(t, t1.Equivalent(t2))
}

func TestListTypeInfoEquivalentNot(t *testing.T) {
	t1 := NewListTypeInfo(NewString("test.Decimal"))
	t2 := NewListTypeInfo(NewString("test.decimal"))
	assert.False(t, t1.Equivalent(t2))
}

func TestNewTupleTypeInfo(t *testing.T) {
	element := NewSysArrayCol(tupleTypeInfoElementTypeSpec, []interface{}{
		NewTupleTypeInfoElement(NewString("val1"), NewString("System.String"), nil),
		NewTupleTypeInfoElement(NewString("val2"), NewString("System.String"), nil),
	})
	ti := NewTupleTypeInfo(NewString("test_ns"), element)
	assert.Equal(t, NewString("test_ns"), ti.Namespace())
	assert.Same(t, element, ti.Element())
}

func TestTupleTypeInfoDataType(t *testing.T) {
	ti := NewTupleTypeInfo(NewString("test_ns"), EmptyCol)
	assert.Equal(t, UndefinedDataType, ti.DataType())
}

func TestTupleTypeInfoSource(t *testing.T) {
	ti := NewTupleTypeInfo(NewString("test_ns"), EmptyCol)
	assert.Nil(t, ti.Source())
}

func TestTupleTypeInfoTypeSpec(t *testing.T) {
	ti := NewTupleTypeInfo(NewString("test_ns"), EmptyCol)
	s := ti.TypeSpec()
	s.Anonymous()
	if assert.NotNil(t, s, "type spec expected") {
		assert.Equal(t, "System.TupleTypeInfo", s.String())
		assert.Nil(t, s.FQBaseName(), "no base name expected")
	}
}

func TestTupleTypeInfoTypeInfo(t *testing.T) {
	ti := NewTupleTypeInfo(NewString("test_ns"), EmptyCol)
	i := ti.TypeInfo()
	if assert.Implements(t, (*ClassInfoAccessor)(nil), i) {
		a := i.(ClassInfoAccessor)
		assert.Equal(t, NewString("System"), a.Namespace())
		assert.Equal(t, NewString("TupleTypeInfo"), a.Name())
		assert.Nil(t, a.BaseType())

		if assert.Equal(t, 1, a.Element().Count()) {
			e := a.Element().Get(0).(TupleTypeInfoElementAccessor)
			assert.Equal(t, NewString("element"), e.Name())
			assert.Equal(t, NewString("List<System.TupleTypeInfoElement>"), e.Type())
			assert.Equal(t, False, e.OneBased())
		}
	}
}

func TestTupleTypeInfoEqual(t *testing.T) {
	t1 := NewTupleTypeInfo(NewString("test_ns"), EmptyCol)
	t2 := NewTupleTypeInfo(NewString("test_ns"), EmptyCol)
	assert.False(t, t1.Equal(t2))
}

func TestTupleTypeInfoEquivalent(t *testing.T) {
	t1 := NewTupleTypeInfo(NewString("test_ns"), EmptyCol)
	t2 := NewTupleTypeInfo(NewString("test_ns"), EmptyCol)
	assert.False(t, t1.Equivalent(t2))
}

func TestNewTupleTypeInfoElement(t *testing.T) {
	ti := NewTupleTypeInfoElement(NewString("valX"), NewString("test.Decimal"), True)
	assert.Equal(t, NewString("valX"), ti.Name())
	assert.Equal(t, NewString("test.Decimal"), ti.Type())
	assert.Equal(t, True, ti.OneBased())
}

func TestTupleTypeInfoElementDataType(t *testing.T) {
	ti := NewTupleTypeInfoElement(NewString("valX"), NewString("test.Decimal"), True)
	assert.Equal(t, UndefinedDataType, ti.DataType())
}

func TestTupleTypeInfoElementSource(t *testing.T) {
	ti := NewTupleTypeInfoElement(NewString("valX"), NewString("test.Decimal"), True)
	assert.Nil(t, ti.Source())
}

func TestTupleTypeInfoElementTypeSpec(t *testing.T) {
	ti := NewTupleTypeInfoElement(NewString("valX"), NewString("test.Decimal"), True)
	s := ti.TypeSpec()
	s.Anonymous()
	if assert.NotNil(t, s, "type spec expected") {
		assert.Equal(t, "System.TupleTypeInfoElement", s.String())
		assert.Nil(t, s.FQBaseName(), "no base name expected")
	}
}

func TestTupleTypeInfoElementTypeInfo(t *testing.T) {
	ti := NewTupleTypeInfoElement(NewString("valX"), NewString("test.Decimal"), True)
	i := ti.TypeInfo()
	if assert.Implements(t, (*ClassInfoAccessor)(nil), i) {
		a := i.(ClassInfoAccessor)
		assert.Equal(t, NewString("System"), a.Namespace())
		assert.Equal(t, NewString("TupleTypeInfoElement"), a.Name())
		assert.Nil(t, a.BaseType())

		if assert.Equal(t, 3, a.Element().Count()) {
			e := a.Element().Get(0).(ClassInfoElementAccessor)
			assert.Equal(t, NewString("name"), e.Name())
			assert.Equal(t, NewString("System.String"), e.Type())
			assert.Nil(t, e.OneBased())
			e = a.Element().Get(1).(ClassInfoElementAccessor)
			assert.Equal(t, NewString("type"), e.Name())
			assert.Equal(t, NewString("System.String"), e.Type())
			assert.Nil(t, e.OneBased())
			e = a.Element().Get(2).(ClassInfoElementAccessor)
			assert.Equal(t, NewString("isOneBased"), e.Name())
			assert.Equal(t, NewString("System.Boolean"), e.Type())
			assert.Nil(t, e.OneBased())
		}
	}
}

func TestTupleTypeInfoElementEqual(t *testing.T) {
	t1 := NewTupleTypeInfoElement(NewString("valX"), NewString("test.Decimal"), True)
	t2 := NewTupleTypeInfoElement(NewString("valX"), NewString("test.Decimal"), True)
	assert.True(t, t1.Equal(t2))
}

func TestTupleTypeInfoElementEqualNotName(t *testing.T) {
	t1 := NewTupleTypeInfoElement(NewString("valX"), NewString("test.Decimal"), True)
	t2 := NewTupleTypeInfoElement(NewString("valY"), NewString("test.Decimal"), True)
	assert.False(t, t1.Equal(t2))
}

func TestTupleTypeInfoElementEqualNotType(t *testing.T) {
	t1 := NewTupleTypeInfoElement(NewString("valX"), NewString("test.Decimal"), True)
	t2 := NewTupleTypeInfoElement(NewString("valX"), NewString("test.decimal"), True)
	assert.False(t, t1.Equal(t2))
}

func TestTupleTypeInfoElementEqualNotBase(t *testing.T) {
	t1 := NewTupleTypeInfoElement(NewString("valX"), NewString("test.Decimal"), True)
	t2 := NewTupleTypeInfoElement(NewString("valX"), NewString("test.Decimal"), False)

	assert.False(t, t1.Equal(t2))
}

func TestTupleTypeInfoElementEqualOther(t *testing.T) {
	t1 := NewTupleTypeInfoElement(NewString("valX"), NewString("test.Decimal"), True)
	assert.False(t, t1.Equal(NewString("test")))
}

func TestTupleTypeInfoElementEquivalent(t *testing.T) {
	t1 := NewTupleTypeInfoElement(NewString("valX"), NewString("test.Decimal"), True)
	t2 := NewTupleTypeInfoElement(NewString("valX"), NewString("test.Decimal"), True)
	assert.True(t, t1.Equivalent(t2))
}

func TestTupleTypeInfoElementEquivalentNot(t *testing.T) {
	t1 := NewTupleTypeInfoElement(NewString("valX"), NewString("test.Decimal"), True)
	t2 := NewTupleTypeInfoElement(NewString("valY"), NewString("test.Decimal"), True)
	assert.False(t, t1.Equivalent(t2))
}
