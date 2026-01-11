package stack

type Stack[T any] struct {
	elements []T
}

// Adds an element to the top of the stack
func (s *Stack[T]) Push(val T) {
	// this will create a new slice with val as the first element in it,
	// then "spread" the elements from s.elements into it.
	s.elements = append([]T{val}, s.elements...)
}

// removes an element from the top of the stack
func (s *Stack[T]) Pop(val T) {
	s.elements = s.elements[1:]
}

// returns the element at the top of the stack
func (s *Stack[T]) Peek() T {
	return s.elements[0]
}

// checks whether the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	if len(s.elements) > 0 {
		return false
	}
	return true
}
