package main

import (
	"fmt"
)

type workflow struct {
	name  string
	rules []rule
}

type rule struct {
	condition    condition
	action       action
	nextWorkflow string
}

type condition struct {
	categoryIndex int
	comparison    comparison
	value         int
}

func (c condition) String() string {
	return fmt.Sprintf("%s %s %d", partStrings[c.categoryIndex], comparisonStrings[c.comparison], c.value)
}

type comparison int

func (c comparison) opposite() comparison {
	switch c {
	case LessThan:
		return MoreThanOrEqual
	case MoreThan:
		return LessThanOrEqual
	case LessThanOrEqual:
		return MoreThan
	case MoreThanOrEqual:
		return LessThan
	case Always:
		return Always
	}
	panic("Error getting opposite comparison")
}

const (
	LessThan comparison = iota
	MoreThan
	LessThanOrEqual
	MoreThanOrEqual
	Always
)

var comparisonValues = map[string]int{
	"<":  int(LessThan),
	"<=": int(LessThanOrEqual),
	">":  int(MoreThan),
	">=": int(MoreThanOrEqual),
}

var comparisonStrings = map[comparison]string{
	LessThan:        "<",
	LessThanOrEqual: "<=",
	MoreThan:        ">",
	MoreThanOrEqual: ">=",
}

type action int

const (
	Rejected action = iota
	Accepted
	NextWorkflow
)

type part [4]int

var partValues = map[string]int{
	"x": 0,
	"m": 1,
	"a": 2,
	"s": 3,
}

var partStrings = map[int]string{
	0: "x",
	1: "m",
	2: "a",
	3: "s",
}

type turn int

const (
	Left turn = iota
	Right
)

func processAction(rule rule, part part) (bool, bool, int) {
	switch rule.action {
	case Rejected:
		return true, false, 0
	case Accepted:
		return true, false, part[0] + part[1] + part[2] + part[3]
	case NextWorkflow:
		return false, true, 0
	}
	panic("Error processing action")
}

func processRule(rule rule, part part) (bool, bool, int) {
	switch rule.condition.comparison {
	case LessThan:
		if part[rule.condition.categoryIndex] < rule.condition.value {
			return processAction(rule, part)
		}
	case MoreThan:
		if part[rule.condition.categoryIndex] > rule.condition.value {
			return processAction(rule, part)
		}
	case Always:
		return processAction(rule, part)
	}
	return false, false, 0
}

func getPartValue(workflows map[string][]rule, part part) int {
	workflowName := "in"
	workflow := workflows[workflowName]
	for {
		for _, rule := range workflow {
			done, nextWorkflow, value := processRule(rule, part)
			workflowName = rule.nextWorkflow
			if done {
				return value
			}
			if nextWorkflow {
				break
			}
		}
		workflow = workflows[workflowName]
	}
}

func puzzle1(workflows map[string][]rule, parts []part) int {
	sum := 0
	for _, part := range parts {
		sum += getPartValue(workflows, part)
		fmt.Println(part, sum)
	}
	return sum
}

func recordPath(path *Stack, acceptedPaths *[][]condition) {
	acceptedPath := make([]condition, len(*path))
	copy(acceptedPath, *path)
	*acceptedPaths = append(*acceptedPaths, acceptedPath)
}

func recordOppositeAndTracePaths(workflows map[string][]rule, workflowName string, ruleIndex int, path *Stack, acceptedPaths *[][]condition) {
	rules := workflows[workflowName]
	condition := rules[ruleIndex].condition
	condition.comparison = condition.comparison.opposite()
	path.Push(condition)
	tracePaths(workflows, workflowName, ruleIndex+1, path, acceptedPaths)
	path.Pop()
}

func tracePaths(workflows map[string][]rule, workflowName string, ruleIndex int, path *Stack, acceptedPaths *[][]condition) {
	rules := workflows[workflowName]
	if rules[ruleIndex].condition.comparison == LessThan || rules[ruleIndex].condition.comparison == MoreThan {
		switch rules[ruleIndex].action {
		case Rejected:
			// record opposite condition and trace that path
			recordOppositeAndTracePaths(workflows, workflowName, ruleIndex, path, acceptedPaths)
			return
		case Accepted:
			// record condition and record path to "Accepted"
			condition := rules[ruleIndex].condition
			path.Push(condition)
			recordPath(path, acceptedPaths) // record path
			path.Pop()
			// record opposite condition and trace that path
			recordOppositeAndTracePaths(workflows, workflowName, ruleIndex, path, acceptedPaths)
			return
		case NextWorkflow:
			// record condition and trace that path
			condition := rules[ruleIndex].condition
			path.Push(condition)
			tracePaths(workflows, rules[ruleIndex].nextWorkflow, 0, path, acceptedPaths)
			path.Pop()
			// record opposite condition and trace that path
			recordOppositeAndTracePaths(workflows, workflowName, ruleIndex, path, acceptedPaths)
			return
		default:
			panic("We should never get here")
		}
	} else if rules[ruleIndex].condition.comparison == Always {
		switch rules[ruleIndex].action {
		case Rejected:
			return
		case Accepted:
			// record path to "Accepted" and return
			recordPath(path, acceptedPaths)
			return
		case NextWorkflow:
			tracePaths(workflows, rules[ruleIndex].nextWorkflow, 0, path, acceptedPaths)
			return
		default:
			panic("We should never get here")
		}
	} else {
		panic("We should never get here")
	}
}

func getRanges(acceptedPaths [][]condition) int {
	sum := 0
	for _, path := range acceptedPaths {
		pathLimits := [][]int{
			{1, 4000},
			{1, 4000},
			{1, 4000},
			{1, 4000},
		}
		for _, condition := range path {
			switch condition.comparison {
			case LessThan:
				if condition.value < pathLimits[condition.categoryIndex][1] {
					pathLimits[condition.categoryIndex][1] = condition.value - 1
				}
			case MoreThan:
				if condition.value > pathLimits[condition.categoryIndex][0] {
					pathLimits[condition.categoryIndex][0] = condition.value + 1
				}
			case LessThanOrEqual:
				if condition.value <= pathLimits[condition.categoryIndex][1] {
					pathLimits[condition.categoryIndex][1] = condition.value
				}
			case MoreThanOrEqual:
				if condition.value >= pathLimits[condition.categoryIndex][0] {
					pathLimits[condition.categoryIndex][0] = condition.value
				}
			}
		}
		pathSum := (pathLimits[0][1] - pathLimits[0][0] + 1) * (pathLimits[1][1] - pathLimits[1][0] + 1) * (pathLimits[2][1] - pathLimits[2][0] + 1) * (pathLimits[3][1] - pathLimits[3][0] + 1)
		sum += pathSum
	}
	return sum
}

func puzzle2(workflows map[string][]rule) int {
	path := make(Stack, 0)
	acceptedPaths := make([][]condition, 0)
	workflowName := "in"
	ruleIndex := 0
	tracePaths(workflows, workflowName, ruleIndex, &path, &acceptedPaths)
	combinations := getRanges(acceptedPaths)
	return combinations
}

func main() {
	// workflows, parts := loadData("input-test.txt")
	workflows, parts := loadData("input.txt")
	fmt.Println("Puzzle 1: ", puzzle1(workflows, parts))
	fmt.Println("Puzzle 2: ", puzzle2(workflows))
}
