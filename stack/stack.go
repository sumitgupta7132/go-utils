package stack

type Stack[T any] struct {
	stk []T
}

func (s *Stack[T]) Empty() bool {
	return len(s.stk) == 0
}

func (s *Stack[T]) Push(value T) {
	s.stk = append(s.stk, value)
}

func (s *Stack[T]) Pop() {
	if s.Empty() {
		return
	}
	s.stk = s.stk[:len(s.stk)-1]
}

// empty check is expected before this
func (s *Stack[T]) Top() T {
	return s.stk[len(s.stk)-1]
}
