package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func checkGame(line string) (int, bool) {
	reGame := regexp.MustCompile(`Game (\d+)`)
	reGems := regexp.MustCompile(`(\d+) (blue|green|red)`)

	firstSplit := strings.Split(line, ":")
	if len(firstSplit) != 2 {
		fmt.Println("Error: firstSplit length is not 2")
		return 0, false
	}
	gameNumberText := reGame.FindStringSubmatch(firstSplit[0])
	gameNumberInt, err := strconv.Atoi(gameNumberText[1])
	if err != nil {
		fmt.Printf("Error converting '%s' to number: %s\n", gameNumberText, err)
		return 0, false
	}

	secondSplit := strings.Split(firstSplit[1], ";")
	possibleGame := true
	for _, subSet := range secondSplit {
		subSetGems := make(map[string]int)
		subSetGemsMatch := reGems.FindAllStringSubmatch(subSet, -1)
		for _, gemMatch := range subSetGemsMatch {
			gemCountInt, err := strconv.Atoi(gemMatch[1])
			if err != nil {
				fmt.Printf("Error converting '%s' to number: %s\n", gemMatch[1], err)
				return 0, false
			}
			subSetGems[gemMatch[2]] = gemCountInt
		}
		if subSetGems["blue"] > 14 || subSetGems["green"] > 13 || subSetGems["red"] > 12 {
			possibleGame = false
		}
	}
	return gameNumberInt, possibleGame
}

func puzzle1() {

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
		gameNum, possibleGame := checkGame(line)
		if possibleGame {
			sum += gameNum
		}
		// fmt.Println(line)
		// fmt.Println("gameNum:", gameNum, "possibleGame:", possibleGame)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}
	fmt.Println("Sum of possible games:", sum)
}
