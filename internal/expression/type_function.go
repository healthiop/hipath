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

type asFunction struct {
	hipathsys.BaseFunction
}

func newAsFunction() *asFunction {
	return &asFunction{
		BaseFunction: hipathsys.NewBaseFunction("as", -1, 1, 1),
	}
}

func (f *asFunction) Execute(ctx hipathsys.ContextAccessor, node interface{}, args []interface{}, _ hipathsys.Looper) (interface{}, error) {
	name, err := stringNode(args[0])
	if name == nil || err != nil {
		return nil, err
	}

	fqName, err := hipathsys.ParseFQTypeName(name.String())
	if err != nil {
		return nil, err
	}

	item := unwrapCollection(node)
	if _, ok := item.(hipathsys.CollectionAccessor); ok {
		return nil, fmt.Errorf("as function cannot be applied on a collection")
	}

	return hipathsys.CastModelType(ctx.ModelAdapter(), item, fqName)
}

type isFunction struct {
	hipathsys.BaseFunction
}

func newIsFunction() *isFunction {
	return &isFunction{
		BaseFunction: hipathsys.NewBaseFunction("is", -1, 1, 1),
	}
}

func (f *isFunction) Execute(ctx hipathsys.ContextAccessor, node interface{}, args []interface{}, _ hipathsys.Looper) (interface{}, error) {
	name, err := stringNode(args[0])
	if name == nil || err != nil {
		return nil, err
	}

	fqName, err := hipathsys.ParseFQTypeName(name.String())
	if err != nil {
		return nil, err
	}

	item := unwrapCollection(node)
	if _, ok := item.(hipathsys.CollectionAccessor); ok {
		return nil, fmt.Errorf("is function cannot be applied on a collection")
	}

	return hipathsys.BooleanOf(hipathsys.HasModelType(ctx.ModelAdapter(), item, fqName)), nil
}
