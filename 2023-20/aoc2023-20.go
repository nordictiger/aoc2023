package main

import (
	"crypto/sha256"
	"fmt"
)

var counter int

func hash(str string) string {
	hasher := sha256.New()
	hasher.Write([]byte(str))
	hashBytes := hasher.Sum(nil)
	hashString := fmt.Sprintf("%x", hashBytes)
	return hashString
}

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
		processSignal(s, mc, &q)
		// fmt.Println(s)
		counter++
	}
}

func puzzle1(mc moduleConfiguration) int {
	sum := 0
	counter = 0
	compileIncoming(&mc)
	initialStateHash := hash(fmt.Sprintf("%v", mc))
	newStateHash := ""
	fmt.Println("Initial state:", initialStateHash)
	pushCounter := 0
	for initialStateHash != newStateHash {
		pushButton(mc)
		pushCounter++
		newStateHash = hash(fmt.Sprintf("%v", mc))
		fmt.Println("New state:", newStateHash, "Pushcounter:", pushCounter, "Counter:", counter)
		if pushCounter > 100 {
			break
		}
	}
	return sum
}

func puzzle2() int {
	sum := 0
	return sum
}

func main() {
	//
	// mc := loadData("input-test.txt")
	// mc := loadData("input-test2.txt")
	mc := loadData("input.txt")
	fmt.Println("Puzzle 1: ", puzzle1(mc))
	fmt.Println("Puzzle 2: ", puzzle2())
}
