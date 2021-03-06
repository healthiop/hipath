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
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/healthiop/hipath/hipathsys"
	"github.com/healthiop/hipath/internal/expression"
	"github.com/healthiop/hipath/internal/parser"
)

func (v *Visitor) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) interface{} {
	return v.visitTree(ctx, 3, visitArithmeticExpression)
}

func (v *Visitor) VisitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) interface{} {
	return v.visitTree(ctx, 3, visitArithmeticExpression)
}

func visitArithmeticExpression(_ antlr.ParserRuleContext, args []interface{}) (hipathsys.Evaluator, error) {
	leftEvaluator := args[0].(hipathsys.Evaluator)
	stringOp := args[1].(string)
	rightEvaluator := args[2].(hipathsys.Evaluator)

	var op hipathsys.ArithmeticOps
	switch stringOp {
	case "+":
		op = hipathsys.AdditionOp
	case "-":
		op = hipathsys.SubtractionOp
	case "*":
		op = hipathsys.MultiplicationOp
	case "/":
		op = hipathsys.DivisionOp
	case "div":
		op = hipathsys.DivOp
	case "mod":
		op = hipathsys.ModOp
	case "&":
		return expression.NewStringConcatExpression(leftEvaluator, rightEvaluator), nil
	default:
		return nil, fmt.Errorf("unsupported arithmetic oparator: %s", stringOp)
	}

	return expression.NewArithmeticExpression(leftEvaluator, op, rightEvaluator), nil
}
