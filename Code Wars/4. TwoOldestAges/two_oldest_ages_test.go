package two_oldest_ages

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TwoOldestAges", func() {
	It("should return 18 and 83 for input []int{6,5,83,5,3,18}", func() {
		Expect(TwoOldestAges1([]int{6, 5, 83, 5, 3, 18})).To(Equal([2]int{18, 83}))
		Expect(TwoOldestAges2([]int{6, 5, 83, 5, 3, 18})).To(Equal([2]int{18, 83}))
	})
	It("should return 45 and 87 for input []int{1,5,87,45,8,8}", func() {
		Expect(TwoOldestAges1([]int{1, 5, 87, 45, 8, 8})).To(Equal([2]int{45, 87}))
		Expect(TwoOldestAges2([]int{1, 5, 87, 45, 8, 8})).To(Equal([2]int{45, 87}))
	})
})
