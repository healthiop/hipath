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
	"strings"
)

var UCUMSystemURI = NewString("http://unitsofmeasure.org")

var QuantityTypeSpec = newAnyTypeSpec("Quantity")

type quantityType struct {
	baseAnyType
	value DecimalAccessor
	unit  StringAccessor
}

type QuantityAccessor interface {
	AnyAccessor
	Comparator
	Stringifier
	DecimalValueAccessor
	Negator
	ArithmeticApplier

	Unit() StringAccessor
	ToUnit(unit StringAccessor) QuantityAccessor
}

func NewQuantity(value DecimalAccessor, unit StringAccessor) QuantityAccessor {
	return NewQuantityWithSource(value, unit, nil)
}

func NewQuantityWithSource(value DecimalAccessor, unit StringAccessor, source interface{}) QuantityAccessor {
	if value == nil {
		panic("value must not be nil")
	}
	return &quantityType{
		baseAnyType: baseAnyType{
			source: source,
		},
		value: value,
		unit:  unit,
	}
}

func (t *quantityType) DataType() DataTypes {
	return QuantityDataType
}

func (t *quantityType) Value() DecimalAccessor {
	return t.value
}

func (t *quantityType) Unit() StringAccessor {
	return t.unit
}

func (t *quantityType) WithValue(node NumberAccessor) DecimalValueAccessor {
	var value DecimalAccessor
	if node == nil {
		return nil
	} else if node.DataType() == DecimalDataType {
		value = node.(DecimalAccessor)
	} else {
		value = NewDecimal(node.Decimal())
	}
	return NewQuantity(value, t.Unit())
}

func (t *quantityType) ArithmeticOpSupported(op ArithmeticOps) bool {
	return op == AdditionOp ||
		op == SubtractionOp ||
		op == MultiplicationOp ||
		op == DivisionOp
}

func (t *quantityType) Negate() AnyAccessor {
	return NewQuantity(t.value.Negate().(DecimalAccessor), t.unit)
}

func (e *quantityType) TypeSpec() TypeSpecAccessor {
	return QuantityTypeSpec
}

func (t *quantityType) Equal(node interface{}) bool {
	return quantityValueEqual(t, node, false)
}

func (t *quantityType) Equivalent(node interface{}) bool {
	return quantityValueEqual(t, node, true)
}

func quantityValueEqual(t QuantityAccessor, node interface{}, equivalent bool) bool {
	if q, ok := node.(QuantityAccessor); ok {
		if Equal(t.Unit(), q.Unit()) {
			return Equal(t.Value(), q.Value())
		}

		u1, exp1 := QuantityUnitWithNameString(t.Unit())
		u2, exp2 := QuantityUnitWithNameString(q.Unit())
		if exp1 != exp2 {
			return false
		}
		v1, v2, u := ConvertUnitToBase(t.Value(), u1, exp1, q.Value(), u2, exp2, !equivalent)
		if u != nil {
			return Equal(v1, v2)
		}
	} else if d, ok := node.(DecimalValueAccessor); ok {
		v1 := t.Value()
		v2 := d.Value()
		if equivalent {
			return Equivalent(v1, v2)
		}
		return Equal(v1, v2)
	}
	return false
}

func (t *quantityType) Compare(comparator Comparator) (int, OperatorStatus) {
	if q, ok := comparator.(QuantityAccessor); ok {
		if !Equal(t.Unit(), q.Unit()) {
			u1, exp1 := QuantityUnitWithNameString(t.Unit())
			u2, exp2 := QuantityUnitWithNameString(q.Unit())
			if exp1 == exp2 {
				v1, v2, u := ConvertUnitToBase(t.Value(), u1, exp1, q.Value(), u2, exp2, true)
				if u != nil {
					return decimalValueCompare(v1, v2)
				}
			}

			return -1, Empty
		}
	}

	return decimalValueCompare(t.value, comparator)
}

