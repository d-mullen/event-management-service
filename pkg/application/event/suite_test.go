package event_test

import (
	"math/rand"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var suiteTestingT *testing.T

func TestEventQueryService(t *testing.T) {
	suiteTestingT = t
	RegisterFailHandler(Fail)

	rand.Seed(GinkgoRandomSeed())
	RunSpecs(t, "EventQuery Service Test Suite")
}
