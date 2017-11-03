package ket

import (
	"testing"
	"fmt"
	"math"
	"math/cmplx"
)

const delta = 1e-10

func TestEqualState(t *testing.T){
	var tests = []struct {
		first, second *State
		want  bool
	}{
		{zero  , &State{1, 0}, true},
		{one   , &State{0, 1}, true},
		{plus  , &State{1, 1}, true},
		{minus , &State{1, -1}, true},
		{plusI , &State{1, 1i}, true},
		{minusI, &State{1, -1i}, true},
		{zero, &State{0, 1}, false},
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
		{zero, one, true},
		{plus, minus, true},
		{plusI, minusI, true},
		{zero, plus, false},
		{zero, minusI, false},
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
		{zero, "|0>"},
		{one, "|1>"},
		{plus, "|+>"},
		{minus, "|->"},
		{plusI, "|+i>"},
		{minusI, "|-i>"},
		{New(3, 4, false), "0.6|0> + 0.8|1>"},
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
		{zero, 0, 0},
		{one, math.Pi, 0},
		{plus, math.Pi/2, 0},
		{minus, math.Pi/2, math.Pi},
		{plusI, math.Pi/2, math.Pi/2},
		{minusI, math.Pi/2, math.Pi*3/2},
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
		{0, 0, zero},
		{math.Pi, 0, one},
		{math.Pi/2, 0, plus},
		{math.Pi/2, math.Pi, minus},
		{math.Pi/2, math.Pi/2, plusI},
		{math.Pi/2, math.Pi*3/2, minusI},
		{math.Pi/3, math.Pi/4, &State{complex(math.Sqrt(3)/2, 0), cmplx.Exp(math.Pi*1i/4)/2}},
	}
	for i, test := range tests {
		if got := ByPolar(test.theta, test.phi); !got.EqualState(test.want, delta){
			t.Errorf("%d番目： ByPolar(%f, %f) = %q != %q", i, test.theta, test.phi, got, test.want)
		}
	}
}

func TestComplexMethod(t *testing.T){
	if !cmplx.IsInf(zero.Complex()) {
		t.Errorf("(|0>).Complex() is not cmplx.Inf()")
	}

	var tests = []struct {
		state *State
		want complex128
	}{
		{one, 0},
		{plus, complex(1, 0)},
		{minus, complex(-1, 0)},
		{plusI, 1i},
		{minusI, -1i},
		{ByPolar(math.Pi/2, math.Pi/4), s2i + s2i*1i},
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
		{cmplx.Inf(), zero},
		{0, one},
		{complex(1, 0), plus},
		{complex(-1, 0), minus},
		{1i, plusI},
		{-1i, minusI},
		{s2i + s2i*1i, ByPolar(math.Pi/2, math.Pi/4)},
		{complex(-math.Sqrt(6)/2, math.Sqrt(6)/2), ByPolar(math.Pi/3, math.Pi*3/4)},
	}
	for i, test := range tests {
		if got := ByComplex(test.c); !got.EqualState(test.want, 1e-10) {
			t.Errorf("%d番目：ByComplex(%g) = %s != %s",	i, test.c, got, test.want)
		}
	}
}

func Example(){
	zero  := zero
	one   := one
	plus  := plus
	minus := minus

	s := New(1, 1, false)

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