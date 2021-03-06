package queue_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mgxian/algo/go/practice/queue"
)

var _ = Describe("Ring", func() {
	Describe("create a ring queue", func() {
		Context("when queue size is zero", func() {
			queue := NewRingQueue(0)
			It("should return nil", func() {
				Expect(queue).To(BeNil())
			})
		})

		Context("when queue size is greater than zero", func() {
			queue := NewRingQueue(1)
			It("should not return nil", func() {
				Expect(queue).NotTo(BeNil())
			})
		})
	})

	Describe("enqueue a element to ring queue", func() {
		Context("when queue size is one", func() {
			queue := NewRingQueue(1)
			It("should return right elements", func() {
				e := queue.Enqueue(1)
				Expect(e).To(Equal(1))
				Expect(queue.IsEmpty()).To(Equal(false))
				Expect(queue.IsFull()).To(Equal(true))

				e = queue.Enqueue(1)
				Expect(e).To(BeNil())
			})
		})

		Context("when queue size is greater than one", func() {
			queue := NewRingQueue(2)
			It("should return right elements", func() {
				e := queue.Enqueue(1)
				Expect(e).To(Equal(1))
				Expect(queue.IsEmpty()).To(Equal(false))
				Expect(queue.IsFull()).To(Equal(false))

				queue.Enqueue(1)
				e = queue.Enqueue(1)
				Expect(e).To(BeNil())
			})
		})
	})

	Describe("dequeue a element from ring queue", func() {
		Context("when queue size is one", func() {
			queue := NewRingQueue(1)
			queue.Enqueue(1)
			It("should return right elements", func() {
				e := queue.Dequeue()
				Expect(e).To(Equal(1))
				Expect(queue.IsEmpty()).To(Equal(true))
				Expect(queue.IsFull()).To(Equal(false))

				e = queue.Dequeue()
				Expect(e).To(BeNil())
			})
		})

		Context("when queue size is greater than one", func() {
			queue := NewRingQueue(2)
			queue.Enqueue(1)
			queue.Enqueue(2)
			It("should return right elements", func() {
				e2 := queue.Dequeue()
				Expect(e2).To(Equal(1))
				e1 := queue.Dequeue()
				Expect(e1).To(Equal(2))

				e := queue.Dequeue()
				Expect(e).To(BeNil())
			})
		})
	})

	Describe("peek a element from ring queue", func() {
		Context("when queue size is one", func() {
			queue := NewRingQueue(1)

			
			It("should return right elements", func() {
				e := queue.Peek()
				Expect(e).To(BeNil())

				queue.Enqueue(1)

				e = queue.Peek()
				Expect(e).To(Equal(1))
				Expect(queue.IsEmpty()).To(Equal(false))
				Expect(queue.IsFull()).To(Equal(true))
			})
		})

		Context("when queue size is greater than one", func() {
			queue := NewRingQueue(2)
			queue.Enqueue(1)
			It("should return right elements", func() {
				e1 := queue.Peek()
				Expect(e1).To(Equal(1))
				Expect(queue.IsEmpty()).To(Equal(false))
				Expect(queue.IsFull()).To(Equal(false))

				queue.Dequeue()
				e := queue.Peek()
				Expect(e).To(BeNil())
			})
		})
	})

	Describe("check if ring queue is empty", func() {
		Context("when queue size is one", func() {
			queue := NewRingQueue(1)

			It("should return true", func() {
				Expect(queue.IsEmpty()).To(Equal(true))
			})

			It("should return false", func() {
				queue.Enqueue(1)
				Expect(queue.IsEmpty()).To(Equal(false))

				queue.Dequeue()
				Expect(queue.IsEmpty()).To(Equal(true))
			})
		})

		Context("when queue size is greater than one", func() {
			queue := NewRingQueue(2)

			It("should return true", func() {
				Expect(queue.IsEmpty()).To(Equal(true))
			})

			It("should return false", func() {
				queue.Enqueue(1)
				queue.Enqueue(2)
				queue.Dequeue()
				Expect(queue.IsEmpty()).To(Equal(false))

				queue.Dequeue()
				Expect(queue.IsEmpty()).To(Equal(true))
			})
		})
	})

	Describe("check if ring queue is full", func() {
		Context("when queue size is one", func() {
			queue := NewRingQueue(1)

			It("should return false", func() {
				Expect(queue.IsFull()).To(Equal(false))
			})

			It("should return true", func() {
				queue.Enqueue(1)
				Expect(queue.IsFull()).To(Equal(true))
			})

			It("should return true", func() {
				queue.Dequeue()
				queue.Enqueue(1)
				Expect(queue.IsFull()).To(Equal(true))
			})
		})

		Context("when queue size is greater than one", func() {
			queue := NewRingQueue(2)

			It("should return false", func() {
				Expect(queue.IsFull()).To(Equal(false))
			})

			It("should return true", func() {
				queue.Enqueue(1)
				queue.Enqueue(2)
				Expect(queue.IsFull()).To(Equal(true))
			})

			It("should return true", func() {
				queue.Dequeue()
				Expect(queue.IsFull()).To(Equal(false))

				queue.Enqueue(1)
				queue.Enqueue(1)
				Expect(queue.IsFull()).To(Equal(true))
			})
		})
	})
})
