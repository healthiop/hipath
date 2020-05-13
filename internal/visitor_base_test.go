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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVisitorAddError(t *testing.T) {
	c := NewErrorItemCollection()
	v := NewVisitor(c)

	v.AddError(newRuleContext(10, 5), "test error")
	v.AddError(newRuleContext(12, 7), "test error2")

	if assert.Len(t, c.Items(), 2) {
		err := c.Items()[0]
		assert.Equal(t, 10, err.Line())
		assert.Equal(t, 5, err.Column())
		assert.Equal(t, "test error", err.Msg())

		err = c.Items()[1]
		assert.Equal(t, 12, err.Line())
		assert.Equal(t, 7, err.Column())
		assert.Equal(t, "test error2", err.Msg())
	}
}

func TestVisitorVisitChildrenNone(t *testing.T) {
	v := NewVisitor(NewErrorItemCollection())
	assert.Nil(t, v.VisitChildren(newRuleNodeMock(make([]antlr.Tree, 0))),
		"without children no data is expected")
}

func TestVisitorVisitChildren(t *testing.T) {
	innerChildren := make([]antlr.Tree, 1)
	innerChildren[0] = newTerminalNodeMock("inner")

	children := make([]antlr.Tree, 3)
	children[0] = newRuleNodeMock(innerChildren)
	children[1] = newTerminalNodeMock("outer")
	children[2] = newOtherNodeMock()

	v := NewVisitor(NewErrorItemCollection())
	r := v.VisitChildren(newRuleNodeMock(children))

	if assert.IsType(t, ([]interface{})(nil), r) {
		a := r.([]interface{})
		if assert.Len(t, a, 3) {
			if assert.IsType(t, ([]interface{})(nil), a[0]) {
				a0 := a[0].([]interface{})
				if assert.Len(t, a0, 1) {
					assert.Equal(t, "inner", a0[0])
				}
			}
			assert.Equal(t, "outer", a[1])
			assert.Nil(t, a[2])
		}
	}
}

func TestVisitorVisitFirstChild(t *testing.T) {
	children := make([]antlr.Tree, 2)
	children[0] = newTerminalNodeMock("first")
	children[1] = newTerminalNodeMock("second")

	v := NewVisitor(NewErrorItemCollection())
	r := v.VisitFirstChild(newRuleNodeMock(children))

	assert.Equal(t, "first", r)
}

func TestVisitorVisitChild(t *testing.T) {
	children := make([]antlr.Tree, 2)
	children[0] = newTerminalNodeMock("first")
	children[1] = newTerminalNodeMock("second")

	v := NewVisitor(NewErrorItemCollection())
	r := v.VisitChild(newRuleNodeMock(children), 1)

	assert.Equal(t, "second", r)
}

func TestVisitorVisitChildNotExist(t *testing.T) {
	children := make([]antlr.Tree, 2)
	children[0] = newTerminalNodeMock("first")
	children[1] = newTerminalNodeMock("second")

	v := NewVisitor(NewErrorItemCollection())
	r := v.VisitChild(newRuleNodeMock(children), 2)

	assert.Nil(t, r, "child does not exist")
}

func TestVisitorVisit(t *testing.T) {
	ctx := newRuleContext(81, 32)
	c := NewErrorItemCollection()
	v := NewVisitor(c)
	r := v.visit(ctx, func(ctx antlr.ParserRuleContext) (interface{}, error) {
		return "test result", nil
	})

	assert.False(t, c.HasErrors(), "no errors expected")
	assert.Equal(t, "test result", r)
}

func TestVisitorVisitError(t *testing.T) {
	ctx := newRuleContext(81, 32)
	c := NewErrorItemCollection()
	v := NewVisitor(c)
	r := v.visit(ctx, func(ctx antlr.ParserRuleContext) (interface{}, error) {
		return nil, fmt.Errorf("test error")
	})

	if assert.Len(t, c.Items(), 1, "errors expected") {
		item := c.Items()[0]
		assert.Equal(t, 81, item.Line())
		assert.Equal(t, 32, item.Column())
		assert.Equal(t, "test error", item.Msg())
	}
	assert.Nil(t, r, "no returned value expected")
}

func TestVisitorTree(t *testing.T) {
	children := make([]antlr.Token, 2)
	children[0] = newTokenMock(32, 2, "first")
	children[1] = newTokenMock(32, 2, "second")

	ctx := newRuleContextWithChildren(81, 32, children)
	c := NewErrorItemCollection()
	v := NewVisitor(c)
	r := v.visitTree(ctx, 2, func(ctx antlr.ParserRuleContext, args []interface{}) (interface{}, error) {
		if assert.Len(t, args, 2) {
			assert.Equal(t, "first", args[0])
			assert.Equal(t, "second", args[1])
		}
		return "test result", nil
	})

	assert.False(t, c.HasErrors(), "no errors expected")
	assert.Equal(t, "test result", r)
}

func TestVisitorTreeMissingArgs(t *testing.T) {
	children := make([]antlr.Token, 1)
	children[0] = newTokenMock(32, 2, "first")

	ctx := newRuleContextWithChildren(81, 32, children)
	c := NewErrorItemCollection()
	v := NewVisitor(c)
	r := v.visitTree(ctx, 2, func(ctx antlr.ParserRuleContext, args []interface{}) (interface{}, error) {
		assert.Fail(t, "function must not be invoked")
		return "test result", nil
	})

	assert.False(t, c.HasErrors(), "no errors expected")
	assert.Nil(t, r, "no returned value expected")
}

