package main

import "fmt"

func checkHorizontal(pattern [][]byte, oldH int) int {
	pattern_width := len(pattern[0])
	pattern_height := len(pattern)

	for i := 0; i < pattern_height-1; i++ {
		k := 0
		line_match := true
		for {
			line_match = true
			for j := 0; j < pattern_width; j++ {
				if pattern[i-k][j] != pattern[i+k+1][j] {
					line_match = false
					break
				}
			}
			if line_match && i-k > 0 && i+k+1 < pattern_height-1 {
				k++
			} else {
				break
			}
		}
		if line_match && (i-k == 0 || i+k+1 == pattern_height-1) {
			if (i+1)*100 != oldH {
				return (i + 1) * 100
			}
		}
	}
	return 0
}

func checkVertical(pattern [][]byte, oldV int) int {
	pattern_width := len(pattern[0])
	pattern_height := len(pattern)

	for i := 0; i < pattern_width-1; i++ {
		k := 0
		line_match := true
		for {
			line_match = true
			for j := 0; j < pattern_height; j++ {
				if pattern[j][i-k] != pattern[j][i+k+1] {
					line_match = false
					break
				}
			}
			if line_match && i-k > 0 && i+k+1 < pattern_width-1 {
				k++
			} else {
				break
			}
		}
		if line_match && (i-k == 0 || i+k+1 == pattern_width-1) {
			if (i + 1) != oldV {
				return i + 1
			}
		}
	}
	return 0
}

func getReflections(pattern [][]byte, oldV int, oldH int) (int, int) {

	vertical := checkVertical(pattern, oldV)
	horizontal := checkHorizontal(pattern, oldH)
	return vertical, horizontal
}

func puzzle1(patterns [][][]byte) int {
	sum := 0
	for _, pattern := range patterns {
		v, h := getReflections(pattern, 0, 0)
		sum += max(v, h)
	}
	return sum
}

func getNewReflection(v int, h int, pattern [][]byte) int {

	for i := 0; i < len(pattern); i++ {
		for j := 0; j < len(pattern[0]); j++ {
			value := pattern[i][j]
			if pattern[i][j] == '#' {
				pattern[i][j] = '.'
			} else {
				pattern[i][j] = '#'
			}

			v2, h2 := getReflections(pattern, v, h)
			if v2 != 0 || h2 != 0 {
				return max(v2, h2)
			}
			pattern[i][j] = value
		}
	}
	return 0
}

func puzzle2(patterns [][][]byte) int {
	sum := 0
	for p, pattern := range patterns {
		v, h := getReflections(pattern, 0, 0)
		newReflection := getNewReflection(v, h, pattern)
		sum += newReflection
		fmt.Println("Pattern: ", p, " old: ", max(v, h), " new: ", newReflection, " sum: ", sum)
	}
	return sum
}

func main() {
	// patterns, _ := loadData("input-test.txt")
	// patterns, _ := loadData("input-test2.txt")
	patterns, _ := loadData("input.txt")
	fmt.Println("Puzzle 1: ", puzzle1(patterns))
	fmt.Println("Puzzle 2: ", puzzle2(patterns))
}
