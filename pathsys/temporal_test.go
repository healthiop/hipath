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
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAddQuantityTemporalDurationYear(t *testing.T) {
	v := time.Date(2019, 7, 14, 18, 44, 21, 982123654, time.UTC)
	res, err := addQuantityTemporalDuration(newDateTemporalAccessorMock(v, NanoTimePrecision),
		NewDecimalFloat64(4.89), YearDatePrecision)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, time.Date(2023, 7, 14, 18, 44, 21, 982123654, time.UTC).UnixNano(), res.UnixNano())
}

func TestAddQuantityTemporalDurationYearNeg(t *testing.T) {
	v := time.Date(2019, 7, 14, 18, 44, 21, 982123654, time.UTC)
	res, err := addQuantityTemporalDuration(newDateTemporalAccessorMock(v, NanoTimePrecision),
		NewDecimalInt(-2), YearDatePrecision)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, time.Date(2017, 7, 14, 18, 44, 21, 982123654, time.UTC).UnixNano(), res.UnixNano())
}

func TestAddQuantityTemporalDurationMonth(t *testing.T) {
	v := time.Date(2019, 7, 14, 18, 44, 21, 982123654, time.Local)
	res, err := addQuantityTemporalDuration(newDateTemporalAccessorMock(v, NanoTimePrecision),
		NewDecimalFloat64(7.89), MonthDatePrecision)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, time.Date(2020, 2, 14, 18, 44, 21, 982123654, time.Local).UnixNano(), res.UnixNano())
}

func TestAddQuantityTemporalDurationDay(t *testing.T) {
	v := time.Date(2019, 7, 14, 18, 44, 21, 982123654, time.UTC)
	res, err := addQuantityTemporalDuration(newDateTemporalAccessorMock(v, NanoTimePrecision),
		NewDecimalFloat64(19.89), DayDatePrecision)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, time.Date(2019, 8, 2, 18, 44, 21, 982123654, time.UTC).UnixNano(), res.UnixNano())
}

func TestAddQuantityTemporalDurationHour(t *testing.T) {
	v := time.Date(2019, 7, 14, 18, 44, 21, 982123654, time.UTC)
	res, err := addQuantityTemporalDuration(newDateTemporalAccessorMock(v, NanoTimePrecision),
		NewDecimalFloat64(26.89), HourTimePrecision)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, time.Date(2019, 7, 15, 20, 44, 21, 982123654, time.UTC).UnixNano(), res.UnixNano())
}

func TestAddQuantityTemporalDurationMinute(t *testing.T) {
	v := time.Date(2019, 7, 14, 18, 44, 21, 982123654, time.UTC)
	res, err := addQuantityTemporalDuration(newDateTemporalAccessorMock(v, NanoTimePrecision),
		NewDecimalFloat64(56.89), MinuteTimePrecision)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, time.Date(2019, 7, 14, 19, 40, 21, 982123654, time.UTC).UnixNano(), res.UnixNano())
}

func TestAddQuantityTemporalDurationSecond(t *testing.T) {
	v := time.Date(2019, 7, 14, 18, 44, 21, 982123654, time.UTC)
	res, err := addQuantityTemporalDuration(newDateTemporalAccessorMock(v, NanoTimePrecision),
		NewDecimalFloat64(59.0), SecondTimePrecision)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, time.Date(2019, 7, 14, 18, 45, 20, 982123654, time.UTC).UnixNano(), res.UnixNano())
}

func TestAddQuantityTemporalDurationSecondFraction(t *testing.T) {
	v := time.Date(2019, 7, 14, 18, 44, 21, 982123654, time.UTC)
	res, err := addQuantityTemporalDuration(newDateTemporalAccessorMock(v, NanoTimePrecision),
		NewDecimalFloat64(58.12), SecondTimePrecision)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, time.Date(2019, 7, 14, 18, 45, 20, 102123654, time.UTC).UnixNano(), res.UnixNano())
}

