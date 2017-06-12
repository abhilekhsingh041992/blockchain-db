package util

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash(data string) string {
	slice := sha256.Sum256([]byte(data))
	return hex.EncodeToString(slice[:])
}
