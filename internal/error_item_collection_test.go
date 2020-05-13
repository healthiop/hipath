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

package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorCollection(t *testing.T) {
	c := NewErrorItemCollection()
	c.AddError(72, 3, "error 1")
	c.AddError(87, 23, "error 2")

	assert.True(t, c.HasErrors(), "collection must have errors")
	if assert.Len(t, c.Items(), 2) {
		err := c.Items()[0]
		assert.Equal(t, 72, err.Line())
		assert.Equal(t, 3, err.Column())
		assert.Equal(t, "error 1", err.Msg())

		err = c.Items()[1]
		assert.Equal(t, 87, err.Line())
		assert.Equal(t, 23, err.Column())
		assert.Equal(t, "error 2", err.Msg())
	}
}

func TestErrorCollectionEmpty(t *testing.T) {
	c := NewErrorItemCollection()
	assert.False(t, c.HasErrors(), "collection must have no errors")
	assert.Len(t, c.Items(), 0)
}
