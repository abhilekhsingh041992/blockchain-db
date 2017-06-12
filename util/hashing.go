package util

import (
	"crypto/sha256"
	"fmt"
	"encoding/hex"
)

func Hash(data string, timestamp int64) string {
	bytes := []byte(data + "-" + fmt.Sprint(timestamp))
	slice := sha256.Sum256(bytes)
	return hex.EncodeToString(slice[:])
}
