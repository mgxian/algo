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

	Describe("factorial", func() {
		Context("n = 0", func() {
			n := 0
			It("should return 1", func() {
				result := Factorial(n)
				Expect(result).To(Equal(1))
			})
		})

		Context("n = 1", func() {
			n := 1
			It("should return 1", func() {
				result := Factorial(n)
				Expect(result).To(Equal(1))
			})
		})

		Context("n = 2", func() {
			n := 2
			It("should return 2", func() {
				result := Factorial(n)
				Expect(result).To(Equal(2))
			})
		})

		Context("n = 3", func() {
			n := 3
			It("should return 6", func() {
				result := Factorial(n)
				Expect(result).To(Equal(6))
			})
		})

		Context("n = 4", func() {
			n := 4
			It("should return 24", func() {
				result := Factorial(n)
				Expect(result).To(Equal(24))
			})
		})
	})

	Describe("permutation", func() {
		Context("[]", func() {
			data := []int{}
			It("should return 1", func() {
				Permutation(data, 0, len(data)-1)
			})
		})

		Context("[0]", func() {
			data := []int{0}
			It("should return 1", func() {
				Permutation(data, 0, len(data)-1)
			})
		})

		Context("[1, 2]", func() {
			data := []int{1, 2}
			It("should return 1", func() {
				Permutation(data, 0, len(data)-1)
			})
		})

		Context("[1, 2, 3]", func() {
			data := []int{1, 2, 3}
			It("should return 1", func() {
				Permutation(data, 0, len(data)-1)
			})
		})

		Context("[1, 2, 3, 4]", func() {
			data := []int{1, 2, 3, 4}
			It("should return 1", func() {
				Permutation(data, 0, len(data)-1)
			})
		})
	})
})
