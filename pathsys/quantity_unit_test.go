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

func TestNewQuantityUnitWithUCUM(t *testing.T) {
	qu := NewQuantityUnitWithUCUM("abc")
	assert.Nil(t, qu.Singular())
	assert.Nil(t, qu.Plural())
	assert.Equal(t, NewString("abc"), qu.UCUM())
}

func TestQuantityUnitWithNewUCUM(t *testing.T) {
	qu, exp := QuantityUnitWithName("abc")
	assert.Nil(t, qu.Singular())
	assert.Nil(t, qu.Plural())
	assert.Equal(t, NewString("abc"), qu.UCUM())
	assert.Equal(t, 1, exp)
}

func TestQuantityUnitWithExistingName(t *testing.T) {
	qu, exp := QuantityUnitWithName("years")
	assert.Same(t, YearQuantityUnit, qu)
	assert.Equal(t, 1, exp)
}

func TestQuantityUnitWithEmptyName(t *testing.T) {
	qu, exp := QuantityUnitWithName("")
	assert.Nil(t, qu)
	assert.Equal(t, 1, exp)
}

func TestQuantityUnitWithNameStringNew(t *testing.T) {
	qu, exp := QuantityUnitWithNameString(NewString("abc"))
	assert.Nil(t, qu.Singular())
	assert.Nil(t, qu.Plural())
	assert.Equal(t, NewString("abc"), qu.UCUM())
	assert.Equal(t, 1, exp)
}

func TestQuantityUnitWithStringExisting(t *testing.T) {
	qu, exp := QuantityUnitWithNameString(NewString("years"))
	assert.Same(t, YearQuantityUnit, qu)
	assert.Equal(t, 1, exp)
}

func TestQuantityUnitWithNameStringNil(t *testing.T) {
	qu, exp := QuantityUnitWithNameString(nil)
	assert.Nil(t, qu)
	assert.Equal(t, 1, exp)
}

func TestQuantityUnitWithNameStringExpDefault(t *testing.T) {
	qu, exp := QuantityUnitWithNameString(NewString("1"))
	assert.Nil(t, qu.Singular())
	assert.Nil(t, qu.Plural())
	assert.Equal(t, NewString("1"), qu.UCUM())
	assert.Equal(t, 1, exp)
}

func TestQuantityUnitWithNameStringExpNew(t *testing.T) {
	qu, exp := QuantityUnitWithNameString(NewString("abc2"))
	assert.Nil(t, qu.Singular())
	assert.Nil(t, qu.Plural())
	assert.Equal(t, NewString("abc"), qu.UCUM())
	assert.Equal(t, 2, exp)
}

func TestQuantityUnitWithStringExpExisting(t *testing.T) {
	qu, exp := QuantityUnitWithNameString(NewString("years2"))
	assert.Same(t, YearQuantityUnit, qu)
	assert.Equal(t, 2, exp)
}

func TestQuantityUnitByNameEmpty(t *testing.T) {
	assert.Nil(t, QuantityUnitByName(""))
}

func TestQuantityUnitByName(t *testing.T) {
	assert.Same(t, MinuteQuantityUnit, QuantityUnitByName("minutes"))
}

func TestQuantityUnitByNameStringNil(t *testing.T) {
	assert.Nil(t, QuantityUnitByNameString(nil))
}

func TestQuantityUnitNanosecond(t *testing.T) {
	assert.Same(t, NanosecondQuantityUnit, QuantityUnitByNameString(NewString("nanosecond")))
	assert.Same(t, NanosecondQuantityUnit, QuantityUnitByNameString(NewString("nanoseconds")))
	assert.Same(t, NanosecondQuantityUnit, QuantityUnitByNameString(NewString("ns")))
	if assert.NotNil(t, NanosecondQuantityUnit.RootBase(), "root base expected") {
		assert.Same(t, SecondQuantityUnit, NanosecondQuantityUnit.RootBase().Unit())
	}
	assert.True(t, IsCalendarDurationUnit(NanosecondQuantityUnit))
}

func TestQuantityUnitNanosecondConvert(t *testing.T) {
	v1, v2, u := ConvertUnitToBase(
		NewDecimalInt(2_087_123_543), NanosecondQuantityUnit, 1,
		NewDecimalInt(14), MillisecondQuantityUnit, 1, true)
	if assert.Same(t, SecondQuantityUnit, u) {
		assert.Equal(t, 2.087_123_543, v1.Float64())
		assert.Equal(t, 0.014, v2.Float64())
	}
}

