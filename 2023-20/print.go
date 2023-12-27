package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func contains[T comparable](slice []T, val T) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func printConfiguration(mc moduleConfiguration, ignoreList []string) {
	if contains(ignoreList, "broadcaster") {
		return
	}
	visited := make(map[string]bool)
	printConfigurationNode(mc, "", "broadcaster", ignoreList, 0, &visited)
}

func printConfigurationNode(mc moduleConfiguration, previousKey, moduleKey string, ignoreList []string, level int, visisted *map[string]bool) {
	for i := 0; i < level*2; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%s%s %v", mc[moduleKey].moduleType, moduleKey, mc[moduleKey].state)
	if mc[moduleKey].moduleType == Conjunction {
		fmt.Printf(" - %s", mc[moduleKey].incoming)
	}
	fmt.Println()
	(*visisted)[previousKey+moduleKey] = true
	for _, node := range mc[moduleKey].outgoing {
		if !(*visisted)[moduleKey+node] && !contains(ignoreList, node) {
			printConfigurationNode(mc, moduleKey, node, ignoreList, level+1, visisted)
		}
	}
}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
