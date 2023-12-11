package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Point struct {
	X int
	Y int
}

func loadData(fileName string, emptyGalazySize int) ([]Point, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	galaxies := make([]Point, 0)
	x := 0
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading from file:", err)
			return nil, err
		}
		noGalaxies := true
		for y, ch := range scanner.Text() {
			if ch == '#' {
				p := Point{X: x, Y: y}
				galaxies = append(galaxies, p)
				noGalaxies = false
			}
		}
		if noGalaxies {
			x += emptyGalazySize
		}
		x++
	}
	return galaxies, nil
}

func add_expansion(galaxies []Point, emptyGalazySize int) []Point {
	newWorld := make([]Point, 0)
	shift := 0
	for index := 0; index <= galaxies[len(galaxies)-1].Y; index++ {
		emptyLine := true
		for _, g := range galaxies {
			if g.Y == index {
				newWorld = append(newWorld, Point{X: g.X, Y: g.Y + shift})
				emptyLine = false
			}
		}
		if emptyLine {
			shift += emptyGalazySize
		}
	}
	return newWorld
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func getDistanceSum(galaxies []Point) int {
	sum := 0
	index := 0
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			distance := abs(galaxies[i].X-galaxies[j].X) + abs(galaxies[i].Y-galaxies[j].Y)
			sum += distance
			index++
		}
	}
	return sum
}

func main() {
	galaxies, _ := loadData("input.txt", 1)
	sort.Slice(galaxies, func(n, m int) bool {
		return galaxies[n].Y < galaxies[m].Y
	})
	newWorld := add_expansion(galaxies, 1)
	sum := getDistanceSum(newWorld)
	fmt.Println("Puzzle 1 - sum:", sum)

	// PART II
	galaxies, _ = loadData("input.txt", 999999)
	sort.Slice(galaxies, func(n, m int) bool {
		return galaxies[n].Y < galaxies[m].Y
	})
	newWorld = add_expansion(galaxies, 999999)
	sum = getDistanceSum(newWorld)

	fmt.Println("Puzzle 2 - sum:", sum)
}
