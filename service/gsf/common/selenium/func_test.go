package selenium

import (
	"testing"
	"time"
)

func TestFuncToken(t *testing.T) {
	c := NewCache()
	GetToken(c)

	time.Sleep(1 * time.Minute)
	//fmt.Println(c.Token.IdToken)
}
