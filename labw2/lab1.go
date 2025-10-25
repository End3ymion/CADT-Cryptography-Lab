
package main
import (
	"bufio"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"strings"
)

var taskCounter int = 1

func dashHeader() {
	total := 50
	taskText := fmt.Sprintf(" Task %d ", taskCounter)
	mid := total/2 - len(taskText)/2

	for i := 0; i < total; i++ {
		if i == mid {
			fmt.Print(taskText)
			i += len(taskText) - 1
		} else {
			fmt.Print("-")
		}
	}
	fmt.Println()
	taskCounter++
}

func addValue(a, b float64) float64 { return a + b }
func subVal(a, b float64) float64 { return a - b }
func mulVal(a, b float64) float64 { return a * b }

func divVal(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func modVal(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("modulo by zero")
	}
	return float64(int(a) % int(b)), nil
}

func calculator(a, b float64, op int) (float64, error) {
	switch op {
	case 1:
		return addValue(a, b), nil
	case 2:
		return subVal(a, b), nil
	case 3:
		return mulVal(a, b), nil
	case 4:
		return divVal(a, b)
	case 5:
		return modVal(a, b)
	default:
		return 0, errors.New("invalid operator")
	}
}

func myAND(a, b int) int { return a & b }
func myOR(a, b int) int { return a | b }
func myXOR(a, b int) int { return a ^ b }

func binConv(s string) string {
	var result string
	for _, b := range []byte(s) {
		result += fmt.Sprintf("%08b ", b)
	}
	return result
}

func decConv(s string) string {
	var result string
	for _, b := range []byte(s) {
		result += fmt.Sprintf("%d ", b)
	}
	return result
}

func b64Conv(s string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(s))
	return encoded
}

func xorEncrypt(s, key string) string {
	data := []byte(s)
	k := []byte(key)
	result := make([]byte, len(data))
	for i := 0; i < len(data); i++ {
		result[i] = data[i] ^ k[i%len(k)]
	}
	return base64.StdEncoding.EncodeToString(result)
}

