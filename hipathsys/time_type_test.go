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

func TestTimeSource(t *testing.T) {
	o := NewTimeWithSource(time.Now(), "abc")
	assert.Equal(t, "abc", o.Source())
}

func TestTimeDataType(t *testing.T) {
	o := NewTime(time.Now())
	dataType := o.DataType()
	assert.Equal(t, TimeDataType, dataType)
}

func TestTimeTypeLowestPrecision(t *testing.T) {
	o := NewTime(time.Now())
	assert.Equal(t, HourTimePrecision, o.LowestPrecision())
}

func TestTimeTypeSpec(t *testing.T) {
	o := NewTime(time.Now())
	i := o.TypeSpec()
	if assert.NotNil(t, i, "type info expected") {
		assert.Equal(t, "System.Time", i.String())
		if assert.NotNil(t, i.FQBaseName(), "base name expected") {
			assert.Equal(t, "System.Any", i.FQBaseName().String())
		}
	}
}

func TestTimeTypeInfo(t *testing.T) {
	o := NewTime(time.Now())
	i := o.TypeInfo()
	if assert.Implements(t, (*SimpleTypeInfoAccessor)(nil), i) {
		a := i.(SimpleTypeInfoAccessor)
		assert.Equal(t, NewString("System"), a.Namespace())
		assert.Equal(t, NewString("Time"), a.Name())
		assert.Equal(t, NewString("System.Any"), a.BaseType())
	}
}

func TestTimeValue(t *testing.T) {
	testTime := time.Now().Add(-time.Hour * 78)
	o := NewTime(testTime)

	assert.Equal(t, testTime.Hour(), o.Hour())
	assert.Equal(t, testTime.Minute(), o.Minute())
	assert.Equal(t, testTime.Second(), o.Second())
	assert.Equal(t, testTime.Nanosecond(), o.Nanosecond())
	assert.Equal(t, NanoTimePrecision, o.Precision())
}

func TestTimeYMD(t *testing.T) {
	o := NewTimeHMSN(16, 28, 47, 837173635)

	assert.Equal(t, 16, o.Hour())
	assert.Equal(t, 28, o.Minute())
	assert.Equal(t, 47, o.Second())
	assert.Equal(t, 837173635, o.Nanosecond())
	assert.Equal(t, NanoTimePrecision, o.Precision())
	assert.Equal(t, "16:28:47.837173635", o.String())
}

func TestParseTimeValueComplete(t *testing.T) {
	source := "test source"
	dt, err := ParseTimeWithSource("13:28:17.239", source)
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected time object") {
		assert.Equal(t, 13, dt.Hour())
		assert.Equal(t, 28, dt.Minute())
		assert.Equal(t, 17, dt.Second())
		assert.Equal(t, 239000000, dt.Nanosecond())
		assert.Equal(t, NanoTimePrecision, dt.Precision())
		assert.Equal(t, source, dt.Source())
	}
}

func TestParseTimeValueInvalid(t *testing.T) {
	dt, err := ParseTime("13:28:17.A")
	assert.Nil(t, dt, "unexpected date object")
	assert.NotNil(t, err, "expected error")
}

func TestParseTimeValueFractionDigits(t *testing.T) {
	dt, err := ParseTime("13:28:17.2397381239")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected time object") {
		assert.Equal(t, 13, dt.Hour())
		assert.Equal(t, 28, dt.Minute())
		assert.Equal(t, 17, dt.Second())
		assert.Equal(t, 239738123, dt.Nanosecond())
		assert.Equal(t, NanoTimePrecision, dt.Precision())
		assert.Nil(t, dt.Source())
	}
}

func TestParseTimeValueNoNanos(t *testing.T) {
	dt, err := ParseTime("13:28:17")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected time object") {
		assert.Equal(t, 13, dt.Hour())
		assert.Equal(t, 28, dt.Minute())
		assert.Equal(t, 17, dt.Second())
		assert.Equal(t, 0, dt.Nanosecond())
		assert.Equal(t, SecondTimePrecision, dt.Precision())
	}
}

