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

func printPath(path *Stack) {
	for i := len(*path) - 1; i >= 0; i-- {
		fmt.Print((*path)[i])
	}
	fmt.Println()
}
