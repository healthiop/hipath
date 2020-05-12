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
	"github.com/stretchr/testify/assert"
	"github.com/volsch/gohimodel/datatype"
	"testing"
)

func TestQuantityLiteralUCUM(t *testing.T) {
	executor, err := ParseQuantityLiteral("-17.4", "ms")

	assert.NoError(t, err, "no error expected")
	assert.NotNil(t, executor, "executor expected")
	if assert.NotNil(t, executor, "executor expected") {
		accessor := executor.Execute(nil)
		assert.NotNil(t, accessor, "accessor expected")
		if assert.Implements(t, (*datatype.QuantityAccessor)(nil), accessor) {
			quantityAccessor := accessor.(datatype.QuantityAccessor)
			if assert.NotNil(t, quantityAccessor.Value()) {
				assert.Equal(t, -17.4, quantityAccessor.Value().Value())
			}
			if assert.NotNil(t, quantityAccessor.System()) {
				assert.Equal(t, "http://unitsofmeasure.org", quantityAccessor.System().Value())
			}
			if assert.NotNil(t, quantityAccessor.Code()) {
				assert.Equal(t, "ms", quantityAccessor.Code().Value())
			}
		}
	}
}

func TestQuantityLiteralNoUCUM(t *testing.T) {
	executor, err := ParseQuantityLiteral("-17.4", "milliseconds")

	assert.NoError(t, err, "no error expected")
	assert.NotNil(t, executor, "executor expected")
	if assert.NotNil(t, executor, "executor expected") {
		accessor := executor.Execute(nil)
		assert.NotNil(t, accessor, "accessor expected")
		if assert.Implements(t, (*datatype.QuantityAccessor)(nil), accessor) {
			quantityAccessor := accessor.(datatype.QuantityAccessor)
			if assert.NotNil(t, quantityAccessor.Value()) {
				assert.Equal(t, -17.4, quantityAccessor.Value().Value())
			}
			assert.Nil(t, quantityAccessor.System())
			if assert.NotNil(t, quantityAccessor.Code()) {
				assert.Equal(t, MillisecondQuantityCode.Value(), quantityAccessor.Code().Value())
			}
		}
	}
}

func TestQuantityLiteralUnitInvalid(t *testing.T) {
	executor, err := ParseQuantityLiteral("-17.4", " test")

	assert.Error(t, err, "error expected")
	assert.Nil(t, executor, "no executor expected")
}

func TestQuantityLiteralUnits(t *testing.T) {
	testParseQuantityUnit(t, "year", YearQuantityCode.Value())
	testParseQuantityUnit(t, "years", YearQuantityCode.Value())
	testParseQuantityUnit(t, "month", MonthQuantityCode.Value())
	testParseQuantityUnit(t, "months", MonthQuantityCode.Value())
	testParseQuantityUnit(t, "week", WeekQuantityCode.Value())
	testParseQuantityUnit(t, "weeks", WeekQuantityCode.Value())
	testParseQuantityUnit(t, "day", DayQuantityCode.Value())
	testParseQuantityUnit(t, "days", DayQuantityCode.Value())
	testParseQuantityUnit(t, "hour", HourQuantityCode.Value())
	testParseQuantityUnit(t, "hours", HourQuantityCode.Value())
	testParseQuantityUnit(t, "minute", MinuteQuantityCode.Value())
	testParseQuantityUnit(t, "minutes", MinuteQuantityCode.Value())
	testParseQuantityUnit(t, "second", SecondQuantityCode.Value())
	testParseQuantityUnit(t, "seconds", SecondQuantityCode.Value())
	testParseQuantityUnit(t, "millisecond", MillisecondQuantityCode.Value())
	testParseQuantityUnit(t, "milliseconds", MillisecondQuantityCode.Value())
}

func testParseQuantityUnit(t *testing.T, unit string, code string) {
	executor, err := ParseQuantityLiteral("-17.4", unit)

	assert.NoError(t, err, "no error expected")
	assert.NotNil(t, executor, "executor expected")
	if assert.NotNil(t, executor, "executor expected") {
		accessor := executor.Execute(nil)
		assert.NotNil(t, accessor, "accessor expected")
		if assert.Implements(t, (*datatype.QuantityAccessor)(nil), accessor) {
			quantityAccessor := accessor.(datatype.QuantityAccessor)
			if assert.NotNil(t, quantityAccessor.Value()) {
				assert.Equal(t, -17.4, quantityAccessor.Value().Value())
			}
			assert.Nil(t, quantityAccessor.System())
			if assert.NotNil(t, quantityAccessor.Code()) {
				assert.Equal(t, code, quantityAccessor.Code().Value())
			}
		}
	}
}
