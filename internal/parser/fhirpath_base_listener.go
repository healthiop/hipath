// Code generated from fhirpath.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // fhirpath

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasefhirpathListener is a complete listener for a parse tree produced by fhirpathParser.
type BasefhirpathListener struct{}

var _ fhirpathListener = &BasefhirpathListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasefhirpathListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasefhirpathListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasefhirpathListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasefhirpathListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterIndexerExpression is called when production indexerExpression is entered.
func (s *BasefhirpathListener) EnterIndexerExpression(ctx *IndexerExpressionContext) {}

// ExitIndexerExpression is called when production indexerExpression is exited.
func (s *BasefhirpathListener) ExitIndexerExpression(ctx *IndexerExpressionContext) {}

// EnterPolarityExpression is called when production polarityExpression is entered.
func (s *BasefhirpathListener) EnterPolarityExpression(ctx *PolarityExpressionContext) {}

// ExitPolarityExpression is called when production polarityExpression is exited.
func (s *BasefhirpathListener) ExitPolarityExpression(ctx *PolarityExpressionContext) {}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *BasefhirpathListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *BasefhirpathListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *BasefhirpathListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *BasefhirpathListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// EnterUnionExpression is called when production unionExpression is entered.
func (s *BasefhirpathListener) EnterUnionExpression(ctx *UnionExpressionContext) {}

// ExitUnionExpression is called when production unionExpression is exited.
func (s *BasefhirpathListener) ExitUnionExpression(ctx *UnionExpressionContext) {}

// EnterOrExpression is called when production orExpression is entered.
func (s *BasefhirpathListener) EnterOrExpression(ctx *OrExpressionContext) {}

// ExitOrExpression is called when production orExpression is exited.
func (s *BasefhirpathListener) ExitOrExpression(ctx *OrExpressionContext) {}

// EnterAndExpression is called when production andExpression is entered.
func (s *BasefhirpathListener) EnterAndExpression(ctx *AndExpressionContext) {}

// ExitAndExpression is called when production andExpression is exited.
func (s *BasefhirpathListener) ExitAndExpression(ctx *AndExpressionContext) {}

// EnterMembershipExpression is called when production membershipExpression is entered.
func (s *BasefhirpathListener) EnterMembershipExpression(ctx *MembershipExpressionContext) {}

// ExitMembershipExpression is called when production membershipExpression is exited.
func (s *BasefhirpathListener) ExitMembershipExpression(ctx *MembershipExpressionContext) {}

// EnterInequalityExpression is called when production inequalityExpression is entered.
func (s *BasefhirpathListener) EnterInequalityExpression(ctx *InequalityExpressionContext) {}

// ExitInequalityExpression is called when production inequalityExpression is exited.
func (s *BasefhirpathListener) ExitInequalityExpression(ctx *InequalityExpressionContext) {}

// EnterInvocationExpression is called when production invocationExpression is entered.
func (s *BasefhirpathListener) EnterInvocationExpression(ctx *InvocationExpressionContext) {}

// ExitInvocationExpression is called when production invocationExpression is exited.
func (s *BasefhirpathListener) ExitInvocationExpression(ctx *InvocationExpressionContext) {}

