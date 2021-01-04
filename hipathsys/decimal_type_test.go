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

package hipathsys

import (
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecimalSource(t *testing.T) {
	o := NewDecimalIntWithSource(10, "abc")
	assert.Equal(t, "abc", o.Source())
}

func TestDecimalImplementsNegator(t *testing.T) {
	o := NewDecimalFloat64(4711.10)
	assert.Implements(t, (*Negator)(nil), o)
}

func TestDecimalDataType(t *testing.T) {
	o := NewDecimalFloat64(4711.10)
	dataType := o.DataType()
	assert.Equal(t, DecimalDataType, dataType)
}

func TestDecimalTypeSpec(t *testing.T) {
	o := NewDecimalInt(0)
	i := o.TypeSpec()
	if assert.NotNil(t, i, "type info expected") {
		assert.Equal(t, "System.Decimal", i.String())
		if assert.NotNil(t, i.FQBaseName(), "base name expected") {
			assert.Equal(t, "System.Any", i.FQBaseName().String())
		}
	}
}

func TestNewDecimal(t *testing.T) {
	o := NewDecimal(decimal.NewFromFloat(-4711.12))
	assert.Equal(t, -4711.12, o.Float64())
	assert.Equal(t, "-4711.12", o.String())
	assert.True(t, decimal.NewFromFloat(-4711.12).Equal(o.Primitive()))
}

func TestNewDecimalInt(t *testing.T) {
	o := NewDecimalInt(-4711)
	assert.Equal(t, -4711.0, o.Float64())
	assert.Equal(t, "-4711", o.String())
}

func TestDecimalInt(t *testing.T) {
	o := NewDecimalFloat64(-4711.831)
	assert.Equal(t, int32(-4711), o.Int())
}

func TestNewDecimalInt64(t *testing.T) {
	o := NewDecimalInt64(-4711)
	assert.Equal(t, -4711.0, o.Float64())
}

func TestDecimalInt64(t *testing.T) {
	o := NewDecimalFloat64(-4711.831)
	assert.Equal(t, int64(-4711), o.Int64())
}

func TestNewDecimalFloat32(t *testing.T) {
	o := NewDecimalFloat64(-4711.6)
	assert.Equal(t, float32(-4711.6), o.Float32())
	assert.Equal(t, "-4711.6", o.String())
}

func TestNewDecimalFloat64(t *testing.T) {
	o := NewDecimalFloat64(-4711.678121)
	assert.Equal(t, -4711.678121, o.Float64())
}

func TestNewDecimalValue(t *testing.T) {
	o := NewDecimalFloat64(-4711.678121)
	assert.Same(t, o, o.Value())
}

func TestDecimalBigFloat(t *testing.T) {
	o, err := ParseDecimal("-4711.83123200")
	assert.Nil(t, err, "no error expected")
	if assert.NotNil(t, o, "value expected") {
		assert.Equal(t, "-4711.831232", o.BigFloat().String())
	}
}

func TestDecimalDecimal(t *testing.T) {
	o, err := ParseDecimal("-4711.831232753400")
	assert.Nil(t, err, "no error expected")
	if assert.NotNil(t, o, "value expected") {
		assert.Equal(t, int32(-12), o.Decimal().Exponent())
		assert.Equal(t, "-4711.831232753400", o.String())
	}
}

func TestParseDecimal(t *testing.T) {
	o, err := ParseDecimal("-83628.85")
	assert.Nil(t, err, "no error expected")
	if assert.NotNil(t, o, "value expected") {
		assert.Equal(t, -83628.85, o.Float64())
	}
}

func TestParseDecimalInvalid(t *testing.T) {
	o, err := ParseDecimal("82737u83")
	assert.Nil(t, o, "value unexpected")
	assert.NotNil(t, err, "error expected")
}

func TestDecimalNegatePos(t *testing.T) {
	o := NewDecimalFloat64(8372.1)
	n := o.Negate()
	assert.NotSame(t, o, n)
	assert.Equal(t, 8372.1, o.Float64())
	if assert.Implements(t, (*DecimalAccessor)(nil), n) {
		assert.Equal(t, -8372.1, n.(DecimalAccessor).Float64())
	}
}

func TestDecimalNegateNeg(t *testing.T) {
	o := NewDecimalFloat64(-8372.1)
	n := o.Negate()
	assert.NotSame(t, o, n)
	assert.Equal(t, -8372.1, o.Float64())
	if assert.Implements(t, (*DecimalAccessor)(nil), n) {
		assert.Equal(t, 8372.1, n.(DecimalAccessor).Float64())
	}
}

func TestDecimalEqualNil(t *testing.T) {
	assert.Equal(t, false, NewDecimalInt(0).Equal(nil))
}

func TestDecimalEqualTypeDiffers(t *testing.T) {
	assert.Equal(t, false, NewDecimalInt(0).Equal(newAccessorMock()))
	assert.Equal(t, false, NewDecimalInt(0).Equivalent(newAccessorMock()))
}

func TestDecimalEqualEqual(t *testing.T) {
	assert.Equal(t, true, NewDecimalFloat64(8274.21).Equal(NewDecimalFloat64(8274.21)))
	assert.Equal(t, true, NewDecimalFloat64(8274.21).Equivalent(NewDecimalFloat64(8274.21)))
}

func TestDecimalEqualInteger(t *testing.T) {
	assert.Equal(t, true, NewDecimalInt(10).Equal(NewInteger(10)))
	assert.Equal(t, true, NewDecimalInt(10).Equivalent(NewInteger(10)))
}

func TestDecimalEqualIntegerNot(t *testing.T) {
	assert.Equal(t, false, NewDecimalInt(10).Equal(NewInteger(11)))
	assert.Equal(t, false, NewDecimalInt(10).Equivalent(NewInteger(11)))
}

func TestDecimalEqualNotEqual(t *testing.T) {
	assert.Equal(t, false, NewDecimalFloat64(8274.21).Equal(NewDecimalFloat64(8274.22)))
	assert.Equal(t, false, NewDecimalFloat64(8274.21).Equivalent(NewDecimalFloat64(8274.22)))
}

func TestDecimalEquivalentLeftPrecision(t *testing.T) {
	d := NewDecimalFloat64(64.12)
	assert.Equal(t, false, NewDecimalFloat64(64.1).Equal(d))
	assert.Equal(t, true, NewDecimalFloat64(64.1).Equivalent(d))
}

func TestDecimalEquivalentRightPrecision(t *testing.T) {
	d := NewDecimalFloat64(64.1)
	assert.Equal(t, false, NewDecimalFloat64(64.12).Equal(d))
	assert.Equal(t, true, NewDecimalFloat64(64.12).Equivalent(d))
}

func TestDecimalEqualQuantity(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(64.12), NewString("cm"))
	assert.Equal(t, true, NewDecimalFloat64(64.12).Equal(q))
	assert.Equal(t, true, NewDecimalFloat64(64.12).Equivalent(q))
}

