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
	"reflect"
	"runtime"
	"testing"
)

var functionTests = []struct {
	name      string
	function  invocationFunc
	minParams int
	maxParams int
}{
	{"empty", emptyPathFunc, 0, 0},
	{"union", unionPathFunc, 1, 1},
	{"combine", combinePathFunc, 1, 1},
}

func TestFunctions(t *testing.T) {
	for _, tt := range functionTests {
		t.Run(tt.name, func(t *testing.T) {
			def, found := functionDefinitionsByName[tt.name]
			if found {
				assert.Equal(t, runtime.FuncForPC(reflect.ValueOf(tt.function).Pointer()).Name(),
					runtime.FuncForPC(reflect.ValueOf(def.function).Pointer()).Name())
				assert.Equal(t, tt.minParams, def.minParams)
				assert.Equal(t, tt.maxParams, def.maxParams)
			} else {
				t.Errorf("function %s has not been defined", tt.name)
			}
		})
	}
}