func TestParseTimeValueNoSeconds(t *testing.T) {
	dt, err := ParseTime("13:28")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected time object") {
		assert.Equal(t, 13, dt.Hour())
		assert.Equal(t, 28, dt.Minute())
		assert.Equal(t, 0, dt.Second())
		assert.Equal(t, 0, dt.Nanosecond())
		assert.Equal(t, MinuteTimePrecision, dt.Precision())
		assert.Equal(t, "13:28", dt.String())
	}
}

func TestParseTimeValueNoMinutes(t *testing.T) {
	dt, err := ParseTime("13")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected time object") {
		assert.Equal(t, 13, dt.Hour())
		assert.Equal(t, 0, dt.Minute())
		assert.Equal(t, 0, dt.Second())
		assert.Equal(t, 0, dt.Nanosecond())
		assert.Equal(t, HourTimePrecision, dt.Precision())
		assert.Equal(t, "13", dt.String())
	}
}

func TestParseNanosecondEmpty(t *testing.T) {
	assert.Equal(t, 0, parseNanosecond(""))
}

func TestTimeEqualNil(t *testing.T) {
	assert.Equal(t, false, NewTime(time.Now()).Equal(nil))
}

func TestTimeEqualTypeDiffers(t *testing.T) {
	assert.Equal(t, false, NewTime(time.Now()).Equal(newAccessorMock()))
	assert.Equal(t, false, NewTime(time.Now()).Equivalent(newAccessorMock()))
}

func TestTimeEqualRightNil(t *testing.T) {
	assert.Equal(t, false, NewTime(time.Now()).Equal(nil))
	assert.Equal(t, false, NewTime(time.Now()).Equivalent(nil))
}

func TestTimeEqualEqual(t *testing.T) {
	now := time.Now()
	assert.Equal(t, true, NewTime(now).Equal(NewTime(now)))
	assert.Equal(t, true, NewTime(now).Equivalent(NewTime(now)))
}

func TestTimeEqualNotEqual(t *testing.T) {
	now := time.Now()
	assert.Equal(t, false, NewTime(now).Equal(NewTime(now.Add(time.Hour))))
	assert.Equal(t, false, NewTime(now).Equivalent(NewTime(now.Add(time.Hour))))
}

func TestTimeEqualPrecisionDiffers(t *testing.T) {
	t1, _ := ParseTime("17:22:00")
	t2, _ := ParseTime("17:22")
	if assert.NotNil(t, t1) && assert.NotNil(t, t2) {
		assert.Equal(t, false, t1.Equal(t2))
		assert.Equal(t, true, t1.Equivalent(t2))
	}
}

func TestTimeEqualSecondNanoPrecisionDiffers(t *testing.T) {
	t1, _ := ParseTime("17:22:21.000")
	t2, _ := ParseTime("17:22:21")
	if assert.NotNil(t, t1) && assert.NotNil(t, t2) {
		assert.Equal(t, true, t1.Equal(t2))
		assert.Equal(t, true, t1.Equivalent(t2))
	}
}

func TestTimeEquivalent(t *testing.T) {
	t1, _ := ParseTime("17:22:00.00")
	t2, _ := ParseTime("17:22")
	if assert.NotNil(t, t1) && assert.NotNil(t, t2) {
		assert.Equal(t, false, t1.Equal(t2))
		assert.Equal(t, true, t1.Equivalent(t2))
	}
}

func TestTimeEqualHourDiffers(t *testing.T) {
	t1 := NewTimeHMSN(17, 23, 41, 231)
	t2 := NewTimeHMSN(18, 23, 41, 231)
	assert.Equal(t, false, t1.Equal(t2))
	assert.Equal(t, false, t1.Equivalent(t2))
}

func TestTimeEqualMinuteDiffers(t *testing.T) {
	t1 := NewTimeHMSN(17, 23, 41, 231)
	t2 := NewTimeHMSN(17, 24, 41, 231)
	assert.Equal(t, false, t1.Equal(t2))
	assert.Equal(t, false, t1.Equivalent(t2))
}

func TestTimeEqualSecondDiffers(t *testing.T) {
	t1 := NewTimeHMSN(17, 23, 41, 231)
	t2 := NewTimeHMSN(17, 23, 42, 231)
	assert.Equal(t, false, t1.Equal(t2))
	assert.Equal(t, false, t1.Equivalent(t2))
}

