package crack
import (
	"crypto/sha512"
)

func sha512Hash(s string) [64]byte {
	return sha512.Sum512([]byte(s))
}

func SHA512Crack(value string, targetHash [64]byte) string {
	valueHash := sha512Hash(value)
	if valueHash == targetHash {
		return value
	}
	return ""
}
