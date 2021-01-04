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
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDateTimeSource(t *testing.T) {
	o := NewDateTimeWithSource(time.Now(), "abc")
	assert.Equal(t, "abc", o.Source())
}

func TestDateTimeDataType(t *testing.T) {
	o := NewDateTime(time.Now())
	dataType := o.DataType()
	assert.Equal(t, DateTimeDataType, dataType)
}

func TestDateTimeTypeLowestPrecision(t *testing.T) {
	o := NewDateTime(time.Now())
	assert.Equal(t, YearDatePrecision, o.LowestPrecision())
}

func TestDateTimeTypeSpec(t *testing.T) {
	o := NewDateTime(time.Now())
	i := o.TypeSpec()
	if assert.NotNil(t, i, "type info expected") {
		assert.Equal(t, "System.DateTime", i.String())
		if assert.NotNil(t, i.FQBaseName(), "base name expected") {
			assert.Equal(t, "System.Any", i.FQBaseName().String())
		}
	}
}

func TestDateTimeValue(t *testing.T) {
	testTime := time.Date(2018, 5, 20, 17, 48, 14, 123, time.Local)
	o := NewDateTime(testTime)
	value := o.Time()
	assert.Equal(t, testTime.Year(), o.Year())
	assert.Equal(t, int(testTime.Month()), o.Month())
	assert.Equal(t, testTime.Day(), o.Day())
	assert.Equal(t, testTime.Hour(), o.Hour())
	assert.Equal(t, testTime.Minute(), o.Minute())
	assert.Equal(t, testTime.Second(), o.Second())
	assert.Equal(t, testTime.Nanosecond(), o.Nanosecond())
	assert.Same(t, testTime.Location(), o.Location())
	assert.Equal(t, NanoTimePrecision, o.Precision())
	assert.True(t, testTime.Equal(value), "expected %d, got %d",
		testTime.UnixNano(), value.UnixNano())
}

func TestDateTimeDate(t *testing.T) {
	testTime := time.Date(2018, 5, 20, 17, 48, 14, 123, time.Local)
	d := NewDateTime(testTime).Date()
	assert.Equal(t, testTime.Year(), d.Year())
	assert.Equal(t, int(testTime.Month()), d.Month())
	assert.Equal(t, testTime.Day(), d.Day())
	assert.Equal(t, DayDatePrecision, d.Precision())
}

func TestDateTimeDatePrecision(t *testing.T) {
	d := NewDateTimeYMDHMSNWithPrecision(2018, 5, 20, 17, 48, 14, 123, time.Local, MonthDatePrecision).Date()
	assert.Equal(t, 2018, d.Year())
	assert.Equal(t, 5, d.Month())
	assert.Equal(t, 1, d.Day())
	assert.Equal(t, MonthDatePrecision, d.Precision())
}

func TestDateTimeDateTime(t *testing.T) {
	testTime := time.Date(2018, 5, 20, 17, 48, 14, 123, time.Local)
	d := NewDateTime(testTime)
	assert.Same(t, d, d.DateTime())
}

func TestParseDateTimeCompleteTzPos(t *testing.T) {
	dt, err := ParseDateTime("2015-02-07T13:28:17.239+02:00")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		value := time.Date(2015, 2, 7, 13, 28, 17, 239000000,
			time.FixedZone("+02:00", 2*60*60))
		assert.True(t, value.Equal(dt.Time()), "expected %d, got %d",
			value.UnixNano(), dt.Time().UnixNano())
		assert.Equal(t, NanoTimePrecision, dt.Precision())
		assert.Equal(t, "2015-02-07T13:28:17.239000000+02:00", dt.String())
	}
}

func TestParseDateTimeCompleteTzNeg(t *testing.T) {
	dt, err := ParseDateTime("2015-02-07T13:28:17.239-05:30")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		value := time.Date(2015, 2, 7, 13, 28, 17, 239000000,
			time.FixedZone("-05:30", -19800))
		assert.True(t, value.Equal(dt.Time()), "expected %d, got %d",
			value.UnixNano(), dt.Time().UnixNano())
		assert.Equal(t, NanoTimePrecision, dt.Precision())
		assert.Equal(t, "2015-02-07T13:28:17.239000000-05:30", dt.String())
	}
}

