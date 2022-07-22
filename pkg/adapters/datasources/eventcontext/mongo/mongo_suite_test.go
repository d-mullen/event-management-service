package mongo_test

import (
	"math/rand"

	. "github.com/onsi/ginkgo/v2"

	. "github.com/onsi/gomega"

	"testing"
)

func TestMongoDB(t *testing.T) {
	RegisterFailHandler(Fail)

	rand.Seed(GinkgoRandomSeed())
	RunSpecs(t, "MongoDB Adapter Test Suite")
}
