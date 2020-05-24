// Copyright (c) 2020, Volker Schmidt (volker@volsch.eu)
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source URI must retain the above copyright notice, this
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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUCUMSystemURI(t *testing.T) {
	assert.Equal(t, "http://unitsofmeasure.org", UCUMSystemURI.String())
}

func TestQuantityDataType(t *testing.T) {
	o := NewQuantity(NewDecimalInt(0), nil)
	dataType := o.DataType()
	assert.Equal(t, QuantityDataType, dataType)
}

func TestQuantityTypeInfo(t *testing.T) {
	o := NewQuantity(NewDecimalInt(10), NewString("g"))
	i := o.TypeInfo()
	if assert.NotNil(t, i, "type info expected") {
		assert.Equal(t, "System.Quantity", i.String())
		if assert.NotNil(t, i.FQBaseName(), "base name expected") {
			assert.Equal(t, "System.Any", i.FQBaseName().String())
		}
	}
}

func TestQuantityValueNil(t *testing.T) {
	assert.Panics(t, func() { NewQuantity(nil, NewString("s")) })
}

func TestQuantity(t *testing.T) {
	o := NewQuantity(NewDecimalFloat64(47.1), NewString("g"))
	if assert.NotNil(t, o.Value()) {
		assert.Equal(t, 47.1, o.Value().Float64())
	}
	if assert.NotNil(t, o.Unit()) {
		assert.Equal(t, "g", o.Unit().String())
	}
}

func TestQuantityNegate(t *testing.T) {
	o := NewQuantity(NewDecimalFloat64(47.1), NewString("g"))
	n := o.Negate().(QuantityAccessor)
	if assert.NotNil(t, n.Value()) {
		assert.Equal(t, -47.1, n.Value().Float64())
	}
	if assert.NotNil(t, n.Unit()) {
		assert.Equal(t, "g", n.Unit().String())
	}
}

func TestQuantityEqual(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.1), NewString("g"))
	q2 := NewQuantity(NewDecimalFloat64(47.1), NewString("g"))
	assert.Equal(t, true, q1.Equal(q2))
	assert.Equal(t, true, q1.Equivalent(q2))
}

func TestQuantityEqualTypeDiffers(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(47.1), NewString("g"))
	assert.Equal(t, false, q.Equal(NewString("")))
	assert.Equal(t, false, q.Equivalent(NewString("")))
}

func TestQuantityEqualValueDiffers(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.1), NewString("g"))
	q2 := NewQuantity(NewDecimalFloat64(47.2), NewString("g"))
	assert.Equal(t, false, q1.Equal(q2))
	assert.Equal(t, false, q1.Equivalent(q2))
}

func TestQuantityEqualUnitDiffers(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.1), NewString("g"))
	q2 := NewQuantity(NewDecimalFloat64(47.1), NewString("kg"))
	assert.Equal(t, false, q1.Equal(q2))
	assert.Equal(t, false, q1.Equivalent(q2))
}

func TestQuantityEqualInteger(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47), NewString("g"))
	assert.Equal(t, true, q1.Equal(NewInteger(47)))
	assert.Equal(t, true, q1.Equivalent(NewInteger(47)))
}

func TestQuantityEqualIntegerNot(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(48), NewString("g"))
	assert.Equal(t, false, q1.Equal(NewInteger(47)))
	assert.Equal(t, false, q1.Equivalent(NewInteger(47)))
}

func TestQuantityEqualDecimal(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.1), NewString("g"))
	assert.Equal(t, true, q1.Equal(NewDecimalFloat64(47.1)))
	assert.Equal(t, true, q1.Equivalent(NewDecimalFloat64(47.1)))
}

func TestQuantityEqualNotEqualDecimal(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.2), NewString("g"))
	assert.Equal(t, false, q1.Equal(NewDecimalFloat64(47.1)))
	assert.Equal(t, false, q1.Equivalent(NewDecimalFloat64(47.1)))
}

func TestQuantityEquivalentDecimal(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.12), NewString("g"))
	assert.Equal(t, false, q1.Equal(NewDecimalFloat64(47.1)))
	assert.Equal(t, true, q1.Equivalent(NewDecimalFloat64(47.1)))
}

func TestQuantityStringValueOnly(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(47.1), nil)
	assert.Equal(t, "47.1", q.String())
}

func TestQuantityString(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(47.1), NewString("g"))
	assert.Equal(t, "47.1 g", q.String())
}

func TestQuantityWithValueNil(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(47.1), NewString("g"))
	r := q.WithValue(nil)
	assert.Nil(t, r)
}