func TestQuantityUnitMillisecond(t *testing.T) {
	assert.Same(t, MillisecondQuantityUnit, QuantityUnitByNameString(NewString("millisecond")))
	assert.Same(t, MillisecondQuantityUnit, QuantityUnitByNameString(NewString("milliseconds")))
	assert.Same(t, MillisecondQuantityUnit, QuantityUnitByNameString(NewString("ms")))
	if assert.NotNil(t, MillisecondQuantityUnit.RootBase(), "root base expected") {
		assert.Same(t, SecondQuantityUnit, MillisecondQuantityUnit.RootBase().Unit())
	}
	assert.True(t, IsCalendarDurationUnit(MillisecondQuantityUnit))
}

func TestQuantityUnitMillisecondConvert(t *testing.T) {
	v1, v2, u := ConvertUnitToBase(
		NewDecimalInt(7_827), MillisecondQuantityUnit, 1,
		NewDecimalInt(8), SecondQuantityUnit, 1, true)
	if assert.Same(t, SecondQuantityUnit, u) {
		assert.Equal(t, 7.827, v1.Float64())
		assert.Equal(t, 8.0, v2.Float64())
	}
}

func TestQuantityUnitSecond(t *testing.T) {
	assert.Same(t, SecondQuantityUnit, QuantityUnitByNameString(NewString("second")))
	assert.Same(t, SecondQuantityUnit, QuantityUnitByNameString(NewString("seconds")))
	assert.Same(t, SecondQuantityUnit, QuantityUnitByNameString(NewString("s")))
	assert.Nil(t, SecondQuantityUnit.RootBase())
	assert.True(t, IsCalendarDurationUnit(SecondQuantityUnit))
}

func TestQuantityUnitSecondConvert(t *testing.T) {
	v1, v2, u := ConvertUnitToBase(
		NewDecimalInt(8), SecondQuantityUnit, 1,
		NewDecimalInt(7_827), MillisecondQuantityUnit, 1, true)
	if assert.Same(t, SecondQuantityUnit, u) {
		assert.Equal(t, 8.0, v1.Float64())
		assert.Equal(t, 7.827, v2.Float64())
	}
}

func TestQuantityUnitMinute(t *testing.T) {
	assert.Same(t, MinuteQuantityUnit, QuantityUnitByNameString(NewString("minute")))
	assert.Same(t, MinuteQuantityUnit, QuantityUnitByNameString(NewString("minutes")))
	if assert.NotNil(t, MinuteQuantityUnit.RootBase(), "root base expected") {
		assert.Same(t, SecondQuantityUnit, MinuteQuantityUnit.RootBase().Unit())
	}
	assert.True(t, IsCalendarDurationUnit(MinuteQuantityUnit))
}

func TestQuantityUnitMinuteConvert(t *testing.T) {
	v1, v2, u := ConvertUnitToBase(
		NewDecimalInt(8), SecondQuantityUnit, 1,
		NewDecimalInt(4), MinuteQuantityUnit, 1, true)
	if assert.Same(t, SecondQuantityUnit, u) {
		assert.Equal(t, 8.0, v1.Float64())
		assert.Equal(t, 240.0, v2.Float64())
	}
}

func TestQuantityUnitHour(t *testing.T) {
	assert.Same(t, HourQuantityUnit, QuantityUnitByNameString(NewString("hour")))
	assert.Same(t, HourQuantityUnit, QuantityUnitByNameString(NewString("hours")))
	if assert.NotNil(t, HourQuantityUnit.RootBase(), "root base expected") {
		assert.Same(t, SecondQuantityUnit, HourQuantityUnit.RootBase().Unit())
	}
	assert.True(t, IsCalendarDurationUnit(HourQuantityUnit))
}

func TestQuantityUnitHourConvert(t *testing.T) {
	v1, v2, u := ConvertUnitToBase(
		NewDecimalInt(8), MinuteQuantityUnit, 1,
		NewDecimalInt(4), HourQuantityUnit, 1, true)
	if assert.Same(t, SecondQuantityUnit, u) {
		assert.Equal(t, 8.0*60.0, v1.Float64())
		assert.Equal(t, 4.0*60.0*60.0, v2.Float64())
	}
}

