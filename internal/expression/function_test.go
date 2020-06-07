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
	"github.com/volsch/gohipath/pathsys"
	"testing"
)

var functionTests = []struct {
	name           string
	executor       pathsys.FunctionExecutor
	evaluatorParam int
	minParams      int
	maxParams      int
}{
	{"empty", newEmptyFunction(), -1, 0, 0},
	{"exists", newExistsFunction(), 0, 0, 1},
	{"all", newAllFunction(), 0, 1, 1},
	{"allTrue", newAllTrueFunction(), -1, 0, 0},
	{"anyTrue", newAnyTrueFunction(), -1, 0, 0},
	{"allFalse", newAllFalseFunction(), -1, 0, 0},
	{"anyFalse", newAnyFalseFunction(), -1, 0, 0},
	{"subsetOf", newSubsetOfFunction(), -1, 1, 1},
	{"supersetOf", newSupersetOfFunction(), -1, 1, 1},
	{"count", newCountFunction(), -1, 0, 0},
	{"distinct", newDistinctFunction(), -1, 0, 0},
	{"isDistinct", newIsDistinctFunction(), -1, 0, 0},
	{"where", newWhereFunction(), 0, 1, 1},
	{"select", newSelectFunction(), 0, 1, 1},
	{"repeat", newRepeatFunction(), 0, 1, 1},
	{"ofType", newOfTypeFunction(), -1, 1, 1},
	{"single", newSingleFunction(), -1, 0, 0},
	{"first", newFirstFunction(), -1, 0, 0},
	{"last", newLastFunction(), -1, 0, 0},
	{"tail", newTailFunction(), -1, 0, 0},
	{"skip", newSkipFunction(), -1, 1, 1},
	{"take", newTakeFunction(), -1, 1, 1},
	{"intersect", newIntersectFunction(), -1, 1, 1},
	{"exclude", newExcludeFunction(), -1, 1, 1},
	{"union", newUnionFunction(), -1, 1, 1},
	{"combine", newCombineFunction(), -1, 1, 1},
	{"iif", newIIfFunction(), -1, 2, 3},
	{"toBoolean", toBooleanFunc, -1, 0, 0},
	{"convertsToBoolean", newConvertsToBooleanFunction(), -1, 0, 0},
	{"toInteger", toIntegerFunc, -1, 0, 0},
	{"convertsToInteger", newConvertsToIntegerFunction(), -1, 0, 0},
	{"toDecimal", toDecimalFunc, -1, 0, 0},
	{"convertsToDecimal", newConvertsToDecimalFunction(), -1, 0, 0},
	{"toDate", toDateFunc, -1, 0, 0},
	{"convertsToDate", newConvertsToDateFunction(), -1, 0, 0},
	{"toDateTime", toDateTimeFunc, -1, 0, 0},
	{"convertsToDateTime", newConvertsToDateTimeFunction(), -1, 0, 0},
	{"toQuantity", toQuantityFunc, -1, 0, 1},
	{"convertsToQuantity", newConvertsToQuantityFunction(), -1, 0, 1},
	{"toString", toStringFunc, -1, 0, 0},
	{"convertsToString", newConvertsToStringFunction(), -1, 0, 0},
	{"toTime", toTimeFunc, -1, 0, 0},
	{"convertsToTime", newConvertsToTimeFunction(), -1, 0, 0},
	{"indexOf", newIndexOfFunction(), -1, 1, 1},
	{"aggregate", newAggregateFunction(), 0, 1, 2},
}

func TestFunctions(t *testing.T) {
	for _, tt := range functionTests {
		t.Run(tt.name, func(t *testing.T) {
			fe, found := functionsByName[tt.name]
			if found {
				assert.Equal(t, tt.executor, fe)
				assert.LessOrEqual(t, fe.EvaluatorParam(), tt.maxParams)
				assert.Equal(t, tt.evaluatorParam, fe.EvaluatorParam())
				assert.Equal(t, tt.minParams, fe.MinParams())
				assert.Equal(t, tt.maxParams, fe.MaxParams())
			} else {
				t.Errorf("executor %s has not been defined", tt.name)
			}
		})
	}
}
