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

func (v *PathVisitor) VisitExternalConstant(ctx *parser.ExternalConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *PathVisitor) VisitMemberInvocation(ctx *parser.MemberInvocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *PathVisitor) VisitFunctionInvocation(ctx *parser.FunctionInvocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *PathVisitor) VisitThisInvocation(ctx *parser.ThisInvocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *PathVisitor) VisitIndexInvocation(ctx *parser.IndexInvocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *PathVisitor) VisitTotalInvocation(ctx *parser.TotalInvocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *PathVisitor) VisitFunction(ctx *parser.FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *PathVisitor) VisitParamList(ctx *parser.ParamListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *PathVisitor) VisitDateTimePrecision(ctx *parser.DateTimePrecisionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *PathVisitor) VisitPluralDateTimePrecision(ctx *parser.PluralDateTimePrecisionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *PathVisitor) VisitTypeSpecifier(ctx *parser.TypeSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *PathVisitor) VisitQualifiedIdentifier(ctx *parser.QualifiedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *PathVisitor) VisitIdentifier(ctx *parser.IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}
