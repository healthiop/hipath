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

package gohipath

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/healthiop/hipath/hipathsys"
	"github.com/healthiop/hipath/internal"
	"github.com/healthiop/hipath/internal/expression"
	"github.com/healthiop/hipath/internal/parser"
)

type Path struct {
	evaluator expression.CollectionExpression
}

func Compile(pathString string) (*Path, *hipathsys.Error) {
	errorItemCollection := internal.NewErrorItemCollection()
	errorListener := internal.NewErrorListener(errorItemCollection)

	is := antlr.NewInputStream(pathString)
	lexer := parser.NewFHIRPathLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewFHIRPathParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)

	v := internal.NewVisitor(errorItemCollection)
	res := p.Expression().Accept(v)

	if errorItemCollection.HasErrors() {
		return nil, hipathsys.NewError(
			"error when parsing path expression", errorItemCollection.Items())
	}

	return &Path{expression.NewCollectionExpression(res.(hipathsys.Evaluator))}, nil
}

func Execute(ctx hipathsys.ContextAccessor, pathString string, node interface{}) (hipathsys.ColAccessor, *hipathsys.Error) {
	path, err := Compile(pathString)
	if err != nil {
		return nil, err
	}

	return path.Execute(ctx, node)
}

func (p *Path) Execute(ctx hipathsys.ContextAccessor, node interface{}) (hipathsys.ColAccessor, *hipathsys.Error) {
	res, err := p.evaluator.Evaluate(ctx, node, nil)
	if err != nil {
		return nil, hipathsys.NewError(err.Error(), nil)
	}
	return res.(hipathsys.ColAccessor), nil
}
