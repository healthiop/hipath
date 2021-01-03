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
	"fmt"
	"github.com/volsch/gohipath/pathsys"
	"regexp"
)

var quantityUnitRegexp = regexp.MustCompile("^[^\\s]+(\\s[^\\s]+)*$")

type QuantityLiteral struct {
	node pathsys.QuantityAccessor
}

func ParseQuantityLiteral(number string, unit string) (pathsys.Evaluator, error) {
	value, err := pathsys.ParseDecimal(number)
	if err != nil {
		return nil, err
	}
	convertedUnit, err := parseQuantityUnit(unit)
	if err != nil {
		return nil, err
	}

	return &QuantityLiteral{pathsys.NewQuantity(value, convertedUnit)}, nil
}

func parseQuantityUnit(unit string) (pathsys.StringAccessor, error) {
	if len(unit) == 0 || unit == "''" {
		return nil, nil
	}

	u := parseStringLiteral(unit, stringDelimiterChar)
	if !quantityUnitRegexp.MatchString(u) {
		return nil, fmt.Errorf("invalid quantity unit: %s", u)
	}
	return pathsys.NewString(u), nil
}

func (e *QuantityLiteral) Evaluate(pathsys.ContextAccessor, interface{}, pathsys.Looper) (interface{}, error) {
	return e.node, nil
}
