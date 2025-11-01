package crack
import (
	"crypto/sha1"
)

func sha1Hash(s string) [20]byte {
	return sha1.Sum([]byte(s))
}

func SHA1Crack(value string, targetHash [20]byte) string {
	valueHash := sha1Hash(value)
	if valueHash == targetHash {
		return value
	}
	return ""
}
