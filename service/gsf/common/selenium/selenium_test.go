package selenium

import (
	"testing"
	"time"
)

func TestSelenium(t *testing.T) {
	gs := NewGsfSelenium()
	gs.Run()
	gs.Token()
	time.Sleep(5 * time.Minute)
}
