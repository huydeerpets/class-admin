package utils

import (
	"time"
	"math/rand"
)
const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	numBytes    = "0123456789"
)
// RandStringNum get a rand number string
func RandStringNum(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, n)
	for i := range b {
		b[i] = numBytes[r.Intn(len(numBytes))]
	}
	return string(b)
}