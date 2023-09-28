package utils

import (
	"math/rand"
	"strings"
	"time"
)

func RandomUrlString() string {
	// return a true random string of 6 characters (a-z, A-Z, 0-9)
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 3)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return strings.ToLower(string(b))
}
