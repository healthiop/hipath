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

func TestCollectionExpressionEmpty(t *testing.T) {
	e := NewCollectionExpression(NewEmptyLiteral())
	node, err := e.Evaluate(nil, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.CollectionAccessor)(nil), node) {
		res := node.(hipathsys.CollectionAccessor)
		assert.Equal(t, 0, res.Count())
	}
}

func TestCollectionExpressionSingle(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewCollectionExpression(NewRawStringLiteral("test 123"))
	node, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.CollectionAccessor)(nil), node) {
		res := node.(hipathsys.CollectionAccessor)
		if assert.Equal(t, 1, res.Count()) {
			assert.Equal(t, hipathsys.NewString("test 123"), res.Get(0))
		}
	}
}

func TestCollectionExpressionMulti(t *testing.T) {
	ctx := test.NewTestContext(t)
	e := NewCollectionExpression(NewUnionExpression(
		NewRawStringLiteral("test 1"), NewRawStringLiteral("test 3")))
	node, err := e.Evaluate(ctx, nil, nil)
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.CollectionAccessor)(nil), node) {
		res := node.(hipathsys.CollectionAccessor)
		if assert.Equal(t, 2, res.Count()) {
			assert.Equal(t, hipathsys.NewString("test 1"), res.Get(0))
			assert.Equal(t, hipathsys.NewString("test 3"), res.Get(1))
		}
	}
}

func TestCollectionExpressionError(t *testing.T) {
	e := NewCollectionExpression(newTestErrorExpression())
	node, err := e.Evaluate(nil, nil, nil)
	assert.Error(t, err, "error expected")
	assert.Nil(t, node, "empty result expected")
}
