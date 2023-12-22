package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
	Undefined
)

func (d Direction) opposite() Direction {
	switch d {
	case Up:
		return Down
	case Down:
		return Up
	case Left:
		return Right
	case Right:
		return Left
	default:
		return Undefined
	}
}

type Point struct {
	x, y int
}

func (p Point) getNextPoint(direction Direction, stepSize int) Point {
	switch direction {
	case Right:
		return Point{p.x, p.y + stepSize}
	case Down:
		return Point{p.x + stepSize, p.y}
	case Up:
		return Point{p.x - stepSize, p.y}
	case Left:
		return Point{p.x, p.y - stepSize}
	default:
		panic("Can not move to unknown direction")
	}
}

type Block struct {
	address          Point
	directionToBlock Direction
	stepsInDirection int
}

type HeapBlock struct {
	Block
	minHeatLoss int
	index       int // heap needs this
}

func processBlock(currentBlock HeapBlock, exploredMap *map[Block]int, queue *PriorityQueue, data [][]int, direction Direction, minSteps int, maxSteps int) {

	currentDirection := currentBlock.directionToBlock
	if currentDirection == direction.opposite() {
		// can not go back
		return
	}

	stepsInDirection := currentBlock.stepsInDirection
	var nextPoint Point
	if currentDirection == direction {
		if stepsInDirection == maxSteps {
			return
		} else {
			stepsInDirection++
		}
		nextPoint = currentBlock.address.getNextPoint(direction, 1)
	} else {
		stepsInDirection = minSteps
		nextPoint = currentBlock.address.getNextPoint(direction, minSteps)
	}

	if nextPoint.x < 0 || nextPoint.x >= len(data) || nextPoint.y < 0 || nextPoint.y >= len(data[0]) {
		// can not go outside of map
		return
	}
	nextBlock := Block{nextPoint, direction, stepsInDirection}
	_, ok := (*exploredMap)[nextBlock]
	if ok {
		// already explored
		return
	}

	newMinHeatLoss := getNewHeatLoss(currentBlock, direction, nextPoint, data)
	newHeapBlock := HeapBlock{nextBlock, newMinHeatLoss, 0}
	// Push takes care of upsert logic
	heap.Push(queue, &newHeapBlock)
}

func calculateDijkstra(data [][]int, minSteps int, maxSteps int) map[Block]int {
	exploredMap := make(map[Block]int)

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &HeapBlock{Block{Point{0, 0}, Undefined, 0}, 0, 0})
	data[0][0] = 0

	for len(pq) > 0 {
		/*
			if len(exploredMap)%10000 == 0 {
				fmt.Println("Explored map size:", len(exploredMap))
				fmt.Println("Heap size:", len(pq))
			}
			if len(pq)%10000 == 0 {
				fmt.Println("Explored map size:", len(exploredMap))
				fmt.Println("Heap size:", len(pq))
			}
		*/
		nextBlock := heap.Pop(&pq).(*HeapBlock)
		exploredMap[nextBlock.Block] = nextBlock.minHeatLoss

		processBlock(*nextBlock, &exploredMap, &pq, data, Right, minSteps, maxSteps)
		processBlock(*nextBlock, &exploredMap, &pq, data, Down, minSteps, maxSteps)
		processBlock(*nextBlock, &exploredMap, &pq, data, Up, minSteps, maxSteps)
		processBlock(*nextBlock, &exploredMap, &pq, data, Left, minSteps, maxSteps)
	}

	return exploredMap
}

func getNewHeatLoss(currentBlock HeapBlock, direction Direction, nextPoint Point, data [][]int) int {
	newHeatLoss := currentBlock.minHeatLoss
	switch direction {
	case Up:
		for x := currentBlock.address.x - 1; x >= nextPoint.x; x-- {
			newHeatLoss += data[x][nextPoint.y]
		}
	case Down:
		for x := currentBlock.address.x + 1; x <= nextPoint.x; x++ {
			newHeatLoss += data[x][nextPoint.y]
		}
	case Left:
		for y := currentBlock.address.y - 1; y >= nextPoint.y; y-- {
			newHeatLoss += data[nextPoint.x][y]
		}
	case Right:
		for y := currentBlock.address.y + 1; y <= nextPoint.y; y++ {
			newHeatLoss += data[nextPoint.x][y]
		}
	}
	return newHeatLoss
}

func puzzle1(data [][]int) int {
	exploredMap := calculateDijkstra(data, 1, 3)
	minHeatLoss := math.MaxInt
	for block, heatLoss := range exploredMap {
		if block.address.x == len(data)-1 && block.address.y == len(data[0])-1 {
			if heatLoss < minHeatLoss {
				minHeatLoss = heatLoss
			}
		}
	}
	return minHeatLoss
}

func puzzle2(data [][]int) int {
	exploredMap := calculateDijkstra(data, 4, 10)
	minHeatLoss := math.MaxInt
	for block, heatLoss := range exploredMap {
		if block.address.x == len(data)-1 && block.address.y == len(data[0])-1 {
			if heatLoss < minHeatLoss {
				minHeatLoss = heatLoss
			}
		}
	}
	return minHeatLoss
}

func main() {
	// cityMap, _ := loadData("input-test.txt")
	cityMap, _ := loadData("input.txt")
	fmt.Println("Puzzle 1: ", puzzle1(cityMap))
	fmt.Println("Puzzle 2: ", puzzle2(cityMap))
}
