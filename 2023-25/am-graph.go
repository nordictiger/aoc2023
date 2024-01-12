package main

import (
	"fmt"
	"math/rand"
)

type amGraph struct {
	matrix   [][]bool
	vertices map[int]bool
}

func NewGraph(maxVertices int) *amGraph {
	matrix := make([][]bool, maxVertices)
	for i := range matrix {
		matrix[i] = make([]bool, maxVertices)
	}
	return &amGraph{
		matrix:   matrix,
		vertices: make(map[int]bool, maxVertices),
	}
}

func (g *amGraph) outOfRange(v1, v2 int) bool {
	if v1 > len(g.matrix)-1 || v2 > len(g.matrix)-1 || v1 < 0 || v2 < 0 {
		panic("Invalid vertex number")
	}
	return false
}

func (g *amGraph) AddEdge(v1, v2 int) {
	if g.outOfRange(v1, v2) {
		return
	}
	g.matrix[v1][v2] = true
	g.matrix[v2][v1] = true
	g.vertices[v1] = true
	g.vertices[v2] = true
}

// MergeNodes merges the connections of v2 into v1 and removes v2
func (g *amGraph) RemoveEdgeMergeV2(v1, v2 int) {
	if g.outOfRange(v1, v2) {
		return
	}

	// Transfer all connections of endNode to startNode
	for v := 0; v < len(g.matrix); v++ {
		if g.matrix[v2][v] {
			g.matrix[v1][v] = true
			g.matrix[v][v1] = true
			g.matrix[v2][v] = false
			g.matrix[v][v2] = false
		}
	}

	// Remove any self-loop that might have been created
	g.matrix[v1][v1] = false
	if _, ok := g.vertices[v2]; !ok {
		panic("Vertex not found")
	}
	delete(g.vertices, v2)
}

func (g *amGraph) Print() {
	for _, row := range g.matrix {
		for _, val := range row {
			if val {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
}

func (g *amGraph) GetEdgeCount() int {
	count := 0
	for i := 0; i < len(g.matrix); i++ {
		for j := 0; j < len(g.matrix[0]); j++ {
			if g.matrix[i][j] {
				count++
			}
		}
	}
	return count / 2
}
func (g *amGraph) GetRandomEdge() (int, int) {
	randomI := rand.Intn(len(g.vertices) - 1)
	v1, v2 := -1, -1
	i := 0
	for k := range g.vertices {
		if i == randomI {
			v1 = k
			break
		}
		i++
	}
	if v1 == -1 {
		panic("Unable to select random edge")
	}
	v1EdgesCount := 0
	for j := 0; j < len(g.matrix[0]); j++ {
		if g.matrix[v1][j] {
			v1EdgesCount++
		}
	}
	randomJ := rand.Intn(v1EdgesCount)
	j := 0
	for k := range g.matrix[v1] {
		if g.matrix[v1][k] {
			if j == randomJ {
				v2 = k
				break
			}
			j++
		}
	}
	if v2 == -1 {
		panic("Unable to select random edge")
	}
	return v1, v2
}

func (g *amGraph) PrintEdges() {
	for i := 0; i < len(g.matrix); i++ {
		for j := 0; j < len(g.matrix[0]); j++ {
			if g.matrix[i][j] {
				fmt.Printf("(%d, %d)\n", i, j)
			}
		}
	}
	fmt.Println()
}

func (g *amGraph) kargerMinCut() int {
	for len(g.vertices) > 2 {
		u, v := g.GetRandomEdge()
		g.RemoveEdgeMergeV2(u, v)
	}
	g.PrintEdges()
	return g.GetEdgeCount()
}
