package activeents_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestActiveEntities(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Active Entities Suite")
}
