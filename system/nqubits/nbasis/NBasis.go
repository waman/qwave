package nbasis

import (
	"github.com/waman/qwave/system/nqubits/nket"
	"github.com/waman/qwave/system/nqubits"
)

func Standard(qbtCount int) NBasis {
  return &nStandard{nqubits.Dim(qbtCount)}
}

type NBasis interface {
	Dim() int
	States() []*nket.NState
	Get(i int) *nket.NState
}

type nStandard struct {
	n int // number of base vectors, dimension of state vector (2^qubitCount)
}

func (b *nStandard) Dim() int {
	return b.n
}

func (b *nStandard) States() []*nket.NState {
	ss := make([]*nket.NState, b.n)
	for i := 0; i < b.n; i++ {
		ss[i] = nket.NewBase(b.n, i)
	}
	return ss
}

func (b *nStandard) Get(i int) *nket.NState {
	return nket.NewBase(b.n, i)
}