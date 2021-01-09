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

func TestExtractIdentifier(t *testing.T) {
	assert.Equal(t, "test", ExtractIdentifier("test"))
}

func TestExtractIdentifierDelimited(t *testing.T) {
	assert.Equal(t, "test", ExtractIdentifier("`test'"))
}

func TestUnwrapCollectionNil(t *testing.T) {
	assert.Nil(t, unwrapCollection(nil))
}

func TestUnwrapCollectionZero(t *testing.T) {
	ctx := test.NewTestContext(t)
	assert.Nil(t, unwrapCollection(ctx.NewCollection()))
}

func TestUnwrapCollectionOne(t *testing.T) {
	ctx := test.NewTestContext(t)
	i := hipathsys.NewString("test")
	c := ctx.NewCollection()
	c.MustAdd(i)

	assert.Same(t, i, unwrapCollection(c))
}

func TestUnwrapCollectionMore(t *testing.T) {
	ctx := test.NewTestContext(t)
	c := ctx.NewCollection()
	c.MustAdd(hipathsys.NewString("test1"))
	c.MustAdd(hipathsys.NewString("test2"))

	assert.Same(t, c, unwrapCollection(c))
}

func TestWrapCollectionNil(t *testing.T) {
	ctx := test.NewTestContext(t)
	col, err := wrapCollection(ctx, nil)
	assert.NoError(t, err)
	if assert.NotNil(t, col, "collection expected") {
		assert.True(t, col.Empty())
	}
}

func TestWrapCollectionCollection(t *testing.T) {
	ctx := test.NewTestContext(t)
	col := ctx.NewCollection()
	wrapped, err := wrapCollection(ctx, col)
	assert.NoError(t, err)
	assert.Same(t, col, wrapped)
}

func TestWrapCollectionNoCollection(t *testing.T) {
	ctx := test.NewTestContext(t)
	item := hipathsys.NewString("test")
	res, err := wrapCollection(ctx, item)
	assert.NoError(t, err)
	if assert.Implements(t, (*hipathsys.CollectionAccessor)(nil), res) {
		col := res.(hipathsys.CollectionAccessor)
		if assert.Equal(t, 1, col.Count()) {
			assert.Same(t, item, col.Get(0))
		}
	}
}

func TestEmptyCollectionNil(t *testing.T) {
	assert.True(t, emptyCollection(nil))
}

func TestEmptyCollectionEmpty(t *testing.T) {
	ctx := test.NewTestContext(t)
	col := ctx.NewCollection()
	assert.True(t, emptyCollection(col))
}

func TestEmptyCollectionNotEmpty(t *testing.T) {
	ctx := test.NewTestContext(t)
	col := ctx.NewCollection()
	col.MustAdd(hipathsys.NewString("test"))
	assert.False(t, emptyCollection(col))
}

func TestEmptyCollectionOther(t *testing.T) {
	assert.False(t, emptyCollection(hipathsys.NewString("test")))
}

func TestUniteCollectionsLeftError(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, err := uniteCollections(ctx, test.NewTestModelErrorNode(), ctx.NewCollection())
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestUniteCollectionsLeftColError(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, err := uniteCollections(ctx, test.NewErrorCollection(), ctx.NewCollection())
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestUniteCollectionsRightError(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, err := uniteCollections(ctx, ctx.NewCollection(), test.NewTestModelErrorNode())
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestUniteCollectionsRightColError(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, err := uniteCollections(ctx, ctx.NewCollection(), test.NewErrorCollection())
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestCombineCollectionsLeftError(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, err := combineCollections(ctx, test.NewTestModelErrorNode(), ctx.NewCollection())
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestCombineCollectionsLeftColError(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, err := combineCollections(ctx, test.NewErrorCollection(), ctx.NewCollection())
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestCombineCollectionsRightError(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, err := combineCollections(ctx, ctx.NewCollection(), test.NewTestModelErrorNode())
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}

func TestCombineCollectionsRightColError(t *testing.T) {
	ctx := test.NewTestContext(t)
	res, err := combineCollections(ctx, ctx.NewCollection(), test.NewErrorCollection())
	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "empty collection expected")
}
