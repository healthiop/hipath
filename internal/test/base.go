// Copyright (c) 2020-2021, Volker Schmidt (volker@volsch.eu)
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
	"fmt"
	"github.com/healthiop/hipath/hipathsys"
	"sort"
	"testing"
)

var testBaseTypeSpec = hipathsys.NewTypeSpec(hipathsys.NewFQTypeName("base", "TEST"))
var testTypeSpec = hipathsys.NewTypeSpecWithBase(hipathsys.NewFQTypeName("type1", "TEST"), testBaseTypeSpec)
var testElementTypeSpec = hipathsys.NewTypeSpecWithBase(hipathsys.NewFQTypeName("Element", "TEST"), testBaseTypeSpec)

type testModelNode struct {
	value float64
	sys   bool
}

type testModelErrorNode struct {
}

type testModelNodeAccessor interface {
	testValue() float64
	testSys() bool
}

type testModelErrorNodeAccessor interface {
	testErr() bool
}

func NewTestModelNode(value float64, sys bool) *testModelNode {
	return &testModelNode{value, sys}
}

func NewTestModelErrorNode() *testModelErrorNode {
	return &testModelErrorNode{}
}

func (n *testModelNode) testValue() float64 {
	return n.value
}

func (n *testModelNode) testSys() bool {
	return n.sys
}

func (n *testModelErrorNode) testErr() bool {
	return true
}

type testModel struct {
	t *testing.T
}

func newTestModel(t *testing.T) hipathsys.ModelAdapter {
	return &testModel{t}
}

func (a *testModel) CastToSystem(node interface{}) (hipathsys.AnyAccessor, error) {
	if _, ok := node.(map[string]interface{}); ok {
		return nil, nil
	}

	if n, ok := node.(testModelNodeAccessor); !ok {
		if _, ok = node.(testModelErrorNodeAccessor); ok {
			return nil, fmt.Errorf("error node cannot be converted")
		}
		a.t.Errorf("not a test model node: %T", node)
		return nil, nil
	} else {
		if n.testSys() {
			return hipathsys.NewDecimalFloat64(n.testValue()), nil
		}
		return nil, nil
	}
}

func (a *testModel) AsType(node interface{}, name hipathsys.FQTypeNameAccessor) (interface{}, error) {
	if n, ok := node.(testModelNodeAccessor); !ok {
		a.t.Errorf("not a test model node: %T", node)
		return nil, nil
	} else {
		if name.Namespace() == "Other" {
			return nil, fmt.Errorf("unsupported conversion")
		}
		if name.Name() == "decimal" && name.Namespace() == "Test" {
			return n, nil
		}
		if name.Name() == "Number" && name.Namespace() == "System" {
			return hipathsys.NewDecimalFloat64(n.testValue()), nil
		}
		return nil, nil
	}
}

func (a *testModel) TypeSpec(node interface{}) hipathsys.TypeSpecAccessor {
	if _, ok := node.(map[string]interface{}); ok {
		return testElementTypeSpec
	}

	if n, ok := node.(testModelNodeAccessor); !ok {
		if _, ok := node.(testModelErrorNodeAccessor); ok {
			return hipathsys.NewTypeSpec(hipathsys.NewFQTypeName("Error", "Test"))
		}
		if _, ok := node.(hipathsys.StringAccessor); !ok {
			a.t.Errorf("not a test model node: %T", node)
		}
		return hipathsys.UndefinedTypeSpec
	} else {
		if n.testSys() {
			a.t.Errorf("type of system node must not be requested")
		}
		return testTypeSpec
	}
}

