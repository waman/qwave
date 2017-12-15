package ket

import (
	"testing"
	"fmt"
	"math"
	"math/cmplx"
	"github.com/waman/qwave/system"
	"github.com/ToQoz/gopwt/assert"
)

const delta = 1e-10
var s2i = system.S2I()

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
		assert.OK(t, test.first.EqualState(test.second, delta) == test.want,
		 "Two states must be equivalent")
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
		assert.OK(t, test.first.IsOrthogonalTo(test.second, delta) == test.want,
			"Two states must be orthgonal.")
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
	for _, test := range tests {
		assert.OK(t, test.state.String() == test.want)
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
	for _, test := range tests {
		gotTheta, gotPhi := test.state.Polar()
		assert.OK(t, EqF10(gotTheta, test.wantTheta) && EqF10(gotPhi, test.wantPhi))
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
	for _, test := range tests {
		assert.OK(t, ByPolar(test.theta, test.phi).EqualState(test.want, delta))
	}
}

func TestComplexMethod(t *testing.T){
	// |1> -> infinity
	assert.OK(t, cmplx.IsInf(one.Complex()), "(|1>).Complex() is not cmplx.Inf()")

	var tests = []struct {
		state *State
		want complex128
	}{
		{zero, 0},
		{plus, 1},
		{minus, -1},
		{plusI, 1i},
		{minusI, -1i},
		{ByPolar(math.Pi/2, math.Pi/4), s2i + s2i*1i},
		{New(3, 4, false), 4.0/3.0},
		{New(5, 12i, false), 2.4i},
	}
	for _, test := range tests {
		assert.OK(t, EqC10(test.state.Complex(), test.want))
	}
}

func TestByComplexFunction(t *testing.T){
	var tests = []struct {
		c complex128
		want *State
	}{
		{0, zero},
		{cmplx.Inf(), one},
		{1, plus},
		{-1, minus},
		{1i, plusI},
		{-1i, minusI},
		{s2i + s2i*1i, ByPolar(math.Pi/2, math.Pi/4)},
		{3.0/4.0, New(4, 3, false)},
		{5.0i/12.0, New(12, 5i, true)},
	}
	for _, test := range tests {
		assert.OK(t, ByComplex(test.c).EqualState(test.want, 1e-10))
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