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

import "strings"

var StringTypeInfo = newAnyTypeInfo("String")

var EmptyString = newString("", nil)

type stringType struct {
	baseAnyType
	value string
}

type StringAccessor interface {
	AnyAccessor
	Comparator
	Stringifier
	Length() int
}

func StringOfNil(value string) StringAccessor {
	if len(value) == 0 {
		return nil
	}
	return NewString(value)
}

func StringOf(value string) StringAccessor {
	if len(value) == 0 {
		return EmptyString
	}
	return NewString(value)
}

func NewString(value string) StringAccessor {
	return NewStringWithSource(value, nil)
}

func NewStringWithSource(value string, source interface{}) StringAccessor {
	return newString(value, source)
}

func newString(value string, source interface{}) StringAccessor {
	return &stringType{
		baseAnyType: baseAnyType{
			source: source,
		},
		value: value,
	}
}

func (t *stringType) DataType() DataTypes {
	return StringDataType
}

func (t *stringType) String() string {
	return t.value
}

func (t *stringType) Length() int {
	return len(t.value)
}

func (e *stringType) TypeInfo() TypeInfoAccessor {
	return StringTypeInfo
}

func (t *stringType) Equal(node interface{}) bool {
	if !SystemAnyTypeEqual(t, node) {
		return false
	}

	return t.String() == node.(Stringifier).String()
}

func (t *stringType) Equivalent(node interface{}) bool {
	if !SystemAnyTypeEqual(t, node) {
		return false
	}

	return NormalizedStringEqual(t.String(), node.(Stringifier).String())
}

func (t *stringType) Compare(comparator Comparator) (int, OperatorStatus) {
	if !TypeEqual(t, comparator) {
		return -1, Inconvertible
	} else {
		return strings.Compare(t.value, comparator.(StringAccessor).String()), Evaluated
	}
}
