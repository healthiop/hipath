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

type ModelAdapter interface {
	ConvertToSystem(node interface{}) interface{}
	TypeInfo(node interface{}) TypeInfoAccessor
	Equal(node1 interface{}, node2 interface{}) bool
	Equivalent(node1 interface{}, node2 interface{}) bool
	Navigate(node interface{}, name string) (interface{}, error)
}

func ModelTypeInfo(adapter ModelAdapter, node interface{}) TypeInfoAccessor {
	if node == nil {
		return nil
	}

	if n, ok := node.(AnyAccessor); ok {
		return n.TypeInfo()
	}
	return adapter.TypeInfo(node)
}

func ModelEqual(adapter ModelAdapter, node1 interface{}, node2 interface{}) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}

	sysNode1, _ := node1.(AnyAccessor)
	sysNode2, _ := node2.(AnyAccessor)
	if sysNode1 != nil && sysNode2 != nil {
		return sysNode1.Equal(sysNode2)
	}
	if sysNode1 != nil || sysNode2 != nil {
		return false
	}

	return adapter.Equal(node1, node2)
}

func ModelEquivalent(adapter ModelAdapter, node1 interface{}, node2 interface{}) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}

	sysNode1, _ := node1.(AnyAccessor)
	sysNode2, _ := node2.(AnyAccessor)
	if sysNode1 != nil && sysNode2 != nil {
		return sysNode1.Equivalent(sysNode2)
	}
	if sysNode1 != nil || sysNode2 != nil {
		return false
	}

	return adapter.Equivalent(node1, node2)
}

func SystemAnyTypeEqual(node1 AnyAccessor, node2 interface{}) bool {
	if node1 == nil || node2 == nil {
		return false
	}
	if sysNode2, ok := node2.(AnyAccessor); ok {
		return node1.TypeInfo().Equal(sysNode2.TypeInfo())
	}
	return false
}

func SystemAnyEqual(node1 AnyAccessor, node2 interface{}) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}

	sysNode2, _ := node2.(AnyAccessor)
	if sysNode2 == nil {
		return false
	}

	return Equal(node1, sysNode2)
}

type ContextAccessor interface {
	EnvVar(name string) (interface{}, bool)
	ContextNode() interface{}
	ModelAdapter() ModelAdapter
	NewCollection() CollectionModifier
}
