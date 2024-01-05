package main

import (
	"fmt"
	"testing"
)

func TestData(t *testing.T) {
	lines := loadData("input-test.txt")
	fmt.Println(lines)
}

func TestPuzzle1(t *testing.T) {
	lines := loadData("input-test.txt")
	fmt.Println("Puzzle 1: ", puzzle1(lines, 7.0, 27.0))
}

func TestPuzzle2(t *testing.T) {
	lines := loadData("input-test.txt")
	fmt.Println("Puzzle 2: ", puzzle2(lines))
}
