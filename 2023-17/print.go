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
	for _, z := range data {
		for _, y := range z {
			fmt.Printf("%1d (%3d) %2d,%2d  | ", y.value, y.minHeatLoss, y.minHeatLossPoint.x, y.minHeatLossPoint.y)
		}
		fmt.Println()
	}
	fmt.Println()
}

func pathContains(path []Point, point Point) bool {
	for _, p := range path {
		if p.x == point.x && p.y == point.y {
			return true
		}
	}
	return false
}

func printMapPath(data [][]CityBlock, path []Point) {
	fmt.Println("Path:", path)
	for x := 0; x < len(data); x++ {
		for y := 0; y < len(data[0]); y++ {
			b := data[x][y]
			if pathContains(path, Point{x, y}) {
				fmt.Printf("%1d (%3d) - %5c | ", b.value, b.minHeatLoss, 'o')
			} else {

				fmt.Printf("%1d (%3d) - %2d,%2d | ", b.value, b.minHeatLoss, b.minHeatLossPoint.x, b.minHeatLossPoint.y)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
