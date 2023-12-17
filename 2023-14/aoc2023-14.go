package main

import "fmt"

func stepWest(data [][]byte) bool {
	isStable := true
	for i := 1; i < len(data[0]); i++ {
		for j := 0; j < len(data); j++ {
			if data[j][i-1] == '.' {
				if data[j][i] == 'O' {
					data[j][i-1] = 'O'
					data[j][i] = '.'
					isStable = false
				}
			}
		}
	}
	return isStable
}

func tiltWest(data [][]byte) {
	for {
		if stepWest(data) {
			break
		}
	}
}

func stepSouth(data [][]byte) bool {
	isStable := true
	for i := len(data) - 2; i >= 0; i-- {
		for j := 0; j < len(data[0]); j++ {
			if data[i+1][j] == '.' {
				if data[i][j] == 'O' {
					data[i+1][j] = 'O'
					data[i][j] = '.'
					isStable = false
				}
			}
		}
	}
	return isStable
}

func tiltSouth(data [][]byte) {
	for {
		if stepSouth(data) {
			break
		}
	}
}

func stepEast(data [][]byte) bool {
	isStable := true
	for i := len(data[0]) - 2; i >= 0; i-- {
		for j := 0; j < len(data); j++ {
			if data[j][i+1] == '.' {
				if data[j][i] == 'O' {
					data[j][i+1] = 'O'
					data[j][i] = '.'
					isStable = false
				}
			}
		}
	}
	return isStable
}

func tiltEast(data [][]byte) {
	for {
		if stepEast(data) {
			break
		}
	}
}

func stepNorth(data [][]byte) bool {
	isStable := true
	for i := 1; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if data[i-1][j] == '.' {
				if data[i][j] == 'O' {
					data[i-1][j] = 'O'
					data[i][j] = '.'
					isStable = false
				}
			}
		}
	}
	return isStable
}

func tiltNorth(data [][]byte) {
	for {
		if stepNorth(data) {
			break
		}
	}
}

func spinCycle(data [][]byte) {
	tiltNorth(data)
	tiltWest(data)
	tiltSouth(data)
	tiltEast(data)
}

func copyData(data [][]byte) [][]byte {
	newData := make([][]byte, len(data))
	for i := 0; i < len(data); i++ {
		newData[i] = make([]byte, len(data[i]))
		copy(newData[i], data[i])
	}
	return newData
}

func compareData(data1 [][]byte, data2 [][]byte) bool {
	for i := 0; i < len(data1); i++ {
		for j := 0; j < len(data1[i]); j++ {
			if data1[i][j] != data2[i][j] {
				return false
			}
		}
	}
	return true
}

func checkCycle(data [][]byte, dataTrail [][][]byte) (int, bool) {
	for index, board := range dataTrail {
		if compareData(data, board) {
			return index, true
		}
	}
	return 0, false
}

func getWeight(data [][]byte) int {
	weight := 0
	for i := 0; i < len(data); i++ {
		lineWeight := 0
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] == 'O' {
				lineWeight++
			}
		}
		weight += lineWeight * (len(data) - i)
	}
	return weight
}

func puzzle1(data [][]byte) int {
	tiltNorth(data)
	sum := getWeight(data)
	return sum
}

func puzzle2(data [][]byte) int {
	dataTrail := make([][][]byte, 0)
	newBoard := copyData(data)
	dataTrail = append(dataTrail, newBoard)
	cycleCounter := 0
	index := 0
	found := false
	for {
		spinCycle(data)
		cycleCounter++
		index, found = checkCycle(data, dataTrail)
		if found {
			break
		}
		newBoard := copyData(data)
		dataTrail = append(dataTrail, newBoard)
		if cycleCounter > 100000 {
			fmt.Println("Too many cycles")
			break
		}
	}
	remining_cycles := (1000000000 - cycleCounter) % (cycleCounter - index)
	for i := 0; i < remining_cycles; i++ {
		spinCycle(data)
	}
	return getWeight(data)
}

func main() {
	// board, _ := loadData("input-test.txt")
	board, _ := loadData("input.txt")
	fmt.Println("Puzzle 1: ", puzzle1(board))
	fmt.Println("Puzzle 2: ", puzzle2(board))
}
