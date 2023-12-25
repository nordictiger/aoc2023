package main

import (
	"fmt"
	"strconv"
)

type instruction struct {
	direction string
	distance  int
	color     string
}

type point struct {
	x int
	y int
}

type turn int

const (
	Left turn = iota
	Right
)

type edge struct {
	start       point
	turnToStart turn
	end         point
	turnFromEnd turn
}

func getTurn(currentDirection, nextDirection string) turn {
	switch currentDirection {
	case "U":
		switch nextDirection {
		case "R":
			return Right
		case "L":
			return Left
		}
	case "R":
		switch nextDirection {
		case "D":
			return Right
		case "U":
			return Left
		}
	case "D":
		switch nextDirection {
		case "L":
			return Right
		case "R":
			return Left
		}
	case "L":
		switch nextDirection {
		case "U":
			return Right
		case "D":
			return Left
		}
	}
	panic("Possible turn not found.")
}

func getShift(lastTurn, nextTurn turn) int {
	// This probably depends on the winding order of the polygon
	switch lastTurn {
	case Left:
		switch nextTurn {
		case Left:
			return -1
		case Right:
			return 0
		}
	case Right:
		switch nextTurn {
		case Left:
			return 0
		case Right:
			return 1
		}
	}
	panic("Possible shift not found.")
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getEdges(data []instruction) []edge {
	edges := make([]edge, 0)
	startingPoint := point{0, 0}
	lastTurn := getTurn(data[len(data)-1].direction, data[0].direction)
	var nextTurn turn
	for i := 0; i < len(data); i++ {
		aciveInstruction := data[i]
		j := i + 1
		if i == len(data)-1 {
			j = 0
		}
		nextInstruction := data[j]
		nextTurn = getTurn(aciveInstruction.direction, nextInstruction.direction)

		var newEdge edge
		shift := getShift(lastTurn, nextTurn)
		newEdge.start = startingPoint
		newEdge.turnToStart = lastTurn
		newEdge.end = startingPoint
		newEdge.turnFromEnd = nextTurn
		switch aciveInstruction.direction {
		case "U":
			newEdge.end.y += (aciveInstruction.distance + shift)
		case "R":
			newEdge.end.x += (aciveInstruction.distance + shift)
		case "D":
			newEdge.end.y -= (aciveInstruction.distance + shift)
		case "L":
			newEdge.end.x -= (aciveInstruction.distance + shift)
		}
		edges = append(edges, newEdge)
		startingPoint = newEdge.end
		lastTurn = nextTurn
	}
	return edges
}

func signedArea(edges []edge) int {
	area := 0
	n := len(edges)
	for i := 0; i < n-1; i++ {
		sa := (edges[i].end.x * edges[i+1].end.y) - (edges[i+1].end.x * edges[i].end.y)
		area += sa
	}
	return area / 2
}

func puzzle1(diggingInstructions []instruction) int {
	edges := getEdges(diggingInstructions)
	signedArea := signedArea(edges)
	fmt.Println("signedArea", signedArea)
	return abs(signedArea)
}

func getDistance(color string) int {
	hexString := color[1 : len(color)-1]
	decimalNumber, err := strconv.ParseInt(hexString, 16, 64)
	if err != nil {
		fmt.Println("Error:", err)
		panic("Error converting hexadecimal string to decimal number.")
	} else {
		return int(decimalNumber)
	}
}

func getDirection(color string) string {
	switch color[len(color)-1:] {
	case "0":
		return "R"
	case "1":
		return "D"
	case "2":
		return "L"
	case "3":
		return "U"
	}
	panic("New instruction not found.")
}

func reprocessInstructions(diggingInstructions []instruction) []instruction {
	newInstructions := make([]instruction, 0)
	for _, i := range diggingInstructions {
		newInstructions = append(newInstructions, instruction{getDirection(i.color), getDistance(i.color), i.color})
	}
	return newInstructions
}
func puzzle2(diggingInstructions []instruction) int {
	newInstructions := reprocessInstructions(diggingInstructions)
	edges := getEdges(newInstructions)
	signedArea := signedArea(edges)
	fmt.Println("signedArea", signedArea)
	return abs(signedArea)
}

func main() {
	// diggingInstructions, _ := loadData("input-test.txt")
	diggingInstructions, _ := loadData("input.txt")
	fmt.Println("Puzzle 1: ", puzzle1(diggingInstructions))
	fmt.Println("Puzzle 2: ", puzzle2(diggingInstructions))

}
