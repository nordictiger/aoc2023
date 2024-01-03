package main

import (
	"fmt"
	"sort"
)

type point3d struct {
	x, y, z int
}

type brick struct {
	start, end point3d
}

type stackedBrick struct {
	id int
	brick
	supports    []int
	supportedBy []int
}

func getCrossings(stack []stackedBrick, b brick) []int {
	crossings := make([]int, 0)
	for i, sb := range stack {
		// check for crossings in each layer
		if b.start.x <= sb.end.x && b.end.x >= sb.start.x &&
			b.start.y <= sb.end.y && b.end.y >= sb.start.y {
			// crossing found
			crossings = append(crossings, i)
		}
	}
	return crossings
}

func getNewPosition(crossings []int, newBrick *stackedBrick, stack []stackedBrick) {
	sortedCrossings := make([]stackedBrick, 0)
	for _, c := range crossings {
		sortedCrossings = append(sortedCrossings, stack[c])
	}
	// crossings found, find highest layer
	sort.Slice(sortedCrossings, func(i, j int) bool {
		return sortedCrossings[i].end.z > sortedCrossings[j].end.z
	})
	newBrickLen := newBrick.end.z - newBrick.start.z
	// get the new Z for the falling brick
	for _, c := range sortedCrossings {
		if c.end.z < newBrick.start.z {
			newBrick.start.z = c.end.z + 1
			newBrick.end.z = newBrick.start.z + newBrickLen
			break
		}
	}
}

func getSupportBricks(crossings []int, newBrick *stackedBrick, stack []stackedBrick) {
	// get all bricks that are supporting by the new brick
	for _, c := range crossings {
		if stack[c].end.z == newBrick.start.z-1 {
			newBrick.supportedBy = append(newBrick.supportedBy, stack[c].id)
			stack[c].supports = append(stack[c].supports, newBrick.id)
		}
	}
}

func compileBricks(sortedBrics []brick) []stackedBrick {
	stack := make([]stackedBrick, 0)
	for i, b := range sortedBrics {
		newBrick := stackedBrick{
			id:          i,
			brick:       b,
			supports:    []int{},
			supportedBy: []int{},
		}
		newBrickLen := newBrick.end.z - newBrick.start.z

		if b.start.z == 1 {
			stack = append(stack, newBrick)
			continue
		}

		crossings := getCrossings(stack, b)
		if len(crossings) == 0 {
			// No crossings found, add brick to layer 1
			newBrick.start.z = 1
			newBrick.end.z = newBrick.start.z + newBrickLen
			stack = append(stack, newBrick)
			continue
		}
		getNewPosition(crossings, &newBrick, stack)

		getSupportBricks(crossings, &newBrick, stack)
		stack = append(stack, newBrick)
	}
	return stack
}

func countRemovableBricks(stack []stackedBrick) int {
	bricksMap := make(map[int]stackedBrick)
	for _, sb := range stack {
		bricksMap[sb.id] = sb
	}
	removableBricks := 0
	for _, sb := range bricksMap {
		if len(sb.supports) == 0 {
			removableBricks++
			continue
		}
		safeToRemove := true
		for _, s := range sb.supports {
			if len(bricksMap[s].supportedBy) == 1 {
				safeToRemove = false
			}
		}
		if safeToRemove {
			removableBricks++
		}
	}
	return removableBricks
}

func puzzle1(bricks []brick) int {
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].start.z < bricks[j].start.z
	})
	stack := compileBricks(bricks)
	return countRemovableBricks(stack)
}

func contains(eliminated map[int]bool, supportedBy []int) bool {

	for _, item := range supportedBy {
		if _, ok := eliminated[item]; !ok {
			return false
		}
	}
	return true
}

func countCascade(sb stackedBrick, bricksMap map[int]stackedBrick) int {
	eliminated := make(map[int]bool, 0)
	eliminated[sb.id] = true
	var inProgress []int
	inProgress = append(inProgress, sb.supports...)
	var nextStep []int
	for len(inProgress) > 0 {
		for _, s := range inProgress {
			if contains(eliminated, bricksMap[s].supportedBy) {
				eliminated[s] = true
				nextStep = append(nextStep, bricksMap[s].supports...)
			}
		}
		inProgress = nextStep
		nextStep = make([]int, 0)
	}
	return len(eliminated) - 1
}

func countFallingBricks(stack []stackedBrick) int {
	fallingBricks := 0
	bricksMap := make(map[int]stackedBrick)
	for _, sb := range stack {
		bricksMap[sb.id] = sb
	}

	for _, sb := range stack {
		fallingBricks += countCascade(sb, bricksMap)
	}
	return fallingBricks
}

func puzzle2(bricks []brick) int {
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].start.z < bricks[j].start.z
	})
	stack := compileBricks(bricks)
	return countFallingBricks(stack)
}

func main() {
	bricks := loadData("input.txt")
	// bricks := loadData("input-test.txt")
	fmt.Println("Puzzle 1: ", puzzle1(bricks))
	fmt.Println("Puzzle 2: ", puzzle2(bricks))
}
