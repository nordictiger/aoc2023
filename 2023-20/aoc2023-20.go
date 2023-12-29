package main

import (
	"fmt"
	"strconv"
	"time"
)

var (
	lowCounter  int
	highCounter int
)

func compileIncoming(mc *moduleConfiguration) {
	for k, node := range *mc {
		if node.moduleType == Conjunction {
			for searchKey, searchNode := range *mc {
				for _, out := range searchNode.outgoing {
					if out == k {
						node.incoming[searchKey] = Low
					}
				}
			}
		}
	}
}

func processSignal(s signal, mc moduleConfiguration, q *Queue) {
	node := mc[s.destination]
	switch node.moduleType {
	case Broadcaster:
		for _, out := range node.outgoing {
			q.Enqueue(signal{s.destination, s.level, out})
		}
	case FlipFlop:
		if s.level == Low {
			node.state = node.state.getReverse()
			mc[s.destination] = node
			for _, out := range node.outgoing {
				q.Enqueue(signal{s.destination, node.state, out})
			}
		}
	case Conjunction:
		node.incoming[s.source] = s.level
		mc[s.destination] = node
		allHigh := true
		for _, in := range node.incoming {
			if in == Low {
				allHigh = false
			}
		}
		sendLevel := High
		if allHigh {
			sendLevel = Low
		}
		for _, out := range node.outgoing {
			q.Enqueue(signal{s.destination, sendLevel, out})
		}
	}
}

func pushButton(mc moduleConfiguration) {
	q := make(Queue, 0)
	q.Enqueue(signal{"button", Low, "broadcaster"})
	for {
		s, ok := q.Dequeue()
		if !ok {
			break
		}
		if s.level == Low {
			lowCounter++
		} else {
			highCounter++
		}
		processSignal(s, mc, &q)
	}
}

func puzzle1(mc moduleConfiguration) int {
	lowCounter = 0
	highCounter = 0
	compileIncoming(&mc)
	for i := 0; i < 1000; i++ {
		pushButton(mc)
	}
	return lowCounter * highCounter
}

func puzzle2(mc moduleConfiguration) int {
	compileIncoming(&mc)
	pushCounter := 0
	for {
		clearScreen()
		printConfiguration(mc, []string{"lq", "jn", "mn"}) //, "hd"
		fmt.Print("Pushcounter:", pushCounter, " <enter> - next, <num> - num times, q - quit:")
		time.Sleep(10 * time.Millisecond)
		var command string
		fmt.Scanln(&command)
		if command == "q" {
			break
		} else if command == "" {
			pushButton(mc)
			pushCounter++
		} else {
			count, err := strconv.Atoi(command)
			if err == nil {
				for i := 0; i < count; i++ {
					pushButton(mc)
					pushCounter++
				}
			}
		}
	}
	return pushCounter
}

func main() {
	// mc := loadData("input-test.txt")
	// mc := loadData("input-test2.txt")
	mc := loadData("input.txt")
	result1 := puzzle1(mc)
	fmt.Println("Puzzle 1: ", result1)
	// mc := loadData("input-test.txt")
	// mc := loadData("input-test2.txt")
	mc = loadData("input.txt")
	// Not a solution, just visualization, not sure its possible to solve this puzzle
	// without analyzing the input
	fmt.Println("Puzzle 2: ", puzzle2(mc))

}
