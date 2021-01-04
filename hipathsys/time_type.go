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
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var timeTypeSpec = newAnyTypeSpec("Time")

var timeRegexp = regexp.MustCompile("^([01]\\d|2[0-3])(?::([0-5]\\d)(?::([0-5]\\d|60)(?:\\.(\\d+))?)?)?$")

type timeType struct {
	temporalType
	hour       int
	minute     int
	second     int
	nanosecond int
}

type TimeAccessor interface {
	TemporalAccessor
	Hour() int
	Minute() int
	Second() int
	Nanosecond() int
}

func NewTime(value time.Time) TimeAccessor {
	return NewTimeWithSource(value, nil)
}

func NewTimeWithSource(value time.Time, source interface{}) TimeAccessor {
	return NewTimeHMSNWithSource(value.Hour(), value.Minute(), value.Second(), value.Nanosecond(), source)
}

func NewTimeHMSN(hour int, minute int, second int, nanosecond int) TimeAccessor {
	return NewTimeHMSNWithSource(hour, minute, second, nanosecond, nil)
}

func NewTimeHMSNWithSource(hour int, minute int, second int, nanosecond int, source interface{}) TimeAccessor {
	return newTime(hour, minute, second, nanosecond, NanoTimePrecision, source)
}

func NewTimeHMSNWithPrecision(hour, minute, second, nanosecond int, precision DateTimePrecisions) TimeAccessor {
	return NewTimeHMSNWithPrecisionAndSource(hour, minute, second, nanosecond, precision, nil)
}

func NewTimeHMSNWithPrecisionAndSource(hour, minute, second int, nanosecond int, precision DateTimePrecisions, source interface{}) TimeAccessor {
	if precision < HourTimePrecision || precision > NanoTimePrecision {
		panic(fmt.Sprintf("invalid time precision %d", precision))
	}

	if precision < MinuteTimePrecision {
		minute = 0
	}
	if precision < SecondTimePrecision {
		second = 0
	}
	if precision < NanoTimePrecision {
		nanosecond = 0
	}

	return newTime(hour, minute, second, nanosecond, precision, source)
}

func ParseTime(value string) (TimeAccessor, error) {
	parts := timeRegexp.FindStringSubmatch(value)
	if parts == nil {
		return nil, fmt.Errorf("not a valid fluent time string: %s", value)
	}
	return newTimeFromParts(parts), nil
}

func newTimeFromParts(parts []string) TimeAccessor {
	hour, _ := strconv.Atoi(parts[1])
	precision := HourTimePrecision

	minute := 0
	if parts[2] != "" {
		minute, _ = strconv.Atoi(parts[2])
		precision = MinuteTimePrecision
	}

	second := 0
	if parts[3] != "" {
		second, _ = strconv.Atoi(parts[3])
		precision = SecondTimePrecision
	}

	nanosecond := 0
	if parts[4] != "" {
		nanosecond = parseNanosecond(parts[4])
		precision = NanoTimePrecision
	}

	return newTime(hour, minute, second, nanosecond, precision, nil)
}

func newTime(hour int, minute int, second int, nanosecond int, precision DateTimePrecisions, source interface{}) TimeAccessor {
	return &timeType{
		temporalType: temporalType{
			baseAnyType: baseAnyType{
				source: source,
			},
			precision: precision,
		},
		hour:       hour,
		minute:     minute,
		second:     second,
		nanosecond: nanosecond,
	}
}

func parseNanosecond(value string) int {
	if value == "" {
		return 0
	}
	nanoValue := value
	if len(nanoValue) > 9 {
		nanoValue = nanoValue[0:9]
	}
	nano, _ := strconv.Atoi(nanoValue)
	nano = nano * int(math.Pow10(9-len(nanoValue)))
	return nano
}

func (t *timeType) DataType() DataTypes {
	return TimeDataType
}

func (t *timeType) LowestPrecision() DateTimePrecisions {
	return HourTimePrecision
}

func (t *timeType) Hour() int {
	return t.hour
}

func (t *timeType) Minute() int {
	return t.minute
}

func (t *timeType) Second() int {
	return t.second
}

func (t *timeType) Nanosecond() int {
	return t.nanosecond
}

func (t *timeType) TypeSpec() TypeSpecAccessor {
	return timeTypeSpec
}

func (t *timeType) Equal(node interface{}) bool {
	if !SystemAnyTypeEqual(t, node) {
		return false
	}

	o := node.(TimeAccessor)
	return TemporalPrecisionEqual(t, o) &&
		t.Hour() == o.Hour() &&
		t.Minute() == o.Minute() &&
		t.Second() == o.Second() &&
		t.Nanosecond() == o.Nanosecond()
}

func (t *timeType) Equivalent(node interface{}) bool {
	if !SystemAnyTypeEqual(t, node) {
		return false
	}
	return timeValueEqual(t, node.(TimeAccessor))
}

func timeValueEqual(t1 TimeAccessor, t2 TimeAccessor) bool {
	return t1.Hour() == t2.Hour() && t1.Minute() == t2.Minute() && t1.Second() == t2.Second() &&
		t1.Nanosecond() == t2.Nanosecond()
}

func (t *timeType) Compare(comparator Comparator) (int, OperatorStatus) {
	if !TypeEqual(t, comparator) {
		return -1, Inconvertible
	}

	o := comparator.(TimeAccessor)
	if !TemporalPrecisionEqual(t, o) {
		return -1, Empty
	}

	v := compareDateTimeValue(t.hour, o.Hour())
	if v != 0 {
		return v, Evaluated
	}
	v = compareDateTimeValue(t.minute, o.Minute())
	if v != 0 {
		return v, Evaluated
	}
	v = compareDateTimeValue(t.second, o.Second())
	if v != 0 {
		return v, Evaluated
	}
	v = compareDateTimeValue(t.nanosecond, o.Nanosecond())
	if v != 0 {
		return v, Evaluated
	}

	return 0, Evaluated
}

func (t *timeType) Add(quantity QuantityAccessor) (TemporalAccessor, error) {
	value, precision, err := quantityDateTimePrecision(quantity)
	if err != nil {
		return nil, err
	}

	if precision < HourTimePrecision {
		return nil, fmt.Errorf("quantity precision not allowd for time type: %s", quantity.String())
	}

	d := time.Date(1, 1, 1, t.hour, t.minute, t.second, t.nanosecond, time.UTC)
	res, err := addQuantityDateTimeDuration(d, t.precision, value, precision)
	if err != nil {
		return nil, err
	}
	if res.Day() != 1 || res.Month() != 1 || res.Year() != 1 {
		return nil, fmt.Errorf("time arithmetic results in an invalid hour")
	}

	return NewTimeHMSNWithPrecision(res.Hour(), res.Minute(), res.Second(), res.Nanosecond(), t.precision), nil
}

func (t *timeType) String() string {
	var b strings.Builder
	b.Grow(19)

	writeStringBuilderInt(&b, t.hour, 2)
	if t.precision >= MinuteTimePrecision {
		b.WriteByte(':')
		writeStringBuilderInt(&b, t.minute, 2)
	}
	if t.precision >= SecondTimePrecision {
		b.WriteByte(':')
		writeStringBuilderInt(&b, t.second, 2)
	}
	if t.precision >= NanoTimePrecision {
		b.WriteByte('.')
		writeStringBuilderInt(&b, t.nanosecond, 9)
	}

	return b.String()
}
