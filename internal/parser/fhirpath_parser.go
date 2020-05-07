// Code generated from FHIRPath.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // FHIRPath

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 66, 152,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 35, 10,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 7, 2, 75, 10, 2, 12, 2, 14, 2, 78, 11, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 87, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 98, 10, 4, 3, 5, 3, 5, 3, 5, 5, 5, 103,
	10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 110, 10, 6, 3, 7, 3, 7, 3, 7,
	5, 7, 115, 10, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 7, 8, 122, 10, 8, 12, 8,
	14, 8, 125, 11, 8, 3, 9, 3, 9, 5, 9, 129, 10, 9, 3, 10, 3, 10, 3, 10, 5,
	10, 134, 10, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14,
	3, 14, 7, 14, 145, 10, 14, 12, 14, 14, 14, 148, 11, 14, 3, 15, 3, 15, 3,
	15, 2, 3, 2, 16, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 2,
	14, 3, 2, 6, 7, 3, 2, 8, 11, 4, 2, 6, 7, 12, 12, 3, 2, 16, 19, 3, 2, 20,
	23, 3, 2, 24, 25, 3, 2, 27, 28, 3, 2, 13, 14, 3, 2, 34, 35, 3, 2, 41, 48,
	3, 2, 49, 56, 5, 2, 13, 14, 24, 25, 60, 61, 2, 171, 2, 34, 3, 2, 2, 2,
	4, 86, 3, 2, 2, 2, 6, 97, 3, 2, 2, 2, 8, 99, 3, 2, 2, 2, 10, 109, 3, 2,
	2, 2, 12, 111, 3, 2, 2, 2, 14, 118, 3, 2, 2, 2, 16, 126, 3, 2, 2, 2, 18,
	133, 3, 2, 2, 2, 20, 135, 3, 2, 2, 2, 22, 137, 3, 2, 2, 2, 24, 139, 3,
	2, 2, 2, 26, 141, 3, 2, 2, 2, 28, 149, 3, 2, 2, 2, 30, 31, 8, 2, 1, 2,
	31, 35, 5, 4, 3, 2, 32, 33, 9, 2, 2, 2, 33, 35, 5, 2, 2, 13, 34, 30, 3,
	2, 2, 2, 34, 32, 3, 2, 2, 2, 35, 76, 3, 2, 2, 2, 36, 37, 12, 12, 2, 2,
	37, 38, 9, 3, 2, 2, 38, 75, 5, 2, 2, 13, 39, 40, 12, 11, 2, 2, 40, 41,
	9, 4, 2, 2, 41, 75, 5, 2, 2, 12, 42, 43, 12, 9, 2, 2, 43, 44, 7, 15, 2,
	2, 44, 75, 5, 2, 2, 10, 45, 46, 12, 8, 2, 2, 46, 47, 9, 5, 2, 2, 47, 75,
	5, 2, 2, 9, 48, 49, 12, 7, 2, 2, 49, 50, 9, 6, 2, 2, 50, 75, 5, 2, 2, 8,
	51, 52, 12, 6, 2, 2, 52, 53, 9, 7, 2, 2, 53, 75, 5, 2, 2, 7, 54, 55, 12,
	5, 2, 2, 55, 56, 7, 26, 2, 2, 56, 75, 5, 2, 2, 6, 57, 58, 12, 4, 2, 2,
	58, 59, 9, 8, 2, 2, 59, 75, 5, 2, 2, 5, 60, 61, 12, 3, 2, 2, 61, 62, 7,
	29, 2, 2, 62, 75, 5, 2, 2, 4, 63, 64, 12, 15, 2, 2, 64, 65, 7, 3, 2, 2,
	65, 75, 5, 10, 6, 2, 66, 67, 12, 14, 2, 2, 67, 68, 7, 4, 2, 2, 68, 69,
	5, 2, 2, 2, 69, 70, 7, 5, 2, 2, 70, 75, 3, 2, 2, 2, 71, 72, 12, 10, 2,
	2, 72, 73, 9, 9, 2, 2, 73, 75, 5, 24, 13, 2, 74, 36, 3, 2, 2, 2, 74, 39,
	3, 2, 2, 2, 74, 42, 3, 2, 2, 2, 74, 45, 3, 2, 2, 2, 74, 48, 3, 2, 2, 2,
	74, 51, 3, 2, 2, 2, 74, 54, 3, 2, 2, 2, 74, 57, 3, 2, 2, 2, 74, 60, 3,
	2, 2, 2, 74, 63, 3, 2, 2, 2, 74, 66, 3, 2, 2, 2, 74, 71, 3, 2, 2, 2, 75,
	78, 3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 3, 3, 2, 2,
	2, 78, 76, 3, 2, 2, 2, 79, 87, 5, 10, 6, 2, 80, 87, 5, 6, 4, 2, 81, 87,
	5, 8, 5, 2, 82, 83, 7, 30, 2, 2, 83, 84, 5, 2, 2, 2, 84, 85, 7, 31, 2,
	2, 85, 87, 3, 2, 2, 2, 86, 79, 3, 2, 2, 2, 86, 80, 3, 2, 2, 2, 86, 81,
	3, 2, 2, 2, 86, 82, 3, 2, 2, 2, 87, 5, 3, 2, 2, 2, 88, 89, 7, 32, 2, 2,
	89, 98, 7, 33, 2, 2, 90, 98, 9, 10, 2, 2, 91, 98, 7, 62, 2, 2, 92, 98,
	7, 63, 2, 2, 93, 98, 7, 57, 2, 2, 94, 98, 7, 58, 2, 2, 95, 98, 7, 59, 2,
	2, 96, 98, 5, 16, 9, 2, 97, 88, 3, 2, 2, 2, 97, 90, 3, 2, 2, 2, 97, 91,
	3, 2, 2, 2, 97, 92, 3, 2, 2, 2, 97, 93, 3, 2, 2, 2, 97, 94, 3, 2, 2, 2,
	97, 95, 3, 2, 2, 2, 97, 96, 3, 2, 2, 2, 98, 7, 3, 2, 2, 2, 99, 102, 7,
	36, 2, 2, 100, 103, 5, 28, 15, 2, 101, 103, 7, 62, 2, 2, 102, 100, 3, 2,
	2, 2, 102, 101, 3, 2, 2, 2, 103, 9, 3, 2, 2, 2, 104, 110, 5, 28, 15, 2,
	105, 110, 5, 12, 7, 2, 106, 110, 7, 37, 2, 2, 107, 110, 7, 38, 2, 2, 108,
	110, 7, 39, 2, 2, 109, 104, 3, 2, 2, 2, 109, 105, 3, 2, 2, 2, 109, 106,
	3, 2, 2, 2, 109, 107, 3, 2, 2, 2, 109, 108, 3, 2, 2, 2, 110, 11, 3, 2,
	2, 2, 111, 112, 5, 28, 15, 2, 112, 114, 7, 30, 2, 2, 113, 115, 5, 14, 8,
	2, 114, 113, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116,
	117, 7, 31, 2, 2, 117, 13, 3, 2, 2, 2, 118, 123, 5, 2, 2, 2, 119, 120,
	7, 40, 2, 2, 120, 122, 5, 2, 2, 2, 121, 119, 3, 2, 2, 2, 122, 125, 3, 2,
	2, 2, 123, 121, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124, 15, 3, 2, 2, 2,
	125, 123, 3, 2, 2, 2, 126, 128, 7, 63, 2, 2, 127, 129, 5, 18, 10, 2, 128,
	127, 3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129, 17, 3, 2, 2, 2, 130, 134, 5,
	20, 11, 2, 131, 134, 5, 22, 12, 2, 132, 134, 7, 62, 2, 2, 133, 130, 3,
	2, 2, 2, 133, 131, 3, 2, 2, 2, 133, 132, 3, 2, 2, 2, 134, 19, 3, 2, 2,
	2, 135, 136, 9, 11, 2, 2, 136, 21, 3, 2, 2, 2, 137, 138, 9, 12, 2, 2, 138,
	23, 3, 2, 2, 2, 139, 140, 5, 26, 14, 2, 140, 25, 3, 2, 2, 2, 141, 146,
	5, 28, 15, 2, 142, 143, 7, 3, 2, 2, 143, 145, 5, 28, 15, 2, 144, 142, 3,
	2, 2, 2, 145, 148, 3, 2, 2, 2, 146, 144, 3, 2, 2, 2, 146, 147, 3, 2, 2,
	2, 147, 27, 3, 2, 2, 2, 148, 146, 3, 2, 2, 2, 149, 150, 9, 13, 2, 2, 150,
	29, 3, 2, 2, 2, 14, 34, 74, 76, 86, 97, 102, 109, 114, 123, 128, 133, 146,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'.'", "'['", "']'", "'+'", "'-'", "'*'", "'/'", "'div'", "'mod'",
	"'&'", "'is'", "'as'", "'|'", "'<='", "'<'", "'>'", "'>='", "'='", "'~'",
	"'!='", "'!~'", "'in'", "'contains'", "'and'", "'or'", "'xor'", "'implies'",
	"'('", "')'", "'{'", "'}'", "'true'", "'false'", "'%'", "'$this'", "'$index'",
	"'$total'", "','", "'year'", "'month'", "'week'", "'day'", "'hour'", "'minute'",
	"'second'", "'millisecond'", "'years'", "'months'", "'weeks'", "'days'",
	"'hours'", "'minutes'", "'seconds'", "'milliseconds'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "DATE", "DATETIME", "TIME", "IDENTIFIER", "DELIMITEDIDENTIFIER", "STRING",
	"NUMBER", "WS", "COMMENT", "LINE_COMMENT",
}

