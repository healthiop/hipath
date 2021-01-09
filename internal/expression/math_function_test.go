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
)

func TestAbsFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newAbsFunction()
	res, err := f.Execute(ctx, nil, []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestAbsFuncOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newAbsFunction()
	res, err := f.Execute(ctx, hipathsys.NewString("test"), []interface{}{}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestAbsFuncIntegerPos(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newAbsFunction()
	res, err := f.Execute(ctx, hipathsys.NewInteger(10), []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.NewInteger(10), res)
}

func TestAbsFuncIntegerNegs(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newAbsFunction()
	res, err := f.Execute(ctx, hipathsys.NewInteger(-10), []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.NewInteger(10), res)
}

func TestAbsFuncIntegerCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.MustAdd(hipathsys.NewInteger(-10))

	f := newAbsFunction()
	res, err := f.Execute(ctx, hipathsys.NewInteger(10), []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.NewInteger(10), res)
}

func TestAbsFuncDecimalPos(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newAbsFunction()
	res, err := f.Execute(ctx, hipathsys.NewDecimalFloat64(10.21), []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.DecimalAccessor)(nil), res) {
		assert.Equal(t, 10.21, res.(hipathsys.DecimalAccessor).Float64())
	}
}

func TestAbsFuncDecimalNeg(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newAbsFunction()
	res, err := f.Execute(ctx, hipathsys.NewDecimalFloat64(-10.21), []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.DecimalAccessor)(nil), res) {
		assert.Equal(t, 10.21, res.(hipathsys.DecimalAccessor).Float64())
	}
}

func TestAbsFuncQuanityPos(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newAbsFunction()
	res, err := f.Execute(ctx, hipathsys.NewQuantity(hipathsys.NewDecimalFloat64(10.21), hipathsys.NewString("cm")), []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.QuantityAccessor)(nil), res) {
		assert.Equal(t, 10.21, res.(hipathsys.QuantityAccessor).Value().Float64())
		assert.Equal(t, hipathsys.NewString("cm"), res.(hipathsys.QuantityAccessor).Unit())
	}
}

func TestAbsFuncQuanityNeg(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newAbsFunction()
	res, err := f.Execute(ctx, hipathsys.NewQuantity(hipathsys.NewDecimalFloat64(-10.21), hipathsys.NewString("cm")), []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.QuantityAccessor)(nil), res) {
		assert.Equal(t, 10.21, res.(hipathsys.QuantityAccessor).Value().Float64())
		assert.Equal(t, hipathsys.NewString("cm"), res.(hipathsys.QuantityAccessor).Unit())
	}
}

func TestCeilingFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newCeilingFunction()
	res, err := f.Execute(ctx, nil, []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestCeilingFuncOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newCeilingFunction()
	res, err := f.Execute(ctx, hipathsys.NewString("test"), []interface{}{}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestCeilingFuncInteger(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newCeilingFunction()
	res, err := f.Execute(ctx, hipathsys.NewInteger(10), []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.IntegerAccessor)(nil), res) {
		assert.Equal(t, 10.0, res.(hipathsys.IntegerAccessor).Float64())
	}
}

func TestCeilingFuncIntegerCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.MustAdd(hipathsys.NewInteger(10))

	f := newCeilingFunction()
	res, err := f.Execute(ctx, col, []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.NewInteger(10), res)
}

func TestCeilingFuncDecimal(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newCeilingFunction()
	res, err := f.Execute(ctx, hipathsys.NewDecimalFloat64(10.21), []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.IntegerAccessor)(nil), res) {
		assert.Equal(t, 11.0, res.(hipathsys.IntegerAccessor).Float64())
	}
}

func TestFloorFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newFloorFunction()
	res, err := f.Execute(ctx, nil, []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestFloorFuncOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newFloorFunction()
	res, err := f.Execute(ctx, hipathsys.NewString("test"), []interface{}{}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestFloorFuncInteger(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newFloorFunction()
	res, err := f.Execute(ctx, hipathsys.NewInteger(10), []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.IntegerAccessor)(nil), res) {
		assert.Equal(t, 10.0, res.(hipathsys.IntegerAccessor).Float64())
	}
}

func TestFloorFuncIntegerCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.MustAdd(hipathsys.NewInteger(10))

	f := newFloorFunction()
	res, err := f.Execute(ctx, col, []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, hipathsys.NewInteger(10), res)
}

