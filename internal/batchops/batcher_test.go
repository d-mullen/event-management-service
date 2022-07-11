package batchops_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/zenoss/event-management-service/internal/batchops"
)

var _ = Describe("Batcher Unit Tests", func() {
	It("should batch up larger payloads into smaller batches and process them", func() {
		i := 0
		err := batchops.Do(2, []int{1, 2, 3, 4, 5, 6, 7},
			func(batch []int) (int, error) {
				GinkgoWriter.Printf("got batch %v\n", batch)
				if i == 0 {
					Ω(batch).Should(ContainElements(1, 2))
				}
				if i == 1 {
					Ω(batch).Should(ContainElements(3, 4))
				}
				i++
				return i, nil
			}, func(r int) (bool, error) {
				Ω(r).Should(Equal(i))
				return true, nil
			},
		)
		Ω(err).ShouldNot(HaveOccurred())
	})
})
