package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/End3ymion/CADT-Cryptography-Lab/labw2"
	"github.com/End3ymion/CADT-Cryptography-Lab/labw3"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n--- Lab Selection Menu ---")
		fmt.Println("1: Lab Week 2")
		fmt.Println("2: Lab Week 3")
		fmt.Println("0: Exit Program")
		fmt.Print("Enter your choice (0, 1, or 2): ")

		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			fmt.Println("\n--- Starting Lab Week 2 ---")
			labw2.RunLab1()

		case "2":
			fmt.Println("\n--- Starting Lab Week 3 ---")
			labw3.RunLab0()
			labw3.RunLab1()
			
			
		case "0":
			fmt.Println("\nExiting program.")
			return
			
		default:
			fmt.Println("\nInvalid choice. Please enter 0, 1, or 2.")
		}
	}
}