func TestDecimalEquivalentQuantity(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(64.12), NewString("cm"))
	assert.Equal(t, false, NewDecimalFloat64(64.1).Equal(q))
	assert.Equal(t, true, NewDecimalFloat64(64.1).Equivalent(q))
}

func TestDecimalEqualNotEqualQuantity(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(64.14), NewString("cm"))
	assert.Equal(t, false, NewDecimalFloat64(64.12).Equal(q))
	assert.Equal(t, false, NewDecimalFloat64(64.12).Equivalent(q))
}

func TestDecimalEquivalentLeft(t *testing.T) {
	assert.Equal(t, true, NewDecimalFloat64(8274.6).Equivalent(NewDecimalFloat64(8274.67)))
}

func TestDecimalEquivalentRight(t *testing.T) {
	assert.Equal(t, true, NewDecimalFloat64(8274.67).Equivalent(NewDecimalFloat64(8274.6)))
}

func TestDecimalEquivalentInteger(t *testing.T) {
	assert.Equal(t, false, NewDecimalFloat64(8274.61).Equal(NewInteger(8274)))
	assert.Equal(t, true, NewDecimalFloat64(8274.61).Equivalent(NewInteger(8274)))
}

func TestDecimalWithValueNil(t *testing.T) {
	assert.Nil(t, NewDecimalFloat64(82763.22).WithValue(nil))
}

