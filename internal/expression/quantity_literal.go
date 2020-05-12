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

package expression

import (
	"github.com/volsch/gohimodel/datatype"
	"github.com/volsch/gohipath/context"
)

var (
	YearQuantityCode        datatype.CodeAccessor = datatype.NewCodeType("year")
	MonthQuantityCode       datatype.CodeAccessor = datatype.NewCodeType("month")
	WeekQuantityCode        datatype.CodeAccessor = datatype.NewCodeType("week")
	DayQuantityCode         datatype.CodeAccessor = datatype.NewCodeType("day")
	HourQuantityCode        datatype.CodeAccessor = datatype.NewCodeType("hour")
	MinuteQuantityCode      datatype.CodeAccessor = datatype.NewCodeType("minute")
	SecondQuantityCode      datatype.CodeAccessor = datatype.NewCodeType("second")
	MillisecondQuantityCode datatype.CodeAccessor = datatype.NewCodeType("millisecond")
)

type QuantityLiteral struct {
	accessor datatype.QuantityAccessor
}

func ParseQuantityLiteral(number string, unit string) (Executor, error) {
	value, err := datatype.ParseDecimalValue(number)
	if err != nil {
		return nil, err
	}

	system, code, err := parseQuantityUnit(unit)
	if err != nil {
		return nil, err
	}

	accessor := datatype.NewQuantityType(value, nil, nil, system, code)
	return &QuantityLiteral{accessor}, nil
}

func parseQuantityUnit(unit string) (system datatype.URIAccessor, code datatype.CodeAccessor, err error) {
	if len(unit) == 0 {
		return
	}
	parsedUnit := parseStringLiteral(unit)

	switch parsedUnit {
	case "year", "years":
		code = YearQuantityCode
	case "month", "months":
		code = MonthQuantityCode
	case "week", "weeks":
		code = WeekQuantityCode
	case "day", "days":
		code = DayQuantityCode
	case "hour", "hours":
		code = HourQuantityCode
	case "minute", "minutes":
		code = MinuteQuantityCode
	case "second", "seconds":
		code = SecondQuantityCode
	case "millisecond", "milliseconds":
		code = MillisecondQuantityCode
	default:
		code, err = datatype.ParseCodeValue(parsedUnit)
		if err != nil {
			return
		}
		system = datatype.UCUMSystemURI
	}

	return
}

func (e *QuantityLiteral) Execute(*context.PathContext) interface{} {
	return e.accessor
}
