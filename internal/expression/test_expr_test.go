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

package expression

import (
	"fmt"
	"github.com/volsch/gohipath/pathsys"
	"testing"
)

type testExpression struct {
	result          interface{}
	invocationCount int
	node            interface{}
	loop            pathsys.Looper
}

func newTestExpression(result interface{}) *testExpression {
	return &testExpression{result: result}
}

func (e *testExpression) Evaluate(ctx pathsys.ContextAccessor, node interface{}, loop pathsys.Looper) (interface{}, error) {
	e.invocationCount = e.invocationCount + 1
	e.node = node
	e.loop = loop
	return e.result, nil
}

type testErrorExpression struct {
}

func newTestErrorExpression() *testErrorExpression {
	return &testErrorExpression{}
}

func (e *testErrorExpression) Evaluate(pathsys.ContextAccessor, interface{}, pathsys.Looper) (interface{}, error) {
	return nil, fmt.Errorf("an error occurred")
}

type testingAccessor interface {
	pathsys.AnyAccessor
	testing() *testing.T
}

type testingType struct {
	t *testing.T
}

func newTestingType(t *testing.T) testingAccessor {
	return &testingType{t}
}

func extractTesting(node interface{}) *testing.T {
	if node == nil {
		panic("cannot extract testing from nil")
	}
	if t, ok := node.(testingAccessor); !ok {
		panic(fmt.Sprintf("cannot extract testing from %T", node))
	} else {
		return t.testing()
	}
}

func (t *testingType) testing() *testing.T {
	return t.t
}

func (t *testingType) DataType() pathsys.DataTypes {
	panic("implement me")
}

func (t *testingType) TypeInfo() pathsys.TypeInfoAccessor {
	panic("implement me")
}

func (t *testingType) Empty() bool {
	panic("implement me")
}

func (t *testingType) Equal(interface{}) bool {
	panic("implement me")
}

func (t *testingType) Equivalent(interface{}) bool {
	panic("implement me")
}
