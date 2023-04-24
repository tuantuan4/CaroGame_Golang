package common

import (
	"math/rand"
	"time"
)

const SIZE_STRING = 6

func RandomString() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	rand.Seed(time.Now().UnixNano())

	b := make([]rune, SIZE_STRING)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GenerateOTP() string {
	otp := rand.Intn(999999)
	return IntToString(otp)
}
