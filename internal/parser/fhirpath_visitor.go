// Code generated from FHIRPath.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // FHIRPath
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by FHIRPathParser.
type FHIRPathVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by FHIRPathParser#indexerExpression.
	VisitIndexerExpression(ctx *IndexerExpressionContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#polarityExpression.
	VisitPolarityExpression(ctx *PolarityExpressionContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#additiveExpression.
	VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#multiplicativeExpression.
	VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#unionExpression.
	VisitUnionExpression(ctx *UnionExpressionContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#orExpression.
	VisitOrExpression(ctx *OrExpressionContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#andExpression.
	VisitAndExpression(ctx *AndExpressionContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#membershipExpression.
	VisitMembershipExpression(ctx *MembershipExpressionContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#inequalityExpression.
	VisitInequalityExpression(ctx *InequalityExpressionContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#invocationExpression.
	VisitInvocationExpression(ctx *InvocationExpressionContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#equalityExpression.
	VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#impliesExpression.
	VisitImpliesExpression(ctx *ImpliesExpressionContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#termExpression.
	VisitTermExpression(ctx *TermExpressionContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#typeExpression.
	VisitTypeExpression(ctx *TypeExpressionContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#invocationTerm.
	VisitInvocationTerm(ctx *InvocationTermContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#literalTerm.
	VisitLiteralTerm(ctx *LiteralTermContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#externalConstantTerm.
	VisitExternalConstantTerm(ctx *ExternalConstantTermContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#parenthesizedTerm.
	VisitParenthesizedTerm(ctx *ParenthesizedTermContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#nullLiteral.
	VisitNullLiteral(ctx *NullLiteralContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#booleanLiteral.
	VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#stringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#numberLiteral.
	VisitNumberLiteral(ctx *NumberLiteralContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#dateLiteral.
	VisitDateLiteral(ctx *DateLiteralContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#dateTimeLiteral.
	VisitDateTimeLiteral(ctx *DateTimeLiteralContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#timeLiteral.
	VisitTimeLiteral(ctx *TimeLiteralContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#quantityLiteral.
	VisitQuantityLiteral(ctx *QuantityLiteralContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#externalConstant.
	VisitExternalConstant(ctx *ExternalConstantContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#memberInvocation.
	VisitMemberInvocation(ctx *MemberInvocationContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#functionInvocation.
	VisitFunctionInvocation(ctx *FunctionInvocationContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#thisInvocation.
	VisitThisInvocation(ctx *ThisInvocationContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#indexInvocation.
	VisitIndexInvocation(ctx *IndexInvocationContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#totalInvocation.
	VisitTotalInvocation(ctx *TotalInvocationContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#function.
	VisitFunction(ctx *FunctionContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#paramList.
	VisitParamList(ctx *ParamListContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#quantity.
	VisitQuantity(ctx *QuantityContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#unit.
	VisitUnit(ctx *UnitContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#dateTimePrecision.
	VisitDateTimePrecision(ctx *DateTimePrecisionContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#pluralDateTimePrecision.
	VisitPluralDateTimePrecision(ctx *PluralDateTimePrecisionContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#typeSpecifier.
	VisitTypeSpecifier(ctx *TypeSpecifierContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#qualifiedIdentifier.
	VisitQualifiedIdentifier(ctx *QualifiedIdentifierContext) interface{}

	// Visit a parse tree produced by FHIRPathParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}
}
