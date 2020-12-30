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
	"github.com/volsch/gohipath/pathsys"
)

type childrenFunction struct {
	pathsys.BaseFunction
}

var childrenFunc = &childrenFunction{
	BaseFunction: pathsys.NewBaseFunction("children", -1, 0, 0),
}

var childrenFuncInvocation = newFunctionInvocation(childrenFunc, []pathsys.Evaluator{})

func (f *childrenFunction) Execute(ctx pathsys.ContextAccessor, node interface{}, _ []interface{}, _ pathsys.Looper) (interface{}, error) {
	if node == nil {
		return nil, nil
	}

	if col, ok := node.(pathsys.CollectionAccessor); ok {
		count := col.Count()
		adapter := ctx.ModelAdapter()
		var children pathsys.CollectionModifier

		for i := 0; i < count; i++ {
			c := col.Get(i)
			if c != nil {
				ccol, err := adapter.Children(c)
				if err != nil {
					return nil, err
				}
				if ccol != nil && !ccol.Empty() {
					if children == nil {
						children = pathsys.NewCollection(adapter)
					}
					children.AddAll(ccol)
				}
			}
		}

		return children, nil
	}

	return ctx.ModelAdapter().Children(node)
}

type descendantsFunction struct {
	pathsys.BaseFunction
}

func newDescendantsFunction() *descendantsFunction {
	return &descendantsFunction{
		BaseFunction: pathsys.NewBaseFunction("descendants", -1, 0, 0),
	}
}

func (f *descendantsFunction) Execute(ctx pathsys.ContextAccessor, node interface{}, args []interface{}, _ pathsys.Looper) (interface{}, error) {
	return repeatFunc.Execute(ctx, node, args, pathsys.NewLoop(childrenFuncInvocation))
}