func TestParseDateTimeCompleteTzZero(t *testing.T) {
	dt, err := ParseDateTime("2015-02-07T13:28:17.239+00:00")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		value := time.Date(2015, 2, 7, 13, 28, 17, 239000000,
			time.UTC)
		assert.True(t, value.Equal(dt.Time()), "expected %d, got %d",
			value.UnixNano(), dt.Time().UnixNano())
		assert.Equal(t, NanoTimePrecision, dt.Precision())
		assert.Equal(t, "2015-02-07T13:28:17.239000000+00:00", dt.String())
	}
}

func TestParseDateTimeCompleteTzUtc(t *testing.T) {
	dt, err := ParseDateTime("2015-02-07T13:28:17.239Z")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		value := time.Date(2015, 2, 7, 13, 28, 17, 239000000, time.UTC)
		assert.True(t, value.Equal(dt.Time()), "expected %d, got %d",
			value.UnixNano(), dt.Time().UnixNano())
		assert.Equal(t, NanoTimePrecision, dt.Precision())
	}
}

func TestParseDateTimeNoTz(t *testing.T) {
	dt, err := ParseDateTime("2015-02-07T13:28:17.239")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		value := time.Date(2015, 2, 7, 13, 28, 17, 239000000, time.Local)
		assert.True(t, value.Equal(dt.Time()), "expected %d, got %d",
			value.UnixNano(), dt.Time().UnixNano())
		assert.Equal(t, NanoTimePrecision, dt.Precision())
	}
}

func TestParseDateTimeFractionDigits(t *testing.T) {
	dt, err := ParseDateTime("2015-02-07T13:28:17.2397381239Z")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		value := time.Date(2015, 2, 7, 13, 28, 17, 239738123, time.UTC)
		assert.True(t, value.Equal(dt.Time()), "expected %d, got %d",
			value.UnixNano(), dt.Time().UnixNano())
		assert.Equal(t, NanoTimePrecision, dt.Precision())
		assert.Equal(t, "2015-02-07T13:28:17.239738123+00:00", dt.String())
	}
}

func TestParseDateTimeNoNanos(t *testing.T) {
	dt, err := ParseDateTime("2015-02-07T13:28:17Z")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		value := time.Date(2015, 2, 7, 13, 28, 17, 0, time.UTC)
		assert.True(t, value.Equal(dt.Time()), "expected %d, got %d",
			value.UnixNano(), dt.Time().UnixNano())
		assert.Equal(t, SecondTimePrecision, dt.Precision())
		assert.Equal(t, "2015-02-07T13:28:17+00:00", dt.String())
	}
}

func TestParseDateTimeNoSeconds(t *testing.T) {
	dt, err := ParseDateTime("2015-02-07T13:28Z")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		value := time.Date(2015, 2, 7, 13, 28, 0, 0, time.UTC)
		assert.True(t, value.Equal(dt.Time()), "expected %d, got %d",
			value.UnixNano(), dt.Time().UnixNano())
		assert.Equal(t, MinuteTimePrecision, dt.Precision())
		assert.Equal(t, "2015-02-07T13:28+00:00", dt.String())
	}
}

func TestParseDateTimeNoMinutes(t *testing.T) {
	dt, err := ParseDateTime("2015-02-07T13Z")
	assert.Nil(t, dt, "unexpected date/time object")
	assert.NotNil(t, err, "expected error")
}

func TestParseDateTimeNoTime(t *testing.T) {
	dt, err := ParseDateTime("2015-02-07T")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		value := time.Date(2015, 2, 7, 0, 0, 0, 0, time.Local)
		assert.True(t, value.Equal(dt.Time()), "expected %d, got %d",
			value.UnixNano(), dt.Time().UnixNano())
		assert.Equal(t, DayDatePrecision, dt.Precision())
	}
}

func TestParseDateTimeNoDay(t *testing.T) {
	dt, err := ParseDateTime("2015-02T")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		value := time.Date(2015, 2, 1, 0, 0, 0, 0, time.Local)
		assert.True(t, value.Equal(dt.Time()), "expected %d, got %d",
			value.UnixNano(), dt.Time().UnixNano())
		assert.Equal(t, MonthDatePrecision, dt.Precision())
	}
}

func TestParseDateTimeNoMonth(t *testing.T) {
	dt, err := ParseDateTime("2015T")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date/time object") {
		value := time.Date(2015, 1, 1, 0, 0, 0, 0, time.Local)
		assert.True(t, value.Equal(dt.Time()), "expected %d, got %d",
			value.UnixNano(), dt.Time().UnixNano())
		assert.Equal(t, YearDatePrecision, dt.Precision())
	}
}

func TestMustEvalLocationInvalid(t *testing.T) {
	assert.Panics(t, func() { mustEvalLocation("X") })
}

