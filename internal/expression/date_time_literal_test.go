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
	"time"
)

func TestFullDateTimeLiteral(t *testing.T) {
	executor, err := ParseDateTimeLiteral("@2014-03-25T14:30:15.559Z")

	assert.NoError(t, err, "no error expected")
	assert.NotNil(t, executor, "executor expected")
	if executor != nil {
		accessor := executor.Execute(nil)
		assert.NotNil(t, accessor, "accessor expected")
		if assert.Implements(t, (*datatype.DateTimeAccessor)(nil), accessor) {
			expected := time.Date(2014, 3, 25, 14, 30, 15, 559000000, time.UTC)
			actual := accessor.(datatype.DateTimeAccessor).Value()
			assert.True(t, expected.Equal(actual), "expected %d, got %d",
				expected.UnixNano(), actual.UnixNano())
		}
	}
}

func TestFluentDateTimeLiteral(t *testing.T) {
	executor, err := ParseDateTimeLiteral("@2014-03-25T")

	assert.NoError(t, err, "no error expected")
	assert.NotNil(t, executor, "executor expected")
	if executor != nil {
		accessor := executor.Execute(nil)
		assert.NotNil(t, accessor, "accessor expected")
		if assert.Implements(t, (*datatype.DateTimeAccessor)(nil), accessor) {
			expected := time.Date(2014, 3, 25, 0, 0, 0, 0, time.Local)
			actual := accessor.(datatype.DateTimeAccessor).Value()
			assert.True(t, expected.Equal(actual), "expected %d, got %d",
				expected.UnixNano(), actual.UnixNano())
		}
	}
}

func TestDateTimeLiteralInvalid(t *testing.T) {
	executor, err := ParseDateTimeLiteral("4-01-25T14:30:14.559Z")

	assert.Error(t, err, "error expected")
	assert.Nil(t, executor, "no executor expected")
}

func TestDateTimeLiteralInvalidStartToken(t *testing.T) {
	executor, err := ParseDateTimeLiteral("2014-01-25T14:30:14.559Z")

	assert.Error(t, err, "error expected")
	assert.Nil(t, executor, "no executor expected")
}

func TestDateTimeLiteralInvalidStart(t *testing.T) {
	executor, err := ParseDateTimeLiteral("@")

	assert.Error(t, err, "error expected")
	assert.Nil(t, executor, "no executor expected")
}
