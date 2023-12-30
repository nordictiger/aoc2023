package main

import (
	"fmt"
)

func puzzle1(data [][]byte) int {
	sum := 0
	startX := -1
	startY := -1
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			fmt.Printf("%c", data[i][j])
			if data[i][j] == 'S' {
				startX = i
				startY = j
			}
		}
		fmt.Println()
	}
	fmt.Println("Start: ", startX, startY)
	return sum
}

func puzzle2(data [][]byte) int {
	sum := 0
	return sum
}

func main() {
	// garden := loadData("input-test.txt")
	garden := loadData("input.txt")
	fmt.Println("Puzzle 1: ", puzzle1(garden))
	fmt.Println("Puzzle 2: ", puzzle2(garden))
}
