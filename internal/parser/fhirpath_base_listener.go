// Code generated from FHIRPath.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // FHIRPath

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFHIRPathListener is a complete listener for a parse tree produced by FHIRPathParser.
type BaseFHIRPathListener struct{}

var _ FHIRPathListener = &BaseFHIRPathListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFHIRPathListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFHIRPathListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFHIRPathListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFHIRPathListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterIndexerExpression is called when production indexerExpression is entered.
func (s *BaseFHIRPathListener) EnterIndexerExpression(ctx *IndexerExpressionContext) {}

// ExitIndexerExpression is called when production indexerExpression is exited.
func (s *BaseFHIRPathListener) ExitIndexerExpression(ctx *IndexerExpressionContext) {}

// EnterPolarityExpression is called when production polarityExpression is entered.
func (s *BaseFHIRPathListener) EnterPolarityExpression(ctx *PolarityExpressionContext) {}

// ExitPolarityExpression is called when production polarityExpression is exited.
func (s *BaseFHIRPathListener) ExitPolarityExpression(ctx *PolarityExpressionContext) {}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *BaseFHIRPathListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *BaseFHIRPathListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *BaseFHIRPathListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *BaseFHIRPathListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// EnterUnionExpression is called when production unionExpression is entered.
func (s *BaseFHIRPathListener) EnterUnionExpression(ctx *UnionExpressionContext) {}

// ExitUnionExpression is called when production unionExpression is exited.
func (s *BaseFHIRPathListener) ExitUnionExpression(ctx *UnionExpressionContext) {}

// EnterOrExpression is called when production orExpression is entered.
func (s *BaseFHIRPathListener) EnterOrExpression(ctx *OrExpressionContext) {}

// ExitOrExpression is called when production orExpression is exited.
func (s *BaseFHIRPathListener) ExitOrExpression(ctx *OrExpressionContext) {}

// EnterAndExpression is called when production andExpression is entered.
func (s *BaseFHIRPathListener) EnterAndExpression(ctx *AndExpressionContext) {}

// ExitAndExpression is called when production andExpression is exited.
func (s *BaseFHIRPathListener) ExitAndExpression(ctx *AndExpressionContext) {}

// EnterMembershipExpression is called when production membershipExpression is entered.
func (s *BaseFHIRPathListener) EnterMembershipExpression(ctx *MembershipExpressionContext) {}

// ExitMembershipExpression is called when production membershipExpression is exited.
func (s *BaseFHIRPathListener) ExitMembershipExpression(ctx *MembershipExpressionContext) {}

// EnterInequalityExpression is called when production inequalityExpression is entered.
func (s *BaseFHIRPathListener) EnterInequalityExpression(ctx *InequalityExpressionContext) {}

// ExitInequalityExpression is called when production inequalityExpression is exited.
func (s *BaseFHIRPathListener) ExitInequalityExpression(ctx *InequalityExpressionContext) {}

// EnterInvocationExpression is called when production invocationExpression is entered.
func (s *BaseFHIRPathListener) EnterInvocationExpression(ctx *InvocationExpressionContext) {}

// ExitInvocationExpression is called when production invocationExpression is exited.
func (s *BaseFHIRPathListener) ExitInvocationExpression(ctx *InvocationExpressionContext) {}

