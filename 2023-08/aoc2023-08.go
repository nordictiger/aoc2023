package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func loadData(fileName string) (string, map[string]map[byte]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "", nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return "", nil, err
	}
	path := scanner.Text()

	network := make(map[string]map[byte]string)

	re := regexp.MustCompile(`(\w{3})`)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading from file:", err)
			return "", nil, err
		}
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		if len(matches) != 3 {
			continue
		}
		node := make(map[byte]string)
		node['L'] = matches[1][0]
		node['R'] = matches[2][0]
		network[matches[0][0]] = node
	}
	return path, network, nil
}

func puzzle1(path string, network map[string]map[byte]string) int {
	fmt.Println("Puzzle 1")
	steps := 0
	index := 0
	activeNode := "AAA"
	for {
		if index > len(path)-1 {
			index = 0
		}
		if activeNode == "ZZZ" {
			break
		}
		activeNode = network[activeNode][path[index]]
		index++
		steps++
	}
	return steps
}

func getPuzzle2Steps(path string, network map[string]map[byte]string, start string) int {
	steps := 0
	index := 0
	activeNode := start
	for {
		if index > len(path)-1 {
			index = 0
		}
		if activeNode[2] == 'Z' {
			break
		}
		activeNode = network[activeNode][path[index]]
		index++
		steps++
	}
	return steps
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func lcmArray(arr []int) int {
	result := arr[0]
	for i := 1; i < len(arr); i++ {
		result = lcm(result, arr[i])
	}
	return result
}

func puzzle2(path string, network map[string]map[byte]string) int {
	fmt.Println("Puzzle 2")
	startNodes := make(map[string]int)
	for key := range network {
		if key[2] == 'A' {
			startNodes[key] = getPuzzle2Steps(path, network, key)
		}
	}

	allSteps := make([]int, 0)
	for _, steps := range startNodes {
		allSteps = append(allSteps, steps)
	}

	return lcmArray(allSteps)
}

func main() {
	path, network, _ := loadData("input.txt")
	steps := puzzle1(path, network)
	fmt.Println("steps:", steps)
	steps = puzzle2(path, network)
	fmt.Println("steps:", steps)
}
