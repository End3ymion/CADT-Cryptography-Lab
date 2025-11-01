package labw3

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/sha3"
	"fmt"
	"os"
)

func dashHeader() {
	total := 50
	for i := 0; i < total; i++ {
		fmt.Print("=")
	}
	fmt.Println()
}

func md5Hash(s string) [16]byte {
	return md5.Sum([]byte(s))
}

func sha1Hash(s string) [20]byte {
	return sha1.Sum([]byte(s))
}

func sha256Hash(s string) [32]byte {
	return sha256.Sum256([]byte(s))
}

func sha512Hash(s string) [64]byte {
	return sha512.Sum512([]byte(s))
}

func sha3Hash(s string) [32]byte {
	return sha3.Sum256([]byte(s))
}

func compareHash(in1, in2 []byte) string {
	if string(in1) == string(in2) {
		return "Match!"
	}
	return "Not Match!"
}

func RunLab0() {
	scanner := bufio.NewScanner(os.Stdin)

	dashHeader()
	fmt.Println("Proof the Hash Program")
	fmt.Println("Step 1: Compute hashes using MD5, SHA1, SHA256, SHA512, SHA3-256")
	fmt.Println("Step 2: Enter two input strings")
	fmt.Println("Step 3: Compare the hash outputs for each algorithm")
	dashHeader()

	fmt.Print("Enter Input 1: ")
	scanner.Scan()
	input1 := scanner.Text()

	fmt.Print("Enter Input 2: ")
	scanner.Scan()
	input2 := scanner.Text()

	md5A := md5Hash(input1)
	md5B := md5Hash(input2)
	dashHeader()
	fmt.Println("MD5 Hash")
	fmt.Printf("Output A: %x\nOutput B: %x\n", md5A, md5B)
	fmt.Printf("=> %s\n", compareHash(md5A[:], md5B[:]))

	sha1A := sha1Hash(input1)
	sha1B := sha1Hash(input2)
	dashHeader()
	fmt.Println("SHA1 Hash")
	fmt.Printf("Output A: %x\nOutput B: %x\n", sha1A, sha1B)
	fmt.Printf("=> %s\n", compareHash(sha1A[:], sha1B[:]))

	sha256A := sha256Hash(input1)
	sha256B := sha256Hash(input2)
	dashHeader()
	fmt.Println("SHA256 Hash")
	fmt.Printf("Output A: %x\nOutput B: %x\n", sha256A, sha256B)
	fmt.Printf("=> %s\n", compareHash(sha256A[:], sha256B[:]))

	sha512A := sha512Hash(input1)
	sha512B := sha512Hash(input2)
	dashHeader()
	fmt.Println("SHA512 Hash")
	fmt.Printf("Output A: %x\nOutput B: %x\n", sha512A, sha512B)
	fmt.Printf("=> %s\n", compareHash(sha512A[:], sha512B[:]))

	sha3A := sha3Hash(input1)
	sha3B := sha3Hash(input2)
	dashHeader()
	fmt.Println("SHA3-256 Hash")
	fmt.Printf("Output A: %x\nOutput B: %x\n", sha3A, sha3B)
	fmt.Printf("=> %s\n", compareHash(sha3A[:], sha3B[:]))
	dashHeader()
}
