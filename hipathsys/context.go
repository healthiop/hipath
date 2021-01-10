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

type ModelAdapter interface {
	ConvertToSystem(node interface{}) (interface{}, error)
	TypeSpec(node interface{}) TypeSpecAccessor
	Cast(node interface{}, name FQTypeNameAccessor) (interface{}, error)
	Equal(node1 interface{}, node2 interface{}) bool
	Equivalent(node1 interface{}, node2 interface{}) bool
	Navigate(node interface{}, name string) (interface{}, error)
	Children(node interface{}) (CollectionAccessor, error)
}

func ModelTypeSpec(adapter ModelAdapter, node interface{}) TypeSpecAccessor {
	if node == nil {
		return nil
	}

	if n, ok := node.(AnyAccessor); ok {
		return n.TypeSpec()
	}
	return adapter.TypeSpec(node)
}

func HasModelType(adapter ModelAdapter, node interface{}, name FQTypeNameAccessor) bool {
	if node == nil {
		return false
	}

	namespace := name.Namespace()
	if systemNamespace(namespace) {
		if n, ok := node.(AnyAccessor); ok {
			if n.TypeSpec().ExtendsName(name) {
				return true
			}
		}
	}

	if namespace != NamespaceName {
		if adapter.TypeSpec(node).ExtendsName(name) {
			return true
		}
	}

	return false
}

func CastModelType(adapter ModelAdapter, node interface{}, name FQTypeNameAccessor) (interface{}, error) {
	if node == nil {
		return node, nil
	}

	sysNode, sys := node.(AnyAccessor)
	if systemNamespace(name.Namespace()) && sys && sysNode.TypeSpec().ExtendsName(name) {
		return node, nil
	}
	if sys {
		// system node cannot be casted by model adapter
		return nil, nil
	}

	return adapter.Cast(node, name)
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
		return node1.TypeSpec().EqualType(sysNode2.TypeSpec())
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

type Tracer interface {
	Enabled(name string) bool
	Trace(name string, col CollectionAccessor)
}

type ContextAccessor interface {
	EnvVar(name string) (interface{}, bool)
	ContextNode() interface{}
	ModelAdapter() ModelAdapter
	NewCollection() CollectionModifier
	NewCollectionWithItem(item interface{}) (CollectionModifier, error)
	Tracer() Tracer
}

func systemNamespace(name string) bool {
	return len(name) == 0 || name == NamespaceName
}
