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
	"fmt"
	"github.com/volsch/gohipath/pathsys"
	"regexp"
	"strings"
	"unicode/utf8"
)

type indexOfFunction struct {
	pathsys.BaseFunction
}

func newIndexOfFunction() *indexOfFunction {
	return &indexOfFunction{
		BaseFunction: pathsys.NewBaseFunction("indexOf", -1, 1, 1),
	}
}

func (f *indexOfFunction) Execute(_ pathsys.ContextAccessor, node interface{}, args []interface{}, _ pathsys.Looper) (interface{}, error) {
	s, err := stringNode(node)
	if s == nil || err != nil {
		return nil, err
	}

	ss, err := stringNode(args[0])
	if ss == nil || err != nil {
		return nil, err
	}

	i := strings.Index(s.String(), ss.String())
	return pathsys.NewInteger(int32(i)), nil
}

type substringFunction struct {
	pathsys.BaseFunction
}

func newSubstringFunction() *substringFunction {
	return &substringFunction{
		BaseFunction: pathsys.NewBaseFunction("substring", -1, 1, 2),
	}
}

func (f *substringFunction) Execute(_ pathsys.ContextAccessor, node interface{}, args []interface{}, _ pathsys.Looper) (interface{}, error) {
	s, err := stringNode(node)
	if s == nil || err != nil {
		return nil, err
	}

	start, err := integerNode(args[0])
	if start == nil || err != nil {
		return nil, err
	}
	startVal := start.Int()
	if startVal < 0 {
		startVal = 0
	}

	var l pathsys.IntegerAccessor = nil
	if len(args) > 1 {
		l, err = integerNode(args[1])
		if err != nil {
			return nil, err
		}
	}

	var lVal int32
	if l != nil {
		lVal = l.Int()
		if lVal <= 0 {
			return nil, nil
		}
	}

	sr := []rune(s.String())
	srLen := int32(len(sr))
	if l == nil {
		lVal = srLen - startVal
	}

	if startVal >= srLen || lVal <= 0 {
		return nil, nil
	}
	if startVal+lVal > srLen {
		lVal = srLen - startVal
	}

	return pathsys.StringOf(string(sr[startVal : startVal+lVal])), nil
}

type startsWithFunction struct {
	pathsys.BaseFunction
}

func newStartsWithFunction() *startsWithFunction {
	return &startsWithFunction{
		BaseFunction: pathsys.NewBaseFunction("startsWith", -1, 1, 1),
	}
}

func (f *startsWithFunction) Execute(_ pathsys.ContextAccessor, node interface{}, args []interface{}, _ pathsys.Looper) (interface{}, error) {
	s, err := stringNode(node)
	if s == nil || err != nil {
		return nil, err
	}

	ss, err := stringNode(args[0])
	if ss == nil || err != nil {
		return nil, err
	}

	return pathsys.BooleanOf(strings.HasPrefix(s.String(), ss.String())), nil
}

type endsWithFunction struct {
	pathsys.BaseFunction
}

func newEndsWithFunction() *endsWithFunction {
	return &endsWithFunction{
		BaseFunction: pathsys.NewBaseFunction("endsWith", -1, 1, 1),
	}
}

func (f *endsWithFunction) Execute(_ pathsys.ContextAccessor, node interface{}, args []interface{}, _ pathsys.Looper) (interface{}, error) {
	s, err := stringNode(node)
	if s == nil || err != nil {
		return nil, err
	}

	ss, err := stringNode(args[0])
	if ss == nil || err != nil {
		return nil, err
	}

	return pathsys.BooleanOf(strings.HasSuffix(s.String(), ss.String())), nil
}

type containsFunction struct {
	pathsys.BaseFunction
}

func newContainsFunction() *containsFunction {
	return &containsFunction{
		BaseFunction: pathsys.NewBaseFunction("contains", -1, 1, 1),
	}
}

func (f *containsFunction) Execute(_ pathsys.ContextAccessor, node interface{}, args []interface{}, _ pathsys.Looper) (interface{}, error) {
	s, err := stringNode(node)
	if s == nil || err != nil {
		return nil, err
	}

	ss, err := stringNode(args[0])
	if ss == nil || err != nil {
		return nil, err
	}

	return pathsys.BooleanOf(strings.Contains(s.String(), ss.String())), nil
}

type upperFunction struct {
	pathsys.BaseFunction
}

func newUpperFunction() *upperFunction {
	return &upperFunction{
		BaseFunction: pathsys.NewBaseFunction("upper", -1, 0, 0),
	}
}

