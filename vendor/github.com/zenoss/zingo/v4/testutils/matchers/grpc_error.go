package matchers

import (
	"errors"

	"github.com/google/go-cmp/cmp"
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
)

var (
	errActualIsNotGoError     = errors.New("actual does not implement error interface")
	errActualIsNotGrpcError   = errors.New("actual is not grpc status error")
	errExpectedIsNotGrpcError = errors.New("expected is not grpc status error")
)

func EqualGrpcError(expected error) types.GomegaMatcher {
	return &matchGrpcError{expected: expected}
}

type matchGrpcError struct {
	expected error
}

func (m matchGrpcError) Match(actual interface{}) (bool, error) {
	if ok, err := handleNilArgs(actual, m.expected); !ok {
		return false, err
	}

	actualMsg, expectedMsg, err := m.getProto(actual)
	if err != nil {
		return false, err
	}

	return proto.Equal(actualMsg.Proto(), expectedMsg.Proto()), nil
}

func (m matchGrpcError) FailureMessage(actual interface{}) (message string) {
	actualMsg, expectedMsg, err := m.getProto(actual)
	if err != nil {
		return err.Error()
	}

	diff := cmp.Diff(actualMsg, expectedMsg, protocmp.Transform())
	prettyDiff := "\nMismatch (-want +got)\n" + format.IndentString(diff, 1)

	return format.Message(
		prototext.Format(actualMsg.Proto()),
		"to equal",
		prototext.Format(expectedMsg.Proto()),
	) + prettyDiff
}

func (m matchGrpcError) NegatedFailureMessage(actual interface{}) (message string) {
	actualMsg, expectedMsg, err := m.getProto(actual)
	if err != nil {
		return err.Error()
	}

	return format.Message(
		prototext.Format(actualMsg.Proto()),
		"not to equal",
		prototext.Format(expectedMsg.Proto()),
	)
}

func (m matchGrpcError) getProto(actual interface{}) (*status.Status, *status.Status, error) {
	err, ok := actual.(error)
	if !ok {
		return nil, nil, errActualIsNotGoError
	}

	actualMsg, ok := status.FromError(err)
	if !ok {
		return nil, nil, errActualIsNotGrpcError
	}

	expectedMsg, ok := status.FromError(m.expected)
	if !ok {
		return nil, nil, errExpectedIsNotGrpcError
	}

	return actualMsg, expectedMsg, nil
}
