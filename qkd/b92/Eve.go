package b92

import (
	"github.com/waman/qwave/qkd"
	"github.com/waman/qwave/system/qubit/ket"
)

func NewEve(n int) *Eve {
	return &Eve{n, nil, make(chan struct{})}
}

type Eve struct {
	n int
	key qkd.Key
	done chan struct{}
}

func (eve *Eve) Key() qkd.Key { return eve.key }
func (eve *Eve) Stop(){ eve.done <- struct{}{} }

func (eve *Eve) Eavesdrop(ch *qkd.InternalOfChannel) {
	var keyBits []bool
	loop:
		for {
			select {
			case qbts := <-ch.QubitsFromAlice():
				keyBits = make([]bool, len(qbts))
				bits := qkd.NewRandomBits(len(qbts))

				for i, qbt := range qbts {
					if bits[i] {
						keyBits[i] = qbt.ObserveInStandardBasis() == ket.One()
					} else {
						keyBits[i] = qbt.ObserveInHadamardBasis() == ket.Plus()
					}
				}
				ch.QubitsToBob() <- qbts

			case resultBits := <-ch.FromBob():
				eve.key = qkd.AppendMatchingBits(eve.key, keyBits, resultBits, eve.n)
				ch.ToAlice() <- resultBits

			case bits := <-ch.FromAlice():
				ch.ToBob() <- bits

			case <-eve.done:
				break loop
			}
		}
}