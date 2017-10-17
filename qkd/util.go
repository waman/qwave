package qkd

import (
	"math/rand"
)

//func NewRandomBitSimply(n int) []bool {
//	bs := make([]bool, n)
//	for i := 0; i < n; i++ {
//		bs[i] = rand.Intn(2) == 1  // [0, 2) => 0 or 1
//	}
//	return bs
//}

const pcBitMinus1 = (32 << (^uint(0) >> 63)) - 1 // == 31 or 63

func NewRandomBit(n int) []bool {
	bs := make([]bool, n)
	index := 0
	for {
		var max int
		var isLast bool
		if m := n-index; m <= pcBitMinus1 {
			max = m
			isLast = true
		}else{
			max = pcBitMinus1
			isLast = false
		}

		r := rand.Int()
		for i := 0; i < max; i++ {
			bs[index] = r & 1 == 1
			r >>= 1
			index++
		}

		if isLast { break }
	}
	return bs
}
