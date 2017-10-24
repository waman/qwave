package bb84

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
		bases := qkd.NewRandomBits(qkd.ProperBitCount)

		ch.Qch() <- encode(bits, bases)
		matches := matchBases(bases, <- ch.FromBob())
		ch.ToBob() <- matches
		alice.key = qkd.AppendMatchingBits(alice.key, bits, matches, alice.n)
	}
}

func encode(bits, bases []bool) []qubit.Qubit {
	var qubits = make([]qubit.Qubit, len(bits))
	for i, bit := range bits {
		if bases[i] {  // 1 -> encoding by the Hadamard basis
			// 1 -> |->
			// 0 -> |+>
			if bit {
				qubits[i] = qubit.NewMinus()
			}else{
				qubits[i] = qubit.NewPlus()
			}

		}else{  // 0 -> encoding by the standard basis
			// 1 -> |1>
			// 0 -> |0>
			if bit {
				qubits[i] = qubit.NewOne()
			}else{
				qubits[i] = qubit.NewZero()
			}
		}
	}
	return qubits
}

func matchBases(bases, bobsBases []bool) []bool {
	var match = make([]bool, len(bases))
	for i, basis := range bases {
		match[i] = basis == bobsBases[i]
	}
	return match
}