package main

import (
	"fmt"
	"os"
	"strings"
)

func loadData(fileName string) ([]string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}
	text := string(data)

	result := strings.Split(text, ",")

	return result, nil
}
