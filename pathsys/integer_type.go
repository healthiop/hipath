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
	"strconv"
)

var IntegerTypeInfo = newAnyTypeInfo("Integer")

type integerType struct {
	baseAnyType
	value        int32
	decimalValue DecimalAccessor
}

type IntegerAccessor interface {
	NumberAccessor
	Primitive() int32
}

func NewInteger(value int32) IntegerAccessor {
	return NewIntegerWithSource(value, nil)
}

func NewIntegerWithSource(value int32, source interface{}) IntegerAccessor {
	return newInteger(value, source)
}

func ParseInteger(value string) (IntegerAccessor, error) {
	if i, err := strconv.Atoi(value); err != nil {
		return nil, fmt.Errorf("not an integer: %s", value)
	} else {
		return NewInteger(int32(i)), nil
	}
}

func newInteger(value int32, source interface{}) IntegerAccessor {
	return &integerType{
		baseAnyType: baseAnyType{
			source: source,
		},
		value: value,
	}
}

func (t *integerType) DataType() DataTypes {
	return IntegerDataType
}

func (t *integerType) Int() int32 {
	return t.value
}

func (t *integerType) Int64() int64 {
	return int64(t.value)
}

func (t *integerType) Float32() float32 {
	return float32(t.value)
}

func (t *integerType) Float64() float64 {
	return float64(t.value)
}

func (t *integerType) BigFloat() *big.Float {
	return t.Decimal().BigFloat()
}

func (t *integerType) Decimal() decimal.Decimal {
	return t.Value().Decimal()
}

func (t *integerType) Primitive() int32 {
	return t.value
}

func (t *integerType) One() bool {
	return t.value == 1
}

func (t *integerType) Positive() bool {
	return t.value > 0
}

func (t *integerType) HasFraction() bool {
	return false
}

func (t *integerType) TypeInfo() TypeInfoAccessor {
	return IntegerTypeInfo
}

func (t *integerType) Value() DecimalAccessor {
	if t.decimalValue == nil {
		t.decimalValue = NewDecimalInt(t.value)
	}
	return t.decimalValue
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
	return newInteger(-t.value, nil)
}

func (t *integerType) Equal(node interface{}) bool {
	if o, ok := node.(IntegerAccessor); ok {
		return t.Int() == o.Int()
	}

	return decimalValueEqual(t, node)
}

func (t *integerType) Equivalent(node interface{}) bool {
	if o, ok := node.(IntegerAccessor); ok {
		return t.Int() == o.Int()
	}

	return decimalValueEquivalent(t, node)
}

func (t *integerType) Compare(comparator Comparator) (int, OperatorStatus) {
	if TypeEqual(t, comparator) {
		l, r := t.value, comparator.(IntegerAccessor).Int()
		if l == r {
			return 0, Evaluated
		}
		if l < r {
			return -1, Evaluated
		}
		return 1, Evaluated
	}

	return decimalValueCompare(t, comparator)
}

func (t *integerType) String() string {
	return strconv.FormatInt(int64(t.value), 10)
}

func (t *integerType) Ceiling() NumberAccessor {
	return t
}

func (t *integerType) Exp() NumberAccessor {
	return NewDecimalFloat64(math.Exp(t.Float64()))
}

func (t *integerType) Floor() NumberAccessor {
	return t
}

func (t *integerType) Ln() (NumberAccessor, error) {
	if t.value <= 0 {
		return nil, fmt.Errorf("logarithmus cannot be applied to non-positive values %d", t.value)
	}
	return NewDecimalFloat64(math.Log(t.Float64())), nil
}

func (t *integerType) Log(base NumberAccessor) (NumberAccessor, error) {
	if t.value <= 0 {
		return nil, fmt.Errorf("logarithmus cannot be applied to non-positive values %d", t.value)
	}
	if !base.Positive() {
		return nil, fmt.Errorf("logarithmus cannot be applied to non-positive base %f", base.Float64())
	}
	return NewDecimalFloat64(math.Log(t.Float64()) / math.Log(base.Float64())), nil
}

func (t *integerType) Power(exponent NumberAccessor) (NumberAccessor, bool) {
	if exponent.One() {
		return t, true
	}
	if exponent.DataType() == IntegerDataType {
		return NewInteger(int32(math.Pow(t.Float64(), exponent.Float64()))), true
	}
	return NewDecimalInt(t.Int()).Power(exponent)
}

func (t *integerType) Round(precision int32) (NumberAccessor, error) {
	if precision < 0 {
		return nil, fmt.Errorf("precision must not be negative %d", precision)
	}
	return t, nil
}

func (t *integerType) Sqrt() (NumberAccessor, bool) {
	r := math.Sqrt(t.Float64())
	if math.IsNaN(r) {
		return nil, false
	}
	return NewDecimalFloat64(r), true
}

func (t *integerType) Truncate(int32) NumberAccessor {
	return t
}

func (t *integerType) Calc(operand DecimalValueAccessor, op ArithmeticOps) (DecimalValueAccessor, error) {
	if operand == nil {
		return nil, nil
	}

	if !t.ArithmeticOpSupported(op) || !operand.ArithmeticOpSupported(op) {
		return nil, fmt.Errorf("arithmetic operator not supported: %c", op)
	}

	if ov, ok := operand.(IntegerAccessor); ok {
		pov := ov.Primitive()
		switch op {
		case AdditionOp:
			return NewInteger(t.Int() + pov), nil
		case SubtractionOp:
			return NewInteger(t.Int() - pov), nil
		case MultiplicationOp:
			return NewInteger(t.Int() * pov), nil
		case DivisionOp:
			if pov == 0 {
				return nil, nil
			}
			return NewDecimalFloat64(float64(t.Int()) / float64(pov)), nil
		case DivOp:
			if pov == 0 {
				return nil, nil
			}
			return NewInteger(t.Int() / pov), nil
		case ModOp:
			if pov == 0 {
				return nil, nil
			}
			return NewInteger(t.Int() % pov), nil
		default:
			panic(fmt.Sprintf("Unhandled operator: %d", op))
		}
	}

	return operand.WithValue(decimalCalc(t, operand.Value(), op)), nil
}

func (t *integerType) Abs() DecimalValueAccessor {
	return NewInteger(int32(math.Abs(float64(t.Int()))))
}

func IntegerValue(node interface{}) interface{} {
	if v, ok := node.(IntegerAccessor); !ok {
		return nil
	} else {
		return v.Int()
	}
}
