package ket

import (
	"testing"
	"fmt"
	"math"
	"math/cmplx"
)

const delta float64 = 1e-10

func TestEqualState(t *testing.T){
	var tests = []struct {
		first, second *State
		want  bool
	}{
		{Zero  , New(1, 0), true},
		{One   , New(0, 1), true},
		{Plus  , New(1, 1), true},
		{Minus , New(1, -1), true},
		{PlusI , New(1, 1i), true},
		{MinusI, New(1, -1i), true},
		{Zero, New(0, 1), false},
	}
	for _, test := range tests {
		if got := test.first.EqualState(test.second, delta); got != test.want {
			t.Errorf("(%v).Equal(%v) = %v", test.first, test.second, got)
		}
	}
}

func TestIsOrthogonalTo(t *testing.T){
	var tests = []struct {
		first, second *State
		want  bool
	}{
		{Zero, One, true},
		{Plus, Minus, true},
		{PlusI, MinusI, true},
		{Zero, Plus, false},
		{Zero, MinusI, false},
	}
	for _, test := range tests {
		if got := test.first.IsOrthogonalTo(test.second, delta); got != test.want {
			t.Errorf("(%v).IsOrthogonalTo(%v) = %v", test.first, test.second, got)
		}
	}
}

func TestStringMethod(t *testing.T){
	var tests = []struct {
		state *State
		want  string
	}{
		{Zero, "|0>"},
		{One, "|1>"},
		{Plus, "|+>"},
		{Minus, "|->"},
		{PlusI, "|+i>"},
		{MinusI, "|-i>"},
		{New(3, 4), "0.6|0> + 0.7999999999999999|1>"},
	}
	for i, test := range tests {
		if got := test.state.String(); got != test.want {
			t.Errorf("%d番目：(%q).String() != %s", i, got, test.want)
		}
	}
}

func EqF10(x, y float64) bool {
	return math.Abs(x - y) < 1e-10
}

func EqC10(x, y complex128) bool {
	return cmplx.Abs(x - y) < 1e-10
}

func TestPolarMethod(t *testing.T){
	var tests = []struct {
		state *State
		wantTheta  float64
		wantPhi float64
	}{
		{Zero, 0, 0},
		{One, math.Pi, 0},
		{Plus, math.Pi/2, 0},
		{Minus, math.Pi/2, math.Pi},
		{PlusI, math.Pi/2, math.Pi/2},
		{MinusI, math.Pi/2, math.Pi*3/2},
	}
	for i, test := range tests {
		if gotTheta, gotPhi := test.state.Polar();
			  !(EqF10(gotTheta, test.wantTheta) && EqF10(gotPhi, test.wantPhi) ){
			t.Errorf("%d番目：(%f, %f) != (%f, %f)",
				i, gotTheta, gotPhi, test.wantTheta, test.wantPhi)
		}
	}
}

func TestByPolarFunction(t *testing.T){
	var tests = []struct {
		theta, phi float64
		want *State
	}{
		{0, 0, Zero},
		{math.Pi, 0, One},
		{math.Pi/2, 0, Plus},
		{math.Pi/2, math.Pi, Minus},
		{math.Pi/2, math.Pi/2, PlusI},
		{math.Pi/2, math.Pi*3/2, MinusI},
		{math.Pi/3, math.Pi/4, newState(math.Sqrt(3)/2, cmplx.Exp(math.Pi*1i/4)/2)},
	}
	for i, test := range tests {
		if got := ByPolar(test.theta, test.phi); !got.EqualState(test.want, delta){
			t.Errorf("%d番目： ByPolar(%f, %f) = %q != %q", i, test.theta, test.phi, got, test.want)
		}
	}
}

func TestComplexMethod(t *testing.T){
	if !cmplx.IsInf(Zero.Complex()) {
		t.Errorf("(|0>).Complex() is not cmplx.Inf()")
	}

	var tests = []struct {
		state *State
		want complex128
	}{
		{One, 0},
		{Plus, complex(1, 0)},
		{Minus, complex(-1, 0)},
		{PlusI, 1i},
		{MinusI, -1i},
		{ByPolar(math.Pi/2, math.Pi/4), complex(s2i, s2i)},
		{ByPolar(math.Pi/3, math.Pi*3/4), complex(-math.Sqrt(6)/2, math.Sqrt(6)/2)},
	}
	for i, test := range tests {
		if got := test.state.Complex(); !EqC10(got, test.want) {
			t.Errorf("%d番目：(%q).Complex() = %g != %g",	i, test.state, got, test.want)
		}
	}
}

func TestByComplexFunction(t *testing.T){
	var tests = []struct {
		c complex128
		want *State
	}{
		{cmplx.Inf(), Zero},
		{0, One},
		{complex(1, 0), Plus},
		{complex(-1, 0), Minus},
		{1i, PlusI},
		{-1i, MinusI},
		{complex(s2i, s2i), ByPolar(math.Pi/2, math.Pi/4)},
		{complex(-math.Sqrt(6)/2, math.Sqrt(6)/2), ByPolar(math.Pi/3, math.Pi*3/4)},
	}
	for i, test := range tests {
		if got := ByComplex(test.c); !got.EqualState(test.want, 1e-10) {
			t.Errorf("%d番目：ByComplex(%g) = %s != %s",	i, test.c, got, test.want)
		}
	}
}

func Example(){
	zero  := Zero
	one   := One
	plus  := Plus
	minus := Minus

	s := New(1, 1)

	delta := 1e-10
	fmt.Println(s.EqualState(plus, delta))
	fmt.Println(zero.IsOrthogonalTo(one, 1e-10))

	p := zero.Probability(minus)
	fmt.Println(math.Abs(p - 0.5) <= delta)
	// Output:
	// true
	// true
	// true
}