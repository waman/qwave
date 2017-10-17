package bb84

import (
	"github.com/waman/qwave/qkd"
	"github.com/waman/qwave/system/qubit/basis"
)

type Eve struct {
	n int
	key qkd.Key
}

func NewEve(n int) *Eve {
  return &Eve{n, nil}
}

func (eve *Eve) Key() qkd.Key {
	return eve.key
}

func (eve *Eve) Eavsedrop(chWithAlice qkd.ChannelOnBob, chWithBob qkd.ChannelOnAlice) {
	nKey := 0
	for nKey < eve.n {
		qubits := <- chWithAlice.Qch()
		for _, qubit := range qubits {
			qubit.Observe(basis.Standard())
		}
		chWithBob.Qch() <- qubits

		chWithAlice.ToAlice() <- <- chWithBob.FromBob()  // nil
		chWithBob.ToBob() <- <- chWithAlice.FromAlice()  // bases

		matches := <- chWithBob.FromBob()
		for _, match := range matches {
			if match { nKey++ }
		}
		chWithAlice.ToAlice() <- matches
	}
}