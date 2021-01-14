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
	"time"
)

type traceFunction struct {
	hipathsys.BaseFunction
}

func newTraceFunction() *traceFunction {
	return &traceFunction{
		BaseFunction: hipathsys.NewBaseFunction("trace", 1, 1, 2),
	}
}

func (f *traceFunction) Execute(ctx hipathsys.ContextAccessor, node interface{}, args []interface{}, loop hipathsys.Looper) (interface{}, error) {
	tracer := ctx.Tracer()
	if tracer == nil {
		return node, nil
	}

	name, err := stringNode(args[0])
	if name == nil || err != nil {
		return node, err
	}

	if !tracer.Enabled(name.String()) {
		return node, nil
	}

	col := wrapCollection(ctx, node)
	count := col.Count()

	var traced hipathsys.ColAccessor
	if count == 0 {
		traced = hipathsys.EmptyCol
	} else if loopEvaluator := loop.Evaluator(); loopEvaluator != nil {
		projected := ctx.NewCol()
		for i := 0; i < count; i++ {
			this := col.Get(i)
			loop.IncIndex(this)

			res, err := loopEvaluator.Evaluate(ctx, this, loop)
			if err != nil {
				return nil, err
			}
			projected.Add(res)
		}
		traced = projected
	} else {
		traced = col
	}

	tracer.Trace(name.String(), traced)
	return node, nil
}

type nowFunction struct {
	hipathsys.BaseFunction
}

func newNowFunction() *nowFunction {
	return &nowFunction{
		BaseFunction: hipathsys.NewBaseFunction("now", -1, 0, 0),
	}
}

func (f *nowFunction) Execute(_ hipathsys.ContextAccessor, _ interface{}, _ []interface{}, _ hipathsys.Looper) (interface{}, error) {
	return hipathsys.NewDateTime(time.Now()), nil
}

type timeOfDayFunction struct {
	hipathsys.BaseFunction
}

func newTimeOfDayFunction() *timeOfDayFunction {
	return &timeOfDayFunction{
		BaseFunction: hipathsys.NewBaseFunction("timeOfDay", -1, 0, 0),
	}
}

func (f *timeOfDayFunction) Execute(_ hipathsys.ContextAccessor, _ interface{}, _ []interface{}, _ hipathsys.Looper) (interface{}, error) {
	return hipathsys.NewTime(time.Now()), nil
}

type todayFunction struct {
	hipathsys.BaseFunction
}

func newTodayFunction() *todayFunction {
	return &todayFunction{
		BaseFunction: hipathsys.NewBaseFunction("today", -1, 0, 0),
	}
}

func (f *todayFunction) Execute(_ hipathsys.ContextAccessor, _ interface{}, _ []interface{}, _ hipathsys.Looper) (interface{}, error) {
	return hipathsys.NewDate(time.Now()), nil
}
