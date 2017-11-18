package nbasis

import (
	"github.com/waman/qwave/system/nqubits/nket"
	"github.com/waman/qwave/system"
	"log"
)

type NBasis interface {
	Dim() int
	States() []nket.NState
	Get(i int) nket.NState
}

//***** Standard basis *****
func Standard(qbtCount int) NBasis {
	return &nStandard{system.Dim(qbtCount)}
}

type nStandard struct {
	n int // number of base vectors, dimension of state vector (2^qubitCount)
}

func (b *nStandard) Dim() int {
	return b.n
}

func (b *nStandard) States() []nket.NState {
	ss := make([]nket.NState, b.n)
	for i := 0; i < b.n; i++ {
		ss[i] = nket.NewBase(b.n, i)
	}
	return ss
}

func (b *nStandard) Get(i int) nket.NState {
	return nket.NewBase(b.n, i)
}

//***** default implementation of basis *****
func New(qbtCount int, states ...nket.NState) NBasis {
	if len(states) != system.Dim(qbtCount) {
		log.Panicf("%d NState objects must be passed as arguments: only %d passed",
			system.Dim(qbtCount), len(states))
	}
	return &defaultNBasis{system.Dim(qbtCount), states}
}

type defaultNBasis struct {
	n int // number of base vectors, dimension of state vector (2^qubitCount)
	states []nket.NState
}

func (b *defaultNBasis) Dim() int {
	return b.n
}

func (b *defaultNBasis) States() []nket.NState {
	states := make([]nket.NState, b.n)
	copy(states, b.states)
	return states
}

func (b *defaultNBasis) Get(i int) nket.NState {
	return b.states[i]
}