func TestDateTimeEqualNil(t *testing.T) {
	assert.Equal(t, false, NewDateTime(time.Now()).Equal(nil))
}

func TestDateTimeEqualTypeDiffers(t *testing.T) {
	assert.Equal(t, false, NewDateTime(time.Now()).Equal(newAccessorMock()))
	assert.Equal(t, false, NewDateTime(time.Now()).Equivalent(newAccessorMock()))
}

func TestDateTimeEqualRightNil(t *testing.T) {
	assert.Equal(t, false, NewDateTime(time.Now()).Equal(nil))
	assert.Equal(t, false, NewDateTime(time.Now()).Equivalent(nil))
}

func TestDateTimeEqualEqual(t *testing.T) {
	now := time.Now()
	assert.Equal(t, true, NewDateTime(now).Equal(NewDateTime(now)))
	assert.Equal(t, true, NewDateTime(now).Equivalent(NewDateTime(now)))
}

func TestDateTimeEqualPrecisionDiffers(t *testing.T) {
	dt1, _ := ParseDateTime("2015-02-07T13:28:00")
	dt2, _ := ParseDateTime("2015-02-07T13:28")
	if assert.NotNil(t, dt1) && assert.NotNil(t, dt2) {
		assert.Equal(t, false, dt1.Equal(dt2))
		assert.Equal(t, true, dt1.Equivalent(dt2))
	}
}

func TestDateTimeEqualSecondNanoPrecisionDiffers(t *testing.T) {
	dt1, _ := ParseDateTime("2015-02-07T13:28:17.000")
	dt2, _ := ParseDateTime("2015-02-07T13:28:17")
	if assert.NotNil(t, dt1) && assert.NotNil(t, dt2) {
		assert.Equal(t, true, dt1.Equal(dt2))
		assert.Equal(t, true, dt1.Equivalent(dt2))
	}
}

func TestDateTimeEquivalent(t *testing.T) {
	dt1, _ := ParseDateTime("2015-02-07T13:28:00.00")
	dt2, _ := ParseDateTime("2015-02-07T13:28")
	if assert.NotNil(t, dt1) && assert.NotNil(t, dt2) {
		assert.Equal(t, false, dt1.Equal(dt2))
		assert.Equal(t, true, dt1.Equivalent(dt2))
	}
}

func TestDateTimeEqualDifferentTemporal(t *testing.T) {
	dt1, _ := ParseDateTime("2015-02-01")
	dt2, _ := ParseDate("2015-02")
	if assert.NotNil(t, dt1) && assert.NotNil(t, dt2) {
		assert.Equal(t, false, dt1.Equal(dt2))
	}
}

func TestDateTimeEquivalentDifferentTemporal(t *testing.T) {
	dt1, _ := ParseDateTime("2015-02-01")
	dt2, _ := ParseDate("2015-02")
	if assert.NotNil(t, dt1) && assert.NotNil(t, dt2) {
		assert.Equal(t, true, dt1.Equivalent(dt2))
	}
}

func TestDateTimeEquivalentTime(t *testing.T) {
	dt1, _ := ParseDateTime("2015-02-01T10:10:10")
	dt2, _ := ParseTime("10:10:10")
	if assert.NotNil(t, dt1) && assert.NotNil(t, dt2) {
		assert.Equal(t, false, dt1.Equivalent(dt2))
	}
}

func TestDateTimeEqualNotEqual(t *testing.T) {
	now := time.Now()
	assert.Equal(t, false, NewDateTime(now).Equal(NewDateTime(now.Add(time.Hour))))
	assert.Equal(t, false, NewDateTime(now).Equivalent(NewDateTime(now.Add(time.Hour))))
}

func TestNewDateTimeYMDHMSNWithInvalidPrecisionYear(t *testing.T) {
	assert.Panics(t, func() { NewDateTimeYMDHMSNWithPrecision(2000, 1, 1, 0, 0, 0, 0, time.UTC, -1) })
}

func TestNewDateTimeYMDHMSNWithInvalidPrecisionNano(t *testing.T) {
	assert.Panics(t, func() { NewDateTimeYMDHMSNWithPrecision(2000, 1, 1, 0, 0, 0, 0, time.UTC, -1) })
}