func TestQuantityWithValueNotNil(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(47.1), NewString("g"))
	r := q.WithValue(NewDecimalFloat64(12.1))
	e := NewQuantity(NewDecimalFloat64(12.1), NewString("g"))
	assert.True(t, e.Equal(r))
}

func TestQuantityWithValueInteger(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(47.1), NewString("g"))
	r := q.WithValue(NewInteger(12))
	e := NewQuantity(NewDecimalFloat64(12), NewString("g"))
	assert.True(t, e.Equal(r))
}

func TestQuantityArithmeticOpSupported(t *testing.T) {
	d := NewQuantity(NewDecimalInt(10), NewString("g"))
	assert.True(t, d.ArithmeticOpSupported(AdditionOp), "addition must be supported")
	assert.True(t, d.ArithmeticOpSupported(SubtractionOp), "subtraction must be supported")
	assert.True(t, d.ArithmeticOpSupported(MultiplicationOp), "multiplication must be supported")
	assert.True(t, d.ArithmeticOpSupported(DivisionOp), "division must be supported")
	assert.False(t, d.ArithmeticOpSupported(DivOp), "DIV must not be supported")
	assert.False(t, d.ArithmeticOpSupported(ModOp), "MOD must not be supported")
}

func TestExtractQuantityCodeExpNil(t *testing.T) {
	unit, exp := extractQuantityCodeExp(nil)
	assert.Equal(t, "", unit)
	assert.Equal(t, 0, exp)
}

func TestExtractQuantityCodeExpSingleCharUnit(t *testing.T) {
	unit, exp := extractQuantityCodeExp(NewString("m"))
	assert.Equal(t, "m", unit)
	assert.Equal(t, 1, exp)
}

func TestExtractQuantityCodeExpMultiCharUnit(t *testing.T) {
	unit, exp := extractQuantityCodeExp(NewString("cm"))
	assert.Equal(t, "cm", unit)
	assert.Equal(t, 1, exp)
}

func TestExtractQuantityCodeExpExp(t *testing.T) {
	unit, exp := extractQuantityCodeExp(NewString("m2"))
	assert.Equal(t, "m", unit)
	assert.Equal(t, 2, exp)
}

func TestMergeQuantityUnitsCodesDiffers(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.2),
		NewString("g"))
	q2 := NewQuantity(NewDecimalFloat64(47.2),
		NewString("m"))
	code, err := mergeQuantityUnits(q1, q2, AdditionOp)
	assert.Error(t, err, "error expected")
	assert.Nil(t, code, "no code expected")
}

func TestMergeQuantityUnitsCodesNil(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.2), nil)
	q2 := NewQuantity(NewDecimalFloat64(47.2), nil)
	code, err := mergeQuantityUnits(q1, q2, AdditionOp)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, code)
}

func TestMergeQuantityUnitsAdditionExpsDiffers(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.2),
		NewString("m"))
	q2 := NewQuantity(NewDecimalFloat64(47.2),
		NewString("m2"))
	code, err := mergeQuantityUnits(q1, q2, AdditionOp)
	assert.Error(t, err, "error expected")
	assert.Nil(t, code, "no code expected")
}

func TestMergeQuantityUnitsSubtractionExpsDiffers(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.2), NewString("m"))
	q2 := NewQuantity(NewDecimalFloat64(47.2), NewString("m2"))
	code, err := mergeQuantityUnits(q1, q2, SubtractionOp)
	assert.Error(t, err, "error expected")
	assert.Nil(t, code, "no code expected")
}

func TestMergeQuantityUnitsMultiplicationExpsError(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.2), NewString("m2"))
	q2 := NewQuantity(NewDecimalFloat64(47.2), NewString("m2"))
	code, err := mergeQuantityUnits(q1, q2, MultiplicationOp)
	assert.Error(t, err, "error expected")
	assert.Nil(t, code, "no code expected")
}

func TestMergeQuantityUnitsDivisionExpError(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.2), NewString("m2"))
	q2 := NewQuantity(NewDecimalFloat64(47.2), NewString("m2"))
	code, err := mergeQuantityUnits(q1, q2, MultiplicationOp)
	assert.Error(t, err, "error expected")
	assert.Nil(t, code, "no code expected")
}

func TestQuantityCalcAddition(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.2), NewString("m"))
	q2 := NewQuantity(NewDecimalFloat64(21.7), NewString("m"))
	r, err := q1.Calc(q2, AdditionOp)
	e := NewQuantity(NewDecimalFloat64(68.9), NewString("m"))
	assert.NoError(t, err, "no error expected")
	assert.True(t, e.Equal(r))
}

