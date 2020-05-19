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
	"github.com/volsch/gohimodel/resource"
	"testing"
)

func TestConvertContextDataNil(t *testing.T) {
	assert.Nil(t, convertContextData(nil))
}

func TestConvertContextDataUnconverted(t *testing.T) {
	assert.Equal(t, datatype.NewString("test"), convertContextData(datatype.NewString("test")))
}

func TestConvertContextDataQuantityNoUCUM(t *testing.T) {
	q := datatype.NewQuantity(datatype.NewDecimalInt(1026), datatype.LessThanQuantityComparator,
		datatype.NewString("a"), nil, datatype.NewCode("a"))
	assert.Same(t, q, convertContextData(q))
}

func TestConvertContextDataQuantityOtherUCUM(t *testing.T) {
	q := datatype.NewQuantity(datatype.NewDecimalInt(1026), datatype.LessThanQuantityComparator,
		datatype.NewString("cm"), datatype.UCUMSystemURI, datatype.NewCode("cm"))
	assert.Same(t, q, convertContextData(q))
}

func TestConvertContextDataQuantityYear(t *testing.T) {
	q := datatype.NewQuantity(datatype.NewDecimalInt(1026), datatype.LessThanQuantityComparator,
		datatype.NewCode("Year"), datatype.UCUMSystemURI, datatype.NewCode("a"))
	e := datatype.NewQuantity(datatype.NewDecimalInt(1026), datatype.LessThanQuantityComparator,
		datatype.NewCode("Year"), nil, datatype.NewCode("year"))
	assert.Equal(t, e, convertContextData(q))
}

func TestConvertContextDataQuantityMonth(t *testing.T) {
	q := datatype.NewQuantity(datatype.NewDecimalInt(1026), nil, nil,
		datatype.UCUMSystemURI, datatype.NewCode("mo"))
	e := datatype.NewQuantity(datatype.NewDecimalInt(1026), nil, nil,
		nil, datatype.NewCode("month"))
	assert.Equal(t, e, convertContextData(q))
}

func TestConvertContextDataQuantityDay(t *testing.T) {
	q := datatype.NewQuantity(datatype.NewDecimalInt(1026), nil, nil,
		datatype.UCUMSystemURI, datatype.NewCode("d"))
	e := datatype.NewQuantity(datatype.NewDecimalInt(1026), nil, nil,
		nil, datatype.NewCode("day"))
	assert.Equal(t, e, convertContextData(q))
}

func TestConvertContextDataQuantityHour(t *testing.T) {
	q := datatype.NewQuantity(datatype.NewDecimalInt(1026), nil, nil,
		datatype.UCUMSystemURI, datatype.NewCode("h"))
	e := datatype.NewQuantity(datatype.NewDecimalInt(1026), nil, nil,
		nil, datatype.NewCode("hour"))
	assert.Equal(t, e, convertContextData(q))
}

func TestConvertContextDataQuantityMinute(t *testing.T) {
	q := datatype.NewQuantity(datatype.NewDecimalInt(1026), nil, nil,
		datatype.UCUMSystemURI, datatype.NewCode("min"))
	e := datatype.NewQuantity(datatype.NewDecimalInt(1026), nil, nil,
		nil, datatype.NewCode("minute"))
	assert.Equal(t, e, convertContextData(q))
}

func TestConvertContextDataQuantitySecond(t *testing.T) {
	q := datatype.NewQuantity(datatype.NewDecimalInt(1026), nil, nil,
		datatype.UCUMSystemURI, datatype.NewCode("s"))
	e := datatype.NewQuantity(datatype.NewDecimalInt(1026), nil, nil,
		nil, datatype.NewCode("second"))
	assert.Equal(t, e, convertContextData(q))
}

func TestUnwrapCollectionNil(t *testing.T) {
	assert.Nil(t, unwrapCollection(nil))
}

func TestUnwrapCollectionZero(t *testing.T) {
	assert.Nil(t, unwrapCollection(datatype.NewCollectionUndefined()))
}