func TestDecimalWithValueDecimal(t *testing.T) {
	d := NewDecimalFloat64(232.123)
	assert.Same(t, d, NewDecimalFloat64(82763.22).WithValue(d))
}

func TestDecimalWithValueInteger(t *testing.T) {
	d := NewInteger(232)
	r := NewDecimalFloat64(82763.22).WithValue(d)
	if assert.Implements(t, (*DecimalAccessor)(nil), r) {
		assert.Equal(t, float64(232), r.(DecimalAccessor).Float64())
	}
}

func TestDecimalArithmeticOpSupported(t *testing.T) {
	d := NewDecimalInt(0)
	assert.True(t, d.ArithmeticOpSupported(AdditionOp), "addition must be supported")
	assert.True(t, d.ArithmeticOpSupported(SubtractionOp), "subtraction must be supported")
	assert.True(t, d.ArithmeticOpSupported(MultiplicationOp), "multiplication must be supported")
	assert.True(t, d.ArithmeticOpSupported(DivisionOp), "division must be supported")
	assert.True(t, d.ArithmeticOpSupported(DivOp), "DIV must be supported")
	assert.True(t, d.ArithmeticOpSupported(ModOp), "MOD must be supported")
}

func TestDecimalCalcNil(t *testing.T) {
	r, err := NewDecimalInt(122).Calc(nil, AdditionOp)
	assert.NoError(t, err)
	assert.Nil(t, r)
}

func TestDecimalCalcAddition(t *testing.T) {
	r, err := NewDecimalFloat64(122.23).Calc(NewDecimalFloat64(23.21), AdditionOp)
	e := NewDecimalFloat64(145.44)
	assert.NoError(t, err)
	if assert.NotNil(t, r.Value()) {
		assert.True(t, e.Decimal().Equal(r.Value().Decimal()),
			"expected %s, got %s", e.String(), r.Value().String())
	}
}

func TestDecimalCalcSubtraction(t *testing.T) {
	r, err := NewDecimalFloat64(150.90).Calc(NewDecimalFloat64(100.6), SubtractionOp)
	e := NewDecimalFloat64(50.3)
	assert.NoError(t, err)
	if assert.NotNil(t, r.Value()) {
		assert.True(t, e.Decimal().Equal(r.Value().Decimal()),
			"expected %s, got %s", e.String(), r.Value().String())
	}
}

func TestDecimalCalcMultiplication(t *testing.T) {
	r, err := NewDecimalFloat64(150.90).Calc(NewDecimalFloat64(1.5), MultiplicationOp)
	e := NewDecimalFloat64(226.35)
	assert.NoError(t, err)
	if assert.NotNil(t, r.Value()) {
		assert.True(t, e.Decimal().Equal(r.Value().Decimal()),
			"expected %s, got %s", e.String(), r.Value().String())
	}
}

func TestDecimalCalcDivision(t *testing.T) {
	r, err := NewDecimalFloat64(150.90).Calc(NewDecimalFloat64(1.5), DivisionOp)
	e := NewDecimalFloat64(100.6)
	assert.NoError(t, err)
	if assert.NotNil(t, r.Value()) {
		assert.True(t, e.Decimal().Equal(r.Value().Decimal()),
			"expected %s, got %s", e.String(), r.Value().String())
	}
}

func TestDecimalCalcDivisionByZero(t *testing.T) {
	r, err := NewDecimalFloat64(150.90).Calc(NewDecimalFloat64(0), DivisionOp)
	assert.NoError(t, err)
	assert.Nil(t, r, "division by zero")
}

