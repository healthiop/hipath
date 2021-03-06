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

func TestThisInvocation(t *testing.T) {
	loop := hipathsys.NewLoop(nil)
	this := hipathsys.NewString("test")
	loop.IncIndex(this)
	evaluator := NewThisInvocation()
	res, err := evaluator.Evaluate(nil, nil, loop)
	assert.NoError(t, err, "error expected")
	assert.Same(t, this, res)
}

func TestThisInvocationOutsideLoop(t *testing.T) {
	evaluator := NewThisInvocation()
	res, err := evaluator.Evaluate(nil, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no res expected")
}

func TestIndexInvocation(t *testing.T) {
	loop := hipathsys.NewLoop(nil)
	loop.IncIndex(hipathsys.NewInteger(10))
	this := hipathsys.NewString("test")
	loop.IncIndex(this)
	evaluator := NewIndexInvocation()
	res, err := evaluator.Evaluate(nil, nil, loop)
	assert.NoError(t, err, "error expected")
	assert.Equal(t, hipathsys.NewInteger(int32(1)), res)
}

func TestIndexInvocationOutsideLoop(t *testing.T) {
	evaluator := NewIndexInvocation()
	res, err := evaluator.Evaluate(nil, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no res expected")
}

func TestTotalInvocation(t *testing.T) {
	loop := hipathsys.NewLoop(nil)
	this := hipathsys.NewString("test")
	loop.IncIndex(this)
	total := hipathsys.NewInteger(20)
	loop.SetTotal(total)
	evaluator := NewTotalInvocation()
	res, err := evaluator.Evaluate(nil, nil, loop)
	assert.NoError(t, err, "error expected")
	assert.Same(t, total, res)
}

func TestTotalInvocationOutsideLoop(t *testing.T) {
	evaluator := NewTotalInvocation()
	res, err := evaluator.Evaluate(nil, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no res expected")
}
