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

func TestDateDataType(t *testing.T) {
	o := NewDate(time.Now())
	dataType := o.DataType()
	assert.Equal(t, DateDataType, dataType)
}

func TestDateTypeLowestPrecision(t *testing.T) {
	o := NewDate(time.Now())
	assert.Equal(t, YearDatePrecision, o.LowestPrecision())
}

func TestDateTypeInfo(t *testing.T) {
	o := NewDate(time.Now())
	i := o.TypeInfo()
	if assert.NotNil(t, i, "type info expected") {
		assert.Equal(t, "System.Date", i.String())
		if assert.NotNil(t, i.FQBaseName(), "base name expected") {
			assert.Equal(t, "System.Any", i.FQBaseName().String())
		}
	}
}

func TestDateValue(t *testing.T) {
	testTime := time.Now().Add(-time.Hour * 78)
	o := NewDate(testTime)

	value := o.Time()
	assert.Equal(t, testTime.Year(), o.Year())
	assert.Equal(t, int(testTime.Month()), o.Month())
	assert.Equal(t, testTime.Day(), o.Day())
	assert.Equal(t, DayDatePrecision, o.Precision())

	expectedTime := time.Date(testTime.Year(), testTime.Month(), testTime.Day(), 0, 0, 0, 0, time.Local)
	assert.True(t, expectedTime.Equal(value), "expected %d, got %d",
		expectedTime.UnixNano(), value.UnixNano())
}

func TestDateYMD(t *testing.T) {
	o := NewDateYMD(2020, 4, 23)

	assert.Equal(t, 2020, o.Year())
	assert.Equal(t, 4, o.Month())
	assert.Equal(t, 23, o.Day())
	assert.Equal(t, DayDatePrecision, o.Precision())
	assert.Equal(t, "2020-04-23", o.String())
}

func TestParseDateComplete(t *testing.T) {
	dt, err := ParseDate("2015-02-07")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date object") {
		value := time.Date(2015, 2, 7, 0, 0, 0, 0, time.Local)
		assert.True(t, value.Equal(dt.Time()), "expected %d, got %d",
			value.UnixNano(), dt.Time().UnixNano())
		assert.Equal(t, DayDatePrecision, dt.Precision())
	}
}

func TestParseDateInvalid(t *testing.T) {
	dt, err := ParseDate("2015-02-0A")
	assert.Nil(t, dt, "unexpected date object")
	assert.NotNil(t, err, "expected error")
}

func TestParseDateNoDay(t *testing.T) {
	dt, err := ParseDate("2015-02")
	assert.Nil(t, err, "unexpected error")
	if assert.NotNil(t, dt, "expected date object") {
		value := time.Date(2015, 2, 1, 0, 0, 0, 0, time.Local)
		assert.True(t, value.Equal(dt.Time()), "expected %d, got %d",
			value.UnixNano(), dt.Time().UnixNano())
		assert.Equal(t, MonthDatePrecision, dt.Precision())
		assert.Equal(t, "2015-02", dt.String())
	}
}

func TestParseDateNoMonth(t *testing.T) {
	dt, err := ParseDate("2015")
	if assert.NotNil(t, dt, "expected date object") {
		assert.Nil(t, err, "unexpected error")
		value := time.Date(2015, 1, 1, 0, 0, 0, 0, time.Local)
		assert.True(t, value.Equal(dt.Time()), "expected %d, got %d",
			value.UnixNano(), dt.Time().UnixNano())
		assert.Equal(t, YearDatePrecision, dt.Precision())
		assert.Equal(t, "2015", dt.String())
	}
}

func TestDateEqualTypeDiffers(t *testing.T) {
	assert.Equal(t, false, NewDate(time.Now()).Equal(newAccessorMock()))
	assert.Equal(t, false, NewDate(time.Now()).Equivalent(newAccessorMock()))
}

func TestDateEqualRightNil(t *testing.T) {
	assert.Equal(t, false, NewDate(time.Now()).Equal(nil))
	assert.Equal(t, false, NewDate(time.Now()).Equivalent(nil))
}

func TestDateEqualEqual(t *testing.T) {
	now := time.Now()
	assert.Equal(t, true, NewDate(now).Equal(NewDate(now)))
	assert.Equal(t, true, NewDate(now).Equivalent(NewDate(now)))
}

func TestDateEqualDateTime(t *testing.T) {
	dt := NewDateTime(time.Date(2018, 7, 28, 0, 0, 0, 0, time.Local))
	d := NewDateYMD(2018, 7, 28)
	assert.Equal(t, false, d.Equal(dt))
	assert.Equal(t, true, d.Equivalent(dt))
}

func TestDateEqualNotTime(t *testing.T) {
	timeOnly := NewTimeHMSN(0, 0, 0, 0)
	d := NewDateYMD(2018, 7, 28)
	assert.Equal(t, false, d.Equal(timeOnly))
	assert.Equal(t, false, d.Equivalent(timeOnly))
}

func TestDateEqualNotEqual(t *testing.T) {
	now := time.Now()
	assert.Equal(t, false, NewDate(now).Equal(NewDate(now.Add(48*time.Hour))))
	assert.Equal(t, false, NewDate(now).Equivalent(NewDate(now.Add(48*time.Hour))))
}

func TestDateEqualPrecisionDiffers(t *testing.T) {
	d1, _ := ParseDate("2015-02-07")
	d2, _ := ParseDate("2015-02")
	if assert.NotNil(t, d1) && assert.NotNil(t, d2) {
		assert.Equal(t, false, d1.Equal(d2))
		assert.Equal(t, false, d1.Equivalent(d2))
	}
}

func TestDateEquivalent(t *testing.T) {
	d1, _ := ParseDate("2015-02-01")
	d2, _ := ParseDate("2015-02")
	if assert.NotNil(t, d1) && assert.NotNil(t, d2) {
		assert.Equal(t, false, d1.Equal(d2))
		assert.Equal(t, true, d1.Equivalent(d2))
	}
}

func TestDateEqualYearDiffers(t *testing.T) {
	d1 := NewDateYMD(2020, 2, 3)
	d2 := NewDateYMD(2021, 2, 3)
	assert.Equal(t, false, d1.Equal(d2))
	assert.Equal(t, false, d1.Equivalent(d2))
}

func TestDateEqualMonthDiffers(t *testing.T) {
	d1 := NewDateYMD(2020, 2, 3)
	d2 := NewDateYMD(2020, 3, 3)
	assert.Equal(t, false, d1.Equal(d2))
	assert.Equal(t, false, d1.Equivalent(d2))
}

func TestDateEqualDayDiffers(t *testing.T) {
	d1 := NewDateYMD(2020, 2, 3)
	d2 := NewDateYMD(2020, 2, 4)
	assert.Equal(t, false, d1.Equal(d2))
	assert.Equal(t, false, d1.Equivalent(d2))
}
