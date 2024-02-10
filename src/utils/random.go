package utils

import (
	"math/rand"
	"time"
)

func Create6digitsRandom() int {
	return newRandomNumber(100000, 999999)
}

func newRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func StringWithCharset(length int, charset string) string {
	var seededRand = rand.New(
		rand.NewSource(time.Now().UnixNano()),
	)
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func CreateRandomString(length int) string {
	return StringWithCharset(length, charset)
}