var ruleNames = []string{
	"expression", "term", "literal", "externalConstant", "invocation", "function",
	"paramList", "quantity", "unit", "dateTimePrecision", "pluralDateTimePrecision",
	"typeSpecifier", "qualifiedIdentifier", "identifier",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type FHIRPathParser struct {
	*antlr.BaseParser
}

func NewFHIRPathParser(input antlr.TokenStream) *FHIRPathParser {
	this := new(FHIRPathParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "FHIRPath.g4"

	return this
}

// FHIRPathParser tokens.
const (
	FHIRPathParserEOF                 = antlr.TokenEOF
	FHIRPathParserT__0                = 1
	FHIRPathParserT__1                = 2
	FHIRPathParserT__2                = 3
	FHIRPathParserT__3                = 4
	FHIRPathParserT__4                = 5
	FHIRPathParserT__5                = 6
	FHIRPathParserT__6                = 7
	FHIRPathParserT__7                = 8
	FHIRPathParserT__8                = 9
	FHIRPathParserT__9                = 10
	FHIRPathParserT__10               = 11
	FHIRPathParserT__11               = 12
	FHIRPathParserT__12               = 13
	FHIRPathParserT__13               = 14
	FHIRPathParserT__14               = 15
	FHIRPathParserT__15               = 16
	FHIRPathParserT__16               = 17
	FHIRPathParserT__17               = 18
	FHIRPathParserT__18               = 19
	FHIRPathParserT__19               = 20
	FHIRPathParserT__20               = 21
	FHIRPathParserT__21               = 22
	FHIRPathParserT__22               = 23
	FHIRPathParserT__23               = 24
	FHIRPathParserT__24               = 25
	FHIRPathParserT__25               = 26
	FHIRPathParserT__26               = 27
	FHIRPathParserT__27               = 28
	FHIRPathParserT__28               = 29
	FHIRPathParserT__29               = 30
	FHIRPathParserT__30               = 31
	FHIRPathParserT__31               = 32
	FHIRPathParserT__32               = 33
	FHIRPathParserT__33               = 34
	FHIRPathParserT__34               = 35
	FHIRPathParserT__35               = 36
	FHIRPathParserT__36               = 37
	FHIRPathParserT__37               = 38
	FHIRPathParserT__38               = 39
	FHIRPathParserT__39               = 40
	FHIRPathParserT__40               = 41
	FHIRPathParserT__41               = 42
	FHIRPathParserT__42               = 43
	FHIRPathParserT__43               = 44
	FHIRPathParserT__44               = 45
	FHIRPathParserT__45               = 46
	FHIRPathParserT__46               = 47
	FHIRPathParserT__47               = 48
	FHIRPathParserT__48               = 49
	FHIRPathParserT__49               = 50
	FHIRPathParserT__50               = 51
	FHIRPathParserT__51               = 52
	FHIRPathParserT__52               = 53
	FHIRPathParserT__53               = 54
	FHIRPathParserDATE                = 55
	FHIRPathParserDATETIME            = 56
	FHIRPathParserTIME                = 57
	FHIRPathParserIDENTIFIER          = 58
	FHIRPathParserDELIMITEDIDENTIFIER = 59
	FHIRPathParserSTRING              = 60
	FHIRPathParserNUMBER              = 61
	FHIRPathParserWS                  = 62
	FHIRPathParserCOMMENT             = 63
	FHIRPathParserLINE_COMMENT        = 64
)

// FHIRPathParser rules.
const (
	FHIRPathParserRULE_expression              = 0
	FHIRPathParserRULE_term                    = 1
	FHIRPathParserRULE_literal                 = 2
	FHIRPathParserRULE_externalConstant        = 3
	FHIRPathParserRULE_invocation              = 4
	FHIRPathParserRULE_function                = 5
	FHIRPathParserRULE_paramList               = 6
	FHIRPathParserRULE_quantity                = 7
	FHIRPathParserRULE_unit                    = 8
	FHIRPathParserRULE_dateTimePrecision       = 9
	FHIRPathParserRULE_pluralDateTimePrecision = 10
	FHIRPathParserRULE_typeSpecifier           = 11
	FHIRPathParserRULE_qualifiedIdentifier     = 12
	FHIRPathParserRULE_identifier              = 13
)

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FHIRPathParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FHIRPathParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IndexerExpressionContext struct {
	*ExpressionContext
}

func NewIndexerExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexerExpressionContext {
	var p = new(IndexerExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *IndexerExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexerExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *IndexerExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IndexerExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterIndexerExpression(s)
	}
}

func (s *IndexerExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitIndexerExpression(s)
	}
}

type PolarityExpressionContext struct {
	*ExpressionContext
}

func NewPolarityExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PolarityExpressionContext {
	var p = new(PolarityExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *PolarityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PolarityExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PolarityExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterPolarityExpression(s)
	}
}

func (s *PolarityExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitPolarityExpression(s)
	}
}

