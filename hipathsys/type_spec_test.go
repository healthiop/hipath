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

package hipathsys

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTypeName(t *testing.T) {
	n := NewTypeName("test")
	assert.Equal(t, "", n.Namespace())
	assert.False(t, n.HasNamespace())
	assert.Equal(t, "test", n.Name())
	assert.Equal(t, "test", n.String())
	assert.False(t, n.Anonymous())
}

func TestNewTypeNameEmpty(t *testing.T) {
	n := NewTypeName("")
	assert.Equal(t, "", n.Namespace())
	assert.False(t, n.HasNamespace())
	assert.Equal(t, "", n.Name())
	assert.Equal(t, "", n.String())
	assert.True(t, n.Anonymous())
}

func TestNewTypeNameWithNamespace(t *testing.T) {
	n := NewFQTypeName("test", "xyz")
	assert.Equal(t, "xyz", n.Namespace())
	assert.True(t, n.HasNamespace())
	assert.Equal(t, "test", n.Name())
	assert.Equal(t, "xyz.test", n.String())
	assert.False(t, n.Anonymous())
}

func TestNewTypeNameWithNamespaceAnonymous(t *testing.T) {
	n := NewFQTypeName("", "xyz")
	assert.Equal(t, "xyz", n.Namespace())
	assert.True(t, n.HasNamespace())
	assert.Equal(t, "", n.Name())
	assert.Equal(t, "", n.String())
	assert.True(t, n.Anonymous())
}

func TestTypeSpecStringNil(t *testing.T) {
	o := NewTypeSpecWithBase(nil, NewTypeSpec(NewFQTypeName("test", "test")))
	assert.Equal(t, "", o.String())
	assert.True(t, o.Anonymous())
}

func TestNewFQTypeNameWithoutNamespace(t *testing.T) {
	n := NewFQTypeName("test", "")
	assert.Equal(t, "", n.Namespace())
	assert.False(t, n.HasNamespace())
	assert.Equal(t, "test", n.Name())
	assert.Equal(t, "test", n.String())
	assert.False(t, n.Anonymous())
}

func TestFQTypeNameEqual(t *testing.T) {
	n1 := NewFQTypeName("test1", "ns1")
	n2 := NewFQTypeName("test1", "ns1")
	assert.Equal(t, true, FQTypeNameEqual(n1, n2))
}

func TestFQTypeNameEqualAnonymousEmpty(t *testing.T) {
	n1 := NewFQTypeName("", "ns1")
	n2 := NewFQTypeName("", "ns1")
	assert.Equal(t, false, FQTypeNameEqual(n1, n2))
}

func TestFQTypeNameEqualNot(t *testing.T) {
	n1 := NewFQTypeName("test1", "ns1")
	n2 := NewFQTypeName("test2", "ns1")
	assert.Equal(t, false, FQTypeNameEqual(n1, n2))
}

func TestFQTypeNameEqualAnonymousNil(t *testing.T) {
	assert.Equal(t, false, FQTypeNameEqual(nil, nil))
}

func TestFQTypeNameEqualLeftNil(t *testing.T) {
	assert.Equal(t, false, FQTypeNameEqual(nil, NewFQTypeName("test1", "ns1")))
}

func TestFQTypeNameEqualBothAnonymous(t *testing.T) {
	assert.Equal(t, false, FQTypeNameEqual(NewFQTypeName("", "ns1"), NewFQTypeName("", "ns1")))
}

func TestFQTypeNameEqualRightNil(t *testing.T) {
	assert.Equal(t, false, FQTypeNameEqual(NewFQTypeName("test1", "ns1"), nil))
}

func TestTypeSpecEqual(t *testing.T) {
	ti1 := NewTypeSpecWithBase(
		NewFQTypeName("test1", "ns1"),
		NewTypeSpec(NewFQTypeName("test2", "ns2")))
	ti2 := NewTypeSpecWithBase(
		NewFQTypeName("test1", "ns1"),
		NewTypeSpec(NewFQTypeName("test2", "ns2")))
	assert.Equal(t, true, ti1.EqualType(ti2))
}

