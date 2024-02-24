package user

import (
	"math/rand"
	"strconv"
	"time"
)

func toShortName(name string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return name + "@" + strconv.Itoa(r.Intn(1000000)+100000)
}
