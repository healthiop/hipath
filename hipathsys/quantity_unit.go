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
	"math"
	"regexp"
	"strconv"
)

var (
	DefaultQuantityUnit = NewQuantityUnitWithUCUM("1")

	EmptyQuantityUnit = NewQuantityUnitWithUCUM("")

	SecondQuantityUnit = NewQuantityUnit("second", "seconds", "s")
	MinuteQuantityUnit = NewQuantityUnit("minute", "minutes", "",
		NewQuantityUnitBase(SecondQuantityUnit, true, 60))
	HourQuantityUnit = NewQuantityUnit("hour", "hours", "",
		NewQuantityUnitBase(SecondQuantityUnit, true, 60*60))
	DayQuantityUnit = NewQuantityUnit("day", "days", "",
		NewQuantityUnitBase(SecondQuantityUnit, true, 24*60*60))
	WeekQuantityUnit = NewQuantityUnit("week", "weeks", "",
		NewQuantityUnitBase(SecondQuantityUnit, true, 7*24*60*60))
	MonthQuantityUnit = NewQuantityUnit("month", "months", "",
		NewQuantityUnitBase(SecondQuantityUnit, true, 30*24*60*60))
	YearQuantityUnit = NewQuantityUnit("year", "years", "",
		NewQuantityUnitBase(SecondQuantityUnit, true, 365*24*60*60),
		NewQuantityUnitBase(MonthQuantityUnit, true, 12))
	MillisecondQuantityUnit = NewQuantityUnit("millisecond", "milliseconds", "ms",
		NewQuantityUnitBase(SecondQuantityUnit, true, .001))
	NanosecondQuantityUnit = NewQuantityUnit("nanosecond", "nanoseconds", "ns",
		NewQuantityUnitBase(SecondQuantityUnit, true, .000000001))
	UCUMMinuteQuantityUnit = NewQuantityUnit("", "", "min",
		NewQuantityUnitBase(SecondQuantityUnit, false, 60))
	UCUMHourQuantityUnit = NewQuantityUnit("", "", "h",
		NewQuantityUnitBase(SecondQuantityUnit, false, 60*60))
	UCUMDayQuantityUnit = NewQuantityUnit("", "", "d",
		NewQuantityUnitBase(SecondQuantityUnit, false, 24*60*60))
	UCUMWeekQuantityUnit = NewQuantityUnit("", "", "wk",
		NewQuantityUnitBase(SecondQuantityUnit, false, 7*24*60*60))
	UCUMMonthQuantityUnit = NewQuantityUnit("", "", "mo",
		NewQuantityUnitBase(SecondQuantityUnit, false, 30*24*60*60))
	UCUMYearQuantityUnit = NewQuantityUnit("", "", "a",
		NewQuantityUnitBase(SecondQuantityUnit, false, 365*24*60*60))
)

var quantityUnitsByName = toQuantityUnitsByName(
	NanosecondQuantityUnit,
	MillisecondQuantityUnit,
	SecondQuantityUnit,
	MinuteQuantityUnit,
	HourQuantityUnit,
	DayQuantityUnit,
	WeekQuantityUnit,
	MonthQuantityUnit,
	YearQuantityUnit,
	UCUMMinuteQuantityUnit,
	UCUMHourQuantityUnit,
	UCUMDayQuantityUnit,
	UCUMWeekQuantityUnit,
	UCUMMonthQuantityUnit,
	UCUMYearQuantityUnit)

var quantityUnitExpRegexp = regexp.MustCompile("^(.*[^\\d])([1-3])$")

func IsCalendarDurationUnit(unit QuantityUnitAccessor) bool {
	return unit == SecondQuantityUnit || unit.HasBase(SecondQuantityUnit, true)
}

func QuantityUnitByName(name string) QuantityUnitAccessor {
	if len(name) == 0 {
		return nil
	}
	return quantityUnitsByName[name]
}

func QuantityUnitByNameString(name StringAccessor) QuantityUnitAccessor {
	if name == nil {
		return nil
	}
	return quantityUnitsByName[name.String()]
}

