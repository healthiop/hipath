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
	"fmt"
	"github.com/volsch/gohimodel/datatype"
)

type IndexerExpression struct {
	exprEvaluator  Evaluator
	indexEvaluator Evaluator
}

func NewIndexerExpression(exprEvaluator Evaluator, indexEvaluator Evaluator) *IndexerExpression {
	return &IndexerExpression{exprEvaluator, indexEvaluator}
}

func (e *IndexerExpression) Evaluate(ctx *EvalContext, obj datatype.Accessor) (datatype.Accessor, error) {
	accessor, err := e.exprEvaluator.Evaluate(ctx, obj)
	if err != nil {
		return nil, err
	}
	index, err := e.indexEvaluator.Evaluate(ctx, obj)
	if err != nil {
		return nil, err
	}

	if accessor == nil || index == nil {
		return nil, nil
	}

	var indexValue int
	if n, ok := index.(datatype.NumberAccessor); !ok {
		return nil, fmt.Errorf("index is not a number: %s", index.TypeInfo().String())
	} else {
		indexValue = int(n.Int())
	}

	if indexValue < 0 {
		return nil, nil
	}

	if c, ok := accessor.(datatype.CollectionAccessor); ok {
		if indexValue >= c.Count() {
			return nil, nil
		}
		return c.Get(indexValue), nil
	}

	if indexValue > 0 {
		return nil, nil
	}
	return accessor, nil
}
