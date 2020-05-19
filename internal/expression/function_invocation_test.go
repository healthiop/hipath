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
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/volsch/gohimodel/datatype"
	"github.com/volsch/gohimodel/resource"
	"github.com/volsch/gohipath/context"
	"testing"
)

func TestFunctionInvocationArgs(t *testing.T) {
	definition := &functionDefinition{"test", testInvocationFunctionArgs, 0, 100}

	ctx := NewEvalContext(resource.NewDynamicResource("Patient"), context.NewContext())
	e := newFunctionInvocation(definition, []Evaluator{
		ParseStringLiteral("test1"), nil, ParseStringLiteral("test2")})

	accessor, err := e.Evaluate(ctx, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*datatype.CollectionAccessor)(nil), accessor) {
		c := accessor.(datatype.CollectionAccessor)
		if assert.Equal(t, 3, c.Count()) {
			assert.Equal(t, datatype.NewString("test1"), c.Get(0))
			assert.Nil(t, c.Get(1))
			assert.Equal(t, datatype.NewString("test2"), c.Get(2))
		}
	}
}

func TestFunctionInvocationArgsError(t *testing.T) {
	definition := &functionDefinition{"test", testInvocationFunctionArgs, 0, 100}

	ctx := NewEvalContext(resource.NewDynamicResource("Patient"), context.NewContext())
	e := newFunctionInvocation(definition, []Evaluator{
		ParseStringLiteral("test1"), ParseExtConstantTerm("xxx"), ParseStringLiteral("test2")})

	accessor, err := e.Evaluate(ctx, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, accessor, "no result expected")
}

func TestFunctionInvocationError(t *testing.T) {
	definition := &functionDefinition{"test", testInvocationFunctionErr, 0, 100}

	ctx := NewEvalContext(resource.NewDynamicResource("Patient"), context.NewContext())
	e := newFunctionInvocation(definition, []Evaluator{})

	accessor, err := e.Evaluate(ctx, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, accessor, "no result expected")
}

func TestLookupFunctionInvocationNotFound(t *testing.T) {
	fi, err := LookupFunctionInvocation("test", make([]Evaluator, 0))
	assert.EqualError(t, err, "function has not been defined: test", "error expected")
	assert.Nil(t, fi, "no function invocation expected")
}

func TestLookupFunctionInvocationTooLessArgs(t *testing.T) {
	fi, err := LookupFunctionInvocation("union", make([]Evaluator, 0))
	assert.EqualError(t, err, "function union requires at least 1 parameters", "error expected")
	assert.Nil(t, fi, "no function invocation expected")
}

func TestLookupFunctionInvocationTooManyArgs(t *testing.T) {
	fi, err := LookupFunctionInvocation("union", make([]Evaluator, 2))
	assert.EqualError(t, err, "function union accepts at most 1 parameters", "error expected")
	assert.Nil(t, fi, "no function invocation expected")
}

func testInvocationFunctionArgs(ctx *EvalContext, obj datatype.Accessor,
	args []datatype.Accessor) (datatype.Accessor, error) {
	c := datatype.NewCollectionUndefined()
	for _, a := range args {
		c.Add(a)
	}
	return c, nil
}

func testInvocationFunctionErr(ctx *EvalContext, obj datatype.Accessor,
	args []datatype.Accessor) (datatype.Accessor, error) {
	return nil, fmt.Errorf("an error occurred")
}
