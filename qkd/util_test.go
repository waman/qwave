package qkd

import (
	"testing"
	"fmt"
	"image/color"
	"math"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
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