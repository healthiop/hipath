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

package internal

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/volsch/gohipath/internal/expression"
	"github.com/volsch/gohipath/internal/parser"
)

type PathVisitor struct {
	parser.BaseFHIRPathVisitor
	errorItemCollection *PathErrorItemCollection
}

func NewPathVisitor(errorItemCollection *PathErrorItemCollection) *PathVisitor {
	v := new(PathVisitor)
	v.errorItemCollection = errorItemCollection
	return v
}

func (v *PathVisitor) AddError(ctx antlr.ParserRuleContext, msg string) expression.Executor {
	v.errorItemCollection.AddError(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), msg)
	return nil
}

func (v *PathVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	c := node.GetChildCount()
	if c == 0 {
		return nil
	}

	r := make([]interface{}, c)
	for pos, child := range node.GetChildren() {
		r[pos] = v.evalNode(child)
	}
	return r
}

func (v *PathVisitor) VisitFirstChild(node antlr.RuleNode) interface{} {
	if node.GetChildCount() == 0 {
		return nil
	}
	return v.evalNode(node.GetChild(0))
}

func (v *PathVisitor) evalNode(node antlr.Tree) interface{} {
	switch n := node.(type) {
	case antlr.RuleNode:
		return n.Accept(v)
	case antlr.TerminalNode:
		return n.GetText()
	}
	return nil
}
