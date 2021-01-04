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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFullTimeLiteral(t *testing.T) {
	evaluator, err := ParseTimeLiteral("@T14:30:15.559")

	assert.NoError(t, err, "no error expected")
	assert.NotNil(t, evaluator, "evaluator expected")
	if evaluator != nil {
		node, err := evaluator.Evaluate(nil, nil, nil)
		assert.NoError(t, err, "no error expected")
		assert.NotNil(t, node, "res expected")
		if assert.Implements(t, (*hipathsys.TimeAccessor)(nil), node) {
			ta := node.(hipathsys.TimeAccessor)
			assert.Equal(t, 14, ta.Hour())
			assert.Equal(t, 30, ta.Minute())
			assert.Equal(t, 15, ta.Second())
			assert.Equal(t, 559000000, ta.Nanosecond())
		}
	}
}

func TestFluentTimeLiteral(t *testing.T) {
	evaluator, err := ParseTimeLiteral("@T14:30")

	assert.NoError(t, err, "no error expected")
	assert.NotNil(t, evaluator, "evaluator expected")
	if evaluator != nil {
		res, err := evaluator.Evaluate(nil, nil, nil)
		assert.NoError(t, err, "no error expected")
		assert.NotNil(t, res, "res expected")
		if assert.Implements(t, (*hipathsys.TimeAccessor)(nil), res) {
			ta := res.(hipathsys.TimeAccessor)
			assert.Equal(t, 14, ta.Hour())
			assert.Equal(t, 30, ta.Minute())
			assert.Equal(t, 0, ta.Second())
			assert.Equal(t, 0, ta.Nanosecond())
		}
	}
}

func TestTimeLiteralInvalid(t *testing.T) {
	evaluator, err := ParseTimeLiteral("@T:30:14.559")

	assert.Error(t, err, "error expected")
	assert.Nil(t, evaluator, "no evaluator expected")
}

func TestTimeLiteralInvalidStartToken(t *testing.T) {
	evaluator, err := ParseTimeLiteral("@14:30:14.559")

	assert.Error(t, err, "error expected")
	assert.Nil(t, evaluator, "no evaluator expected")
}

func TestTimeLiteralInvalidStart(t *testing.T) {
	evaluator, err := ParseTimeLiteral("@T")

	assert.Error(t, err, "error expected")
	assert.Nil(t, evaluator, "no evaluator expected")
}