type AdditiveExpressionContext struct {
	*ExpressionContext
}

func NewAdditiveExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AdditiveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AdditiveExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AdditiveExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitAdditiveExpression(s)
	}
}

type MultiplicativeExpressionContext struct {
	*ExpressionContext
}

func NewMultiplicativeExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MultiplicativeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MultiplicativeExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MultiplicativeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitMultiplicativeExpression(s)
	}
}

type UnionExpressionContext struct {
	*ExpressionContext
}

func NewUnionExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnionExpressionContext {
	var p = new(UnionExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *UnionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *UnionExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnionExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterUnionExpression(s)
	}
}

func (s *UnionExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitUnionExpression(s)
	}
}

type OrExpressionContext struct {
	*ExpressionContext
}

func NewOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExpressionContext {
	var p = new(OrExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *OrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *OrExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterOrExpression(s)
	}
}

func (s *OrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitOrExpression(s)
	}
}

type AndExpressionContext struct {
	*ExpressionContext
}

func NewAndExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExpressionContext {
	var p = new(AndExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AndExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterAndExpression(s)
	}
}

func (s *AndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitAndExpression(s)
	}
}

type MembershipExpressionContext struct {
	*ExpressionContext
}

func NewMembershipExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MembershipExpressionContext {
	var p = new(MembershipExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MembershipExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MembershipExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MembershipExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MembershipExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterMembershipExpression(s)
	}
}

func (s *MembershipExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitMembershipExpression(s)
	}
}

type InequalityExpressionContext struct {
	*ExpressionContext
}

func NewInequalityExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InequalityExpressionContext {
	var p = new(InequalityExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *InequalityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InequalityExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *InequalityExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *InequalityExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterInequalityExpression(s)
	}
}

func (s *InequalityExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitInequalityExpression(s)
	}
}

type InvocationExpressionContext struct {
	*ExpressionContext
}

func NewInvocationExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InvocationExpressionContext {
	var p = new(InvocationExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *InvocationExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvocationExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *InvocationExpressionContext) Invocation() IInvocationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInvocationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInvocationContext)
}

func (s *InvocationExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterInvocationExpression(s)
	}
}

func (s *InvocationExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitInvocationExpression(s)
	}
}

type EqualityExpressionContext struct {
	*ExpressionContext
}

func NewEqualityExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *EqualityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *EqualityExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EqualityExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitEqualityExpression(s)
	}
}

type ImpliesExpressionContext struct {
	*ExpressionContext
}

func NewImpliesExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ImpliesExpressionContext {
	var p = new(ImpliesExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ImpliesExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImpliesExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ImpliesExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ImpliesExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterImpliesExpression(s)
	}
}

