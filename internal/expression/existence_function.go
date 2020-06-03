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

type emptyFunction struct {
	pathsys.BaseFunction
}

func newEmptyFunction() *emptyFunction {
	return &emptyFunction{
		BaseFunction: pathsys.NewBaseFunction("empty", -1, 0, 0),
	}
}

func (f *emptyFunction) Execute(ctx pathsys.ContextAccessor, node interface{}, args []interface{}, loop pathsys.Looper) (interface{}, error) {
	if node == nil {
		return pathsys.True, nil
	}

	if col, ok := node.(pathsys.CollectionAccessor); ok {
		return pathsys.BooleanOf(col.Empty()), nil
	} else {
		return pathsys.False, nil
	}
}

type existsFunction struct {
	pathsys.BaseFunction
}

func newExistsFunction() *existsFunction {
	return &existsFunction{
		BaseFunction: pathsys.NewBaseFunction("exists", 0, 0, 1),
	}
}

func (f *existsFunction) Execute(ctx pathsys.ContextAccessor, node interface{}, args []interface{}, loop pathsys.Looper) (interface{}, error) {
	if node == nil {
		return pathsys.False, nil
	}

	loopEvaluator := loop.Evaluator()
	col, ok := node.(pathsys.CollectionAccessor)
	if !ok {
		if loopEvaluator == nil {
			return pathsys.True, nil
		}
		col = ctx.NewCollectionWithItem(node)
	}
	count := col.Count()

	found := false
	if loopEvaluator == nil {
		found = count > 0
	} else {
		for i := 0; i < count; i++ {
			this := col.Get(i)
			loop.IncIndex(this)

			res, err := loopEvaluator.Evaluate(ctx, this, loop)
			if err != nil {
				return nil, err
			}
			if res != nil {
				if b, ok := res.(pathsys.BooleanAccessor); !ok {
					return nil, fmt.Errorf("filter expression must return boolean, but returned %T", res)
				} else if b.Bool() {
					return pathsys.True, nil
				}
			}
		}
	}

	return pathsys.BooleanOf(found), nil
}

type allFunction struct {
	pathsys.BaseFunction
}

func newAllFunction() *allFunction {
	return &allFunction{
		BaseFunction: pathsys.NewBaseFunction("all", 0, 1, 1),
	}
}

func (f *allFunction) Execute(ctx pathsys.ContextAccessor, node interface{}, args []interface{}, loop pathsys.Looper) (interface{}, error) {
	loopEvaluator := loop.Evaluator()
	col := wrapCollection(ctx, node)
	count := col.Count()
	for i := 0; i < count; i++ {
		this := col.Get(i)
		loop.IncIndex(this)

		res, err := loopEvaluator.Evaluate(ctx, this, loop)
		if err != nil {
			return nil, err
		}
		if b, ok := res.(pathsys.BooleanAccessor); !ok {
			return nil, fmt.Errorf("parameter expression must return boolean, but returned %T", res)
		} else if !b.Bool() {
			return pathsys.False, nil
		}
	}

	return pathsys.True, nil
}

type allAnyTrueFalseFunction struct {
	pathsys.BaseFunction
	all bool
	t   bool
}

func newAllAnyTrueFalseFunction(name string, all, t bool) *allAnyTrueFalseFunction {
	return &allAnyTrueFalseFunction{
		BaseFunction: pathsys.NewBaseFunction(name, -1, 0, 0),
		all:          all,
		t:            t,
	}
}

func newAllTrueFunction() *allAnyTrueFalseFunction {
	return newAllAnyTrueFalseFunction("allTrue", true, true)
}

func newAnyTrueFunction() *allAnyTrueFalseFunction {
	return newAllAnyTrueFalseFunction("anyTrue", false, true)
}

func newAllFalseFunction() *allAnyTrueFalseFunction {
	return newAllAnyTrueFalseFunction("allFalse", true, false)
}

func newAnyFalseFunction() *allAnyTrueFalseFunction {
	return newAllAnyTrueFalseFunction("anyFalse", false, false)
}

