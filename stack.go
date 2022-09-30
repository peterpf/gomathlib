package gomathlib

// LIFOStack provides an implementation of a First-In-Last-Out (FILO) data structure.
type LIFOStack[T any] struct {
	Elements          []*T
	currentElementIdx int // Keeps track of where to insert a new element in the fixed-sized stack.
	maxSize           int // Defines the maximum number of elements the stack can hold.
}

// NewLIFOStack initializes a new stack with the given maximum number of elements (maxSize).
func NewLIFOStack[T any](maxSize int) *LIFOStack[T] {
	dataQueue := make([]*T, maxSize)
	return &LIFOStack[T]{
		Elements: dataQueue,
		maxSize:  maxSize,
	}
}

// Push adds the new element to the stack and enforces the maxSize limit by removing the last item.
func (s *LIFOStack[T]) Push(data *T) {
	if s.currentElementIdx >= s.maxSize {
		s.currentElementIdx = 0
	}
	s.Elements[s.currentElementIdx] = data
	s.currentElementIdx++
}

// Size returns the number of non-nil elements in the stack.
func (s *LIFOStack[T]) Size() int {
	var count int
	for _, i := range s.Elements {
		if i != nil {
			count++
		}
	}
	return count
}

// IsFull returns true when no nil-elements are in the stack.
func (s *LIFOStack[T]) IsFull() bool {
	return s.Size() == s.maxSize
}
