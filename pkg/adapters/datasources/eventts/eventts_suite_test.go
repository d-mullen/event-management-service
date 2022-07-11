package eventts_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestEventTimeseries(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Event Timeseries Data Source Suite")
}
