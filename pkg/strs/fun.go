package strs

import (
	r "math/rand"
	"time"
)

var rand *r.Rand

func init() {
	rand = r.New(r.NewSource(time.Now().UnixNano()))
}

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func Random(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = rune(letterBytes[rand.Intn(len(letterBytes))])
	}
	return string(b)
}