func TestUnwrapCollectionOne(t *testing.T) {
	i := datatype.NewString("test")
	c := datatype.NewStringCollection()
	c.Add(i)

	assert.Same(t, i, unwrapCollection(c))
}

func TestUnwrapCollectionMore(t *testing.T) {
	c := datatype.NewStringCollection()
	c.Add(datatype.NewString("test1"))
	c.Add(datatype.NewString("test2"))

	assert.Same(t, c, unwrapCollection(c))
}

func TestCommonAccessorBaseTypeColOnly(t *testing.T) {
	c1 := datatype.NewCollectionUndefined()
	c1.Add(datatype.NewPositiveInt(10))
	c1.Add(datatype.NewPositiveInt(11))

	c2 := datatype.NewCollectionUndefined()
	c2.Add(datatype.NewUnsignedInt(12))
	c2.Add(datatype.NewUnsignedInt(14))

	ti := commonAccessorBaseType([]datatype.Accessor{c1, c2})
	if assert.NotNil(t, ti, "common type expected") {
		assert.Equal(t, "FHIR.integer", ti.String())
	}
}

func TestCommonAccessorBaseTypeColOnlyNone(t *testing.T) {
	c1 := datatype.NewCollectionUndefined()
	c1.Add(datatype.NewPositiveInt(10))
	c1.Add(datatype.NewPositiveInt(11))

	c2 := datatype.NewCollectionUndefined()
	c2.Add(resource.NewDynamicResource("Patient"))

	ti := commonAccessorBaseType([]datatype.Accessor{c1, c2})
	assert.Nil(t, ti, "no common type expected")
}

func TestCommonAccessorBaseTypeDiffer(t *testing.T) {
	a1 := datatype.NewPositiveInt(10)
	a2 := datatype.NewUnsignedInt(12)

	ti := commonAccessorBaseType([]datatype.Accessor{a1, a2})
	if assert.NotNil(t, ti, "common type expected") {
		assert.Equal(t, "FHIR.integer", ti.String())
	}
}

func TestCommonAccessorBaseTypeSame(t *testing.T) {
	a1 := datatype.NewUnsignedInt(10)
	a2 := datatype.NewUnsignedInt(12)

	ti := commonAccessorBaseType([]datatype.Accessor{a1, a2})
	if assert.NotNil(t, ti, "common type expected") {
		assert.Equal(t, "FHIR.unsignedInt", ti.String())
	}
}

func TestCommonAccessorBaseTypeNone(t *testing.T) {
	a1 := datatype.NewPositiveInt(10)
	a2 := resource.NewDynamicResource("Patient")

	ti := commonAccessorBaseType([]datatype.Accessor{a1, a2})
	assert.Nil(t, ti, "no common type expected")
}

func TestCommonAccessorBaseTypeMixed(t *testing.T) {
	c1 := datatype.NewCollectionUndefined()
	c1.Add(datatype.NewPositiveInt(10))
	c1.Add(datatype.NewPositiveInt(11))

	a2 := datatype.NewString("test")

	ti := commonAccessorBaseType([]datatype.Accessor{c1, a2})
	if assert.NotNil(t, ti, "common type expected") {
		assert.Equal(t, "FHIR.Element", ti.String())
	}
}

func TestNewCollectionWithAccessorType(t *testing.T) {
	c1 := datatype.NewCollectionUndefined()
	c1.Add(datatype.NewPositiveInt(10))
	c1.Add(datatype.NewPositiveInt(11))

	a2 := datatype.NewUnsignedInt(14)

	result := newCollectionWithAccessorTypes([]datatype.Accessor{c1, a2})
	assert.Equal(t, "FHIR.integer", result.ItemTypeInfo().String())
}

func TestNewCollectionWithAccessorTypeUndefined(t *testing.T) {
	c1 := datatype.NewCollectionUndefined()
	c1.Add(datatype.NewPositiveInt(10))
	c1.Add(datatype.NewPositiveInt(11))

	a2 := resource.NewDynamicResource("Patient")

	result := newCollectionWithAccessorTypes([]datatype.Accessor{c1, a2})
	assert.Equal(t, "", result.ItemTypeInfo().String())
}
