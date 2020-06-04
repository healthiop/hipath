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
	"github.com/volsch/gohipath/pathsys"
)

var functions = []pathsys.FunctionExecutor{
	// existence
	newEmptyFunction(),
	newExistsFunction(),
	newAllFunction(),
	newAllTrueFunction(),
	newAnyTrueFunction(),
	newAllFalseFunction(),
	newAnyFalseFunction(),
	newSubsetOfFunction(),
	newSupersetOfFunction(),
	newCountFunction(),
	newDistinctFunction(),
	newIsDistinctFunction(),
	// filtering and projection
	newWhereFunction(),
	newSelectFunction(),
	newRepeatFunction(),
	newOfTypeFunction(),
	// combining
	newUnionFunction(),
	newCombineFunction(),
	// aggregate
	newAggregateFunction(),
}

var functionsByName = createFunctionsByName(functions)

type FunctionInvocation struct {
	executor        pathsys.FunctionExecutor
	paramEvaluators []pathsys.Evaluator
}

func LookupFunctionInvocation(name string, paramEvaluators []pathsys.Evaluator) (*FunctionInvocation, error) {
	executor, found := functionsByName[name]
	if !found {
		return nil, fmt.Errorf("executor has not been defined: %s", name)
	}

	if len(paramEvaluators) < executor.MinParams() {
		return nil, fmt.Errorf("executor %s requires at least %d parameters", name, executor.MinParams())
	}
	if len(paramEvaluators) > executor.MaxParams() {
		return nil, fmt.Errorf("executor %s accepts at most %d parameters", name, executor.MaxParams())
	}

	return newFunctionInvocation(executor, paramEvaluators), nil
}

func newFunctionInvocation(executor pathsys.FunctionExecutor, argEvaluators []pathsys.Evaluator) *FunctionInvocation {
	return &FunctionInvocation{executor, argEvaluators}
}

func (f *FunctionInvocation) Evaluate(ctx pathsys.ContextAccessor, node interface{}, loop pathsys.Looper) (interface{}, error) {
	evaluatorParam := f.executor.EvaluatorParam()
	args := make([]interface{}, len(f.paramEvaluators))

	var loopEvaluator pathsys.Evaluator
	for pos, argEvaluator := range f.paramEvaluators {
		if evaluatorParam == pos {
			loopEvaluator = argEvaluator
		} else {
			if argEvaluator != nil {
				if arg, err := argEvaluator.Evaluate(ctx, node, loop); err != nil {
					return nil, fmt.Errorf("error in argument %d of executor invocation %s: %v",
						pos, f.executor.Name(), err)
				} else {
					args[pos] = arg
				}
			}
		}
	}

	if evaluatorParam >= 0 {
		loop = pathsys.NewLoop(loopEvaluator)
	}

	return f.executor.Execute(ctx, node, args, loop)
}

func createFunctionsByName(functions []pathsys.FunctionExecutor) map[string]pathsys.FunctionExecutor {
	functionsByName := make(map[string]pathsys.FunctionExecutor)
	for _, f := range functions {
		functionsByName[f.Name()] = f
	}
	return functionsByName
}