func TestAddQuantityTemporalDurationNanosecond(t *testing.T) {
	v := time.Date(2019, 7, 14, 18, 44, 21, 982123654, time.UTC)
	res, err := addQuantityTemporalDuration(newDateTemporalAccessorMock(v, NanoTimePrecision),
		NewDecimalFloat64(120000000.0), NanoTimePrecision)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, time.Date(2019, 7, 14, 18, 44, 22, 102123654, time.UTC).UnixNano(), res.UnixNano())
}

func TestAddQuantityDateTimePrecisionInvalid(t *testing.T) {
	v := time.Date(2019, 7, 14, 18, 44, 21, 982123654, time.UTC)
	assert.Panics(t, func() {
		_, _ = addQuantityTemporalDuration(newDateTemporalAccessorMock(v, NanoTimePrecision),
			NewDecimalFloat64(120000000.0), -1)
	})
}

func TestAddQuantityDateTimePrecisionYearTooSmall(t *testing.T) {
	v := time.Date(2019, 7, 14, 18, 44, 21, 982123654, time.UTC)
	_, err := addQuantityTemporalDuration(newDateTemporalAccessorMock(v, NanoTimePrecision),
		NewDecimalInt(-2020), YearDatePrecision)
	assert.Error(t, err, "error expected")
}

func TestAddQuantityDateTimePrecisionYearTooBig(t *testing.T) {
	v := time.Date(2019, 7, 14, 18, 44, 21, 982123654, time.UTC)
	_, err := addQuantityTemporalDuration(newDateTemporalAccessorMock(v, NanoTimePrecision),
		NewDecimalInt(7981), YearDatePrecision)
	assert.Error(t, err, "error expected")
}

func TestAddQuantityTemporalDurationNanosecondPrecisionSecond(t *testing.T) {
	v := time.Date(2019, 7, 14, 18, 44, 21, 0, time.UTC)
	res, err := addQuantityTemporalDuration(newDateTemporalAccessorMock(v, SecondTimePrecision),
		NewDecimalFloat64(2_800_000_000), NanoTimePrecision)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, time.Date(2019, 7, 14, 18, 44, 23, 0, time.UTC).UnixNano(), res.UnixNano())
}

func TestAddQuantityTemporalDurationSecondPrecisionMinute(t *testing.T) {
	v := time.Date(2019, 7, 14, 18, 44, 0, 0, time.UTC)
	res, err := addQuantityTemporalDuration(newDateTemporalAccessorMock(v, MinuteTimePrecision),
		NewDecimalFloat64(140), SecondTimePrecision)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, time.Date(2019, 7, 14, 18, 46, 0, 0, time.UTC).UnixNano(), res.UnixNano())
}

func TestAddQuantityTemporalDurationMinutePrecisionHour(t *testing.T) {
	v := time.Date(2019, 7, 14, 18, 0, 0, 0, time.UTC)
	res, err := addQuantityTemporalDuration(newDateTemporalAccessorMock(v, HourTimePrecision),
		NewDecimalFloat64(140), MinuteTimePrecision)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, time.Date(2019, 7, 14, 20, 0, 0, 0, time.UTC).UnixNano(), res.UnixNano())
}

func TestAddQuantityTemporalDurationHourPrecisionDay(t *testing.T) {
	v := time.Date(2019, 7, 14, 0, 0, 0, 0, time.UTC)
	res, err := addQuantityTemporalDuration(newDateTemporalAccessorMock(v, DayDatePrecision),
		NewDecimalFloat64(50), HourTimePrecision)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, time.Date(2019, 7, 16, 0, 0, 0, 0, time.UTC).UnixNano(), res.UnixNano())
}

func TestAddQuantityTemporalDurationDayPrecisionMonth(t *testing.T) {
	v := time.Date(2019, 7, 1, 0, 0, 0, 0, time.UTC)
	res, err := addQuantityTemporalDuration(newDateTemporalAccessorMock(v, MonthDatePrecision),
		NewDecimalFloat64(60), DayDatePrecision)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, time.Date(2019, 9, 1, 0, 0, 0, 0, time.UTC).UnixNano(), res.UnixNano())
}

