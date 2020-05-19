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
	"github.com/volsch/gohimodel/resource"
	"github.com/volsch/gohipath/context"
	"testing"
)

func TestNewEvalContext(t *testing.T) {
	res := resource.NewDynamicResource("Patient")
	ctx := NewEvalContext(res, context.NewContext())
	assert.Same(t, res, ctx.ContextObj())
	v, _ := ctx.EnvVar("context")
	assert.Same(t, res, v)
	v, _ = ctx.EnvVar("resource")
	assert.Same(t, res, v)
	v, _ = ctx.EnvVar("rootResource")
	assert.Same(t, res, v)
}

func TestNewEvalContextWithRoot(t *testing.T) {
	res := resource.NewDynamicResource("Patient")
	root := resource.NewDynamicResource("Other")
	ctx := NewEvalContextWithRoot(res, root, context.NewContext())
	assert.Same(t, res, ctx.ContextObj())
	v, _ := ctx.EnvVar("context")
	assert.Same(t, res, v)
	v, _ = ctx.EnvVar("resource")
	assert.Same(t, res, v)
	v, _ = ctx.EnvVar("rootResource")
	assert.Same(t, root, v)
}

func TestNewEvalContextWithData(t *testing.T) {
	dt := datatype.NewString("test")
	res := resource.NewDynamicResource("Patient")
	ctx := NewEvalContextWithData(dt, res, context.NewContext())
	assert.Same(t, dt, ctx.ContextObj())
	v, _ := ctx.EnvVar("context")
	assert.Same(t, dt, v)
	v, _ = ctx.EnvVar("resource")
	assert.Same(t, res, v)
	v, _ = ctx.EnvVar("rootResource")
	assert.Same(t, res, v)
}

func TestNewEvalContextWithDataAndRoot(t *testing.T) {
	dt := datatype.NewString("test")
	res := resource.NewDynamicResource("Patient")
	root := resource.NewDynamicResource("Other")
	ctx := NewEvalContextWithDataAndRoot(dt, res, root, context.NewContext())
	assert.Same(t, dt, ctx.ContextObj())
	v, _ := ctx.EnvVar("context")
	assert.Same(t, dt, v)
	v, _ = ctx.EnvVar("resource")
	assert.Same(t, res, v)
	v, _ = ctx.EnvVar("rootResource")
	assert.Same(t, root, v)
}

func TestNewEvalContextWithQuantity(t *testing.T) {
	dt := datatype.NewQuantity(datatype.NewDecimalInt(10), nil, nil,
		datatype.UCUMSystemURI, datatype.NewString("mo"))
	res := resource.NewDynamicResource("Patient")
	root := resource.NewDynamicResource("Other")
	ctx := NewEvalContextWithDataAndRoot(dt, res, root, context.NewContext())
	assert.NotSame(t, dt, ctx.ContextObj())
	if assert.Implements(t, (*datatype.QuantityAccessor)(nil), ctx.ContextObj()) {
		q := ctx.ContextObj().(datatype.QuantityAccessor)
		assert.Nil(t, q.System(), "system must have been reset")
		if assert.NotNil(t, q.Code()) {
			assert.Equal(t, "month", q.Code().String())
		}
	}
	v, _ := ctx.EnvVar("context")
	assert.Same(t, ctx.ContextObj(), v)
	v, _ = ctx.EnvVar("resource")
	assert.Same(t, res, v)
	v, _ = ctx.EnvVar("rootResource")
	assert.Same(t, root, v)
}

func TestEvalContextEnvVar(t *testing.T) {
	res := resource.NewDynamicResource("Patient")
	ctx := NewEvalContext(res, context.NewContext())
	v, found := ctx.EnvVar("loinc")
	if assert.True(t, found, "LOINC system must be defined") {
		assert.Equal(t, datatype.LOINCSystemURI, v)
	}
}
