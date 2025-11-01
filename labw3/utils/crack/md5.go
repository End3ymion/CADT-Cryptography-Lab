package crack

import (
	"crypto/md5"
)

func md5Hash(s string) [16]byte {
	return md5.Sum([]byte(s))
}

func MD5Crack(value string, targetHash [16]byte) string {
	valueHash := md5Hash(value)
	if valueHash == targetHash {
		return value
	}
	return ""
}