func (s *ImpliesExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitImpliesExpression(s)
	}
}

type TermExpressionContext struct {
	*ExpressionContext
}

func NewTermExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TermExpressionContext {
	var p = new(TermExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *TermExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermExpressionContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *TermExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterTermExpression(s)
	}
}

func (s *TermExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitTermExpression(s)
	}
}

type TypeExpressionContext struct {
	*ExpressionContext
}

func NewTypeExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeExpressionContext {
	var p = new(TypeExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *TypeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TypeExpressionContext) TypeSpecifier() ITypeSpecifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecifierContext)
}

func (s *TypeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterTypeExpression(s)
	}
}

func (s *TypeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitTypeExpression(s)
	}
}

func (p *FHIRPathParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *FHIRPathParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, FHIRPathParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(32)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FHIRPathParserT__10, FHIRPathParserT__11, FHIRPathParserT__21, FHIRPathParserT__22, FHIRPathParserT__27, FHIRPathParserT__29, FHIRPathParserT__31, FHIRPathParserT__32, FHIRPathParserT__33, FHIRPathParserT__34, FHIRPathParserT__35, FHIRPathParserT__36, FHIRPathParserDATE, FHIRPathParserDATETIME, FHIRPathParserTIME, FHIRPathParserIDENTIFIER, FHIRPathParserDELIMITEDIDENTIFIER, FHIRPathParserSTRING, FHIRPathParserNUMBER:
		localctx = NewTermExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(29)
			p.Term()
		}

	case FHIRPathParserT__3, FHIRPathParserT__4:
		localctx = NewPolarityExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(30)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FHIRPathParserT__3 || _la == FHIRPathParserT__4) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(31)
			p.expression(11)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(72)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultiplicativeExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FHIRPathParserRULE_expression)
				p.SetState(34)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(35)
					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FHIRPathParserT__5)|(1<<FHIRPathParserT__6)|(1<<FHIRPathParserT__7)|(1<<FHIRPathParserT__8))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(36)
					p.expression(11)
				}

			case 2:
				localctx = NewAdditiveExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FHIRPathParserRULE_expression)
				p.SetState(37)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(38)
					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FHIRPathParserT__3)|(1<<FHIRPathParserT__4)|(1<<FHIRPathParserT__9))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(39)
					p.expression(10)
				}

			case 3:
				localctx = NewUnionExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FHIRPathParserRULE_expression)
				p.SetState(40)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(41)
					p.Match(FHIRPathParserT__12)
				}
				{
					p.SetState(42)
					p.expression(8)
				}

			case 4:
				localctx = NewInequalityExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FHIRPathParserRULE_expression)
				p.SetState(43)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(44)
					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FHIRPathParserT__13)|(1<<FHIRPathParserT__14)|(1<<FHIRPathParserT__15)|(1<<FHIRPathParserT__16))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(45)
					p.expression(7)
				}

			case 5:
				localctx = NewEqualityExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FHIRPathParserRULE_expression)
				p.SetState(46)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(47)
					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FHIRPathParserT__17)|(1<<FHIRPathParserT__18)|(1<<FHIRPathParserT__19)|(1<<FHIRPathParserT__20))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(48)
					p.expression(6)
				}

			case 6:
				localctx = NewMembershipExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FHIRPathParserRULE_expression)
				p.SetState(49)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(50)
					_la = p.GetTokenStream().LA(1)

					if !(_la == FHIRPathParserT__21 || _la == FHIRPathParserT__22) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(51)
					p.expression(5)
				}

			case 7:
				localctx = NewAndExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FHIRPathParserRULE_expression)
				p.SetState(52)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(53)
					p.Match(FHIRPathParserT__23)
				}
				{
					p.SetState(54)
					p.expression(4)
				}

			case 8:
				localctx = NewOrExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FHIRPathParserRULE_expression)
				p.SetState(55)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(56)
					_la = p.GetTokenStream().LA(1)

					if !(_la == FHIRPathParserT__24 || _la == FHIRPathParserT__25) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(57)
					p.expression(3)
				}

			case 9:
				localctx = NewImpliesExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FHIRPathParserRULE_expression)
				p.SetState(58)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(59)
					p.Match(FHIRPathParserT__26)
				}
				{
					p.SetState(60)
					p.expression(2)
				}

			case 10:
				localctx = NewInvocationExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FHIRPathParserRULE_expression)
				p.SetState(61)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(62)
					p.Match(FHIRPathParserT__0)
				}
				{
					p.SetState(63)
					p.Invocation()
				}

			case 11:
				localctx = NewIndexerExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FHIRPathParserRULE_expression)
				p.SetState(64)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(65)
					p.Match(FHIRPathParserT__1)
				}
				{
					p.SetState(66)
					p.expression(0)
				}
				{
					p.SetState(67)
					p.Match(FHIRPathParserT__2)
				}

			case 12:
				localctx = NewTypeExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FHIRPathParserRULE_expression)
				p.SetState(69)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(70)
					_la = p.GetTokenStream().LA(1)

					if !(_la == FHIRPathParserT__10 || _la == FHIRPathParserT__11) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(71)
					p.TypeSpecifier()
				}

			}

		}
		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FHIRPathParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FHIRPathParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) CopyFrom(ctx *TermContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExternalConstantTermContext struct {
	*TermContext
}

func NewExternalConstantTermContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExternalConstantTermContext {
	var p = new(ExternalConstantTermContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *ExternalConstantTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExternalConstantTermContext) ExternalConstant() IExternalConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExternalConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExternalConstantContext)
}

func (s *ExternalConstantTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterExternalConstantTerm(s)
	}
}

func (s *ExternalConstantTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitExternalConstantTerm(s)
	}
}

type LiteralTermContext struct {
	*TermContext
}

func NewLiteralTermContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralTermContext {
	var p = new(LiteralTermContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *LiteralTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralTermContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterLiteralTerm(s)
	}
}

func (s *LiteralTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitLiteralTerm(s)
	}
}

type ParenthesizedTermContext struct {
	*TermContext
}

func NewParenthesizedTermContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesizedTermContext {
	var p = new(ParenthesizedTermContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *ParenthesizedTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesizedTermContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenthesizedTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterParenthesizedTerm(s)
	}
}

func (s *ParenthesizedTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitParenthesizedTerm(s)
	}
}

type InvocationTermContext struct {
	*TermContext
}

func NewInvocationTermContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InvocationTermContext {
	var p = new(InvocationTermContext)

	p.TermContext = NewEmptyTermContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TermContext))

	return p
}

func (s *InvocationTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvocationTermContext) Invocation() IInvocationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInvocationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInvocationContext)
}

func (s *InvocationTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterInvocationTerm(s)
	}
}

func (s *InvocationTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitInvocationTerm(s)
	}
}

func (p *FHIRPathParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FHIRPathParserRULE_term)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(84)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FHIRPathParserT__10, FHIRPathParserT__11, FHIRPathParserT__21, FHIRPathParserT__22, FHIRPathParserT__34, FHIRPathParserT__35, FHIRPathParserT__36, FHIRPathParserIDENTIFIER, FHIRPathParserDELIMITEDIDENTIFIER:
		localctx = NewInvocationTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(77)
			p.Invocation()
		}

	case FHIRPathParserT__29, FHIRPathParserT__31, FHIRPathParserT__32, FHIRPathParserDATE, FHIRPathParserDATETIME, FHIRPathParserTIME, FHIRPathParserSTRING, FHIRPathParserNUMBER:
		localctx = NewLiteralTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(78)
			p.Literal()
		}

	case FHIRPathParserT__33:
		localctx = NewExternalConstantTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(79)
			p.ExternalConstant()
		}

	case FHIRPathParserT__27:
		localctx = NewParenthesizedTermContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(80)
			p.Match(FHIRPathParserT__27)
		}
		{
			p.SetState(81)
			p.expression(0)
		}
		{
			p.SetState(82)
			p.Match(FHIRPathParserT__28)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FHIRPathParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FHIRPathParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) CopyFrom(ctx *LiteralContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TimeLiteralContext struct {
	*LiteralContext
}

func NewTimeLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TimeLiteralContext {
	var p = new(TimeLiteralContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *TimeLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeLiteralContext) TIME() antlr.TerminalNode {
	return s.GetToken(FHIRPathParserTIME, 0)
}

func (s *TimeLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterTimeLiteral(s)
	}
}

func (s *TimeLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitTimeLiteral(s)
	}
}

type NullLiteralContext struct {
	*LiteralContext
}

func NewNullLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullLiteralContext {
	var p = new(NullLiteralContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *NullLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterNullLiteral(s)
	}
}

func (s *NullLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitNullLiteral(s)
	}
}

type DateTimeLiteralContext struct {
	*LiteralContext
}

func NewDateTimeLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DateTimeLiteralContext {
	var p = new(DateTimeLiteralContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *DateTimeLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateTimeLiteralContext) DATETIME() antlr.TerminalNode {
	return s.GetToken(FHIRPathParserDATETIME, 0)
}

func (s *DateTimeLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterDateTimeLiteral(s)
	}
}

func (s *DateTimeLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitDateTimeLiteral(s)
	}
}

type StringLiteralContext struct {
	*LiteralContext
}

func NewStringLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLiteralContext {
	var p = new(StringLiteralContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) STRING() antlr.TerminalNode {
	return s.GetToken(FHIRPathParserSTRING, 0)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

type DateLiteralContext struct {
	*LiteralContext
}

func NewDateLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DateLiteralContext {
	var p = new(DateLiteralContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *DateLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateLiteralContext) DATE() antlr.TerminalNode {
	return s.GetToken(FHIRPathParserDATE, 0)
}

func (s *DateLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterDateLiteral(s)
	}
}

func (s *DateLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitDateLiteral(s)
	}
}

type BooleanLiteralContext struct {
	*LiteralContext
}

func NewBooleanLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitBooleanLiteral(s)
	}
}

type NumberLiteralContext struct {
	*LiteralContext
}

func NewNumberLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberLiteralContext {
	var p = new(NumberLiteralContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *NumberLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberLiteralContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(FHIRPathParserNUMBER, 0)
}

func (s *NumberLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterNumberLiteral(s)
	}
}

func (s *NumberLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitNumberLiteral(s)
	}
}

type QuantityLiteralContext struct {
	*LiteralContext
}

func NewQuantityLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *QuantityLiteralContext {
	var p = new(QuantityLiteralContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *QuantityLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuantityLiteralContext) Quantity() IQuantityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuantityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuantityContext)
}

func (s *QuantityLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterQuantityLiteral(s)
	}
}

func (s *QuantityLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitQuantityLiteral(s)
	}
}

func (p *FHIRPathParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FHIRPathParserRULE_literal)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNullLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(86)
			p.Match(FHIRPathParserT__29)
		}
		{
			p.SetState(87)
			p.Match(FHIRPathParserT__30)
		}

	case 2:
		localctx = NewBooleanLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(88)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FHIRPathParserT__31 || _la == FHIRPathParserT__32) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 3:
		localctx = NewStringLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(89)
			p.Match(FHIRPathParserSTRING)
		}

	case 4:
		localctx = NewNumberLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(90)
			p.Match(FHIRPathParserNUMBER)
		}

	case 5:
		localctx = NewDateLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(91)
			p.Match(FHIRPathParserDATE)
		}

	case 6:
		localctx = NewDateTimeLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(92)
			p.Match(FHIRPathParserDATETIME)
		}

	case 7:
		localctx = NewTimeLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(93)
			p.Match(FHIRPathParserTIME)
		}

	case 8:
		localctx = NewQuantityLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(94)
			p.Quantity()
		}

	}

	return localctx
}

