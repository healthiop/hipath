// Copyright (c) 2020, Volker Schmidt (volker@volsch.eu)
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source URI must retain the above copyright notice, this
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
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNormalizedStringEqual(t *testing.T) {
	assert.Equal(t, true, NormalizedStringEqual("This is a test", "This is a test"))
}

func TestNormalizedStringEqualLeftEmpty(t *testing.T) {
	assert.Equal(t, false, NormalizedStringEqual("", "This is a test"))
}

func TestNormalizedStringEqualRightEmpty(t *testing.T) {
	assert.Equal(t, false, NormalizedStringEqual("This is a test", ""))
}

func TestNormalizedStringEqualCaseInsensitive(t *testing.T) {
	assert.Equal(t, true, NormalizedStringEqual("TeSt", "tEsT"))
}

func TestNormalizedStringEqualWhitespaceLeft(t *testing.T) {
	assert.Equal(t, true, NormalizedStringEqual("\r\nTeSt\r  \tUnder\r ", "Test Under"))
}

func TestNormalizedStringEqualWhitespaceRight(t *testing.T) {
	assert.Equal(t, true, NormalizedStringEqual("Test Under", "\r\nTeSt\r  \tUnder\r "))
}

func TestNormalizedStringEqualDiffers(t *testing.T) {
	assert.Equal(t, false, NormalizedStringEqual("Test", "Turn"))
}

func TestWriteStringBuilderInt(t *testing.T) {
	var b strings.Builder
	writeStringBuilderInt(&b, 24, 4)
	assert.Equal(t, "0024", b.String())
}

func TestWriteStringBuilderIntSufficientDigits(t *testing.T) {
	var b strings.Builder
	writeStringBuilderInt(&b, 8724, 4)
	assert.Equal(t, "8724", b.String())
}