func TestQuantityUnitDay(t *testing.T) {
	assert.Same(t, DayQuantityUnit, QuantityUnitByNameString(NewString("day")))
	assert.Same(t, DayQuantityUnit, QuantityUnitByNameString(NewString("days")))
	if assert.NotNil(t, DayQuantityUnit.RootBase(), "root base expected") {
		assert.Same(t, SecondQuantityUnit, DayQuantityUnit.RootBase().Unit())
	}
	assert.True(t, IsCalendarDurationUnit(DayQuantityUnit))
}

func TestQuantityUnitDayConvert(t *testing.T) {
	v1, v2, u := ConvertUnitToBase(
		NewDecimalInt(8), HourQuantityUnit, 1,
		NewDecimalInt(4), DayQuantityUnit, 1, true)
	if assert.Same(t, SecondQuantityUnit, u) {
		assert.Equal(t, 8.0*60.0*60.0, v1.Float64())
		assert.Equal(t, 4.0*24.0*60.0*60.0, v2.Float64())
	}
}

func TestQuantityUnitWeek(t *testing.T) {
	assert.Same(t, WeekQuantityUnit, QuantityUnitByNameString(NewString("week")))
	assert.Same(t, WeekQuantityUnit, QuantityUnitByNameString(NewString("weeks")))
	if assert.NotNil(t, WeekQuantityUnit.RootBase(), "root base expected") {
		assert.Same(t, SecondQuantityUnit, WeekQuantityUnit.RootBase().Unit())
	}
	assert.True(t, IsCalendarDurationUnit(WeekQuantityUnit))
}

func TestQuantityUnitWeekConvert(t *testing.T) {
	v1, v2, u := ConvertUnitToBase(
		NewDecimalInt(8), DayQuantityUnit, 1,
		NewDecimalInt(4), WeekQuantityUnit, 1, true)
	if assert.Same(t, SecondQuantityUnit, u) {
		assert.Equal(t, 8.0*24.0*60.0*60.0, v1.Float64())
		assert.Equal(t, 4.0*7.0*24.0*60.0*60.0, v2.Float64())
	}
}

func TestQuantityUnitMonth(t *testing.T) {
	assert.Same(t, MonthQuantityUnit, QuantityUnitByNameString(NewString("month")))
	assert.Same(t, MonthQuantityUnit, QuantityUnitByNameString(NewString("months")))
	if assert.NotNil(t, MonthQuantityUnit.RootBase(), "root base expected") {
		assert.Same(t, SecondQuantityUnit, MonthQuantityUnit.RootBase().Unit())
	}
	assert.True(t, IsCalendarDurationUnit(MonthQuantityUnit))
}

func TestQuantityUnitMonthConvert(t *testing.T) {
	v1, v2, u := ConvertUnitToBase(
		NewDecimalInt(8), DayQuantityUnit, 1,
		NewDecimalInt(4), MonthQuantityUnit, 1, true)
	if assert.Same(t, SecondQuantityUnit, u) {
		assert.Equal(t, 8.0*24.0*60.0*60.0, v1.Float64())
		assert.Equal(t, 4.0*30.0*24.0*60.0*60.0, v2.Float64())
	}
}

func TestQuantityUnitYear(t *testing.T) {
	assert.Same(t, YearQuantityUnit, QuantityUnitByNameString(NewString("year")))
	assert.Same(t, YearQuantityUnit, QuantityUnitByNameString(NewString("years")))
	if assert.NotNil(t, YearQuantityUnit.RootBase(), "root base expected") {
		assert.Same(t, SecondQuantityUnit, YearQuantityUnit.RootBase().Unit())
	}
	assert.True(t, IsCalendarDurationUnit(YearQuantityUnit))
}

func TestQuantityUnitYearConvert(t *testing.T) {
	v1, v2, u := ConvertUnitToBase(
		NewDecimalInt(8), DayQuantityUnit, 1,
		NewDecimalInt(4), YearQuantityUnit, 1, true)
	if assert.Same(t, SecondQuantityUnit, u) {
		assert.Equal(t, 8.0*24.0*60.0*60.0, v1.Float64())
		assert.Equal(t, 4.0*365.0*24.0*60.0*60.0, v2.Float64())
	}
}

func TestQuantityUnitYearMonthConvert(t *testing.T) {
	v1, v2, u := ConvertUnitToBase(
		NewDecimalInt(8), MonthQuantityUnit, 1,
		NewDecimalInt(4), YearQuantityUnit, 1, true)
	if assert.Same(t, MonthQuantityUnit, u) {
		assert.Equal(t, 8.0, v1.Float64())
		assert.Equal(t, 4.0*12.0, v2.Float64())
	}
}

