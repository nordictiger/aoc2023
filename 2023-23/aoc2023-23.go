package main

import (
	"fmt"
	"strings"
)

type direction int

const (
	up direction = iota
	down
	left
	right
)

type point struct {
	x int
	y int
}

func tracePath(hikingMap [][]byte, start point, end point, sum int, dir direction, longestPath *int, slippery bool) {
	if start.x < 0 || start.x >= len(hikingMap) {
		return
	}
	if start.y < 0 || start.y >= len(hikingMap[0]) {
		return
	}

	if hikingMap[start.x][start.y] == '#' {
		return
	}
	if hikingMap[start.x][start.y] == 'O' {
		return
	}
	if slippery {
		if hikingMap[start.x][start.y] == '>' && dir == left {
			return
		}
		if hikingMap[start.x][start.y] == '<' && dir == right {
			return
		}
		if hikingMap[start.x][start.y] == '^' && dir == down {
			return
		}
		if hikingMap[start.x][start.y] == 'v' && dir == up {
			return
		}
	}

	if start.x == end.x && start.y == end.y {
		sum++
		if sum > *longestPath {
			fmt.Println("New longest path: ", sum)
			*longestPath = sum
		}
		return
	}

	if strings.Contains("<>v^.", string(hikingMap[start.x][start.y])) {
		sum++
	}

	mapSymbol := hikingMap[start.x][start.y]
	hikingMap[start.x][start.y] = 'O'
	tracePath(hikingMap, point{start.x - 1, start.y}, end, sum, up, longestPath, slippery)
	tracePath(hikingMap, point{start.x + 1, start.y}, end, sum, down, longestPath, slippery)
	tracePath(hikingMap, point{start.x, start.y - 1}, end, sum, left, longestPath, slippery)
	tracePath(hikingMap, point{start.x, start.y + 1}, end, sum, right, longestPath, slippery)
	hikingMap[start.x][start.y] = mapSymbol
}

func puzzle1(hikingMap [][]byte) int {
	sum := 0
	tracePath(hikingMap, point{0, 1}, point{len(hikingMap) - 1, len(hikingMap[0]) - 2}, 0, down, &sum, true)
	return sum - 1
}

func puzzle2(hikingMap [][]byte) int {
	sum := 0
	tracePath(hikingMap, point{0, 1}, point{len(hikingMap) - 1, len(hikingMap[0]) - 2}, 0, down, &sum, false)
	return sum - 1
}

func main() {
	hikingMap := loadData("input.txt")
	// hikingMap := loadData("input-test.txt")
	fmt.Println("Puzzle 1: ", puzzle1(hikingMap))
	fmt.Println("Puzzle 2: ", puzzle2(hikingMap))
}
