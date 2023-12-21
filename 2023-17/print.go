package main

import (
	"fmt"
)

func printData(data [][]int) {
	for _, z := range data {
		for _, y := range z {
			fmt.Print(y)
		}
		fmt.Println()
	}
	fmt.Println()
}

func printMap(data [][]CityBlock) {
	for x, line := range data {
		for y, block := range line {
			fmt.Printf("%1d (%3d) %2d,%2d  | ", block.value, block.minHeatLoss, x, y)
		}
		fmt.Println()
	}
	fmt.Println()
}

func printPath(path *Stack) {
	for i := len(*path) - 1; i >= 0; i-- {
		fmt.Print((*path)[i])
	}
	fmt.Println()

}

func pathContains(path []PointStep, point PointStep) bool {
	for _, p := range path {
		if p.X == point.X && p.Y == point.Y {
			return true
		}
	}
	return false
}

func printMapPath(data [][]CityBlock, path []PointStep) {
	fmt.Println("Path:", path)
	for x := 0; x < len(data); x++ {
		for y := 0; y < len(data[0]); y++ {
			b := data[x][y]
			if pathContains(path, PointStep{x, y, 0}) {
				fmt.Printf("%1d (%3d) %c | ", b.value, b.minHeatLoss, 'o')
			} else {

				fmt.Printf("%1d (%3d)   | ", b.value, b.minHeatLoss)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
