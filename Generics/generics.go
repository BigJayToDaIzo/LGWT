package generics

import "errors"

type Stack[T any] struct {
	values []T
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var nilValue T
		return nilValue, errors.New("cannot peek empty stack")
	}
	return s.values[len(s.values)-1], nil
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var nilValue T
		return nilValue, errors.New("cannot pop empty stack")
	}
	// we don't need to recapture isEmpty() error here
	v, _ := s.Peek()
	s.values = s.values[:len(s.values)-1]
	return v, nil
}

type StackOfInts struct {
	values []int
}

func (s *StackOfInts) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *StackOfInts) Push(value int) {
	s.values = append(s.values, value)
}

func (s *StackOfInts) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("cannot peek empty stack")
	}
	return s.values[len(s.values)-1], nil
}

func (s *StackOfInts) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("cannot pop empty stack")
	}
	value, _ := s.Peek()
	s.values = s.values[:len(s.values)-1]
	return value, nil
}

type StackOfStrings struct {
	values []string
}

func (s *StackOfStrings) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *StackOfStrings) Push(val string) {
	s.values = append(s.values, val)
}

func (s *StackOfStrings) Peek() (string, error) {
	if s.IsEmpty() {
		return "", errors.New("cannot peek empty stack")
	}
	return s.values[len(s.values)-1], nil
}

func (s *StackOfStrings) Pop() (string, error) {
	if s.IsEmpty() {
		return "", errors.New("cannot pop empty stack")
	}
	v := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return v, nil
}
