package main

import "fmt"

func copyGarden(garden [][]byte) [][]byte {
	newGarden := make([][]byte, len(garden))
	for i := 0; i < len(garden); i++ {
		newGarden[i] = make([]byte, len(garden[i]))
		copy(newGarden[i], garden[i])
	}
	return newGarden
}

func getStart(garden [][]byte) (int, int) {
	for i := 0; i < len(garden); i++ {
		for j := 0; j < len(garden[i]); j++ {
			if garden[i][j] == 'S' {
				return i, j
			}
		}
	}
	return -1, -1
}

func step(steps [][]byte, garden [][]byte) [][]byte {
	newSteps := copyGarden(garden)
	for i := 0; i < len(steps); i++ {
		for j := 0; j < len(steps[i]); j++ {
			if steps[i][j] == 'O' || steps[i][j] == 'S' {
				updateSteps(newSteps, i, j, garden)
			}
		}
	}
	return newSteps
}

func updateSteps(steps [][]byte, x, y int, garden [][]byte) {
	if x > 0 && garden[x-1][y] == '.' {
		steps[x-1][y] = 'O'
	}
	if x < len(garden)-1 && garden[x+1][y] == '.' {
		steps[x+1][y] = 'O'
	}
	if y > 0 && garden[x][y-1] == '.' {
		steps[x][y-1] = 'O'
	}
	if y < len(garden[x])-1 && garden[x][y+1] == '.' {
		steps[x][y+1] = 'O'
	}
}

func countVisitedPlots(garden [][]byte) int {
	count := 0
	for i := 0; i < len(garden); i++ {
		for j := 0; j < len(garden[i]); j++ {
			if garden[i][j] == 'O' {
				count++
			}
		}
	}
	return count
}

func runSteps(garden [][]byte, reminingSteps, startX, startY int) int {
	steps := copyGarden(garden)
	// printGarden(steps)
	steps[startX][startY] = 'S'
	for i := 0; i < reminingSteps; i++ {
		steps = step(steps, garden)
	}
	// fmt.Println()
	// printGarden(steps)
	sum := countVisitedPlots(steps)
	return sum
}

func puzzle1(data [][]byte, reminingSteps int) int {
	startX, startY := getStart(data)
	data[startX][startY] = '.'
	sum := runSteps(data, reminingSteps, startX, startY)
	return sum
}
func puzzle2(data [][]byte, reminingSteps int) int {
	// Prepare garden for getting reachable plots
	garden := copyGarden(data)
	startX, startY := getStart(garden)
	garden[startX][startY] = '.'

	fullMapPlotsOdd := runSteps(garden, 131+131+65, startX, startY)
	fullMapPlotsEven := runSteps(garden, 131+131+65+1, startX, startY)

	// there are two stable counts flip-floping 7344, 7451
	// center is 7344 - let's call it odd, hence 7451 is even
	sum := fullMapPlotsOdd // central point map

	// how many maps we need to add for moving straight up, right, down, left
	mapsAdded := reminingSteps / 131
	// last map is "arrowhead", so not complete
	fullMaps := mapsAdded - 1
	evenMaps := fullMaps/2 + 1
	oddMaps := fullMaps / 2
	triangle := evenMaps*evenMaps*fullMapPlotsEven + oddMaps*(oddMaps+1)*fullMapPlotsOdd
	sum += 4 * triangle

	// get arrowheads
	arrowUp := runSteps(garden, 130, 130, 65)
	arrowRight := runSteps(garden, 130, 65, 0)
	arrowDown := runSteps(garden, 130, 0, 65)
	arrowLeft := runSteps(garden, 130, 65, 130)

	// get corners
	upperLeftFull := runSteps(garden, 195, 0, 0)
	upperLeft := runSteps(garden, 64, 0, 0)
	upperRightFull := runSteps(garden, 195, 0, 130)
	upperRight := runSteps(garden, 64, 0, 130)
	lowerRight := runSteps(garden, 64, 130, 130)
	lowerRightFull := runSteps(garden, 195, 130, 130)
	lowerLeft := runSteps(garden, 64, 130, 0)
	lowerLeftFull := runSteps(garden, 195, 130, 0)

	// add arrowheads
	sum += arrowUp
	sum += arrowRight
	sum += arrowDown
	sum += arrowLeft

	// add diagonals
	sum += fullMaps*lowerLeftFull + mapsAdded*lowerLeft
	sum += fullMaps*upperLeftFull + mapsAdded*upperLeft
	sum += fullMaps*upperRightFull + mapsAdded*upperRight
	sum += fullMaps*lowerRightFull + mapsAdded*lowerRight

	return sum
}

func main() {
	// garden := loadData("input-test.txt")
	// fmt.Println("Puzzle 1: ", puzzle1(garden, 6))
	garden := loadData("input.txt")
	fmt.Println("Puzzle 1: ", puzzle1(garden, 64))
	garden = loadData("input.txt")
	fmt.Println("Puzzle 2: ", puzzle2(garden, 26501365))
}
