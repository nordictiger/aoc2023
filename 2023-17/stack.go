package main

type Point struct {
	X, Y     int
	Straight int
}

type Stack []Point

// Push adds an element to the top of the stack
func (s *Stack) Push(v Point) {
	*s = append(*s, v)
}

// Pop removes and returns the top element of the stack
func (s *Stack) Pop() (Point, bool) {
	if len(*s) == 0 {
		return Point{}, false
	}

	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}
