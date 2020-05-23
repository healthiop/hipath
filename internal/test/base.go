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

package test

import (
	"github.com/volsch/gohipath/pathsys"
	"testing"
)

var testBaseTypeInfo = pathsys.NewTypeInfo(pathsys.NewFQTypeName("base", "TEST"))
var testTypeInfo = pathsys.NewTypeInfoWithBase(pathsys.NewFQTypeName("type1", "TEST"), testBaseTypeInfo)

type testModelNode struct {
	value float64
	sys   bool
}

type testModelNodeAccessor interface {
	testValue() float64
	testSys() bool
}

func NewTestModelNode(value float64, sys bool) *testModelNode {
	return &testModelNode{value, sys}
}

func (n *testModelNode) testValue() float64 {
	return n.value
}

func (n *testModelNode) testSys() bool {
	return n.sys
}

type testModel struct {
	t *testing.T
}

func newTestModel(t *testing.T) pathsys.ModelAdapter {
	return &testModel{t}
}

func (a *testModel) ConvertToSystem(node interface{}) interface{} {
	if n, ok := node.(testModelNodeAccessor); !ok {
		a.t.Errorf("not a test model node: %T", node)
		return nil
	} else {
		if n.testSys() {
			return pathsys.NewDecimalFloat64(n.testValue())
		}
		return n
	}
}

func (a *testModel) TypeInfo(node interface{}) pathsys.TypeInfoAccessor {
	if n, ok := node.(testModelNodeAccessor); !ok {
		a.t.Errorf("not a test model node: %T", node)
		return pathsys.UndefinedTypeInfo
	} else {
		if n.testSys() {
			a.t.Errorf("type of system node must not be requested")
		}
		return testTypeInfo
	}
}

func (a *testModel) Equal(node1 interface{}, node2 interface{}) bool {
	n1, ok := node1.(testModelNodeAccessor)
	if !ok {
		a.t.Errorf("not a test model node: %T", node1)
	}
	n2, ok := node2.(testModelNodeAccessor)
	if !ok {
		a.t.Errorf("not a test model node: %T", node2)
	}

	if n1.testSys() || n2.testSys() {
		a.t.Errorf("equality of system node must not be requested")
	}
	return n1.testValue() == n2.testValue()
}

func (a *testModel) Equivalent(node1 interface{}, node2 interface{}) bool {
	n1, ok := node1.(testModelNodeAccessor)
	if !ok {
		a.t.Errorf("not a test model node: %T", node1)
	}
	n2, ok := node2.(testModelNodeAccessor)
	if !ok {
		a.t.Errorf("not a test model node: %T", node2)
	}

	if n1.testSys() || n2.testSys() {
		a.t.Errorf("equality of system node must not be requested")
	}
	return int64(n1.testValue()) == int64(n2.testValue())
}

type testContext struct {
	modelAdapter pathsys.ModelAdapter
	node         interface{}
}

func NewTestContext(t *testing.T) pathsys.ContextAccessor {
	return &testContext{modelAdapter: newTestModel(t)}
}

func NewTestContextWithNode(t *testing.T, node interface{}) pathsys.ContextAccessor {
	return &testContext{modelAdapter: newTestModel(t), node: node}
}

func (t *testContext) EnvVar(name string) (interface{}, bool) {
	if name == "ucum" {
		return pathsys.UCUMSystemURI, true
	}
	return nil, false
}

func (t *testContext) ModelAdapter() pathsys.ModelAdapter {
	return t.modelAdapter
}

func (t *testContext) NewCollection() pathsys.CollectionModifier {
	return pathsys.NewCollection(t.modelAdapter)
}

func (t *testContext) ContextNode() interface{} {
	return t.node
}