func TestDecimalCalcDiv(t *testing.T) {
	r, err := NewDecimalFloat64(150.90).Calc(NewDecimalFloat64(1.5), DivOp)
	e := NewDecimalFloat64(100.0)
	assert.NoError(t, err)
	if assert.NotNil(t, r.Value()) {
		assert.True(t, e.Decimal().Equal(r.Value().Decimal()),
			"expected %s, got %s", e.String(), r.Value().String())
	}
}

func TestDecimalCalcDivByZero(t *testing.T) {
	r, err := NewDecimalFloat64(150.90).Calc(NewDecimalFloat64(0), DivOp)
	assert.NoError(t, err)
	assert.Nil(t, r, "division by zero")
}

func TestDecimalCalcMod(t *testing.T) {
	r, err := NewDecimalFloat64(150.90).Calc(NewDecimalFloat64(1.5), ModOp)
	e := NewDecimalFloat64(0.9)
	assert.NoError(t, err)
	if assert.NotNil(t, r.Value()) {
		assert.True(t, e.Decimal().Equal(r.Value().Decimal()),
			"expected %s, got %s", e.String(), r.Value().String())
	}
}

func TestDecimalCalcModByZero(t *testing.T) {
	r, err := NewDecimalFloat64(150.90).Calc(NewDecimalFloat64(0), ModOp)
	assert.NoError(t, err)
	assert.Nil(t, r, "division by zero")
}

func TestDecimalCalcUnknownOp(t *testing.T) {
	assert.Panics(t, func() { _, _ = NewDecimalFloat64(150.90).Calc(NewDecimalFloat64(1.5), 'x') })
}

func TestDecimalCalcRightNil(t *testing.T) {
	r, err := NewDecimalFloat64(1.5).Calc(nil, AdditionOp)
	assert.NoError(t, err)
	assert.Nil(t, r)
}

func TestDecimalCalcRightNilValue(t *testing.T) {
	r, err := NewDecimalFloat64(1.5).Calc(newDecimalValueAccessorMock(), AdditionOp)
	assert.NoError(t, err)
	assert.Nil(t, r)
}

func TestDecimalCalcQuantity(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(47.2), NewString("m"))
	r, err := NewDecimalFloat64(1.5).Calc(q, AdditionOp)
	e := NewQuantity(NewDecimalFloat64(48.7), NewString("m"))
	assert.NoError(t, err, "error expected")
	assert.True(t, e.Equal(r))
}

func TestDecimalCalcNotSupportedOp(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(47.2), NewString("m"))
	r, err := NewDecimalFloat64(1.5).Calc(q, ModOp)
	assert.Error(t, err, "error expected")
	assert.Nil(t, r)
}

type decimalValueAccessorMock struct {
}

func newDecimalValueAccessorMock() DecimalValueAccessor {
	return &decimalValueAccessorMock{}
}

func (d *decimalValueAccessorMock) Source() interface{} {
	panic("implement me")
}

func (d *decimalValueAccessorMock) Value() DecimalAccessor {
	return nil
}

func (d *decimalValueAccessorMock) WithValue(node NumberAccessor) DecimalValueAccessor {
	return node
}

func (d *decimalValueAccessorMock) ArithmeticOpSupported(ArithmeticOps) bool {
	return true
}

func (d *decimalValueAccessorMock) Equal(interface{}) bool {
	panic("implement me")
}

func (d *decimalValueAccessorMock) Equivalent(interface{}) bool {
	panic("implement me")
}

func (d *decimalValueAccessorMock) DataType() DataTypes {
	panic("implement me")
}

func (d *decimalValueAccessorMock) TypeSpec() TypeSpecAccessor {
	panic("implement me")
}

func (d *decimalValueAccessorMock) String() string {
	panic("implement me")
}

func TestDecimalValueFloat64Nil(t *testing.T) {
	assert.Nil(t, DecimalValueFloat64(nil))
}

