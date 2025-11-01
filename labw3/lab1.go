package labw3

import (
	"bufio"
	"fmt"
	"os"
	"encoding/hex"
	"strings" 
	"github.com/End3ymion/CADT-Cryptography-Lab/labw3/utils/crack"
)

func printVerboseAttempt(candidate string) string {
    output := fmt.Sprintf("Trying: %s\n", candidate)
    fmt.Print(output) 
    return output
}

func writeLogToFile(logContent string) {
    err := os.WriteFile("labw3/verbose.txt", []byte(logContent), 0644)
    if err != nil {
        fmt.Printf("\nError saving log file: %v\n", err)
    }
}

func RunLab1() {
    scanner := bufio.NewScanner(os.Stdin)
	
    for {
        var hashType string
        var expectedLen int
        var defaultHashStr string

        fmt.Println("\n=== Hash Type Selector ===")
        fmt.Println("1) MD5 (32 hex chars)")
        fmt.Println("2) SHA1 (40 hex chars)")
        fmt.Println("3) SHA512 (128 hex chars)")
        fmt.Println("0) Exit Hash Cracker")
        fmt.Print("Select hash type (0, 1, 2, or 3): ")
        scanner.Scan()
        selection := strings.TrimSpace(scanner.Text())
        
        switch selection {
        case "1":
            hashType = "MD5"
            expectedLen = 16
            defaultHashStr = "6a85dfd77d9cb35770c9dc6728d73d3f"
        case "2":
            hashType = "SHA1"
            expectedLen = 20
            defaultHashStr = "aa1c7d931cf140bb35a5a16adeb83a551649c3b9"
        case "3":
            hashType = "SHA512"
            expectedLen = 64
            defaultHashStr = "485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38"
        case "0":
            fmt.Println("\nExiting Hash Cracker.")
            return
        default:
            fmt.Println("Invalid selection. Please enter 0, 1, 2, or 3.")
            continue
        }

        fmt.Printf("Selected Hash Type: %s\n", hashType)
        
        var targetHashStr = defaultHashStr
        
        fmt.Print("Enter target hash (enter for default): ")
        scanner.Scan()
        
        userInput := strings.TrimSpace(scanner.Text())

        if userInput != "" {
            targetHashStr = userInput
        }

        targetHashSlice, err := hex.DecodeString(targetHashStr)
        if err != nil || len(targetHashSlice) != expectedLen {
            fmt.Println("Error: Target hash must be a valid hex string. Check your input or selected type.\n")
            continue
        }
        
        var fixedTargetHash []byte = targetHashSlice
        
        var verboseMode bool = false
        var logBuffer string = "--- Starting Verbose Log ---\n"
        
        fmt.Print("Verbose Mode (y/n): ")
        scanner.Scan()
        userInput = strings.TrimSpace(scanner.Text())
        
        if strings.ToLower(userInput) == "y" {
            verboseMode = true
            fmt.Println("Verbose mode ENABLED.")
        } else {
            fmt.Println("Verbose mode DISABLED.")
        }

        f, err := os.Open("labw3/utils/crack/nord_vpn.txt")
        if err != nil {
            errMsg := fmt.Sprintf("Error: Could not open wordlist file: %v\n", err)
            if verboseMode {
                logBuffer += errMsg
                writeLogToFile(logBuffer)
            }
            fmt.Print(errMsg)
            continue
        }
        defer f.Close()

        fileScanner := bufio.NewScanner(f)
        
        var found bool = false

        for fileScanner.Scan() {
            candidate := strings.TrimSpace(fileScanner.Text())
            if candidate == "" {
                continue
            }
            
            if verboseMode {
                logBuffer += printVerboseAttempt(candidate)
            }
            
            var cracked string
            
            switch hashType {
            case "SHA1":
                var targetArray [20]byte
                copy(targetArray[:], fixedTargetHash)
                cracked = crack.SHA1Crack(candidate, targetArray)
            case "SHA512":
                var targetArray [64]byte
                copy(targetArray[:], fixedTargetHash)
                cracked = crack.SHA512Crack(candidate, targetArray)
            case "MD5":
                var targetArray [16]byte
                copy(targetArray[:], fixedTargetHash)
                cracked = crack.MD5Crack(candidate, targetArray)
            }

            if cracked != "" {
                found = true
                fmt.Printf("\033[92mCRACKED: %s\033[0m\n", cracked)
                if verboseMode {
                    logBuffer += fmt.Sprintf("\nCRACKED: %s\n", cracked)
                }
                break
            }
        }

        if err := fileScanner.Err(); err != nil {
            errMsg := fmt.Sprintf("Error reading wordlist: %v\n", err)
            if verboseMode {
                logBuffer += errMsg
            }
            fmt.Print(errMsg)
        } else if !found {
            if verboseMode {
                logBuffer += "\nNot found in wordlist.\n"
            }
            fmt.Println("Not found in wordlist.")
        }
        
        if verboseMode {
            writeLogToFile(logBuffer)
            fmt.Println("Verbose output saved to verbose.txt")
        }
    }
}
