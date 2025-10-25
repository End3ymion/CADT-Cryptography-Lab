package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter multi-line text (Ctrl+D or Ctrl+Z to end):")

	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		return
	}

	fmt.Println("\nYour input:")
	for i, line := range lines {
		fmt.Printf("Line %d: %s\n", i+1, line)
	}
}
