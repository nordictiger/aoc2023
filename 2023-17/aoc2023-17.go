package main

import (
	"fmt"
	"math"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Point struct {
	x, y int
}

type CityBlock struct {
	value       int
	visited     bool
	minHeatLoss int
}

func createSolvingMap(data [][]int) [][]CityBlock {
	solvingMap := make([][]CityBlock, len(data))
	for x := 0; x < len(data); x++ {
		solvingMap[x] = make([]CityBlock, len(data[0]))
		for y := 0; y < len(data[0]); y++ {
			solvingMap[x][y] = CityBlock{data[x][y], false, math.MaxInt}
		}
	}
	return solvingMap
}

func getSmallestBlock(queue map[Point]int, solvingMap [][]CityBlock) Point {
	min := math.MaxInt
	var minPoint Point
	for point, heatLoss := range queue {
		if heatLoss < min {
			min = heatLoss
			minPoint = point
		} else if heatLoss == min {
			if point.y > minPoint.y {
				min = heatLoss
				minPoint = point
			} else if point.y == minPoint.y {
				if point.x > minPoint.x {
					min = heatLoss
					minPoint = point
				}
			}
		}
	}
	delete(queue, minPoint)
	return minPoint
}

func processBlock(currentPoint Point, nextPoint Point, solvingMap [][]CityBlock, queue *map[Point]int) {
	nextBlock := solvingMap[nextPoint.x][nextPoint.y]
	newMinHeatLoss := nextBlock.value + solvingMap[currentPoint.x][currentPoint.y].minHeatLoss
	if nextBlock.minHeatLoss > newMinHeatLoss {
		solvingMap[nextPoint.x][nextPoint.y].minHeatLoss = newMinHeatLoss
		(*queue)[nextPoint] = newMinHeatLoss
	}
}

func resolveUp(currentPoint Point, solvingMap [][]CityBlock, queue *map[Point]int) {
	nextX := currentPoint.x - 1
	nextY := currentPoint.y
	if nextX >= 0 && !solvingMap[nextX][nextY].visited {
		// Moving Up is valid move
		processBlock(currentPoint, Point{nextX, nextY}, solvingMap, queue)
	}
}

func resolveRight(currentPoint Point, solvingMap [][]CityBlock, queue *map[Point]int) {
	nextX := currentPoint.x
	nextY := currentPoint.y + 1

	if nextY < len(solvingMap[0]) && !solvingMap[nextX][nextY].visited {
		// Moving Right is valid move
		processBlock(currentPoint, Point{nextX, nextY}, solvingMap, queue)
	}
}

func resolveDown(currentPoint Point, solvingMap [][]CityBlock, queue *map[Point]int) {
	nextX := currentPoint.x + 1
	nextY := currentPoint.y
	if nextX < len(solvingMap) && !solvingMap[nextX][nextY].visited {
		// Moving Down is valid move
		processBlock(currentPoint, Point{nextX, nextY}, solvingMap, queue)
	}
}

func resolveLeft(currentPoint Point, solvingMap [][]CityBlock, queue *map[Point]int) {
	nextX := currentPoint.x
	nextY := currentPoint.y - 1
	if nextY >= 0 && !solvingMap[nextX][nextY].visited {
		// Moving Left is valid move
		processBlock(currentPoint, Point{nextX, nextY}, solvingMap, queue)
	}
}

func calculateDijkstra(solvingMap [][]CityBlock) {
	queue := make(map[Point]int)
	/*
		solvingMap[0][0].value = 0
		solvingMap[0][0].minHeatLoss = 0
		queue[Point{0, 0}] = 0
	*/
	solvingMap[len(solvingMap)-1][len(solvingMap[0])-1].value = 0
	solvingMap[len(solvingMap)-1][len(solvingMap[0])-1].minHeatLoss = 0
	queue[Point{len(solvingMap) - 1, len(solvingMap[0]) - 1}] = 0

	for len(queue) > 0 {
		nextBlock := getSmallestBlock(queue, solvingMap)
		solvingMap[nextBlock.x][nextBlock.y].visited = true
		/*
			if nextBlock.x == len(solvingMap)-1 && nextBlock.y == len(solvingMap[0])-1 {
				break
			}
		*/
		resolveRight(nextBlock, solvingMap, &queue)
		resolveDown(nextBlock, solvingMap, &queue)
		resolveUp(nextBlock, solvingMap, &queue)
		resolveLeft(nextBlock, solvingMap, &queue)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func puzzle1(data [][]int) int {
	solvingMap := createSolvingMap(data)
	calculateDijkstra(solvingMap)
	for x := 0; x < len(solvingMap); x++ {
		for y := 0; y < len(solvingMap[0]); y++ {
			solvingMap[x][y].visited = false
		}
	}
	solvingMap[len(solvingMap)-1][len(solvingMap[0])-1].value = data[len(data)-1][len(data[0])-1]
	solvingMap[0][0].value = 0
	solvingMap[0][0].minHeatLoss = 0

	printData(data)
	printMap(solvingMap)
	// printMapPath(solvingMap, path)
	path := Stack{}

	min := max(solvingMap[1][1].minHeatLoss, solvingMap[0][1].minHeatLoss)
	// get first proxy number
	path = Stack{}
	found := false

	for !found {
		findPath(solvingMap, 1, 0, Down, 1, 0, &min, &found, &path)
		fmt.Println("min: ", min, "found: ", found)
		min++
	}
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
