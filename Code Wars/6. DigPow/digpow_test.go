package digpow

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func dotest(n, p int, exp int) {
	var ans = DigPow(n, p)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		dotest(89, 1, 1)
		dotest(92, 1, -1)
	})
})
