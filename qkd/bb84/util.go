package bb84

import "math/rand"

func NewRandomBit(n int) []bool {
	bs := make([]bool, n)
	for i := 0; i < n; i++ {
		bs[i] = rand.Intn(1) == 1
	}
	return bs
}