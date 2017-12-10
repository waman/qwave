package nqubits

import (
	"github.com/waman/qwave/system"
	"github.com/waman/qwave/system/nqubits/nbasis"
	"github.com/waman/qwave/system/nqubits/nket"
	"github.com/waman/qwave/system/nqubits/nop"
	"github.com/waman/qwave/system/qubit"
	"github.com/waman/qwave/system/qubit/basis"
	"github.com/waman/qwave/system/qubit/ket"
	"github.com/waman/qwave/system/qubit/op"
	"sync"
)

type MultiQubitSystem interface {
	QubitCount() int
	Dim() int // = 2^QubitCount()

	Observe(b nbasis.NBasis) nket.NState
	ObserveInStandardBasis() nket.NState

	ObserveQubit(i int, b *basis.Basis) *ket.State
	ObserveQubitInStandardBasis(i int) *ket.State
	ObserveQubitInHadamardBasis(i int) *ket.State

	Qubit(i int) qubit.Qubit

	Apply(u nop.Matrix)
	ApplyToQubit(i int, u op.Matrix2x2)
}

func New(qbtCount int, isNormalized bool, cs ...complex128) MultiQubitSystem {
	if !isNormalized { system.Normalize(cs)	}

	var _cs []complex128
	if dim := system.Dim(qbtCount); len(cs) != dim {
		_cs = make([]complex128, dim)
		copy(_cs, cs)
	} else {
		_cs = cs // DO NOT create copy (differ to NewDense function)
	}

	var mu sync.Mutex
	return &denseMultiQubitSystem{qbtCount, &mu, _cs}
}

type qubitImpl struct {
	parent MultiQubitSystem
	i      int
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

func (qbts *qubitImpl) Apply(u op.Matrix2x2) {
	qbts.parent.ApplyToQubit(qbts.i, u)
}

func U(m nop.Matrix, ms ...nop.Matrix) func(nqbts MultiQubitSystem) MultiQubitSystem {
	return func(nqbts MultiQubitSystem) MultiQubitSystem {
		for i := len(ms) - 1; i >= 0; i-- {
			nqbts.Apply(ms[i])
		}
		nqbts.Apply(m)
		return nqbts
	}
}
