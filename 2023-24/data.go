package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func loadData(fileName string) []line {
	file, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("Error opening file %s: %v", fileName, err))
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := make([]line, 0)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic(fmt.Sprintf("Error reading from file %s: %v", fileName, err))
		}
		trimmedLine := strings.ReplaceAll(scanner.Text(), " ", "")
		parts := strings.Split(trimmedLine, "@")

		position := strings.Split(parts[0], ",")
		x1, _ := strconv.ParseFloat(position[0], 64)
		y1, _ := strconv.ParseFloat(position[1], 64)
		z1, _ := strconv.ParseFloat(position[2], 64)

		velocity := strings.Split(parts[1], ",")
		x2, _ := strconv.ParseFloat(velocity[0], 64)
		y2, _ := strconv.ParseFloat(velocity[1], 64)
		z2, _ := strconv.ParseFloat(velocity[2], 64)

		lines = append(lines, line{point{x: x1, y: y1, z: z1}, delta{x: x2, y: y2, z: z2}})
	}
	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading from file %s: %v", fileName, err))
	}
	return lines
}
