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

type singleFunction struct {
	hipathsys.BaseFunction
}

func newSingleFunction() *singleFunction {
	return &singleFunction{
		BaseFunction: hipathsys.NewBaseFunction("single", -1, 0, 0),
	}
}

func (f *singleFunction) Execute(ctx hipathsys.ContextAccessor, node interface{}, _ []interface{}, _ hipathsys.Looper) (interface{}, error) {
	col := wrapCollection(ctx, node)
	count := col.Count()
	if count == 0 {
		return nil, nil
	}
	if count > 1 {
		return nil, fmt.Errorf("expected collection with one item: %d", count)
	}
	return col.Get(0), nil
}

type firstFunction struct {
	hipathsys.BaseFunction
}

func newFirstFunction() *firstFunction {
	return &firstFunction{
		BaseFunction: hipathsys.NewBaseFunction("first", -1, 0, 0),
	}
}

func (f *firstFunction) Execute(ctx hipathsys.ContextAccessor, node interface{}, _ []interface{}, _ hipathsys.Looper) (interface{}, error) {
	col := wrapCollection(ctx, node)
	if col.Empty() {
		return nil, nil
	}
	return col.Get(0), nil
}

type lastFunction struct {
	hipathsys.BaseFunction
}

func newLastFunction() *lastFunction {
	return &lastFunction{
		BaseFunction: hipathsys.NewBaseFunction("last", -1, 0, 0),
	}
}

func (f *lastFunction) Execute(ctx hipathsys.ContextAccessor, node interface{}, _ []interface{}, _ hipathsys.Looper) (interface{}, error) {
	col := wrapCollection(ctx, node)
	count := col.Count()
	if count == 0 {
		return nil, nil
	}
	return col.Get(count - 1), nil
}

type tailFunction struct {
	hipathsys.BaseFunction
}

func newTailFunction() *tailFunction {
	return &tailFunction{
		BaseFunction: hipathsys.NewBaseFunction("tail", -1, 0, 0),
	}
}

func (f *tailFunction) Execute(ctx hipathsys.ContextAccessor, node interface{}, _ []interface{}, _ hipathsys.Looper) (interface{}, error) {
	col := wrapCollection(ctx, node)
	count := col.Count()
	if count < 2 {
		return nil, nil
	}

	tail := ctx.NewCollection()
	for i := 1; i < count; i++ {
		tail.Add(col.Get(i))
	}

	return tail, nil
}

type skipFunction struct {
	hipathsys.BaseFunction
}

func newSkipFunction() *skipFunction {
	return &skipFunction{
		BaseFunction: hipathsys.NewBaseFunction("skip", -1, 1, 1),
	}
}

func (f *skipFunction) Execute(ctx hipathsys.ContextAccessor, node interface{}, args []interface{}, _ hipathsys.Looper) (interface{}, error) {
	var num int
	if n, ok := unwrapCollection(args[0]).(hipathsys.NumberAccessor); !ok {
		return nil, fmt.Errorf("argument must be an integer: %T", args[0])
	} else {
		num = int(n.Int())
	}

	col := wrapCollection(ctx, node)
	if num <= 0 {
		return col, nil
	}

	count := col.Count()
	if count <= num {
		return nil, nil
	}

	res := ctx.NewCollection()
	for i := num; i < count; i++ {
		res.Add(col.Get(i))
	}
	return res, nil
}

type takeFunction struct {
	hipathsys.BaseFunction
}

func newTakeFunction() *takeFunction {
	return &takeFunction{
		BaseFunction: hipathsys.NewBaseFunction("take", -1, 1, 1),
	}
}

func (f *takeFunction) Execute(ctx hipathsys.ContextAccessor, node interface{}, args []interface{}, _ hipathsys.Looper) (interface{}, error) {
	var num int
	if n, ok := unwrapCollection(args[0]).(hipathsys.NumberAccessor); !ok {
		return nil, fmt.Errorf("argument must be an integer: %T", args[0])
	} else {
		num = int(n.Int())
	}
	if num <= 0 {
		return nil, nil
	}

	col := wrapCollection(ctx, node)
	count := col.Count()
	if count == 0 {
		return nil, nil
	}
	if count <= num {
		return col, nil
	}

	res := ctx.NewCollection()
	for i := 0; i < num; i++ {
		res.Add(col.Get(i))
	}
	return res, nil
}

type intersectFunction struct {
	hipathsys.BaseFunction
}

func newIntersectFunction() *intersectFunction {
	return &intersectFunction{
		BaseFunction: hipathsys.NewBaseFunction("intersect", -1, 1, 1),
	}
}

func (f *intersectFunction) Execute(ctx hipathsys.ContextAccessor, node interface{}, args []interface{}, _ hipathsys.Looper) (interface{}, error) {
	other := wrapCollection(ctx, args[0])
	if other.Empty() {
		return nil, nil
	}

	col := wrapCollection(ctx, node)
	if col.Empty() {
		return nil, nil
	}

	if col.Count() > other.Count() {
		x := col
		col = other
		other = x
	}
	count := col.Count()

	res := ctx.NewCollection()
	for i := 0; i < count; i++ {
		n := col.Get(i)
		if other.Contains(n) {
			res.AddUnique(n)
		}
	}
	return res, nil
}

type excludeFunction struct {
	hipathsys.BaseFunction
}

func newExcludeFunction() *excludeFunction {
	return &excludeFunction{
		BaseFunction: hipathsys.NewBaseFunction("exclude", -1, 1, 1),
	}
}

func (f *excludeFunction) Execute(ctx hipathsys.ContextAccessor, node interface{}, args []interface{}, _ hipathsys.Looper) (interface{}, error) {
	other := wrapCollection(ctx, args[0])
	col := wrapCollection(ctx, node)
	count := col.Count()
	if count == 0 {
		return nil, nil
	}
	if other.Empty() {
		return col, nil
	}

	res := ctx.NewCollection()
	for i := 0; i < count; i++ {
		n := col.Get(i)
		if !other.Contains(n) {
			res.Add(n)
		}
	}
	return res, nil
}
