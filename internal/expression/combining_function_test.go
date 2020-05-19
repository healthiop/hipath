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
	"github.com/volsch/gohimodel/datatype"
	"testing"
)

func TestUnionPathFunc(t *testing.T) {
	c1 := datatype.NewCollectionUndefined()
	c1.Add(datatype.NewPositiveInt(10))
	c1.Add(datatype.NewPositiveInt(11))
	c1.Add(datatype.NewPositiveInt(14))

	c2 := datatype.NewCollectionUndefined()
	c2.Add(datatype.NewUnsignedInt(11))
	c2.Add(datatype.NewUnsignedInt(12))

	result, err := unionPathFunc(nil, c1, []datatype.Accessor{c2})
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*datatype.CollectionAccessor)(nil), result) {
		c := result.(datatype.CollectionAccessor)
		if assert.Equal(t, 4, c.Count()) {
			assert.Equal(t, datatype.NewPositiveInt(10), c.Get(0))
			assert.Equal(t, datatype.NewPositiveInt(11), c.Get(1))
			assert.Equal(t, datatype.NewPositiveInt(14), c.Get(2))
			assert.Equal(t, datatype.NewUnsignedInt(12), c.Get(3))
		}
		assert.Equal(t, "FHIR.integer", c.ItemTypeInfo().String())
	}
}

func TestUnionPathFuncNoCollection(t *testing.T) {
	c1 := datatype.NewString("test1")
	c2 := datatype.NewString("test2")

	result, err := unionPathFunc(nil, c1, []datatype.Accessor{c2})
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*datatype.CollectionAccessor)(nil), result) {
		c := result.(datatype.CollectionAccessor)
		if assert.Equal(t, 2, c.Count()) {
			assert.Equal(t, datatype.NewString("test1"), c.Get(0))
			assert.Equal(t, datatype.NewString("test2"), c.Get(1))
		}
		assert.Equal(t, "FHIR.string", c.ItemTypeInfo().String())
	}
}

func TestUnionPathFuncArgNil(t *testing.T) {
	c1 := datatype.NewString("test1")

	result, err := unionPathFunc(nil, c1, []datatype.Accessor{nil})
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*datatype.CollectionAccessor)(nil), result) {
		c := result.(datatype.CollectionAccessor)
		if assert.Equal(t, 1, c.Count()) {
			assert.Equal(t, datatype.NewString("test1"), c.Get(0))
		}
		assert.Equal(t, "FHIR.string", c.ItemTypeInfo().String())
	}
}

func TestUnionPathFuncObjNil(t *testing.T) {
	c1 := datatype.NewString("test1")

	result, err := unionPathFunc(nil, nil, []datatype.Accessor{c1})
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*datatype.CollectionAccessor)(nil), result) {
		c := result.(datatype.CollectionAccessor)
		if assert.Equal(t, 1, c.Count()) {
			assert.Equal(t, datatype.NewString("test1"), c.Get(0))
		}
		assert.Equal(t, "FHIR.string", c.ItemTypeInfo().String())
	}
}

func TestUnionPathFuncBothNil(t *testing.T) {
	result, err := unionPathFunc(nil, nil, []datatype.Accessor{nil})
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, result, "empty result expected")
}

func TestUnionPathFuncBothEmpty(t *testing.T) {
	result, err := unionPathFunc(nil, datatype.NewCollectionUndefined(),
		[]datatype.Accessor{datatype.NewCollectionUndefined()})
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, result, "empty result expected")
}

func TestCombinePathFunc(t *testing.T) {
	c1 := datatype.NewCollectionUndefined()
	c1.Add(datatype.NewPositiveInt(10))
	c1.Add(datatype.NewPositiveInt(11))
	c1.Add(datatype.NewPositiveInt(14))

	c2 := datatype.NewCollectionUndefined()
	c2.Add(datatype.NewUnsignedInt(11))
	c2.Add(datatype.NewUnsignedInt(12))

	result, err := combinePathFunc(nil, c1, []datatype.Accessor{c2})
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*datatype.CollectionAccessor)(nil), result) {
		c := result.(datatype.CollectionAccessor)
		if assert.Equal(t, 5, c.Count()) {
			assert.Equal(t, datatype.NewPositiveInt(10), c.Get(0))
			assert.Equal(t, datatype.NewPositiveInt(11), c.Get(1))
			assert.Equal(t, datatype.NewPositiveInt(14), c.Get(2))
			assert.Equal(t, datatype.NewUnsignedInt(11), c.Get(3))
			assert.Equal(t, datatype.NewUnsignedInt(12), c.Get(4))
		}
		assert.Equal(t, "FHIR.integer", c.ItemTypeInfo().String())
	}
}

func TestCombinePathFuncNoCollection(t *testing.T) {
	c1 := datatype.NewString("test1")
	c2 := datatype.NewString("test2")

	result, err := combinePathFunc(nil, c1, []datatype.Accessor{c2})
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*datatype.CollectionAccessor)(nil), result) {
		c := result.(datatype.CollectionAccessor)
		if assert.Equal(t, 2, c.Count()) {
			assert.Equal(t, datatype.NewString("test1"), c.Get(0))
			assert.Equal(t, datatype.NewString("test2"), c.Get(1))
		}
		assert.Equal(t, "FHIR.string", c.ItemTypeInfo().String())
	}
}

func TestCombinePathFuncArgNil(t *testing.T) {
	c1 := datatype.NewString("test1")

	result, err := combinePathFunc(nil, c1, []datatype.Accessor{nil})
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*datatype.CollectionAccessor)(nil), result) {
		c := result.(datatype.CollectionAccessor)
		if assert.Equal(t, 1, c.Count()) {
			assert.Equal(t, datatype.NewString("test1"), c.Get(0))
		}
		assert.Equal(t, "FHIR.string", c.ItemTypeInfo().String())
	}
}

func TestCombinePathFuncObjNil(t *testing.T) {
	c1 := datatype.NewString("test1")

	result, err := combinePathFunc(nil, nil, []datatype.Accessor{c1})
	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*datatype.CollectionAccessor)(nil), result) {
		c := result.(datatype.CollectionAccessor)
		if assert.Equal(t, 1, c.Count()) {
			assert.Equal(t, datatype.NewString("test1"), c.Get(0))
		}
		assert.Equal(t, "FHIR.string", c.ItemTypeInfo().String())
	}
}

func TestCombinePathFuncBothNil(t *testing.T) {
	result, err := combinePathFunc(nil, nil, []datatype.Accessor{nil})
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, result, "empty result expected")
}

func TestCombinePathFuncBothEmpty(t *testing.T) {
	result, err := combinePathFunc(nil, datatype.NewCollectionUndefined(),
		[]datatype.Accessor{datatype.NewCollectionUndefined()})
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, result, "empty result expected")
}
