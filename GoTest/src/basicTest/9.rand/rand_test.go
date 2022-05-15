package rand

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRand(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	intn := rand.Intn(12)
	fmt.Println(intn)
	perm := rand.Perm(12)
	fmt.Println(perm)
}
