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
	"github.com/volsch/gohipath/pathsys"
)

func (v *Visitor) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) interface{} {
	return v.visitTree(ctx, 3, visitArithmeticExpression)
}

func (v *Visitor) VisitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) interface{} {
	return v.visitTree(ctx, 3, visitArithmeticExpression)
}

func visitArithmeticExpression(_ antlr.ParserRuleContext, args []interface{}) (pathsys.Evaluator, error) {
	leftEvaluator := args[0].(pathsys.Evaluator)
	stringOp := args[1].(string)
	rightEvaluator := args[2].(pathsys.Evaluator)

	var op pathsys.ArithmeticOps
	switch stringOp {
	case "+":
		op = pathsys.AdditionOp
	case "-":
		op = pathsys.SubtractionOp
	case "*":
		op = pathsys.MultiplicationOp
	case "/":
		op = pathsys.DivisionOp
	case "div":
		op = pathsys.DivOp
	case "mod":
		op = pathsys.ModOp
	case "&":
		return expression.NewStringConcatExpression(leftEvaluator, rightEvaluator), nil
	default:
		return nil, fmt.Errorf("unsupported arithmetic oparator: %s", stringOp)
	}

	return expression.NewArithmeticExpression(leftEvaluator, op, rightEvaluator), nil
}
