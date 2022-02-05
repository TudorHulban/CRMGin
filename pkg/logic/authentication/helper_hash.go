package authentication

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/TudorHulban/GinCRM/cmd/setup"
	"golang.org/x/crypto/bcrypt"
)

/*
File provides global helpers for salt and hashing.
*/

// GenerateSALT Generates salt based on settings value.
func GenerateSALT() string {
	return randomString(setup.LenSALT)
}

// GenerateSessionID Generates session ID based on settings value and UNIX time.
func generateSessionID() string {
	return UXSecs() + randomString(setup.LenSessionID)
}

// HASHPassword Uses bcrypt to generate password hash.
func HASHPassword(password, salt string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password+salt), setup.HASHCost)
}

// UXSecs Returns UNIX time in seconds.
func UXSecs() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

// UXNano Returns UNIX time in nanoseconds.
func UXNano() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

func checkPasswordHash(password, salt, hashedPassword string) bool {
	errCompare := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password+salt))
	return errCompare == nil
}

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	res := make([]byte, length)
	for i := range res {
		res[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(res)
}
