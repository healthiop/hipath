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

func TestQuantitySource(t *testing.T) {
	o := NewQuantityWithSource(NewDecimalInt(10), nil, "abc")
	assert.Equal(t, "abc", o.Source())
}

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

func TestQuantityEqualExpDiffers(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.1), NewString("m2"))
	q2 := NewQuantity(NewDecimalFloat64(47.1), NewString("m3"))
	assert.Equal(t, false, q1.Equal(q2))
	assert.Equal(t, false, q1.Equivalent(q2))
}

func TestQuantityEqualUnitFactor(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(7), NewString("days"))
	q2 := NewQuantity(NewDecimalFloat64(1), NewString("week"))
	assert.Equal(t, true, q1.Equal(q2))
	assert.Equal(t, true, q1.Equivalent(q2))
}

func TestQuantityEqualUnitFactorExpLeft(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(1), NewString("week2"))
	q2 := NewQuantity(NewDecimalFloat64(49), NewString("days2"))
	assert.Equal(t, true, q1.Equal(q2))
	assert.Equal(t, true, q1.Equivalent(q2))
}

func TestQuantityEqualUnitFactorExpRight(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(49), NewString("days2"))
	q2 := NewQuantity(NewDecimalFloat64(1), NewString("week2"))
	assert.Equal(t, true, q1.Equal(q2))
	assert.Equal(t, true, q1.Equivalent(q2))
}

func TestQuantityEqualUnitFactorNot(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(7), NewString("days"))
	q2 := NewQuantity(NewDecimalFloat64(2), NewString("week"))
	assert.Equal(t, false, q1.Equal(q2))
	assert.Equal(t, false, q1.Equivalent(q2))
}

func TestQuantityEqualUnitFactorEquivalent(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(7), NewString("d"))
	q2 := NewQuantity(NewDecimalFloat64(7), NewString("days"))
	assert.Equal(t, false, q1.Equal(q2))
	assert.Equal(t, true, q1.Equivalent(q2))
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
	assert.Equal(t, "47.1 'g'", q.String())
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

func TestMergeQuantityUnitsCodesDiffers(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.2),
		NewString("g"))
	q2 := NewQuantity(NewDecimalFloat64(47.2),
		NewString("m"))
	_, _, _, _, err := mergeQuantityUnits(q1, q2, AdditionOp)
	assert.Error(t, err, "error expected")
}

func TestMergeQuantityUnitsCodesNil(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.2), nil)
	q2 := NewQuantity(NewDecimalFloat64(49.2), nil)
	v1, v2, unit, exp, err := mergeQuantityUnits(q1, q2, AdditionOp)
	if assert.NoError(t, err, "no error expected") {
		assert.Equal(t, 47.2, v1.Float64())
		assert.Equal(t, 49.2, v2.Float64())
		assert.Same(t, EmptyQuantityUnit, unit)
		assert.Equal(t, 1, exp)
	}
}

func TestMergeQuantityUnitsAdditionExpsDiffers(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.2),
		NewString("m"))
	q2 := NewQuantity(NewDecimalFloat64(47.2),
		NewString("m2"))
	_, _, _, _, err := mergeQuantityUnits(q1, q2, AdditionOp)
	assert.Error(t, err, "error expected")
}

func TestMergeQuantityUnitsSubtractionExpsDiffers(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.2), NewString("m"))
	q2 := NewQuantity(NewDecimalFloat64(47.2), NewString("m2"))
	_, _, _, _, err := mergeQuantityUnits(q1, q2, SubtractionOp)
	assert.Error(t, err, "error expected")
}

func TestMergeQuantityUnitsMultiplicationExpsError(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.2), NewString("m2"))
	q2 := NewQuantity(NewDecimalFloat64(47.2), NewString("m2"))
	_, _, _, _, err := mergeQuantityUnits(q1, q2, MultiplicationOp)
	assert.Error(t, err, "error expected")
}

func TestMergeQuantityUnitsDivisionExpError(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.2), NewString("m2"))
	q2 := NewQuantity(NewDecimalFloat64(47.2), NewString("m2"))
	_, _, _, _, err := mergeQuantityUnits(q1, q2, MultiplicationOp)
	assert.Error(t, err, "error expected")
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

func TestQuantityCalcAdditionConvertUnit(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(26), NewString("hours"))
	q2 := NewQuantity(NewDecimalFloat64(2), NewString("days"))
	r, err := q1.Calc(q2, AdditionOp)
	e := NewQuantity(NewDecimalFloat64(74), NewString("hours"))
	assert.NoError(t, err, "no error expected")
	assert.True(t, e.Equal(r))
}

func TestQuantityCalcAdditionConvertUnitWithExpLeft(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(26), NewString("hours2"))
	q2 := NewQuantity(NewDecimalFloat64(2), NewString("days2"))
	r, err := q1.Calc(q2, AdditionOp)
	e := NewQuantity(NewDecimalFloat64(1178), NewString("hours2"))
	assert.NoError(t, err, "no error expected")
	assert.True(t, e.Equal(r))
}

