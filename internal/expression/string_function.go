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
	"github.com/healthiop/hipath/hipathsys"
	"regexp"
	"strings"
	"unicode/utf8"
)

type indexOfFunction struct {
	hipathsys.BaseFunction
}

func newIndexOfFunction() *indexOfFunction {
	return &indexOfFunction{
		BaseFunction: hipathsys.NewBaseFunction("indexOf", -1, 1, 1),
	}
}

func (f *indexOfFunction) Execute(_ hipathsys.ContextAccessor, node interface{}, args []interface{}, _ hipathsys.Looper) (interface{}, error) {
	s, err := stringNode(node)
	if s == nil || err != nil {
		return nil, err
	}

	ss, err := stringNode(args[0])
	if ss == nil || err != nil {
		return nil, err
	}

	i := strings.Index(s.String(), ss.String())
	return hipathsys.NewInteger(int32(i)), nil
}

type substringFunction struct {
	hipathsys.BaseFunction
}

func newSubstringFunction() *substringFunction {
	return &substringFunction{
		BaseFunction: hipathsys.NewBaseFunction("substring", -1, 1, 2),
	}
}

func (f *substringFunction) Execute(_ hipathsys.ContextAccessor, node interface{}, args []interface{}, _ hipathsys.Looper) (interface{}, error) {
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

	var l hipathsys.IntegerAccessor = nil
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

	return hipathsys.StringOf(string(sr[startVal : startVal+lVal])), nil
}

type startsWithFunction struct {
	hipathsys.BaseFunction
}

func newStartsWithFunction() *startsWithFunction {
	return &startsWithFunction{
		BaseFunction: hipathsys.NewBaseFunction("startsWith", -1, 1, 1),
	}
}

func (f *startsWithFunction) Execute(_ hipathsys.ContextAccessor, node interface{}, args []interface{}, _ hipathsys.Looper) (interface{}, error) {
	s, err := stringNode(node)
	if s == nil || err != nil {
		return nil, err
	}

	ss, err := stringNode(args[0])
	if ss == nil || err != nil {
		return nil, err
	}

	return hipathsys.BooleanOf(strings.HasPrefix(s.String(), ss.String())), nil
}

type endsWithFunction struct {
	hipathsys.BaseFunction
}

func newEndsWithFunction() *endsWithFunction {
	return &endsWithFunction{
		BaseFunction: hipathsys.NewBaseFunction("endsWith", -1, 1, 1),
	}
}

func (f *endsWithFunction) Execute(_ hipathsys.ContextAccessor, node interface{}, args []interface{}, _ hipathsys.Looper) (interface{}, error) {
	s, err := stringNode(node)
	if s == nil || err != nil {
		return nil, err
	}

	ss, err := stringNode(args[0])
	if ss == nil || err != nil {
		return nil, err
	}

	return hipathsys.BooleanOf(strings.HasSuffix(s.String(), ss.String())), nil
}

type containsFunction struct {
	hipathsys.BaseFunction
}

func newContainsFunction() *containsFunction {
	return &containsFunction{
		BaseFunction: hipathsys.NewBaseFunction("contains", -1, 1, 1),
	}
}

func (f *containsFunction) Execute(_ hipathsys.ContextAccessor, node interface{}, args []interface{}, _ hipathsys.Looper) (interface{}, error) {
	s, err := stringNode(node)
	if s == nil || err != nil {
		return nil, err
	}

	ss, err := stringNode(args[0])
	if ss == nil || err != nil {
		return nil, err
	}

	return hipathsys.BooleanOf(strings.Contains(s.String(), ss.String())), nil
}

type upperFunction struct {
	hipathsys.BaseFunction
}

func newUpperFunction() *upperFunction {
	return &upperFunction{
		BaseFunction: hipathsys.NewBaseFunction("upper", -1, 0, 0),
	}
}

func (f *upperFunction) Execute(_ hipathsys.ContextAccessor, node interface{}, _ []interface{}, _ hipathsys.Looper) (interface{}, error) {
	s, err := stringNode(node)
	if s == nil || err != nil {
		return nil, err
	}

	return hipathsys.NewString(strings.ToUpper(s.String())), nil
}