type quantityUnitBase struct {
	unit          QuantityUnitAccessor
	equal         bool
	factor        float64
	decimalFactor DecimalAccessor
}

type QuantityUnitBaseAccessor interface {
	Unit() QuantityUnitAccessor
	Equal() bool
	Factor() float64
	DecimalFactor() DecimalAccessor
}

type quantityUnit struct {
	singular StringAccessor
	plural   StringAccessor
	ucum     StringAccessor
	rootBase QuantityUnitBaseAccessor
	bases    []QuantityUnitBaseAccessor
}

type QuantityUnitAccessor interface {
	Singular() StringAccessor
	Plural() StringAccessor
	UCUM() StringAccessor
	Equal(unit QuantityUnitAccessor) bool
	Name(value DecimalAccessor) StringAccessor
	NameWithExp(value DecimalAccessor, exp int) StringAccessor
	RootBase() QuantityUnitBaseAccessor
	HasBase(unit QuantityUnitAccessor, equal bool) bool
	CommonBase(other QuantityUnitAccessor, equal bool) QuantityUnitAccessor
	Factor(other QuantityUnitAccessor, exp int) NumberAccessor
}

func NewQuantityUnitBase(unit QuantityUnitAccessor, equal bool, factor float64) QuantityUnitBaseAccessor {
	return &quantityUnitBase{
		unit:          unit,
		equal:         equal,
		factor:        factor,
		decimalFactor: NewDecimalFloat64(factor),
	}
}

func NewQuantityUnit(singular, plural, ucum string, bases ...QuantityUnitBaseAccessor) QuantityUnitAccessor {
	var rootBase QuantityUnitBaseAccessor
	if bases != nil {
		for _, b := range bases {
			if b.Unit().RootBase() == nil {
				rootBase = b
				break
			}
		}
	}

	return &quantityUnit{
		singular: StringOfNil(singular),
		plural:   StringOfNil(plural),
		ucum:     StringOfNil(ucum),
		rootBase: rootBase,
		bases:    bases,
	}
}

func NewQuantityUnitWithUCUM(ucum string) QuantityUnitAccessor {
	return &quantityUnit{
		ucum: StringOfNil(ucum),
	}
}

func QuantityUnitWithName(name string) (QuantityUnitAccessor, int) {
	l := len(name)
	if l == 0 {
		return nil, 1
	}

	exp := 1
	if l > 1 {
		parts := quantityUnitExpRegexp.FindStringSubmatch(name)
		if parts != nil {
			name = parts[1]
			exp, _ = strconv.Atoi(parts[2])
		}
	}

	unit := QuantityUnitByName(name)
	if unit == nil {
		unit = NewQuantityUnitWithUCUM(name)
	}
	return unit, exp
}

func QuantityUnitWithNameString(name StringAccessor) (QuantityUnitAccessor, int) {
	if name == nil {
		return nil, 1
	}
	return QuantityUnitWithName(name.String())
}

func (q *quantityUnitBase) Unit() QuantityUnitAccessor {
	return q.unit
}

func (q *quantityUnitBase) Equal() bool {
	return q.equal
}

func (q *quantityUnitBase) Factor() float64 {
	return q.factor
}

func (q *quantityUnitBase) DecimalFactor() DecimalAccessor {
	return q.decimalFactor
}

func (q *quantityUnit) Singular() StringAccessor {
	return q.singular
}

func (q *quantityUnit) Plural() StringAccessor {
	return q.plural
}

func (q *quantityUnit) UCUM() StringAccessor {
	return q.ucum
}

func (q *quantityUnit) Equal(unit QuantityUnitAccessor) bool {
	if unit == nil {
		return false
	}
	return (q.UCUM() != nil && Equal(q.UCUM(), unit.UCUM())) ||
		(q.Singular() != nil && Equal(q.Singular(), unit.Singular()) && Equal(q.Plural(), unit.Plural()))
}

func (q *quantityUnit) Name(value DecimalAccessor) StringAccessor {
	if DecimalOne.Equal(value) {
		if q.singular != nil {
			return q.singular
		}
	} else if q.plural != nil {
		return q.plural
	}
	return q.ucum
}