func (a *testModel) Equal(node1 interface{}, node2 interface{}) bool {
	if m1, ok := node1.(map[string]interface{}); ok {
		if m2, ok := node2.(map[string]interface{}); !ok {
			return false
		} else {
			id := m1["id"]
			return id != nil && id == m2["id"]
		}
	}

	n1, ok := node1.(testModelNodeAccessor)
	if !ok {
		return false
	}
	n2, ok := node2.(testModelNodeAccessor)
	if !ok {
		return false
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

func (a *testModel) Navigate(node interface{}, name string) (interface{}, error) {
	model, ok := node.(map[string]interface{})
	if !ok {
		a.t.Fatal(fmt.Sprintf("cannot be cast to map: %T", node))
	}

	result, found := model[name]
	if !found {
		return nil, fmt.Errorf("path cannot be evaluated on model: %s", name)
	}
	return result, nil
}

func (a *testModel) Children(node interface{}) (hipathsys.ColAccessor, error) {
	if _, ok := node.(hipathsys.AnyAccessor); ok {
		return nil, nil
	}

	model, ok := node.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("cannot be cast to map: %T", node)
	}

	res := hipathsys.NewCol(a)
	keys := make([]string, 0)
	for k := range model {
		if k == "errorCollection" {
			return NewErrorCollection(), nil
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := model[k]
		if v != nil {
			res.Add(model[k])
		}
	}
	return res, nil
}

type testContext struct {
	modelAdapter hipathsys.ModelAdapter
	tracer       hipathsys.Tracer
	node         interface{}
}

func NewTestContext(t *testing.T) hipathsys.ContextAccessor {
	return &testContext{modelAdapter: newTestModel(t)}
}

func NewTestContextWithNode(t *testing.T, node interface{}) hipathsys.ContextAccessor {
	return &testContext{modelAdapter: newTestModel(t), node: node}
}

func NewTestContextWithNodeAndTracer(t *testing.T, node interface{}, tracer hipathsys.Tracer) hipathsys.ContextAccessor {
	return &testContext{
		modelAdapter: newTestModel(t),
		tracer:       tracer,
		node:         node,
	}
}

func (t *testContext) EnvVar(name string) (interface{}, bool) {
	if name == "ucum" {
		return hipathsys.UCUMSystemURI, true
	}
	return nil, false
}

func (t *testContext) ModelAdapter() hipathsys.ModelAdapter {
	return t.modelAdapter
}

func (t *testContext) NewCol() hipathsys.ColModifier {
	return hipathsys.NewCol(t.modelAdapter)
}

func (t *testContext) NewColWithItem(item interface{}) hipathsys.ColModifier {
	return hipathsys.NewColWithItem(t.modelAdapter, item)
}

func (t *testContext) ContextNode() interface{} {
	return t.node
}

func (t *testContext) Tracer() hipathsys.Tracer {
	return t.tracer
}

type errorCollection struct {
}

func NewErrorCollection() hipathsys.ColAccessor {
	return &errorCollection{}
}

func (c *errorCollection) DataType() hipathsys.DataTypes {
	return hipathsys.ColDataType
}

func (c *errorCollection) Empty() bool {
	return false
}

func (c *errorCollection) Count() int {
	return 3
}

func (c *errorCollection) Get(i int) interface{} {
	switch i {
	case 0:
		return hipathsys.NewString("test item 1")
	case 1:
		return NewTestModelErrorNode()
	case 2:
		return hipathsys.NewString("test item 2")
	default:
		panic("invalid item index")
	}
}

func (c *errorCollection) Contains(node interface{}) bool {
	if _, ok := node.(testModelErrorNodeAccessor); ok {
		return true
	}
	return false
}

func (c *errorCollection) TypeSpec() hipathsys.TypeSpecAccessor {
	panic("implement me")
}

func (c *errorCollection) TypeInfo() hipathsys.TypeInfoAccessor {
	panic("implement me")
}

func (c *errorCollection) Source() interface{} {
	panic("implement me")
}

func (c *errorCollection) Equal(_ interface{}) bool {
	panic("implement me")
}

func (c *errorCollection) Equivalent(_ interface{}) bool {
	panic("implement me")
}

func (c *errorCollection) ItemTypeSpec() hipathsys.TypeSpecAccessor {
	panic("implement me")
}
