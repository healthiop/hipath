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
	"github.com/stretchr/testify/assert"
	"github.com/volsch/gohipath/internal/test"
	"github.com/volsch/gohipath/pathsys"
	"testing"
)

func TestUnionPathFunc(t *testing.T) {
	ctx := test.NewTestContext(t)
	c1 := ctx.NewCollection()
	c1.Add(pathsys.NewInteger(10))
	c1.Add(pathsys.NewInteger(11))
	c1.Add(pathsys.NewInteger(14))

	c2 := ctx.NewCollection()
	c2.Add(pathsys.NewDecimalInt(11))
	c2.Add(pathsys.NewDecimalInt(12))

	f := newUnionFunction()
	res, err := f.Execute(ctx, c1, []interface{}{c2}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.CollectionAccessor)(nil), res) {
		c := res.(pathsys.CollectionAccessor)
		if assert.Equal(t, 4, c.Count()) {
			assert.Equal(t, pathsys.NewInteger(10), c.Get(0))
			assert.Equal(t, pathsys.NewInteger(11), c.Get(1))
			assert.Equal(t, pathsys.NewInteger(14), c.Get(2))
			assert.Condition(t, func() bool {
				return pathsys.NewDecimalInt(12).Equal(c.Get(3))
			})
		}
		assert.Equal(t, "System.Any", c.ItemTypeInfo().String())
	}
}

func TestUnionPathFuncNoCollection(t *testing.T) {
	ctx := test.NewTestContext(t)
	c1 := pathsys.NewString("test1")
	c2 := pathsys.NewString("test2")

	f := newUnionFunction()
	res, err := f.Execute(ctx, c1, []interface{}{c2}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.CollectionAccessor)(nil), res) {
		c := res.(pathsys.CollectionAccessor)
		if assert.Equal(t, 2, c.Count()) {
			assert.Equal(t, pathsys.NewString("test1"), c.Get(0))
			assert.Equal(t, pathsys.NewString("test2"), c.Get(1))
		}
		assert.Equal(t, "System.String", c.ItemTypeInfo().String())
	}
}

func TestUnionPathFuncArgNil(t *testing.T) {
	ctx := test.NewTestContext(t)
	c1 := pathsys.NewString("test1")

	f := newUnionFunction()
	res, err := f.Execute(ctx, c1, []interface{}{nil}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.CollectionAccessor)(nil), res) {
		c := res.(pathsys.CollectionAccessor)
		if assert.Equal(t, 1, c.Count()) {
			assert.Equal(t, pathsys.NewString("test1"), c.Get(0))
		}
		assert.Equal(t, "System.String", c.ItemTypeInfo().String())
	}
}

func TestUnionPathFuncObjNil(t *testing.T) {
	ctx := test.NewTestContext(t)
	c1 := pathsys.NewString("test1")

	f := newUnionFunction()
	res, err := f.Execute(ctx, nil, []interface{}{c1}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.CollectionAccessor)(nil), res) {
		c := res.(pathsys.CollectionAccessor)
		if assert.Equal(t, 1, c.Count()) {
			assert.Equal(t, pathsys.NewString("test1"), c.Get(0))
		}
		assert.Equal(t, "System.String", c.ItemTypeInfo().String())
	}
}

func TestUnionPathFuncBothNil(t *testing.T) {
	f := newUnionFunction()
	res, err := f.Execute(nil, nil, []interface{}{nil}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty res expected")
}

func TestUnionPathFuncBothEmpty(t *testing.T) {
	ctx := test.NewTestContext(t)
	f := newUnionFunction()
	res, err := f.Execute(ctx, ctx.NewCollection(), []interface{}{ctx.NewCollection()}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty res expected")
}

func TestCombinePathFunc(t *testing.T) {
	ctx := test.NewTestContext(t)
	c1 := ctx.NewCollection()
	c1.Add(pathsys.NewInteger(10))
	c1.Add(pathsys.NewInteger(11))
	c1.Add(pathsys.NewInteger(14))

	c2 := ctx.NewCollection()
	c2.Add(pathsys.NewDecimalInt(11))
	c2.Add(pathsys.NewDecimalInt(12))

	f := newCombineFunction()
	res, err := f.Execute(ctx, c1, []interface{}{c2}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.CollectionAccessor)(nil), res) {
		c := res.(pathsys.CollectionAccessor)
		if assert.Equal(t, 5, c.Count()) {
			assert.Equal(t, pathsys.NewInteger(10), c.Get(0))
			assert.Equal(t, pathsys.NewInteger(11), c.Get(1))
			assert.Equal(t, pathsys.NewInteger(14), c.Get(2))
			assert.Condition(t, func() bool {
				return pathsys.NewDecimalInt(11).Equal(c.Get(3))
			})
			assert.Condition(t, func() bool {
				return pathsys.NewDecimalInt(12).Equal(c.Get(4))
			})
		}
		assert.Equal(t, "System.Any", c.ItemTypeInfo().String())
	}
}

func TestCombinePathFuncNoCollection(t *testing.T) {
	ctx := test.NewTestContext(t)
	c1 := pathsys.NewString("test1")
	c2 := pathsys.NewString("test2")

	f := newCombineFunction()
	res, err := f.Execute(ctx, c1, []interface{}{c2}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.CollectionAccessor)(nil), res) {
		c := res.(pathsys.CollectionAccessor)
		if assert.Equal(t, 2, c.Count()) {
			assert.Equal(t, pathsys.NewString("test1"), c.Get(0))
			assert.Equal(t, pathsys.NewString("test2"), c.Get(1))
		}
		assert.Equal(t, "System.String", c.ItemTypeInfo().String())
	}
}

func TestCombinePathFuncArgNil(t *testing.T) {
	ctx := test.NewTestContext(t)
	c1 := pathsys.NewString("test1")

	f := newCombineFunction()
	res, err := f.Execute(ctx, c1, []interface{}{nil}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.CollectionAccessor)(nil), res) {
		c := res.(pathsys.CollectionAccessor)
		if assert.Equal(t, 1, c.Count()) {
			assert.Equal(t, pathsys.NewString("test1"), c.Get(0))
		}
		assert.Equal(t, "System.String", c.ItemTypeInfo().String())
	}
}

func TestCombinePathFuncObjNil(t *testing.T) {
	ctx := test.NewTestContext(t)
	c1 := pathsys.NewString("test1")

	f := newCombineFunction()
	res, err := f.Execute(ctx, nil, []interface{}{c1}, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*pathsys.CollectionAccessor)(nil), res) {
		c := res.(pathsys.CollectionAccessor)
		if assert.Equal(t, 1, c.Count()) {
			assert.Equal(t, pathsys.NewString("test1"), c.Get(0))
		}
		assert.Equal(t, "System.String", c.ItemTypeInfo().String())
	}
}

func TestCombinePathFuncBothNil(t *testing.T) {
	ctx := test.NewTestContext(t)
	f := newCombineFunction()
	res, err := f.Execute(ctx, nil, []interface{}{nil}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty res expected")
}

func TestCombinePathFuncBothEmpty(t *testing.T) {
	ctx := test.NewTestContext(t)
	f := newCombineFunction()
	res, err := f.Execute(ctx, ctx.NewCollection(), []interface{}{ctx.NewCollection()}, nil)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "empty res expected")
}