func (q *quantityUnit) NameWithExp(value DecimalAccessor, exp int) StringAccessor {
	name := q.Name(value)
	if name == nil {
		return nil
	}

	if exp == 1 {
		return name
	}

	return NewString(name.String() + strconv.FormatInt(int64(exp), 10))
}

func (q *quantityUnit) RootBase() QuantityUnitBaseAccessor {
	return q.rootBase
}

func (q *quantityUnit) HasBase(unit QuantityUnitAccessor, equal bool) bool {
	if q.bases != nil {
		for _, b := range q.bases {
			if b.Unit() == unit && ((equal && b.Equal()) || !equal) {
				return true
			}
		}
	}
	return false
}

func (q *quantityUnit) CommonBase(other QuantityUnitAccessor, equal bool) QuantityUnitAccessor {
	if other.HasBase(q, equal) {
		return q
	}

	var found QuantityUnitAccessor
	factor := math.MaxFloat64
	if q.bases != nil {
		for _, b := range q.bases {
			if (equal && b.Equal()) || !equal {
				u := b.Unit()
				if (other.HasBase(u, equal) || other == u) && b.Factor() < factor {
					found, factor = b.Unit(), b.Factor()
				}
			}
		}
	}
	return found
}

func (q *quantityUnit) Factor(other QuantityUnitAccessor, exp int) NumberAccessor {
	if q == other {
		return DecimalOne
	}
	if q.bases != nil {
		for _, b := range q.bases {
			if b.Unit() == other {
				if exp == 1 {
					return b.DecimalFactor()
				}
				f, _ := b.DecimalFactor().Power(DecimalOfInt(int32(exp)))
				return f
			}
		}
	}
	return nil
}

func ConvertUnitToBase(v1 DecimalAccessor, u1 QuantityUnitAccessor, exp1 int, v2 DecimalAccessor, u2 QuantityUnitAccessor, exp2 int, equal bool) (DecimalAccessor, DecimalAccessor, QuantityUnitAccessor) {
	if u1 == nil || u2 == nil {
		return nil, nil, nil
	}

	u := u1.CommonBase(u2, equal)
	if u == nil {
		return nil, nil, nil
	}

	f1, f2 := u1.Factor(u, exp1), u2.Factor(u, exp2)
	r1, _ := v1.Calc(f1, MultiplicationOp)
	r2, _ := v2.Calc(f2, MultiplicationOp)
	return r1.Value(), r2.Value(), u
}

func ConvertUnitToMostGranular(v1 DecimalAccessor, u1 QuantityUnitAccessor, exp1 int, v2 DecimalAccessor, u2 QuantityUnitAccessor, exp2 int, equal bool) (DecimalAccessor, DecimalAccessor, QuantityUnitAccessor) {
	if u1 == nil || u2 == nil {
		return nil, nil, nil
	}

	u := u1.CommonBase(u2, equal)
	if u == nil {
		return nil, nil, nil
	}

	var r1, r2 DecimalValueAccessor
	f1, f2 := u1.Factor(u, exp1), u2.Factor(u, exp2)
	if c, _ := f1.Compare(f2); c < 0 {
		r1 = v1
		r2, _ = v2.Calc(f2, MultiplicationOp)
		r2, _ = r2.Value().Calc(f1, DivisionOp)
		u = u1
	} else {
		r1, _ = v1.Calc(f1, MultiplicationOp)
		r1, _ = r1.Value().Calc(f2, DivisionOp)
		r2 = v2
		u = u2
	}
	return r1.Value(), r2.Value(), u
}

func toQuantityUnitsByName(units ...QuantityUnitAccessor) map[string]QuantityUnitAccessor {
	m := make(map[string]QuantityUnitAccessor)
	for _, unit := range units {
		if unit.Singular() != nil {
			m[unit.Singular().String()] = unit
		}
		if unit.Plural() != nil {
			m[unit.Plural().String()] = unit
		}
		if unit.UCUM() != nil {
			m[unit.UCUM().String()] = unit
		}
	}
	return m
}
