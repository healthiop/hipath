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

package pathsys

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTypeName(t *testing.T) {
	n := NewTypeName("test")
	assert.Equal(t, "", n.Namespace())
	assert.Equal(t, "test", n.Name())
	assert.Equal(t, "test", n.String())
}

func TestTypeInfoStringNil(t *testing.T) {
	o := NewTypeInfoWithBase(nil, NewTypeInfo(NewFQTypeName("test", "test")))
	assert.Equal(t, "", o.String())
}

func TestNewFQTypeNameWithoutNamespace(t *testing.T) {
	n := NewFQTypeName("test", "")
	assert.Equal(t, "", n.Namespace())
	assert.Equal(t, "test", n.Name())
	assert.Equal(t, "test", n.String())
}

func TestFQTypeNameEqual(t *testing.T) {
	n1 := NewFQTypeName("test1", "ns1")
	n2 := NewFQTypeName("test1", "ns1")
	assert.Equal(t, true, FQTypeNameEqual(n1, n2))
}

func TestFQTypeNameEqualNot(t *testing.T) {
	n1 := NewFQTypeName("test1", "ns1")
	n2 := NewFQTypeName("test2", "ns1")
	assert.Equal(t, false, FQTypeNameEqual(n1, n2))
}

func TestFQTypeNameEqualNil(t *testing.T) {
	assert.Equal(t, true, FQTypeNameEqual(nil, nil))
}

func TestFQTypeNameEqualLeftNil(t *testing.T) {
	assert.Equal(t, false, FQTypeNameEqual(nil, NewFQTypeName("test1", "ns1")))
}

func TestFQTypeNameEqualRightNil(t *testing.T) {
	assert.Equal(t, false, FQTypeNameEqual(NewFQTypeName("test1", "ns1"), nil))
}

func TestTypeInfoEqual(t *testing.T) {
	ti1 := NewTypeInfoWithBase(
		NewFQTypeName("test1", "ns1"),
		NewTypeInfo(NewFQTypeName("test2", "ns2")))
	ti2 := NewTypeInfoWithBase(
		NewFQTypeName("test1", "ns1"),
		NewTypeInfo(NewFQTypeName("test2", "ns2")))
	assert.Equal(t, true, ti1.Equal(ti2))
}

func TestTypeInfoEqualDiffers(t *testing.T) {
	ti1 := NewTypeInfoWithBase(
		NewFQTypeName("test1", "ns3"),
		NewTypeInfo(NewFQTypeName("test2", "ns2")))
	ti2 := NewTypeInfoWithBase(
		NewFQTypeName("test1", "ns1"),
		NewTypeInfo(NewFQTypeName("test2", "ns2")))
	assert.Equal(t, false, ti1.Equal(ti2))
}

func TestNewTypeInfo(t *testing.T) {
	ti := NewTypeInfo(NewFQTypeName("test1", "ns3"))
	assert.Equal(t, "ns3.test1", ti.String())
	assert.Nil(t, ti.Base(), "no base expected")
	assert.Nil(t, ti.FQBaseName(), "no base name expected")
}

func TestNewTypeInfoWithBase(t *testing.T) {
	ti := NewTypeInfoWithBase(
		NewFQTypeName("test1", "ns3"),
		NewTypeInfo(NewFQTypeName("test2", "ns2")))
	assert.Equal(t, "ns3.test1", ti.String())
	if assert.NotNil(t, ti.Base(), "base expected") {
		assert.Equal(t, "ns2.test2", ti.Base().String())
		assert.Equal(t, NewFQTypeName("test2", "ns2"), ti.FQBaseName())
	}
}

func TestCommonBaseTypeSame(t *testing.T) {
	resource := NewTypeInfo(NewFQTypeName("Resource", "FHIR"))
	domainResource := NewTypeInfoWithBase(NewFQTypeName("DomainResource", "FHIR"), resource)
	patient := NewTypeInfoWithBase(NewFQTypeName("Patient", "FHIR"), domainResource)
	assert.Same(t, patient, CommonBaseType(patient, patient))
}

func TestCommonBaseTypeMiddle(t *testing.T) {
	resource := NewTypeInfo(NewFQTypeName("Resource", "FHIR"))
	domainResource := NewTypeInfoWithBase(NewFQTypeName("DomainResource", "FHIR"), resource)
	patient := NewTypeInfoWithBase(NewFQTypeName("Patient", "FHIR"), domainResource)
	person := NewTypeInfoWithBase(NewFQTypeName("Person", "FHIR"), domainResource)
	assert.Same(t, domainResource, CommonBaseType(patient, person))
}

func TestCommonBaseTypeRoot(t *testing.T) {
	resource := NewTypeInfo(NewFQTypeName("Resource", "FHIR"))
	domainResource := NewTypeInfoWithBase(NewFQTypeName("DomainResource", "FHIR"), resource)
	patient := NewTypeInfoWithBase(NewFQTypeName("Patient", "FHIR"), domainResource)
	medication := NewTypeInfoWithBase(NewFQTypeName("Medication", "FHIR"), resource)
	assert.Same(t, resource, CommonBaseType(patient, medication))
}

func TestCommonBaseTypeNone(t *testing.T) {
	resource := NewTypeInfo(NewFQTypeName("Resource", "FHIR"))
	domainResource := NewTypeInfoWithBase(NewFQTypeName("DomainResource", "FHIR"), resource)
	patient := NewTypeInfoWithBase(NewFQTypeName("Patient", "FHIR"), domainResource)
	other := NewTypeInfo(NewFQTypeName("Other", ""))
	assert.Nil(t, CommonBaseType(patient, other), "no common base type expected")
}
