package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func GenerateShortName(name string) string {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	name = "@" + strings.ToLower(name) + strconv.Itoa(rand.Intn(99999)+12345)
	return name
}
