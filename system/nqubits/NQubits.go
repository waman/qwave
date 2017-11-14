package nqubits

import (
	"github.com/waman/qwave/system/nqubits/nket"
	"github.com/waman/qwave/system/nqubits/nbasis"
	"github.com/waman/qwave/system/qubit/ket"
	"github.com/waman/qwave/system/qubit/basis"
	"github.com/waman/qwave/system/qubit"
)

type NQubits interface {
	Observe(b nbasis.NBasis) nket.NState
	ObserveInStandardBasis() nket.NState

	ObserveQubit(i int, b basis.Basis) ket.State
	ObserveQubitInStandardBasis() ket.State
	ObserveQubitInHadamardBasis() ket.State

	Qubit(i int) qubit.Qubit
}