type lowerFunction struct {
	hipathsys.BaseFunction
}

func newLowerFunction() *lowerFunction {
	return &lowerFunction{
		BaseFunction: hipathsys.NewBaseFunction("lower", -1, 0, 0),
	}
}

func (f *lowerFunction) Execute(_ hipathsys.ContextAccessor, node interface{}, _ []interface{}, _ hipathsys.Looper) (interface{}, error) {
	s, err := stringNode(node)
	if s == nil || err != nil {
		return nil, err
	}

	return hipathsys.NewString(strings.ToLower(s.String())), nil
}

type replaceFunction struct {
	hipathsys.BaseFunction
}

func newReplaceFunction() *replaceFunction {
	return &replaceFunction{
		BaseFunction: hipathsys.NewBaseFunction("replace", -1, 2, 2),
	}
}

func (f *replaceFunction) Execute(_ hipathsys.ContextAccessor, node interface{}, args []interface{}, _ hipathsys.Looper) (interface{}, error) {
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
	return hipathsys.StringOf(res), nil
}

type matchesFunction struct {
	hipathsys.BaseFunction
}

func newMatchesFunction() *matchesFunction {
	return &matchesFunction{
		BaseFunction: hipathsys.NewBaseFunction("matches", -1, 1, 1),
	}
}

func (f *matchesFunction) Execute(_ hipathsys.ContextAccessor, node interface{}, args []interface{}, _ hipathsys.Looper) (interface{}, error) {
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
	return hipathsys.BooleanOf(b), nil
}

type replaceMatchesFunction struct {
	hipathsys.BaseFunction
}

func newReplaceMatchesFunction() *replaceMatchesFunction {
	return &replaceMatchesFunction{
		BaseFunction: hipathsys.NewBaseFunction("replaceMatches", -1, 2, 2),
	}
}

func (f *replaceMatchesFunction) Execute(_ hipathsys.ContextAccessor, node interface{}, args []interface{}, _ hipathsys.Looper) (interface{}, error) {
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

	return hipathsys.StringOf(re.ReplaceAllString(s.String(), substitution.String())), nil
}

type lengthFunction struct {
	hipathsys.BaseFunction
}

func newLengthFunction() *lengthFunction {
	return &lengthFunction{
		BaseFunction: hipathsys.NewBaseFunction("length", -1, 0, 0),
	}
}

func (f *lengthFunction) Execute(_ hipathsys.ContextAccessor, node interface{}, _ []interface{}, _ hipathsys.Looper) (interface{}, error) {
	s, err := stringNode(node)
	if s == nil || err != nil {
		return nil, err
	}

	return hipathsys.NewInteger(int32(utf8.RuneCountInString(s.String()))), nil
}

type toCharsFunction struct {
	hipathsys.BaseFunction
}

func newToCharsFunction() *toCharsFunction {
	return &toCharsFunction{
		BaseFunction: hipathsys.NewBaseFunction("toChars", -1, 0, 0),
	}
}

func (f *toCharsFunction) Execute(ctx hipathsys.ContextAccessor, node interface{}, _ []interface{}, _ hipathsys.Looper) (interface{}, error) {
	s, err := stringNode(node)
	if s == nil || err != nil {
		return nil, err
	}

	if s.Length() == 0 {
		return hipathsys.NewEmptyCol(), nil
	}

	col := ctx.NewCol()
	sr := []rune(s.String())
	for _, c := range sr {
		col.Add(hipathsys.StringOf(string(c)))
	}

	return col, nil
}

func stringNode(node interface{}) (hipathsys.StringAccessor, error) {
	value := unwrapCollection(node)
	if value == nil {
		return nil, nil
	}

	if s, ok := value.(hipathsys.StringAccessor); !ok {
		return nil, fmt.Errorf("not a string: %T", value)
	} else {
		return s, nil
	}
}
