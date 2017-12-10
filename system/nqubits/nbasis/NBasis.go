package nbasis

import (
	"github.com/waman/qwave/system/nqubits/nket"
	"github.com/waman/qwave/system"
	"log"
	"bytes"
)

type NBasis interface {
	QubitCount() int
	Dim() int
	States() []nket.NState
	At(i int) nket.NState

	String() string
}

func toString(b NBasis) string {
	var buf bytes.Buffer
	buf.WriteString("{")
	sep := ""
	for _, s := range b.States() {
		buf.WriteString(sep)
		buf.WriteString(s.String())
		sep = ", "
	}
	buf.WriteString("}")
	return buf.String()
}

//***** Standard basis *****
func Standard(qbtCount int) NBasis {
	return &nStandard{qbtCount}
}

type nStandard struct {
	qubitCount int
}

func (b *nStandard) QubitCount() int {
	return b.qubitCount
}

func (b *nStandard) Dim() int {
	return system.Dim(b.qubitCount)
}

func (b *nStandard) States() []nket.NState {
	n := b.Dim()
	ss := make([]nket.NState, n)
	for i := 0; i < n; i++ {
		ss[i] = nket.NewBase(n, i)
	}
	return ss
}

func (b *nStandard) At(i int) nket.NState {
	return nket.NewBase(b.qubitCount, i)
}

func (b *nStandard) String() string {
	return toString(b)
}

//***** default implementation of basis *****
// This method does not check orthogonality between nket.NStates.
func New(qbtCount int, states ...nket.NState) NBasis {
	if n := system.Dim(qbtCount); len(states) != n {
		log.Panicf("%d NState objects must be passed as arguments: only %d passed", n, len(states))
	}

	for i, state := range states {
		if state.QubitCount() != qbtCount {
			log.Panicf("%d-th NState has %d qubits, but this must be %d",
				i, state.QubitCount(), qbtCount)
		}
	}

	return &defaultNBasis{qbtCount, states}
}

type defaultNBasis struct {
	qubitCount int
	states []nket.NState
}

func (b *defaultNBasis) QubitCount() int {
	return b.qubitCount
}

func (b *defaultNBasis) Dim() int {
	return system.Dim(b.qubitCount)
}

func (b *defaultNBasis) States() []nket.NState {
	states := make([]nket.NState, b.Dim())
	copy(states, b.states)
	return states
}

func (b *defaultNBasis) At(i int) nket.NState {
	return b.states[i]
}

func (b *defaultNBasis) String() string {
	return toString(b)
}