func (t *quantityType) ToUnit(unit StringAccessor) QuantityAccessor {
	u2, exp2 := QuantityUnitWithNameString(unit)
	if u2 == nil {
		return nil
	}

	u1, exp1 := QuantityUnitWithNameString(t.Unit())
	if u1 == nil || exp1 != exp2 {
		return nil
	}

	if u1.Equal(u2) {
		return t
	}

	u := u1.CommonBase(u2, true)
	if u == nil {
		return nil
	}

	f1, f2 := u1.Factor(u, exp1), u2.Factor(u, exp2)
	v, _ := t.Value().Calc(f1, MultiplicationOp)
	v, _ = v.Value().Calc(f2, DivisionOp)

	val := v.Value()
	return NewQuantity(val, u2.NameWithExp(val, exp2))
}

func (t *quantityType) String() string {
	var b strings.Builder
	b.Grow(32)
	b.WriteString(t.value.String())
	if t.unit != nil {
		b.WriteByte(' ')
		b.WriteByte('\'')
		b.WriteString(t.unit.String())
		b.WriteByte('\'')
	}
	return b.String()
}

func (t *quantityType) Calc(operand DecimalValueAccessor, op ArithmeticOps) (DecimalValueAccessor, error) {
	if t.value == nil || operand == nil {
		return nil, nil
	}

	if !t.ArithmeticOpSupported(op) || !operand.ArithmeticOpSupported(op) {
		return nil, fmt.Errorf("arithmetic operator not supported: %c", op)
	}

	var valLeft, varRight DecimalAccessor
	var unit QuantityUnitAccessor
	var exp int
	if q, ok := operand.(QuantityAccessor); !ok {
		valLeft, varRight = t.Value(), operand.Value()
		unit, exp = QuantityUnitWithNameString(t.Unit())
	} else {
		var err error
		valLeft, varRight, unit, exp, err = mergeQuantityUnits(t, q, op)
		if err != nil {
			return nil, err
		}
	}

	value, _ := valLeft.Calc(varRight, op)
	return NewQuantity(value.Value(), unit.NameWithExp(value.Value(), exp)), nil
}

func mergeQuantityUnits(l QuantityAccessor, r QuantityAccessor, op ArithmeticOps) (DecimalAccessor, DecimalAccessor, QuantityUnitAccessor, int, error) {
	leftVal, rightVal := l.Value(), r.Value()
	leftUnit, leftExp := QuantityUnitWithNameString(l.Unit())
	rightUnit, rightExp := QuantityUnitWithNameString(r.Unit())

	var unit QuantityUnitAccessor
	if leftUnit == nil && rightUnit == nil {
		return leftVal, rightVal, EmptyQuantityUnit, 1, nil
	}
	if leftUnit != nil && leftUnit.Equal(rightUnit) {
		unit = leftUnit
	} else {
		leftVal, rightVal, unit = ConvertUnitToMostGranular(
			leftVal, leftUnit, leftExp, rightVal, rightUnit, rightExp, true)
		if unit == nil {
			return nil, nil, nil, 1, fmt.Errorf("units are not equal: %s != %s",
				leftUnit, rightUnit)
		}
	}

	exp := leftExp
	switch op {
	case AdditionOp, SubtractionOp:
		if leftExp != rightExp {
			return nil, nil, nil, 1, fmt.Errorf("units exponents are not equal: %d != %d",
				leftExp, rightExp)
		}
	case MultiplicationOp:
		exp = leftExp + rightExp
	case DivisionOp:
		exp = leftExp - rightExp
	}

	if exp < 1 || exp > 3 {
		return nil, nil, nil, 1, fmt.Errorf("resulting unit exponent is invalid (must be between 1 and 3): %d", exp)
	}

	return leftVal, rightVal, unit, exp, nil
}

func (t *quantityType) Abs() DecimalValueAccessor {
	return NewQuantity(t.Value().Abs().(DecimalAccessor), t.Unit())
}
