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

package internal

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/healthiop/hipath/hipathsys"
	"github.com/healthiop/hipath/internal/expression"
	"github.com/healthiop/hipath/internal/parser"
)

func (v *Visitor) VisitFunctionInvocation(ctx *parser.FunctionInvocationContext) interface{} {
	return v.VisitFirstChild(ctx)
}

func (v *Visitor) VisitFunction(ctx *parser.FunctionContext) interface{} {
	return v.visitTree(ctx, 3, visitFunction)
}

func visitFunction(_ antlr.ParserRuleContext, args []interface{}) (hipathsys.Evaluator, error) {
	name := args[0].(string)

	var paramEvaluators []hipathsys.Evaluator
	if len(args) < 4 {
		paramEvaluators = []hipathsys.Evaluator{}
	} else if name == "as" || name == "is" {
		typeSpec := args[2].(string)
		paramEvaluators = []hipathsys.Evaluator{expression.NewRawStringLiteral(typeSpec)}
	} else {
		paramList := args[2].([]interface{})
		// commas need to removed from argument list
		paramEvaluators = make([]hipathsys.Evaluator, (len(paramList)+1)/2)
		for pos, param := range paramList {
			if pos%2 == 0 {
				paramEvaluators[pos/2] = param.(hipathsys.Evaluator)
			}
		}
	}

	return expression.LookupFunctionInvocation(expression.ExtractIdentifier(name), paramEvaluators)
}

func (v *Visitor) VisitParamList(ctx *parser.ParamListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitThisInvocation(*parser.ThisInvocationContext) interface{} {
	return expression.NewThisInvocation()
}

func (v *Visitor) VisitIndexInvocation(*parser.IndexInvocationContext) interface{} {
	return expression.NewIndexInvocation()
}

func (v *Visitor) VisitTotalInvocation(*parser.TotalInvocationContext) interface{} {
	return expression.NewTotalInvocation()
}

func (v *Visitor) VisitMemberInvocation(ctx *parser.MemberInvocationContext) interface{} {
	return v.visitTree(ctx, 1, visitMemberInvocation)
}

func visitMemberInvocation(_ antlr.ParserRuleContext, args []interface{}) (hipathsys.Evaluator, error) {
	name := args[0].(string)
	return expression.NewMemberInvocation(expression.ExtractIdentifier(name)), nil
}
