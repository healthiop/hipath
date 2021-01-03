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
	"github.com/volsch/gohipath/internal/expression"
	"github.com/volsch/gohipath/internal/parser"
	"github.com/volsch/gohipath/pathsys"
)

func (v *Visitor) VisitNullLiteral(*parser.NullLiteralContext) interface{} {
	return expression.NewEmptyLiteral()
}

func (v *Visitor) VisitBooleanLiteral(ctx *parser.BooleanLiteralContext) interface{} {
	return v.visit(ctx, visitBooleanLiteral)
}

func visitBooleanLiteral(ctx antlr.ParserRuleContext) (pathsys.Evaluator, error) {
	return expression.ParseBooleanLiteral(ctx.GetText())
}

func (v *Visitor) VisitStringLiteral(ctx *parser.StringLiteralContext) interface{} {
	return expression.ParseStringLiteral(ctx.GetText())
}

func (v *Visitor) VisitNumberLiteral(ctx *parser.NumberLiteralContext) interface{} {
	return v.visit(ctx, visitNumberLiteral)
}

func visitNumberLiteral(ctx antlr.ParserRuleContext) (pathsys.Evaluator, error) {
	return expression.ParseNumberLiteral(ctx.GetText())
}

func (v *Visitor) VisitDateLiteral(ctx *parser.DateLiteralContext) interface{} {
	return v.visit(ctx, visitDateLiteral)
}

func visitDateLiteral(ctx antlr.ParserRuleContext) (pathsys.Evaluator, error) {
	return expression.ParseDateLiteral(ctx.GetText())
}

func (v *Visitor) VisitDateTimeLiteral(ctx *parser.DateTimeLiteralContext) interface{} {
	return v.visit(ctx, visitDateTimeLiteral)
}

func visitDateTimeLiteral(ctx antlr.ParserRuleContext) (pathsys.Evaluator, error) {
	return expression.ParseDateTimeLiteral(ctx.GetText())
}

func (v *Visitor) VisitTimeLiteral(ctx *parser.TimeLiteralContext) interface{} {
	return v.visit(ctx, visitTimeLiteral)
}

func visitTimeLiteral(ctx antlr.ParserRuleContext) (pathsys.Evaluator, error) {
	return expression.ParseTimeLiteral(ctx.GetText())
}

func (v *Visitor) VisitQuantity(ctx *parser.QuantityContext) interface{} {
	return v.visitTree(ctx, 2, visitQuantity)
}

func visitQuantity(_ antlr.ParserRuleContext, args []interface{}) (pathsys.Evaluator, error) {
	number := args[0].(string)
	unit := args[1].(string)
	return expression.ParseQuantityLiteral(number, unit)
}

func (v *Visitor) VisitQuantityLiteral(ctx *parser.QuantityLiteralContext) interface{} {
	return v.VisitFirstChild(ctx)
}

func (v *Visitor) VisitUnit(ctx *parser.UnitContext) interface{} {
	return ctx.GetText()
}