// EnterEqualityExpression is called when production equalityExpression is entered.
func (s *BasefhirpathListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production equalityExpression is exited.
func (s *BasefhirpathListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterImpliesExpression is called when production impliesExpression is entered.
func (s *BasefhirpathListener) EnterImpliesExpression(ctx *ImpliesExpressionContext) {}

// ExitImpliesExpression is called when production impliesExpression is exited.
func (s *BasefhirpathListener) ExitImpliesExpression(ctx *ImpliesExpressionContext) {}

// EnterTermExpression is called when production termExpression is entered.
func (s *BasefhirpathListener) EnterTermExpression(ctx *TermExpressionContext) {}

// ExitTermExpression is called when production termExpression is exited.
func (s *BasefhirpathListener) ExitTermExpression(ctx *TermExpressionContext) {}

// EnterTypeExpression is called when production typeExpression is entered.
func (s *BasefhirpathListener) EnterTypeExpression(ctx *TypeExpressionContext) {}

// ExitTypeExpression is called when production typeExpression is exited.
func (s *BasefhirpathListener) ExitTypeExpression(ctx *TypeExpressionContext) {}

// EnterInvocationTerm is called when production invocationTerm is entered.
func (s *BasefhirpathListener) EnterInvocationTerm(ctx *InvocationTermContext) {}

// ExitInvocationTerm is called when production invocationTerm is exited.
func (s *BasefhirpathListener) ExitInvocationTerm(ctx *InvocationTermContext) {}

// EnterLiteralTerm is called when production literalTerm is entered.
func (s *BasefhirpathListener) EnterLiteralTerm(ctx *LiteralTermContext) {}

// ExitLiteralTerm is called when production literalTerm is exited.
func (s *BasefhirpathListener) ExitLiteralTerm(ctx *LiteralTermContext) {}

// EnterExternalConstantTerm is called when production externalConstantTerm is entered.
func (s *BasefhirpathListener) EnterExternalConstantTerm(ctx *ExternalConstantTermContext) {}

// ExitExternalConstantTerm is called when production externalConstantTerm is exited.
func (s *BasefhirpathListener) ExitExternalConstantTerm(ctx *ExternalConstantTermContext) {}

// EnterParenthesizedTerm is called when production parenthesizedTerm is entered.
func (s *BasefhirpathListener) EnterParenthesizedTerm(ctx *ParenthesizedTermContext) {}

// ExitParenthesizedTerm is called when production parenthesizedTerm is exited.
func (s *BasefhirpathListener) ExitParenthesizedTerm(ctx *ParenthesizedTermContext) {}

// EnterNullLiteral is called when production nullLiteral is entered.
func (s *BasefhirpathListener) EnterNullLiteral(ctx *NullLiteralContext) {}

// ExitNullLiteral is called when production nullLiteral is exited.
func (s *BasefhirpathListener) ExitNullLiteral(ctx *NullLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BasefhirpathListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BasefhirpathListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BasefhirpathListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BasefhirpathListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterNumberLiteral is called when production numberLiteral is entered.
func (s *BasefhirpathListener) EnterNumberLiteral(ctx *NumberLiteralContext) {}

// ExitNumberLiteral is called when production numberLiteral is exited.
func (s *BasefhirpathListener) ExitNumberLiteral(ctx *NumberLiteralContext) {}

// EnterDateLiteral is called when production dateLiteral is entered.
func (s *BasefhirpathListener) EnterDateLiteral(ctx *DateLiteralContext) {}

// ExitDateLiteral is called when production dateLiteral is exited.
func (s *BasefhirpathListener) ExitDateLiteral(ctx *DateLiteralContext) {}

// EnterDateTimeLiteral is called when production dateTimeLiteral is entered.
func (s *BasefhirpathListener) EnterDateTimeLiteral(ctx *DateTimeLiteralContext) {}

// ExitDateTimeLiteral is called when production dateTimeLiteral is exited.
func (s *BasefhirpathListener) ExitDateTimeLiteral(ctx *DateTimeLiteralContext) {}

// EnterTimeLiteral is called when production timeLiteral is entered.
func (s *BasefhirpathListener) EnterTimeLiteral(ctx *TimeLiteralContext) {}

// ExitTimeLiteral is called when production timeLiteral is exited.
func (s *BasefhirpathListener) ExitTimeLiteral(ctx *TimeLiteralContext) {}

// EnterQuantityLiteral is called when production quantityLiteral is entered.
func (s *BasefhirpathListener) EnterQuantityLiteral(ctx *QuantityLiteralContext) {}

// ExitQuantityLiteral is called when production quantityLiteral is exited.
func (s *BasefhirpathListener) ExitQuantityLiteral(ctx *QuantityLiteralContext) {}

// EnterExternalConstant is called when production externalConstant is entered.
func (s *BasefhirpathListener) EnterExternalConstant(ctx *ExternalConstantContext) {}

// ExitExternalConstant is called when production externalConstant is exited.
func (s *BasefhirpathListener) ExitExternalConstant(ctx *ExternalConstantContext) {}

// EnterMemberInvocation is called when production memberInvocation is entered.
func (s *BasefhirpathListener) EnterMemberInvocation(ctx *MemberInvocationContext) {}

// ExitMemberInvocation is called when production memberInvocation is exited.
func (s *BasefhirpathListener) ExitMemberInvocation(ctx *MemberInvocationContext) {}

// EnterFunctionInvocation is called when production functionInvocation is entered.
func (s *BasefhirpathListener) EnterFunctionInvocation(ctx *FunctionInvocationContext) {}

// ExitFunctionInvocation is called when production functionInvocation is exited.
func (s *BasefhirpathListener) ExitFunctionInvocation(ctx *FunctionInvocationContext) {}

// EnterThisInvocation is called when production thisInvocation is entered.
func (s *BasefhirpathListener) EnterThisInvocation(ctx *ThisInvocationContext) {}

// ExitThisInvocation is called when production thisInvocation is exited.
func (s *BasefhirpathListener) ExitThisInvocation(ctx *ThisInvocationContext) {}

// EnterIndexInvocation is called when production indexInvocation is entered.
func (s *BasefhirpathListener) EnterIndexInvocation(ctx *IndexInvocationContext) {}

// ExitIndexInvocation is called when production indexInvocation is exited.
func (s *BasefhirpathListener) ExitIndexInvocation(ctx *IndexInvocationContext) {}

// EnterTotalInvocation is called when production totalInvocation is entered.
func (s *BasefhirpathListener) EnterTotalInvocation(ctx *TotalInvocationContext) {}

// ExitTotalInvocation is called when production totalInvocation is exited.
func (s *BasefhirpathListener) ExitTotalInvocation(ctx *TotalInvocationContext) {}

// EnterFunction is called when production function is entered.
func (s *BasefhirpathListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BasefhirpathListener) ExitFunction(ctx *FunctionContext) {}

// EnterParamList is called when production paramList is entered.
func (s *BasefhirpathListener) EnterParamList(ctx *ParamListContext) {}

// ExitParamList is called when production paramList is exited.
func (s *BasefhirpathListener) ExitParamList(ctx *ParamListContext) {}

// EnterQuantity is called when production quantity is entered.
func (s *BasefhirpathListener) EnterQuantity(ctx *QuantityContext) {}

// ExitQuantity is called when production quantity is exited.
func (s *BasefhirpathListener) ExitQuantity(ctx *QuantityContext) {}

// EnterUnit is called when production unit is entered.
func (s *BasefhirpathListener) EnterUnit(ctx *UnitContext) {}

// ExitUnit is called when production unit is exited.
func (s *BasefhirpathListener) ExitUnit(ctx *UnitContext) {}

// EnterDateTimePrecision is called when production dateTimePrecision is entered.
func (s *BasefhirpathListener) EnterDateTimePrecision(ctx *DateTimePrecisionContext) {}

// ExitDateTimePrecision is called when production dateTimePrecision is exited.
func (s *BasefhirpathListener) ExitDateTimePrecision(ctx *DateTimePrecisionContext) {}

// EnterPluralDateTimePrecision is called when production pluralDateTimePrecision is entered.
func (s *BasefhirpathListener) EnterPluralDateTimePrecision(ctx *PluralDateTimePrecisionContext) {}

// ExitPluralDateTimePrecision is called when production pluralDateTimePrecision is exited.
func (s *BasefhirpathListener) ExitPluralDateTimePrecision(ctx *PluralDateTimePrecisionContext) {}

// EnterTypeSpecifier is called when production typeSpecifier is entered.
func (s *BasefhirpathListener) EnterTypeSpecifier(ctx *TypeSpecifierContext) {}

// ExitTypeSpecifier is called when production typeSpecifier is exited.
func (s *BasefhirpathListener) ExitTypeSpecifier(ctx *TypeSpecifierContext) {}

// EnterQualifiedIdentifier is called when production qualifiedIdentifier is entered.
func (s *BasefhirpathListener) EnterQualifiedIdentifier(ctx *QualifiedIdentifierContext) {}

// ExitQualifiedIdentifier is called when production qualifiedIdentifier is exited.
func (s *BasefhirpathListener) ExitQualifiedIdentifier(ctx *QualifiedIdentifierContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BasefhirpathListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BasefhirpathListener) ExitIdentifier(ctx *IdentifierContext) {}
