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
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegerDataType(t *testing.T) {
	o := NewInteger(4711)
	dataType := o.DataType()
	assert.Equal(t, IntegerDataType, dataType)
}

func TestIntegerTypeInfo(t *testing.T) {
	o := NewInteger(0)
	i := o.TypeInfo()
	if assert.NotNil(t, i, "type info expected") {
		assert.Equal(t, "System.Integer", i.String())
		if assert.NotNil(t, i.FQBaseName(), "base name expected") {
			assert.Equal(t, "System.Any", i.FQBaseName().String())
		}
	}
}

func TestIntegerNegatePos(t *testing.T) {
	o := NewInteger(8372)
	n := o.Negate()
	assert.NotSame(t, o, n)
	assert.Equal(t, float64(8372), o.Float64())
	if assert.Implements(t, (*IntegerAccessor)(nil), n) {
		assert.Equal(t, float64(-8372), n.(IntegerAccessor).Float64())
	}
}

func TestIntegerNegateNeg(t *testing.T) {
	o := NewInteger(-8372)
	n := o.Negate()
	assert.NotSame(t, o, n)
	assert.Equal(t, float64(-8372), o.Float64())
	if assert.Implements(t, (*IntegerAccessor)(nil), n) {
		assert.Equal(t, float64(8372), n.(IntegerAccessor).Float64())
	}
}

func TestIntegerValue(t *testing.T) {
	o := NewInteger(-4711)
	assert.Equal(t, int32(-4711), o.Int())
	assert.Equal(t, "-4711", o.String())
}

func TestIntegerValueDecimal(t *testing.T) {
	o := NewInteger(-4711)
	r := o.Value()
	assert.Same(t, r, o.Value())
}

func TestIntegerInt64(t *testing.T) {
	o := NewInteger(-4711)
	assert.Equal(t, int64(-4711), o.Int64())
}

func TestIntegerFloat64Value(t *testing.T) {
	o := NewInteger(-4711)
	value := o.Float64()
	assert.Equal(t, float64(-4711), value)
}

func TestIntegerDecimalValue(t *testing.T) {
	o := NewInteger(-4711)
	value := o.Decimal()
	expected := decimal.NewFromInt32(-4711)
	assert.True(t, expected.Equal(value), "expected %s, got %s", expected.String(), value.String())
}

func TestParseInteger(t *testing.T) {
	o, err := ParseInteger("-83628")
	assert.Nil(t, err, "no error expected")
	if assert.NotNil(t, o, "value expected") {
		assert.Equal(t, int32(-83628), o.Int())
	}
}

func TestParseIntegerInvalid(t *testing.T) {
	o, err := ParseInteger("8273.3")
	assert.Nil(t, o, "value unexpected")
	assert.NotNil(t, err, "error expected")
}

func TestIntegerWithValueNil(t *testing.T) {
	assert.Nil(t, NewDecimalFloat64(82763.22).WithValue(nil))
}

func TestIntegerWithValueInteger(t *testing.T) {
	d := NewInteger(232)
	assert.Same(t, d, NewInteger(82763).WithValue(d))
}

func TestIntegerWithValueDecimal(t *testing.T) {
	d := NewDecimalFloat64(232.72)
	r := NewInteger(123).WithValue(d)
	if assert.Implements(t, (*IntegerAccessor)(nil), r) {
		assert.Equal(t, float64(232), r.(IntegerAccessor).Float64())
	}
}

func TestIntegerEqualNil(t *testing.T) {
	assert.Equal(t, false, NewInteger(0).Equal(nil))
}

func TestIntegerEqualTypeDiffers(t *testing.T) {
	assert.Equal(t, false, NewInteger(0).Equal(newAccessorMock()))
	assert.Equal(t, false, NewInteger(0).Equivalent(newAccessorMock()))
}

func TestIntegerEqualRightNil(t *testing.T) {
	assert.Equal(t, false, NewInteger(0).Equal(nil))
	assert.Equal(t, false, NewInteger(0).Equivalent(nil))
}

func TestIntegerEqualEqual(t *testing.T) {
	assert.Equal(t, true, NewInteger(8274).Equal(NewInteger(8274)))
	assert.Equal(t, true, NewInteger(8274).Equivalent(NewInteger(8274)))
}

func TestIntegerEqualDecimal(t *testing.T) {
	assert.Equal(t, true, NewInteger(10).Equal(NewDecimalInt(10)))
	assert.Equal(t, true, NewInteger(10).Equivalent(NewDecimalInt(10)))
}

func TestIntegerEquivalentDecimal(t *testing.T) {
	assert.Equal(t, false, NewInteger(10).Equal(NewDecimalFloat64(10.1)))
	assert.Equal(t, true, NewInteger(10).Equivalent(NewDecimalFloat64(10.1)))
}