func TestTimeEqualNanosecondDiffers(t *testing.T) {
	t1 := NewTimeHMSN(17, 23, 41, 231)
	t2 := NewTimeHMSN(17, 23, 41, 232)
	assert.Equal(t, false, t1.Equal(t2))
	assert.Equal(t, false, t1.Equivalent(t2))
}

func TestNewTimeHMSNWithPrecisionDay(t *testing.T) {
	assert.Panics(t, func() { NewTimeHMSNWithPrecision(0, 0, 0, 0, DayDatePrecision) })
}

func TestNewTimeHMSNWithPrecisionInvalid(t *testing.T) {
	assert.Panics(t, func() { NewTimeHMSNWithPrecision(0, 0, 0, 0, 100) })
}

func TestNewTimeHMSNWithPrecisionHour(t *testing.T) {
	v := NewTimeHMSNWithPrecision(17, 28, 31, 823123876, HourTimePrecision)
	assert.Equal(t, 17, v.Hour())
	assert.Equal(t, 0, v.Minute())
	assert.Equal(t, 0, v.Second())
	assert.Equal(t, 0, v.Nanosecond())
	assert.Equal(t, HourTimePrecision, v.Precision())
}

func TestNewTimeHMSNWithPrecisionMinute(t *testing.T) {
	v := NewTimeHMSNWithPrecision(17, 28, 31, 823123876, MinuteTimePrecision)
	assert.Equal(t, 17, v.Hour())
	assert.Equal(t, 28, v.Minute())
	assert.Equal(t, 0, v.Second())
	assert.Equal(t, 0, v.Nanosecond())
	assert.Equal(t, MinuteTimePrecision, v.Precision())
}

func TestNewTimeHMSNWithPrecisionSecond(t *testing.T) {
	v := NewTimeHMSNWithPrecision(17, 28, 31, 823123876, SecondTimePrecision)
	assert.Equal(t, 17, v.Hour())
	assert.Equal(t, 28, v.Minute())
	assert.Equal(t, 31, v.Second())
	assert.Equal(t, 0, v.Nanosecond())
	assert.Equal(t, SecondTimePrecision, v.Precision())
}

func TestNewTimeHMSNWithPrecisionNano(t *testing.T) {
	v := NewTimeHMSNWithPrecision(17, 28, 31, 823123876, NanoTimePrecision)
	assert.Equal(t, 17, v.Hour())
	assert.Equal(t, 28, v.Minute())
	assert.Equal(t, 31, v.Second())
	assert.Equal(t, 823123876, v.Nanosecond())
	assert.Equal(t, NanoTimePrecision, v.Precision())
}

func TestTimeAdd(t *testing.T) {
	v := NewTimeHMSNWithPrecision(17, 28, 31, 823123876, NanoTimePrecision)
	res, err := v.Add(NewQuantity(NewDecimalInt(34), NewString("minute")))

	assert.NoError(t, err)
	if assert.Implements(t, (*TimeAccessor)(nil), res) {
		r := res.(TimeAccessor)
		assert.Equal(t, 18, r.Hour())
		assert.Equal(t, 2, r.Minute())
		assert.Equal(t, 31, r.Second())
		assert.Equal(t, 823123876, r.Nanosecond())
		assert.Equal(t, NanoTimePrecision, r.Precision())
	}
}

func TestTimeAddPrecision(t *testing.T) {
	v := NewTimeHMSNWithPrecision(17, 28, 31, 823123876, HourTimePrecision)
	res, err := v.Add(NewQuantity(NewDecimalInt(80), NewString("minute")))

	assert.NoError(t, err)
	if assert.Implements(t, (*TimeAccessor)(nil), res) {
		r := res.(TimeAccessor)
		assert.Equal(t, 18, r.Hour())
		assert.Equal(t, 0, r.Minute())
		assert.Equal(t, 0, r.Second())
		assert.Equal(t, 0, r.Nanosecond())
		assert.Equal(t, HourTimePrecision, r.Precision())
	}
}

