package x

import (
	"crypto/md5"
	"encoding/hex"
)

func IsBlank(s string) bool {
	for i := range s {
		if s[i] != ' ' && s[i] != '\t' && s[i] != '\n' && s[i] != '\r' {
			return false
		}
	}
	return true
}

// Md5 returns the MD5 string which is 32 bytes long only containing hexadecimal characters.
func Md5(data string) string {
	hasher := md5.New()
	hasher.Write([]byte(data))
	sum := hasher.Sum(nil)
	return hex.EncodeToString(sum)
}