func TestFloorFuncDecimal(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newFloorFunction()
	res, err := f.Execute(ctx, hipathsys.NewDecimalFloat64(10.81), []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.IntegerAccessor)(nil), res) {
		assert.Equal(t, 10.0, res.(hipathsys.IntegerAccessor).Float64())
	}
}

func TestExpFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newExpFunction()
	res, err := f.Execute(ctx, nil, []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestExpFuncOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newExpFunction()
	res, err := f.Execute(ctx, hipathsys.NewString("test"), []interface{}{}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestExpFuncInteger(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newExpFunction()
	res, err := f.Execute(ctx, hipathsys.NewInteger(10), []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.DecimalAccessor)(nil), res) {
		assert.InDelta(t, 22026.46579, res.(hipathsys.DecimalAccessor).Float64(), .000005)
	}
}

func TestExpFuncIntegerCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.MustAdd(hipathsys.NewInteger(10))

	f := newExpFunction()
	res, err := f.Execute(ctx, col, []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.DecimalAccessor)(nil), res) {
		assert.InDelta(t, 22026.46579, res.(hipathsys.DecimalAccessor).Float64(), .000005)
	}
}

func TestExpFuncDecimal(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newExpFunction()
	res, err := f.Execute(ctx, hipathsys.NewDecimalFloat64(10.21), []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.DecimalAccessor)(nil), res) {
		assert.InDelta(t, 27173.567589, res.(hipathsys.DecimalAccessor).Float64(), .0000005)
	}
}

func TestLnFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newLnFunction()
	res, err := f.Execute(ctx, nil, []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestLnFuncOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newLnFunction()
	res, err := f.Execute(ctx, hipathsys.NewString("test"), []interface{}{}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestLnFuncInteger(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newLnFunction()
	res, err := f.Execute(ctx, hipathsys.NewInteger(10), []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.DecimalAccessor)(nil), res) {
		assert.InDelta(t, 2.302585, res.(hipathsys.DecimalAccessor).Float64(), .0000001)
	}
}

func TestLnFuncIntegerCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.MustAdd(hipathsys.NewInteger(10))

	f := newLnFunction()
	res, err := f.Execute(ctx, col, []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.DecimalAccessor)(nil), res) {
		assert.InDelta(t, 2.302585, res.(hipathsys.DecimalAccessor).Float64(), .0000001)
	}
}

func TestLnFuncDecimal(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newLnFunction()
	res, err := f.Execute(ctx, hipathsys.NewDecimalFloat64(10.81), []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.DecimalAccessor)(nil), res) {
		assert.InDelta(t, 2.38047, res.(hipathsys.DecimalAccessor).Float64(), .000002)
	}
}

func TestLnFuncError(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newLnFunction()
	res, err := f.Execute(ctx, hipathsys.NewDecimalFloat64(0), []interface{}{}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "NaN expected")
}

func TestLogFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newLogFunction()
	res, err := f.Execute(ctx, nil, []interface{}{hipathsys.NewInteger(1)}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestLogFuncBaseNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newLogFunction()
	res, err := f.Execute(ctx, hipathsys.NewInteger(1), []interface{}{nil}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestLogFuncOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newLogFunction()
	res, err := f.Execute(ctx, hipathsys.NewString("test"), []interface{}{hipathsys.NewInteger(1)}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestLogFuncBaseOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newLogFunction()
	res, err := f.Execute(ctx, hipathsys.NewInteger(1), []interface{}{hipathsys.NewString("test")}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestLogFuncInteger(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newLogFunction()
	res, err := f.Execute(ctx, hipathsys.NewInteger(10), []interface{}{hipathsys.NewInteger(5)}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.DecimalAccessor)(nil), res) {
		assert.InDelta(t, 1.430676, res.(hipathsys.DecimalAccessor).Float64(), .0000006)
	}
}

func TestLogFuncIntegerCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.MustAdd(hipathsys.NewInteger(10))

	f := newLogFunction()
	res, err := f.Execute(ctx, col, []interface{}{hipathsys.NewInteger(5)}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.DecimalAccessor)(nil), res) {
		assert.InDelta(t, 1.430676, res.(hipathsys.DecimalAccessor).Float64(), .0000006)
	}
}

func TestLogFuncIntegerBaseCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.MustAdd(hipathsys.NewInteger(5))

	f := newLogFunction()
	res, err := f.Execute(ctx, hipathsys.NewInteger(10), []interface{}{col}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.DecimalAccessor)(nil), res) {
		assert.InDelta(t, 1.430676, res.(hipathsys.DecimalAccessor).Float64(), .0000006)
	}
}

