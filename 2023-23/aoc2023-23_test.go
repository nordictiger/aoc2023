package main

import (
	"fmt"
	"testing"
)

func TestData(t *testing.T) {
	hikingMap := loadData("input-test.txt")
	for _, line := range hikingMap {
		fmt.Println(string(line))
	}
}

func TestPuzzle1(t *testing.T) {
	hikingMap := loadData("input-test.txt")
	fmt.Println("Puzzle 1: ", puzzle1(hikingMap))
}

func TestPuzzle2(t *testing.T) {
	hikingMap := loadData("input-test.txt")
	fmt.Println("Puzzle 2: ", puzzle2(hikingMap))
}
