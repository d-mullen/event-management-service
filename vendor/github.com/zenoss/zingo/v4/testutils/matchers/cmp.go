package matchers

import (
	"github.com/google/go-cmp/cmp"
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
)

func EqualCmp(expected interface{}, options ...cmp.Option) types.GomegaMatcher {
	return &equalCmpMatcher{
		expected: expected,
		options:  options,
	}
}

type equalCmpMatcher struct {
	expected interface{}
	options  cmp.Options
}

func (matcher *equalCmpMatcher) Match(actual interface{}) (success bool, err error) {
	if actual == nil && matcher.expected == nil {
		return false, errBothArgsNil
	}
	return cmp.Equal(actual, matcher.expected, matcher.options), nil
}

func (matcher *equalCmpMatcher) FailureMessage(actual interface{}) (message string) {
	actualString, actualOk := actual.(string)
	expectedString, expectedOk := matcher.expected.(string)
	if actualOk && expectedOk {
		return format.MessageWithDiff(actualString, "to equal", expectedString)
	}
	diff := cmp.Diff(actual, matcher.expected, matcher.options)
	return format.Message(actual, "to equal", matcher.expected) +
		"\n\nDiff:\n" + format.IndentString(diff, 1)
}

func (matcher *equalCmpMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	diff := cmp.Diff(actual, matcher.expected, matcher.options)
	return format.Message(actual, "not to equal", matcher.expected) +
		"\n\nDiff:\n" + format.IndentString(diff, 1)
}
