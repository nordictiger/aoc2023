package main

import (
	"math/rand"
)

func contractTracking(vertices []string, edges map[edge]int, edgeMapping map[edge]map[edge]bool) ([]string, map[edge]int, map[edge]map[edge]bool) {
	randomIndex := rand.Intn(len(edges))
	i := 0
	var randomEdge edge
	for k := range edges {
		if i == randomIndex {
			randomEdge = k
			break
		}
		i++
	}
	u, v := randomEdge.u, randomEdge.v

	newEdges := make(map[edge]int, 3600)
	newEdgeMapping := make(map[edge]map[edge]bool, 3600)

	for e, count := range edges {
		a, b := e.u, e.v
		if a == v {
			a = u
		}
		if b == v {
			b = u
		}
		if a != b {
			newEdge := edge{a, b}
			newEdges[newEdge] += count
			if newEdgeMapping[newEdge] == nil {
				newEdgeMapping[newEdge] = make(map[edge]bool)
			}
			for originalEdge := range edgeMapping[e] {
				newEdgeMapping[newEdge][originalEdge] = true
			}
		}
	}

	var newVertices []string
	for _, vertex := range vertices {
		if vertex != v {
			newVertices = append(newVertices, vertex)
		}
	}

	return newVertices, newEdges, newEdgeMapping
}

func kargerMinCutTracking(vertices []string, edges map[edge]int) (int, map[edge]bool) {
	edgeMapping := make(map[edge]map[edge]bool)
	for e := range edges {
		edgeMapping[e] = map[edge]bool{e: true}
	}

	for len(vertices) > 2 {
		vertices, edges, edgeMapping = contractTracking(vertices, edges, edgeMapping)
	}

	finalCutEdges := make(map[edge]bool)
	for _, mapping := range edgeMapping {
		for edge := range mapping {
			finalCutEdges[edge] = true
		}
	}
	return len(finalCutEdges), finalCutEdges
}
