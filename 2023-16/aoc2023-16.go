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

type Tile struct {
	visitingRayDirection [4]bool
}

type Position struct {
	x         int
	y         int
	direction Direction
}

func getSolvingField(field [][]byte) [][]Tile {
	data := make([][]Tile, 0)
	for _, line := range field {
		newLine := make([]Tile, len(line))
		data = append(data, newLine)
	}
	return data
}

func tracePath(field [][]byte, solvingField [][]Tile, positionQueue *Stack, position Position) {
	for {
		if position.x < 0 || position.x > len(field)-1 || position.y < 0 || position.y > len(field[position.x])-1 {
			// out of the field
			return
		}
		// check for loop
		if solvingField[position.x][position.y].visitingRayDirection[position.direction] {
			return
		}
		// mark the tile as visited
		solvingField[position.x][position.y].visitingRayDirection[position.direction] = true
		// figure out the next action
		switch field[position.x][position.y] {
		case '.':
			// continue in the same direction
		case '/':
			switch position.direction {
			case Up:
				position.direction = Right
			case Down:
				position.direction = Left
			case Left:
				position.direction = Down
			case Right:
				position.direction = Up
			}
		case '\\':
			switch position.direction {
			case Up:
				position.direction = Left
			case Down:
				position.direction = Right
			case Left:
				position.direction = Up
			case Right:
				position.direction = Down
			}
		case '|':
			if position.direction == Left || position.direction == Right {
				solvingField[position.x][position.y].visitingRayDirection[position.direction] = true
				positionQueue.Push(Position{position.x - 1, position.y, Up})
				positionQueue.Push(Position{position.x + 1, position.y, Down})
				return
			}
		case '-':
			if position.direction == Up || position.direction == Down {
				solvingField[position.x][position.y].visitingRayDirection[position.direction] = true
				positionQueue.Push(Position{position.x, position.y - 1, Left})
				positionQueue.Push(Position{position.x, position.y + 1, Right})
				return
			}
		default:
			panic("Error: unknown tile type: " + string(field[position.y][position.x]))
		}

		// figure out the next position
		switch position.direction {
		case Up:
			position.x--
		case Down:
			position.x++
		case Left:
			position.y--
		case Right:
			position.y++
		}
	}
}

func solvePuzzle(field [][]byte, solvingField [][]Tile, positionQueue *Stack) {
	for {
		position, ok := positionQueue.Pop()
		if !ok {
			break
		}
		tracePath(field, solvingField, positionQueue, position)
	}
}

func getEnergizedTileCount(solvingField [][]Tile) int {
	sum := 0
	for _, line := range solvingField {
		for _, t := range line {
			if t.visitingRayDirection[0] || t.visitingRayDirection[1] || t.visitingRayDirection[2] || t.visitingRayDirection[3] {
				sum++
			}
		}
	}
	return sum
}

func puzzle1(data [][]byte) int {
	positionQueue := Stack{}
	solvingField := getSolvingField(data)

	startingPosition := Position{0, 0, Right}
	positionQueue.Push(startingPosition)

	solvePuzzle(data, solvingField, &positionQueue)
	sum := getEnergizedTileCount(solvingField)

	return sum
}

func puzzle2(data [][]byte) int {
	max := 0
	for i := 0; i < len(data[0]); i++ {
		// check from the top
		positionQueue := Stack{}
		solvingField := getSolvingField(data)
		startingPosition := Position{0, i, Down}
		positionQueue.Push(startingPosition)
		solvePuzzle(data, solvingField, &positionQueue)
		count := getEnergizedTileCount(solvingField)
		if count > max {
			max = count
		}
	}
	for i := 0; i < len(data[0]); i++ {
		// check from the bottom
		positionQueue := Stack{}
		solvingField := getSolvingField(data)
		startingPosition := Position{len(data) - 1, i, Up}
		positionQueue.Push(startingPosition)
		solvePuzzle(data, solvingField, &positionQueue)
		count := getEnergizedTileCount(solvingField)
		if count > max {
			max = count
		}
	}
	for i := 0; i < len(data); i++ {
		// check from the left
		positionQueue := Stack{}
		solvingField := getSolvingField(data)
		startingPosition := Position{i, 0, Right}
		positionQueue.Push(startingPosition)
		solvePuzzle(data, solvingField, &positionQueue)
		count := getEnergizedTileCount(solvingField)
		if count > max {
			max = count
		}
	}
	for i := 0; i < len(data); i++ {
		// check from the right
		positionQueue := Stack{}
		solvingField := getSolvingField(data)
		startingPosition := Position{i, len(data[0]) - 1, Left}
		positionQueue.Push(startingPosition)
		solvePuzzle(data, solvingField, &positionQueue)
		count := getEnergizedTileCount(solvingField)
		if count > max {
			max = count
		}
	}
	return max
}

func main() {
	// contraption, _ := loadData("input-test.txt")
	contraption, _ := loadData("input.txt")
	fmt.Println("Puzzle 1: ", puzzle1(contraption))
	fmt.Println("Puzzle 2: ", puzzle2(contraption))
}
