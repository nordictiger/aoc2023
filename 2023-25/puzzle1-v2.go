package main

func getIntMaps(diagram map[string][]string) (map[string]int, map[int]string) {
	tmpNodes := make(map[string]bool, 1600)
	for k, v := range diagram {
		for _, vv := range v {
			tmpNodes[k] = true
			tmpNodes[vv] = true
		}
	}
	nodesString := make(map[string]int, 1600)
	nodesInt := make(map[int]string, 1600)
	index := 0
	for k := range tmpNodes {
		nodesString[k] = index
		nodesInt[index] = k
		index++
	}
	return nodesString, nodesInt
}

func getAmGraph(diagram map[string][]string, nodesString map[string]int) *amGraph {
	g := NewGraph(len(nodesString))
	for k, v := range diagram {
		for _, vv := range v {
			g.AddEdge(nodesString[k], nodesString[vv])
		}
	}
	return g
}

func puzzle1V2(diagram map[string][]string) int {
	nodesString, _ := getIntMaps(diagram)
	g := getAmGraph(diagram, nodesString)
	// g.PrintEdges()
	result := g.kargerMinCut()
	/*
		nodes, edges := getGraphTracking(diagram)
		minCut := len(nodes) * len(nodes)
		for i := 0; i < 100000; i++ {
			nodes, edges = getGraphTracking(diagram)
			cut, _ := kargerMinCutTracking(nodes, edges)
			if cut < minCut {
				minCut = cut
			}
		}
		fmt.Println("Puzzle 1:", minCut)
	*/
	return result
}
