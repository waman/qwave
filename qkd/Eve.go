package qkd

import "github.com/waman/qwave/system/qubit/basis"

type Eve interface {
	Eavesdrop(ch *InternalOfChannel)
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

func (eve *observingEve) Eavesdrop(ch *InternalOfChannel) {
	loop:
		for {
			select {
			case qbts := <-ch.QubitsFromAlice():
				bases := NewRandomBits(len(qbts))
				for i, qbt := range qbts {
					if bases[i] {
						qbt.Observe(basis.Hadamard())
					} else {
						qbt.Observe(basis.Standard())
					}
				}
				ch.QubitsToBob() <- qbts

			case bits := <-ch.FromBob():
				ch.ToAlice() <- bits

			case bits := <-ch.FromAlice():
				ch.ToBob() <- bits

			case <-eve.done:
				break loop
			}
		}
}