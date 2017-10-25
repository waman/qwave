package qkd

import "github.com/waman/qwave/system/qubit/basis"

type Eve interface {
	Eavesdrop(in *InternalOfChannel)
	Stop()
}

func NewObservingEve() Eve {
	return &observingEve{make(chan struct{})}
}

type observingEve struct {
	done chan struct{}
}

func (eve *observingEve) Stop(){
	eve.done <- struct{}{}
}

func (eve *observingEve) Eavesdrop(in *InternalOfChannel) {
	loop:
		for {
			select {
			case qbts := <-in.QubitsFromAlice():
				bases := NewRandomBits(len(qbts))
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