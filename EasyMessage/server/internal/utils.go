package data

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(hashedPassword), err
}

func HashPasswordCheck(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err == nil {
		return true
	}
	return false
}

func GenerateShortName(name string) string {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	name = "@" + strings.ToLower(name) + strconv.Itoa(rand.Intn(99999)+12345)
	return name
}
