package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var rulesRe = regexp.MustCompile(`(\w+){(.+)}`)
var singleRuleRe = regexp.MustCompile(`([xmas])([<>])(\d+):([RA]|\w+)|([RA]|\w+)`)

var partRe = regexp.MustCompile(`{(.+)}`)
var singlePartRe = regexp.MustCompile(`([xmas])=(\d+)`)

func getActionForRule(actionString string, rule *rule) {
	switch actionString {
	case "R":
		rule.action = Rejected
	case "A":
		rule.action = Accepted
	default:
		rule.action = NextWorkflow
		rule.nextWorkflow = actionString
	}
}

func getWorkflow(line string) workflow {
	rules := make([]rule, 0)
	match := rulesRe.FindStringSubmatch(line)
	if len(match) != 3 {
		panic("Error parsing line: " + line)
	}
	rulesStrings := strings.Split(match[2], ",")
	for _, ruleString := range rulesStrings {
		var rule rule

		rulePieces := singleRuleRe.FindStringSubmatch(ruleString)

		if len(rulePieces[5]) != 0 {
			rule.condition.comparison = Always
			getActionForRule(rulePieces[5], &rule)
		} else if len(rulePieces[1]) != 0 {
			rule.condition.categoryIndex = partValues[rulePieces[1]]
			rule.condition.comparison = comparison(comparisonValues[rulePieces[2]])
			value, err := strconv.Atoi(rulePieces[3])
			if err != nil {
				panic("Error converting to int: " + rulePieces[3])
			}
			rule.condition.value = value
			getActionForRule(rulePieces[4], &rule)
		} else {
			panic("Error parsing rule: " + ruleString)
		}
		rules = append(rules, rule)
	}
	return workflow{match[1], rules}
}

func getPart(line string) part {
	partString := partRe.FindStringSubmatch(line)
	if len(partString) != 2 {
		panic("Error parsing line: " + line)
	}
	partPieces := strings.Split(partString[1], ",")
	if len(partPieces) != 4 {
		panic("Error parsing line: " + line)
	}
	var part part
	for _, partPiece := range partPieces {
		assignment := singlePartRe.FindStringSubmatch(partPiece)
		value, err := strconv.Atoi(assignment[2])
		if err != nil {
			panic("Error converting to int: " + assignment[2])
		}
		part[partValues[assignment[1]]] = value
	}
	return part
}

func loadData(fileName string) (map[string][]rule, []part) {
	file, err := os.Open(fileName)
	if err != nil {
		panic("Error opening file: " + fileName + " " + err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	workflows := make(map[string][]rule)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic("Error reading from file: " + err.Error())
		}
		if scanner.Text() == "" {
			break
		}
		workflow := getWorkflow(scanner.Text())
		workflows[workflow.name] = workflow.rules
	}

	parts := make([]part, 0)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic("Error reading from file: " + err.Error())
		}
		part := getPart(scanner.Text())
		parts = append(parts, part)
	}
	return workflows, parts
}
