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
	"github.com/volsch/gohipath/pathsys"
	"strconv"
	"strings"
)

const stringDelimiterChar = '\''

type StringLiteral struct {
	node pathsys.StringAccessor
}

func NewRawStringLiteral(value string) pathsys.Evaluator {
	return &StringLiteral{pathsys.NewString(value)}
}

func ParseStringLiteral(value string) pathsys.Evaluator {
	return &StringLiteral{pathsys.NewString(parseStringLiteral(value, stringDelimiterChar))}
}

func parseStringLiteral(value string, delimiter byte) string {
	l := len(value)
	if l < 2 || value[0] != delimiter || value[l-1] != delimiter {
		return value
	}

	var b strings.Builder
	b.Grow(l)

	var esc bool
	var unicode bool
	var unicodeValue strings.Builder
	for _, char := range value[1 : l-1] {
		if unicode {
			unicodeValue.WriteRune(char)
			if unicodeValue.Len() == 4 {
				if r, err := strconv.ParseInt(unicodeValue.String(), 16, 32); err != nil {
					b.WriteRune('u')
					b.WriteString(unicodeValue.String())
				} else {
					b.WriteRune(int32(r))
				}
				unicode = false
			}
		} else if esc {
			switch char {
			case 'r':
				b.WriteRune('\r')
			case 'n':
				b.WriteRune('\n')
			case 't':
				b.WriteRune('\t')
			case 'f':
				b.WriteRune('\f')
			case 'u':
				unicode = true
				unicodeValue.Grow(4)
				unicodeValue.Reset()
			default:
				b.WriteRune(char)
			}
			esc = false
		} else if char == '\\' {
			esc = true
		} else {
			b.WriteRune(char)
		}
	}

	if unicode {
		b.WriteRune('u')
		b.WriteString(unicodeValue.String())
	}

	return b.String()
}

func (e *StringLiteral) Evaluate(pathsys.ContextAccessor, interface{}, pathsys.Looper) (interface{}, error) {
	return e.node, nil
}
