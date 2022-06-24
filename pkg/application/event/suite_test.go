package event_test

import (
	"math/rand"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

var (
	suiteTestingT *testing.T
)

func TestEventQueryService(t *testing.T) {
	suiteTestingT = t
	RegisterFailHandler(Fail)

	rand.Seed(GinkgoRandomSeed())
	RunSpecs(t, "EventQuery Service Test Suite")
}
