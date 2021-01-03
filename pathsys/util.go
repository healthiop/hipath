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
	"strconv"
	"strings"
	"unicode"
)

type normalizedStringStream struct {
	value []rune
	pos   int
	size  int
}

func (s *normalizedStringStream) next() (rune, bool) {
	begin := s.pos == 0
	ws := false
	for s.pos < s.size {
		c := s.value[s.pos]
		if c == ' ' || c == '\r' || c == '\n' || c == '\t' {
			s.pos = s.pos + 1
			ws = true
		} else if ws && !begin {
			return ' ', true
		} else {
			s.pos = s.pos + 1
			return unicode.ToLower(c), true
		}
	}
	return 0, false
}

func NormalizedStringEqual(value1 string, value2 string) bool {
	if value1 == value2 {
		return true
	}
	if len(value1) == 0 || len(value2) == 0 {
		return false
	}

	r1 := []rune(value1)
	s1 := normalizedStringStream{value: r1, size: len(r1)}
	r2 := []rune(value2)
	s2 := normalizedStringStream{value: r2, size: len(r2)}

	for c1, ok1 := s1.next(); ok1; c1, ok1 = s1.next() {
		c2, ok2 := s2.next()
		if !ok2 || c1 != c2 {
			return false
		}
	}
	_, ok := s2.next()
	return !ok
}

func writeStringBuilderInt(b *strings.Builder, value int, digits int) {
	formatted := strconv.FormatInt(int64(value), 10)
	l := len(formatted)
	for i := l; i < digits; i++ {
		b.WriteByte('0')
	}
	b.WriteString(formatted)
}