func TestNewDateTimeYMDHMSNWithPrecisionYear(t *testing.T) {
	v := NewDateTimeYMDHMSNWithPrecision(2019, 8, 21, 14, 38, 49, 827362627, time.Local, YearDatePrecision)
	date := v.Time()
	assert.Equal(t, 2019, date.Year())
	assert.Equal(t, 1, int(date.Month()))
	assert.Equal(t, 1, date.Day())
	assert.Equal(t, 0, date.Hour())
	assert.Equal(t, 0, date.Minute())
	assert.Equal(t, 0, date.Second())
	assert.Equal(t, 0, date.Nanosecond())
	assert.Same(t, time.Local, date.Location())
	assert.Equal(t, YearDatePrecision, v.Precision())
}

func TestNewDateTimeYMDHMSNWithPrecisionMonth(t *testing.T) {
	v := NewDateTimeYMDHMSNWithPrecision(2019, 8, 21, 14, 38, 49, 827362627, time.UTC, MonthDatePrecision)
	date := v.Time()
	assert.Equal(t, 2019, date.Year())
	assert.Equal(t, 8, int(date.Month()))
	assert.Equal(t, 1, date.Day())
	assert.Equal(t, 0, date.Hour())
	assert.Equal(t, 0, date.Minute())
	assert.Equal(t, 0, date.Second())
	assert.Equal(t, 0, date.Nanosecond())
	assert.Same(t, time.UTC, date.Location())
	assert.Equal(t, MonthDatePrecision, v.Precision())
}

func TestNewDateTimeYMDHMSNWithPrecisionDay(t *testing.T) {
	v := NewDateTimeYMDHMSNWithPrecision(2019, 8, 21, 14, 38, 49, 827362627, time.Local, DayDatePrecision)
	date := v.Time()
	assert.Equal(t, 2019, date.Year())
	assert.Equal(t, 8, int(date.Month()))
	assert.Equal(t, 21, date.Day())
	assert.Equal(t, 0, date.Hour())
	assert.Equal(t, 0, date.Minute())
	assert.Equal(t, 0, date.Second())
	assert.Equal(t, 0, date.Nanosecond())
	assert.Equal(t, DayDatePrecision, v.Precision())
}

func TestNewDateTimeYMDHMSNWithPrecisionHour(t *testing.T) {
	v := NewDateTimeYMDHMSNWithPrecision(2019, 8, 21, 14, 38, 49, 827362627, time.Local, HourTimePrecision)
	date := v.Time()
	assert.Equal(t, 2019, date.Year())
	assert.Equal(t, 8, int(date.Month()))
	assert.Equal(t, 21, date.Day())
	assert.Equal(t, 14, date.Hour())
	assert.Equal(t, 0, date.Minute())
	assert.Equal(t, 0, date.Second())
	assert.Equal(t, 0, date.Nanosecond())
	assert.Equal(t, HourTimePrecision, v.Precision())
}

func TestNewDateTimeYMDHMSNWithPrecisionMinute(t *testing.T) {
	v := NewDateTimeYMDHMSNWithPrecision(2019, 8, 21, 14, 38, 49, 827362627, time.Local, MinuteTimePrecision)
	date := v.Time()
	assert.Equal(t, 2019, date.Year())
	assert.Equal(t, 8, int(date.Month()))
	assert.Equal(t, 21, date.Day())
	assert.Equal(t, 14, date.Hour())
	assert.Equal(t, 38, date.Minute())
	assert.Equal(t, 0, date.Second())
	assert.Equal(t, 0, date.Nanosecond())
	assert.Equal(t, MinuteTimePrecision, v.Precision())
}

func TestNewDateTimeYMDHMSNWithPrecisionSecond(t *testing.T) {
	v := NewDateTimeYMDHMSNWithPrecision(2019, 8, 21, 14, 38, 49, 827362627, time.Local, SecondTimePrecision)
	date := v.Time()
	assert.Equal(t, 2019, date.Year())
	assert.Equal(t, 8, int(date.Month()))
	assert.Equal(t, 21, date.Day())
	assert.Equal(t, 14, date.Hour())
	assert.Equal(t, 38, date.Minute())
	assert.Equal(t, 49, date.Second())
	assert.Equal(t, 0, date.Nanosecond())
	assert.Equal(t, SecondTimePrecision, v.Precision())
}

func TestNewDateTimeYMDHMSNWithPrecisionNano(t *testing.T) {
	v := NewDateTimeYMDHMSNWithPrecision(2019, 8, 21, 14, 38, 49, 827362627, time.Local, NanoTimePrecision)
	date := v.Time()
	assert.Equal(t, 2019, date.Year())
	assert.Equal(t, 8, int(date.Month()))
	assert.Equal(t, 21, date.Day())
	assert.Equal(t, 14, date.Hour())
	assert.Equal(t, 38, date.Minute())
	assert.Equal(t, 49, date.Second())
	assert.Equal(t, 827362627, date.Nanosecond())
	assert.Equal(t, NanoTimePrecision, v.Precision())
}

