package ket

import (
	"math/cmplx"
	"fmt"
	"math"
	"log"
)

// 1/√2 (sqrt 2 inverse)
var s2i float64 = 1/math.Sqrt(2)

var (
	Zero   = newState(1, 0)
	One    = newState(0, 1)
	Plus   = newState(s2i, complex(s2i, 0))
	Minus  = newState(s2i, -complex(s2i, 0))
	PlusI  = newState(s2i, complex(s2i, 0)*1i)
	MinusI = newState(s2i, -complex(s2i, 0)*1i)
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

	aAbs, bAbs := cmplx.Abs(a), cmplx.Abs(b)
	norm := math.Sqrt(aAbs*aAbs + bAbs*bAbs)  // = √(|a|^2 + |b|^2)
	newA := aAbs/norm  // = |a|/norm
	newB := complex(newA, 0)*b/a  // | b|a|/(a*norm)
	return newState(newA, newB)
}

// <x|y>
func (x *State) Prod(y *State) complex128 {
	return complex(x.a* y.a, 0) + cmplx.Conj(x.b) * y.b
}

// |<x|y>|
func (x *State) Amplitude(y *State) float64 {
	return cmplx.Abs(x.Prod(y))
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

// a|0> + b|1> -> (2*Acos(a), Phase(b))   (Phase(b) = Atan(b_i/b_r))
func (s *State) Polar() (theta, phi float64) {
	theta = math.Acos(s.a)*2
	phi = cmplx.Phase(s.b)
	if phi < 0 {
		phi += math.Pi*2
	}
	return
}

// (theta, phi) -> cos(theta/2)|0> + exp[i*phi]*sin(theta/2)|1>
func ByPolar(theta, phi float64) *State {
	if theta == 0 {
		return Zero
	} else if theta == math.Pi {
		return One
	}// else if theta == math.Pi/2 {
	//	if phi == 0 {
	//		return Plus
	//	} else if phi == math.Pi {
	//		return Minus
	//	} else if phi == math.Pi/2 {
	//		return PlusI
	//	} else if phi == math.Pi*3/2 {
	//		return MinusI
	//	}
	//}

	s, c := math.Sincos(theta/2)
	return newState(c, cmplx.Exp(complex(0, phi))*complex(s, 0))
}

// a|0> + b|1> -> ab/(1-a^2)
// |0> -> cmplx.Nan()
// |1> -> 0
func (s *State) Complex() complex128 {
	if s.a == 0 {
		return 0
	} else if s.a == 1 {
		return cmplx.Inf()
	}

  return complex(s.a/(1-s.a*s.a), 0)*s.b
}

// c -> |c|/√(|c|^2+1)|0> + c/(|c|√(|c|^2+1))|1>
func ByComplex(c complex128) *State {
  if cmplx.IsInf(c) {
  	return Zero
	} else if c == 0 {
		return One
	}

	cAbs := cmplx.Abs(c)
	f := 1/math.Sqrt(cAbs*cAbs+1)  // = 1/√(|c|^2+1)
	return newState(cAbs*f, c*complex(f/cAbs, 0))
}