func TestQuantityUnitMonthYearConvert(t *testing.T) {
	v1, v2, u := ConvertUnitToBase(
		NewDecimalInt(8), YearQuantityUnit, 1,
		NewDecimalInt(4), MonthQuantityUnit, 1, true)
	if assert.Same(t, MonthQuantityUnit, u) {
		assert.Equal(t, 8.0*12.0, v1.Float64())
		assert.Equal(t, 4.0, v2.Float64())
	}
}

func TestQuantityUnitFactorNoBase(t *testing.T) {
	assert.Nil(t, MonthQuantityUnit.Factor(NewQuantityUnit("", "", ""), 1),
		"no factor expected")
}

func TestConvertUnitToBaseNoCommonBase(t *testing.T) {
	v1, v2, u := ConvertUnitToBase(
		NewDecimalInt(8), YearQuantityUnit, 1,
		NewDecimalInt(4), NewQuantityUnit("", "", ""), 1, true)
	assert.Nil(t, v1, "no value expected")
	assert.Nil(t, v2, "no value expected")
	assert.Nil(t, u, "no unit expected")
}

func TestConvertUnitToBaseLeftUnitNil(t *testing.T) {
	v1, v2, u := ConvertUnitToBase(
		NewDecimalInt(8), nil, 1,
		NewDecimalInt(4), YearQuantityUnit, 1, true)
	assert.Nil(t, v1, "no value expected")
	assert.Nil(t, v2, "no value expected")
	assert.Nil(t, u, "no unit expected")
}

func TestConvertUnitToBaseRightUnitNil(t *testing.T) {
	v1, v2, u := ConvertUnitToBase(
		NewDecimalInt(8), YearQuantityUnit, 1,
		NewDecimalInt(4), nil, 1, true)
	assert.Nil(t, v1, "no value expected")
	assert.Nil(t, v2, "no value expected")
	assert.Nil(t, u, "no unit expected")
}

func TestConvertUnitToMostGranularNoCommonBase(t *testing.T) {
	v1, v2, u := ConvertUnitToMostGranular(
		NewDecimalInt(8), YearQuantityUnit, 1,
		NewDecimalInt(4), NewQuantityUnit("", "", ""), 1, true)
	assert.Nil(t, v1, "no value expected")
	assert.Nil(t, v2, "no value expected")
	assert.Nil(t, u, "no unit expected")
}

func TestConvertUnitToMostGranularLeftUnitNil(t *testing.T) {
	v1, v2, u := ConvertUnitToMostGranular(
		NewDecimalInt(8), nil, 1,
		NewDecimalInt(4), YearQuantityUnit, 1, true)
	assert.Nil(t, v1, "no value expected")
	assert.Nil(t, v2, "no value expected")
	assert.Nil(t, u, "no unit expected")
}

func TestConvertUnitToMostGranularRightUnitNil(t *testing.T) {
	v1, v2, u := ConvertUnitToMostGranular(
		NewDecimalInt(8), YearQuantityUnit, 1,
		NewDecimalInt(4), nil, 1, true)
	assert.Nil(t, v1, "no value expected")
	assert.Nil(t, v2, "no value expected")
	assert.Nil(t, u, "no unit expected")
}

func TestConvertUnitToMostGranularLeft(t *testing.T) {
	v1, v2, u := ConvertUnitToMostGranular(
		NewDecimalInt(15), DayQuantityUnit, 1,
		NewDecimalInt(4), WeekQuantityUnit, 1, true)
	if assert.Same(t, DayQuantityUnit, u) {
		assert.Equal(t, 15.0, v1.Float64())
		assert.Equal(t, 28.0, v2.Float64())
	}
}

func TestConvertUnitToMostGranularRight(t *testing.T) {
	v1, v2, u := ConvertUnitToMostGranular(
		NewDecimalInt(2), YearQuantityUnit, 1,
		NewDecimalInt(5), MonthQuantityUnit, 1, true)
	if assert.Same(t, MonthQuantityUnit, u) {
		assert.Equal(t, 24.0, v1.Float64())
		assert.Equal(t, 5.0, v2.Float64())
	}
}

func TestQuantityUnitNameSingular(t *testing.T) {
	assert.Equal(t, NewString("second"), SecondQuantityUnit.Name(NewDecimalInt(1)))
}

func TestQuantityUnitNamePlural(t *testing.T) {
	assert.Equal(t, NewString("seconds"), SecondQuantityUnit.Name(NewDecimalInt(2)))
}

