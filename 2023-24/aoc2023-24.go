package main

import (
	"fmt"
)

type point struct {
	x float64
	y float64
	z float64
}

type delta struct {
	x float64
	y float64
	z float64
}

type line struct {
	point
	delta
}

func findIntersection(l1, l2 line) (point, bool) {
	if l1.delta.x == 0 && l2.delta.x == 0 {
		return point{}, false
	} else if l1.delta.x == 0 {
		// l1 is vertical
		x := l1.point.x
		y := (x-l2.point.x)*l2.delta.y/l2.delta.x + l2.point.y
		return point{x, y, 0}, true
	} else if l2.delta.x == 0 {
		// l2 is vertical
		x := l2.point.x
		y := (x-l1.point.x)*l1.delta.y/l1.delta.x + l1.point.y
		return point{x, y, 0}, true
	}

	m1 := l1.delta.y / l1.delta.x
	b1 := l1.point.y - m1*l1.point.x
	m2 := l2.delta.y / l2.delta.x
	b2 := l2.point.y - m2*l2.point.x

	if m1 == m2 {
		return point{}, false
	}

	x := (b2 - b1) / (m1 - m2)
	y := m1*x + b1

	return point{x, y, 0}, true
}

func isIntersectionInRange(intersection point, minPos, maxPos float64) bool {
	if intersection.x < minPos || intersection.x > maxPos || intersection.y < minPos || intersection.y > maxPos {
		return false
	}
	return true
}

func sameSign(a, b float64) bool {
	if (a >= 0 && b >= 0) || (a <= 0 && b <= 0) {
		return true
	}
	return false
}

func isIntersectionInFuture(l1 line, l2 line, intersection point) bool {
	l1Ok := false
	l2Ok := false
	sameX := sameSign(intersection.x-l1.point.x, l1.delta.x)
	sameY := sameSign(intersection.y-l1.point.y, l1.delta.y)
	if sameX && sameY {
		l1Ok = true
	}

	sameX = sameSign(intersection.x-l2.point.x, l2.delta.x)
	sameY = sameSign(intersection.y-l2.point.y, l2.delta.y)
	if sameX && sameY {
		l2Ok = true
	}

	return l1Ok && l2Ok
}

func calculateIntrsectionsInRange(lines []line, minPos, maxPos float64) int {
	count := 0
	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			intersection, ok := findIntersection(lines[i], lines[j])
			if !ok {
				continue
			}
			if !isIntersectionInRange(intersection, minPos, maxPos) {
				continue
			}
			if !isIntersectionInFuture(lines[i], lines[j], intersection) {
				continue
			}
			// fmt.Println("Match:", lines[i], lines[j], intersection)
			count++
		}
	}
	return count
}

func puzzle1(lines []line, minPos, maxPos float64) int {
	sum := calculateIntrsectionsInRange(lines, minPos, maxPos)
	return sum
}

func puzzle2(lines []line) int {
	sum := 0
	return sum
}

func main() {
	lines := loadData("input.txt")
	fmt.Println("Puzzle 1: ", puzzle1(lines, 200000000000000.0, 400000000000000.0))
	fmt.Println("Puzzle 2: ", puzzle2(lines))
}
