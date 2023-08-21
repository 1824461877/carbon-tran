package retire_cert

import (
	"github.com/google/uuid"
	"math/rand"
	"strconv"
)

func RCID(name string, onTime string) string {
	return uuid.NewMD5(uuid.NameSpaceDNS, []byte(name+"RCID"+"TIME"+onTime+strconv.Itoa(rand.Int())+uuid.New().String())).String()
}
