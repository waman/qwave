package bb84

import (
	"github.com/waman/qwave/qubit"
	"github.com/waman/qwave/qkd"
)

type Alice struct{
	n int
	key qkd.Key
}

func NewAlice(n int) *Alice {
  return &Alice{n, nil}
}

func (alice *Alice) EstablishKey(ch qkd.ChannelOnAlice, done chan<- struct{}){
	for len(alice.key) < alice.n {
		bits  := qkd.NewRandomBitSimply(alice.n)
		bases := qkd.NewRandomBitSimply(alice.n)

		sendQubits(bits, bases, ch.Qch())
		<- ch.FromBob()

		ch.ToBob() <- bases
		matches := <- ch.FromBob()
		alice.key = AppendMatchingBit(alice.key, bits, matches, alice.n)
	}
	done <- struct{}{}
}

func sendQubits(bits, bases []bool, qch chan<- []qubit.Qubit){
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
	qch <- qubits
}

func (alice *Alice) Key() qkd.Key {
  return alice.key
}