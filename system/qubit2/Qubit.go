package qubit2

import (
	"github.com/waman/qwave/system/qubit2/ket2"
	"github.com/waman/qwave/system/nqubits"
	"github.com/waman/qwave/system/nqubits/nket"
	"github.com/waman/qwave/system/qubit"
)

func New(a, b, c, d complex128, isNormalized bool) nqubits.NQubits {
	return nqubits.New(2, isNormalized, a, b, c, d)
}

func NewWith(s nket.NState) nqubits.NQubits {
	return nqubits.New(2, true, s.Coefficients()...)
}

func NewZero() nqubits.NQubits {
	return NewWith(ket2.Zero())
}

func NewOne() nqubits.NQubits {
	return NewWith(ket2.One())
}

func NewTwo() nqubits.NQubits {
	return NewWith(ket2.Two())
}

func NewThree() nqubits.NQubits {
	return NewWith(ket2.Three())
}

func NewPhiPlus() nqubits.NQubits {
	return NewWith(ket2.PhiPlus())
}

func NewPhiMinus() nqubits.NQubits {
	return NewWith(ket2.PhiMinus())
}

func NewPsiPlus() nqubits.NQubits {
	return NewWith(ket2.PsiPlus())
}

func NewPsiMinus() nqubits.NQubits {
	return NewWith(ket2.PsiMinus())
}

func SplitQubits(q2s []nqubits.NQubits) (firsts, seconds []qubit.Qubit) {
	n := len(q2s)
	firsts = make([]qubit.Qubit, n)
	seconds = make([]qubit.Qubit, n)
	for i, q2 := range q2s {
		firsts[i] = q2.Qubit(0)
		seconds[i] = q2.Qubit(1)
	}
	return
}