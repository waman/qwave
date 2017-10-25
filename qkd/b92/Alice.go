package b92

import (
	"github.com/waman/qwave/system/qubit"
	"github.com/waman/qwave/qkd"
)

func NewAlice(n int) *Alice {
	return &Alice{n, nil}
}

type Alice struct{
	n int
	key qkd.Key
}

func (alice *Alice) Key() qkd.Key {
	return alice.key
}

func (alice *Alice) EstablishKey(ch qkd.ChannelOnAlice){
	for len(alice.key) < alice.n {
		bits  := qkd.NewRandomBits(qkd.ProperBitCount)
		ch.Qch() <- encode(bits)

		matches := <- ch.FromBob()
		alice.key, _ = qkd.AppendMatchingBits(alice.key, bits, matches, alice.n)
	}
}

func encode(bits []bool) []qubit.Qubit {
	var qbts = make([]qubit.Qubit, len(bits))
	for i, bit := range bits {
		if bit {  // 1 -> |+>
			qbts[i] = qubit.NewPlus()
		} else {  // 0 -> |0>
			qbts[i] = qubit.NewZero()
		}
	}
	return qbts
}