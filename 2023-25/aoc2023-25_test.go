package main

import (
	"fmt"
	"testing"
)

func TestData(t *testing.T) {
	diagram := loadData("input-test.txt")
	nodes, edges := getGraphTracking(diagram)
	fmt.Println(diagram)
	fmt.Println(nodes)
	fmt.Println(edges)
}

func TestPuzzle1(t *testing.T) {
	diagram := loadData("input-test.txt")
	fmt.Println("Puzzle 1: ", puzzle1(diagram))
}
