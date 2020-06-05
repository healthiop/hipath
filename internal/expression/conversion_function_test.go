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
	"github.com/volsch/gohipath/internal/test"
	"github.com/volsch/gohipath/pathsys"
	"testing"
	"time"
)

func TestIIfPathFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newIIfFunction()
	res, err := f.Execute(ctx, nil, []interface{}{nil, nil}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestIIfPathFuncNilOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newIIfFunction()
	res, err := f.Execute(ctx, nil, []interface{}{nil, nil, pathsys.NewString("other")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewString("other"), res)
}

func TestIIfPathFuncTrue(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newIIfFunction()
	res, err := f.Execute(ctx, nil, []interface{}{pathsys.True, pathsys.NewString("match"), pathsys.NewString("other")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewString("match"), res)
}

func TestIIfPathFuncFalse(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newIIfFunction()
	res, err := f.Execute(ctx, nil, []interface{}{pathsys.False, pathsys.NewString("match"), pathsys.NewString("other")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewString("other"), res)
}

func TestIIfPathFuncFalseNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newIIfFunction()
	res, err := f.Execute(ctx, nil, []interface{}{pathsys.False, pathsys.NewString("match")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestIIfPathFuncInvalidType(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newIIfFunction()
	res, err := f.Execute(ctx, nil, []interface{}{pathsys.NewString(""), pathsys.NewString("match")}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestToBooleanFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toBooleanFunc
	res, err := f.Execute(ctx, nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestToBooleanFuncOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toBooleanFunc
	res, err := f.Execute(ctx, "test", nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestToBooleanFuncMultiCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.True)
	col.Add(pathsys.True)

	f := toBooleanFunc
	res, err := f.Execute(ctx, col, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestToBooleanFuncBooleanTrue(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toBooleanFunc
	res, err := f.Execute(ctx, pathsys.True, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.True, res)
}

func TestToBooleanFuncBooleanFalse(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toBooleanFunc
	res, err := f.Execute(ctx, pathsys.False, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.False, res)
}

func TestToBooleanFuncStringTrue(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toBooleanFunc
	res, err := f.Execute(ctx, pathsys.NewString("YeS"), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.True, res)
}

func TestToBooleanFuncStringFalse(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toBooleanFunc
	res, err := f.Execute(ctx, pathsys.NewString("No"), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.False, res)
}

func TestToBooleanFuncStringOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toBooleanFunc
	res, err := f.Execute(ctx, pathsys.NewString("xyz"), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestToBooleanFuncIntegerTrue(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toBooleanFunc
	res, err := f.Execute(ctx, pathsys.NewInteger(1), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.True, res)
}

func TestToBooleanFuncIntegerFalse(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toBooleanFunc
	res, err := f.Execute(ctx, pathsys.NewInteger(0), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.False, res)
}

func TestToBooleanFuncIntegerOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toBooleanFunc
	res, err := f.Execute(ctx, pathsys.NewInteger(2), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestToBooleanFuncDecimalTrue(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toBooleanFunc
	res, err := f.Execute(ctx, pathsys.NewDecimalInt(1), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.True, res)
}

func TestToBooleanFuncDecimalFalse(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toBooleanFunc
	res, err := f.Execute(ctx, pathsys.NewDecimalInt(0), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.False, res)
}

func TestToBooleanFuncDecimalOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toBooleanFunc
	res, err := f.Execute(ctx, pathsys.NewDecimalInt(2), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestConvertsToBooleanFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newConvertsToBooleanFunction()
	res, err := f.Execute(ctx, nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.False, res)
}

func TestConvertsToBooleanFuncTooMany(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.True)
	col.Add(pathsys.True)

	f := newConvertsToBooleanFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestConvertToBoolean(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newConvertsToBooleanFunction()
	res, err := f.Execute(ctx, pathsys.NewString("No"), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.True, res)
}

func TestConvertToBooleanNot(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newConvertsToBooleanFunction()
	res, err := f.Execute(ctx, pathsys.NewString("Other"), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.False, res)
}

func TestToIntegerFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toIntegerFunc
	res, err := f.Execute(ctx, nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestToIntegerFuncOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toIntegerFunc
	res, err := f.Execute(ctx, "test", nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestToIntegerFuncDecimal(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toIntegerFunc
	res, err := f.Execute(ctx, pathsys.NewDecimalInt(123), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestToIntegerFuncMultiCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewInteger(10))
	col.Add(pathsys.NewInteger(10))

	f := toIntegerFunc
	res, err := f.Execute(ctx, col, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestToIntegerFuncInteger(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toIntegerFunc
	res, err := f.Execute(ctx, pathsys.NewInteger(123), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewInteger(123), res)
}

func TestToIntegerFuncStringPos(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toIntegerFunc
	res, err := f.Execute(ctx, pathsys.NewString("+123"), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewInteger(123), res)
}

func TestToIntegerFuncStringNeg(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toIntegerFunc
	res, err := f.Execute(ctx, pathsys.NewString("-123"), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewInteger(-123), res)
}

func TestToIntegerFuncBooleanTrue(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toIntegerFunc
	res, err := f.Execute(ctx, pathsys.True, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewInteger(1), res)
}

func TestToIntegerFuncBooleanFalse(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toIntegerFunc
	res, err := f.Execute(ctx, pathsys.False, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewInteger(0), res)
}

func TestConvertsToIntegerFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newConvertsToIntegerFunction()
	res, err := f.Execute(ctx, nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.False, res)
}

func TestConvertsToIntegerFuncTooMany(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewInteger(10))
	col.Add(pathsys.NewInteger(10))

	f := newConvertsToIntegerFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestConvertToInteger(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newConvertsToIntegerFunction()
	res, err := f.Execute(ctx, pathsys.NewString("123"), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.True, res)
}

func TestConvertToIntegerNot(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newConvertsToIntegerFunction()
	res, err := f.Execute(ctx, pathsys.NewString("Other"), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.False, res)
}

func TestToDecimalFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toDecimalFunc
	res, err := f.Execute(ctx, nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestToDecimalFuncOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toDecimalFunc
	res, err := f.Execute(ctx, "test", nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestToDecimalFuncInteger(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toDecimalFunc
	res, err := f.Execute(ctx, pathsys.NewInteger(123), nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.DecimalAccessor)(nil), res) {
		d := res.(pathsys.DecimalAccessor)
		assert.Equal(t, 123.0, d.Float64())
	}
}

func TestToDecimalFuncMultiCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewDecimalInt(10))
	col.Add(pathsys.NewDecimalInt(10))

	f := toDecimalFunc
	res, err := f.Execute(ctx, col, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestToDecimalFuncDecimal(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toDecimalFunc
	res, err := f.Execute(ctx, pathsys.NewDecimalFloat64(123.56), nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.DecimalAccessor)(nil), res) {
		d := res.(pathsys.DecimalAccessor)
		assert.Equal(t, 123.56, d.Float64())
	}
}

func TestToDecimalFuncStringPos(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toDecimalFunc
	res, err := f.Execute(ctx, pathsys.NewString("+123.56"), nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.DecimalAccessor)(nil), res) {
		d := res.(pathsys.DecimalAccessor)
		assert.Equal(t, 123.56, d.Float64())
	}
}

func TestToDecimalFuncStringNeg(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toDecimalFunc
	res, err := f.Execute(ctx, pathsys.NewString("-123.56"), nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.DecimalAccessor)(nil), res) {
		d := res.(pathsys.DecimalAccessor)
		assert.Equal(t, -123.56, d.Float64())
	}
}

func TestToDecimalFuncBooleanTrue(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toDecimalFunc
	res, err := f.Execute(ctx, pathsys.True, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.DecimalAccessor)(nil), res) {
		d := res.(pathsys.DecimalAccessor)
		assert.Equal(t, 1.0, d.Float64())
	}
}

func TestToDecimalFuncBooleanFalse(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toDecimalFunc
	res, err := f.Execute(ctx, pathsys.False, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.DecimalAccessor)(nil), res) {
		d := res.(pathsys.DecimalAccessor)
		assert.Equal(t, 0.0, d.Float64())
	}
}

func TestConvertsToDecimalFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newConvertsToDecimalFunction()
	res, err := f.Execute(ctx, nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.False, res)
}

func TestConvertsToDecimalFuncTooMany(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewDecimalInt(10))
	col.Add(pathsys.NewDecimalInt(10))

	f := newConvertsToDecimalFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestConvertToDecimal(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newConvertsToDecimalFunction()
	res, err := f.Execute(ctx, pathsys.NewString("123.56"), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.True, res)
}

func TestConvertToDecimalNot(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newConvertsToDecimalFunction()
	res, err := f.Execute(ctx, pathsys.NewString("Other"), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.False, res)
}

func TestToDateFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toDateFunc
	res, err := f.Execute(ctx, nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestToDateFuncOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toDateFunc
	res, err := f.Execute(ctx, "test", nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestToDateFuncDateTime(t *testing.T) {
	ctx := test.NewTestContext(t)

	now := time.Now()

	f := toDateFunc
	res, err := f.Execute(ctx, pathsys.NewDateTime(now), nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.DateAccessor)(nil), res) {
		d := res.(pathsys.DateAccessor)
		assert.Equal(t, now.Year(), d.Year())
		assert.Equal(t, int(now.Month()), d.Month())
		assert.Equal(t, now.Day(), d.Day())
		assert.Equal(t, pathsys.DayDatePrecision, d.Precision())
	}
}

func TestToDateFuncMultiCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewDateTime(time.Now()))
	col.Add(pathsys.NewDateTime(time.Now()))

	f := toDateFunc
	res, err := f.Execute(ctx, col, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestToDateFuncDate(t *testing.T) {
	ctx := test.NewTestContext(t)
	d := pathsys.NewDate(time.Now())

	f := toDateFunc
	res, err := f.Execute(ctx, d, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Same(t, d, res)
}

func TestToDateFuncString(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toDateFunc
	res, err := f.Execute(ctx, pathsys.NewString("2020-08-27"), nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.DateAccessor)(nil), res) {
		d := res.(pathsys.DateAccessor)
		assert.Equal(t, 2020, d.Year())
		assert.Equal(t, 8, d.Month())
		assert.Equal(t, 27, d.Day())
		assert.Equal(t, pathsys.DayDatePrecision, d.Precision())
	}
}

func TestToDateFuncStringPrecision(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toDateFunc
	res, err := f.Execute(ctx, pathsys.NewString("2020-08"), nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.DateAccessor)(nil), res) {
		d := res.(pathsys.DateAccessor)
		assert.Equal(t, 2020, d.Year())
		assert.Equal(t, 8, d.Month())
		assert.Equal(t, 1, d.Day())
		assert.Equal(t, pathsys.MonthDatePrecision, d.Precision())
	}
}

func TestToDateFuncStringInvalid(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toDateFunc
	res, err := f.Execute(ctx, pathsys.NewString("test"), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestConvertsToDateFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newConvertsToDateFunction()
	res, err := f.Execute(ctx, nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.False, res)
}

func TestConvertsToDateFuncTooMany(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewDateTime(time.Now()))
	col.Add(pathsys.NewDateTime(time.Now()))

	f := newConvertsToDateFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestConvertToDate(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newConvertsToDateFunction()
	res, err := f.Execute(ctx, pathsys.NewString("2018-11-27"), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.True, res)
}

func TestConvertToDateNot(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newConvertsToDateFunction()
	res, err := f.Execute(ctx, pathsys.NewString("Other"), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.False, res)
}

func TestToDateTimeFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toDateTimeFunc
	res, err := f.Execute(ctx, nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestToDateTimeFuncOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toDateTimeFunc
	res, err := f.Execute(ctx, "test", nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestToDateTimeFuncDate(t *testing.T) {
	ctx := test.NewTestContext(t)

	now := time.Now()

	f := toDateTimeFunc
	res, err := f.Execute(ctx, pathsys.NewDate(now), nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.DateTimeAccessor)(nil), res) {
		d := res.(pathsys.DateTimeAccessor)
		assert.Equal(t, now.Year(), d.Year())
		assert.Equal(t, int(now.Month()), d.Month())
		assert.Equal(t, now.Day(), d.Day())
		assert.Equal(t, pathsys.DayDatePrecision, d.Precision())
	}
}

func TestToDateTimeFuncMultiCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewDateTime(time.Now()))
	col.Add(pathsys.NewDateTime(time.Now()))

	f := toDateTimeFunc
	res, err := f.Execute(ctx, col, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestToDateTimeFuncDateTime(t *testing.T) {
	ctx := test.NewTestContext(t)
	d := pathsys.NewDateTime(time.Now())

	f := toDateTimeFunc
	res, err := f.Execute(ctx, d, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Same(t, d, res)
}

func TestToDateTimeFuncString(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toDateTimeFunc
	res, err := f.Execute(ctx, pathsys.NewString("2020-08-27T14:32:17"), nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.DateTimeAccessor)(nil), res) {
		d := res.(pathsys.DateTimeAccessor)
		assert.Equal(t, 2020, d.Year())
		assert.Equal(t, 8, d.Month())
		assert.Equal(t, 27, d.Day())
		assert.Equal(t, 14, d.Hour())
		assert.Equal(t, 32, d.Minute())
		assert.Equal(t, 17, d.Second())
		assert.Equal(t, pathsys.SecondTimePrecision, d.Precision())
	}
}

func TestToDateTimeFuncStringPrecision(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toDateTimeFunc
	res, err := f.Execute(ctx, pathsys.NewString("2020-08-27T14:32"), nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.DateTimeAccessor)(nil), res) {
		d := res.(pathsys.DateTimeAccessor)
		assert.Equal(t, 2020, d.Year())
		assert.Equal(t, 8, d.Month())
		assert.Equal(t, 27, d.Day())
		assert.Equal(t, 14, d.Hour())
		assert.Equal(t, 32, d.Minute())
		assert.Equal(t, 0, d.Second())
		assert.Equal(t, pathsys.MinuteTimePrecision, d.Precision())
	}
}

func TestToDateTimeFuncStringInvalid(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toDateTimeFunc
	res, err := f.Execute(ctx, pathsys.NewString("test"), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestConvertsToDateTimeFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newConvertsToDateTimeFunction()
	res, err := f.Execute(ctx, nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.False, res)
}

func TestConvertsToDateTimeFuncTooMany(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewDateTime(time.Now()))
	col.Add(pathsys.NewDateTime(time.Now()))

	f := newConvertsToDateTimeFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestConvertToDateTime(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newConvertsToDateTimeFunction()
	res, err := f.Execute(ctx, pathsys.NewString("2018-11-27T14:32:17.123+01:00"), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.True, res)
}

func TestConvertToDateTimeNot(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newConvertsToDateTimeFunction()
	res, err := f.Execute(ctx, pathsys.NewString("Other"), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.False, res)
}

func TestToQuantityFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toQuantityFunc
	res, err := f.Execute(ctx, nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestToQuantityFuncOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toQuantityFunc
	res, err := f.Execute(ctx, "test", nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestToQuantityFuncMultiCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewQuantity(pathsys.NewDecimalInt(10), pathsys.DayQuantityUnit.Plural()))
	col.Add(pathsys.NewQuantity(pathsys.NewDecimalInt(10), pathsys.DayQuantityUnit.Plural()))

	f := toQuantityFunc
	res, err := f.Execute(ctx, col, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestToQuantityFuncQuantityTrue(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toQuantityFunc
	res, err := f.Execute(ctx, pathsys.True, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.QuantityAccessor)(nil), res) {
		q := res.(pathsys.QuantityAccessor)
		assert.Equal(t, 1.0, q.Value().Float64())
		if assert.NotNil(t, q.Unit()) {
			assert.Equal(t, "1", q.Unit().String())
		}
	}
}

func TestToQuantityFuncQuantityFalse(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toQuantityFunc
	res, err := f.Execute(ctx, pathsys.False, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.QuantityAccessor)(nil), res) {
		q := res.(pathsys.QuantityAccessor)
		assert.Equal(t, 0.0, q.Value().Float64())
		if assert.NotNil(t, q.Unit()) {
			assert.Equal(t, "1", q.Unit().String())
		}
	}
}

func TestToQuantityFuncDecimal(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toQuantityFunc
	res, err := f.Execute(ctx, pathsys.NewDecimalFloat64(87.12), nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.QuantityAccessor)(nil), res) {
		q := res.(pathsys.QuantityAccessor)
		assert.Equal(t, 87.12, q.Value().Float64())
		if assert.NotNil(t, q.Unit()) {
			assert.Equal(t, "1", q.Unit().String())
		}
	}
}

func TestToQuantityFuncInteger(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toQuantityFunc
	res, err := f.Execute(ctx, pathsys.NewInteger(87), nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.QuantityAccessor)(nil), res) {
		q := res.(pathsys.QuantityAccessor)
		assert.Equal(t, 87.0, q.Value().Float64())
		if assert.NotNil(t, q.Unit()) {
			assert.Equal(t, "1", q.Unit().String())
		}
	}
}

func TestToQuantityFuncQuantity(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toQuantityFunc
	res, err := f.Execute(ctx, pathsys.NewQuantity(
		pathsys.NewDecimalFloat64(87.12), pathsys.YearQuantityUnit.Plural()), nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.QuantityAccessor)(nil), res) {
		q := res.(pathsys.QuantityAccessor)
		assert.Equal(t, 87.12, q.Value().Float64())
		if assert.NotNil(t, q.Unit()) {
			assert.Equal(t, "years", q.Unit().String())
		}
	}
}

func TestToQuantityFuncQuantityConvert(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toQuantityFunc
	res, err := f.Execute(ctx, pathsys.NewQuantity(
		pathsys.NewDecimalFloat64(2), pathsys.WeekQuantityUnit.Plural()),
		[]interface{}{pathsys.NewString("day")}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.QuantityAccessor)(nil), res) {
		q := res.(pathsys.QuantityAccessor)
		assert.Equal(t, 14.0, q.Value().Float64())
		if assert.NotNil(t, q.Unit()) {
			assert.Equal(t, "days", q.Unit().String())
		}
	}
}

func TestToQuantityFuncQuantityConvertInvalidUnit(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toQuantityFunc
	res, err := f.Execute(ctx, pathsys.NewQuantity(
		pathsys.NewDecimalFloat64(2), pathsys.WeekQuantityUnit.Plural()),
		[]interface{}{"day"}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestToQuantityFuncString(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toQuantityFunc
	res, err := f.Execute(ctx, pathsys.NewString("10.5 years"), nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.QuantityAccessor)(nil), res) {
		q := res.(pathsys.QuantityAccessor)
		assert.Equal(t, 10.5, q.Value().Float64())
		if assert.NotNil(t, q.Unit()) {
			assert.Equal(t, "years", q.Unit().String())
		}
	}
}

func TestToQuantityFuncStringWithoutUnit(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toQuantityFunc
	res, err := f.Execute(ctx, pathsys.NewString("10.5"), nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.QuantityAccessor)(nil), res) {
		q := res.(pathsys.QuantityAccessor)
		assert.Equal(t, 10.5, q.Value().Float64())
		assert.Nil(t, q.Unit())
	}
}

func TestToQuantityFuncStringInvalid(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toQuantityFunc
	res, err := f.Execute(ctx, pathsys.NewString("10.5 years2"), nil, nil)
	assert.NoError(t, err, "mo error expected")
	assert.Nil(t, res, "empty result expected")
}

func TestToQuantityFuncStringUCUM(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toQuantityFunc
	res, err := f.Execute(ctx, pathsys.NewString("-10.5 'cm2'"), nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.QuantityAccessor)(nil), res) {
		q := res.(pathsys.QuantityAccessor)
		assert.Equal(t, -10.5, q.Value().Float64())
		if assert.NotNil(t, q.Unit()) {
			assert.Equal(t, "cm2", q.Unit().String())
		}
	}
}

func TestConvertsToQuantityFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newConvertsToQuantityFunction()
	res, err := f.Execute(ctx, nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.False, res)
}

func TestConvertsToQuantityFuncTooMany(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewQuantity(pathsys.NewDecimalInt(10), pathsys.DayQuantityUnit.Plural()))
	col.Add(pathsys.NewQuantity(pathsys.NewDecimalInt(10), pathsys.DayQuantityUnit.Plural()))

	f := newConvertsToQuantityFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestConvertToQuantity(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newConvertsToQuantityFunction()
	res, err := f.Execute(ctx, pathsys.NewString("10 days"), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.True, res)
}

func TestConvertToQuantityInconvertibleUnit(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newConvertsToQuantityFunction()
	res, err := f.Execute(ctx, pathsys.NewString("10 days"), []interface{}{pathsys.NewString("cm")}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.False, res)
}

func TestConvertToQuantityNot(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newConvertsToQuantityFunction()
	res, err := f.Execute(ctx, pathsys.NewString("No"), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.False, res)
}

func TestToStringFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toStringFunc
	res, err := f.Execute(ctx, nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestToStringFuncOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toStringFunc
	res, err := f.Execute(ctx, "test", nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestToStringFuncMultiCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.True)
	col.Add(pathsys.True)

	f := toStringFunc
	res, err := f.Execute(ctx, col, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestToStringFuncString(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toStringFunc
	res, err := f.Execute(ctx, pathsys.NewString("Test"), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewString("Test"), res)
}

func TestToStringFuncBoolean(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toStringFunc
	res, err := f.Execute(ctx, pathsys.True, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewString("true"), res)
}

func TestToStringFuncInteger(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toStringFunc
	res, err := f.Execute(ctx, pathsys.NewInteger(8263), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewString("8263"), res)
}

func TestToStringFuncDecimal(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toStringFunc
	res, err := f.Execute(ctx, pathsys.NewDecimalFloat64(-18.82), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewString("-18.82"), res)
}

func TestToStringFuncQuantity(t *testing.T) {
	ctx := test.NewTestContext(t)
	q := pathsys.NewQuantity(pathsys.NewDecimalFloat64(10.5), pathsys.DayQuantityUnit.Plural())

	f := toStringFunc
	res, err := f.Execute(ctx, q, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewString("10.5 'days'"), res)
}

func TestToStringFuncDateTime(t *testing.T) {
	ctx := test.NewTestContext(t)
	d := pathsys.NewDateTimeYMDHMSNWithPrecision(2018, 8, 17, 21, 46, 6, 872673212, time.FixedZone("test", 120*60), pathsys.NanoTimePrecision)

	f := toStringFunc
	res, err := f.Execute(ctx, d, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewString("2018-08-17T21:46:06.872673212+02:00"), res)
}

func TestToStringFuncDateTimeUTC(t *testing.T) {
	ctx := test.NewTestContext(t)
	d := pathsys.NewDateTimeYMDHMSNWithPrecision(2018, 8, 17, 21, 46, 6, 872673212, time.UTC, pathsys.NanoTimePrecision)

	f := toStringFunc
	res, err := f.Execute(ctx, d, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewString("2018-08-17T21:46:06.872673212+00:00"), res)
}

func TestToStringFuncDate(t *testing.T) {
	ctx := test.NewTestContext(t)
	d := pathsys.NewDateYMDWithPrecision(2018, 8, 17, pathsys.DayDatePrecision)

	f := toStringFunc
	res, err := f.Execute(ctx, d, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewString("2018-08-17"), res)
}

func TestToStringFuncTime(t *testing.T) {
	ctx := test.NewTestContext(t)
	d := pathsys.NewTimeHMSNWithPrecision(21, 46, 6, 872673212, pathsys.NanoTimePrecision)

	f := toStringFunc
	res, err := f.Execute(ctx, d, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.NewString("21:46:06.872673212"), res)
}

func TestConvertsToStringFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newConvertsToStringFunction()
	res, err := f.Execute(ctx, nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.False, res)
}

func TestConvertsToStringFuncTooMany(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.True)
	col.Add(pathsys.True)

	f := newConvertsToStringFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestConvertToString(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newConvertsToStringFunction()
	res, err := f.Execute(ctx, pathsys.NewString("No"), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.True, res)
}

func TestConvertToStringNot(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newConvertsToStringFunction()
	res, err := f.Execute(ctx, "test", nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.False, res)
}

func TestToTimeFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toTimeFunc
	res, err := f.Execute(ctx, nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestToTimeFuncOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toTimeFunc
	res, err := f.Execute(ctx, "test", nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestToTimeFuncMultiCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewTime(time.Now()))
	col.Add(pathsys.NewTime(time.Now()))

	f := toTimeFunc
	res, err := f.Execute(ctx, col, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestToTimeFuncTime(t *testing.T) {
	ctx := test.NewTestContext(t)
	d := pathsys.NewTime(time.Now())

	f := toTimeFunc
	res, err := f.Execute(ctx, d, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Same(t, d, res)
}

func TestToTimeFuncString(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toTimeFunc
	res, err := f.Execute(ctx, pathsys.NewString("14:36:49.726126128726126128"), nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.TimeAccessor)(nil), res) {
		d := res.(pathsys.TimeAccessor)
		assert.Equal(t, 14, d.Hour())
		assert.Equal(t, 36, d.Minute())
		assert.Equal(t, 49, d.Second())
		assert.Equal(t, 726126128, d.Nanosecond())
		assert.Equal(t, pathsys.NanoTimePrecision, d.Precision())
	}
}

func TestToTimeFuncStringPrecision(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toTimeFunc
	res, err := f.Execute(ctx, pathsys.NewString("14:36"), nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.TimeAccessor)(nil), res) {
		d := res.(pathsys.TimeAccessor)
		assert.Equal(t, 14, d.Hour())
		assert.Equal(t, 36, d.Minute())
		assert.Equal(t, 0, d.Second())
		assert.Equal(t, 0, d.Nanosecond())
		assert.Equal(t, pathsys.MinuteTimePrecision, d.Precision())
	}
}

func TestToTimeFuncStringInvalid(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := toTimeFunc
	res, err := f.Execute(ctx, pathsys.NewString("test"), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestConvertsToTimeFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newConvertsToTimeFunction()
	res, err := f.Execute(ctx, nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.False, res)
}

func TestConvertsToTimeFuncTooMany(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.Add(pathsys.NewTime(time.Now()))
	col.Add(pathsys.NewTime(time.Now()))

	f := newConvertsToTimeFunction()
	res, err := f.Execute(ctx, col, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
}

func TestConvertToTime(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newConvertsToTimeFunction()
	res, err := f.Execute(ctx, pathsys.NewString("14:28:39"), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.True, res)
}

func TestConvertToTimeNot(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newConvertsToTimeFunction()
	res, err := f.Execute(ctx, pathsys.NewString("Other"), nil, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, pathsys.False, res)
}
