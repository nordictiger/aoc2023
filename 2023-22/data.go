package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func loadData(fileName string) []brick {
	file, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("Error opening file %s: %v", fileName, err))
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	data := make([]brick, 0)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic(fmt.Sprintf("Error reading from file %s: %v", fileName, err))
		}
		ends := strings.Split(scanner.Text(), "~")

		start := strings.Split(ends[0], ",")
		x1, _ := strconv.Atoi(start[0])
		y1, _ := strconv.Atoi(start[1])
		z1, _ := strconv.Atoi(start[2])

		end := strings.Split(ends[1], ",")
		x2, _ := strconv.Atoi(end[0])
		y2, _ := strconv.Atoi(end[1])
		z2, _ := strconv.Atoi(end[2])

		brick := brick{start: point3d{x: x1, y: y1, z: z1}, end: point3d{x: x2, y: y2, z: z2}}
		if brick.end.z < brick.start.z {
			brick.start.z, brick.end.z = brick.end.z, brick.start.z
			fmt.Printf("Swapped brick start-end z %v\n", brick)
		}
		data = append(data, brick)
	}
	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading from file %s: %v", fileName, err))
	}
	return data
}
