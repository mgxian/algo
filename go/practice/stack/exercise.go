package stack

import (
	"errors"
)

// Browser represents a browser.
type Browser struct {
	forward     Stack
	back        Stack
	currentPage string
	size        int
}

// NewBrowser creates a new initial browser.
func NewBrowser(n int) *Browser {
	return &Browser{
		forward: NewArrayStack(n),
		back:    NewArrayStack(n),
		size:    n,
	}
}

// Goto go to new page.
func (b *Browser) Goto(url string) error {
	if b.currentPage == "" {
		b.currentPage = url
		return nil
	}

	page := b.back.Push(b.currentPage)
	if page == nil {
		return errors.New("push current page to back stack error")
	}
	b.currentPage = url
	b.forward = NewArrayStack(b.size)
	return nil
}

// Forward forward to the last page.
func (b *Browser) Forward() error {
	forwardPage := b.forward.Pop()
	if forwardPage == nil {
		return errors.New("no forward page")
	}

	currentPage := b.back.Push(b.currentPage)
	if currentPage == nil {
		return errors.New("push current page to back stack error")
	}

	b.currentPage = forwardPage.(string)
	return nil
}

// Back back to the last page.
func (b *Browser) Back() error {
	backPage := b.back.Pop()
	if backPage == nil {
		return errors.New("no back page")
	}

	currentPage := b.forward.Push(b.currentPage)
	if currentPage == nil {
		return errors.New("push current page to forward stack error")
	}

	b.currentPage = backPage.(string)
	return nil
}

// CurrentPage returns current page url.
func (b *Browser) CurrentPage() string {
	return b.currentPage
}
