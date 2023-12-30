package main

import (
	"fmt"
)

func copyGarden(garden [][]byte) [][]byte {
	newGarden := make([][]byte, len(garden))
	for i := 0; i < len(garden); i++ {
		newGarden[i] = make([]byte, len(garden[i]))
		copy(newGarden[i], garden[i])
	}
	return newGarden
}

func getStart(garden [][]byte) (int, int) {
	for i := 0; i < len(garden); i++ {
		for j := 0; j < len(garden[i]); j++ {
			if garden[i][j] == 'S' {
				return i, j
			}
		}
	}
	return -1, -1
}

func step(steps [][]byte, garden [][]byte) [][]byte {
	newSteps := copyGarden(garden)
	for i := 0; i < len(steps); i++ {
		for j := 0; j < len(steps[i]); j++ {
			if steps[i][j] == 'O' || steps[i][j] == 'S' {
				updateSteps(newSteps, i, j, garden)
			}
		}
	}
	return newSteps
}

func updateSteps(steps [][]byte, x, y int, garden [][]byte) {
	if x > 0 && garden[x-1][y] == '.' {
		steps[x-1][y] = 'O'
	}
	if x < len(garden)-1 && garden[x+1][y] == '.' {
		steps[x+1][y] = 'O'
	}
	if y > 0 && garden[x][y-1] == '.' {
		steps[x][y-1] = 'O'
	}
	if y < len(garden[x])-1 && garden[x][y+1] == '.' {
		steps[x][y+1] = 'O'
	}
}

func countPlots(garden [][]byte) int {
	count := 0
	for i := 0; i < len(garden); i++ {
		for j := 0; j < len(garden[i]); j++ {
			if garden[i][j] == 'O' {
				count++
			}
		}
	}
	return count
}

func puzzle1(data [][]byte, reminingSteps int) int {
	startX := -1
	startY := -1
	startX, startY = getStart(data)
	fmt.Println("Start: ", startX, startY)
	steps := copyGarden(data)
	printGarden(steps)
	data[startX][startY] = '.'
	for i := 0; i < reminingSteps; i++ {
		steps = step(steps, data)
	}
	fmt.Println()
	printGarden(steps)
	sum := countPlots(steps)
	return sum
}

func puzzle2(data [][]byte) int {
	sum := 0
	return sum
}

func main() {
	// garden := loadData("input-test.txt")
	// fmt.Println("Puzzle 1: ", puzzle1(garden, 6))
	garden := loadData("input.txt")
	fmt.Println("Puzzle 1: ", puzzle1(garden, 64))
	// 3660
	fmt.Println("Puzzle 2: ", puzzle2(garden))
}
