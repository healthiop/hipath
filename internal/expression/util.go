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

import "github.com/healthiop/hipath/hipathsys"

const delimitedIdentifierChar = '`'

func ExtractIdentifier(value string) string {
	resultingValue := value
	if len(resultingValue) > 1 && value[0] == delimitedIdentifierChar {
		resultingValue = resultingValue[1 : len(resultingValue)-1]
	}
	return resultingValue
}

func uniteCollections(ctx hipathsys.ContextAccessor, n1 interface{}, n2 interface{}) hipathsys.CollectionModifier {
	if n1 == nil && n2 == nil {
		return nil
	}

	c := ctx.NewCollection()
	addUniqueCollectionItems(c, n1)
	addUniqueCollectionItems(c, n2)

	if c.Count() == 0 {
		return nil
	}
	return c
}

func addUniqueCollectionItems(collection hipathsys.CollectionModifier, node interface{}) {
	if node == nil {
		return
	}
	if c, ok := node.(hipathsys.CollectionAccessor); ok {
		collection.AddAllUnique(c)
	} else {
		collection.AddUnique(node)
	}
}

func combineCollections(ctx hipathsys.ContextAccessor, n1 interface{}, n2 interface{}) hipathsys.CollectionModifier {
	if n1 == nil && n2 == nil {
		return nil
	}

	c := ctx.NewCollection()
	addCollectionItems(c, n1)
	addCollectionItems(c, n2)

	if c.Count() == 0 {
		return nil
	}
	return c
}

func addCollectionItems(collection hipathsys.CollectionModifier, node interface{}) {
	if node == nil {
		return
	}
	if c, ok := node.(hipathsys.CollectionAccessor); ok {
		collection.AddAll(c)
	} else {
		collection.Add(node)
	}
}

func unwrapCollection(node interface{}) interface{} {
	if node == nil {
		return nil
	}
	if c, ok := node.(hipathsys.CollectionAccessor); !ok {
		return node
	} else {
		count := c.Count()
		if count == 0 {
			return nil
		}
		if count == 1 {
			return c.Get(0)
		}
		return c
	}
}

func wrapCollection(ctx hipathsys.ContextAccessor, node interface{}) hipathsys.CollectionAccessor {
	if node == nil {
		return hipathsys.EmptyCollection
	}

	if col, ok := node.(hipathsys.CollectionAccessor); ok {
		return col
	}

	return ctx.NewCollectionWithItem(node)
}

func emptyCollection(node interface{}) bool {
	if node == nil {
		return true
	}
	if col, ok := node.(hipathsys.CollectionAccessor); ok {
		return col.Empty()
	}
	return false
}