func TestDecimalValueInt(t *testing.T) {
	assert.Equal(t, 209.123, DecimalValueFloat64(NewDecimalFloat64(209.123)))
}

func TestDecimalCompareEqual(t *testing.T) {
	res, status := NewDecimalFloat64(10.21).Compare(NewDecimalFloat64(10.21))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, 0, res)
}

func TestDecimalCompareEqualTypeDiffers(t *testing.T) {
	res, status := NewDecimalInt(10).Compare(NewString("test1"))
	assert.Equal(t, Inconvertible, status)
	assert.Equal(t, -1, res)
}

func TestDecimalCompareLessThan(t *testing.T) {
	res, status := NewDecimalFloat64(10.64).Compare(NewDecimalFloat64(10.71))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, -1, res)
}

func TestDecimalCompareGreaterThan(t *testing.T) {
	res, status := NewDecimalFloat64(10.81).Compare(NewDecimalFloat64(10.71))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, 1, res)
}

func TestDecimalCompareInteger(t *testing.T) {
	res, status := NewDecimalFloat64(10.64).Compare(NewInteger(11))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, -1, res)
}

func TestDecimalCompareQuantity(t *testing.T) {
	res, status := NewDecimalFloat64(10.64).Compare(
		NewQuantity(NewDecimalFloat64(10.71), NewString("cm")))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, -1, res)
}

func TestDecimalOfIntZero(t *testing.T) {
	res := DecimalOfInt(0)
	assert.Equal(t, 0.0, res.Float64())
}

func TestDecimalOfIntOne(t *testing.T) {
	res := DecimalOfInt(1)
	assert.Equal(t, 1.0, res.Float64())
}

func TestDecimalOfIntTwo(t *testing.T) {
	res := DecimalOfInt(2)
	assert.Equal(t, 2.0, res.Float64())
}

func TestDecimalOfIntThree(t *testing.T) {
	res := DecimalOfInt(3)
	assert.Equal(t, 3.0, res.Float64())
}

func TestDecimalOfIntFour(t *testing.T) {
	res := DecimalOfInt(4)
	assert.Equal(t, 4.0, res.Float64())
}

func TestDecimalAbsPos(t *testing.T) {
	res := NewDecimalFloat64(10.21).Abs()
	assert.Equal(t, 10.21, res.Value().Float64())
}

func TestDecimalAbsNeg(t *testing.T) {
	res := NewDecimalFloat64(-10.21).Abs()
	assert.Equal(t, 10.21, res.Value().Float64())
}

func TestDecimalIsOne(t *testing.T) {
	assert.True(t, NewDecimalInt(1).One())
}

func TestDecimalIsOneNot(t *testing.T) {
	assert.False(t, NewDecimalFloat64(1.01).One())
}

func TestDecimalHasFractionZero(t *testing.T) {
	assert.False(t, NewDecimalInt(0).HasFraction())
}

func TestDecimalHasFractionIntPos(t *testing.T) {
	assert.False(t, NewDecimalInt(2).HasFraction())
}

func TestDecimalHasFractionIntNeg(t *testing.T) {
	assert.False(t, NewDecimalInt(-2).HasFraction())
}

func TestDecimalHasFraction(t *testing.T) {
	assert.True(t, NewDecimalFloat64(2.00000000001).HasFraction())
}

func TestDecimalHasFractionDot(t *testing.T) {
	assert.True(t, NewDecimalFloat64(.5).HasFraction())
}

func TestDecimalHasFractionMinusDot(t *testing.T) {
	assert.True(t, NewDecimalFloat64(-.5).HasFraction())
}

func TestDecimalPos(t *testing.T) {
	assert.True(t, NewDecimalFloat64(.000000000001).Positive())
}

func TestDecimalPosNot(t *testing.T) {
	assert.False(t, NewDecimalFloat64(0).Positive())
}

