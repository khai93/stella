package random

import (
	"fmt"
	"math/rand"
	"time"
)

func NewToken(len int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:len]
}
