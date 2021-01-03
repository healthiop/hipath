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
	"fmt"
)

var BooleanTypeSpec = newAnyTypeSpec("Boolean")

type booleanType struct {
	baseAnyType
	value bool
}

type BooleanAccessor interface {
	AnyAccessor
	Comparator
	Stringifier
	Negator
	Bool() bool
}

var (
	True  = NewBooleanWithSource(true, nil)
	False = NewBooleanWithSource(false, nil)
)

func BooleanOf(value bool) BooleanAccessor {
	if value {
		return True
	}
	return False
}

func NewBoolean(value bool) BooleanAccessor {
	return NewBooleanWithSource(value, nil)
}

func NewBooleanWithSource(value bool, source interface{}) BooleanAccessor {
	return newBoolean(value, source)
}

func newBoolean(value bool, source interface{}) BooleanAccessor {
	return &booleanType{
		baseAnyType: baseAnyType{
			source: source,
		},
		value: value,
	}
}

func ParseBoolean(value string) (BooleanAccessor, error) {
	switch value {
	case "true":
		return True, nil
	case "false":
		return False, nil
	}
	return nil, fmt.Errorf("not a boolean: %s", value)
}

func (t *booleanType) DataType() DataTypes {
	return BooleanDataType
}

func (t *booleanType) Bool() bool {
	return t.value
}

func (t *booleanType) String() string {
	if t.value {
		return "true"
	}
	return "false"
}

func (e *booleanType) TypeSpec() TypeSpecAccessor {
	return BooleanTypeSpec
}

func (t *booleanType) Equal(node interface{}) bool {
	if !SystemAnyTypeEqual(t, node) {
		return false
	}
	return t.Bool() == node.(BooleanAccessor).Bool()
}

func (t *booleanType) Equivalent(node interface{}) bool {
	return t.Equal(node)
}

func (t *booleanType) Compare(comparator Comparator) (int, OperatorStatus) {
	if !TypeEqual(t, comparator) {
		return -1, Inconvertible
	} else {
		if t.value == comparator.(BooleanAccessor).Bool() {
			return 0, Evaluated
		}
		if !t.value {
			return -1, Evaluated
		}
		return 1, Evaluated
	}
}

func (t *booleanType) Negate() AnyAccessor {
	return BooleanOf(!t.value)
}
