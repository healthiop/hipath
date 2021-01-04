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

type loop struct {
	evaluator Evaluator
	this      interface{}
	index     int
	total     interface{}
}

type Looper interface {
	Evaluator() Evaluator
	This() interface{}
	Index() int
	IncIndex(this interface{}) int
	Total() interface{}
	SetTotal(total interface{})
}

func NewLoop(evaluator Evaluator) Looper {
	return &loop{evaluator, nil, -1, nil}
}

func NewLoopWithIndex(evaluator Evaluator, index int) Looper {
	return &loop{evaluator, nil, index - 1, nil}
}

func (c *loop) Evaluator() Evaluator {
	return c.evaluator
}

func (c *loop) This() interface{} {
	if c.index < 0 {
		panic("index has not yet been incremented")
	}
	return c.this
}

func (c *loop) Index() int {
	return c.index
}

func (c *loop) IncIndex(this interface{}) int {
	c.this = this
	c.index = c.index + 1
	return c.index
}

func (c *loop) Total() interface{} {
	return c.total
}

func (c *loop) SetTotal(total interface{}) {
	c.total = total
}

type Evaluator interface {
	Evaluate(ctx ContextAccessor, node interface{}, loop Looper) (interface{}, error)
}

type BaseFunction struct {
	name           string
	evaluatorParam int
	minParams      int
	maxParams      int
}

type FunctionExecutor interface {
	Name() string
	EvaluatorParam() int
	MinParams() int
	MaxParams() int
	Execute(ctx ContextAccessor, node interface{}, args []interface{}, loop Looper) (interface{}, error)
}

func NewBaseFunction(name string, evaluatorParam int, minParams int, maxParams int) BaseFunction {
	return BaseFunction{name, evaluatorParam, minParams, maxParams}
}

func (f *BaseFunction) Name() string {
	return f.name
}

func (f *BaseFunction) EvaluatorParam() int {
	return f.evaluatorParam
}

func (f *BaseFunction) MinParams() int {
	return f.minParams
}

func (f *BaseFunction) MaxParams() int {
	return f.maxParams
}
