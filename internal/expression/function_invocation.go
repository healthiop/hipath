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
	"github.com/volsch/gohimodel/datatype"
)

type invocationFunc func(ctx *EvalContext, obj datatype.Accessor, args []datatype.Accessor) (datatype.Accessor, error)

type functionDefinition struct {
	name      string
	function  invocationFunc
	minParams int
	maxParams int
}

var functionDefinitions = []*functionDefinition{
	// existence
	{"empty", emptyPathFunc, 0, 0},
	// combining
	{"union", unionPathFunc, 1, 1},
	{"combine", combinePathFunc, 1, 1},
}

var functionDefinitionsByName = createFunctionDefinitionsByName(functionDefinitions)

type FunctionInvocation struct {
	definition      *functionDefinition
	paramEvaluators []Evaluator
}

func LookupFunctionInvocation(name string, paramEvaluators []Evaluator) (*FunctionInvocation, error) {
	definition, found := functionDefinitionsByName[name]
	if !found {
		return nil, fmt.Errorf("function has not been defined: %s", name)
	}

	if len(paramEvaluators) < definition.minParams {
		return nil, fmt.Errorf("function %s requires at least %d parameters", name, definition.minParams)
	}
	if len(paramEvaluators) > definition.maxParams {
		return nil, fmt.Errorf("function %s accepts at most %d parameters", name, definition.maxParams)
	}

	return newFunctionInvocation(definition, paramEvaluators), nil
}

func newFunctionInvocation(definition *functionDefinition, argEvaluators []Evaluator) *FunctionInvocation {
	return &FunctionInvocation{definition, argEvaluators}
}

func (f *FunctionInvocation) Evaluate(ctx *EvalContext, obj datatype.Accessor) (datatype.Accessor, error) {
	args := make([]datatype.Accessor, len(f.paramEvaluators))
	for pos, argEvaluator := range f.paramEvaluators {
		if argEvaluator == nil {
			args[pos] = nil
		} else {
			if arg, err := argEvaluator.Evaluate(ctx, obj); err != nil {
				return nil, fmt.Errorf("error in argument %d of function invocation %s: %v",
					pos, f.definition.name, err)
			} else {
				args[pos] = arg
			}
		}
	}

	return f.definition.function(ctx, obj, args)
}

func createFunctionDefinitionsByName(definitions []*functionDefinition) map[string]*functionDefinition {
	definitionsByName := make(map[string]*functionDefinition)
	for _, definition := range definitions {
		definitionsByName[definition.name] = definition
	}
	return definitionsByName
}