func TestTypeSpecEqualDiffers(t *testing.T) {
	ti1 := NewTypeSpecWithBase(
		NewFQTypeName("test1", "ns3"),
		NewTypeSpec(NewFQTypeName("test2", "ns2")))
	ti2 := NewTypeSpecWithBase(
		NewFQTypeName("test1", "ns1"),
		NewTypeSpec(NewFQTypeName("test2", "ns2")))
	assert.Equal(t, false, ti1.EqualType(ti2))
}

func TestNewTypeSpec(t *testing.T) {
	ti := NewTypeSpec(NewFQTypeName("test1", "ns3"))
	assert.Equal(t, "ns3.test1", ti.String())
	assert.Nil(t, ti.Base(), "no base expected")
	assert.Nil(t, ti.FQBaseName(), "no base name expected")
	assert.False(t, ti.Anonymous())
}

func TestNewTypeSpecAnonymousNil(t *testing.T) {
	ti := NewTypeSpec(nil)
	assert.Equal(t, "", ti.String())
	assert.Nil(t, ti.Base(), "no base expected")
	assert.Nil(t, ti.FQBaseName(), "no base name expected")
	assert.True(t, ti.Anonymous())
}

func TestNewTypeSpecAnonymousEmpty(t *testing.T) {
	ti := NewTypeSpec(NewFQTypeName("", "ns3"))
	assert.Equal(t, "", ti.String())
	assert.Nil(t, ti.Base(), "no base expected")
	assert.Nil(t, ti.FQBaseName(), "no base name expected")
	assert.True(t, ti.Anonymous())
}

func TestNewTypeSpecWithBase(t *testing.T) {
	ti := NewTypeSpecWithBase(
		NewFQTypeName("test1", "ns3"),
		NewTypeSpec(NewFQTypeName("test2", "ns2")))
	assert.False(t, ti.Anonymous())
	assert.Equal(t, "ns3.test1", ti.String())
	if assert.NotNil(t, ti.Base(), "base expected") {
		assert.Equal(t, "ns2.test2", ti.Base().String())
		assert.Equal(t, NewFQTypeName("test2", "ns2"), ti.FQBaseName())
	}
}

func TestCommonBaseTypeSame(t *testing.T) {
	resource := NewTypeSpec(NewFQTypeName("Resource", "FHIR"))
	domainResource := NewTypeSpecWithBase(NewFQTypeName("DomainResource", "FHIR"), resource)
	patient := NewTypeSpecWithBase(NewFQTypeName("Patient", "FHIR"), domainResource)
	assert.Same(t, patient, CommonBaseType(patient, patient))
}

func TestCommonBaseTypeMiddle(t *testing.T) {
	resource := NewTypeSpec(NewFQTypeName("Resource", "FHIR"))
	domainResource := NewTypeSpecWithBase(NewFQTypeName("DomainResource", "FHIR"), resource)
	patient := NewTypeSpecWithBase(NewFQTypeName("Patient", "FHIR"), domainResource)
	person := NewTypeSpecWithBase(NewFQTypeName("Person", "FHIR"), domainResource)
	assert.Same(t, domainResource, CommonBaseType(patient, person))
}

func TestCommonBaseTypeRoot(t *testing.T) {
	resource := NewTypeSpec(NewFQTypeName("Resource", "FHIR"))
	domainResource := NewTypeSpecWithBase(NewFQTypeName("DomainResource", "FHIR"), resource)
	patient := NewTypeSpecWithBase(NewFQTypeName("Patient", "FHIR"), domainResource)
	medication := NewTypeSpecWithBase(NewFQTypeName("Medication", "FHIR"), resource)
	assert.Same(t, resource, CommonBaseType(patient, medication))
}

