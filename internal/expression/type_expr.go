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

type AsTypeExpression struct {
	exprEvaluator hipathsys.Evaluator
	fqName        hipathsys.FQTypeNameAccessor
}

func NewAsTypeExpression(exprEvaluator hipathsys.Evaluator, name string) (*AsTypeExpression, error) {
	fqName, err := hipathsys.ParseFQTypeName(name)
	if err != nil {
		return nil, err
	}

	return &AsTypeExpression{exprEvaluator, fqName}, nil
}

func (e *AsTypeExpression) Evaluate(ctx hipathsys.ContextAccessor, node interface{}, loop hipathsys.Looper) (interface{}, error) {
	value, err := e.exprEvaluator.Evaluate(ctx, node, loop)
	if err != nil {
		return nil, err
	}

	if value == nil {
		return nil, nil
	}

	item := unwrapCollection(value)
	if _, ok := item.(hipathsys.CollectionAccessor); ok {
		return nil, fmt.Errorf("as operator cannot be applied on a collection")
	}

	return hipathsys.CastModelType(ctx.ModelAdapter(), item, e.fqName)
}

type IsTypeExpression struct {
	exprEvaluator hipathsys.Evaluator
	fqName        hipathsys.FQTypeNameAccessor
}

func NewIsTypeExpression(exprEvaluator hipathsys.Evaluator, name string) (*IsTypeExpression, error) {
	fqName, err := hipathsys.ParseFQTypeName(name)
	if err != nil {
		return nil, err
	}

	return &IsTypeExpression{exprEvaluator, fqName}, nil
}

func (e *IsTypeExpression) Evaluate(ctx hipathsys.ContextAccessor, node interface{}, loop hipathsys.Looper) (interface{}, error) {
	value, err := e.exprEvaluator.Evaluate(ctx, node, loop)
	if err != nil {
		return nil, err
	}

	if value == nil {
		return nil, nil
	}

	item := unwrapCollection(value)
	if _, ok := item.(hipathsys.CollectionAccessor); ok {
		return nil, fmt.Errorf("is operator cannot be applied on a collection")
	}

	return hipathsys.BooleanOf(hipathsys.HasModelType(ctx.ModelAdapter(), item, e.fqName)), nil
}
