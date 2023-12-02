package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func checkGame2(line string) int {
	reGems := regexp.MustCompile(`(\d+) (blue|green|red)`)

	firstSplit := strings.Split(line, ":")
	if len(firstSplit) != 2 {
		fmt.Println("Error: firstSplit length is not 2")
		return -1
	}
	secondSplit := strings.Split(firstSplit[1], ";")
	minGameGems := make(map[string]int)
	for _, subSet := range secondSplit {
		subSetGemsMatch := reGems.FindAllStringSubmatch(subSet, -1)
		for _, gemMatch := range subSetGemsMatch {
			gemCountInt, err := strconv.Atoi(gemMatch[1])
			if err != nil {
				fmt.Printf("Error converting '%s' to number: %s\n", gemMatch[1], err)
				return -1
			}
			if minGameGems[gemMatch[2]] < gemCountInt {
				minGameGems[gemMatch[2]] = gemCountInt
			}
		}
	}
	return minGameGems["blue"] * minGameGems["green"] * minGameGems["red"]
}

func puzzle2() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		gameProduct := checkGame2(line)
		sum += gameProduct
		// fmt.Println(line)x
		// fmt.Println("gameProduct:", gameProduct)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}
	fmt.Println("Sum of possible games:", sum)
}
