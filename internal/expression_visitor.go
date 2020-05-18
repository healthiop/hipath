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
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/volsch/gohipath/internal/expression"
	"github.com/volsch/gohipath/internal/parser"
)

func (v *Visitor) VisitTermExpression(ctx *parser.TermExpressionContext) interface{} {
	return v.VisitFirstChild(ctx)
}

func (v *Visitor) VisitPolarityExpression(ctx *parser.PolarityExpressionContext) interface{} {
	return v.visitTree(ctx, 2, visitPolarityExpression)
}

func visitPolarityExpression(ctx antlr.ParserRuleContext, args []interface{}) (expression.Evaluator, error) {
	op := args[0].(string)
	evaluator := args[1].(expression.Evaluator)

	if op == "-" && evaluator != nil {
		evaluator = expression.NewNegatorExpression(evaluator.(expression.Evaluator))
	}
	return evaluator, nil
}

func (v *Visitor) VisitEqualityExpression(ctx *parser.EqualityExpressionContext) interface{} {
	return v.visitTree(ctx, 3, visitEqualityExpression)
}

func visitEqualityExpression(ctx antlr.ParserRuleContext, args []interface{}) (expression.Evaluator, error) {
	evalLeft := args[0].(expression.Evaluator)
	op := args[1].(string)
	evalRight := args[2].(expression.Evaluator)

	not := false
	equivalent := false
	switch op {
	case "=":
	case "!=":
		not = true
	case "~":
		equivalent = true
	case "!~":
		not = true
		equivalent = true
	default:
		return nil, fmt.Errorf("invalid equality operator: %s", op)
	}
	return expression.NewEqualityExpression(not, equivalent,
		evalLeft, evalRight), nil
}

func (v *Visitor) VisitUnionExpression(ctx *parser.UnionExpressionContext) interface{} {
	return v.visitTree(ctx, 3, visitUnionExpression)
}

func visitUnionExpression(ctx antlr.ParserRuleContext, args []interface{}) (expression.Evaluator, error) {
	evalLeft := args[0].(expression.Evaluator)
	evalRight := args[2].(expression.Evaluator)

	return expression.NewUnionExpression(evalLeft, evalRight), nil
}

func (v *Visitor) VisitIndexerExpression(ctx *parser.IndexerExpressionContext) interface{} {
	return v.visitTree(ctx, 4, visitIndexerExpression)
}

func visitIndexerExpression(ctx antlr.ParserRuleContext, args []interface{}) (expression.Evaluator, error) {
	exprEvaluator := args[0].(expression.Evaluator)
	indexEvaluator := args[2].(expression.Evaluator)

	return expression.NewIndexerExpression(exprEvaluator, indexEvaluator), nil
}
