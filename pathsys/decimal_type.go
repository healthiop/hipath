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
	"math/big"
)

var DecimalTypeInfo = newAnyTypeInfo("Decimal")

type decimalType struct {
	value decimal.Decimal
}

type DecimalAccessor interface {
	NumberAccessor
	Int() int32
	Int64() int64
	Float32() float32
	Float64() float64
	BigFloat() *big.Float
	Decimal() decimal.Decimal
}

func NewDecimal(value decimal.Decimal) DecimalAccessor {
	return newDecimal(value)
}

func NewDecimalInt(value int32) DecimalAccessor {
	return newDecimal(decimal.NewFromInt32(value))
}

func NewDecimalInt64(value int64) DecimalAccessor {
	return newDecimal(decimal.NewFromInt(value))
}

func NewDecimalFloat64(value float64) DecimalAccessor {
	return newDecimal(decimal.NewFromFloat(value))
}

func ParseDecimal(value string) (DecimalAccessor, error) {
	if d, err := decimal.NewFromString(value); err != nil {
		return nil, fmt.Errorf("not a decimal: %s", value)
	} else {
		return newDecimal(d), nil
	}
}

func newDecimal(value decimal.Decimal) DecimalAccessor {
	return &decimalType{
		value: value,
	}
}

func (t *decimalType) DataType() DataTypes {
	return DecimalDataType
}

func (t *decimalType) Int() int32 {
	return int32(t.value.IntPart())
}

func (t *decimalType) Int64() int64 {
	return t.value.IntPart()
}

func (t *decimalType) Float32() float32 {
	v, _ := t.value.Float64()
	return float32(v)
}

func (t *decimalType) Float64() float64 {
	v, _ := t.value.Float64()
	return v
}

func (t *decimalType) BigFloat() *big.Float {
	return t.value.BigFloat()
}

func (t *decimalType) Decimal() decimal.Decimal {
	return t.value
}

func (t *decimalType) Value() DecimalAccessor {
	return t
}

func (t *decimalType) WithValue(node NumberAccessor) DecimalValueAccessor {
	if node == nil || node.DataType() == DecimalDataType {
		return node
	}

	return NewDecimal(node.Decimal())
}

func (t *decimalType) ArithmeticOpSupported(ArithmeticOps) bool {
	return true
}

func (t *decimalType) TypeInfo() TypeInfoAccessor {
	return DecimalTypeInfo
}

func (t *decimalType) Negate() AnyAccessor {
	return newDecimal(t.value.Neg())
}

func (t *decimalType) Equal(node interface{}) bool {
	return decimalValueEqual(t, node)
}

func decimalValueEqual(t NumberAccessor, node interface{}) bool {
	var d NumberAccessor
	if da, ok := node.(NumberAccessor); ok {
		d = da
	} else if da, ok := node.(DecimalValueAccessor); ok {
		d = da.Value()
		if d == nil {
			return false
		}
	} else {
		return false
	}

	return t.Decimal().Equal(d.Decimal())
}

func (t *decimalType) Equivalent(node interface{}) bool {
	return decimalValueEquivalent(t, node)
}

func decimalValueEquivalent(t NumberAccessor, node interface{}) bool {
	var d DecimalAccessor
	if da, ok := node.(DecimalAccessor); ok {
		d = da
	} else if da, ok := node.(DecimalValueAccessor); ok {
		d = da.Value()
		if d == nil {
			return false
		}
	} else {
		return false
	}

	d1, d2 := leastPrecisionDecimal(t.Decimal(), d.Decimal())
	return d1.Equal(d2)
}

func (t *decimalType) String() string {
	exp := t.value.Exponent()
	if exp >= 0 {
		return t.value.String()
	}
	return t.value.StringFixed(-exp)
}

func (t *decimalType) Truncate(precision int32) NumberAccessor {
	return NewDecimal(t.Decimal().Truncate(precision))
}

func (t *decimalType) Calc(operand DecimalValueAccessor, op ArithmeticOps) (DecimalValueAccessor, error) {
	if operand == nil {
		return nil, nil
	}

	if !t.ArithmeticOpSupported(op) || !operand.ArithmeticOpSupported(op) {
		return nil, fmt.Errorf("arithmetic operator not supported: %c", op)
	}

	return operand.WithValue(decimalCalc(t, operand.Value(), op)), nil
}

func decimalCalc(leftOperand NumberAccessor, rightOperand DecimalAccessor, op ArithmeticOps) DecimalAccessor {
	if leftOperand == nil || rightOperand == nil {
		return nil
	}

	leftOperandValue := leftOperand.Decimal()
	rightOperandValue := rightOperand.Decimal()
	switch op {
	case AdditionOp:
		return NewDecimal(leftOperandValue.Add(rightOperandValue))
	case SubtractionOp:
		return NewDecimal(leftOperandValue.Sub(rightOperandValue))
	case MultiplicationOp:
		return NewDecimal(leftOperandValue.Mul(rightOperandValue))
	case DivisionOp:
		if rightOperandValue.IsZero() {
			return nil
		}
		return NewDecimal(leftOperandValue.Div(rightOperandValue))
	case DivOp:
		if rightOperandValue.IsZero() {
			return nil
		}
		return NewDecimal(leftOperandValue.Div(rightOperandValue).Truncate(0))
	case ModOp:
		if rightOperandValue.IsZero() {
			return nil
		}
		return NewDecimal(leftOperandValue.Mod(rightOperandValue))
	default:
		panic(fmt.Sprintf("Unhandled operator: %d", op))
	}
}
