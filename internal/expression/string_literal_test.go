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
	"github.com/volsch/gohimodel/datatype"
	"testing"
)

func TestStringLiteral(t *testing.T) {
	evaluator := ParseStringLiteral(
		"'x\\ra\\nb\\tc\\fd\\\\e\\'f\\\"g\\`h\\u0076i\\u23DAj\\pk'")

	assert.NotNil(t, evaluator, "evaluator expected")
	if evaluator != nil {
		accessor, err := evaluator.Evaluate(nil, nil)
		assert.NoError(t, err, "no error expected")
		assert.NotNil(t, accessor, "accessor expected")
		if assert.Implements(t, (*datatype.StringAccessor)(nil), accessor) {
			assert.Equal(t, "x\ra\nb\tc\fd\\e'f\"g`hvi⏚jpk",
				accessor.(datatype.StringAccessor).Value())
		}
	}
}

func TestParseStringLiteralShortUnicode(t *testing.T) {
	assert.Equal(t, "u005", parseStringLiteral("'\\u005'", stringDelimiterChar))
}

func TestParseStringLiteralInvalidUnicode(t *testing.T) {
	assert.Equal(t, "aux005b", parseStringLiteral("'a\\ux005b'", stringDelimiterChar))
}

func TestParseStringLiteralNoEscapedChar(t *testing.T) {
	assert.Equal(t, "", parseStringLiteral("'\\'", stringDelimiterChar))
}

func TestParseStringLiteralEmpty(t *testing.T) {
	assert.Equal(t, "", parseStringLiteral("", stringDelimiterChar))
}

func TestParseStringLiteralDelimited(t *testing.T) {
	assert.Equal(t, "Test", parseStringLiteral("`Test`", '`'))
}
