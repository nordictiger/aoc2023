package main

import (
	"fmt"
)

type edge struct {
	u string
	v string
}

func getGraphTracking(diagram map[string][]string) ([]string, map[edge]int) {
	nodes := make([]string, 0)
	tmpNodes := make(map[string]bool, 1600)
	edges := make(map[edge]int, 3600)
	for k, v := range diagram {
		for _, vv := range v {
			if k < vv {
				edges[edge{k, vv}] = 1
			} else {
				edges[edge{vv, k}] = 1
			}
			tmpNodes[k] = true
			tmpNodes[vv] = true
		}
	}
	for k := range tmpNodes {
		nodes = append(nodes, k)
	}

	return nodes, edges
}

func getEdges(diagram map[string][]string) []edge {
	edges := make([]edge, 0)
	tmpEdges := make(map[edge]int)

	for k, v := range diagram {
		for _, vv := range v {
			if k < vv {
				tmpEdges[edge{k, vv}] = 1
			} else {
				tmpEdges[edge{vv, k}] = 1
			}
		}
	}
	for k := range tmpEdges {
		edges = append(edges, k)
	}

	return edges
}

func cutEdges(edges []edge, cut map[edge]bool) []edge {
	cutEdges := make([]edge, 0, len(edges))
	for _, e := range edges {
		if _, ok := cut[e]; ok {
			continue
		}
		cutEdges = append(cutEdges, e)
	}
	return cutEdges
}

func getGraphMap(edges []edge) map[string]map[string]bool {
	result := make(map[string]map[string]bool, len(edges)*2)
	for _, e := range edges {
		if result[e.u] == nil {
			result[e.u] = make(map[string]bool)
		}
		result[e.u][e.v] = true
		if result[e.v] == nil {
			result[e.v] = make(map[string]bool)
		}
		result[e.v][e.u] = true
	}
	return result
}

func countPartitions(start string, graph map[string]map[string]bool, visited *map[string]bool) int {
	initialCount := len(*visited)
	queue := []string{start}
	(*visited)[start] = true
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for k := range graph[node] {
			if _, ok := (*visited)[k]; !ok {
				queue = append(queue, k)
				(*visited)[k] = true
			}
		}
	}
	return len(*visited) - initialCount
}

func getPartitionSizes(graph map[string]map[string]bool) []int {
	visited := make(map[string]bool)
	var start string
	for k := range graph {
		if _, ok := visited[k]; !ok {
			start = k
			break
		}
	}
	part1 := countPartitions(start, graph, &visited)
	fmt.Println("Partition 1 size: ", part1)
	for k := range graph {
		if _, ok := visited[k]; !ok {
			start = k
			break
		}
	}
	part2 := countPartitions(start, graph, &visited)
	fmt.Println("Partition 2 size: ", part2)
	for k := range graph {
		if _, ok := visited[k]; !ok {
			fmt.Println("There should be no unvisited nodes left")
			break
		}
	}
	fmt.Println("All good, all nodes visited")

	return []int{part1, part2}
}

func puzzle1(diagram map[string][]string) int {
	nodes, edges := getGraphTracking(diagram)
	fmt.Println("Nodes: ", len(nodes), "Edges: ", len(edges))
	counter := 0
	count := 0
	var edgesToCut map[edge]bool
	fmt.Println("Starting loop. It may take up to few hours.")
	for {
		count, edgesToCut = kargerMinCutTracking(nodes, edges)
		if count == 3 {
			fmt.Println("Found cut")
			fmt.Println(edgesToCut)
			break
		}
		counter++
		fmt.Print(count, ",")
	}
	/*
		edgesToCut := map[edge]bool{
			{u: "dgc", v: "fqn"}: true,
			{u: "htp", v: "vps"}: true,
			{u: "rpd", v: "ttj"}: true,
		}
	*/
	edgesSlice := getEdges(diagram)
	cutEdges := cutEdges(edgesSlice, edgesToCut)
	fmt.Println("Cut edges: ", len(cutEdges))
	graphAfterCut := getGraphMap(cutEdges)
	fmt.Println("Graph after cut: ", len(graphAfterCut))
	partitionSizes := getPartitionSizes(graphAfterCut)
	return partitionSizes[0] * partitionSizes[1]
}

func main() {

	diagram := loadData("input.txt")
	fmt.Println("Puzzle 1: ", puzzle1(diagram))
	// fmt.Println("Puzzle 1: ", puzzle1V2(diagram))

}
