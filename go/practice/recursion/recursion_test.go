package recursion_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mgxian/algo/go/practice/recursion"
)

var _ = Describe("Recursion", func() {
	Describe("fib", func() {
		Context("n = 0", func() {
			n := 0
			It("should return 1", func() {
				result := Fib(n)
				Expect(result).To(Equal(1))
			})
		})

		Context("n = 1", func() {
			n := 1
			It("should return 1", func() {
				result := Fib(n)
				Expect(result).To(Equal(1))
			})
		})

		Context("n = 2", func() {
			n := 2
			It("should return 1", func() {
				result := Fib(n)
				Expect(result).To(Equal(2))
			})
		})

		Context("n = 3", func() {
			n := 3
			It("should return 1", func() {
				result := Fib(n)
				Expect(result).To(Equal(3))
			})
		})

		Context("n = 4", func() {
			n := 4
			It("should return 1", func() {
				result := Fib(n)
				Expect(result).To(Equal(5))
			})
		})
	})
})