func TestDecimalCeiling(t *testing.T) {
	assert.Equal(t, 8.0, NewDecimalFloat64(7.1).Ceiling().Float64())
}

func TestDecimalExp(t *testing.T) {
	assert.InDelta(t, 54.59815, NewDecimalFloat64(4).Exp().Float64(), .000001)
}

func TestDecimalExpFraction(t *testing.T) {
	assert.InDelta(t, 90.01713, NewDecimalFloat64(4.5).Exp().Float64(), .000002)
}

func TestDecimalFloor(t *testing.T) {
	assert.Equal(t, 7.0, NewDecimalFloat64(7.9).Floor().Float64())
}

func TestDecimalLn(t *testing.T) {
	res, err := NewDecimalFloat64(7.5).Ln()
	assert.NoError(t, err, "no error expected")
	if assert.NotNil(t, res) {
		assert.InDelta(t, 2.0149, res.Float64(), .00001)
	}
}

func TestDecimalLnZero(t *testing.T) {
	res, err := NewDecimalInt(0).Ln()
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestDecimalLog(t *testing.T) {
	res, err := NewDecimalFloat64(7.5).Log(NewDecimalFloat64(2.2))
	assert.NoError(t, err, "no error expected")
	if assert.NotNil(t, res) {
		assert.InDelta(t, 2.5555, res.Float64(), .00001)
	}
}

func TestDecimalLogZero(t *testing.T) {
	res, err := NewDecimalFloat64(0).Log(NewInteger(2))
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestDecimalLogBaseZero(t *testing.T) {
	res, err := NewDecimalFloat64(7).Log(NewInteger(0))
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestDecimalPower(t *testing.T) {
	res, ok := NewDecimalFloat64(20.5).Power(NewDecimalFloat64(3.5))
	assert.True(t, ok)
	assert.InDelta(t, 39006.6374, res.Float64(), 0.00005)
}

func TestDecimalPowerExp1(t *testing.T) {
	res, ok := NewDecimalFloat64(20.5).Power(NewDecimalInt(1))
	assert.True(t, ok)
	assert.Equal(t, 20.5, res.Float64())
}

func TestDecimalPowerNaN(t *testing.T) {
	res, ok := NewDecimalInt(-1).Power(NewDecimalFloat64(0.5))
	assert.False(t, ok)
	assert.Nil(t, res, "NaN expected")
}

func TestDecimalRound(t *testing.T) {
	res, err := NewDecimalFloat64(7.82481).Round(3)
	assert.NoError(t, err, "no error expected")
	if assert.NotNil(t, res) {
		assert.Equal(t, 7.825, res.Float64())
	}
}

func TestDecimalRoundPrecisionZero(t *testing.T) {
	res, err := NewDecimalFloat64(7.812).Round(0)
	assert.NoError(t, err, "no error expected")
	if assert.NotNil(t, res) {
		assert.Equal(t, 8.0, res.Float64())
	}
}

func TestDecimalRoundPrecisionNeg(t *testing.T) {
	res, err := NewDecimalFloat64(7).Round(-1)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestDecimalSqrt(t *testing.T) {
	res, ok := NewDecimalFloat64(16.6).Sqrt()
	assert.True(t, ok)
	if assert.NotNil(t, res) {
		assert.InDelta(t, 4.0743, res.Float64(), .00001)
	}
}

func TestDecimalSqrtNan(t *testing.T) {
	res, ok := NewDecimalFloat64(-1).Sqrt()
	assert.False(t, ok)
	assert.Nil(t, res, "NaN expected")
}

func TestDecimalSqrtNanFraction(t *testing.T) {
	res, ok := NewDecimalFloat64(-.5).Sqrt()
	assert.False(t, ok)
	assert.Nil(t, res, "NaN expected")
}

func TestDecimalTruncate(t *testing.T) {
	v := NewDecimalFloat64(23223.187636)
	r, _ := v.Truncate(2).Decimal().Float64()
	assert.Equal(t, 23223.18, r)
}
