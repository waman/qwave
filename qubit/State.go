package qubit

import (
	. "math/cmplx"
	"fmt"
	"math"
	"log"
)

func toComplex(re float64) complex128 {
	return complex(re, 0)
}

// 1/√2
var sqrt2inv complex128 = toComplex(1/math.Sqrt(2))

var (
	Zero   *State = NewState(1, 0)
	One    *State = NewState(0, 1)
	Plus   *State = NewState(sqrt2inv, sqrt2inv)
	Minus  *State = NewState(sqrt2inv, -sqrt2inv)
	PlusI  *State = NewState(sqrt2inv, sqrt2inv*1i)
	MinusI *State = NewState(sqrt2inv, -sqrt2inv*1i)
)

// a|0> + b|1>
type State struct {
	a float64
	b complex128
}

// a and b must satisfy that a >= 0 and |a|^2 + |b|^2 == 1.
// If a == 0, b > 0.
func newState(a float64, b complex128) *State{
	return &State{a, b}
}

func NewState(a, b complex128) *State {
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
	newB := toComplex(newA)*b/a  // | b|a|/(a*norm)
	return newState(newA, newB)
}

func (x *State) Prod(y *State) complex128 {
  return toComplex(x.a * y.a) + Conj(x.b) * y.b
}

func (x *State) Amplitude(y *State) float64 {
	return Abs(x.Prod(y))
}

func (x *State) Probability(y *State) float64 {
	amp := x.Amplitude(y)
	return amp*amp
}

func (x *State) EqualState(y *State, delta float64) bool {
	return x.Probability(y) >= 1-delta
}

func (s *State) String() string {
	if s.b == 0 {
		return "|0>"

	} else if s.a == 0 {
		return "|1>"

	} else {
		// Both a and b are non-zero
		if imag(s.b) == 0 {
			if re_b := real(s.b); re_b > 0 {
				return fmt.Sprintf("%v|0> + %v|1>", s.a, re_b)
			}else{
				return fmt.Sprintf("%v|0> - %v|1>", s.a, -re_b)
			}

		} else if real(s.b) == 0 {
			if im_b := imag(s.b); im_b > 0 {
				return fmt.Sprintf("%v|0> + %vi|1>", s.a, im_b)
			}else{
				return fmt.Sprintf("%v|0> - %vi|1>", s.a, -im_b)
			}
		}

	  return fmt.Sprintf("%v|0> + %v|1>", s.a, s.b)
	}
}