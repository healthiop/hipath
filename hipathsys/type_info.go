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

var namespaceNameString = NewString(NamespaceName)

var simpleTypeInfoTypeSpec = NewTypeSpec(NewFQTypeName("SimpleTypeInfo", NamespaceName))
var classInfoTypeSpec = NewTypeSpec(NewFQTypeName("ClassInfo", NamespaceName))
var classInfoElementTypeSpec = NewTypeSpec(NewFQTypeName("ClassInfoElement", NamespaceName))
var listTypeInfoTypeSpec = NewTypeSpec(NewFQTypeName("ListTypeInfo", NamespaceName))
var tupleTypeInfoTypeSpec = NewTypeSpec(NewFQTypeName("TupleTypeInfo", NamespaceName))
var tupleTypeInfoElementTypeSpec = NewTypeSpec(NewFQTypeName("TupleTypeInfoElement", NamespaceName))

var simpleTypeInfoTypeInfo = NewClassInfo(namespaceNameString, NewString("SimpleTypeInfo"), nil,
	NewSysArrayCol(classInfoElementTypeSpec, []interface{}{
		NewClassInfoElement(NewString("namespace"), NewString("System.String"), nil),
		NewClassInfoElement(NewString("name"), NewString("System.String"), nil),
		NewClassInfoElement(NewString("baseType"), NewString("System.String"), nil),
	}))
var classInfoTypeInfo = NewClassInfo(namespaceNameString, NewString("ClassInfo"), nil,
	NewSysArrayCol(classInfoElementTypeSpec, []interface{}{
		NewClassInfoElement(NewString("namespace"), NewString("System.String"), nil),
		NewClassInfoElement(NewString("name"), NewString("System.String"), nil),
		NewClassInfoElement(NewString("baseType"), NewString("System.String"), nil),
		NewClassInfoElement(NewString("element"), NewString("List<System.ClassInfoElement>"), False),
	}))
var classInfoElementTypeInfo = NewClassInfo(namespaceNameString, NewString("ClassInfoElement"), nil,
	NewSysArrayCol(classInfoElementTypeSpec, []interface{}{
		NewClassInfoElement(NewString("name"), NewString("System.String"), nil),
		NewClassInfoElement(NewString("type"), NewString("System.String"), nil),
		NewClassInfoElement(NewString("isOneBased"), NewString("System.Boolean"), nil),
	}))
var listTypeInfoTypeInfo = NewClassInfo(namespaceNameString, NewString("ListTypeInfo"), nil,
	NewSysArrayCol(classInfoElementTypeSpec, []interface{}{
		NewClassInfoElement(NewString("elementType"), NewString("System.String"), nil),
	}))
var tupleTypeInfoTypeInfo = NewClassInfo(namespaceNameString, NewString("TupleTypeInfo"), nil,
	NewSysArrayCol(classInfoElementTypeSpec, []interface{}{
		NewClassInfoElement(NewString("element"), NewString("List<System.TupleTypeInfoElement>"), False),
	}))
var tupleTypeInfoElementTypeInfo = NewClassInfo(namespaceNameString, NewString("TupleTypeInfoElement"), nil,
	NewSysArrayCol(classInfoElementTypeSpec, []interface{}{
		NewClassInfoElement(NewString("name"), NewString("System.String"), nil),
		NewClassInfoElement(NewString("type"), NewString("System.String"), nil),
		NewClassInfoElement(NewString("isOneBased"), NewString("System.Boolean"), nil),
	}))

type TypeInfoAccessor interface {
	AnyAccessor
	Namespace() StringAccessor
}

type TypeInfoElementAccessor interface {
	AnyAccessor
	Name() StringAccessor
	Type() StringAccessor
	OneBased() BooleanAccessor
}

type SimpleTypeInfoAccessor interface {
	TypeInfoAccessor
	Name() StringAccessor
	BaseType() StringAccessor
}

type ClassInfoAccessor interface {
	TypeInfoAccessor
	Name() StringAccessor
	BaseType() StringAccessor
	Element() ColAccessor
}

type ClassInfoElementAccessor interface {
	TypeInfoElementAccessor
}

type ListTypeInfoAccessor interface {
	TypeInfoAccessor
	ElementType() StringAccessor
}

type TupleTypeInfoAccessor interface {
	TypeInfoAccessor
	Element() ColAccessor
}

type TupleTypeInfoElementAccessor interface {
	TypeInfoElementAccessor
}

type typeInfo struct {
	namespace StringAccessor
}
type typeInfoElement struct {
	name     StringAccessor
	typeName StringAccessor
	oneBased BooleanAccessor
}

type simpleTypeInfo struct {
	typeInfo
	name     StringAccessor
	baseType StringAccessor
}

type classInfo struct {
	typeInfo
	name     StringAccessor
	baseType StringAccessor
	element  ColAccessor
}

type classInfoElement struct {
	typeInfoElement
}

type listTypeInfo struct {
	typeInfo
	elementType StringAccessor
}

type tupleTypeInfo struct {
	typeInfo
	element ColAccessor
}

type tupleTypeInfoElement struct {
	typeInfoElement
}

func NewSimpleTypeInfo(namespace StringAccessor, name StringAccessor, baseType StringAccessor) SimpleTypeInfoAccessor {
	return &simpleTypeInfo{
		typeInfo{
			namespace,
		},
		name,
		baseType,
	}
}

func NewClassInfo(namespace StringAccessor, name StringAccessor, baseType StringAccessor, element ColAccessor) ClassInfoAccessor {
	return &classInfo{
		typeInfo{
			namespace,
		},
		name,
		baseType,
		element,
	}
}