func TestQuantityUnitNameUCUM(t *testing.T) {
	assert.Equal(t, NewString("abc"),
		NewQuantityUnit("", "", "abc").Name(NewDecimalInt(2)))
}

func TestQuantityUnitNameWithExpEmpty(t *testing.T) {
	assert.Nil(t, EmptyQuantityUnit.NameWithExp(NewDecimalInt(2), 2))
}

func TestQuantityUnitNameWithExp1(t *testing.T) {
	assert.Equal(t, NewString("years"), YearQuantityUnit.NameWithExp(NewDecimalInt(10), 1))
}

func TestQuantityUnitNameWithExp2(t *testing.T) {
	assert.Equal(t, NewString("years2"), YearQuantityUnit.NameWithExp(NewDecimalInt(10), 2))
}

func TestQuantityUnitNameWithExpSingular(t *testing.T) {
	assert.Equal(t, NewString("year"), YearQuantityUnit.NameWithExp(NewDecimalInt(1), 1))
}

func TestNewQuantityUnitEqualNil(t *testing.T) {
	qu1 := NewQuantityUnitWithUCUM("abc")
	assert.Equal(t, false, qu1.Equal(nil))
}

func TestNewQuantityUnitEqualUCUM(t *testing.T) {
	qu1 := NewQuantityUnitWithUCUM("abc")
	qu2 := NewQuantityUnitWithUCUM("abc")
	assert.Equal(t, true, qu1.Equal(qu2))
}

func TestNewQuantityUnitEqualNotUCUM(t *testing.T) {
	qu1 := NewQuantityUnitWithUCUM("abc")
	qu2 := NewQuantityUnitWithUCUM("abcd")
	assert.Equal(t, false, qu1.Equal(qu2))
}

func TestNewQuantityUnitEqualSingularPlural(t *testing.T) {
	qu1 := NewQuantityUnit("test", "tests", "")
	qu2 := NewQuantityUnit("test", "tests", "")
	assert.Equal(t, true, qu1.Equal(qu2))
}

func TestNewQuantityUnitEqualNotSingular(t *testing.T) {
	qu1 := NewQuantityUnit("test", "tests", "")
	qu2 := NewQuantityUnit("testx", "tests", "")
	assert.Equal(t, false, qu1.Equal(qu2))
}

func TestNewQuantityUnitEqualNotPlural(t *testing.T) {
	qu1 := NewQuantityUnit("test", "tests", "")
	qu2 := NewQuantityUnit("test", "testsx", "")
	assert.Equal(t, false, qu1.Equal(qu2))
}

var timeUnitEquivalenceTests = []struct {
	name        string
	ucumUnit    QuantityUnitAccessor
	nonUcumUnit QuantityUnitAccessor
	factor      float64
}{
	{"minute", UCUMMinuteQuantityUnit, MinuteQuantityUnit, 60},
	{"hour", UCUMHourQuantityUnit, HourQuantityUnit, 60 * 60},
	{"day", UCUMDayQuantityUnit, DayQuantityUnit, 24 * 60 * 60},
	{"week", UCUMWeekQuantityUnit, WeekQuantityUnit, 7 * 24 * 60 * 60},
	{"month", UCUMMonthQuantityUnit, MonthQuantityUnit, 30 * 24 * 60 * 60},
	{"year", UCUMYearQuantityUnit, YearQuantityUnit, 365 * 24 * 60 * 60},
}

func TestFunctions(t *testing.T) {
	for _, tt := range timeUnitEquivalenceTests {
		t.Run(tt.name, func(t *testing.T) {
			ucum := QuantityUnitByNameString(tt.ucumUnit.UCUM())
			assert.NotNil(t, ucum, "UCUM unit could not be looked up: %s", tt.ucumUnit.UCUM())

			v1, v2, u := ConvertUnitToBase(
				NewDecimalInt(10), tt.ucumUnit, 1,
				NewDecimalInt(10), tt.nonUcumUnit, 1, true)
			assert.Nil(t, v1, "no value expected")
			assert.Nil(t, v2, "no value expected")
			assert.Nil(t, u, "no unit expected")

			v1, v2, u = ConvertUnitToBase(
				NewDecimalInt(10), tt.ucumUnit, 1,
				NewDecimalInt(10), tt.nonUcumUnit, 1, false)
			if assert.Same(t, SecondQuantityUnit, u) {
				assert.Equal(t, 10.0*tt.factor, v1.Float64())
				assert.Equal(t, 10.0*tt.factor, v2.Float64())
			}
		})
	}
}