func TestDateTimeAdd(t *testing.T) {
	v := NewDateTimeYMDHMSNWithPrecision(2019, 8, 21, 14, 38, 49, 827362627, time.Local, NanoTimePrecision)
	res, err := v.Add(NewQuantity(NewDecimalInt(12), NewString("day")))

	assert.NoError(t, err)
	if assert.Implements(t, (*DateTimeAccessor)(nil), res) {
		date := res.(DateTimeAccessor).Time()
		assert.Equal(t, 2019, date.Year())
		assert.Equal(t, 9, int(date.Month()))
		assert.Equal(t, 2, date.Day())
		assert.Equal(t, 14, date.Hour())
		assert.Equal(t, 38, date.Minute())
		assert.Equal(t, 49, date.Second())
		assert.Equal(t, 827362627, date.Nanosecond())
		assert.Equal(t, NanoTimePrecision, res.Precision())
	}
}

func TestDateTimeAddPrecision(t *testing.T) {
	v := NewDateTimeYMDHMSNWithPrecision(2019, 8, 21, 14, 38, 49, 827362627, time.Local, HourTimePrecision)
	res, err := v.Add(NewQuantity(NewDecimalInt(12), NewString("day")))

	assert.NoError(t, err)
	if assert.Implements(t, (*DateTimeAccessor)(nil), res) {
		date := res.(DateTimeAccessor).Time()
		assert.Equal(t, 2019, date.Year())
		assert.Equal(t, 9, int(date.Month()))
		assert.Equal(t, 2, date.Day())
		assert.Equal(t, 14, date.Hour())
		assert.Equal(t, 0, date.Minute())
		assert.Equal(t, 0, date.Second())
		assert.Equal(t, 0, date.Nanosecond())
		assert.Equal(t, HourTimePrecision, res.Precision())
	}
}

func TestDateTimeAddInvalidUnit(t *testing.T) {
	v := NewDateTimeYMDHMSNWithPrecision(2019, 8, 21, 14, 38, 49, 827362627, time.Local, NanoTimePrecision)
	_, err := v.Add(NewQuantity(NewDecimalInt(14), NewString("x")))

	assert.Error(t, err)
}

func TestDateTimeAddExceedsYear(t *testing.T) {
	v := NewDateTimeYMDHMSNWithPrecision(2019, 8, 21, 14, 38, 49, 827362627, time.Local, NanoTimePrecision)
	_, err := v.Add(NewQuantity(NewDecimalInt(-2020), NewString("year")))

	assert.Error(t, err)
}

func TestDateTimeCompareEqual(t *testing.T) {
	now := time.Now()
	res, status := NewDateTime(now).Compare(NewDateTime(now))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, 0, res)
}

func TestDateTimeCompareEqualTypeDiffers(t *testing.T) {
	res, status := NewDateTime(time.Now()).Compare(NewString("test1"))
	assert.Equal(t, Inconvertible, status)
	assert.Equal(t, -1, res)
}

func TestDateTimeCompareLessThan(t *testing.T) {
	now := time.Now()
	res, status := NewDateTime(now.Add(-time.Hour)).Compare(NewDateTime(now))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, -1, res)
}

func TestDateTimeCompareGreaterThan(t *testing.T) {
	now := time.Now()
	res, status := NewDateTime(now).Compare(NewDateTime(now.Add(-time.Hour)))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, 1, res)
}

func TestDateTimeComparePrecisionDiffers(t *testing.T) {
	res, status := NewDateTimeYMDHMSNWithPrecision(2020, 7, 21, 16, 0, 0, 0, time.Local, MinuteTimePrecision).
		Compare(NewDateTimeYMDHMSNWithPrecision(2020, 7, 21, 16, 0, 0, 0, time.Local, HourTimePrecision))
	assert.Equal(t, Empty, status)
	assert.Equal(t, -1, res)
}

func TestDateTimeCompareNanoSecondPrecision(t *testing.T) {
	res, status := NewDateTimeYMDHMSNWithPrecision(2020, 7, 21, 16, 18, 21, 0, time.Local, NanoTimePrecision).
		Compare(NewDateTimeYMDHMSNWithPrecision(2020, 7, 21, 16, 18, 21, 0, time.Local, SecondTimePrecision))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, 0, res)
}
