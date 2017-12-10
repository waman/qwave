package ket2

import (
	. "github.com/waman/qwave/system"
	"github.com/waman/qwave/system/nqubits/nket"
)

var (
	zero  = nket.NewBase(2, 0) // = |00>
	one   = nket.NewBase(2, 1) // = |01>
	two   = nket.NewBase(2, 2) // = |10>
	three = nket.NewBase(2, 3) // = |11>

	phiPlus  = nket.New(2, true, S2I(), 0, 0,  S2I())  // = (|00> + |11>)/√2
	phiMinus = nket.New(2, true, S2I(), 0, 0, -S2I())  // = (|00> - |11>)/√2
	psiPlus  = nket.New(2, true, 0, S2I(),  S2I(), 0)  // = (|01> + |10>)/√2
	psiMinus = nket.New(2, true, 0, S2I(), -S2I(), 0)  // = (|01> - |10>)/√2
)

func New(a, b, c, d complex128, isNormalized bool) nket.NState {
	return nket.New(2, isNormalized, a, b, c, d)
}

func Zero()  nket.NState { return zero }
func One()   nket.NState { return one }
func Two()   nket.NState { return two }
func Three() nket.NState { return three }

func PhiPlus()  nket.NState { return phiPlus }
func PhiMinus() nket.NState { return phiMinus }
func PsiPlus()  nket.NState { return psiPlus }
func PsiMinus() nket.NState { return psiMinus }
