package gomathlib

import "testing"

func TestPush(t *testing.T) {
	// Arrange
	stack := NewLIFOStack[int](5)
	for i := 0; i < 5; i++ {
		stack.Push(intToPtr(i))
	}
	// Act
	stack.Push(intToPtr(6))

	// Assert
	if stack.Size() != 5 {
		t.Errorf("expected stack to have 5 elements, got %d", stack.Size())
	}
}

func TestIsFull(t *testing.T) {
	// Arrange
	stack := NewLIFOStack[int](5)
	for i := 0; i < 5; i++ {
		stack.Push(intToPtr(i))
	}
	// Act && Assert
	if stack.IsFull() != true {
		t.Errorf("expected stack to be full, got %v", stack.Elements)
	}
}

// Convert integer to pointer.
func intToPtr(i int) *int {
	ptr := i
	return &ptr
}
