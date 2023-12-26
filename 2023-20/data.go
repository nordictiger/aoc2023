package main

import (
	"bufio"
	"os"
	"strings"
)

func loadData(fileName string) moduleConfiguration {
	file, err := os.Open(fileName)
	if err != nil {
		panic("Error opening file: " + err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	data := make(moduleConfiguration, 0)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic("Error reading from file: " + err.Error())
		}
		parts := strings.Split(scanner.Text(), " -> ")
		var m moduleType
		key := ""
		switch parts[0][0] {
		case '%':
			m = FlipFlop
			key = parts[0][1:]
		case '&':
			m = Conjunction
			key = parts[0][1:]
		default:
			m = Broadcaster
			key = parts[0]
		}
		outgoing := make(map[string]state, 0)
		data[key] = node{key, m, Low, outgoing, strings.Split(parts[1], ", ")}
	}
	return data
}
