package main
import (
	"fmt"
	"crypto/md5"
)
func md5Hash(s string) string {
	s = md5.New()
	s = s.Sum(nil)
	return string(s)
}
func main () {
	var s string
	fmt.Println("Enter string:")
	fmt.Scan(&s)
	md5Hash(s)
}
