package qubits2

import (
	"github.com/waman/qwave/system/qubits2/ket2"
	"github.com/waman/qwave/system/nqubits"
	"github.com/waman/qwave/system/nqubits/nket"
	"github.com/waman/qwave/system/qubit"
	"log"
	"github.com/waman/qwave/system/qubits2/basis2"
)

type TwoQubitSystem interface {
	nqubits.MultiQubitSystem
	ObserveInBellBasis() nket.NState
	Split() (qbt0, qbt1 qubit.Qubit)
}

func New(a, b, c, d complex128, isNormalized bool) TwoQubitSystem {
	return &defaultTwoQubitSystem{nqubits.New(2, isNormalized, a, b, c, d)}
}

func NewWith(s nket.NState) TwoQubitSystem {
	if s.QubitCount() != 2 {
		log.Panicf("The argument nket.NState object must have QubitCount 2: %d appears.",
			s.QubitCount())
	}
	return &defaultTwoQubitSystem{nqubits.New(2, true, s.Coefficients()...)}
}

type defaultTwoQubitSystem struct {
	nqubits.MultiQubitSystem
}

func (qbts *defaultTwoQubitSystem) ObserveInBellBasis() nket.NState {
	return qbts.Observe(basis2.Bell())
}

func (qbts *defaultTwoQubitSystem) Split() (qbt0, qbt1 qubit.Qubit) {
	qbt0, qbt1 = qbts.Qubit(0), qbts.Qubit(1)
	return
}

func NewZero() TwoQubitSystem {
	return NewWith(ket2.Zero())
}

func NewOne() TwoQubitSystem {
	return NewWith(ket2.One())
}

func NewTwo() TwoQubitSystem {
	return NewWith(ket2.Two())
}

func NewThree() TwoQubitSystem {
	return NewWith(ket2.Three())
}

func NewPhiPlus() TwoQubitSystem {
	return NewWith(ket2.PhiPlus())
}

func NewPhiMinus() TwoQubitSystem {
	return NewWith(ket2.PhiMinus())
}

func NewPsiPlus() TwoQubitSystem {
	return NewWith(ket2.PsiPlus())
}

func NewPsiMinus() TwoQubitSystem {
	return NewWith(ket2.PsiMinus())
}