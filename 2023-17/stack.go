package main

type PointStep struct {
	X, Y  int
	Steps int
}

type Stack []PointStep

// Push adds an element to the top of the stack
func (s *Stack) Push(v PointStep) {
	*s = append(*s, v)
}

// Pop removes and returns the top element of the stack
func (s *Stack) Pop() (PointStep, bool) {
	if len(*s) == 0 {
		return PointStep{}, false
	}

	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}