func TestLogFuncDecimal(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newLogFunction()
	res, err := f.Execute(ctx, hipathsys.NewDecimalFloat64(10.81), []interface{}{hipathsys.NewDecimalFloat64(5.12)}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.DecimalAccessor)(nil), res) {
		assert.InDelta(t, 1.45759, res.(hipathsys.DecimalAccessor).Float64(), .00002)
	}
}

func TestLogFuncError(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newLogFunction()
	res, err := f.Execute(ctx, hipathsys.NewDecimalFloat64(0), []interface{}{hipathsys.NewDecimalFloat64(2)}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "NaN expected")
}

func TestLogFuncErrorBase(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newLogFunction()
	res, err := f.Execute(ctx, hipathsys.NewDecimalFloat64(10), []interface{}{hipathsys.NewDecimalFloat64(0)}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "NaN expected")
}

func TestPowerFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newPowerFunction()
	res, err := f.Execute(ctx, nil, []interface{}{hipathsys.NewInteger(1)}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestPowerFuncExpNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newPowerFunction()
	res, err := f.Execute(ctx, hipathsys.NewInteger(1), []interface{}{nil}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestPowerFuncOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newPowerFunction()
	res, err := f.Execute(ctx, hipathsys.NewString("test"), []interface{}{hipathsys.NewInteger(1)}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestPowerFuncExpOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newPowerFunction()
	res, err := f.Execute(ctx, hipathsys.NewInteger(1), []interface{}{hipathsys.NewString("test")}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestPowerFuncInteger(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newPowerFunction()
	res, err := f.Execute(ctx, hipathsys.NewInteger(4), []interface{}{hipathsys.NewInteger(3)}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.IntegerAccessor)(nil), res) {
		assert.Equal(t, 64.0, res.(hipathsys.IntegerAccessor).Float64())
	}
}

func TestPowerFuncIntegerCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.MustAdd(hipathsys.NewInteger(4))

	f := newPowerFunction()
	res, err := f.Execute(ctx, col, []interface{}{hipathsys.NewInteger(3)}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.IntegerAccessor)(nil), res) {
		assert.Equal(t, 64.0, res.(hipathsys.IntegerAccessor).Float64())
	}
}

func TestPowerFuncIntegerExpCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.MustAdd(hipathsys.NewInteger(3))

	f := newPowerFunction()
	res, err := f.Execute(ctx, hipathsys.NewInteger(4), []interface{}{col}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.IntegerAccessor)(nil), res) {
		assert.Equal(t, 64.0, res.(hipathsys.IntegerAccessor).Float64())
	}
}

func TestPowerFuncDecimal(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newPowerFunction()
	res, err := f.Execute(ctx, hipathsys.NewDecimalFloat64(4.5), []interface{}{hipathsys.NewDecimalFloat64(3.2)}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.DecimalAccessor)(nil), res) {
		assert.InDelta(t, 123.10623, res.(hipathsys.DecimalAccessor).Float64(), .000004)
	}
}

func TestPowerFuncEmpty(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newPowerFunction()
	res, err := f.Execute(ctx, hipathsys.NewDecimalFloat64(-1), []interface{}{hipathsys.NewDecimalFloat64(0.5)}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "NaN expected")
}

func TestRoundFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newRoundFunction()
	res, err := f.Execute(ctx, nil, []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestRoundFuncPewcisionNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newRoundFunction()
	res, err := f.Execute(ctx, hipathsys.NewInteger(1), []interface{}{nil}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestRoundFuncOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newRoundFunction()
	res, err := f.Execute(ctx, hipathsys.NewString("test"), []interface{}{}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestRoundFuncPrecisionOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newRoundFunction()
	res, err := f.Execute(ctx, hipathsys.NewInteger(1), []interface{}{hipathsys.NewString("test")}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestRoundFuncInteger(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newRoundFunction()
	res, err := f.Execute(ctx, hipathsys.NewInteger(4), []interface{}{hipathsys.NewInteger(3)}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.DecimalAccessor)(nil), res) {
		assert.Equal(t, 4.0, res.(hipathsys.DecimalAccessor).Float64())
	}
}

func TestRoundFuncIntegerCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.MustAdd(hipathsys.NewInteger(4))

	f := newRoundFunction()
	res, err := f.Execute(ctx, col, []interface{}{hipathsys.NewInteger(3)}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.DecimalAccessor)(nil), res) {
		assert.Equal(t, 4.0, res.(hipathsys.DecimalAccessor).Float64())
	}
}

func TestRoundFuncPrecisionCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.MustAdd(hipathsys.NewInteger(2))

	f := newRoundFunction()
	res, err := f.Execute(ctx, hipathsys.NewDecimalFloat64(3.255), []interface{}{col}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.DecimalAccessor)(nil), res) {
		assert.Equal(t, 3.26, res.(hipathsys.DecimalAccessor).Float64())
	}
}