func TestQuantityCalcAdditionConvertUnitWithExpRight(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(2), NewString("days2"))
	q2 := NewQuantity(NewDecimalFloat64(26), NewString("hours2"))
	r, err := q1.Calc(q2, AdditionOp)
	e := NewQuantity(NewDecimalFloat64(1178), NewString("hours2"))
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

func TestQuantityCompareEqualUnitFactor(t *testing.T) {
	res, status := NewQuantity(NewDecimalFloat64(1.0), NewString("hour")).
		Compare(NewQuantity(NewDecimalFloat64(60.0), NewString("minutes")))
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

func TestQuantityCompareLessThanUnitFactor(t *testing.T) {
	res, status := NewQuantity(NewDecimalFloat64(1), NewString("week")).
		Compare(NewQuantity(NewDecimalFloat64(8), NewString("day")))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, -1, res)
}

func TestQuantityCompareLessThanUnitFactorExpLeft(t *testing.T) {
	res, status := NewQuantity(NewDecimalFloat64(1), NewString("week2")).
		Compare(NewQuantity(NewDecimalFloat64(50), NewString("day2")))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, -1, res)
}

func TestQuantityCompareLessThanUnitFactorExpRight(t *testing.T) {
	res, status := NewQuantity(NewDecimalFloat64(48), NewString("day2")).
		Compare(NewQuantity(NewDecimalFloat64(1), NewString("week2")))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, -1, res)
}

func TestQuantityCompareGreaterThan(t *testing.T) {
	res, status := NewQuantity(NewDecimalFloat64(10.21), NewString("cm")).
		Compare(NewQuantity(NewDecimalFloat64(10.11), NewString("cm")))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, 1, res)
}

func TestQuantityCompareGreaterThanUnitFactor(t *testing.T) {
	res, status := NewQuantity(NewDecimalFloat64(1), NewString("week")).
		Compare(NewQuantity(NewDecimalFloat64(6), NewString("days")))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, 1, res)
}

func TestQuantityCompareGreaterThanUnitFactorEquivalent(t *testing.T) {
	res, status := NewQuantity(NewDecimalFloat64(2), NewString("d")).
		Compare(NewQuantity(NewDecimalFloat64(1), NewString("days")))
	assert.Equal(t, Empty, status)
	assert.Equal(t, -1, res)
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

func TestQuantityToUnitToNil(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(10.64), NewString("cm"))
	assert.Nil(t, q.ToUnit(nil))
}

func TestQuantityToUnitNil(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(10.64), nil)
	assert.Nil(t, q.ToUnit(NewString("cm")))
}

func TestQuantityToUnitExpDiffer(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(10.64), NewString("years2"))
	assert.Nil(t, q.ToUnit(NewString("days3")))
}

func TestQuantityToUnitNoCommonBase(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(10.64), NewString("years"))
	assert.Nil(t, q.ToUnit(NewString("d")))
}

func TestQuantityToUnitEqual(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(10.64), NewString("cm"))
	assert.Same(t, q, q.ToUnit(NewString("cm")))
}

func TestQuantityToUnit(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(10.5), NewString("days"))
	q = q.ToUnit(NewString("week"))
	if assert.NotNil(t, q) {
		if assert.NotNil(t, q.Value()) {
			assert.Equal(t, 1.5, q.Value().Float64())
		}
		if assert.NotNil(t, q.Unit()) {
			assert.Equal(t, "weeks", q.Unit().String())
		}
	}
}

func TestQuantityToUnitExp(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(1), NewString("weeks2"))
	q = q.ToUnit(NewString("days2"))
	if assert.NotNil(t, q) {
		if assert.NotNil(t, q.Value()) {
			assert.Equal(t, 49.0, q.Value().Float64())
		}
		if assert.NotNil(t, q.Unit()) {
			assert.Equal(t, "days2", q.Unit().String())
		}
	}
}

func TestQuantityAbsPos(t *testing.T) {
	res := NewQuantity(NewDecimalFloat64(2.1), NewString("mg")).Abs()
	if assert.Implements(t, (*QuantityAccessor)(nil), res) {
		q := res.(QuantityAccessor)
		assert.Equal(t, 2.1, q.Value().Float64())
		if assert.NotNil(t, q.Unit()) {
			assert.Equal(t, "mg", q.Unit().String())
		}
	}
}

func TestQuantityAbsNeg(t *testing.T) {
	res := NewQuantity(NewDecimalFloat64(-2.1), NewString("mg")).Abs()
	if assert.Implements(t, (*QuantityAccessor)(nil), res) {
		q := res.(QuantityAccessor)
		assert.Equal(t, 2.1, q.Value().Float64())
		if assert.NotNil(t, q.Unit()) {
			assert.Equal(t, "mg", q.Unit().String())
		}
	}
}
