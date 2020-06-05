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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringOfEmpty(t *testing.T) {
	assert.Nil(t, StringOf(""))
}

func TestStringOf(t *testing.T) {
	assert.Equal(t, NewString("test"), StringOf("test"))
}

func TestStringSource(t *testing.T) {
	o := NewStringWithSource("xyz", "abc")
	assert.Equal(t, "abc", o.Source())
}

func TestStringDataType(t *testing.T) {
	o := NewString("Test")
	dataType := o.DataType()
	assert.Equal(t, StringDataType, dataType)
}

func TestStringTypeInfo(t *testing.T) {
	o := NewString("Test")
	i := o.TypeInfo()
	if assert.NotNil(t, i, "type info expected") {
		assert.Equal(t, "System.String", i.String())
		if assert.NotNil(t, i.FQBaseName(), "base name expected") {
			assert.Equal(t, "System.Any", i.FQBaseName().String())
		}
	}
}

func TestStringValue(t *testing.T) {
	o := NewString("Test")
	value := o.String()
	assert.Equal(t, "Test", value)
}

func TestStringEqualNil(t *testing.T) {
	assert.Equal(t, false, NewString("").Equal(nil))
}

func TestStringEqualTypeDiffers(t *testing.T) {
	assert.Equal(t, false, NewString("").Equal(newAccessorMock()))
	assert.Equal(t, false, NewString("").Equivalent(newAccessorMock()))
}

func TestStringEqualStringTypeDiffers(t *testing.T) {
	assert.Equal(t, false, NewString("").Equal(False))
	assert.Equal(t, false, NewString("").Equivalent(False))
}

func TestStringEqualEqual(t *testing.T) {
	assert.Equal(t, true, NewString("test").Equal(NewString("test")))
	assert.Equal(t, true, NewString("test").Equivalent(NewString("test")))
}

func TestStringEquivalent(t *testing.T) {
	assert.Equal(t, true, NewString("TEST\nvalue").Equivalent(NewString("test VALUE")))
}

func TestStringEqualNotEqual(t *testing.T) {
	assert.Equal(t, false, NewString("test1").Equal(NewString("test2")))
	assert.Equal(t, false, NewString("test1").Equivalent(NewString("test2")))
}

func TestStringCompareEqual(t *testing.T) {
	res, status := NewString("test1").Compare(NewString("test1"))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, 0, res)
}

func TestStringCompareEqualTypeDiffers(t *testing.T) {
	res, status := NewString("test1").Compare(False)
	assert.Equal(t, Inconvertible, status)
	assert.Equal(t, -1, res)
}

func TestStringCompareLessThan(t *testing.T) {
	res, status := NewString("test1").Compare(NewString("test5"))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, -1, res)
}

func TestStringCompareGreaterThan(t *testing.T) {
	res, status := NewString("test5").Compare(NewString("test1"))
	assert.Equal(t, Evaluated, status)
	assert.Equal(t, 1, res)
}
