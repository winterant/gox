package xstring

import (
	"crypto/rand"
	"io"
)

const randomCharSet = "abcdefghijklmnopqrstuvwxyz0123456789"

func Random(prefix string, totalLength int) string {
	var bytes = make([]byte, totalLength)
	if _, err := io.ReadFull(rand.Reader, bytes); err != nil {
		return "__RandomStrError__"
	}
	for i, b := range bytes {
		bytes[i] = randomCharSet[b%byte(len(randomCharSet))]
	}
	return prefix + string(bytes)
}
