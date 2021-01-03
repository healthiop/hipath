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

package internal

import (
	"github.com/stretchr/testify/assert"
	"github.com/volsch/gohipath/internal/expression"
	"github.com/volsch/gohipath/pathsys"
	"testing"
	"time"
)

func TestParseNullLiteral(t *testing.T) {
	res, errorItemCollection := testParse("{}")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	assert.IsType(t, (*expression.EmptyLiteral)(nil), res)
}

func TestParseBooleanLiteral(t *testing.T) {
	res, errorItemCollection := testParse("true")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.BooleanLiteral)(nil), res) {
		res, _ := res.(pathsys.Evaluator).Evaluate(nil, nil, nil)
		assert.Equal(t, true, res.(pathsys.BooleanAccessor).Bool())
	}
}

func TestParseStringLiteral(t *testing.T) {
	res, errorItemCollection := testParse("'Test \\nValue'")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.StringLiteral)(nil), res) {
		res, _ := res.(pathsys.Evaluator).Evaluate(nil, nil, nil)
		assert.Equal(t, "Test \nValue", res.(pathsys.StringAccessor).String())
	}
}

func TestParseNumberLiteral(t *testing.T) {
	res, errorItemCollection := testParse("183.2889")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.NumberLiteral)(nil), res) {
		res, _ := res.(pathsys.Evaluator).Evaluate(nil, nil, nil)
		assert.Equal(t, 183.2889, res.(pathsys.NumberAccessor).Float64())
	}
}

func TestParseDateTimeLiteral(t *testing.T) {
	res, errorItemCollection := testParse("@2014-05-25T14:30:14.559Z")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.DateTimeLiteral)(nil), res) {
		res, _ := res.(pathsys.Evaluator).Evaluate(nil, nil, nil)
		assert.Equal(t, time.Date(2014, 5, 25, 14, 30, 14, 559000000, time.UTC),
			res.(pathsys.DateTimeAccessor).Time())
	}
}

func TestParseDateLiteral(t *testing.T) {
	res, errorItemCollection := testParse("@2014-05-25")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.DateLiteral)(nil), res) {
		res, _ := res.(pathsys.Evaluator).Evaluate(nil, nil, nil)
		assert.Equal(t, time.Date(2014, 5, 25, 0, 0, 0, 0, time.Local),
			res.(pathsys.DateAccessor).Time())
	}
}

func TestParseTimeLiteral(t *testing.T) {
	res, errorItemCollection := testParse("@T14:30:17.559")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.TimeLiteral)(nil), res) {
		res, _ := res.(pathsys.Evaluator).Evaluate(nil, nil, nil)
		timeRes := res.(pathsys.TimeAccessor)
		assert.Equal(t, 14, timeRes.Hour())
		assert.Equal(t, 30, timeRes.Minute())
		assert.Equal(t, 17, timeRes.Second())
		assert.Equal(t, 559000000, timeRes.Nanosecond())
	}
}

func TestParseQuantityLiteralPlural(t *testing.T) {
	res, errorItemCollection := testParse("736.2321 years")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.QuantityLiteral)(nil), res) {
		res, _ := res.(pathsys.Evaluator).Evaluate(nil, nil, nil)
		q := res.(pathsys.QuantityAccessor)
		if assert.NotNil(t, q.Value(), "quantity value expected") {
			assert.Equal(t, 736.2321, q.Value().Float64())
		}
		if assert.NotNil(t, q.Unit(), "quantity code expected") {
			assert.Equal(t, "years", q.Unit().String())
		}
	}
}

func TestParseQuantityLiteralSingular(t *testing.T) {
	res, errorItemCollection := testParse("1 year")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.QuantityLiteral)(nil), res) {
		res, _ := res.(pathsys.Evaluator).Evaluate(nil, nil, nil)
		q := res.(pathsys.QuantityAccessor)
		if assert.NotNil(t, q.Value(), "quantity value expected") {
			assert.Equal(t, 1.0, q.Value().Float64())
		}
		if assert.NotNil(t, q.Unit(), "quantity code expected") {
			assert.Equal(t, "year", q.Unit().String())
		}
	}
}

func TestParseQuantityLiteralUCUM(t *testing.T) {
	res, errorItemCollection := testParse("736.2321 'cm'")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.QuantityLiteral)(nil), res) {
		res, _ := res.(pathsys.Evaluator).Evaluate(nil, nil, nil)
		q := res.(pathsys.QuantityAccessor)
		if assert.NotNil(t, q.Value(), "quantity value expected") {
			assert.Equal(t, 736.2321, q.Value().Float64())
		}
		if assert.NotNil(t, q.Unit(), "quantity code expected") {
			assert.Equal(t, "cm", q.Unit().String())
		}
	}
}

func TestParseQuantityLiteralUnitInvalid(t *testing.T) {
	res, errorItemCollection := testParse("736.2321 ' cm'")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.True(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	assert.Nil(t, res, "no res expected")
}
