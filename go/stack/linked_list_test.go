package stack

import "testing"

func TestLinkedListStackPush(t *testing.T) {
	elements := []int{5, 4, 3, 2, 1}
	expected := 1

	var aStack Stack
	aStack = NewLinkedListStack(len(elements))
	for _, e := range elements {
		if err := aStack.Push(e); err != nil {
			t.Errorf("push %v error %s", e, err)
		}
	}

	actual, err := aStack.Top()
	if err != nil {
		t.Errorf("Top error %s", err)
	}

	if actual != expected {
		t.Errorf("expected to be %v, but instead got %v", expected, actual)
	}
}

func TestLinkedListStackPop(t *testing.T) {
	elements := []int{1, 2, 3, 4, 5}
	expected := 5

	var aStack Stack
	aStack = NewLinkedListStack(len(elements))
	for _, e := range elements {
		if err := aStack.Push(e); err != nil {
			t.Errorf("push %v error %s", e, err)
		}
	}

	actual, err := aStack.Pop()
	if err != nil {
		t.Errorf("Pop error %s", err)
	}

	if actual != expected {
		t.Errorf("expected to be %v, but instead got %v", expected, actual)
	}
}

func TestLinkedListStackIsEmpty(t *testing.T) {
	elements := []int{5}
	expected := []bool{false, true}

	var aStack Stack
	aStack = NewLinkedListStack(len(elements))
	for _, e := range elements {
		if err := aStack.Push(e); err != nil {
			t.Errorf("push %v error %s", e, err)
		}
	}

	actual := aStack.isEmpty()
	if actual != expected[0] {
		t.Errorf("expected to be %v, but instead got %v", expected, actual)
	}

	_, err := aStack.Pop()
	if err != nil {
		t.Errorf("Pop error %s", err)
	}

	actual = aStack.isEmpty()
	if actual != expected[1] {
		t.Errorf("expected to be %v, but instead got %v", expected, actual)
	}
}

func TestLinkedListStackIsFull(t *testing.T) {
	elements := []int{5, 4, 3, 2, 1}
	expected := []bool{true, false}

	var aStack Stack
	aStack = NewLinkedListStack(len(elements))
	for _, e := range elements {
		if err := aStack.Push(e); err != nil {
			t.Errorf("push %v error %s", e, err)
		}
	}

	actual := aStack.isFull()
	if actual != expected[0] {
		t.Errorf("expected to be %v, but instead got %v", expected, actual)
	}

	_, err := aStack.Pop()
	if err != nil {
		t.Errorf("Pop error %s", err)
	}

	actual = aStack.isFull()
	if actual != expected[1] {
		t.Errorf("expected to be %v, but instead got %v", expected, actual)
	}
}
