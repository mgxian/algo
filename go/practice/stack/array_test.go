package stack_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mgxian/algo/go/practice/stack"
)

var _ = Describe("Array", func() {
	var (
		stack *ArrayStack
	)

	Describe("Zero size array stack", func() {
		BeforeEach(func() {
			stack = NewArrayStack(0)
		})

		It("should return nil", func() {
			Expect(stack).To(Equal(nil))
		})

		It("should return true", func() {
			Expect(stack.IsEmpty()).To(Equal(true))
		})
	})
})
