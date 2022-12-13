package minstack

type Stack struct {
	thestack []int
	minstack []int
}

func (s *Stack) IsEmpty() bool {
	return len(s.minstack) == 0
}

func (s *Stack) Push(a int) {
	s.thestack = append(s.thestack, a)
	if s.IsEmpty() {
		s.minstack = append(s.minstack, a)
	} else if a < s.Peek() {
		s.minstack = append(s.minstack, a)
	}
}

func (s *Stack) Pop() {
	i := len(s.thestack) - 1
	popped := s.thestack[i]
	j := (len(s.minstack) - 1)
	s.thestack = s.thestack[:i]
	if popped == s.Peek() {
		s.minstack = s.minstack[:j]
	}
}

func (s *Stack) Peek() int {
	i := len(s.minstack) - 1
	return s.minstack[i]
}

func (s *Stack) min() int {
	i := len(s.minstack) - 1
	return s.minstack[i]
}