func TestIntegerEqualDecimalNot(t *testing.T) {
	assert.Equal(t, false, NewInteger(10).Equal(NewDecimalInt(11)))
	assert.Equal(t, false, NewInteger(10).Equivalent(NewDecimalInt(11)))
}

func TestIntegerEqualEqualDecimal(t *testing.T) {
	assert.Equal(t, true, NewInteger(8274).Equal(NewDecimalInt(8274)))
	assert.Equal(t, true, NewInteger(8274).Equivalent(NewDecimalInt(8274)))
}

func TestIntegerEqualNotEqual(t *testing.T) {
	assert.Equal(t, false, NewInteger(8274).Equal(NewInteger(8275)))
	assert.Equal(t, false, NewInteger(8274).Equivalent(NewInteger(8275)))
}

func TestIntegerEqualNotEqualDecimal(t *testing.T) {
	assert.Equal(t, false, NewInteger(8274).Equal(NewDecimalInt(8275)))
	assert.Equal(t, false, NewInteger(8274).Equivalent(NewDecimalInt(8275)))
}

func TestIntegerEquivalent(t *testing.T) {
	assert.Equal(t, false, NewInteger(8274).Equal(NewDecimalFloat64(8274.8237)))
	assert.Equal(t, true, NewInteger(8274).Equivalent(NewDecimalFloat64(8274.8237)))
}

func TestIntegerEqualQuantity(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(64), NewString("cm"))
	assert.Equal(t, true, NewInteger(64).Equal(q))
	assert.Equal(t, true, NewInteger(64).Equivalent(q))
}

func TestIntegerEquivalentQuantity(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(64.1), NewString("cm"))
	assert.Equal(t, false, NewInteger(64).Equal(q))
	assert.Equal(t, true, NewInteger(64).Equivalent(q))
}

func TestIntegerEqualNotEqualQuantity(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(65), NewString("cm"))
	assert.Equal(t, false, NewInteger(64).Equal(q))
	assert.Equal(t, false, NewInteger(64).Equivalent(q))
}

func TestIntegerArithmeticOpSupported(t *testing.T) {
	d := NewInteger(0)
	assert.True(t, d.ArithmeticOpSupported(AdditionOp), "addition must be supported")
	assert.True(t, d.ArithmeticOpSupported(SubtractionOp), "subtraction must be supported")
	assert.True(t, d.ArithmeticOpSupported(MultiplicationOp), "multiplication must be supported")
	assert.True(t, d.ArithmeticOpSupported(DivisionOp), "division must be supported")
	assert.True(t, d.ArithmeticOpSupported(DivOp), "DIV must be supported")
	assert.True(t, d.ArithmeticOpSupported(ModOp), "MOD must be supported")
}

func TestIntegerCalcNil(t *testing.T) {
	r, err := NewInteger(122).Calc(nil, AdditionOp)
	assert.NoError(t, err)
	assert.Nil(t, r)
}

func TestIntegerCalcAddition(t *testing.T) {
	r, err := NewInteger(122).Calc(NewInteger(23), AdditionOp)
	e := NewInteger(145)
	assert.NoError(t, err)
	assert.Implements(t, (*IntegerAccessor)(nil), r)
	if assert.NotNil(t, r.Value()) {
		assert.True(t, e.Decimal().Equal(r.Value().Decimal()),
			"expected %s, got %s", e.String(), r.Value().String())
	}
}

func TestIntegerCalcAdditionDecimal(t *testing.T) {
	r, err := NewInteger(122).Calc(NewDecimalFloat64(23.15), AdditionOp)
	e := NewDecimalFloat64(145.15)
	assert.NoError(t, err)
	assert.Implements(t, (*IntegerAccessor)(nil), r)
	if assert.NotNil(t, r.Value()) {
		assert.True(t, e.Decimal().Equal(r.Value().Decimal()),
			"expected %s, got %s", e.String(), r.Value().String())
	}
}

func TestIntegerCalcSubtraction(t *testing.T) {
	r, err := NewInteger(150).Calc(NewInteger(101), SubtractionOp)
	e := NewInteger(49)
	assert.NoError(t, err)
	assert.Implements(t, (*IntegerAccessor)(nil), r)
	if assert.NotNil(t, r.Value()) {
		assert.True(t, e.Decimal().Equal(r.Value().Decimal()),
			"expected %s, got %s", e.String(), r.Value().String())
	}
}

func TestIntegerCalcMultiplication(t *testing.T) {
	r, err := NewInteger(150).Calc(NewInteger(3), MultiplicationOp)
	e := NewInteger(450)
	assert.NoError(t, err)
	assert.Implements(t, (*IntegerAccessor)(nil), r)
	if assert.NotNil(t, r.Value()) {
		assert.True(t, e.Decimal().Equal(r.Value().Decimal()),
			"expected %s, got %s", e.String(), r.Value().String())
	}
}

