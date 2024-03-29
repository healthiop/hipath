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
	"regexp"
	"strconv"
	"strings"
	"time"
)

var dateTypeSpec = newAnyTypeSpec("Date")

var dateTypeInfo = NewSimpleTypeInfo(namespaceNameString, NewString("Date"), anyTypeNameString)

var dateRegexp = regexp.MustCompile("^(\\d(?:\\d(?:\\d[1-9]|[1-9]0)|[1-9]00)|[1-9]000)(?:-(0[1-9]|1[0-2])(?:-(0[1-9]|[1-2]\\d|3[0-1]))?)?$")

type dateType struct {
	temporalType
	year  int
	month int
	day   int
}

type DateAccessor interface {
	DateTemporalAccessor
	Year() int
	Month() int
	Day() int
}

func NewDate(value time.Time) DateAccessor {
	return NewDateWithSource(value, nil)
}

func NewDateWithSource(value time.Time, source interface{}) DateAccessor {
	return NewDateYMDWithSource(value.Year(), int(value.Month()), value.Day(), source)
}

func NewDateYMD(year int, month int, day int) DateAccessor {
	return NewDateYMDWithSource(year, month, day, nil)
}

func NewDateYMDWithSource(year int, month int, day int, source interface{}) DateAccessor {
	return newDate(year, month, day, DayDatePrecision, source)
}

func NewDateYMDWithPrecision(year, month, day int, precision DateTimePrecisions) DateAccessor {
	return NewDateYMDWithPrecisionAndSource(year, month, day, precision, nil)
}

func NewDateYMDWithPrecisionAndSource(year, month, day int, precision DateTimePrecisions, source interface{}) DateAccessor {
	if precision < YearDatePrecision || precision > DayDatePrecision {
		panic(fmt.Sprintf("invalid date precision %d", precision))
	}

	if precision < MonthDatePrecision {
		month = 1
	}
	if precision < DayDatePrecision {
		day = 1
	}

	return newDate(year, month, day, precision, source)
}

func ParseDate(value string) (DateAccessor, error) {
	return ParseDateWithSource(value, nil)
}

func ParseDateWithSource(value string, source interface{}) (DateAccessor, error) {
	parts := dateRegexp.FindStringSubmatch(value)
	if parts == nil {
		return nil, fmt.Errorf("not a valid date string: %s", value)
	}
	return newDateFromParts(parts, source), nil
}

func newDateFromParts(parts []string, source interface{}) DateAccessor {
	year, _ := strconv.Atoi(parts[1])
	precision := YearDatePrecision

	month := 1
	if parts[2] != "" {
		month, _ = strconv.Atoi(parts[2])
		precision = MonthDatePrecision
	}

	day := 1
	if parts[3] != "" {
		day, _ = strconv.Atoi(parts[3])
		precision = DayDatePrecision
	}

	return newDate(year, month, day, precision, source)
}

func newDate(year int, month int, day int, precision DateTimePrecisions, source interface{}) *dateType {
	return &dateType{
		temporalType: temporalType{
			baseAnyType: baseAnyType{
				source: source,
			},
			precision: precision,
		},
		year:  year,
		month: month,
		day:   day,
	}
}

func (t *dateType) DataType() DataTypes {
	return DateDataType
}

func (t *dateType) LowestPrecision() DateTimePrecisions {
	return YearDatePrecision
}

func (t *dateType) Date() DateAccessor {
	return t
}

func (t *dateType) DateTime() DateTimeAccessor {
	return NewDateTimeYMDHMSNWithPrecision(t.year, t.month, t.day, 0, 0, 0, 0, time.Local, t.precision)
}

func (t *dateType) Year() int {
	return t.year
}

func (t *dateType) Month() int {
	return t.month
}

func (t *dateType) Day() int {
	return t.day
}

func (t *dateType) Time() time.Time {
	return time.Date(t.year, time.Month(t.month), t.day, 0, 0, 0, 0, time.Local)
}

func (t *dateType) TypeSpec() TypeSpecAccessor {
	return dateTypeSpec
}

func (e *dateType) TypeInfo() TypeInfoAccessor {
	return dateTypeInfo
}

func (t *dateType) Equal(node interface{}) bool {
	if o, ok := node.(DateTemporalAccessor); !ok {
		return false
	} else {
		return TemporalPrecisionEqual(t, o) && dateValueEqual(t, o)
	}
}

func (t *dateType) Equivalent(node interface{}) bool {
	if o, ok := node.(DateTemporalAccessor); !ok {
		return false
	} else {
		return t.Time().Equal(o.Time())
	}
}

func dateValueEqual(dt1 DateTemporalAccessor, dt2 DateTemporalAccessor) bool {
	return dt1.Year() == dt2.Year() && dt1.Month() == dt2.Month() && dt1.Day() == dt2.Day()
}

func (t *dateType) Compare(comparator Comparator) (int, OperatorStatus) {
	if !TypeEqual(t, comparator) {
		return -1, Inconvertible
	}

	o := comparator.(DateAccessor)
	if !TemporalPrecisionEqual(t, o) {
		return -1, Empty
	}

	v := compareDateTimeValue(t.year, o.Year())
	if v != 0 {
		return v, Evaluated
	}
	v = compareDateTimeValue(t.month, o.Month())
	if v != 0 {
		return v, Evaluated
	}
	v = compareDateTimeValue(t.day, o.Day())
	if v != 0 {
		return v, Evaluated
	}

	return 0, Evaluated
}

func (t *dateType) Add(quantity QuantityAccessor) (TemporalAccessor, error) {
	value, precision, err := quantityDateTimePrecision(quantity)
	if err != nil {
		return nil, err
	}

	if precision > DayDatePrecision {
		return nil, fmt.Errorf("quantity precision not allowd for date type: %s", quantity.String())
	}

	res, err := addQuantityTemporalDuration(t, value, precision)
	if err != nil {
		return nil, err
	}
	return NewDateYMDWithPrecision(res.Year(), int(res.Month()), res.Day(), t.precision), nil
}

func (t *dateType) String() string {
	var b strings.Builder
	b.Grow(10)

	writeStringBuilderInt(&b, t.year, 4)
	if t.precision >= MonthDatePrecision {
		b.WriteByte('-')
		writeStringBuilderInt(&b, int(t.month), 2)
	}
	if t.precision >= DayDatePrecision {
		b.WriteByte('-')
		writeStringBuilderInt(&b, t.day, 2)
	}

	return b.String()
}