func main() {
	dashHeader()
	fmt.Println("\nTask 1: Write a Go program to demonstrate the use of assignment operators. The program should take two integer inputs and perform various assignment operations such as =, +=, -=, *=, /=, and %=. Display the result after each operation.\n")
	fmt.Println("Answer:\n")

	var number int = 10
	number += 1
	fmt.Println(number)
	number -= 1
	fmt.Println(number)
	number *= 1
	fmt.Println(number)
	number /= 1
	fmt.Println(number)
	number %= 1
	fmt.Println(number)
	fmt.Println()

	dashHeader()
	fmt.Println("\nTask2: Write a Go program to demonstrate the use of logical operators such as &&, ||, and !. The program should take two integer inputs and evaluate logical expressions like:\n\t• both positive (&&)\n\t• one greater than the other (||)\n\t• not equal (!)\nDisplay whether each condition is true or false.\n")
	fmt.Println("Answer:\n")

	var number1 int = 1
	var number2 int = 1
	fmt.Println(number1 > 0 && number2 > 0)
	fmt.Println((number1 > number2) || (number2 > number1))
	fmt.Println(number1 != number2)
	fmt.Println()

	dashHeader()
	fmt.Println("\nTask3: Write a Go program to demonstrate the use of bitwise and assignment operators. The program should perform AND, OR, XOR, NOT, left shift, and right shift operations on two integers using functions and display the results. Also, show the effect of assignment operators on variable values.\n")
	fmt.Println("Answer:\n")

	a, b := 1, 0
	fmt.Printf("Bitwise AND: %d & %d = %d\n", a, b, myAND(a, b))
	fmt.Printf("Bitwise OR: %d | %d = %d\n", a, b, myOR(a, b))
	fmt.Printf("Bitwise XOR: %d ^ %d = %d\n", a, b, myXOR(a, b))
	fmt.Println()

	dashHeader()
	fmt.Println("\nTask 4:  Write a Go program to implement a mini calculator using functions. The program should read two numbers and an operator from the user, perform the operation, and return/print the result. Support at least the operators: +, -, *, /, %, Handle division by zero and invalid operators gracefully. (Optional) Make it menu-driven and repeat until the user chooses Exit.\n")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n===== Mini Calculator =====")
		fmt.Println("Select operation:")
		fmt.Println("1: Addition")
		fmt.Println("2: Subtraction")
		fmt.Println("3: Multiplication")
		fmt.Println("4: Division")
		fmt.Println("5: Modulo")
		fmt.Println("0: Exit")
		fmt.Println("===========================\n")

		fmt.Print("Enter your choice: ")
		scanner.Scan()
		opInput := strings.TrimSpace(scanner.Text())
		var op int
		_, err := fmt.Sscanf(opInput, "%d", &op)
		if err != nil || op < 0 || op > 5 {
			fmt.Println("\nInvalid input, enter a number 0-5\n")
			continue
		}
		if op == 0 {
			fmt.Println("\nExiting calculator...\n")
			break
		}

		var a, b float64

		for {
			fmt.Print("Enter first number: ")
			scanner.Scan()
			input := strings.TrimSpace(scanner.Text())
			_, err := fmt.Sscanf(input, "%f", &a)
			if err != nil {
				fmt.Println("Please enter a number.")
				continue
			}
			break
		}

		for {
			fmt.Print("Enter second number: ")
			scanner.Scan()
			input := strings.TrimSpace(scanner.Text())
			_, err := fmt.Sscanf(input, "%f", &b)
			if err != nil {
				fmt.Println("Please enter a number.")
				continue
			}
			break
		}
		result, err := calculator(a, b, op)
		if err != nil {
			fmt.Println("\nError:", err, "\n")
			continue
		}

		var opSymbol string
		switch op {
		case 1:
			opSymbol = "+"
		case 2:
			opSymbol = "-"
		case 3:
			opSymbol = "*"
		case 4:
			opSymbol = "/"
		case 5:
			opSymbol = "%"
		}

		fmt.Println()
		fmt.Printf("%v %s %v = %v\n", a, opSymbol, b, result)
		fmt.Println()
	}

	dashHeader()
	fmt.Println("\nTask 5: Write a Go program to demonstrate how text data can be represented in binary, hexadecimal, and Base64 formats. The program should take a string input and display its equivalent binary, hexadecimal, and Base64 representations\n")

	var input string

	for {
		fmt.Print("Enter string for Binary Format: ")
		scanner.Scan()
		input = strings.TrimSpace(scanner.Text())
		if len(input) == 0 {
			fmt.Println("Please enter a string.")
			continue
		}
		fmt.Println("\nBinary Format:", binConv(input), "\n")
		break
	}

	for {
		fmt.Print("Enter string for Decimal Format: ")
		scanner.Scan()
		input = strings.TrimSpace(scanner.Text())
		if len(input) == 0 {
			fmt.Println("Please enter a string.")
			continue
		}
		fmt.Println("\nDecimal Format:", decConv(input), "\n")
		break
	}

	for {
		fmt.Print("Enter string for Base64 Format: ")
		scanner.Scan()
		input = strings.TrimSpace(scanner.Text())
		if len(input) == 0 {
			fmt.Println("Please enter a string.")
			continue
		}
		fmt.Println("\nBase64 Format:", b64Conv(input), "\n")
		break
	}
	dashHeader()
	fmt.Println("\nTask 6: Write a Go program to perform simple XOR encryption and decryption. The program should take a plaintext message and a key (single byte or repeating key) and produce a ciphertext by XORing them. Use the same function to decrypt the ciphertext to recover the original message.\n")

	for {
		scanner := bufio.NewScanner(os.Stdin)
		var lines []string
		fmt.Println("Enter string (Ctrl+D to end): ")
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			return
		}
		input = strings.Join(lines, "\n")
		if len(strings.TrimSpace(input)) == 0 {
			fmt.Println("Please enter a string.")
			continue
		}
		break
	}

	var key string
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter key: ")
		scanner.Scan()
		key = strings.TrimSpace(scanner.Text())
		if len(key) == 0 {
			fmt.Println("Error: key cannot be empty")
			continue
		}
		break
	}

	encrypted := xorEncrypt(input, key)
	fmt.Println("\nEncrypted (Base64):", encrypted)

	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		fmt.Println("Decryption Error (Base64 Decode):", err)
		return
	}

	decrypted := xorEncrypt(string(ciphertext), key)
	finalDecrypted, err := base64.StdEncoding.DecodeString(decrypted)
	if err != nil {
		fmt.Println("Decryption Error (Final Base64 Decode):", err)
		return
	}

	fmt.Print("\nDecrypted:\n", string(finalDecrypted))

}
