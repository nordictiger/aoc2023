package main

import (
	"fmt"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

func getFirstHeatLoss(data [][]int) int {
	minHeatLoss := 0
	for x := 0; x < len(data)-1; x++ {
		minHeatLoss += data[x+1][x]
		minHeatLoss += data[x+1][x+1]
	}
	return minHeatLoss
}

func findPath(data [][]int, x, y int, direction Direction, straightSteps int, sum int, minHeatLoss *int, path *Stack) {
	if x < 0 || y < 0 || x >= len(data) || y >= len(data[0]) {
		return
	}
	if data[x][y] < 0 {
		return
	}
	sum += data[x][y]
	if sum > *minHeatLoss {
		return
	}
	if x == len(data)-1 && y == len(data[0])-1 {
		fmt.Println("Found path: ", sum)
		// printPath(path)
		if sum < *minHeatLoss {
			*minHeatLoss = sum
		}
		return
	}
	data[x][y] -= 100
	// path.Push(Point{x, y, straightSteps})
	switch direction {
	case Up:
		if straightSteps < 3 {
			findPath(data, x-1, y, Up, straightSteps+1, sum, minHeatLoss, path)
		}
		findPath(data, x, y-1, Left, 1, sum, minHeatLoss, path)
		findPath(data, x, y+1, Right, 1, sum, minHeatLoss, path)
	case Down:
		if straightSteps < 3 {
			findPath(data, x+1, y, Down, straightSteps+1, sum, minHeatLoss, path)
		}
		findPath(data, x, y-1, Left, 1, sum, minHeatLoss, path)
		findPath(data, x, y+1, Right, 1, sum, minHeatLoss, path)
	case Left:
		if straightSteps < 3 {
			findPath(data, x, y-1, Left, straightSteps+1, sum, minHeatLoss, path)
		}
		findPath(data, x-1, y, Up, 1, sum, minHeatLoss, path)
		findPath(data, x+1, y, Down, 1, sum, minHeatLoss, path)
	case Right:
		if straightSteps < 3 {
			findPath(data, x, y+1, Right, straightSteps+1, sum, minHeatLoss, path)
		}
		findPath(data, x-1, y, Up, 1, sum, minHeatLoss, path)
		findPath(data, x+1, y, Down, 1, sum, minHeatLoss, path)
	}
	data[x][y] += 100
	// path.Pop()
}

func puzzle1(data [][]int) int {
	min := getFirstHeatLoss(data)
	path := Stack{}
	// min = 104
	fmt.Println("First proxy: ", min)
	// get first proxy number
	minDown := min
	findPath(data, 1, 0, Down, 1, 0, &minDown, &path)
	minRight := minDown
	// minRight := min
	// findPath(data, 0, 1, Right, 1, 0, &minRight, &path)

	if minDown < minRight {
		min = minDown
	} else {
		min = minRight
	}
	// printData(data)
	return min
}

func puzzle2(data [][]int) int {
	sum := 0
	return sum
}

func main() {
	// cityMap, _ := loadData("input-test2.txt")
	// cityMap, _ := loadData("input-test.txt")
	cityMap, _ := loadData("input.txt")
	fmt.Println("Puzzle 1: ", puzzle1(cityMap))
	fmt.Println("Puzzle 2: ", puzzle2(cityMap))
}
