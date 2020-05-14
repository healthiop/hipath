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

package internal

import (
	"github.com/stretchr/testify/assert"
	"github.com/volsch/gohimodel/datatype"
	"github.com/volsch/gohipath/internal/expression"
	"testing"
	"time"
)

func TestParseNullLiteral(t *testing.T) {
	result, errorItemCollection := testParse("{}")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	assert.IsType(t, (*expression.NullLiteral)(nil), result)
}

func TestParseBooleanLiteral(t *testing.T) {
	result, errorItemCollection := testParse("true")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.BooleanLiteral)(nil), result) {
		a, _ := result.(expression.Evaluator).Evaluate(nil, nil)
		assert.Equal(t, true, a.(datatype.BooleanAccessor).Value())
	}
}

func TestParseStringLiteral(t *testing.T) {
	result, errorItemCollection := testParse("'Test \\nValue'")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.StringLiteral)(nil), result) {
		a, _ := result.(expression.Evaluator).Evaluate(nil, nil)
		assert.Equal(t, "Test \nValue", a.(datatype.StringAccessor).Value())
	}
}

func TestParseNumberLiteral(t *testing.T) {
	result, errorItemCollection := testParse("183.2889")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.NumberLiteral)(nil), result) {
		a, _ := result.(expression.Evaluator).Evaluate(nil, nil)
		assert.Equal(t, 183.2889, a.(datatype.NumberAccessor).Float64())
	}
}

func TestParseDateTimeLiteral(t *testing.T) {
	result, errorItemCollection := testParse("@2014-05-25T14:30:14.559Z")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.DateTimeLiteral)(nil), result) {
		a, _ := result.(expression.Evaluator).Evaluate(nil, nil)
		assert.Equal(t, time.Date(2014, 5, 25, 14, 30, 14, 559000000, time.UTC),
			a.(datatype.DateTimeAccessor).Value())
	}
}

func TestParseDateLiteral(t *testing.T) {
	result, errorItemCollection := testParse("@2014-05-25")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.DateLiteral)(nil), result) {
		a, _ := result.(expression.Evaluator).Evaluate(nil, nil)
		assert.Equal(t, time.Date(2014, 5, 25, 0, 0, 0, 0, time.Local),
			a.(datatype.DateAccessor).Value())
	}
}

func TestParseTimeLiteral(t *testing.T) {
	result, errorItemCollection := testParse("@T14:30:14.559")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.TimeLiteral)(nil), result) {
		a, _ := result.(expression.Evaluator).Evaluate(nil, nil)
		now := time.Now()
		assert.Equal(t, time.Date(now.Year(), now.Month(), now.Day(), 14, 30, 14, 559000000, time.Local),
			a.(datatype.TimeAccessor).Value())
	}
}

func TestParseQuantityLiteral(t *testing.T) {
	result, errorItemCollection := testParse("736.2321 years")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.QuantityLiteral)(nil), result) {
		a, _ := result.(expression.Evaluator).Evaluate(nil, nil)
		qa := a.(datatype.QuantityAccessor)
		if assert.NotNil(t, qa.Value(), "quantity value expected") {
			assert.Equal(t, 736.2321, qa.Value().Float64())
		}
		assert.Nil(t, qa.System(), "no quantity unit system expected")
		if assert.NotNil(t, qa.Code(), "quantity code expected") {
			assert.Equal(t, "year", qa.Code().Value())
		}
	}
}

func TestParseQuantityLiteralUCUM(t *testing.T) {
	result, errorItemCollection := testParse("736.2321 'cm'")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.False(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	if assert.IsType(t, (*expression.QuantityLiteral)(nil), result) {
		a, _ := result.(expression.Evaluator).Evaluate(nil, nil)
		qa := a.(datatype.QuantityAccessor)
		if assert.NotNil(t, qa.Value(), "quantity value expected") {
			assert.Equal(t, 736.2321, qa.Value().Float64())
		}
		if assert.NotNil(t, qa.System(), "quantity unit system expected") {
			assert.Equal(t, datatype.UCUMSystemURI, qa.System())
		}
		if assert.NotNil(t, qa.Code(), "quantity code expected") {
			assert.Equal(t, "cm", qa.Code().Value())
		}
	}
}

func TestParseQuantityLiteralUnitInvalid(t *testing.T) {
	result, errorItemCollection := testParse("736.2321 ' cm'")

	if assert.NotNil(t, errorItemCollection, "error item collection must have been initialized") {
		assert.True(t, errorItemCollection.HasErrors(), "no errors expected")
	}
	assert.Nil(t, result, "no result expected")
}
