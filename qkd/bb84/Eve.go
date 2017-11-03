package bb84

import (
	"github.com/waman/qwave/qkd"
	"github.com/waman/qwave/system/qubit/ket"
)

func NewEve(n int) *Eve {
	return &Eve{n, nil, nil, make(chan struct{})}
}

type Eve struct {
	n int
	key qkd.Key
	sureKeyBits []bool
	done chan struct{}
}

func (eve *Eve) Key() qkd.Key {
	return eve.key
}

func (eve *Eve) SureKeyBitCount() int {
	sure := 0
	for i := 0; i < eve.n; i++ {
		if eve.sureKeyBits[i] { sure++ }
	}
	return sure
}

func (eve *Eve) Stop(){
	eve.done <- struct{}{}
}

func (eve *Eve) Eavesdrop(ch *qkd.InternalOfChannel) {
	var keyBits []bool
	var bases []bool
	var matchesBE []bool  // matches betwenn Bob and Eve
	loop:
		for {
			select {
			case qbts := <-ch.QubitsFromAlice():
				keyBits = make([]bool, len(qbts))
				bases = qkd.NewRandomBits(len(qbts))

				for i, qbt := range qbts {
					if bases[i] {
						keyBits[i] = qbt.ObserveInHadamardBasis() == ket.Minus()
					} else {
						keyBits[i] = qbt.ObserveInStandardBasis() == ket.One()
					}
				}
				ch.QubitsToBob() <- qbts

			case bobBases := <-ch.FromBob():
				matchesBE = matchBases(bases, bobBases)
				ch.ToAlice() <- bobBases

			case matches := <-ch.FromAlice():
				eve.key = qkd.AppendMatchingBits(eve.key, keyBits, matches, eve.n)
				eve.sureKeyBits = qkd.AppendMatchingBits(eve.sureKeyBits, matchesBE, matches, eve.n)
				ch.ToBob() <- matches

			case <-eve.done:
				break loop
			}
		}
}