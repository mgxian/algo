package stack_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mgxian/algo/go/practice/stack"
)

var _ = Describe("List", func() {
	Describe("create a list stack", func() {
		Context("when stack size is zero", func() {
			stack := NewListStack(0)
			It("should return nil", func() {
				Expect(stack).To(BeNil())
			})
		})

		Context("when stack size is greater than zero", func() {
			stack := NewListStack(1)
			It("should not return nil", func() {
				Expect(stack).NotTo(BeNil())
			})
		})
	})

	Describe("push a element to list stack", func() {
		Context("when stack size is one", func() {
			stack := NewListStack(1)
			It("should return right elements", func() {
				e := stack.Push(1)
				Expect(e).To(Equal(1))
				Expect(stack.IsEmpty()).To(Equal(false))
				Expect(stack.IsFull()).To(Equal(true))

				e = stack.Push(1)
				Expect(e).To(BeNil())
			})
		})

		Context("when stack size is greater than one", func() {
			stack := NewListStack(2)
			It("should return right elements", func() {
				e := stack.Push(1)
				Expect(e).To(Equal(1))
				Expect(stack.IsEmpty()).To(Equal(false))
				Expect(stack.IsFull()).To(Equal(false))

				stack.Push(1)
				e = stack.Push(1)
				Expect(e).To(BeNil())
			})
		})
	})

	Describe("pop a element from list stack", func() {
		Context("when stack size is one", func() {
			stack := NewListStack(1)
			stack.Push(1)
			It("should return right elements", func() {
				e := stack.Pop()
				Expect(e).To(Equal(1))
				Expect(stack.IsEmpty()).To(Equal(true))
				Expect(stack.IsFull()).To(Equal(false))

				e = stack.Pop()
				Expect(e).To(BeNil())
			})
		})

		Context("when stack size is greater than one", func() {
			stack := NewListStack(2)
			stack.Push(1)
			stack.Push(2)
			It("should return right elements", func() {
				e2 := stack.Pop()
				Expect(e2).To(Equal(2))
				e1 := stack.Pop()
				Expect(e1).To(Equal(1))

				e := stack.Pop()
				Expect(e).To(BeNil())
			})
		})
	})

	Describe("peek a element from list stack", func() {
		Context("when stack size is one", func() {
			stack := NewListStack(1)

			
			It("should return right elements", func() {
				e := stack.Peek()
				Expect(e).To(BeNil())

				stack.Push(1)

				e = stack.Peek()
				Expect(e).To(Equal(1))
				Expect(stack.IsEmpty()).To(Equal(false))
				Expect(stack.IsFull()).To(Equal(true))
			})
		})

		Context("when stack size is greater than one", func() {
			stack := NewListStack(2)
			stack.Push(1)
			It("should return right elements", func() {
				e1 := stack.Peek()
				Expect(e1).To(Equal(1))
				Expect(stack.IsEmpty()).To(Equal(false))
				Expect(stack.IsFull()).To(Equal(false))

				stack.Pop()
				e := stack.Peek()
				Expect(e).To(BeNil())
			})
		})
	})

	Describe("check if list stack is empty", func() {
		Context("when stack size is one", func() {
			stack := NewListStack(1)

			It("should return true", func() {
				Expect(stack.IsEmpty()).To(Equal(true))
			})

			It("should return false", func() {
				stack.Push(1)
				Expect(stack.IsEmpty()).To(Equal(false))

				stack.Pop()
				Expect(stack.IsEmpty()).To(Equal(true))
			})
		})

		Context("when stack size is greater than one", func() {
			stack := NewListStack(2)

			It("should return true", func() {
				Expect(stack.IsEmpty()).To(Equal(true))
			})

			It("should return false", func() {
				stack.Push(1)
				stack.Push(2)
				stack.Pop()
				Expect(stack.IsEmpty()).To(Equal(false))

				stack.Pop()
				Expect(stack.IsEmpty()).To(Equal(true))
			})
		})
	})

	Describe("check if list stack is full", func() {
		Context("when stack size is one", func() {
			stack := NewListStack(1)

			It("should return false", func() {
				Expect(stack.IsFull()).To(Equal(false))
			})

			It("should return true", func() {
				stack.Push(1)
				Expect(stack.IsFull()).To(Equal(true))
			})

			It("should return true", func() {
				stack.Pop()
				stack.Push(1)
				Expect(stack.IsFull()).To(Equal(true))
			})
		})

		Context("when stack size is greater than one", func() {
			stack := NewListStack(2)

			It("should return false", func() {
				Expect(stack.IsFull()).To(Equal(false))
			})

			It("should return true", func() {
				stack.Push(1)
				stack.Push(2)
				Expect(stack.IsFull()).To(Equal(true))
			})

			It("should return true", func() {
				e := stack.Pop()
				Expect(e).To(Equal(2))
				
				stack.Pop()
				stack.Push(1)
				stack.Push(1)
				Expect(stack.IsFull()).To(Equal(true))
			})
		})
	})
})
