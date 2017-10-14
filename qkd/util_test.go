package qkd

import (
	"testing"
	"fmt"
)

func BenchmarkNewRandomBitSimply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewRandomBitSimply(63)
	}
}

func ExampleNewRandomBit() {
	bs := NewRandomBit(94)
	fmt.Println(Key(bs).String())
}

func BenchmarkNewRandomBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewRandomBit(63)
	}
}
