package qkd

import (
	"github.com/waman/qwave/system/qubit/basis"
)

// NewObservingEve function returns an Eve observing qubits in
// the standard basis or the Hadamard basis randomly.
func NewObservingEve() Eve {
	return &observingEve{make(chan struct{})}
}

type Eve interface {
	Eavesdrop(ch *InternalOfChannel)
	Stop()
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
						qbt.ObserveInHadamardBasis()
					} else {
						qbt.ObserveInStandardBasis()
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

// NewObservingEveR function returns an Eve observing qubits in
// real random bases.
func NewObservingEveR() Eve {
	return &observingEveR{make(chan struct{})}
}

type EveR interface {
	Eavesdrop(ch *InternalOfChannel)
	Stop()
}

type observingEveR struct {
	done chan struct{}
}

func (eve *observingEveR) Stop(){
	eve.done <- struct{}{}
}

func (eve *observingEveR) Eavesdrop(ch *InternalOfChannel) {
	loop:
		for {
			select {
			case qbts := <-ch.QubitsFromAlice():
				for _, qbt := range qbts {
					qbt.Observe(basis.NewRandomRealBasis())
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