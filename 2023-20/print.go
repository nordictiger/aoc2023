package main

import "fmt"

func printDeps(mc moduleConfiguration, key string, level int) {
	fmt.Println(key, "depends on:")
	deps := make([]string, 0)
	deps = append(deps, key)
	for i := 0; i < level; i++ {
		fmt.Println("level: ", i+1)
		nextDeps := make([]string, 0)
		for _, dep := range deps {
			for _, node := range mc {
				for _, value := range node.outgoing {
					if value == dep {
						nextDeps = append(nextDeps, node.name)
						break
					}
				}
			}
		}
		for _, dep := range nextDeps {
			fmt.Println(mc[dep].name, mc[dep].moduleType, mc[dep].state)
		}
		deps = nextDeps
	}
}
