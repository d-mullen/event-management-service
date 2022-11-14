package matchers

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
)

var (
	msgActualIsNotProto = "actual is not a proto.Message"
)

func EqualProtoSlice(expected interface{}) types.GomegaMatcher {
	return EqualCmp(expected, protocmp.Transform())
}

func EqualProto(expected proto.Message, options ...cmp.Option) types.GomegaMatcher {
	options = append(
		options,
		cmp.Comparer(proto.Equal),
	)

	return &protoMessageEqual{
		expected: expected,
		options:  options,
	}
}

type protoMessageEqual struct {
	expected proto.Message
	options  cmp.Options
}

func (matcher *protoMessageEqual) Match(actual interface{}) (bool, error) {
	if ok, err := handleNilArgs(actual, matcher.expected); !ok {
		return false, err
	}

	message, ok := actual.(proto.Message)
	if !ok {
		return false, fmt.Errorf("matcher expects a proto message.  Got:\n%s", format.Object(actual, 1))
	}

	return cmp.Equal(message, matcher.expected, matcher.options), nil
}

func (matcher *protoMessageEqual) FailureMessage(actual interface{}) string {
	actualMessage, ok := actual.(proto.Message)
	if !ok {
		return msgActualIsNotProto
	}

	diff := cmp.Diff(actual, matcher.expected, protocmp.Transform())
	prettyDiff := "\nMismatch (-want +got)\n" + format.IndentString(diff, 1)

	return format.Message(prototext.Format(actualMessage), "to equal", prototext.Format(matcher.expected)) + prettyDiff
}

func (matcher *protoMessageEqual) NegatedFailureMessage(actual interface{}) string {
	actualMessage, ok := actual.(proto.Message)
	if !ok {
		return msgActualIsNotProto
	}

	return format.Message(prototext.Format(actualMessage), "not to equal", prototext.Format(matcher.expected))
}
