package utilities

import "crypto/rand"

func GenerateRandomBytes() int32 {
	b := make([]byte, 1, 1)
	_, _ = rand.Read(b)
	return int32(b[0])
}

