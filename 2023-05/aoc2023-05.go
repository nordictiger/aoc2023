package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getSeeds(scanner *bufio.Scanner) ([]int, error) {
	for {
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading from file:", err)
			return nil, err
		}

		if strings.HasPrefix(scanner.Text(), "seeds:") {
			break
		}
	}

	seedNumbers := strings.Fields(scanner.Text())
	result := make([]int, 0, len(seedNumbers))

	for _, seed := range seedNumbers[1:] {
		num, err := strconv.Atoi(seed)
		if err != nil {
			fmt.Println("Error converting seed to int:", err)
			return nil, err
		}
		result = append(result, num)
	}
	return result, nil
}

func getMaps(scanner *bufio.Scanner) ([][][]int, error) {
	result := make([][][]int, 0)
	for scanner.Scan() {
		// Skip until map prefix
		for {
			if err := scanner.Err(); err != nil {
				fmt.Println("Error reading from file:", err)
				return nil, err
			}

			if strings.Contains(scanner.Text(), ":") {
				break
			}
			scanner.Scan()
		}
		// Load map
		mapNumbers := make([][]int, 0)
		for {
			scanner.Scan()
			if err := scanner.Err(); err != nil {
				fmt.Println("Error reading from file:", err)
				return nil, err
			}
			if scanner.Text() == "" {
				break
			}
			mapLineNumbers := strings.Fields(scanner.Text())
			mapLine := make([]int, 0, len(mapLineNumbers))

			for _, seed := range mapLineNumbers {
				num, err := strconv.Atoi(seed)
				if err != nil {
					fmt.Println("Error converting seed to int:", err)
					return nil, err
				}
				mapLine = append(mapLine, num)
			}
			mapNumbers = append(mapNumbers, mapLine)
		}
		result = append(result, mapNumbers)
	}

	return result, nil
}

func loadData(fileName string) ([]int, [][][]int) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	seeds, _ := getSeeds(scanner)
	seedMaps, _ := getMaps(scanner)

	return seeds, seedMaps
}

func solvePuzzle1(seeds []int, maps [][][]int) int {
	lowestLocation := 1<<31 - 1
	for _, seed := range seeds {
		baton := seed
		for _, seedMap := range maps {
			for _, mapLine := range seedMap {
				if (mapLine[1] <= baton) && (baton < mapLine[1]+mapLine[2]) {
					baton = mapLine[0] + (baton - mapLine[1])
					break
				}
			}

		}
		if baton < lowestLocation {
			lowestLocation = baton
		}
	}
	return lowestLocation
}

func solvePuzzle2(seeds []int, maps [][][]int) int {
	lowestLocation := 1<<31 - 1
	ranges := make([][]int, 0)
	for sp := 0; sp < len(seeds); sp += 2 {
		pair := make([]int, 2)
		pair[0] = seeds[sp]
		pair[1] = seeds[sp] + seeds[sp+1]
		ranges = append(ranges, pair)
	}
	for _, seedMap := range maps {
		ranges = getPairs(ranges, seedMap)
	}

	for _, pair := range ranges {
		if pair[0] < lowestLocation {
			lowestLocation = pair[0]
		}
	}
	return lowestLocation
}

func getPairs(ranges [][]int, seedMap [][]int) [][]int {
	newRanges := make([][]int, 0)
	for _, pair := range ranges {
		for _, mapLine := range seedMap {
			// find all possible new pairs for this pair and mapLine[1] and mapLine[2]
			if (pair[0] >= mapLine[1]) && (pair[1] < mapLine[1]+mapLine[2]) {
				newPair := make([]int, 2)
				newPair[0] = mapLine[0] + (pair[0] - mapLine[1])
				newPair[1] = mapLine[0] + (pair[0] - mapLine[1]) + (pair[1] - pair[0])
				newRanges = append(newRanges, newPair)
			} else if (pair[0] < mapLine[1]) && (pair[1] >= mapLine[1]+mapLine[2]) {
				newPair := make([]int, 2)
				newPair[0] = mapLine[0]
				newPair[1] = mapLine[0] + mapLine[2]
				newRanges = append(newRanges, newPair)
			} else if (pair[0] >= mapLine[1]) && (pair[0] < mapLine[1]+mapLine[2]) {
				newPair := make([]int, 2)
				newPair[0] = mapLine[0] + (pair[0] - mapLine[1])
				newPair[1] = mapLine[0] + mapLine[2]
				newRanges = append(newRanges, newPair)
			} else if (pair[1] >= mapLine[1]) && (pair[1] <= mapLine[1]+mapLine[2]) {
				newPair := make([]int, 2)
				newPair[0] = mapLine[0]
				newPair[1] = mapLine[0] + (pair[1] - mapLine[1])
				newRanges = append(newRanges, newPair)
			}
		}
	}
	return newRanges
}

func puzzle1() {
	fmt.Println("Puzzle 1")
	seeds, maps := loadData("input.txt")
	solution := solvePuzzle1(seeds, maps)
	fmt.Println("Lowest location number is ", solution)
}

func puzzle2() {
	fmt.Println("Puzzle 2")
	seeds, maps := loadData("input.txt")
	solution := solvePuzzle2(seeds, maps)
	fmt.Println("Lowest location number is ", solution)
}

func main() {
	puzzle1()
	puzzle2()
}
