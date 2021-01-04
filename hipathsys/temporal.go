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
	"fmt"
	"time"
)

type DateTimePrecisions int

const (
	YearDatePrecision DateTimePrecisions = iota
	MonthDatePrecision
	DayDatePrecision
	HourTimePrecision
	MinuteTimePrecision
	SecondTimePrecision
	NanoTimePrecision
)

var milliNanosecondFactor = NewDecimalInt(1_000_000)
var weekDayFactor = NewDecimalInt(7)
var secondNanosecondFactor = NewDecimalInt(1_000_000_000)
var minuteNanosecondFactor = NewDecimalInt64(60 * 1_000_000_000)
var hourNanosecondFactor = NewDecimalInt64(60 * 60 * 1_000_000_000)
var dayNanosecondFactor = NewDecimalInt64(24 * 60 * 60 * 1_000_000_000)
var monthNanosecondFactor = NewDecimalInt64(30 * 24 * 60 * 60 * 1_000_000_000)
var yearNanosecondFactor = NewDecimalInt64(365 * 24 * 60 * 60 * 1_000_000_000)

type temporalType struct {
	baseAnyType
	precision DateTimePrecisions
}

type TemporalAccessor interface {
	AnyAccessor
	Comparator
	Stringifier
	Precision() DateTimePrecisions
	LowestPrecision() DateTimePrecisions
	Add(quantity QuantityAccessor) (TemporalAccessor, error)
}

func TemporalPrecisionEqual(t1 TemporalAccessor, t2 TemporalAccessor) bool {
	return t1.Precision() == t2.Precision() ||
		(t1.Precision() >= SecondTimePrecision && t2.Precision() >= SecondTimePrecision)
}

type DateTemporalAccessor interface {
	TemporalAccessor
	Time() time.Time
	Year() int
	Month() int
	Day() int
	Date() DateAccessor
	DateTime() DateTimeAccessor
}

func (t *temporalType) Precision() DateTimePrecisions {
	return t.precision
}

func compareDateTimeValue(left, right int) int {
	if left > right {
		return 1
	}
	if left < right {
		return -1
	}
	return 0
}

func addQuantityTemporalDuration(temporal DateTemporalAccessor, quantityValue NumberAccessor,
	quantityPrecision DateTimePrecisions) (time.Time, error) {
	return addQuantityDateTimeDuration(temporal.Time(), temporal.Precision(), quantityValue, quantityPrecision)
}

func addQuantityDateTimeDuration(t time.Time, precision DateTimePrecisions,
	quantityValue NumberAccessor, quantityPrecision DateTimePrecisions) (time.Time, error) {
	if precision < quantityPrecision {
		nanos := quantityValueNanos(quantityValue, quantityPrecision)

		var res DecimalValueAccessor
		switch precision {
		case YearDatePrecision:
			res, _ = nanos.Calc(yearNanosecondFactor, DivisionOp)
		case MonthDatePrecision:
			res, _ = nanos.Calc(monthNanosecondFactor, DivisionOp)
		case DayDatePrecision:
			res, _ = nanos.Calc(dayNanosecondFactor, DivisionOp)
		case HourTimePrecision:
			res, _ = nanos.Calc(hourNanosecondFactor, DivisionOp)
		case MinuteTimePrecision:
			res, _ = nanos.Calc(minuteNanosecondFactor, DivisionOp)
		case SecondTimePrecision:
			res, _ = nanos.Calc(secondNanosecondFactor, DivisionOp)
		default:
			panic(fmt.Sprintf("invalid date/time precision: %d", precision))
		}

		quantityValue = res.Value().Truncate(0)
		quantityPrecision = precision
	} else if quantityPrecision < SecondTimePrecision {
		quantityValue = quantityValue.Truncate(0)
	}

	switch quantityPrecision {
	case YearDatePrecision:
		t = t.AddDate(int(quantityValue.Int()), 0, 0)
	case MonthDatePrecision:
		t = t.AddDate(0, int(quantityValue.Int()), 0)
	case DayDatePrecision:
		t = t.AddDate(0, 0, int(quantityValue.Int()))
	case HourTimePrecision:
		t = t.Add(time.Duration(quantityValue.Int64()) * time.Hour)
	case MinuteTimePrecision:
		t = t.Add(time.Duration(quantityValue.Int64()) * time.Minute)
	case SecondTimePrecision:
		t = t.Add(time.Duration(quantityValue.Float64() * float64(time.Second)))
	case NanoTimePrecision:
		t = t.Add(time.Duration(quantityValue.Int64()))
	default:
		panic(fmt.Sprintf("invalid date/time precision: %d", quantityPrecision))
	}

	year := t.Year()
	if year < 0 || year > 9999 {
		return time.Time{}, fmt.Errorf("date/time arithmetic results in invalid year: %d", year)
	}

	return t, nil
}

func quantityValueNanos(value NumberAccessor, precision DateTimePrecisions) NumberAccessor {
	var d DecimalValueAccessor

	switch precision {
	case MonthDatePrecision:
		d, _ = value.Calc(monthNanosecondFactor, MultiplicationOp)
	case DayDatePrecision:
		d, _ = value.Calc(dayNanosecondFactor, MultiplicationOp)
	case HourTimePrecision:
		d, _ = value.Calc(hourNanosecondFactor, MultiplicationOp)
	case MinuteTimePrecision:
		d, _ = value.Calc(minuteNanosecondFactor, MultiplicationOp)
	case SecondTimePrecision:
		d, _ = value.Calc(secondNanosecondFactor, MultiplicationOp)
	case NanoTimePrecision:
		d = value
	default:
		panic(fmt.Sprintf("invalid date/time precision: %d", precision))
	}

	return d.Value()
}

func quantityDateTimePrecision(q QuantityAccessor) (NumberAccessor, DateTimePrecisions, error) {
	value, unit := q.Value(), q.Unit()
	if unit == nil {
		return nil, NanoTimePrecision, fmt.Errorf("quantity has no date/time unit")
	}

	qu := QuantityUnitByNameString(unit)
	if qu == nil {
		return nil, NanoTimePrecision, fmt.Errorf(
			"quantity has no valid date/time unit: %s", unit.String())
	}

	switch qu {
	case YearQuantityUnit:
		return value, YearDatePrecision, nil
	case MonthQuantityUnit:
		return value, MonthDatePrecision, nil
	case WeekQuantityUnit:
		v, _ := value.Calc(weekDayFactor, MultiplicationOp)
		return v.Value(), DayDatePrecision, nil
	case DayQuantityUnit:
		return value, DayDatePrecision, nil
	case HourQuantityUnit:
		return value, HourTimePrecision, nil
	case MinuteQuantityUnit:
		return value, MinuteTimePrecision, nil
	case SecondQuantityUnit:
		return value, SecondTimePrecision, nil
	case MillisecondQuantityUnit:
		v, _ := value.Calc(milliNanosecondFactor, MultiplicationOp)
		return v.Value(), NanoTimePrecision, nil
	case NanosecondQuantityUnit:
		return value, NanoTimePrecision, nil
	default:
		return nil, NanoTimePrecision, fmt.Errorf(
			"quantity has no valid date/time unit: %s", unit.String())
	}
}
