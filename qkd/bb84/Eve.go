package bb84

import (
	"github.com/waman/qwave/qkd"
	"github.com/waman/qwave/system/qubit/basis"
)

func NewEve(n int) *Eve {
	return &Eve{n, nil, make(chan struct{})}
}

type Eve struct {
	n int
	key qkd.Key
	done chan struct{}
}

func (eve *Eve) Key() qkd.Key {
	return eve.key
}

func (eve *Eve) Stop(){
	eve.done <- struct{}{}
}

func (eve *Eve) Eavesdrop(in *qkd.InternalOfChannel) {
	loop:
		for {
			select {
			case qbts := <-in.QubitsFromAlice():
				bases := qkd.NewRandomBits(len(qbts))
				for i, qbt := range qbts {
					if bases[i] {
						qbt.Observe(basis.Hadamard)
					} else {
						qbt.Observe(basis.Standard)
					}
				}
				in.QubitsToBob() <- qbts

			case bits := <-in.FromAlice():
				in.ToBob() <- bits

			case bits := <-in.FromBob():
				in.ToAlice() <- bits

			case <-eve.done:
				break loop
			}
		}
}