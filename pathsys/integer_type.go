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

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strconv"
)

var IntegerTypeInfo = newAnyTypeInfo("Integer")

type integerType struct {
	value int32
}

type IntegerAccessor interface {
	NumberAccessor
}

func NewInteger(value int32) IntegerAccessor {
	return newInteger(value)
}

func ParseInteger(value string) (IntegerAccessor, error) {
	if i, err := strconv.Atoi(value); err != nil {
		return nil, fmt.Errorf("not an integer: %s", value)
	} else {
		return NewInteger(int32(i)), nil
	}
}

func newInteger(value int32) IntegerAccessor {
	return &integerType{
		value: value,
	}
}

func (t *integerType) DataType() DataTypes {
	return IntegerDataType
}

func (t *integerType) Int() int32 {
	return t.value
}

func (t *integerType) Float64() float64 {
	return float64(t.value)
}

func (t *integerType) Decimal() decimal.Decimal {
	return decimal.NewFromInt32(t.value)
}

func (t *integerType) TypeInfo() TypeInfoAccessor {
	return IntegerTypeInfo
}

func (t *integerType) Value() DecimalAccessor {
	return NewDecimalInt(t.value)
}

func (t *integerType) WithValue(node NumberAccessor) DecimalValueAccessor {
	if node == nil || node.DataType() == IntegerDataType {
		return node
	}

	return NewInteger(node.Int())
}

func (t *integerType) ArithmeticOpSupported(ArithmeticOps) bool {
	return true
}

func (t *integerType) Negate() AnyAccessor {
	return newInteger(-t.value)
}

func (t *integerType) Equal(node interface{}) bool {
	if o, ok := node.(AnyAccessor); ok {
		if o.DataType() == IntegerDataType {
			return t.Int() == o.(IntegerAccessor).Int()
		}
	}

	return decimalValueEqual(t, node)
}

func (t *integerType) Equivalent(node interface{}) bool {
	if o, ok := node.(AnyAccessor); ok {
		if o.DataType() == IntegerDataType {
			return t.Int() == o.(IntegerAccessor).Int()
		}
	}

	return decimalValueEquivalent(t, node)
}

func (t *integerType) String() string {
	return strconv.FormatInt(int64(t.value), 10)
}

func (t *integerType) Calc(operand DecimalValueAccessor, op ArithmeticOps) (DecimalValueAccessor, error) {
	if operand == nil {
		return nil, nil
	}

	if !t.ArithmeticOpSupported(op) || !operand.ArithmeticOpSupported(op) {
		return nil, fmt.Errorf("arithmetic operator not supported: %c", op)
	}

	if operand.DataType() == IntegerDataType {
		operandValue := operand.(IntegerAccessor).Int()
		switch op {
		case AdditionOp:
			return NewInteger(t.Int() + operandValue), nil
		case SubtractionOp:
			return NewInteger(t.Int() - operandValue), nil
		case MultiplicationOp:
			return NewInteger(t.Int() * operandValue), nil
		case DivisionOp:
			if operandValue == 0 {
				return nil, nil
			}
			return NewDecimalFloat64(float64(t.Int()) / float64(operandValue)), nil
		case DivOp:
			if operandValue == 0 {
				return nil, nil
			}
			return NewInteger(t.Int() / operandValue), nil
		case ModOp:
			if operandValue == 0 {
				return nil, nil
			}
			return NewInteger(t.Int() % operandValue), nil
		default:
			panic(fmt.Sprintf("Unhandled operator: %d", op))
		}
	}

	return operand.WithValue(decimalCalc(t, operand.Value(), op)), nil
}
