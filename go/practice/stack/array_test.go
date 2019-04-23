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

	BeforeEach(func() {
		stack = NewArrayStack(1)
	})

	Describe("Test array stack length", func() {
		It("should be a true", func() {
			Expect(stack.IsEmpty()).To(Equal(true))
		})
	})
})