func NewClassInfoElement(name StringAccessor, typeName StringAccessor, oneBased BooleanAccessor) ClassInfoElementAccessor {
	return &classInfoElement{
		typeInfoElement{
			name,
			typeName,
			oneBased,
		},
	}
}

func NewListTypeInfo(elementType StringAccessor) ListTypeInfoAccessor {
	return &listTypeInfo{
		typeInfo{
			namespaceNameString,
		},
		elementType,
	}
}

func NewTupleTypeInfo(namespace StringAccessor, element ColAccessor) TupleTypeInfoAccessor {
	return &tupleTypeInfo{
		typeInfo{
			namespace,
		},
		element,
	}
}

func NewTupleTypeInfoElement(name StringAccessor, typeName StringAccessor, oneBased BooleanAccessor) TupleTypeInfoElementAccessor {
	return &tupleTypeInfoElement{
		typeInfoElement{
			name,
			typeName,
			oneBased,
		},
	}
}

func (t *typeInfo) DataType() DataTypes {
	return UndefinedDataType
}

func (t *typeInfo) Source() interface{} {
	return nil
}

func (t *typeInfo) Namespace() StringAccessor {
	return t.namespace
}

func (t *simpleTypeInfo) TypeSpec() TypeSpecAccessor {
	return simpleTypeInfoTypeSpec
}

func (t *simpleTypeInfo) TypeInfo() TypeInfoAccessor {
	return simpleTypeInfoTypeInfo
}

func (t *simpleTypeInfo) Equal(node interface{}) bool {
	if ti, ok := node.(SimpleTypeInfoAccessor); !ok {
		return false
	} else {
		return Equal(t.Namespace(), ti.Namespace()) &&
			Equal(t.Name(), ti.Name())
	}
}

func (t *simpleTypeInfo) Equivalent(node interface{}) bool {
	return t.Equal(node)
}

func (t *simpleTypeInfo) Name() StringAccessor {
	return t.name
}

func (t *simpleTypeInfo) BaseType() StringAccessor {
	return t.baseType
}

func (t *classInfo) TypeSpec() TypeSpecAccessor {
	return classInfoTypeSpec
}

func (t *classInfo) TypeInfo() TypeInfoAccessor {
	return classInfoTypeInfo
}

func (t *classInfo) Equal(node interface{}) bool {
	if ti, ok := node.(ClassInfoAccessor); !ok {
		return false
	} else {
		return Equal(t.Namespace(), ti.Namespace()) &&
			Equal(t.Name(), ti.Name())
	}
}

func (t *classInfo) Equivalent(node interface{}) bool {
	return t.Equal(node)
}

func (t *classInfo) Name() StringAccessor {
	return t.name
}

func (t *classInfo) BaseType() StringAccessor {
	return t.baseType
}

func (t *classInfo) Element() ColAccessor {
	return t.element
}

func (t *typeInfoElement) DataType() DataTypes {
	return UndefinedDataType
}

func (t *typeInfoElement) Source() interface{} {
	return nil
}

func (t *typeInfoElement) Equal(node interface{}) bool {
	if ti, ok := node.(TypeInfoElementAccessor); !ok {
		return false
	} else {
		return Equal(t.Name(), ti.Name()) &&
			Equal(t.Type(), ti.Type()) &&
			Equal(t.OneBased(), ti.OneBased())
	}
}

func (t *typeInfoElement) Equivalent(node interface{}) bool {
	return t.Equal(node)
}

func (t *typeInfoElement) Name() StringAccessor {
	return t.name
}

func (t *typeInfoElement) Type() StringAccessor {
	return t.typeName
}

func (t *typeInfoElement) OneBased() BooleanAccessor {
	return t.oneBased
}

func (t *classInfoElement) TypeSpec() TypeSpecAccessor {
	return classInfoElementTypeSpec
}

func (t *classInfoElement) TypeInfo() TypeInfoAccessor {
	return classInfoElementTypeInfo
}

func (t *listTypeInfo) TypeSpec() TypeSpecAccessor {
	return listTypeInfoTypeSpec
}

func (t *listTypeInfo) TypeInfo() TypeInfoAccessor {
	return listTypeInfoTypeInfo
}

func (t *listTypeInfo) Equal(node interface{}) bool {
	if ti, ok := node.(ListTypeInfoAccessor); !ok {
		return false
	} else {
		return t.elementType != nil && Equal(t.ElementType(), ti.ElementType())
	}
}

func (t *listTypeInfo) Equivalent(node interface{}) bool {
	return t.Equal(node)
}

func (t *listTypeInfo) ElementType() StringAccessor {
	return t.elementType
}

func (t *tupleTypeInfo) TypeSpec() TypeSpecAccessor {
	return tupleTypeInfoTypeSpec
}

func (t *tupleTypeInfo) TypeInfo() TypeInfoAccessor {
	return tupleTypeInfoTypeInfo
}

func (t *tupleTypeInfo) Equal(_ interface{}) bool {
	return false
}

func (t *tupleTypeInfo) Equivalent(node interface{}) bool {
	return t.Equal(node)
}

func (t *tupleTypeInfo) Element() ColAccessor {
	return t.element
}

func (t *tupleTypeInfoElement) TypeSpec() TypeSpecAccessor {
	return tupleTypeInfoElementTypeSpec
}

func (t *tupleTypeInfoElement) TypeInfo() TypeInfoAccessor {
	return tupleTypeInfoElementTypeInfo
}
