package utils

import (
	"github.com/google/uuid"
	"math/rand"
	"strconv"
)

func PID() string {
	return uuid.NewMD5(uuid.NameSpaceDNS, []byte("-PID@#"+strconv.Itoa(rand.Int())+uuid.New().String())).String()
}
