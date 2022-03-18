package service_test

import (
	"math/rand"

	. "github.com/onsi/ginkgo/v2"

	. "github.com/onsi/gomega"

	"testing"
)

func TestService(t *testing.T) {
	RegisterFailHandler(Fail)

	rand.Seed(GinkgoRandomSeed())
	RunSpecs(t, "Event-Management-Service Service Suite")
}