// IExternalConstantContext is an interface to support dynamic dispatch.
type IExternalConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExternalConstantContext differentiates from other interfaces.
	IsExternalConstantContext()
}

type ExternalConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExternalConstantContext() *ExternalConstantContext {
	var p = new(ExternalConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FHIRPathParserRULE_externalConstant
	return p
}

func (*ExternalConstantContext) IsExternalConstantContext() {}

func NewExternalConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExternalConstantContext {
	var p = new(ExternalConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FHIRPathParserRULE_externalConstant

	return p
}

func (s *ExternalConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ExternalConstantContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ExternalConstantContext) STRING() antlr.TerminalNode {
	return s.GetToken(FHIRPathParserSTRING, 0)
}

func (s *ExternalConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExternalConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExternalConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterExternalConstant(s)
	}
}

func (s *ExternalConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitExternalConstant(s)
	}
}

func (p *FHIRPathParser) ExternalConstant() (localctx IExternalConstantContext) {
	localctx = NewExternalConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FHIRPathParserRULE_externalConstant)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.Match(FHIRPathParserT__33)
	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FHIRPathParserT__10, FHIRPathParserT__11, FHIRPathParserT__21, FHIRPathParserT__22, FHIRPathParserIDENTIFIER, FHIRPathParserDELIMITEDIDENTIFIER:
		{
			p.SetState(98)
			p.Identifier()
		}

	case FHIRPathParserSTRING:
		{
			p.SetState(99)
			p.Match(FHIRPathParserSTRING)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInvocationContext is an interface to support dynamic dispatch.
type IInvocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInvocationContext differentiates from other interfaces.
	IsInvocationContext()
}

type InvocationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInvocationContext() *InvocationContext {
	var p = new(InvocationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FHIRPathParserRULE_invocation
	return p
}

func (*InvocationContext) IsInvocationContext() {}

func NewInvocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InvocationContext {
	var p = new(InvocationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FHIRPathParserRULE_invocation

	return p
}

func (s *InvocationContext) GetParser() antlr.Parser { return s.parser }

func (s *InvocationContext) CopyFrom(ctx *InvocationContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *InvocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TotalInvocationContext struct {
	*InvocationContext
}

func NewTotalInvocationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TotalInvocationContext {
	var p = new(TotalInvocationContext)

	p.InvocationContext = NewEmptyInvocationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InvocationContext))

	return p
}

func (s *TotalInvocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TotalInvocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterTotalInvocation(s)
	}
}

func (s *TotalInvocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitTotalInvocation(s)
	}
}

type ThisInvocationContext struct {
	*InvocationContext
}

func NewThisInvocationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ThisInvocationContext {
	var p = new(ThisInvocationContext)

	p.InvocationContext = NewEmptyInvocationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InvocationContext))

	return p
}

func (s *ThisInvocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThisInvocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterThisInvocation(s)
	}
}

func (s *ThisInvocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitThisInvocation(s)
	}
}

type IndexInvocationContext struct {
	*InvocationContext
}

func NewIndexInvocationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexInvocationContext {
	var p = new(IndexInvocationContext)

	p.InvocationContext = NewEmptyInvocationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InvocationContext))

	return p
}

func (s *IndexInvocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexInvocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterIndexInvocation(s)
	}
}

func (s *IndexInvocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitIndexInvocation(s)
	}
}

type FunctionInvocationContext struct {
	*InvocationContext
}

func NewFunctionInvocationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionInvocationContext {
	var p = new(FunctionInvocationContext)

	p.InvocationContext = NewEmptyInvocationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InvocationContext))

	return p
}

func (s *FunctionInvocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionInvocationContext) Function() IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *FunctionInvocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterFunctionInvocation(s)
	}
}

func (s *FunctionInvocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitFunctionInvocation(s)
	}
}

type MemberInvocationContext struct {
	*InvocationContext
}

func NewMemberInvocationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MemberInvocationContext {
	var p = new(MemberInvocationContext)

	p.InvocationContext = NewEmptyInvocationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InvocationContext))

	return p
}

func (s *MemberInvocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberInvocationContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *MemberInvocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterMemberInvocation(s)
	}
}

func (s *MemberInvocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitMemberInvocation(s)
	}
}

func (p *FHIRPathParser) Invocation() (localctx IInvocationContext) {
	localctx = NewInvocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FHIRPathParserRULE_invocation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewMemberInvocationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(102)
			p.Identifier()
		}

	case 2:
		localctx = NewFunctionInvocationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(103)
			p.Function()
		}

	case 3:
		localctx = NewThisInvocationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(104)
			p.Match(FHIRPathParserT__34)
		}

	case 4:
		localctx = NewIndexInvocationContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(105)
			p.Match(FHIRPathParserT__35)
		}

	case 5:
		localctx = NewTotalInvocationContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(106)
			p.Match(FHIRPathParserT__36)
		}

	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FHIRPathParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FHIRPathParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *FunctionContext) ParamList() IParamListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamListContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (p *FHIRPathParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FHIRPathParserRULE_function)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Identifier()
	}
	{
		p.SetState(110)
		p.Match(FHIRPathParserT__27)
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FHIRPathParserT__3)|(1<<FHIRPathParserT__4)|(1<<FHIRPathParserT__10)|(1<<FHIRPathParserT__11)|(1<<FHIRPathParserT__21)|(1<<FHIRPathParserT__22)|(1<<FHIRPathParserT__27)|(1<<FHIRPathParserT__29))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(FHIRPathParserT__31-32))|(1<<(FHIRPathParserT__32-32))|(1<<(FHIRPathParserT__33-32))|(1<<(FHIRPathParserT__34-32))|(1<<(FHIRPathParserT__35-32))|(1<<(FHIRPathParserT__36-32))|(1<<(FHIRPathParserDATE-32))|(1<<(FHIRPathParserDATETIME-32))|(1<<(FHIRPathParserTIME-32))|(1<<(FHIRPathParserIDENTIFIER-32))|(1<<(FHIRPathParserDELIMITEDIDENTIFIER-32))|(1<<(FHIRPathParserSTRING-32))|(1<<(FHIRPathParserNUMBER-32)))) != 0) {
		{
			p.SetState(111)
			p.ParamList()
		}

	}
	{
		p.SetState(114)
		p.Match(FHIRPathParserT__28)
	}

	return localctx
}

