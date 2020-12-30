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
	"math"
	"math/big"
)

var DecimalTypeInfo = newAnyTypeInfo("Decimal")

var (
	DecimalZero    = NewDecimalInt(0)
	DecimalOne     = NewDecimalInt(1)
	DecimalTwo     = NewDecimalInt(2)
	DecimalThree   = NewDecimalInt(3)
	DecimalDotFive = NewDecimalFloat64(.5)
	DecimalE       = NewDecimalFloat64(math.E)
)

type decimalType struct {
	baseAnyType
	value decimal.Decimal
}

type DecimalAccessor interface {
	NumberAccessor
	Primitive() decimal.Decimal
}

func NewDecimal(value decimal.Decimal) DecimalAccessor {
	return NewDecimalWithSource(value, nil)
}

func NewDecimalWithSource(value decimal.Decimal, source interface{}) DecimalAccessor {
	return newDecimal(value, source)
}

func NewDecimalInt(value int32) DecimalAccessor {
	return NewDecimalIntWithSource(value, nil)
}

func NewDecimalIntWithSource(value int32, source interface{}) DecimalAccessor {
	return newDecimal(decimal.NewFromInt32(value), source)
}

func NewDecimalInt64(value int64) DecimalAccessor {
	return NewDecimalInt64WithSource(value, nil)
}

func NewDecimalInt64WithSource(value int64, source interface{}) DecimalAccessor {
	return newDecimal(decimal.NewFromInt(value), source)
}

func NewDecimalFloat64(value float64) DecimalAccessor {
	return NewDecimalFloat64WithSource(value, nil)
}

func NewDecimalFloat64WithSource(value float64, source interface{}) DecimalAccessor {
	return newDecimal(decimal.NewFromFloat(value), source)
}

func DecimalOfInt(value int32) DecimalAccessor {
	switch value {
	case 0:
		return DecimalZero
	case 1:
		return DecimalOne
	case 2:
		return DecimalTwo
	case 3:
		return DecimalThree
	default:
		return NewDecimalInt(value)
	}
}

func ParseDecimal(value string) (DecimalAccessor, error) {
	if d, err := decimal.NewFromString(value); err != nil {
		return nil, fmt.Errorf("not a decimal: %s", value)
	} else {
		return newDecimal(d, nil), nil
	}
}

func newDecimal(value decimal.Decimal, source interface{}) DecimalAccessor {
	return &decimalType{
		baseAnyType: baseAnyType{
			source: source,
		},
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

func (t *decimalType) Primitive() decimal.Decimal {
	return t.value
}

func (t *decimalType) One() bool {
	return DecimalOne.Decimal().Equal(t.value)
}

func (t *decimalType) Positive() bool {
	return t.value.IsPositive()
}

func (t *decimalType) HasFraction() bool {
	if t.value.Exponent() >= 0 {
		return false
	}
	return !t.value.Equal(t.value.Truncate(0))
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
	return newDecimal(t.value.Neg(), nil)
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
	} else {
		return false
	}

	d1, d2 := leastPrecisionDecimal(t.Decimal(), d.Decimal())
	return d1.Equal(d2)
}

func (t *decimalType) Compare(comparator Comparator) (int, OperatorStatus) {
	return decimalValueCompare(t, comparator)
}

func decimalValueCompare(t NumberAccessor, comparator Comparator) (int, OperatorStatus) {
	var d DecimalAccessor
	if da, ok := comparator.(DecimalAccessor); ok {
		d = da
	} else if da, ok := comparator.(DecimalValueAccessor); ok {
		d = da.Value()
	} else {
		return -1, Inconvertible
	}

	return t.Decimal().Cmp(d.Decimal()), Evaluated
}

func (t *decimalType) String() string {
	exp := t.value.Exponent()
	if exp >= 0 {
		return t.value.String()
	}
	return t.value.StringFixed(-exp)
}

func (t *decimalType) Ceiling() NumberAccessor {
	return NewInteger(int32(t.value.Ceil().IntPart()))
}

func (t *decimalType) Exp() NumberAccessor {
	if t.HasFraction() {
		return NewDecimalFloat64(math.Exp(t.Float64()))
	}
	return NewDecimal(DecimalE.Decimal().Pow(t.value))
}

func (t *decimalType) Floor() NumberAccessor {
	return NewInteger(int32(t.value.Floor().IntPart()))
}

func (t *decimalType) Ln() (NumberAccessor, error) {
	if !t.value.IsPositive() {
		return nil, fmt.Errorf("logarithmus cannot be applied to non-positive values %d", t.value)
	}
	return NewDecimalFloat64(math.Log(t.Float64())), nil
}

func (t *decimalType) Log(base NumberAccessor) (NumberAccessor, error) {
	if !t.value.IsPositive() {
		return nil, fmt.Errorf("logarithmus cannot be applied to non-positive values %f", t.Float64())
	}
	if !base.Positive() {
		return nil, fmt.Errorf("logarithmus cannot be applied to non-positive base %f", base.Float64())
	}
	return NewDecimalFloat64(math.Log(t.Float64()) / math.Log(base.Float64())), nil
}

func (t *decimalType) Power(exponent NumberAccessor) (NumberAccessor, bool) {
	if exponent.One() {
		return t, true
	}
	if exponent.HasFraction() {
		r := math.Pow(t.Float64(), exponent.Float64())
		if math.IsNaN(r) {
			return nil, false
		}
		return NewDecimalFloat64(r), true
	}
	return NewDecimal(t.value.Pow(exponent.Decimal())), true
}

func (t *decimalType) Round(precision int32) (NumberAccessor, error) {
	if precision < 0 {
		return nil, fmt.Errorf("precision must not be negative %d", precision)
	}
	return NewDecimal(t.value.Round(precision)), nil
}

func (t *decimalType) Sqrt() (NumberAccessor, bool) {
	return t.Power(DecimalDotFive)
}

func (t *decimalType) Truncate(precision int32) NumberAccessor {
	return NewDecimal(t.value.Truncate(precision))
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

func (t *decimalType) Abs() DecimalValueAccessor {
	return NewDecimal(t.value.Abs())
}

func DecimalValueFloat64(node interface{}) interface{} {
	if v, ok := node.(DecimalAccessor); !ok {
		return nil
	} else {
		return v.Float64()
	}
}
