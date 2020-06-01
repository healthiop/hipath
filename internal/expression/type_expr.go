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

type AsTypeExpression struct {
	exprEvaluator pathsys.Evaluator
	fqName        pathsys.FQTypeNameAccessor
}

func NewAsTypeExpression(exprEvaluator pathsys.Evaluator, name string) (*AsTypeExpression, error) {
	fqName, err := pathsys.ParseFQTypeName(name)
	if err != nil {
		return nil, err
	}

	return &AsTypeExpression{exprEvaluator, fqName}, nil
}

func (e *AsTypeExpression) Evaluate(ctx pathsys.ContextAccessor, node interface{}, loop pathsys.Looper) (interface{}, error) {
	value, err := e.exprEvaluator.Evaluate(ctx, node, loop)
	if err != nil {
		return nil, err
	}

	if value == nil {
		return nil, nil
	}

	item := unwrapCollection(value)
	if _, ok := item.(pathsys.CollectionAccessor); ok {
		return nil, fmt.Errorf("as operator cannot be applied on a collection")
	}

	if pathsys.HasModelType(ctx.ModelAdapter(), item, e.fqName) {
		return value, nil
	}
	return nil, nil
}

type IsTypeExpression struct {
	exprEvaluator pathsys.Evaluator
	fqName        pathsys.FQTypeNameAccessor
}

func NewIsTypeExpression(exprEvaluator pathsys.Evaluator, name string) (*IsTypeExpression, error) {
	fqName, err := pathsys.ParseFQTypeName(name)
	if err != nil {
		return nil, err
	}

	return &IsTypeExpression{exprEvaluator, fqName}, nil
}

func (e *IsTypeExpression) Evaluate(ctx pathsys.ContextAccessor, node interface{}, loop pathsys.Looper) (interface{}, error) {
	value, err := e.exprEvaluator.Evaluate(ctx, node, loop)
	if err != nil {
		return nil, err
	}

	if value == nil {
		return nil, nil
	}

	item := unwrapCollection(value)
	if _, ok := item.(pathsys.CollectionAccessor); ok {
		return nil, fmt.Errorf("is operator cannot be applied on a collection")
	}

	return pathsys.BooleanOf(pathsys.HasModelType(ctx.ModelAdapter(), item, e.fqName)), nil
}
