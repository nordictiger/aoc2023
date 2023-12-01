package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // Open the file
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close() // Ensure the file is closed after function returns

    // Create a new scanner to read the file
    scanner := bufio.NewScanner(file)

    // Read the file line by line
    for scanner.Scan() {
        line := scanner.Text()

        // Process the line (for now, just print it)
        fmt.Println(line)
    }

    // Check for errors during scanning
    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading from file:", err)
    }
}

