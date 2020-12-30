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

package pathsys

import (
	"github.com/shopspring/decimal"
	"math"
)

var decimalTen = decimal.NewFromInt32(10)

type ArithmeticOps byte

const (
	AdditionOp       ArithmeticOps = '+'
	SubtractionOp    ArithmeticOps = '-'
	MultiplicationOp ArithmeticOps = '*'
	DivisionOp       ArithmeticOps = '/'
	DivOp            ArithmeticOps = 'D'
	ModOp            ArithmeticOps = 'M'
)

type DecimalValueAccessor interface {
	AnyAccessor
	Value() DecimalAccessor
	WithValue(node NumberAccessor) DecimalValueAccessor
	ArithmeticOpSupported(op ArithmeticOps) bool
}

type ArithmeticApplier interface {
	Calc(operand DecimalValueAccessor, op ArithmeticOps) (DecimalValueAccessor, error)
	Abs() DecimalValueAccessor
}

type NumberAccessor interface {
	AnyAccessor
	Comparator
	Stringifier
	DecimalValueAccessor
	Negator
	ArithmeticApplier
	Truncate(precision int32) NumberAccessor
	Int() int32
	Int64() int64
	Float64() float64
	Decimal() decimal.Decimal
}

func leastPrecisionDecimal(d1 decimal.Decimal, d2 decimal.Decimal) (decimal.Decimal, decimal.Decimal) {
	p1, p2 := decimalPrecision(d1), decimalPrecision(d2)
	if p1 == p2 {
		return d1, d2
	}

	if p1 < p2 {
		return d1, d2.Truncate(p1)
	}
	return d1.Truncate(p2), d2
}

func decimalPrecision(d decimal.Decimal) int32 {
	precision := -d.Exponent()
	if precision <= 0 {
		return 0
	}

	v := d.Mul(decimal.NewFromInt(int64(math.Pow(10.0, float64(precision)))))
	for precision > 0 {
		m := v.Mod(decimalTen)
		if !m.IsZero() {
			return precision
		}
		v = v.Div(decimalTen)
		precision = precision - 1
	}
	return precision
}
