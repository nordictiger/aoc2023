package main

type Stack []condition

// Push adds an element to the top of the stack
func (s *Stack) Push(v condition) {
	*s = append(*s, v)
}

// Pop removes and returns the top element of the stack
func (s *Stack) Pop() (condition, bool) {
	if len(*s) == 0 {
		return condition{}, false
	}

	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}
