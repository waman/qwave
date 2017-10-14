package qkd

import (
	"math/rand"
)

func NewRandomBit(n int) []bool {
	bs := make([]bool, n)
	for i := 0; i < n; i++ {
		bs[i] = rand.Intn(2) == 1  // [0, 2) => 0 or 1
	}
	return bs
}