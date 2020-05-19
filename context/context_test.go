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

package context

import (
	"github.com/stretchr/testify/assert"
	"github.com/volsch/gohimodel/datatype"
	"testing"
)

func TestNewContext(t *testing.T) {
	c := NewContext()
	value, found := c.EnvVar("ucum")
	if assert.True(t, found, "ucum environment variable must be available") {
		assert.Equal(t, datatype.UCUMSystemURI, value)
	}
	value, found = c.EnvVar("sct")
	if assert.True(t, found, "sct environment variable must be available") {
		assert.Equal(t, datatype.SCTSystemURI, value)
	}
	value, found = c.EnvVar("loinc")
	if assert.True(t, found, "loinc environment variable must be available") {
		assert.Equal(t, datatype.LOINCSystemURI, value)
	}
	value, found = c.EnvVar("test")
	assert.False(t, found, "test environment variable must not be available")
}

func TestNewContextWithEnvVars(t *testing.T) {
	stringAccessor := datatype.NewString("Test")
	envVars := make(map[string]datatype.Accessor)
	envVars["test"] = stringAccessor

	c := NewContextWithEnvVars(envVars)
	assert.Len(t, envVars, 1, "passed map must not have changed")

	value, found := c.EnvVar("ucum")
	if assert.True(t, found, "ucum environment variable must be available") {
		assert.Equal(t, datatype.UCUMSystemURI, value)
	}
	value, found = c.EnvVar("test")
	if assert.True(t, found, "test environment variable must be available") {
		assert.Same(t, stringAccessor, value)
	}
}