func TestTimeAddInvalidUnit(t *testing.T) {
	v := NewTimeHMSNWithPrecision(17, 28, 31, 823123876, NanoTimePrecision)
	_, err := v.Add(NewQuantity(NewDecimalInt(34), NewString("x")))

	assert.Error(t, err)
}

func TestTimeAddUnsupportedPrecision(t *testing.T) {
	v := NewTimeHMSNWithPrecision(17, 28, 31, 823123876, NanoTimePrecision)
	_, err := v.Add(NewQuantity(NewDecimalInt(34), NewString("day")))

	assert.Error(t, err)
}

func TestTimeAddExceedsYear(t *testing.T) {
	v := NewTimeHMSNWithPrecision(17, 28, 31, 823123876, NanoTimePrecision)
	_, err := v.Add(NewQuantity(NewDecimalInt(-15_000), NewString("hour")))

	assert.Error(t, err)
}

func TestTimeAddExceedsDay(t *testing.T) {
	v := NewTimeHMSNWithPrecision(17, 28, 31, 823123876, NanoTimePrecision)
	_, err := v.Add(NewQuantity(NewDecimalInt(9), NewString("hour")))

	assert.Error(t, err)
}

func TestTimeAddExceedsMonth(t *testing.T) {
	v := NewTimeHMSNWithPrecision(17, 28, 31, 823123876, NanoTimePrecision)
	_, err := v.Add(NewQuantity(NewDecimalInt(744), NewString("hour")))

	assert.Error(t, err)
}

func TestTimeCompareEqual(t *testing.T) {
	now := time.Now()
	res, status := NewTime(now).Compare(NewTime(now))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, 0, res)
}

func TestTimeCompareEqualTypeDiffers(t *testing.T) {
	res, status := NewTime(time.Now()).Compare(NewString("test1"))
	assert.Equal(t, Inconvertible, status)
	assert.Equal(t, -1, res)
}

func TestTimeCompareLessThan(t *testing.T) {
	res, status := NewTimeHMSN(17, 18, 21, 0).Compare(NewTimeHMSN(19, 15, 17, 0))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, -1, res)
}

func TestTimeCompareGreaterThan(t *testing.T) {
	res, status := NewTimeHMSN(19, 15, 17, 0).Compare(NewTimeHMSN(17, 18, 21, 0))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, 1, res)
}

func TestTimeComparePrecisionDiffers(t *testing.T) {
	res, status := NewTimeHMSNWithPrecision(14, 28, 0, 0, SecondTimePrecision).
		Compare(NewTimeHMSNWithPrecision(14, 28, 0, 0, MinuteTimePrecision))
	assert.Equal(t, Empty, status)
	assert.Equal(t, -1, res)
}

func TestTimeCompareNanoSecondPrecision(t *testing.T) {
	res, status := NewTimeHMSNWithPrecision(14, 28, 0, 0, NanoTimePrecision).
		Compare(NewTimeHMSNWithPrecision(14, 28, 0, 0, SecondTimePrecision))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, 0, res)
}

func TestTimeCompareHourDiffers(t *testing.T) {
	res, status := NewTimeHMSNWithPrecision(14, 28, 17, 987, NanoTimePrecision).
		Compare(NewTimeHMSNWithPrecision(15, 28, 17, 987, NanoTimePrecision))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, -1, res)
}

func TestTimeCompareMinuteDiffers(t *testing.T) {
	res, status := NewTimeHMSNWithPrecision(14, 29, 17, 987, NanoTimePrecision).
		Compare(NewTimeHMSNWithPrecision(14, 28, 17, 987, NanoTimePrecision))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, 1, res)
}

func TestTimeCompareSecondDiffers(t *testing.T) {
	res, status := NewTimeHMSNWithPrecision(14, 28, 17, 987, NanoTimePrecision).
		Compare(NewTimeHMSNWithPrecision(14, 28, 18, 987, NanoTimePrecision))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, -1, res)
}

func TestTimeCompareNanoDiffers(t *testing.T) {
	res, status := NewTimeHMSNWithPrecision(14, 28, 17, 988, NanoTimePrecision).
		Compare(NewTimeHMSNWithPrecision(14, 28, 17, 987, NanoTimePrecision))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, 1, res)
}
