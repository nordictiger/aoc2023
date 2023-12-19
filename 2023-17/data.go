package main

import (
	"bufio"
	"fmt"
	"os"
)

func loadData(fileName string) ([][]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	data := make([][]int, 0)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading from file:", err)
			return nil, err
		}
		line := make([]int, 0)
		for i := 0; i < len(scanner.Text()); i++ {
			line = append(line, int(scanner.Text()[i]-'0'))
		}
		data = append(data, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return nil, err
	}
	return data, nil
}