func (f *upperFunction) Execute(_ pathsys.ContextAccessor, node interface{}, _ []interface{}, _ pathsys.Looper) (interface{}, error) {
	s, err := stringNode(node)
	if s == nil || err != nil {
		return nil, err
	}

	return pathsys.NewString(strings.ToUpper(s.String())), nil
}

type lowerFunction struct {
	pathsys.BaseFunction
}

func newLowerFunction() *lowerFunction {
	return &lowerFunction{
		BaseFunction: pathsys.NewBaseFunction("lower", -1, 0, 0),
	}
}

func (f *lowerFunction) Execute(_ pathsys.ContextAccessor, node interface{}, _ []interface{}, _ pathsys.Looper) (interface{}, error) {
	s, err := stringNode(node)
	if s == nil || err != nil {
		return nil, err
	}

	return pathsys.NewString(strings.ToLower(s.String())), nil
}

type replaceFunction struct {
	pathsys.BaseFunction
}

func newReplaceFunction() *replaceFunction {
	return &replaceFunction{
		BaseFunction: pathsys.NewBaseFunction("replace", -1, 2, 2),
	}
}

func (f *replaceFunction) Execute(_ pathsys.ContextAccessor, node interface{}, args []interface{}, _ pathsys.Looper) (interface{}, error) {
	s, err := stringNode(node)
	if s == nil || err != nil {
		return nil, err
	}

	pattern, err := stringNode(args[0])
	if pattern == nil || err != nil {
		return nil, err
	}

	substitution, err := stringNode(args[1])
	if substitution == nil || err != nil {
		return nil, err
	}

	res := strings.ReplaceAll(s.String(), pattern.String(), substitution.String())
	return pathsys.StringOf(res), nil
}

type matchesFunction struct {
	pathsys.BaseFunction
}

func newMatchesFunction() *matchesFunction {
	return &matchesFunction{
		BaseFunction: pathsys.NewBaseFunction("matches", -1, 1, 1),
	}
}

func (f *matchesFunction) Execute(_ pathsys.ContextAccessor, node interface{}, args []interface{}, _ pathsys.Looper) (interface{}, error) {
	s, err := stringNode(node)
	if s == nil || err != nil {
		return nil, err
	}

	re, err := stringNode(args[0])
	if re == nil || err != nil {
		return nil, err
	}

	b, err := regexp.MatchString(re.String(), s.String())
	if err != nil {
		return nil, err
	}
	return pathsys.BooleanOf(b), nil
}

type replaceMatchesFunction struct {
	pathsys.BaseFunction
}

func newReplaceMatchesFunction() *replaceMatchesFunction {
	return &replaceMatchesFunction{
		BaseFunction: pathsys.NewBaseFunction("replaceMatches", -1, 2, 2),
	}
}

func (f *replaceMatchesFunction) Execute(_ pathsys.ContextAccessor, node interface{}, args []interface{}, _ pathsys.Looper) (interface{}, error) {
	s, err := stringNode(node)
	if s == nil || err != nil {
		return nil, err
	}

	regex, err := stringNode(args[0])
	if regex == nil || err != nil {
		return nil, err
	}

	substitution, err := stringNode(args[1])
	if substitution == nil || err != nil {
		return nil, err
	}

	re, err := regexp.Compile(regex.String())
	if err != nil {
		return nil, err
	}

	return pathsys.StringOf(re.ReplaceAllString(s.String(), substitution.String())), nil
}

type lengthFunction struct {
	pathsys.BaseFunction
}

func newLengthFunction() *lengthFunction {
	return &lengthFunction{
		BaseFunction: pathsys.NewBaseFunction("length", -1, 0, 0),
	}
}

func (f *lengthFunction) Execute(_ pathsys.ContextAccessor, node interface{}, _ []interface{}, _ pathsys.Looper) (interface{}, error) {
	s, err := stringNode(node)
	if s == nil || err != nil {
		return nil, err
	}

	return pathsys.NewInteger(int32(utf8.RuneCountInString(s.String()))), nil
}

type toCharsFunction struct {
	pathsys.BaseFunction
}

func newToCharsFunction() *toCharsFunction {
	return &toCharsFunction{
		BaseFunction: pathsys.NewBaseFunction("toChars", -1, 0, 0),
	}
}

func (f *toCharsFunction) Execute(ctx pathsys.ContextAccessor, node interface{}, _ []interface{}, _ pathsys.Looper) (interface{}, error) {
	s, err := stringNode(node)
	if s == nil || err != nil {
		return nil, err
	}

	if s.Length() == 0 {
		return pathsys.NewEmptyCollection(), nil
	}

	col := pathsys.NewCollection(ctx.ModelAdapter())
	sr := []rune(s.String())
	for _, c := range sr {
		col.Add(pathsys.StringOf(string(c)))
	}

	return col, nil
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
