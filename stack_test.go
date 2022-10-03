package gomathlib

import "testing"

func TestPush(t *testing.T) {
	// Arrange
	stack := NewFIFOStack[int](5)
	for i := 0; i < 5; i++ {
		stack.Push(intToPtr(i + 1))
	}
	expectedElements := []int{3, 4, 5, 6, 7}
	// Act
	stack.Push(intToPtr(6))
	stack.Push(intToPtr(7))

	// Assert
	if stack.Size() != 5 {
		t.Errorf("expected stack to have 5 elements, got %d", stack.Size())
	}
	gotElements := stack.Elements()
	for i := 0; i < 5; i++ {
		if *gotElements[i] != expectedElements[i] {
			t.Errorf("expected element %d at position %d, got %d", expectedElements[i], i, *stack.elements[i])
		}
	}
}

func TestIsFull(t *testing.T) {
	// Arrange
	stack := NewFIFOStack[int](5)
	for i := 0; i < 5; i++ {
		stack.Push(intToPtr(i))
	}
	// Act && Assert
	if stack.IsFull() != true {
		t.Errorf("expected stack to be full, got %v", stack.elements)
	}
}

// Convert integer to pointer.
func intToPtr(i int) *int {
	ptr := i
	return &ptr
}