func TestIntegerCalcDivision(t *testing.T) {
	r, err := NewInteger(150).Calc(NewInteger(8), DivisionOp)
	e := NewDecimalFloat64(18.75)
	assert.NoError(t, err)
	assert.Implements(t, (*DecimalAccessor)(nil), r)
	if assert.NotNil(t, r.Value()) {
		assert.True(t, e.Decimal().Equal(r.Value().Decimal()),
			"expected %s, got %s", e.String(), r.Value().String())
	}
}

func TestIntegerCalcDivisionByZero(t *testing.T) {
	r, err := NewInteger(150).Calc(NewInteger(0), DivisionOp)
	assert.NoError(t, err)
	assert.Nil(t, r, "division by zero")
}

func TestIntegerCalcDiv(t *testing.T) {
	r, err := NewInteger(150).Calc(NewInteger(8), DivOp)
	e := NewDecimalInt(18)
	assert.NoError(t, err)
	assert.Implements(t, (*IntegerAccessor)(nil), r)
	if assert.NotNil(t, r.Value()) {
		assert.True(t, e.Decimal().Equal(r.Value().Decimal()),
			"expected %s, got %s", e.String(), r.Value().String())
	}
}

func TestIntegerCalcDivByZero(t *testing.T) {
	r, err := NewInteger(150).Calc(NewInteger(0), DivOp)
	assert.NoError(t, err)
	assert.Nil(t, r, "division by zero")
}

func TestIntegerCalcMod(t *testing.T) {
	r, err := NewInteger(150).Calc(NewInteger(8), ModOp)
	e := NewDecimalFloat64(6)
	assert.NoError(t, err)
	assert.Implements(t, (*IntegerAccessor)(nil), r)
	if assert.NotNil(t, r.Value()) {
		assert.True(t, e.Decimal().Equal(r.Value().Decimal()),
			"expected %s, got %s", e.String(), r.Value().String())
	}
}

func TestIntegerCalcModByZero(t *testing.T) {
	r, err := NewInteger(150).Calc(NewInteger(0), ModOp)
	assert.NoError(t, err)
	assert.Nil(t, r, "division by zero")
}

func TestIntegerCalcUnknownOp(t *testing.T) {
	assert.Panics(t, func() { _, _ = NewInteger(150).Calc(NewInteger(1), 'x') })
}

func TestIntegerCalcRightNil(t *testing.T) {
	r, err := NewInteger(1).Calc(nil, AdditionOp)
	assert.NoError(t, err)
	assert.Nil(t, r)
}

func TestIntegerCalcQuantity(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(47.2), NewString("m"))
	r, err := NewInteger(2).Calc(q, AdditionOp)
	e := NewQuantity(NewDecimalFloat64(49.2), NewString("m"))
	assert.NoError(t, err, "error expected")
	assert.True(t, e.Equal(r))
}

func TestIntegerCalcNotSupportedOp(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(47.2), NewString("m"))
	r, err := NewInteger(2).Calc(q, ModOp)
	assert.Error(t, err, "error expected")
	assert.Nil(t, r)
}

func TestIntegerTruncate(t *testing.T) {
	v := NewInteger(23223)
	assert.Same(t, v, v.Truncate(2))
}

func TestIntegerValueNil(t *testing.T) {
	assert.Nil(t, IntegerValue(nil))
}

func TestIntegerValueInt(t *testing.T) {
	assert.Equal(t, int32(20), IntegerValue(NewInteger(20)))
}

func TestIntegerCompareEqual(t *testing.T) {
	res, status := NewInteger(10).Compare(NewInteger(10))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, 0, res)
}

func TestIntegerCompareEqualTypeDiffers(t *testing.T) {
	res, status := NewDecimalInt(10).Compare(NewString("test1"))
	assert.Equal(t, Inconvertible, status)
	assert.Equal(t, -1, res)
}

func TestIntegerCompareLessThan(t *testing.T) {
	res, status := NewInteger(10).Compare(NewInteger(11))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, -1, res)
}

func TestIntegerCompareGreaterThan(t *testing.T) {
	res, status := NewInteger(10).Compare(NewInteger(9))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, 1, res)
}

func TestIntegerCompareDecimal(t *testing.T) {
	res, status := NewInteger(10).Compare(NewDecimalFloat64(10.61))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, -1, res)
}

func TestIntegerCompareQuantity(t *testing.T) {
	res, status := NewInteger(10).Compare(
		NewQuantity(NewDecimalFloat64(10.71), NewString("cm")))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, -1, res)
}
