package main

import (
	"fmt"
	"strconv"
	"strings"
)

type lensStack []lens

type lens struct {
	label string
	focal int
}

func (s *lensStack) Add(label string, focal int) {
	for i, lens := range *s {
		if lens.label == label {
			(*s)[i].focal = focal
			return
		}
	}
	v := lens{label: label, focal: focal}
	*s = append(*s, v)
}

func (s *lensStack) Remove(label string) {
	for i, lens := range *s {
		if lens.label == label {
			*s = append((*s)[:i], (*s)[i+1:]...)
			return
		}
	}
}

func parseStep(step string) (string, string, int) {
	command := ""
	label := ""
	focal := 0
	stepParts := strings.Split(step, "=")
	if len(stepParts) == 1 {
		byteArray := []byte(step)
		command = "remove"
		label = string(byteArray[:len(byteArray)-1])
		focal = 0
		return command, label, focal
	}
	command = "add"
	label = stepParts[0]
	num, err := strconv.Atoi(stepParts[1])
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	}
	focal = num
	return command, label, focal
}

func calcHash(step string) int {
	hashValue := 0
	for i := range step {
		asciiValue := int(step[i])
		hashValue += asciiValue
		hashValue *= 17
		hashValue %= 256
	}
	return hashValue
}

func puzzle1(data []string) int {
	sum := 0
	for _, v := range data {
		sum += calcHash(v)
	}
	return sum
}

func puzzle2(data []string) int {
	sum := 0
	boxes := make([]lensStack, 0)
	for i := 0; i < 256; i++ {
		boxes = append(boxes, make(lensStack, 0))
	}

	for _, v := range data {
		command, label, focal := parseStep(v)
		fmt.Println(command, label, focal)
		switch command {
		case "add":
			boxes[calcHash(label)].Add(label, focal)
		case "remove":
			boxes[calcHash(label)].Remove(label)
		}
	}
	for index, box := range boxes {
		if len(box) != 0 {
			fmt.Print(index, ": ")
			for order, lens := range box {
				fmt.Print(lens.label, lens.focal, " ")
				focusingPower := (index + 1) * (order + 1) * lens.focal
				fmt.Print(focusingPower, " ")
				sum += focusingPower
			}
			fmt.Println()
		}
	}

	return sum
}

func main() {
	// sequence, _ := loadData("input-test.txt")
	sequence, _ := loadData("input.txt")
	fmt.Println("Puzzle 1: ", puzzle1(sequence))
	fmt.Println("Puzzle 2: ", puzzle2(sequence))
}
