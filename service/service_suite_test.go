package service_test

import (
	"math/rand"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"

	"testing"
)

func TestService(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("junit.xml")
	rand.Seed(GinkgoRandomSeed())
	RunSpecsWithDefaultAndCustomReporters(t, "Event-Management-Service Service Suite", []Reporter{junitReporter})
}
