package public

import (
	"crypto/sha256"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"regexp"
	"strconv"
)

func WID() string {
	return uuid.NewMD5(uuid.NameSpaceDNS, []byte(strconv.Itoa(rand.Int())+"@WID"+uuid.New().String())).String()
}

func AID() string {
	return uuid.NewMD5(uuid.NameSpaceDNS, []byte(strconv.Itoa(rand.Int())+"s-fAID-Te.fas"+uuid.New().String())).String()
}

func HID() string {
	return fmt.Sprintf("H%v", uuid.New().ID())
}

func EID() string {
	return uuid.NewMD5(uuid.NameSpaceDNS, []byte("ex-EID-ex"+strconv.Itoa(rand.Int())+uuid.New().String())).String()
}

func RID() string {
	return uuid.NewMD5(uuid.NameSpaceDNS, []byte("rs-RID-rs"+strconv.Itoa(rand.Int())+uuid.New().String())).String()
}

func GenSaltPassword(salt, password string) string {
	return Salt(salt, password)
}

func CheckPassword(ps string) bool {
	if len(ps) < 9 {
		return false
	}
	num := `[0-9]{1}`
	a_z := `[a-z]{1}`
	A_Z := `[A-Z]{1}`
	symbol := `[!@#~$%^&*()+|_.]{1}`
	if b, err := regexp.MatchString(num, ps); !b || err != nil {
		return false
	}
	if b, err := regexp.MatchString(a_z, ps); !b || err != nil {
		return false
	}
	if b, err := regexp.MatchString(A_Z, ps); !b || err != nil {
		return false
	}
	if b, err := regexp.MatchString(symbol, ps); !b || err != nil {
		return false
	}
	return true
}

func Salt(salt, val string) string {
	s1 := sha256.New()
	s1.Write([]byte(val))
	str1 := fmt.Sprintf("%x", s1.Sum(nil))
	s2 := sha256.New()
	s2.Write([]byte(str1 + salt))
	return fmt.Sprintf("%x", s2.Sum(nil))
}
