package main

import (
	"fmt"
	"testing"
)

func TestPuzzle1(t *testing.T) {
	bricks := loadData("input-test.txt")
	fmt.Println("Puzzle 1: ", puzzle1(bricks))
}

func TestPuzzle2(t *testing.T) {
	bricks := loadData("input-test.txt")
	fmt.Println("Puzzle 2: ", puzzle2(bricks))
}
