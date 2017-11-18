package nqubits

import (
	"github.com/waman/qwave/system/nqubits/nket"
	"github.com/waman/qwave/system/nqubits/nbasis"
	"github.com/waman/qwave/system/qubit/ket"
	"github.com/waman/qwave/system/qubit/basis"
	"github.com/waman/qwave/system/qubit"
	"github.com/waman/qwave/system"
	"sync"
)

type NQubits interface {
	QubitCount() int
	Dim() int  // = 2^QubitCount()

	Observe(b nbasis.NBasis) nket.NState
  ObserveInStandardBasis() nket.NState

	ObserveQubit(i int, b *basis.Basis) *ket.State
  ObserveQubitInStandardBasis(i int) *ket.State
  ObserveQubitInHadamardBasis(i int) *ket.State

  Qubit(i int) qubit.Qubit

  ApplyToQubit(i int, u qubit.Matrix2)
}

func New(qbtCount int, isNormalized bool, cs ...complex128) NQubits {
	if !isNormalized { system.Normalize(cs) }

	var cs_ []complex128
	if dim := system.Dim(qbtCount); len(cs) != dim {
		cs_ = make([]complex128, dim)
		copy(cs_, cs)
	} else {
		cs_ = cs  // DO NOT create copy (differ to NewDense function)
	}

	var mu sync.Mutex
	return &denseNQubits{qbtCount, mu, cs_}
}

type qubitImpl struct {
	parent NQubits
	i int
}

func (qbts *qubitImpl) Observe(b *basis.Basis) *ket.State {
	return qbts.parent.ObserveQubit(qbts.i, b)
}

func (qbts *qubitImpl) ObserveInStandardBasis() *ket.State {
	return qbts.parent.ObserveQubitInStandardBasis(qbts.i)
}

func (qbts *qubitImpl) ObserveInHadamardBasis() *ket.State {
	return qbts.parent.ObserveQubitInHadamardBasis(qbts.i)
}

func (qbts *qubitImpl) Apply(u qubit.Matrix2) {
	qbts.parent.ApplyToQubit(qbts.i, u)
}