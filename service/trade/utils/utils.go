package utils

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"strconv"
)

func TID() string {
	return uuid.NewMD5(uuid.NameSpaceDNS, []byte("-TID@#"+strconv.Itoa(rand.Int())+uuid.New().String())).String()
}

func AID() string {
	return uuid.NewMD5(uuid.NameSpaceDNS, []byte(strconv.Itoa(rand.Int())+"s-fAID-Te.fas"+uuid.New().String())).String()
}

func HID() string {
	return fmt.Sprintf("H%v", uuid.New().ID())
}