func (f *allAnyTrueFalseFunction) Execute(ctx pathsys.ContextAccessor, node interface{}, args []interface{}, loop pathsys.Looper) (interface{}, error) {
	col := wrapCollection(ctx, node)
	count := col.Count()
	for i := 0; i < count; i++ {
		this := col.Get(i)
		if b, ok := this.(pathsys.BooleanAccessor); !ok {
			return nil, fmt.Errorf("collection must contain only boolean values, but contains %T", this)
		} else if f.all && f.t != b.Bool() {
			return pathsys.False, nil
		} else if !f.all && f.t == b.Bool() {
			return pathsys.True, nil
		}
	}

	if f.all {
		return pathsys.True, nil
	}
	return pathsys.False, nil
}

type subsetOfFunction struct {
	pathsys.BaseFunction
}

func newSubsetOfFunction() *subsetOfFunction {
	return &subsetOfFunction{
		BaseFunction: pathsys.NewBaseFunction("subsetOf", -1, 1, 1),
	}
}

func (f *subsetOfFunction) Execute(ctx pathsys.ContextAccessor, node interface{}, args []interface{}, loop pathsys.Looper) (interface{}, error) {
	col := wrapCollection(ctx, node)
	count := col.Count()
	if count > 0 {
		otherCol := wrapCollection(ctx, args[0])
		for i := 0; i < count; i++ {
			if !otherCol.Contains(col.Get(i)) {
				return pathsys.False, nil
			}
		}
	}
	return pathsys.True, nil
}

type supersetOfFunction struct {
	pathsys.BaseFunction
}

func newSupersetOfFunction() *supersetOfFunction {
	return &supersetOfFunction{
		BaseFunction: pathsys.NewBaseFunction("supersetOf", -1, 1, 1),
	}
}

func (f *supersetOfFunction) Execute(ctx pathsys.ContextAccessor, node interface{}, args []interface{}, loop pathsys.Looper) (interface{}, error) {
	otherCol := wrapCollection(ctx, args[0])
	count := otherCol.Count()
	if count > 0 {
		col := wrapCollection(ctx, node)
		for i := 0; i < count; i++ {
			if !col.Contains(otherCol.Get(i)) {
				return pathsys.False, nil
			}
		}
	}
	return pathsys.True, nil
}

type countFunction struct {
	pathsys.BaseFunction
}

func newCountFunction() *countFunction {
	return &countFunction{
		BaseFunction: pathsys.NewBaseFunction("count", -1, 0, 0),
	}
}

func (f *countFunction) Execute(ctx pathsys.ContextAccessor, node interface{}, args []interface{}, loop pathsys.Looper) (interface{}, error) {
	col := wrapCollection(ctx, node)
	return pathsys.NewInteger(int32(col.Count())), nil
}

type distinctFunction struct {
	pathsys.BaseFunction
}

func newDistinctFunction() *distinctFunction {
	return &distinctFunction{
		BaseFunction: pathsys.NewBaseFunction("distinct", -1, 0, 0),
	}
}

func (f *distinctFunction) Execute(ctx pathsys.ContextAccessor, node interface{}, args []interface{}, loop pathsys.Looper) (interface{}, error) {
	col := wrapCollection(ctx, node)
	if col.Count() < 2 {
		return col, nil
	}

	res := ctx.NewCollection()
	res.AddAllUnique(col)
	return res, nil
}

type isDistinctFunction struct {
	pathsys.BaseFunction
}

func newIsDistinctFunction() *isDistinctFunction {
	return &isDistinctFunction{
		BaseFunction: pathsys.NewBaseFunction("isDistinct", -1, 0, 0),
	}
}

func (f *isDistinctFunction) Execute(ctx pathsys.ContextAccessor, node interface{}, args []interface{}, loop pathsys.Looper) (interface{}, error) {
	col := wrapCollection(ctx, node)
	if col.Empty() {
		return nil, nil
	}

	if col.Count() == 1 {
		return pathsys.True, nil
	}

	res := ctx.NewCollection()
	res.AddAllUnique(col)
	return pathsys.BooleanOf(col.Count() == res.Count()), nil
}
