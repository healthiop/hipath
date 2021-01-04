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

package expression

import (
	"github.com/healthiop/hipath/hipathsys"
	"github.com/healthiop/hipath/internal/test"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestArithmeticExpression(t *testing.T) {
	e := NewArithmeticExpression(NewNumberLiteralInt(87),
		hipathsys.AdditionOp, NewNumberLiteralFloat64(12.43))
	node, err := e.Evaluate(nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.DecimalAccessor)(nil), node) {
		res := node.(hipathsys.DecimalAccessor)
		assert.Equal(t, 99.43, res.Float64())
	}
}

func TestArithmeticExpressionLeftError(t *testing.T) {
	e := NewArithmeticExpression(newTestErrorExpression(),
		hipathsys.AdditionOp, NewNumberLiteralFloat64(12.43))
	node, err := e.Evaluate(nil, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, node, "no res expected")
}

func TestArithmeticExpressionRightError(t *testing.T) {
	e := NewArithmeticExpression(NewNumberLiteralFloat64(12.43),
		hipathsys.AdditionOp, newTestErrorExpression())
	node, err := e.Evaluate(nil, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, node, "no res expected")
}

func TestArithmeticExpressionLeftNil(t *testing.T) {
	e := NewArithmeticExpression(newTestExpression(nil),
		hipathsys.AdditionOp, NewNumberLiteralFloat64(12.43))
	node, err := e.Evaluate(nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, node, "empty res expected")
}

func TestArithmeticExpressionRightNil(t *testing.T) {
	e := NewArithmeticExpression(NewNumberLiteralFloat64(12.43),
		hipathsys.AdditionOp, newTestExpression(nil))
	node, err := e.Evaluate(nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, node, "empty res expected")
}

func TestArithmeticExpressionLeftInvalidType(t *testing.T) {
	e := NewArithmeticExpression(newTestExpression(test.NewTestModelNode(10, false)),
		hipathsys.AdditionOp, NewNumberLiteralFloat64(12.43))
	node, err := e.Evaluate(nil, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, node, "no res expected")
}

func TestArithmeticExpressionRightInvalidType(t *testing.T) {
	e := NewArithmeticExpression(NewNumberLiteralFloat64(12.43),
		hipathsys.AdditionOp, newTestExpression(test.NewTestModelNode(10, false)))
	node, err := e.Evaluate(nil, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, node, "no res expected")
}

func TestArithmeticExpressionTemporalAddition(t *testing.T) {
	date := time.Date(2019, 8, 21, 14, 38, 49, 827362627, time.Local)
	e := NewArithmeticExpression(newTestExpression(hipathsys.NewDateTime(date)), hipathsys.AdditionOp,
		newTestExpression(hipathsys.NewQuantity(hipathsys.NewDecimalInt(12), hipathsys.NewString("day"))))
	res, err := e.Evaluate(nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.DateTimeAccessor)(nil), res) {
		assert.Equal(t, date.Add(12*24*time.Hour).UnixNano(),
			res.(hipathsys.DateTimeAccessor).Time().UnixNano())
	}
}

func TestArithmeticExpressionTemporalSubtraction(t *testing.T) {
	date := time.Date(2019, 8, 21, 14, 38, 49, 827362627, time.Local)
	e := NewArithmeticExpression(newTestExpression(hipathsys.NewDateTime(date)), hipathsys.SubtractionOp,
		newTestExpression(hipathsys.NewQuantity(hipathsys.NewDecimalInt(12), hipathsys.NewString("day"))))
	res, err := e.Evaluate(nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.DateTimeAccessor)(nil), res) {
		assert.Equal(t, date.Add(-12*24*time.Hour).UnixNano(),
			res.(hipathsys.DateTimeAccessor).Time().UnixNano())
	}
}

func TestArithmeticExpressionTemporalWithString(t *testing.T) {
	date := time.Date(2019, 8, 21, 14, 38, 49, 827362627, time.Local)
	e := NewArithmeticExpression(newTestExpression(hipathsys.NewDateTime(date)), hipathsys.SubtractionOp,
		NewRawStringLiteral("10"))
	res, err := e.Evaluate(nil, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestArithmeticExpressionTemporalInvalidUnit(t *testing.T) {
	date := time.Date(2019, 8, 21, 14, 38, 49, 827362627, time.Local)
	e := NewArithmeticExpression(newTestExpression(hipathsys.NewDateTime(date)), hipathsys.SubtractionOp,
		newTestExpression(hipathsys.NewQuantity(hipathsys.NewDecimalInt(12), hipathsys.NewString("x"))))
	res, err := e.Evaluate(nil, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestArithmeticExpressionBothString(t *testing.T) {
	e := NewArithmeticExpression(NewRawStringLiteral("Test1"),
		hipathsys.AdditionOp, NewRawStringLiteral("Test2"))
	res, err := e.Evaluate(nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.StringAccessor)(nil), res) {
		assert.Equal(t, hipathsys.NewString("Test1Test2"), res)
	}
}

func TestArithmeticExpressionBothStringSubtraction(t *testing.T) {
	e := NewArithmeticExpression(NewRawStringLiteral("Test1"),
		hipathsys.SubtractionOp, NewRawStringLiteral("Test2"))
	res, err := e.Evaluate(nil, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no res expected")
}

func TestArithmeticExpressionStringNil(t *testing.T) {
	e := NewArithmeticExpression(NewRawStringLiteral("Test1"),
		hipathsys.AdditionOp, NewRawStringLiteral("Test2"))
	res, err := e.Evaluate(nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.StringAccessor)(nil), res) {
		assert.Equal(t, hipathsys.NewString("Test1Test2"), res)
	}
}
