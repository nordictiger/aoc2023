package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func loadData(fileName string) map[string][]string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("Error opening file %s: %v", fileName, err))
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	diagram := make(map[string][]string)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic(fmt.Sprintf("Error reading from file %s: %v", fileName, err))
		}
		parts := strings.Split(scanner.Text(), ": ")
		diagram[parts[0]] = strings.Split(parts[1], " ")
	}
	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading from file %s: %v", fileName, err))
	}
	return diagram
}
