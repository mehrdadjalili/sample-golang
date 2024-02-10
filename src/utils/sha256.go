package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

func NewSHA256(data []byte) string {
	hash := sha1.New()
	hash.Write(data)
	return hex.EncodeToString(hash.Sum(nil))
}
