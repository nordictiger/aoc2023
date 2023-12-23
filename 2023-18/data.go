package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func loadData(fileName string) ([]instruction, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	data := make([]instruction, 0)
	re := regexp.MustCompile(`([URDL]) (\d+) \((.+)\)`)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading from file:", err)
			return nil, err
		}
		matches := re.FindStringSubmatch(scanner.Text())
		if len(matches) != 4 {
			continue
		}
		meters, err := strconv.Atoi(matches[2])
		if err != nil {
			fmt.Println("Error converting to int:", err)
			return nil, err
		}
		data = append(data, instruction{matches[1], meters, matches[3]})
	}
	return data, nil

}
