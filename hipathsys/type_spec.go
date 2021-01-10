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
	"fmt"
	"strings"
)

const NamespaceName = "System"

var UndefinedTypeSpec = NewTypeSpecWithBase(nil, nil)
var anyTypeSpec = NewTypeSpec(NewFQTypeName("Any", NamespaceName))

type fqTypeName struct {
	namespace string
	name      string
	fqName    string
}

type FQTypeNameAccessor interface {
	Anonymous() bool
	HasNamespace() bool
	Namespace() string
	Name() string
	String() string
	Equal(name FQTypeNameAccessor) bool
}

type typeSpec struct {
	base   TypeSpecAccessor
	fqName FQTypeNameAccessor
}

type TypeSpecAccessor interface {
	Anonymous() bool
	Base() TypeSpecAccessor
	FQName() FQTypeNameAccessor
	FQBaseName() FQTypeNameAccessor
	String() string
	EqualType(other TypeSpecAccessor) bool
	ExtendsName(name FQTypeNameAccessor) bool
}

func ParseFQTypeName(fqName string) (FQTypeNameAccessor, error) {
	if len(fqName) == 0 {
		return nil, fmt.Errorf("type name must not be empty")
	}

	fi := strings.IndexRune(fqName, '.')
	if fi < 0 {
		return NewTypeName(fqName), nil
	}
	if fi == 0 || fi+1 == len(fqName) {
		return nil, fmt.Errorf("invalid type name: %s", fqName)
	}

	name := fqName[fi+1:]
	li := strings.IndexRune(name, '.')
	if li >= 0 {
		return nil, fmt.Errorf("invalid type name: %s", fqName)
	}

	return NewFQTypeName(name, fqName[:fi]), nil
}

func NewFQTypeName(name string, namespace string) FQTypeNameAccessor {
	var fqName string
	if len(name) == 0 {
		fqName = ""
	} else if len(namespace) > 0 {
		fqName = namespace + "." + name
	} else {
		fqName = name
	}

	return &fqTypeName{
		namespace: namespace,
		name:      name,
		fqName:    fqName,
	}
}

func NewTypeName(name string) FQTypeNameAccessor {
	return &fqTypeName{
		namespace: "",
		name:      name,
		fqName:    name,
	}
}

func FQTypeNameEqual(t1 FQTypeNameAccessor, t2 FQTypeNameAccessor) bool {
	return (t1 == t2 && t1 != nil) || (t1 != nil && t2 != nil && t1.Equal(t2))
}

func NewTypeSpec(fqName FQTypeNameAccessor) TypeSpecAccessor {
	return NewTypeSpecWithBase(fqName, nil)
}

func NewTypeSpecWithBase(fqName FQTypeNameAccessor, base TypeSpecAccessor) TypeSpecAccessor {
	return &typeSpec{
		base:   base,
		fqName: fqName,
	}
}

func (t *fqTypeName) Anonymous() bool {
	return len(t.name) == 0
}

func (t *fqTypeName) HasNamespace() bool {
	return len(t.namespace) > 0
}

func (t *fqTypeName) Namespace() string {
	return t.namespace
}

func (t *fqTypeName) Name() string {
	return t.name
}

func (t *fqTypeName) String() string {
	return t.fqName
}

func (t *fqTypeName) Equal(name FQTypeNameAccessor) bool {
	return t.String() == name.String() && !t.Anonymous()
}

func (t *typeSpec) Anonymous() bool {
	return t.fqName == nil || t.fqName.Anonymous()
}

func (t *typeSpec) Base() TypeSpecAccessor {
	return t.base
}

func (t *typeSpec) FQName() FQTypeNameAccessor {
	return t.fqName
}

func (t *typeSpec) FQBaseName() FQTypeNameAccessor {
	if t.base == nil {
		return nil
	}
	return t.base.FQName()
}

func (t *typeSpec) String() string {
	if t.fqName == nil {
		return ""
	}
	return t.fqName.String()
}

func (t *typeSpec) EqualType(other TypeSpecAccessor) bool {
	return FQTypeNameEqual(t.FQName(), other.FQName())
}

func (t *typeSpec) ExtendsName(name FQTypeNameAccessor) bool {
	if t.fqName == nil {
		return false
	}

	if name.HasNamespace() {
		if name.Equal(t.fqName) {
			return true
		}
	} else {
		if name.Name() == t.fqName.Name() && !t.fqName.Anonymous() {
			return true
		}
	}

	if t.base != nil {
		return t.base.ExtendsName(name)
	}

	return false
}

func CommonBaseType(ti1 TypeSpecAccessor, ti2 TypeSpecAccessor) TypeSpecAccessor {
	for t1 := ti1; t1 != nil; t1 = t1.Base() {
		for t2 := ti2; t2 != nil; t2 = t2.Base() {
			if t1.EqualType(t2) {
				return t1
			}
		}
	}
	return nil
}

func newAnyTypeSpec(name string) TypeSpecAccessor {
	return newAnyTypeSpecWithBase(name, anyTypeSpec)
}

func newAnyTypeSpecWithBase(name string, base TypeSpecAccessor) TypeSpecAccessor {
	return NewTypeSpecWithBase(NewFQTypeName(name, NamespaceName), base)
}
