package ket

import (
	"math/cmplx"
	"math"
	"log"
	"math/rand"
	. "github.com/waman/qwave/system"
	"github.com/waman/qwave/system/qubit/op"
)

var (
	zero   = &State{1, 0}
	one    = &State{0, 1}
	plus   = &State{S2I(), S2I()}
	minus  = &State{S2I(), -S2I()}
	plusI  = &State{S2I(), S2I()*1i}
	minusI = &State{S2I(), -S2I()*1i}
)

func Zero()   *State { return zero }
func One()    *State { return one }
func Plus()   *State { return plus }
func Minus()  *State { return minus }
func PlusI()  *State { return plusI }
func MinusI() *State { return minusI }

// State represent an immutable ket vector a|0> + b|1>
type State struct {
	a, b complex128
}

func (x *State) Coefficients() (complex128, complex128) {
	return x.a, x.b
}

func (x *State) A() complex128 { return x.a }
func (x *State) B() complex128 { return x.b }

func (x *State) At(i int) complex128 {
	switch i {
	case 0: return x.a
	case 1: return x.b
	default:
		log.Panicf("The index must be in [%d, %d)", 0, 2)
	  return cmplx.NaN()
	}
}

// New function return the state a|0> + b|1> (when isNormalized is true).
// If isNormalized is false, the returned state is (a|0> + b|1>)/√(a^2+b^2).
func New(a, b complex128, isNormalized bool) *State {
	if b == 0 {
		if a == 0 {
			log.Panic("Both coefficients are zero.")
		}
		return zero

	} else if a == 0 {
		return one
	}

	if isNormalized {
		return &State{a, b}

	} else {
		aAbs, bAbs := cmplx.Abs(a), cmplx.Abs(b)
		norm := complex(math.Sqrt(aAbs*aAbs + bAbs*bAbs), 0)  // = √(|a|^2 + |b|^2)
		return &State{a/norm, b/norm}
	}
}

// <x|y>
func (x *State) InnerProduct(y *State) complex128 {
	return cmplx.Conj(x.a) * y.a + cmplx.Conj(x.b) * y.b
}

// |<x|y>|
func (x *State) Amplitude(y *State) float64 {
	return cmplx.Abs(x.InnerProduct(y))
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

// |x><y|
func (x *State) OuterProduct(y *State) op.Matrix2x2 {
	return op.NewMatrix2x2(
		cmplx.Conj(x.a)*y.a, cmplx.Conj(x.a)*y.b,
		cmplx.Conj(x.b)*y.a, cmplx.Conj(x.b)*y.b)
}

func (s *State) String() string {
	if s.b == 0 {
		return "|0>"

	} else if s.a == 0 {
		return "|1>"

	} else {
		// Both a and b are non-zero
		if s.b == s.a {
			return "|+>"
		} else if s.b == -s.a {
			return "|->"
		} else if s.b == s.a*1i {
			return "|+i>"
		} else if s.b == -s.a*1i {
			return "|-i>"
		}
	}

	return ToString(s.Coefficients())
}

func (x *State) Vertical() *State {
	return New(cmplx.Conj(x.b), -cmplx.Conj(x.a), true)
}

// Polar method returns the polar coordinates on the Bloch sphere.
// a|0> + b|1> -> (2*Acos(a), Phase(b))   (Phase(b) = Atan(b_i/b_r))
func (s *State) Polar() (theta, phi float64) {
	if s.b == 0 {
		return 0, 0
	} else if s.a == 0 {
		return math.Pi, 0
	}

	phase := s.a/complex(cmplx.Abs(s.a), 0)  // a/|a|
	a, b := s.a/phase, s.b/phase

	theta = math.Acos(real(a))*2
	phi = cmplx.Phase(b)
	if phi < 0 {
		phi += math.Pi*2
	}
	return
}

// (theta, phi) -> cos(theta/2)|0> + exp[i*phi]*sin(theta/2)|1>
func ByPolar(theta, phi float64) *State {
	if theta == 0 {
		return Zero()
	} else if theta == math.Pi {
		return One()
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
	return &State{complex(c, 0), cmplx.Exp(complex(0, phi))*complex(s, 0)}
}

// a|0> + b|1> -> b/a
// |0> -> 0
// |1> -> cmplx.Inf()
func (s *State) Complex() complex128 {
  return s.b/s.a
}

// c -> 1/√(1+|c|^2)|0> + (c/√(1+|c|^2))|1>
func ByComplex(c complex128) *State {
  if cmplx.IsInf(c) {
  	return One()
	} else if c == 0 {
		return Zero()
	}

	cAbs := cmplx.Abs(c)
	f := 1/math.Sqrt(1+cAbs*cAbs)  // = 1/√(1+|c|^2)
	return &State{complex(f, 0), c*complex(f, 0)}
}

func NewRandomState() *State {
	cosTheta := rand.Float64()*2-1
	phi := rand.Float64()*2*math.Pi
	return ByPolar(math.Acos(cosTheta), phi)
}

func NewRandomRealState() *State {
	theta := rand.Float64()*math.Pi
	return ByPolar(theta, 0)
}