func TestRoundFuncDecimal(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newRoundFunction()
	res, err := f.Execute(ctx, hipathsys.NewDecimalFloat64(3.255), []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.DecimalAccessor)(nil), res) {
		assert.InDelta(t, 3.0, res.(hipathsys.DecimalAccessor).Float64(), .000004)
	}
}

func TestRoundFuncDecimalPrecision(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newRoundFunction()
	res, err := f.Execute(ctx, hipathsys.NewDecimalFloat64(3.255), []interface{}{hipathsys.NewInteger(2)}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.DecimalAccessor)(nil), res) {
		assert.InDelta(t, 3.26, res.(hipathsys.DecimalAccessor).Float64(), .000004)
	}
}

func TestRoundFuncError(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newRoundFunction()
	res, err := f.Execute(ctx, hipathsys.NewDecimalFloat64(-1), []interface{}{hipathsys.NewInteger(-1)}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "NaN expected")
}

func TestSqrtFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newSqrtFunction()
	res, err := f.Execute(ctx, nil, []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestSqrtFuncOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newSqrtFunction()
	res, err := f.Execute(ctx, hipathsys.NewString("test"), []interface{}{}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestSqrtFuncInteger(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newSqrtFunction()
	res, err := f.Execute(ctx, hipathsys.NewInteger(10), []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.DecimalAccessor)(nil), res) {
		assert.InDelta(t, 3.16227, res.(hipathsys.DecimalAccessor).Float64(), .000008)
	}
}

func TestSqrtFuncIntegerCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.MustAdd(hipathsys.NewInteger(10))

	f := newSqrtFunction()
	res, err := f.Execute(ctx, col, []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.DecimalAccessor)(nil), res) {
		assert.InDelta(t, 3.16227, res.(hipathsys.DecimalAccessor).Float64(), .000008)
	}
}

func TestSqrtFuncDecimal(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newSqrtFunction()
	res, err := f.Execute(ctx, hipathsys.NewDecimalFloat64(10.81), []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.DecimalAccessor)(nil), res) {
		assert.InDelta(t, 3.287856, res.(hipathsys.DecimalAccessor).Float64(), .0000005)
	}
}

func TestSqrtFuncNaN(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newSqrtFunction()
	res, err := f.Execute(ctx, hipathsys.NewDecimalFloat64(-1), []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty result expected")
}

func TestTruncateFuncNil(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newTruncateFunction()
	res, err := f.Execute(ctx, nil, []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestTruncateFuncOther(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newTruncateFunction()
	res, err := f.Execute(ctx, hipathsys.NewString("test"), []interface{}{}, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestTruncateFuncInteger(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newTruncateFunction()
	res, err := f.Execute(ctx, hipathsys.NewInteger(10), []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.IntegerAccessor)(nil), res) {
		assert.InDelta(t, 10, res.(hipathsys.IntegerAccessor).Float64(), .0000001)
	}
}

func TestTruncateFuncIntegerCol(t *testing.T) {
	ctx := test.NewTestContext(t)

	col := ctx.NewCollection()
	col.MustAdd(hipathsys.NewInteger(10))

	f := newTruncateFunction()
	res, err := f.Execute(ctx, col, []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.IntegerAccessor)(nil), res) {
		assert.InDelta(t, 10, res.(hipathsys.IntegerAccessor).Float64(), .0000001)
	}
}

func TestTruncateFuncDecimal(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newTruncateFunction()
	res, err := f.Execute(ctx, hipathsys.NewDecimalFloat64(10.81), []interface{}{}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.IntegerAccessor)(nil), res) {
		assert.InDelta(t, 10.0, res.(hipathsys.IntegerAccessor).Float64(), .000002)
	}
}
