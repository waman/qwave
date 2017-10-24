package qkd

import (
	"testing"
	"fmt"
)

func ExampleNewRandomBit() {
	bs := NewRandomBits(94)
	fmt.Println(Key(bs).String())
}

func BenchmarkNewRandomBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewRandomBits(63)
	}
}
