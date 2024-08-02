package geriatrics

type StackOfInts struct {
	stack []int
}

func (s *StackOfInts) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *StackOfInts) Push(i int) {
	s.stack = append(s.stack, i)
}

func (s *StackOfInts) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	return s.stack[len(s.stack)-1], true
}

func (s *StackOfInts) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	popped := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return popped, true
}

type StackOfStrings struct {
	stack []string
}

func (s *StackOfStrings) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *StackOfStrings) Push(i string) {
	s.stack = append(s.stack, i)
}

func (s *StackOfStrings) Peek() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}
	return s.stack[len(s.stack)-1], true
}

func (s *StackOfStrings) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}
	popped := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return popped, true
}
