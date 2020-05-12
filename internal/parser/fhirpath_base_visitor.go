// Code generated from FHIRPath.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // FHIRPath
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseFHIRPathVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseFHIRPathVisitor) VisitIndexerExpression(ctx *IndexerExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitPolarityExpression(ctx *PolarityExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitUnionExpression(ctx *UnionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitOrExpression(ctx *OrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitAndExpression(ctx *AndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitMembershipExpression(ctx *MembershipExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitInequalityExpression(ctx *InequalityExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitInvocationExpression(ctx *InvocationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitEqualityExpression(ctx *EqualityExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitImpliesExpression(ctx *ImpliesExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitTermExpression(ctx *TermExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitTypeExpression(ctx *TypeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitInvocationTerm(ctx *InvocationTermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitLiteralTerm(ctx *LiteralTermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitExternalConstantTerm(ctx *ExternalConstantTermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitParenthesizedTerm(ctx *ParenthesizedTermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitNullLiteral(ctx *NullLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitNumberLiteral(ctx *NumberLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitDateLiteral(ctx *DateLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitDateTimeLiteral(ctx *DateTimeLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitTimeLiteral(ctx *TimeLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitQuantityLiteral(ctx *QuantityLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitExternalConstant(ctx *ExternalConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitMemberInvocation(ctx *MemberInvocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitFunctionInvocation(ctx *FunctionInvocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitThisInvocation(ctx *ThisInvocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitIndexInvocation(ctx *IndexInvocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitTotalInvocation(ctx *TotalInvocationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitFunction(ctx *FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitParamList(ctx *ParamListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitQuantity(ctx *QuantityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitUnit(ctx *UnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitDateTimePrecision(ctx *DateTimePrecisionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitPluralDateTimePrecision(ctx *PluralDateTimePrecisionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitTypeSpecifier(ctx *TypeSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitQualifiedIdentifier(ctx *QualifiedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFHIRPathVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}
