package grpc_test

import (
	"math/rand"

	. "github.com/onsi/ginkgo/v2"

	. "github.com/onsi/gomega"

	"testing"
)

func TestEventQueryGRPC(t *testing.T) {
	RegisterFailHandler(Fail)

	rand.Seed(GinkgoRandomSeed())
	RunSpecs(t, "EventQuery GRPC Adapter Test Suite")
}
