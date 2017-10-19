package ket

import (
	. "math/cmplx"
	"fmt"
	"math"
	"log"
)

// 1/√2 (sqrt 2 inverse)
var s2i float64 = 1/math.Sqrt(2)

var (
	Zero   *State = newState(1, 0)
	One    *State = newState(0, 1)
	Plus   *State = newState(s2i, complex(s2i, 0))
	Minus  *State = newState(s2i, -complex(s2i, 0))
	PlusI  *State = newState(s2i, complex(s2i, 0)*1i)
	MinusI *State = newState(s2i, -complex(s2i, 0)*1i)
)

// State represent an immutable ket vector a|0> + b|1>
// a is a non-negative number
// If a is zero, b is a positive number
type State struct {
	a float64
	b complex128
}

func (x *State) A() float64   {	return x.a }
func (x *State) B() complex128 { return x.b }

// a and b must satisfy that a >= 0 and |a|^2 + |b|^2 == 1.
// If a == 0, b > 0.
func newState(a float64, b complex128) *State{
	return &State{a, b}
}

func New(a, b complex128) *State {
	if b == 0 {
		if a == 0 {
			log.Panicf("Both coefficients are zero.")
		}
		return newState(1, 0)

	} else if a == 0 {
		return newState(0, 1)
	}

	aAbs, bAbs := Abs(a), Abs(b)
	norm := math.Sqrt(aAbs*aAbs + bAbs*bAbs)  // = √(|a|^2 + |b|^2)
	newA := aAbs/norm  // = |a|/norm
	newB := complex(newA, 0)*b/a  // | b|a|/(a*norm)
	return newState(newA, newB)
}

// <x|y>
func (x *State) Prod(y *State) complex128 {
	return complex(x.a* y.a, 0) + Conj(x.b) * y.b
}

// |<x|y>|
func (x *State) Amplitude(y *State) float64 {
	return Abs(x.Prod(y))
}

// |<x|y>|^2
func (x *State) Probability(y *State) float64 {
	amp := x.Amplitude(y)
	return amp*amp
}

// |<x|y>| == 1 (|<x|y>| >= 1-delta)
func (x *State) EqualState(y *State, delta float64) bool {
	if x == y { return true }
	return x.Amplitude(y) >= 1-delta
}

// |<x|y>| == 0 (|<x|y>| <= delta)
func (x *State) IsOrthogonalTo(y *State, delta float64) bool {
	return x.Amplitude(y) <= delta
}

func (s *State) String() string {
	if s.b == 0 {
		return "|0>"

	} else if s.a == 0 {
		return "|1>"

	} else {
		// Both a and b are non-zero
		re_b, im_b := real(s.b), imag(s.b)
		if im_b == 0 {
			if s.a == re_b {
				return "|+>"

			} else if s.a == -re_b {
				return "|->"

			} else if re_b > 0 {
				return fmt.Sprintf("%v|0> + %v|1>", s.a, re_b)

			}else{
				return fmt.Sprintf("%v|0> - %v|1>", s.a, -re_b)

			}

		} else if re_b == 0 {
			if s.a == im_b {
				return "|+i>"

			} else if s.a == -im_b {
				return "|-i>"

			} else if im_b := imag(s.b); im_b > 0 {
				return fmt.Sprintf("%v|0> + %vi|1>", s.a, im_b)

			}else{
				return fmt.Sprintf("%v|0> - %vi|1>", s.a, -im_b)

			}
		}

		return fmt.Sprintf("%v|0> + %v|1>", s.a, s.b)
	}
}