func TestVisitorTreeErrorChild(t *testing.T) {
	ctx := newRuleContextWithErrorChild(81, 32, newTokenMock(87, 32, "test"))
	c := NewErrorItemCollection()
	v := NewVisitor(c)
	r := v.visitTree(ctx, 1, func(ctx antlr.ParserRuleContext, args []interface{}) (interface{}, error) {
		assert.Fail(t, "function must not be invoked")
		return "test result", nil
	})

	assert.False(t, c.HasErrors(), "no errors expected")
	assert.Nil(t, r, "no returned value expected")
}

func newRuleContext(line int, column int) *antlr.BaseParserRuleContext {
	c := antlr.NewBaseParserRuleContext(nil, -1)
	c.SetStart(newTokenMock(line, column, "_"))
	return c
}

func newRuleContextWithChildren(line int, column int, children []antlr.Token) *antlr.BaseParserRuleContext {
	c := newRuleContext(line, column)
	for _, child := range children {
		c.AddTokenNode(child)
	}
	return c
}

func newRuleContextWithErrorChild(line int, column int, child antlr.Token) *antlr.BaseParserRuleContext {
	c := newRuleContext(line, column)
	c.AddErrorNode(child)
	return c
}

type tokenMock struct {
	line   int
	column int
	text   string
}

func newTokenMock(line int, column int, text string) antlr.Token {
	return &tokenMock{line, column, text}
}

func (t *tokenMock) GetLine() int {
	return t.line
}

func (t *tokenMock) GetColumn() int {
	return t.column
}

func (t *tokenMock) GetSource() *antlr.TokenSourceCharStreamPair {
	panic("implement me")
}

func (t *tokenMock) GetTokenType() int {
	panic("implement me")
}

func (t *tokenMock) GetChannel() int {
	panic("implement me")
}

func (t *tokenMock) GetStart() int {
	panic("implement me")
}

func (t *tokenMock) GetStop() int {
	panic("implement me")
}

func (t *tokenMock) GetText() string {
	return t.text
}

func (t *tokenMock) SetText(string) {
	panic("implement me")
}

func (t *tokenMock) GetTokenIndex() int {
	panic("implement me")
}

func (t *tokenMock) SetTokenIndex(int) {
	panic("implement me")
}

func (t *tokenMock) GetTokenSource() antlr.TokenSource {
	panic("implement me")
}

func (t *tokenMock) GetInputStream() antlr.CharStream {
	panic("implement me")
}

type ruleNodeMock struct {
	children []antlr.Tree
}

func newRuleNodeMock(children []antlr.Tree) antlr.RuleNode {
	return &ruleNodeMock{children}
}

func (r ruleNodeMock) GetParent() antlr.Tree {
	panic("implement me")
}

func (r ruleNodeMock) SetParent(antlr.Tree) {
	panic("implement me")
}

func (r ruleNodeMock) GetPayload() interface{} {
	panic("implement me")
}

func (r ruleNodeMock) GetChild(i int) antlr.Tree {
	return r.children[i]
}

func (r ruleNodeMock) GetChildCount() int {
	return len(r.children)
}

func (r ruleNodeMock) GetChildren() []antlr.Tree {
	return r.children
}

func (r ruleNodeMock) GetSourceInterval() *antlr.Interval {
	panic("implement me")
}

func (r ruleNodeMock) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	return visitor.VisitChildren(r)
}

func (r ruleNodeMock) GetText() string {
	panic("implement me")
}

func (r ruleNodeMock) ToStringTree(strings []string, recognizer antlr.Recognizer) string {
	panic("implement me")
}

func (r ruleNodeMock) GetRuleContext() antlr.RuleContext {
	panic("implement me")
}

func (r ruleNodeMock) GetBaseRuleContext() *antlr.BaseRuleContext {
	panic("implement me")
}

type terminalNodeMock struct {
	text string
}

func newTerminalNodeMock(text string) antlr.TerminalNode {
	return &terminalNodeMock{text}
}

func (t terminalNodeMock) GetParent() antlr.Tree {
	panic("implement me")
}

func (t terminalNodeMock) SetParent(antlr.Tree) {
	panic("implement me")
}

func (t terminalNodeMock) GetPayload() interface{} {
	panic("implement me")
}

func (t terminalNodeMock) GetChild(int) antlr.Tree {
	panic("implement me")
}

func (t terminalNodeMock) GetChildCount() int {
	panic("implement me")
}

func (t terminalNodeMock) GetChildren() []antlr.Tree {
	panic("implement me")
}

func (t terminalNodeMock) GetSourceInterval() *antlr.Interval {
	panic("implement me")
}

func (t terminalNodeMock) Accept(antlr.ParseTreeVisitor) interface{} {
	panic("implement me")
}

func (t terminalNodeMock) GetText() string {
	return t.text
}

func (t terminalNodeMock) ToStringTree(strings []string, recognizer antlr.Recognizer) string {
	panic("implement me")
}

func (t terminalNodeMock) GetSymbol() antlr.Token {
	panic("implement me")
}

type otherNodeMock struct {
}

func newOtherNodeMock() antlr.Tree {
	return &otherNodeMock{}
}

func (e otherNodeMock) GetParent() antlr.Tree {
	panic("implement me")
}

func (e otherNodeMock) SetParent(antlr.Tree) {
	panic("implement me")
}

func (e otherNodeMock) GetPayload() interface{} {
	panic("implement me")
}

func (e otherNodeMock) GetChild(int) antlr.Tree {
	panic("implement me")
}

func (e otherNodeMock) GetChildCount() int {
	panic("implement me")
}

func (e otherNodeMock) GetChildren() []antlr.Tree {
	panic("implement me")
}
