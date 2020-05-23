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
	"regexp"
	"strconv"
	"strings"
)

var UCUMSystemURI = NewString("http://unitsofmeasure.org")

var QuantityTypeInfo = newAnyTypeInfo("Quantity")

var quantityCodeExpRegexp = regexp.MustCompile("^(.*[^\\d])([1-3])$")

type quantityType struct {
	value DecimalAccessor
	unit  StringAccessor
}

type QuantityAccessor interface {
	AnyAccessor
	Stringifier
	DecimalValueAccessor
	Negator
	ArithmeticApplier

	Unit() StringAccessor
}

func NewQuantity(value DecimalAccessor, unit StringAccessor) QuantityAccessor {
	return &quantityType{
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
		value = nil
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

func (e *quantityType) TypeInfo() TypeInfoAccessor {
	return QuantityTypeInfo
}

func (t *quantityType) Equal(node interface{}) bool {
	return quantityValueEqual(t, node, false)
}

func (t *quantityType) Equivalent(node interface{}) bool {
	return quantityValueEqual(t, node, true)
}

func quantityValueEqual(t QuantityAccessor, node interface{}, equivalent bool) bool {
	if o, ok := node.(QuantityAccessor); ok {
		return Equal(t.Value(), o.Value()) &&
			Equal(t.Unit(), o.Unit())
	}
	if d, ok := node.(DecimalValueAccessor); ok {
		v1 := t.Value()
		v2 := d.Value()
		if equivalent {
			return Equivalent(v1, v2)
		}
		return Equal(v1, v2)
	}
	return false
}

func (t *quantityType) String() string {
	var b strings.Builder
	b.Grow(32)
	if t.value != nil {
		b.WriteString(t.value.String())
	}
	if t.unit != nil {
		if b.Len() > 0 {
			b.WriteByte(' ')
		}
		b.WriteString(t.unit.String())
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

	var unit = t.unit
	if q, ok := operand.(QuantityAccessor); ok {
		var err error
		unit, err = mergeQuantityUnits(t, q, op)
		if err != nil {
			return nil, err
		}
	}

	value, _ := t.Value().Calc(operand.Value(), op)
	if value == nil {
		return nil, nil
	}

	return NewQuantity(value.Value(), unit), nil
}

func mergeQuantityUnits(l QuantityAccessor, r QuantityAccessor, op ArithmeticOps) (StringAccessor, error) {
	leftUnit, leftExp := extractQuantityCodeExp(l.Unit())
	rightUnit, rightExp := extractQuantityCodeExp(r.Unit())

	if leftUnit != rightUnit {
		return nil, fmt.Errorf("units are not equal: %s != %s",
			leftUnit, rightUnit)
	}

	if len(leftUnit) == 0 {
		return nil, nil
	}

	exp := leftExp
	switch op {
	case AdditionOp, SubtractionOp:
		if leftExp != rightExp {
			return nil, fmt.Errorf("units exponents are not equal: %d != %d",
				leftExp, rightExp)
		}
	case MultiplicationOp:
		exp = leftExp + rightExp
	case DivisionOp:
		exp = leftExp - rightExp
	}

	if exp < 1 || exp > 3 {
		return nil, fmt.Errorf("resulting unit exponent is invalid (must be between 1 and 3): %d", exp)
	}

	if exp == 1 {
		return NewString(leftUnit), nil
	}
	return NewString(leftUnit + strconv.FormatInt(int64(exp), 10)), nil
}

func extractQuantityCodeExp(string StringAccessor) (string, int) {
	if string == nil {
		return "", 0
	}

	value := string.String()
	if len(value) < 2 {
		return value, 1
	}

	parts := quantityCodeExpRegexp.FindStringSubmatch(value)
	if parts == nil {
		return value, 1
	}

	exp, _ := strconv.Atoi(parts[2])
	return parts[1], exp
}
