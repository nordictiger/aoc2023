package main

// Queue represents a FIFO queue
type Queue []signal

// Enqueue adds an element to the end of the queue
func (q *Queue) Enqueue(v signal) {
	*q = append(*q, v)
}

// Dequeue removes and returns the front element of the queue
func (q *Queue) Dequeue() (signal, bool) {
	if len(*q) == 0 {
		return signal{}, false
	}
	element := (*q)[0]
	*q = (*q)[1:]
	return element, true
}
