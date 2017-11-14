package bb84

import (
	"github.com/waman/qwave/qkd"
	"github.com/waman/qwave/system/qubit/basis"
)

func NewEveR(n int) qkd.Eve {
	return &EveR{n, nil, make(chan struct{})}
}

type EveR struct {
	n int
	key qkd.Key
	done chan struct{}
}

func (eve *EveR) Key() qkd.Key {
	return eve.key
}

func (eve *EveR) Stop(){
	eve.done <- struct{}{}
}

func (eve *EveR) Eavesdrop(ch *qkd.InternalOfChannel) {
	var keyBits []bool
	loop:
		for {
			select {
			case qbts := <-ch.QubitsFromAlice():
				keyBits = make([]bool, len(qbts))
				for i, qbt := range qbts {
					b := basis.NewRandomRealBasis()
					keyBits[i] = qbt.Observe(b) == b.Second()
				}
				ch.QubitsToBob() <- qbts

			case bobBases := <-ch.FromBob():
				ch.ToAlice() <- bobBases

			case matches := <-ch.FromAlice():
				eve.key = qkd.AppendMatchingBits(eve.key, keyBits, matches, eve.n)
				ch.ToBob() <- matches

			case <-eve.done:
				break loop
			}
		}
}