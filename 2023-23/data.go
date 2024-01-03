package main

import (
	"bufio"
	"fmt"
	"os"
)

func loadData(fileName string) [][]byte {
	file, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("Error opening file %s: %v", fileName, err))
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	data := make([][]byte, 0)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic(fmt.Sprintf("Error reading from file %s: %v", fileName, err))
		}
		line := scanner.Text()
		data = append(data, []byte(line))
	}
	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading from file %s: %v", fileName, err))
	}
	return data
}
