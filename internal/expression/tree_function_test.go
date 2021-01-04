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

package expression

import (
	"github.com/healthiop/hipath/hipathsys"
	"github.com/healthiop/hipath/internal/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChildrenFuncNoTracer(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := childrenFunc
	res, err := f.Execute(ctx, nil, []interface{}{}, nil)

	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestChildrenFuncNoChildren(t *testing.T) {
	ctx := test.NewTestContext(t)

	n1 := hipathsys.NewString("test1")

	f := childrenFunc
	res, err := f.Execute(ctx, n1, []interface{}{}, nil)

	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestChildrenFuncColNoChildren(t *testing.T) {
	ctx := test.NewTestContext(t)

	n1 := ctx.NewCollection()
	n1.Add(hipathsys.NewString("test1"))
	n1.Add(hipathsys.NewString("test2"))

	f := childrenFunc
	res, err := f.Execute(ctx, n1, []interface{}{}, nil)

	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestChildrenFuncChildren(t *testing.T) {
	ctx := test.NewTestContext(t)

	n1 := hipathsys.NewString("test1")
	n2 := hipathsys.NewInteger(12)
	n3 := hipathsys.NewString("test3")

	l1c := map[string]interface{}{}
	l1c["x"] = n3

	l1 := map[string]interface{}{}
	l1["a"] = n1
	l1["b"] = nil
	l1["c"] = l1c
	l1["d"] = n2

	f := childrenFunc
	res, err := f.Execute(ctx, l1, []interface{}{}, nil)

	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.CollectionAccessor)(nil), res) {
		col := res.(hipathsys.CollectionAccessor)
		if assert.Equal(t, 3, col.Count()) {
			assert.Same(t, n1, col.Get(0))
			assert.Equal(t, l1c, col.Get(1))
			assert.Same(t, n2, col.Get(2))
		}
	}
}

func TestChildrenFuncColChildren(t *testing.T) {
	ctx := test.NewTestContext(t)

	n1 := hipathsys.NewString("test1")
	n2 := hipathsys.NewInteger(12)
	n3 := hipathsys.NewString("test3")
	n4 := hipathsys.NewString("test4")

	l1c := map[string]interface{}{}
	l1c["x"] = n3

	l1 := map[string]interface{}{}
	l1["a"] = n1
	l1["b"] = nil
	l1["c"] = l1c
	l1["d"] = n2

	l2 := map[string]interface{}{}
	l2["a"] = n4

	n := ctx.NewCollection()
	n.Add(l1)
	n.Add(l2)

	f := childrenFunc
	res, err := f.Execute(ctx, n, []interface{}{}, nil)

	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.CollectionAccessor)(nil), res) {
		col := res.(hipathsys.CollectionAccessor)
		if assert.Equal(t, 4, col.Count()) {
			assert.Same(t, n1, col.Get(0))
			assert.Equal(t, l1c, col.Get(1))
			assert.Same(t, n2, col.Get(2))
			assert.Same(t, n4, col.Get(3))
		}
	}
}

func TestChildrenFuncChildrenError(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := childrenFunc
	res, err := f.Execute(ctx, "test", []interface{}{}, nil)

	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestChildrenFuncColChildrenError(t *testing.T) {
	ctx := test.NewTestContext(t)

	n := ctx.NewCollection()
	n.Add(hipathsys.NewString("test1"))
	n.Add(test.NewTestModelNode(0, false))

	f := childrenFunc
	res, err := f.Execute(ctx, n, []interface{}{}, nil)

	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestDescendantsFuncNoTracer(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newDescendantsFunction()
	res, err := f.Execute(ctx, nil, []interface{}{}, nil)

	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestDescendantsFuncNoDescendants(t *testing.T) {
	ctx := test.NewTestContext(t)

	n1 := hipathsys.NewString("test1")

	f := newDescendantsFunction()
	res, err := f.Execute(ctx, n1, []interface{}{}, nil)

	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestDescendantsFuncColNoDescendants(t *testing.T) {
	ctx := test.NewTestContext(t)

	n1 := ctx.NewCollection()
	n1.Add(hipathsys.NewString("test1"))
	n1.Add(hipathsys.NewString("test2"))

	f := newDescendantsFunction()
	res, err := f.Execute(ctx, n1, []interface{}{}, nil)

	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestDescendantsFuncDescendants(t *testing.T) {
	ctx := test.NewTestContext(t)

	n1 := hipathsys.NewString("test1")
	n2 := hipathsys.NewInteger(12)
	n3 := hipathsys.NewString("test3")
	n4 := hipathsys.NewString("test4")

	l1cy := map[string]interface{}{}
	l1cy["x"] = n4

	l1c := map[string]interface{}{}
	l1c["x"] = n3
	l1c["y"] = l1cy

	l1 := map[string]interface{}{}
	l1["a"] = n1
	l1["b"] = nil
	l1["c"] = l1c
	l1["d"] = n2

	f := newDescendantsFunction()
	res, err := f.Execute(ctx, l1, []interface{}{}, nil)

	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.CollectionAccessor)(nil), res) {
		col := res.(hipathsys.CollectionAccessor)
		if assert.Equal(t, 6, col.Count()) {
			assert.Same(t, n1, col.Get(0))
			assert.Equal(t, l1c, col.Get(1))
			assert.Same(t, n3, col.Get(2))
			assert.Equal(t, l1cy, col.Get(3))
			assert.Same(t, n4, col.Get(4))
			assert.Same(t, n2, col.Get(5))
		}
	}
}

func TestDescendantsFuncColDescendants(t *testing.T) {
	ctx := test.NewTestContext(t)

	n1 := hipathsys.NewString("test1")
	n2 := hipathsys.NewInteger(12)
	n3 := hipathsys.NewString("test3")
	n4 := hipathsys.NewString("test4")

	l1c := map[string]interface{}{}
	l1c["x"] = n3

	l1 := map[string]interface{}{}
	l1["a"] = n1
	l1["b"] = nil
	l1["c"] = l1c
	l1["d"] = n2

	l2 := map[string]interface{}{}
	l2["a"] = n4

	n := ctx.NewCollection()
	n.Add(l1)
	n.Add(l2)

	f := newDescendantsFunction()
	res, err := f.Execute(ctx, n, []interface{}{}, nil)

	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.CollectionAccessor)(nil), res) {
		col := res.(hipathsys.CollectionAccessor)
		if assert.Equal(t, 5, col.Count()) {
			assert.Same(t, n1, col.Get(0))
			assert.Equal(t, l1c, col.Get(1))
			assert.Same(t, n3, col.Get(2))
			assert.Same(t, n2, col.Get(3))
			assert.Same(t, n4, col.Get(4))
		}
	}
}

func TestDescendantsFuncDescendantsError(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newDescendantsFunction()
	res, err := f.Execute(ctx, test.NewTestModelNode(0, false), []interface{}{}, nil)

	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestDescendantsFuncColDescendantsError(t *testing.T) {
	ctx := test.NewTestContext(t)

	n := ctx.NewCollection()
	n.Add(hipathsys.NewString("test1"))
	n.Add(test.NewTestModelNode(0, false))

	f := newDescendantsFunction()
	res, err := f.Execute(ctx, n, []interface{}{}, nil)

	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}
