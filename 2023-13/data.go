package main

import (
	"bufio"
	"fmt"
	"os"
)

func loadData(fileName string) ([][][]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	data := make([][][]byte, 0)
	pattern := make([][]byte, 0)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading from file:", err)
			return nil, err
		}
		if len(scanner.Text()) == 0 {
			data = append(data, pattern)
			pattern = make([][]byte, 0)
			continue
		}
		line := make([]byte, 0)
		for i := 0; i < len(scanner.Text()); i++ {
			line = append(line, scanner.Text()[i])
		}
		pattern = append(pattern, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return nil, err
	} else {
		data = append(data, pattern)
	}
	return data, nil
}