func TestAddQuantityTemporalDurationDayPrecisionOneMonth(t *testing.T) {
	v := time.Date(2019, 7, 1, 0, 0, 0, 0, time.UTC)
	res, err := addQuantityTemporalDuration(newDateTemporalAccessorMock(v, MonthDatePrecision),
		NewDecimalFloat64(59), DayDatePrecision)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, time.Date(2019, 8, 1, 0, 0, 0, 0, time.UTC).UnixNano(), res.UnixNano())
}

func TestAddQuantityTemporalDurationMonthPrecisionYear(t *testing.T) {
	v := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	res, err := addQuantityTemporalDuration(newDateTemporalAccessorMock(v, YearDatePrecision),
		NewDecimalFloat64(26), MonthDatePrecision)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano(), res.UnixNano())
}

func TestAddQuantityTemporalDurationInvalidPrecision(t *testing.T) {
	v := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	assert.Panics(t, func() {
		_, _ = addQuantityTemporalDuration(newDateTemporalAccessorMock(v, -1), NewDecimalFloat64(26), MonthDatePrecision)
	})
}

func TestQuantityValueNanosInvalidPrecision(t *testing.T) {
	assert.Panics(t, func() { _ = quantityValueNanos(NewDecimalInt(10), -1) })
}

func TestQuantityDateTimePrecisionUnitNil(t *testing.T) {
	q := NewQuantity(NewDecimalInt(10), nil)
	_, _, err := quantityDateTimePrecision(q)
	assert.Error(t, err, "error expected")
}

func TestQuantityDateTimePrecisionYear(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(2.5), NewString("year"))
	v, p, err := quantityDateTimePrecision(q)
	assert.NoError(t, err, "error expected")
	if assert.NotNil(t, v) {
		assert.Equal(t, 2.5, v.Float64())
	}
	assert.Equal(t, YearDatePrecision, p)
}

func TestQuantityDateTimePrecisionMonth(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(2.5), NewString("month"))
	v, p, err := quantityDateTimePrecision(q)
	assert.NoError(t, err, "error expected")
	if assert.NotNil(t, v) {
		assert.Equal(t, 2.5, v.Float64())
	}
	assert.Equal(t, MonthDatePrecision, p)
}

func TestQuantityDateTimePrecisionWeek(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(2.5), NewString("week"))
	v, p, err := quantityDateTimePrecision(q)
	assert.NoError(t, err, "error expected")
	if assert.NotNil(t, v) {
		assert.Equal(t, 17.5, v.Float64())
	}
	assert.Equal(t, DayDatePrecision, p)
}

func TestQuantityDateTimePrecisionDay(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(2.5), NewString("day"))
	v, p, err := quantityDateTimePrecision(q)
	assert.NoError(t, err, "error expected")
	if assert.NotNil(t, v) {
		assert.Equal(t, 2.5, v.Float64())
	}
	assert.Equal(t, DayDatePrecision, p)
}

func TestQuantityDateTimePrecisionHour(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(2.5), NewString("hour"))
	v, p, err := quantityDateTimePrecision(q)
	assert.NoError(t, err, "error expected")
	if assert.NotNil(t, v) {
		assert.Equal(t, 2.5, v.Float64())
	}
	assert.Equal(t, HourTimePrecision, p)
}

func TestQuantityDateTimePrecisionMinute(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(2.5), NewString("minute"))
	v, p, err := quantityDateTimePrecision(q)
	assert.NoError(t, err, "error expected")
	if assert.NotNil(t, v) {
		assert.Equal(t, 2.5, v.Float64())
	}
	assert.Equal(t, MinuteTimePrecision, p)
}

