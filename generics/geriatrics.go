package geriatrics

type Stack[T any] struct {
	stack []T
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *Stack[T]) Push(value T) {
	s.stack = append(s.stack, value)
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	return s.stack[len(s.stack)-1], true
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	popped := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return popped, true
}
