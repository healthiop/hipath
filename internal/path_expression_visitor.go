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
	"github.com/volsch/gohipath/internal/parser"
)

func (v *PathVisitor) VisitIndexerExpression(ctx *parser.IndexerExpressionContext) interface{} {
	return v.VisitFirstChild(ctx)
}

func (v *PathVisitor) VisitPolarityExpression(ctx *parser.PolarityExpressionContext) interface{} {
	return v.VisitFirstChild(ctx)
}

func (v *PathVisitor) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) interface{} {
	return v.VisitFirstChild(ctx)
}

func (v *PathVisitor) VisitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) interface{} {
	return v.VisitFirstChild(ctx)
}

func (v *PathVisitor) VisitUnionExpression(ctx *parser.UnionExpressionContext) interface{} {
	return v.VisitFirstChild(ctx)
}

func (v *PathVisitor) VisitOrExpression(ctx *parser.OrExpressionContext) interface{} {
	return v.VisitFirstChild(ctx)
}

func (v *PathVisitor) VisitAndExpression(ctx *parser.AndExpressionContext) interface{} {
	return v.VisitFirstChild(ctx)
}

func (v *PathVisitor) VisitMembershipExpression(ctx *parser.MembershipExpressionContext) interface{} {
	return v.VisitFirstChild(ctx)
}

func (v *PathVisitor) VisitInequalityExpression(ctx *parser.InequalityExpressionContext) interface{} {
	return v.VisitFirstChild(ctx)
}

func (v *PathVisitor) VisitInvocationExpression(ctx *parser.InvocationExpressionContext) interface{} {
	return v.VisitFirstChild(ctx)
}

func (v *PathVisitor) VisitEqualityExpression(ctx *parser.EqualityExpressionContext) interface{} {
	return v.VisitFirstChild(ctx)
}

func (v *PathVisitor) VisitImpliesExpression(ctx *parser.ImpliesExpressionContext) interface{} {
	return v.VisitFirstChild(ctx)
}

func (v *PathVisitor) VisitTermExpression(ctx *parser.TermExpressionContext) interface{} {
	return v.VisitFirstChild(ctx)
}

func (v *PathVisitor) VisitTypeExpression(ctx *parser.TypeExpressionContext) interface{} {
	return v.VisitFirstChild(ctx)
}
