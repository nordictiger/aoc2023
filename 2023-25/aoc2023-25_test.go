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

func TestPuzzle1V2(t *testing.T) {
	diagram := loadData("input-test.txt")
	fmt.Println("Puzzle 1 v2: ", puzzle1V2(diagram))
}

func TestPuzzle1V2full(t *testing.T) {
	diagram := loadData("input.txt")
	fmt.Println("Puzzle 1 v2: ", puzzle1V2(diagram))
}

func TestAmGraph(t *testing.T) {
	g := NewGraph(5)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)

	fmt.Println("Initial Graph:")
	g.PrintEdges()

	// Merge node 1 into node 0
	g.RemoveEdgeMergeV2(0, 1)

	fmt.Println("Graph after merging nodes 0 and 1:")
	g.PrintEdges()

}
