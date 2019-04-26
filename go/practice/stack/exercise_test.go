package stack_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mgxian/algo/go/practice/stack"
)

var _ = Describe("Browser", func() {
	Describe("create a browser", func() {
		Context("when browser stack size is smaller than go to pages", func() {
			browser := NewBrowser(1)
			It("should not return nil", func() {
				Expect(browser).NotTo(BeNil())
			})

			It("should not return error", func() {
				err := browser.Goto("A")
				Expect(err).NotTo(HaveOccurred())
			})

			It("should return error", func() {
				err := browser.Goto("A")
				err = browser.Goto("B")
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when browser stack size is greater than go to pages", func() {
			browser := NewBrowser(10)
			It("should not return nil", func() {
				Expect(browser).NotTo(BeNil())
			})

			It("should not return error", func() {
				browser.Goto("A")
				browser.Goto("B")
				browser.Goto("C")
				err := browser.Goto("D")
				Expect(err).NotTo(HaveOccurred())

				err = browser.Back()
				Expect(err).NotTo(HaveOccurred())

				curr := browser.CurrentPage()
				Expect(curr).To(Equal("C"))

				err = browser.Goto("E")
				Expect(err).NotTo(HaveOccurred())

				err = browser.Goto("F")
				Expect(err).NotTo(HaveOccurred())

				curr = browser.CurrentPage()
				Expect(curr).To(Equal("F"))

				err = browser.Back()
				Expect(err).NotTo(HaveOccurred())

				curr = browser.CurrentPage()
				Expect(curr).To(Equal("E"))

				err = browser.Back()
				Expect(err).NotTo(HaveOccurred())

				curr = browser.CurrentPage()
				Expect(curr).To(Equal("C"))

				err = browser.Back()
				Expect(err).NotTo(HaveOccurred())

				curr = browser.CurrentPage()
				Expect(curr).To(Equal("B"))

				err = browser.Back()
				Expect(err).NotTo(HaveOccurred())

				curr = browser.CurrentPage()
				Expect(curr).To(Equal("A"))

				err = browser.Back()
				Expect(err).To(HaveOccurred())

				curr = browser.CurrentPage()
				Expect(curr).To(Equal("A"))

				err = browser.Forward()
				Expect(err).NotTo(HaveOccurred())

				curr = browser.CurrentPage()
				Expect(curr).To(Equal("B"))

				err = browser.Forward()
				Expect(err).NotTo(HaveOccurred())

				curr = browser.CurrentPage()
				Expect(curr).To(Equal("C"))

				err = browser.Forward()
				Expect(err).NotTo(HaveOccurred())

				curr = browser.CurrentPage()
				Expect(curr).To(Equal("E"))

				err = browser.Back()
				Expect(err).NotTo(HaveOccurred())

				curr = browser.CurrentPage()
				Expect(curr).To(Equal("C"))

				err = browser.Forward()
				Expect(err).NotTo(HaveOccurred())

				curr = browser.CurrentPage()
				Expect(curr).To(Equal("E"))
			})
		})
	})
})