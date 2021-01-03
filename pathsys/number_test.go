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

package pathsys

import (
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLestPrecisionDecimalEqual(t *testing.T) {
	d1, d2 := leastPrecisionDecimal(
		decimal.NewFromFloat(-7283.18),
		decimal.NewFromFloat(82737263.28))
	assert.Equal(t, "-7283.18", d1.String())
	assert.Equal(t, "82737263.28", d2.String())
}

func TestLestPrecisionDecimalLeft(t *testing.T) {
	d1, d2 := leastPrecisionDecimal(
		decimal.NewFromFloat(-7283.1),
		decimal.NewFromFloat(82737263.28))
	assert.Equal(t, "-7283.1", d1.String())
	assert.Equal(t, "82737263.2", d2.String())
}

func TestLestPrecisionDecimalRight(t *testing.T) {
	d1, d2 := leastPrecisionDecimal(
		decimal.NewFromFloat(-7283.18),
		decimal.NewFromFloat(82737263.2))
	assert.Equal(t, "-7283.1", d1.String())
	assert.Equal(t, "82737263.2", d2.String())
}

func TestLestPrecisionDecimalNone(t *testing.T) {
	v1, _ := decimal.NewFromString("-7283.00000")
	v2, _ := decimal.NewFromString("82737263.00")
	d1, d2 := leastPrecisionDecimal(v1, v2)
	assert.Equal(t, "-7283", d1.String())
	assert.Equal(t, "82737263", d2.String())
}

func TestDecimalPrecisionZeroPrecision(t *testing.T) {
	assert.Equal(t, int32(0), decimalPrecision(decimal.NewFromInt32(83720)))
}

func TestDecimalPrecisionWithFractionDigits(t *testing.T) {
	assert.Equal(t, int32(3), decimalPrecision(decimal.NewFromFloat(87267.232)))
}

func TestDecimalPrecisionWithTrailingZeros(t *testing.T) {
	v, err := decimal.NewFromString("-87267.232000000000000")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, int32(3), decimalPrecision(v))
}

func TestDecimalPrecisionZeros(t *testing.T) {
	v, err := decimal.NewFromString("0.000000000000000")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, int32(0), decimalPrecision(v))
}
