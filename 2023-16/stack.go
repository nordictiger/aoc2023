package main

type Stack []Position

// Push adds an element to the top of the stack
func (s *Stack) Push(v Position) {
	*s = append(*s, v)
}

// Pop removes and returns the top element of the stack
func (s *Stack) Pop() (Position, bool) {
	if len(*s) == 0 {
		return Position{}, false
	}

	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}
