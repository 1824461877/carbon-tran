package utils

import (
	"github.com/google/uuid"
	"math/rand"
	"strconv"
)

func TID() string {
	return uuid.NewMD5(uuid.NameSpaceDNS, []byte("-TID@#"+strconv.Itoa(rand.Int())+uuid.New().String())).String()
}
