package main

import (
	"fmt"
	"sort"
)

type PointDirection struct {
	x, y      int
	direction Direction
	value     int
}

func getNextSteps(solvingMap [][]CityBlock, x, y int) []PointDirection {
	var nextSteps []PointDirection
	if x > 0 {
		nextSteps = append(nextSteps, PointDirection{x - 1, y, Up, solvingMap[x-1][y].value})
	}
	if x < len(solvingMap)-1 {
		nextSteps = append(nextSteps, PointDirection{x + 1, y, Down, solvingMap[x+1][y].value})
	}
	if y > 0 {
		nextSteps = append(nextSteps, PointDirection{x, y - 1, Left, solvingMap[x][y-1].value})
	}
	if y < len(solvingMap[0])-1 {
		nextSteps = append(nextSteps, PointDirection{x, y + 1, Right, solvingMap[x][y+1].value})
	}
	sort.Slice(nextSteps, func(i, j int) bool {
		return nextSteps[i].value < nextSteps[j].value
	})

	return nextSteps

}

func findPath(solvingMap [][]CityBlock, x, y int, direction Direction, straightSteps int, sum int, minHeatLoss *int, found *bool, path *Stack) {
	if x < 0 || y < 0 || x >= len(solvingMap) || y >= len(solvingMap[0]) {
		return
	}
	if solvingMap[x][y].visited {
		return
	}
	sum += solvingMap[x][y].value
	if sum+solvingMap[x][y].minHeatLoss > *minHeatLoss {
		return
	}
	if x == len(solvingMap)-1 && y == len(solvingMap[0])-1 {
		*found = true
		fmt.Println("Found path: ", sum)
		printMapPath(solvingMap, *path)
		if sum < *minHeatLoss {
			*minHeatLoss = sum
		}
		return
	}
	solvingMap[x][y].visited = true
	path.Push(PointStep{x, y, straightSteps})
	nextSteps := getNextSteps(solvingMap, x, y)
	switch direction {
	case Up:
		for _, nextStep := range nextSteps {
			switch nextStep.direction {
			case Up:
				if straightSteps < 3 {
					findPath(solvingMap, x-1, y, Up, straightSteps+1, sum, minHeatLoss, found, path)
				}
			case Left:
				findPath(solvingMap, x, y-1, Left, 1, sum, minHeatLoss, found, path)
			case Right:
				findPath(solvingMap, x, y+1, Right, 1, sum, minHeatLoss, found, path)
			}
		}
	case Down:
		for _, nextStep := range nextSteps {
			switch nextStep.direction {
			case Down:
				if straightSteps < 3 {
					findPath(solvingMap, x+1, y, Down, straightSteps+1, sum, minHeatLoss, found, path)
				}
			case Left:
				findPath(solvingMap, x, y-1, Left, 1, sum, minHeatLoss, found, path)
			case Right:
				findPath(solvingMap, x, y+1, Right, 1, sum, minHeatLoss, found, path)
			}
		}

	case Left:
		for _, nextStep := range nextSteps {
			switch nextStep.direction {
			case Left:
				if straightSteps < 3 {
					findPath(solvingMap, x, y-1, Left, straightSteps+1, sum, minHeatLoss, found, path)
				}
			case Up:
				findPath(solvingMap, x-1, y, Up, 1, sum, minHeatLoss, found, path)
			case Down:
				findPath(solvingMap, x+1, y, Down, 1, sum, minHeatLoss, found, path)
			}
		}
	case Right:
		for _, nextStep := range nextSteps {
			switch nextStep.direction {
			case Right:
				if straightSteps < 3 {
					findPath(solvingMap, x, y+1, Right, straightSteps+1, sum, minHeatLoss, found, path)
				}
			case Up:
				findPath(solvingMap, x-1, y, Up, 1, sum, minHeatLoss, found, path)
			case Down:
				findPath(solvingMap, x+1, y, Down, 1, sum, minHeatLoss, found, path)
			}
		}

	}
	solvingMap[x][y].visited = false
	path.Pop()
}
