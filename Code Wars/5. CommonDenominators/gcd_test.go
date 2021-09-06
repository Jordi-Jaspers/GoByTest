package gcd

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func dotest(a [][]int, exp string) {
	var ans = ConvertFracts(a)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("ConvertFracts", func() {
	It("Basic tests", func() {
		var lst = [][]int{{1, 2}, {1, 3}, {1, 4}}
		dotest(lst, "(6,12)(4,12)(3,12)")

		lst = [][]int{{69, 130}, {87, 1310}, {30, 40}}
		dotest(lst, "(18078,34060)(2262,34060)(25545,34060)")
	})
})