// IParamListContext is an interface to support dynamic dispatch.
type IParamListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamListContext differentiates from other interfaces.
	IsParamListContext()
}

type ParamListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamListContext() *ParamListContext {
	var p = new(ParamListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FHIRPathParserRULE_paramList
	return p
}

func (*ParamListContext) IsParamListContext() {}

func NewParamListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamListContext {
	var p = new(ParamListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FHIRPathParserRULE_paramList

	return p
}

func (s *ParamListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamListContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ParamListContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParamListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterParamList(s)
	}
}

func (s *ParamListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitParamList(s)
	}
}

func (p *FHIRPathParser) ParamList() (localctx IParamListContext) {
	localctx = NewParamListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FHIRPathParserRULE_paramList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.expression(0)
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FHIRPathParserT__37 {
		{
			p.SetState(117)
			p.Match(FHIRPathParserT__37)
		}
		{
			p.SetState(118)
			p.expression(0)
		}

		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IQuantityContext is an interface to support dynamic dispatch.
type IQuantityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuantityContext differentiates from other interfaces.
	IsQuantityContext()
}

type QuantityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuantityContext() *QuantityContext {
	var p = new(QuantityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FHIRPathParserRULE_quantity
	return p
}

func (*QuantityContext) IsQuantityContext() {}

func NewQuantityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuantityContext {
	var p = new(QuantityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FHIRPathParserRULE_quantity

	return p
}

func (s *QuantityContext) GetParser() antlr.Parser { return s.parser }

func (s *QuantityContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(FHIRPathParserNUMBER, 0)
}

func (s *QuantityContext) Unit() IUnitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnitContext)
}

func (s *QuantityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuantityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuantityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterQuantity(s)
	}
}

func (s *QuantityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitQuantity(s)
	}
}

func (p *FHIRPathParser) Quantity() (localctx IQuantityContext) {
	localctx = NewQuantityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FHIRPathParserRULE_quantity)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)
		p.Match(FHIRPathParserNUMBER)
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(125)
			p.Unit()
		}

	}

	return localctx
}

// IUnitContext is an interface to support dynamic dispatch.
type IUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnitContext differentiates from other interfaces.
	IsUnitContext()
}

type UnitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnitContext() *UnitContext {
	var p = new(UnitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FHIRPathParserRULE_unit
	return p
}

func (*UnitContext) IsUnitContext() {}

func NewUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnitContext {
	var p = new(UnitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FHIRPathParserRULE_unit

	return p
}

func (s *UnitContext) GetParser() antlr.Parser { return s.parser }

func (s *UnitContext) DateTimePrecision() IDateTimePrecisionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateTimePrecisionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDateTimePrecisionContext)
}

func (s *UnitContext) PluralDateTimePrecision() IPluralDateTimePrecisionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPluralDateTimePrecisionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPluralDateTimePrecisionContext)
}

func (s *UnitContext) STRING() antlr.TerminalNode {
	return s.GetToken(FHIRPathParserSTRING, 0)
}

func (s *UnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterUnit(s)
	}
}

func (s *UnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitUnit(s)
	}
}

func (p *FHIRPathParser) Unit() (localctx IUnitContext) {
	localctx = NewUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FHIRPathParserRULE_unit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(131)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FHIRPathParserT__38, FHIRPathParserT__39, FHIRPathParserT__40, FHIRPathParserT__41, FHIRPathParserT__42, FHIRPathParserT__43, FHIRPathParserT__44, FHIRPathParserT__45:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(128)
			p.DateTimePrecision()
		}

	case FHIRPathParserT__46, FHIRPathParserT__47, FHIRPathParserT__48, FHIRPathParserT__49, FHIRPathParserT__50, FHIRPathParserT__51, FHIRPathParserT__52, FHIRPathParserT__53:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)
			p.PluralDateTimePrecision()
		}

	case FHIRPathParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(130)
			p.Match(FHIRPathParserSTRING)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDateTimePrecisionContext is an interface to support dynamic dispatch.
type IDateTimePrecisionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDateTimePrecisionContext differentiates from other interfaces.
	IsDateTimePrecisionContext()
}

type DateTimePrecisionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateTimePrecisionContext() *DateTimePrecisionContext {
	var p = new(DateTimePrecisionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FHIRPathParserRULE_dateTimePrecision
	return p
}

func (*DateTimePrecisionContext) IsDateTimePrecisionContext() {}

func NewDateTimePrecisionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateTimePrecisionContext {
	var p = new(DateTimePrecisionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FHIRPathParserRULE_dateTimePrecision

	return p
}

func (s *DateTimePrecisionContext) GetParser() antlr.Parser { return s.parser }
func (s *DateTimePrecisionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateTimePrecisionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateTimePrecisionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterDateTimePrecision(s)
	}
}

func (s *DateTimePrecisionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitDateTimePrecision(s)
	}
}