func TestCommonBaseTypeNone(t *testing.T) {
	resource := NewTypeSpec(NewFQTypeName("Resource", "FHIR"))
	domainResource := NewTypeSpecWithBase(NewFQTypeName("DomainResource", "FHIR"), resource)
	patient := NewTypeSpecWithBase(NewFQTypeName("Patient", "FHIR"), domainResource)
	other := NewTypeSpec(NewFQTypeName("Other", ""))
	assert.Nil(t, CommonBaseType(patient, other), "no common base type expected")
}

func TestTypeSpecExtends(t *testing.T) {
	resource := NewTypeSpec(NewFQTypeName("Resource", "FHIR"))
	assert.True(t, resource.ExtendsName(NewFQTypeName("Resource", "FHIR")))
}

func TestTypeSpecExtendsNil(t *testing.T) {
	resource := NewTypeSpec(nil)
	assert.False(t, resource.ExtendsName(NewFQTypeName("Resource", "FHIR")))
}

func TestTypeSpecExtendsBase(t *testing.T) {
	resource := NewTypeSpec(NewFQTypeName("Resource", "FHIR"))
	domainResource := NewTypeSpecWithBase(NewFQTypeName("DomainResource", "FHIR"), resource)
	assert.True(t, domainResource.ExtendsName(NewFQTypeName("Resource", "FHIR")))
}

func TestTypeSpecExtendsNot(t *testing.T) {
	resource := NewTypeSpec(NewFQTypeName("Resource", "FHIR"))
	domainResource := NewTypeSpecWithBase(NewFQTypeName("DomainResource", "FHIR"), resource)
	assert.False(t, domainResource.ExtendsName(NewFQTypeName("Patient", "FHIR")))
}

func TestTypeSpecExtendsBaseWithoutNamespace(t *testing.T) {
	resource := NewTypeSpec(NewFQTypeName("Resource", "FHIR"))
	domainResource := NewTypeSpecWithBase(NewFQTypeName("DomainResource", "FHIR"), resource)
	assert.True(t, domainResource.ExtendsName(NewTypeName("Resource")))
}

func TestTypeSpecExtendsNotWithoutNamespace(t *testing.T) {
	resource := NewTypeSpec(NewFQTypeName("Resource", "FHIR"))
	domainResource := NewTypeSpecWithBase(NewFQTypeName("DomainResource", "FHIR"), resource)
	assert.False(t, domainResource.ExtendsName(NewTypeName("Patient")))
}

func TestParseFQTypeName(t *testing.T) {
	tn, err := ParseFQTypeName("TEST.data")
	assert.NoError(t, err, "no error expected")
	if assert.NotNil(t, tn, "full qualified type name expected") {
		assert.Equal(t, "TEST", tn.Namespace())
		assert.Equal(t, "data", tn.Name())
	}
}

func TestParseFQTypeNameEmpty(t *testing.T) {
	tn, err := ParseFQTypeName("")
	assert.Error(t, err, "error expected")
	assert.Nil(t, tn, "no result expected")
}

func TestParseFQTypeNameWithoutNamespace(t *testing.T) {
	tn, err := ParseFQTypeName("data")
	assert.NoError(t, err, "no error expected")
	if assert.NotNil(t, tn, "full qualified type name expected") {
		assert.False(t, tn.HasNamespace())
		assert.Equal(t, "data", tn.Name())
	}
}

func TestParseFQTypeNameThreeComponents(t *testing.T) {
	tn, err := ParseFQTypeName("TEST.data.other")
	assert.Error(t, err, "error expected")
	assert.Nil(t, tn, "no result expected")
}

func TestParseFQTypeNameNoNamespace(t *testing.T) {
	tn, err := ParseFQTypeName(".data")
	assert.Error(t, err, "error expected")
	assert.Nil(t, tn, "no result expected")
}

func TestParseFQTypeNameNoName(t *testing.T) {
	tn, err := ParseFQTypeName("TEST.")
	assert.Error(t, err, "error expected")
	assert.Nil(t, tn, "no result expected")
}