func TestQuantityCalcAdditionValue(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.2), NewString("m"))
	r, err := q1.Calc(NewDecimalInt(20), AdditionOp)
	e := NewQuantity(NewDecimalFloat64(67.2), NewString("m"))
	assert.NoError(t, err, "no error expected")
	assert.True(t, e.Equal(r))
}

func TestQuantityCalcSubtraction(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.2), NewString("m"))
	q2 := NewQuantity(NewDecimalFloat64(21.7), NewString("m"))
	r, err := q1.Calc(q2, SubtractionOp)
	e := NewQuantity(NewDecimalFloat64(25.5), NewString("m"))
	assert.NoError(t, err, "no error expected")
	assert.True(t, e.Equal(r))
}

func TestQuantityCalcMultiplication(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.2), NewString("m"))
	q2 := NewQuantity(NewDecimalFloat64(21.7), NewString("m"))
	r, err := q1.Calc(q2, MultiplicationOp)
	e := NewQuantity(NewDecimalFloat64(1024.24), NewString("m2"))
	assert.NoError(t, err, "no error expected")
	assert.True(t, e.Equal(r))
}

func TestQuantityCalcDivision(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(48.75), NewString("m3"))
	q2 := NewQuantity(NewDecimalFloat64(2.5), NewString("m"))
	r, err := q1.Calc(q2, DivisionOp)
	e := NewQuantity(NewDecimalFloat64(19.5), NewString("m2"))
	assert.NoError(t, err, "no error expected")
	assert.True(t, e.Equal(r))
}

func TestQuantityCalcNotSupportedOp(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(48.75), NewString("m3"))
	q2 := NewQuantity(NewDecimalFloat64(2.5), NewString("m"))
	r, err := q1.Calc(q2, DivOp)
	assert.Error(t, err, "error expected")
	assert.Nil(t, r, "no res expected")
}

func TestQuantityCalcCodesDiffers(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(48.75), NewString("m"))
	q2 := NewQuantity(NewDecimalFloat64(2.5), NewString("g"))
	r, err := q1.Calc(q2, AdditionOp)
	assert.Error(t, err, "error expected")
	assert.Nil(t, r, "no res expected")
}

func TestQuantityCalcRightNil(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(2.5), NewString("g"))
	r, err := q1.Calc(nil, AdditionOp)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, r, "empty res expected")
}

func TestQuantityCompareEqual(t *testing.T) {
	res, status := NewQuantity(NewDecimalFloat64(10.21), NewString("cm")).
		Compare(NewQuantity(NewDecimalFloat64(10.21), NewString("cm")))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, 0, res)
}

func TestQuantityCompareEqualUnitNil(t *testing.T) {
	res, status := NewQuantity(NewDecimalFloat64(10.21), nil).
		Compare(NewQuantity(NewDecimalFloat64(10.21), nil))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, 0, res)
}

func TestQuantityCompareEqualNotOneUnitNil(t *testing.T) {
	res, status := NewQuantity(NewDecimalFloat64(10.21), NewString("cm")).
		Compare(NewQuantity(NewDecimalFloat64(10.21), nil))
	assert.Equal(t, Empty, status)
	assert.Equal(t, -1, res)
}

func TestQuantityCompareEqualNotUnit(t *testing.T) {
	res, status := NewQuantity(NewDecimalFloat64(10.21), NewString("cm")).
		Compare(NewQuantity(NewDecimalFloat64(10.21), NewString("m")))
	assert.Equal(t, Empty, status)
	assert.Equal(t, -1, res)
}

func TestQuantityCompareEqualTypeDiffers(t *testing.T) {
	res, status := NewQuantity(NewDecimalFloat64(10.21), NewString("cm")).
		Compare(NewString("test1"))
	assert.Equal(t, Inconvertible, status)
	assert.Equal(t, -1, res)
}

func TestQuantityCompareLessThan(t *testing.T) {
	res, status := NewQuantity(NewDecimalFloat64(10.21), NewString("cm")).
		Compare(NewQuantity(NewDecimalFloat64(10.25), NewString("cm")))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, -1, res)
}

func TestQuantityCompareGreaterThan(t *testing.T) {
	res, status := NewQuantity(NewDecimalFloat64(10.21), NewString("cm")).
		Compare(NewQuantity(NewDecimalFloat64(10.11), NewString("cm")))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, 1, res)
}

func TestQuantityCompareInteger(t *testing.T) {
	res, status := NewQuantity(NewDecimalFloat64(10.64), NewString("cm")).
		Compare(NewInteger(11))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, -1, res)
}

func TestQuantityCompareQuantity(t *testing.T) {
	res, status := NewQuantity(NewDecimalFloat64(10.64), NewString("cm")).Compare(
		NewQuantity(NewDecimalFloat64(10.71), NewString("cm")))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, -1, res)
}