func (p *FHIRPathParser) DateTimePrecision() (localctx IDateTimePrecisionContext) {
	localctx = NewDateTimePrecisionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FHIRPathParserRULE_dateTimePrecision)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-39)&-(0x1f+1)) == 0 && ((1<<uint((_la-39)))&((1<<(FHIRPathParserT__38-39))|(1<<(FHIRPathParserT__39-39))|(1<<(FHIRPathParserT__40-39))|(1<<(FHIRPathParserT__41-39))|(1<<(FHIRPathParserT__42-39))|(1<<(FHIRPathParserT__43-39))|(1<<(FHIRPathParserT__44-39))|(1<<(FHIRPathParserT__45-39)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IPluralDateTimePrecisionContext is an interface to support dynamic dispatch.
type IPluralDateTimePrecisionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPluralDateTimePrecisionContext differentiates from other interfaces.
	IsPluralDateTimePrecisionContext()
}

type PluralDateTimePrecisionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPluralDateTimePrecisionContext() *PluralDateTimePrecisionContext {
	var p = new(PluralDateTimePrecisionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FHIRPathParserRULE_pluralDateTimePrecision
	return p
}

func (*PluralDateTimePrecisionContext) IsPluralDateTimePrecisionContext() {}

func NewPluralDateTimePrecisionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PluralDateTimePrecisionContext {
	var p = new(PluralDateTimePrecisionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FHIRPathParserRULE_pluralDateTimePrecision

	return p
}

func (s *PluralDateTimePrecisionContext) GetParser() antlr.Parser { return s.parser }
func (s *PluralDateTimePrecisionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PluralDateTimePrecisionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PluralDateTimePrecisionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterPluralDateTimePrecision(s)
	}
}

func (s *PluralDateTimePrecisionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitPluralDateTimePrecision(s)
	}
}

func (p *FHIRPathParser) PluralDateTimePrecision() (localctx IPluralDateTimePrecisionContext) {
	localctx = NewPluralDateTimePrecisionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, FHIRPathParserRULE_pluralDateTimePrecision)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-47)&-(0x1f+1)) == 0 && ((1<<uint((_la-47)))&((1<<(FHIRPathParserT__46-47))|(1<<(FHIRPathParserT__47-47))|(1<<(FHIRPathParserT__48-47))|(1<<(FHIRPathParserT__49-47))|(1<<(FHIRPathParserT__50-47))|(1<<(FHIRPathParserT__51-47))|(1<<(FHIRPathParserT__52-47))|(1<<(FHIRPathParserT__53-47)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITypeSpecifierContext is an interface to support dynamic dispatch.
type ITypeSpecifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSpecifierContext differentiates from other interfaces.
	IsTypeSpecifierContext()
}

type TypeSpecifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSpecifierContext() *TypeSpecifierContext {
	var p = new(TypeSpecifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FHIRPathParserRULE_typeSpecifier
	return p
}

func (*TypeSpecifierContext) IsTypeSpecifierContext() {}

func NewTypeSpecifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSpecifierContext {
	var p = new(TypeSpecifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FHIRPathParserRULE_typeSpecifier

	return p
}

func (s *TypeSpecifierContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSpecifierContext) QualifiedIdentifier() IQualifiedIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifiedIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualifiedIdentifierContext)
}

func (s *TypeSpecifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSpecifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterTypeSpecifier(s)
	}
}

func (s *TypeSpecifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitTypeSpecifier(s)
	}
}

func (p *FHIRPathParser) TypeSpecifier() (localctx ITypeSpecifierContext) {
	localctx = NewTypeSpecifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, FHIRPathParserRULE_typeSpecifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(137)
		p.QualifiedIdentifier()
	}

	return localctx
}

// IQualifiedIdentifierContext is an interface to support dynamic dispatch.
type IQualifiedIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQualifiedIdentifierContext differentiates from other interfaces.
	IsQualifiedIdentifierContext()
}

type QualifiedIdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualifiedIdentifierContext() *QualifiedIdentifierContext {
	var p = new(QualifiedIdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FHIRPathParserRULE_qualifiedIdentifier
	return p
}

func (*QualifiedIdentifierContext) IsQualifiedIdentifierContext() {}

func NewQualifiedIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifiedIdentifierContext {
	var p = new(QualifiedIdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FHIRPathParserRULE_qualifiedIdentifier

	return p
}

func (s *QualifiedIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *QualifiedIdentifierContext) AllIdentifier() []IIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifierContext)(nil)).Elem())
	var tst = make([]IIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifierContext)
		}
	}

	return tst
}

func (s *QualifiedIdentifierContext) Identifier(i int) IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *QualifiedIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifiedIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualifiedIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterQualifiedIdentifier(s)
	}
}

func (s *QualifiedIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitQualifiedIdentifier(s)
	}
}

func (p *FHIRPathParser) QualifiedIdentifier() (localctx IQualifiedIdentifierContext) {
	localctx = NewQualifiedIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, FHIRPathParserRULE_qualifiedIdentifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		p.Identifier()
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(140)
				p.Match(FHIRPathParserT__0)
			}
			{
				p.SetState(141)
				p.Identifier()
			}

		}
		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	return localctx
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FHIRPathParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FHIRPathParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FHIRPathParserIDENTIFIER, 0)
}

func (s *IdentifierContext) DELIMITEDIDENTIFIER() antlr.TerminalNode {
	return s.GetToken(FHIRPathParserDELIMITEDIDENTIFIER, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FHIRPathListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *FHIRPathParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, FHIRPathParserRULE_identifier)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FHIRPathParserT__10)|(1<<FHIRPathParserT__11)|(1<<FHIRPathParserT__21)|(1<<FHIRPathParserT__22))) != 0) || _la == FHIRPathParserIDENTIFIER || _la == FHIRPathParserDELIMITEDIDENTIFIER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *FHIRPathParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *FHIRPathParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
