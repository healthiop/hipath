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

type whereFunction struct {
	hipathsys.BaseFunction
}

func newWhereFunction() *whereFunction {
	return &whereFunction{
		BaseFunction: hipathsys.NewBaseFunction("where", 0, 1, 1),
	}
}

func (f *whereFunction) Execute(ctx hipathsys.ContextAccessor, node interface{}, _ []interface{}, loop hipathsys.Looper) (interface{}, error) {
	col := wrapCollection(ctx, node)
	count := col.Count()
	if count == 0 {
		return nil, nil
	}

	var filtered hipathsys.ColModifier
	loopEvaluator := loop.Evaluator()
	for i := 0; i < count; i++ {
		this := col.Get(i)
		loop.IncIndex(this)

		res, err := loopEvaluator.Evaluate(ctx, this, loop)
		if err != nil {
			return nil, err
		}
		if res != nil {
			if b, ok := res.(hipathsys.BooleanAccessor); !ok {
				return nil, fmt.Errorf("filter expression must return boolean, but returned %T", res)
			} else if b.Bool() {
				if filtered == nil {
					filtered = ctx.NewCol()
				}
				filtered.Add(this)
			}
		}
	}

	return filtered, nil
}

type selectFunction struct {
	hipathsys.BaseFunction
}

func newSelectFunction() *selectFunction {
	return &selectFunction{
		BaseFunction: hipathsys.NewBaseFunction("select", 0, 1, 1),
	}
}

func (f *selectFunction) Execute(ctx hipathsys.ContextAccessor, node interface{}, _ []interface{}, loop hipathsys.Looper) (interface{}, error) {
	col := wrapCollection(ctx, node)
	count := col.Count()
	if count == 0 {
		return nil, nil
	}

	var projected hipathsys.ColModifier
	loopEvaluator := loop.Evaluator()
	for i := 0; i < count; i++ {
		this := col.Get(i)
		loop.IncIndex(this)

		res, err := loopEvaluator.Evaluate(ctx, this, loop)
		if err != nil {
			return nil, err
		}
		if res != nil {
			if projected == nil {
				projected = ctx.NewCol()
			}

			if c, ok := res.(hipathsys.ColAccessor); ok {
				projected.AddAll(c)
			} else {
				projected.Add(res)
			}
		}
	}

	return projected, nil
}

type repeatFunction struct {
	hipathsys.BaseFunction
}

var repeatFunc = &repeatFunction{
	BaseFunction: hipathsys.NewBaseFunction("repeat", 0, 1, 1),
}

func (f *repeatFunction) Execute(ctx hipathsys.ContextAccessor, node interface{}, _ []interface{}, loop hipathsys.Looper) (interface{}, error) {
	projected := ctx.NewCol()
	err := repeat(ctx, node, loop, projected)

	if err != nil || projected.Empty() {
		projected = nil
	}

	return projected, err
}

func repeat(ctx hipathsys.ContextAccessor, node interface{}, loop hipathsys.Looper, projected hipathsys.ColModifier) error {
	col := wrapCollection(ctx, node)
	count := col.Count()
	if count == 0 {
		return nil
	}

	loopEvaluator := loop.Evaluator()
	for i := 0; i < count; i++ {
		this := col.Get(i)
		loop.IncIndex(this)

		res, err := loopEvaluator.Evaluate(ctx, this, loop)
		if err != nil {
			return err
		}
		if res != nil {
			err := repeatRecursively(ctx, res, loop, projected)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func repeatRecursively(ctx hipathsys.ContextAccessor, node interface{}, loop hipathsys.Looper, projected hipathsys.ColModifier) error {
	if col, ok := node.(hipathsys.ColAccessor); ok {
		count := col.Count()
		for i := 0; i < count; i++ {
			n := col.Get(i)
			if n != nil {
				added := projected.AddUnique(n)
				if added {
					err := repeat(ctx, n, hipathsys.NewLoopWithIndex(
						loop.Evaluator(), i), projected)
					if err != nil {
						return err
					}
				}
			}
		}
	} else if projected.AddUnique(node) {
		err := repeat(ctx, node, hipathsys.NewLoop(
			loop.Evaluator()), projected)
		if err != nil {
			return err
		}
	}
	return nil
}

type ofTypeFunction struct {
	hipathsys.BaseFunction
}

func newOfTypeFunction() *ofTypeFunction {
	return &ofTypeFunction{
		BaseFunction: hipathsys.NewBaseFunction("ofType", -1, 1, 1),
	}
}

func (f *ofTypeFunction) Execute(ctx hipathsys.ContextAccessor, node interface{}, args []interface{}, _ hipathsys.Looper) (interface{}, error) {
	var typeSpec hipathsys.StringAccessor
	var ok bool
	if typeSpec, ok = unwrapCollection(args[0]).(hipathsys.StringAccessor); !ok {
		return nil, fmt.Errorf("not a valid type specifier: %T", args[0])
	}

	var typeName hipathsys.FQTypeNameAccessor
	var err error
	if typeName, err = hipathsys.ParseFQTypeName(typeSpec.String()); err != nil {
		return nil, fmt.Errorf("not a valid type specifier: %s", typeSpec)
	}

	col := wrapCollection(ctx, node)
	count := col.Count()
	if count == 0 {
		return nil, nil
	}

	var filtered hipathsys.ColModifier
	adapter := ctx.ModelAdapter()
	for i := 0; i < count; i++ {
		n := col.Get(i)
		if hipathsys.HasModelType(adapter, n, typeName) {
			if filtered == nil {
				filtered = ctx.NewCol()
			}
			filtered.Add(n)
		}
	}

	return filtered, nil
}
