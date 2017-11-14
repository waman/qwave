package ket2

import (
	"math"
	"log"
	"math/cmplx"
	. "github.com/waman/qwave/system"
)

var (
	zero  = &State{1, 0, 0, 0} // = |00>
	one   = &State{0, 1, 0, 0} // = |01>
	two   = &State{0, 0, 1, 0} // = |10>
	three = &State{0, 0, 0, 1} // = |11>

	phiPlus  = &State{S2I(), 0, 0,  S2I()}  // = (|00> + |11>)/√2
	phiMinus = &State{S2I(), 0, 0, -S2I()}  // = (|00> - |11>)/√2
	psiPlus  = &State{0, S2I(),  S2I(), 0}  // = (|01> + |10>)/√2
	psiMinus = &State{0, S2I(), -S2I(), 0}  // = (|01> - |10>)/√2
)

func Zero()  *State { return zero }
func One()   *State { return one }
func Two()   *State { return two }
func Three() *State { return three }

func PhiPlus()  *State { return phiPlus }
func PhiMinus() *State { return phiMinus }
func PsiPlus()  *State { return psiPlus }
func PsiMinus() *State { return psiMinus }

// a|00> + b|01> + c|10> + d|11>
// a must be non-negative real number.
type State struct{
	a, b, c, d complex128
}

func (x *State) Coefficients() (complex128, complex128, complex128, complex128) {
	return x.a, x.b, x.c, x.d
}

func (s *State) Get(i int) complex128 {
	switch i {
	case 0: return s.a
	case 1: return s.b
	case 2: return s.c
	case 3: return s.d
	}
	log.Panicf("Index is out of bounds [0, 4): %d", i)
	return cmplx.NaN()
}

func New(a, b, c, d complex128, isNormalized bool) *State {
	p := nonZeroBitSet(a, b, c ,d)

	// states in the standard basis
	if p == 8 { return zero }
	if p == 4 { return one }
	if p == 2 { return two }
	if p == 1 { return three }

	// states in the Bell basis
	if p == 12 {
		if a == b {
			return phiPlus
		} else if a == -b {
			return phiMinus
		}
	} else if p == 3 {
		if c == d {
			return psiPlus
		} else if c == -d {
			return psiMinus
		}
	}

	if isNormalized {
		return &State{a, b, c, d}

	} else {
		aAbs, bAbs, cAbs, dAbs := cmplx.Abs(a), cmplx.Abs(b), cmplx.Abs(c), cmplx.Abs(d)
		norm := complex(math.Sqrt(aAbs*aAbs + bAbs*bAbs + cAbs*cAbs + dAbs*dAbs), 0)
		return &State{a/norm, b/norm, c/norm, d/norm}
	}
}

func nonZeroBitSet(a, b, c, d complex128) int {
	p := 0
	if a != 0 { p += 8 }
	if b != 0 { p += 4 }
	if c != 0 { p += 2 }
	if d != 0 { p += 1 }

	if p == 0 { log.Panic("All coefficients are zero.") }
	return p
}

// <x|y>
func (x *State) Prod(y *State) complex128 {
	return cmplx.Conj(x.a) * y.a +
		       cmplx.Conj(x.b) * y.b +
		         cmplx.Conj(x.c) * y.c +
		           cmplx.Conj(x.d) * y.d
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
	p := nonZeroBitSet(s.a, s.b, s.c, s.d)

	// states in the standard basis
	if p == 8 { return "|00>" }
	if p == 4 { return "|01>" }
	if p == 2 { return "|10>" }
	if p == 1 { return "|11>" }

	// states in the Bell basis
	if p == 12 {
		if s.a == s.b {
			return "|Φ+>"
		} else if s.a == -s.b {
			return "|Φ->"
		}
	} else if p == 3 {
		if s.c == s.d {
			return "|Ψ+>"
		} else if s.c == -s.d {
			return "|Ψ->"
		}
	}

	return ToString(s.Coefficients())
}