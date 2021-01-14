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
	"github.com/healthiop/hipath/internal/test"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTraceFuncNoTracer(t *testing.T) {
	ctx := test.NewTestContext(t)
	node := hipathsys.NewString("value")

	f := newTraceFunction()
	res, err := f.Execute(ctx, node, []interface{}{hipathsys.NewString("test-tracer")},
		hipathsys.NewLoop(newTestErrorExpression()))

	assert.NoError(t, err, "no error expected")
	assert.Same(t, node, res, "unchanged node expected")
}

func TestTraceFuncNameNil(t *testing.T) {
	tracer := newTestingTracer()
	ctx := test.NewTestContextWithNodeAndTracer(t, hipathsys.NewString("test"), tracer)
	node := hipathsys.NewString("value")

	f := newTraceFunction()
	res, err := f.Execute(ctx, node, []interface{}{nil},
		hipathsys.NewLoop(newTestErrorExpression()))

	assert.NoError(t, err, "no error expected")
	assert.Same(t, res, node, "unchanged node expected")
	assert.Equal(t, 0, tracer.count)
}

func TestTraceFuncEmptyInput(t *testing.T) {
	tracer := newTestingTracer()
	ctx := test.NewTestContextWithNodeAndTracer(t, hipathsys.NewString("test"), tracer)

	f := newTraceFunction()
	res, err := f.Execute(ctx, nil, []interface{}{hipathsys.NewString("test-tracer")},
		hipathsys.NewLoop(newTestErrorExpression()))

	assert.NoError(t, err, "no error expected")
	assert.Nil(t, res, "unchanged node expected")
	assert.Equal(t, 1, tracer.count)
	assert.Equal(t, "test-tracer", tracer.name)
	if assert.NotNil(t, tracer.col, "traced collection expected") {
		assert.Equal(t, 0, tracer.col.Count(), "empty collection expected")
	}
}

func TestTraceFunc(t *testing.T) {
	tracer := newTestingTracer()
	ctx := test.NewTestContextWithNodeAndTracer(t, hipathsys.NewString("test"), tracer)

	node := ctx.NewCol()
	node.Add(hipathsys.NewString("value1"))
	node.Add(hipathsys.NewString("value2"))

	f := newTraceFunction()
	res, err := f.Execute(ctx, node, []interface{}{hipathsys.NewString("test-tracer")},
		hipathsys.NewLoop(nil))

	assert.NoError(t, err, "no error expected")
	assert.Same(t, node, res, "unchanged node expected")
	assert.Equal(t, 1, tracer.count)
	assert.Equal(t, "test-tracer", tracer.name)
	assert.Same(t, node, tracer.col, "traced collection expected")
}

func TestTraceFuncDisabled(t *testing.T) {
	tracer := newTestingTracer()
	ctx := test.NewTestContextWithNodeAndTracer(t, hipathsys.NewString("test"), tracer)

	node := ctx.NewCol()
	node.Add(hipathsys.NewString("value1"))
	node.Add(hipathsys.NewString("value2"))

	f := newTraceFunction()
	res, err := f.Execute(ctx, node, []interface{}{hipathsys.NewString("other-tracer")},
		hipathsys.NewLoop(nil))

	assert.NoError(t, err, "no error expected")
	assert.Same(t, node, res, "unchanged node expected")
	assert.Equal(t, 0, tracer.count)
}

func TestTraceFuncProjection(t *testing.T) {
	tracer := newTestingTracer()
	ctx := test.NewTestContextWithNodeAndTracer(t, hipathsys.NewString("test"), tracer)

	nodeN := make(map[string]interface{})
	nodeN["id"] = nil
	nodeN["item"] = "testN"

	id1 := ctx.NewCol()
	id1.Add(hipathsys.NewString("1"))
	node1 := make(map[string]interface{})
	node1["id"] = id1
	node1["item"] = "test1"

	node7 := make(map[string]interface{})
	node7["id"] = hipathsys.NewString("7")
	node7["item"] = "test7"

	node9 := make(map[string]interface{})
	node9["id"] = hipathsys.NewString("9")
	node9["item"] = "test9"

	node := ctx.NewCol()
	node.Add(node1)
	node.Add(node9)
	node.Add(nodeN)
	node.Add(node7)

	f := newTraceFunction()
	res, err := f.Execute(ctx, node, []interface{}{hipathsys.NewString("test-tracer")},
		hipathsys.NewLoop(NewMemberInvocation("id")))

	assert.NoError(t, err, "no error expected")
	assert.Same(t, node, res, "unchanged node expected")
	assert.Equal(t, 1, tracer.count)
	assert.Equal(t, "test-tracer", tracer.name)
	if assert.NotNil(t, tracer.col, "traced collection expected") {
		if assert.Equal(t, 4, tracer.col.Count()) {
			assert.Same(t, id1, tracer.col.Get(0))
			assert.Equal(t, hipathsys.NewString("9"), tracer.col.Get(1))
			assert.Nil(t, tracer.col.Get(2))
			assert.Equal(t, hipathsys.NewString("7"), tracer.col.Get(3))
		}
	}
}

func TestTraceFuncProjectionError(t *testing.T) {
	tracer := newTestingTracer()
	ctx := test.NewTestContextWithNodeAndTracer(t, hipathsys.NewString("test"), tracer)

	node := ctx.NewCol()
	node.Add(hipathsys.NewString("value1"))
	node.Add(hipathsys.NewString("value2"))

	f := newTraceFunction()
	res, err := f.Execute(ctx, node, []interface{}{hipathsys.NewString("test-tracer")},
		hipathsys.NewLoop(newTestErrorExpression()))

	assert.Error(t, err, "error expected")
	assert.Nil(t, res, "no result expected")
	assert.Equal(t, 0, tracer.count)
}

func TestNowFunc(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newNowFunction()
	b := time.Now()
	res, err := f.Execute(ctx, nil, []interface{}{}, nil)
	e := time.Now()

	assert.NoError(t, err, "no error expected")
	if assert.Implements(t, (*hipathsys.DateTimeAccessor)(nil), res) {
		assert.False(t, res.(hipathsys.DateTimeAccessor).Time().Before(b))
		assert.False(t, res.(hipathsys.DateTimeAccessor).Time().After(e))
	}
}

func TestTimeOfDayFunc(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newTimeOfDayFunction()
	res, err := f.Execute(ctx, nil, []interface{}{}, nil)

	assert.NoError(t, err, "no error expected")
	assert.Implements(t, (*hipathsys.TimeAccessor)(nil), res)
}

func TestTodayFunc(t *testing.T) {
	ctx := test.NewTestContext(t)

	f := newTodayFunction()
	res, err := f.Execute(ctx, nil, []interface{}{}, nil)

	assert.NoError(t, err, "no error expected")
	assert.Implements(t, (*hipathsys.DateAccessor)(nil), res)
}

type testingTracer struct {
	count int
	name  string
	col   hipathsys.ColAccessor
}

func newTestingTracer() *testingTracer {
	return &testingTracer{}
}

func (t *testingTracer) Enabled(name string) bool {
	return name == "test-tracer"
}

func (t *testingTracer) Trace(name string, col hipathsys.ColAccessor) {
	t.count++
	t.name = name
	t.col = col
}
