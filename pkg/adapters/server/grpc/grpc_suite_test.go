package grpc_test

import (
	"math/rand"
	"testing"

	. "github.com/onsi/ginkgo/v2"

	. "github.com/onsi/gomega"
)

func TestEventQueryGRPC(t *testing.T) {
	RegisterFailHandler(Fail)

	rand.Seed(GinkgoRandomSeed())
	RunSpecs(t, "EventQuery gRPC Service Adapter Test Suite")
}
