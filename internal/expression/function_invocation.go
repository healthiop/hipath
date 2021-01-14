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
	"fmt"
	"github.com/healthiop/hipath/hipathsys"
)

var functions = []hipathsys.FunctionExecutor{
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
	repeatFunc,
	newOfTypeFunction(),
	// sub-setting
	newSingleFunction(),
	newFirstFunction(),
	newLastFunction(),
	newTailFunction(),
	newSkipFunction(),
	newTakeFunction(),
	newIntersectFunction(),
	newExcludeFunction(),
	// combining
	newUnionFunction(),
	newCombineFunction(),
	// conversion
	newIIfFunction(),
	toBooleanFunc,
	newConvertsToBooleanFunction(),
	toIntegerFunc,
	newConvertsToIntegerFunction(),
	toDateFunc,
	newConvertsToDateFunction(),
	toDateTimeFunc,
	newConvertsToDateTimeFunction(),
	toDecimalFunc,
	newConvertsToDecimalFunction(),
	toQuantityFunc,
	newConvertsToQuantityFunction(),
	toStringFunc,
	newConvertsToStringFunction(),
	toTimeFunc,
	newConvertsToTimeFunction(),
	// string manipulation
	newIndexOfFunction(),
	newSubstringFunction(),
	newStartsWithFunction(),
	newEndsWithFunction(),
	newContainsFunction(),
	newUpperFunction(),
	newLowerFunction(),
	newReplaceFunction(),
	newMatchesFunction(),
	newReplaceMatchesFunction(),
	newLengthFunction(),
	newToCharsFunction(),
	// math
	newAbsFunction(),
	newCeilingFunction(),
	newExpFunction(),
	newFloorFunction(),
	newLnFunction(),
	newLogFunction(),
	newPowerFunction(),
	newRoundFunction(),
	newSqrtFunction(),
	newTruncateFunction(),
	// tree navigation
	childrenFunc,
	newDescendantsFunction(),
	// utility
	newTraceFunction(),
	newNowFunction(),
	newTimeOfDayFunction(),
	newTodayFunction(),
	// type
	newAsFunction(),
	newIsFunction(),
	// aggregate
	newAggregateFunction(),
}

var functionsByName = createFunctionsByName(functions)

type FunctionInvocation struct {
	executor        hipathsys.FunctionExecutor
	paramEvaluators []hipathsys.Evaluator
}

func LookupFunctionInvocation(name string, paramEvaluators []hipathsys.Evaluator) (*FunctionInvocation, error) {
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

func newFunctionInvocation(executor hipathsys.FunctionExecutor, argEvaluators []hipathsys.Evaluator) *FunctionInvocation {
	return &FunctionInvocation{executor, argEvaluators}
}

func (f *FunctionInvocation) Evaluate(ctx hipathsys.ContextAccessor, node interface{}, loop hipathsys.Looper) (interface{}, error) {
	var args []interface{}
	ac := len(f.paramEvaluators)
	if ac == 0 {
		args = nil
	} else {
		evaluatorParam := f.executor.EvaluatorParam()
		args = make([]interface{}, len(f.paramEvaluators))

		var loopEvaluator hipathsys.Evaluator
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
			loop = hipathsys.NewLoop(loopEvaluator)
		}
	}

	return f.executor.Execute(ctx, node, args, loop)
}

func createFunctionsByName(functions []hipathsys.FunctionExecutor) map[string]hipathsys.FunctionExecutor {
	functionsByName := make(map[string]hipathsys.FunctionExecutor)
	for _, f := range functions {
		functionsByName[f.Name()] = f
	}
	return functionsByName
}
