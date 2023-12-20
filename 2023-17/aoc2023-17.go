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

type DirectionCounter struct {
	direction   Direction
	count       int
	minHeatLoss int
}

type CityBlock struct {
	value            int
	visited          bool
	minHeatLoss      int
	minHeatLossPoint Point
}

func createSolvingMap(data [][]int) [][]CityBlock {
	solvingMap := make([][]CityBlock, len(data))
	for x := 0; x < len(data); x++ {
		solvingMap[x] = make([]CityBlock, len(data[0]))
		for y := 0; y < len(data[0]); y++ {
			solvingMap[x][y] = CityBlock{data[x][y], false, math.MaxInt, Point{-1, -1}}
		}
	}
	return solvingMap
}

func getSmallestBlock(queue map[Point]DirectionCounter) (Point, DirectionCounter) {
	min := math.MaxInt
	maxX := 0
	maxY := 0
	var minPoint Point
	var minDirection DirectionCounter
	for point, dc := range queue {
		if dc.minHeatLoss < min {
			min = dc.minHeatLoss
			maxX = point.x
			maxY = point.y
			minPoint = point
			minDirection = dc
		} else if dc.minHeatLoss == min {
			if point.y > maxY {
				min = dc.minHeatLoss
				maxX = point.x
				maxY = point.y
				minPoint = point
				minDirection = dc
			} else if point.y == maxY {
				if point.x > maxX {
					min = dc.minHeatLoss
					maxY = point.y
					maxX = point.x
					minPoint = point
					minDirection = dc
				}
			}
		}
	}
	delete(queue, minPoint)
	return minPoint, minDirection
}

func processBlock(currentPoint Point, nextX int, nextY int, currentPointDirection DirectionCounter, dir Direction, solvingMap [][]CityBlock, queue *map[Point]DirectionCounter) {
	stepCounter := 1
	if currentPointDirection.direction == dir {
		stepCounter = currentPointDirection.count + 1
	}
	stepCounter = 1
	nextBlock := solvingMap[nextX][nextY]
	newMinHeatLoss := nextBlock.value + currentPointDirection.minHeatLoss
	if nextBlock.minHeatLoss > newMinHeatLoss {
		solvingMap[nextX][nextY].minHeatLoss = newMinHeatLoss
		solvingMap[nextX][nextY].minHeatLossPoint = Point{currentPoint.x, currentPoint.y}
		(*queue)[Point{nextX, nextY}] = DirectionCounter{dir, stepCounter, newMinHeatLoss}
	}
}

func resolveUp(currentPoint Point, currentPointDirection DirectionCounter, solvingMap [][]CityBlock, queue *map[Point]DirectionCounter) {
	nextX := currentPoint.x - 1
	nextY := currentPoint.y
	if nextX >= 0 && !solvingMap[nextX][nextY].visited {
		// Moving Up is valid move
		processBlock(currentPoint, nextX, nextY, currentPointDirection, Up, solvingMap, queue)
	}
}

func resolveRight(currentPoint Point, currentPointDirection DirectionCounter, solvingMap [][]CityBlock, queue *map[Point]DirectionCounter) {
	nextX := currentPoint.x
	nextY := currentPoint.y + 1

	if nextY < len(solvingMap[0]) && !solvingMap[nextX][nextY].visited {
		// Moving Right is valid move
		processBlock(currentPoint, nextX, nextY, currentPointDirection, Right, solvingMap, queue)
	}
}

func resolveDown(currentPoint Point, currentPointDirection DirectionCounter, solvingMap [][]CityBlock, queue *map[Point]DirectionCounter) {
	nextX := currentPoint.x + 1
	nextY := currentPoint.y
	if nextX < len(solvingMap) && !solvingMap[nextX][nextY].visited {
		// Moving Down is valid move
		processBlock(currentPoint, nextX, nextY, currentPointDirection, Down, solvingMap, queue)
	}
}

func resolveLeft(currentPoint Point, currentPointDirection DirectionCounter, solvingMap [][]CityBlock, queue *map[Point]DirectionCounter) {
	nextX := currentPoint.x
	nextY := currentPoint.y - 1
	if nextY >= 0 && !solvingMap[nextX][nextY].visited {
		// Moving Left is valid move
		processBlock(currentPoint, nextX, nextY, currentPointDirection, Left, solvingMap, queue)
	}
}

func calculateDijkstra(solvingMap [][]CityBlock) {
	queue := make(map[Point]DirectionCounter)

	solvingMap[0][0].value = 0
	solvingMap[0][0].minHeatLoss = 0
	queue[Point{0, 0}] = DirectionCounter{Right, 0, 0}

	for len(queue) > 0 {
		nba, nbd := getSmallestBlock(queue)
		solvingMap[nba.x][nba.y].visited = true
		/*
			if nba.x == len(solvingMap)-1 && nba.y == len(solvingMap[0])-1 {
				break
			}
		*/
		switch nbd.direction {
		case Up:
			resolveRight(nba, nbd, solvingMap, &queue)
			if nbd.count < 3 {
				resolveUp(nba, nbd, solvingMap, &queue)
			}
			resolveLeft(nba, nbd, solvingMap, &queue)

		case Down:
			resolveRight(nba, nbd, solvingMap, &queue)
			if nbd.count < 3 {
				resolveDown(nba, nbd, solvingMap, &queue)
			}
			resolveLeft(nba, nbd, solvingMap, &queue)
		case Left:
			resolveDown(nba, nbd, solvingMap, &queue)
			resolveUp(nba, nbd, solvingMap, &queue)
			if nbd.count < 3 {
				resolveLeft(nba, nbd, solvingMap, &queue)
			}
		case Right:
			if nbd.count < 3 {
				resolveRight(nba, nbd, solvingMap, &queue)
			}
			resolveDown(nba, nbd, solvingMap, &queue)
			resolveUp(nba, nbd, solvingMap, &queue)
		}
	}
}
func tracePath(solvingMap [][]CityBlock) []Point {
	path := make([]Point, 0)
	x := len(solvingMap) - 1
	y := len(solvingMap[0]) - 1
	path = append(path, Point{x, y})
	for x != 0 || y != 0 {
		p := solvingMap[x][y].minHeatLossPoint
		x = p.x
		y = p.y
		path = append(path, Point{x, y})
	}
	return path
}

func puzzle1(data [][]int) int {
	min := 0
	solvingMap := createSolvingMap(data)
	calculateDijkstra(solvingMap)
	printData(data)
	printMap(solvingMap)
	path := tracePath(solvingMap)
	printMapPath(solvingMap, path)
	return min
}

func puzzle2(data [][]int) int {
	sum := 0
	return sum
}

func main() {
	// cityMap, _ := loadData("input-test2.txt")
	cityMap, _ := loadData("input-test.txt")
	// cityMap, _ := loadData("input.txt")
	fmt.Println("Puzzle 1: ", puzzle1(cityMap))
	fmt.Println("Puzzle 2: ", puzzle2(cityMap))
}
