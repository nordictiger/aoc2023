package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func printGarden(garden [][]byte) {
	for i := 0; i < len(garden); i++ {
		for j := 0; j < len(garden[i]); j++ {
			fmt.Printf("%c", garden[i][j])
		}
		fmt.Println()
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