func TestQuantityDateTimePrecisionSecond(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(2.5), NewString("second"))
	v, p, err := quantityDateTimePrecision(q)
	assert.NoError(t, err, "error expected")
	if assert.NotNil(t, v) {
		assert.Equal(t, 2.5, v.Float64())
	}
	assert.Equal(t, SecondTimePrecision, p)
}

func TestQuantityDateTimePrecisionMillisecond(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(2.5), NewString("millisecond"))
	v, p, err := quantityDateTimePrecision(q)
	assert.NoError(t, err, "error expected")
	if assert.NotNil(t, v) {
		assert.Equal(t, 2_500_000.0, v.Float64())
	}
	assert.Equal(t, NanoTimePrecision, p)
}

func TestQuantityDateTimePrecisionInvalidUnit(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(2.5), NewString("x"))
	_, _, err := quantityDateTimePrecision(q)
	assert.Error(t, err, "error expected")
}

func TestTemporalPrecisionEqual(t *testing.T) {
	assert.True(t, TemporalPrecisionEqual(
		newDateTemporalAccessorMock(time.Now(), MinuteTimePrecision),
		newDateTemporalAccessorMock(time.Now(), MinuteTimePrecision)))
}

func TestTemporalPrecisionEqualNot(t *testing.T) {
	assert.False(t, TemporalPrecisionEqual(
		newDateTemporalAccessorMock(time.Now(), MinuteTimePrecision),
		newDateTemporalAccessorMock(time.Now(), HourTimePrecision)))
}

func TestTemporalPrecisionEqualSecondNano(t *testing.T) {
	assert.True(t, TemporalPrecisionEqual(
		newDateTemporalAccessorMock(time.Now(), SecondTimePrecision),
		newDateTemporalAccessorMock(time.Now(), NanoTimePrecision)))
}

func TestTemporalPrecisionEqualSecondOther(t *testing.T) {
	assert.False(t, TemporalPrecisionEqual(
		newDateTemporalAccessorMock(time.Now(), SecondTimePrecision),
		newDateTemporalAccessorMock(time.Now(), MinuteTimePrecision)))
}

func TestTemporalPrecisionEqualNanoOther(t *testing.T) {
	assert.False(t, TemporalPrecisionEqual(
		newDateTemporalAccessorMock(time.Now(), MinuteTimePrecision),
		newDateTemporalAccessorMock(time.Now(), NanoTimePrecision)))
}

type dateTemporalAccessorMock struct {
	time      time.Time
	precision DateTimePrecisions
}

func newDateTemporalAccessorMock(time time.Time, precision DateTimePrecisions) DateTemporalAccessor {
	return &dateTemporalAccessorMock{time, precision}
}

func (d *dateTemporalAccessorMock) Source() interface{} {
	panic("implement me")
}

func (d *dateTemporalAccessorMock) Time() time.Time {
	return d.time
}

func (d dateTemporalAccessorMock) Precision() DateTimePrecisions {
	return d.precision
}

func (d *dateTemporalAccessorMock) Year() int {
	panic("implement me")
}

func (d *dateTemporalAccessorMock) Month() int {
	panic("implement me")
}

func (d *dateTemporalAccessorMock) Day() int {
	panic("implement me")
}

func (d dateTemporalAccessorMock) DataType() DataTypes {
	panic("implement me")
}

func (d dateTemporalAccessorMock) TypeInfo() TypeInfoAccessor {
	panic("implement me")
}

func (d dateTemporalAccessorMock) Equal(interface{}) bool {
	panic("implement me")
}

func (d dateTemporalAccessorMock) Equivalent(interface{}) bool {
	panic("implement me")
}

func (d *dateTemporalAccessorMock) Compare(Comparator) (int, OperatorStatus) {
	panic("implement me")
}

func (d dateTemporalAccessorMock) String() string {
	panic("implement me")
}

func (d dateTemporalAccessorMock) LowestPrecision() DateTimePrecisions {
	panic("implement me")
}

func (d *dateTemporalAccessorMock) Add(QuantityAccessor) (TemporalAccessor, error) {
	panic("implement me")
}