// EnterEqualityExpression is called when production equalityExpression is entered.
func (s *BaseFHIRPathListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production equalityExpression is exited.
func (s *BaseFHIRPathListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterImpliesExpression is called when production impliesExpression is entered.
func (s *BaseFHIRPathListener) EnterImpliesExpression(ctx *ImpliesExpressionContext) {}

// ExitImpliesExpression is called when production impliesExpression is exited.
func (s *BaseFHIRPathListener) ExitImpliesExpression(ctx *ImpliesExpressionContext) {}

// EnterTermExpression is called when production termExpression is entered.
func (s *BaseFHIRPathListener) EnterTermExpression(ctx *TermExpressionContext) {}

// ExitTermExpression is called when production termExpression is exited.
func (s *BaseFHIRPathListener) ExitTermExpression(ctx *TermExpressionContext) {}

// EnterTypeExpression is called when production typeExpression is entered.
func (s *BaseFHIRPathListener) EnterTypeExpression(ctx *TypeExpressionContext) {}

// ExitTypeExpression is called when production typeExpression is exited.
func (s *BaseFHIRPathListener) ExitTypeExpression(ctx *TypeExpressionContext) {}

// EnterInvocationTerm is called when production invocationTerm is entered.
func (s *BaseFHIRPathListener) EnterInvocationTerm(ctx *InvocationTermContext) {}

// ExitInvocationTerm is called when production invocationTerm is exited.
func (s *BaseFHIRPathListener) ExitInvocationTerm(ctx *InvocationTermContext) {}

// EnterLiteralTerm is called when production literalTerm is entered.
func (s *BaseFHIRPathListener) EnterLiteralTerm(ctx *LiteralTermContext) {}

// ExitLiteralTerm is called when production literalTerm is exited.
func (s *BaseFHIRPathListener) ExitLiteralTerm(ctx *LiteralTermContext) {}

// EnterExternalConstantTerm is called when production externalConstantTerm is entered.
func (s *BaseFHIRPathListener) EnterExternalConstantTerm(ctx *ExternalConstantTermContext) {}

// ExitExternalConstantTerm is called when production externalConstantTerm is exited.
func (s *BaseFHIRPathListener) ExitExternalConstantTerm(ctx *ExternalConstantTermContext) {}

// EnterParenthesizedTerm is called when production parenthesizedTerm is entered.
func (s *BaseFHIRPathListener) EnterParenthesizedTerm(ctx *ParenthesizedTermContext) {}

// ExitParenthesizedTerm is called when production parenthesizedTerm is exited.
func (s *BaseFHIRPathListener) ExitParenthesizedTerm(ctx *ParenthesizedTermContext) {}

// EnterNullLiteral is called when production nullLiteral is entered.
func (s *BaseFHIRPathListener) EnterNullLiteral(ctx *NullLiteralContext) {}

// ExitNullLiteral is called when production nullLiteral is exited.
func (s *BaseFHIRPathListener) ExitNullLiteral(ctx *NullLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseFHIRPathListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseFHIRPathListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseFHIRPathListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseFHIRPathListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterNumberLiteral is called when production numberLiteral is entered.
func (s *BaseFHIRPathListener) EnterNumberLiteral(ctx *NumberLiteralContext) {}

// ExitNumberLiteral is called when production numberLiteral is exited.
func (s *BaseFHIRPathListener) ExitNumberLiteral(ctx *NumberLiteralContext) {}

// EnterDateLiteral is called when production dateLiteral is entered.
func (s *BaseFHIRPathListener) EnterDateLiteral(ctx *DateLiteralContext) {}

// ExitDateLiteral is called when production dateLiteral is exited.
func (s *BaseFHIRPathListener) ExitDateLiteral(ctx *DateLiteralContext) {}

// EnterDateTimeLiteral is called when production dateTimeLiteral is entered.
func (s *BaseFHIRPathListener) EnterDateTimeLiteral(ctx *DateTimeLiteralContext) {}

// ExitDateTimeLiteral is called when production dateTimeLiteral is exited.
func (s *BaseFHIRPathListener) ExitDateTimeLiteral(ctx *DateTimeLiteralContext) {}

// EnterTimeLiteral is called when production timeLiteral is entered.
func (s *BaseFHIRPathListener) EnterTimeLiteral(ctx *TimeLiteralContext) {}

// ExitTimeLiteral is called when production timeLiteral is exited.
func (s *BaseFHIRPathListener) ExitTimeLiteral(ctx *TimeLiteralContext) {}

// EnterQuantityLiteral is called when production quantityLiteral is entered.
func (s *BaseFHIRPathListener) EnterQuantityLiteral(ctx *QuantityLiteralContext) {}

// ExitQuantityLiteral is called when production quantityLiteral is exited.
func (s *BaseFHIRPathListener) ExitQuantityLiteral(ctx *QuantityLiteralContext) {}

// EnterExternalConstant is called when production externalConstant is entered.
func (s *BaseFHIRPathListener) EnterExternalConstant(ctx *ExternalConstantContext) {}

// ExitExternalConstant is called when production externalConstant is exited.
func (s *BaseFHIRPathListener) ExitExternalConstant(ctx *ExternalConstantContext) {}

// EnterMemberInvocation is called when production memberInvocation is entered.
func (s *BaseFHIRPathListener) EnterMemberInvocation(ctx *MemberInvocationContext) {}

// ExitMemberInvocation is called when production memberInvocation is exited.
func (s *BaseFHIRPathListener) ExitMemberInvocation(ctx *MemberInvocationContext) {}

// EnterFunctionInvocation is called when production functionInvocation is entered.
func (s *BaseFHIRPathListener) EnterFunctionInvocation(ctx *FunctionInvocationContext) {}

// ExitFunctionInvocation is called when production functionInvocation is exited.
func (s *BaseFHIRPathListener) ExitFunctionInvocation(ctx *FunctionInvocationContext) {}

// EnterThisInvocation is called when production thisInvocation is entered.
func (s *BaseFHIRPathListener) EnterThisInvocation(ctx *ThisInvocationContext) {}

// ExitThisInvocation is called when production thisInvocation is exited.
func (s *BaseFHIRPathListener) ExitThisInvocation(ctx *ThisInvocationContext) {}

// EnterIndexInvocation is called when production indexInvocation is entered.
func (s *BaseFHIRPathListener) EnterIndexInvocation(ctx *IndexInvocationContext) {}

// ExitIndexInvocation is called when production indexInvocation is exited.
func (s *BaseFHIRPathListener) ExitIndexInvocation(ctx *IndexInvocationContext) {}

// EnterTotalInvocation is called when production totalInvocation is entered.
func (s *BaseFHIRPathListener) EnterTotalInvocation(ctx *TotalInvocationContext) {}

// ExitTotalInvocation is called when production totalInvocation is exited.
func (s *BaseFHIRPathListener) ExitTotalInvocation(ctx *TotalInvocationContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseFHIRPathListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseFHIRPathListener) ExitFunction(ctx *FunctionContext) {}

// EnterParamList is called when production paramList is entered.
func (s *BaseFHIRPathListener) EnterParamList(ctx *ParamListContext) {}

// ExitParamList is called when production paramList is exited.
func (s *BaseFHIRPathListener) ExitParamList(ctx *ParamListContext) {}

// EnterQuantity is called when production quantity is entered.
func (s *BaseFHIRPathListener) EnterQuantity(ctx *QuantityContext) {}

// ExitQuantity is called when production quantity is exited.
func (s *BaseFHIRPathListener) ExitQuantity(ctx *QuantityContext) {}

// EnterUnit is called when production unit is entered.
func (s *BaseFHIRPathListener) EnterUnit(ctx *UnitContext) {}

// ExitUnit is called when production unit is exited.
func (s *BaseFHIRPathListener) ExitUnit(ctx *UnitContext) {}

// EnterDateTimePrecision is called when production dateTimePrecision is entered.
func (s *BaseFHIRPathListener) EnterDateTimePrecision(ctx *DateTimePrecisionContext) {}

// ExitDateTimePrecision is called when production dateTimePrecision is exited.
func (s *BaseFHIRPathListener) ExitDateTimePrecision(ctx *DateTimePrecisionContext) {}

// EnterPluralDateTimePrecision is called when production pluralDateTimePrecision is entered.
func (s *BaseFHIRPathListener) EnterPluralDateTimePrecision(ctx *PluralDateTimePrecisionContext) {}

// ExitPluralDateTimePrecision is called when production pluralDateTimePrecision is exited.
func (s *BaseFHIRPathListener) ExitPluralDateTimePrecision(ctx *PluralDateTimePrecisionContext) {}

// EnterTypeSpecifier is called when production typeSpecifier is entered.
func (s *BaseFHIRPathListener) EnterTypeSpecifier(ctx *TypeSpecifierContext) {}

// ExitTypeSpecifier is called when production typeSpecifier is exited.
func (s *BaseFHIRPathListener) ExitTypeSpecifier(ctx *TypeSpecifierContext) {}

// EnterQualifiedIdentifier is called when production qualifiedIdentifier is entered.
func (s *BaseFHIRPathListener) EnterQualifiedIdentifier(ctx *QualifiedIdentifierContext) {}

// ExitQualifiedIdentifier is called when production qualifiedIdentifier is exited.
func (s *BaseFHIRPathListener) ExitQualifiedIdentifier(ctx *QualifiedIdentifierContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseFHIRPathListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseFHIRPathListener) ExitIdentifier(ctx *IdentifierContext) {}
