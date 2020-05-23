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
		return pathsys.NewBoolean(true), nil
	}

	if col, ok := node.(pathsys.CollectionAccessor); ok {
		return pathsys.NewBoolean(col.Empty()), nil
	} else {
		return pathsys.NewBoolean(false), nil
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
		return pathsys.NewBoolean(false), nil
	}

	loopEvaluator := loop.Evaluator()
	col, ok := node.(pathsys.CollectionAccessor)
	if !ok {
		if loopEvaluator == nil {
			return pathsys.NewBoolean(true), nil
		}

		nc := ctx.NewCollection()
		nc.Add(node)
		col = nc
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
					return pathsys.NewBoolean(true), nil
				}
			}
		}
	}

	return pathsys.NewBoolean(found), nil
}
