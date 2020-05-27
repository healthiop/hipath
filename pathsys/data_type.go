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

type DataTypes int

const UndefinedDataType DataTypes = 0x0001
const CollectionDataType DataTypes = 0x0002

const LiteralDataType DataTypes = 0x0200

const (
	BooleanDataType = iota + LiteralDataType
	IntegerDataType
	DecimalDataType
	StringDataType
	DateDataType
	DateTimeDataType
	TimeDataType
	QuantityDataType
)

type AnyAccessor interface {
	DataType() DataTypes
	TypeInfo() TypeInfoAccessor
	Source() interface{}
	Equal(node interface{}) bool
	Equivalent(node interface{}) bool
}

type baseAnyType struct {
	source interface{}
}

func (a *baseAnyType) Source() interface{} {
	return a.source
}

func TypeEqual(n1 AnyAccessor, n2 AnyAccessor) bool {
	return n1 != nil && n2 != nil && n1.DataType() == n2.DataType()
}

func Equal(n1 AnyAccessor, n2 AnyAccessor) bool {
	return n1 == n2 || (n1 != nil && n2 != nil && n1.Equal(n2))
}

func Equivalent(n1 AnyAccessor, n2 AnyAccessor) bool {
	return n1 == n2 || (n1 != nil && n2 != nil && n1.Equivalent(n2))
}

type OperatorStatus int

const (
	Inconvertible OperatorStatus = iota
	Empty
	Evaluated
)

type Comparator interface {
	AnyAccessor
	Compare(comparator Comparator) (int, OperatorStatus)
}

type Negator interface {
	AnyAccessor
	Negate() AnyAccessor
}

type Stringifier interface {
	AnyAccessor
	String() string
}
