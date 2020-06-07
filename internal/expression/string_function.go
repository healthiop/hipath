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
	"github.com/volsch/gohipath/pathsys"
	"strings"
)

type indexOfFunction struct {
	pathsys.BaseFunction
}

func newIndexOfFunction() *indexOfFunction {
	return &indexOfFunction{
		BaseFunction: pathsys.NewBaseFunction("indexOf", -1, 1, 1),
	}
}

func (f *indexOfFunction) Execute(ctx pathsys.ContextAccessor, node interface{}, args []interface{}, loop pathsys.Looper) (interface{}, error) {
	s, err := stringNode(node)
	if s == nil || err != nil {
		return nil, err
	}

	var ss pathsys.StringAccessor
	var ok bool
	if ss, ok = args[0].(pathsys.StringAccessor); !ok {
		return nil, fmt.Errorf("substring is not a string: %T", args[0])
	}

	i := strings.Index(s.String(), ss.String())
	return pathsys.NewInteger(int32(i)), nil
}

func stringNode(node interface{}) (pathsys.StringAccessor, error) {
	value := unwrapCollection(node)
	if value == nil {
		return nil, nil
	}

	if s, ok := value.(pathsys.StringAccessor); !ok {
		return nil, fmt.Errorf("not a string: %T", value)
	} else {
		return s, nil
	}
}
