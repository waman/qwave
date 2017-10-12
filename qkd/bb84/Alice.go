package bb84

import (
	"github.com/waman/qwave/qubit"
)

type Alice struct{
	n int
	key []bool
}

func NewAlice(n int) *Alice {
  return &Alice{n, nil}
}

func (alice *Alice) EstablishKey(ch ChannelOnAlice, done chan<- struct{}){

	bits  := NewRandomBit(alice.n)
	bases := NewRandomBit(alice.n)
	sendQubits(bits, bases, ch.Qch())
}

func sendQubits(bits, bases []bool, qch chan<- qubit.Qubit){
	for i, bit := range bits {
		if bases[i] {  // 1 -> encoding by the Hadamard basis
			// 1 -> |->
			// 0 -> |+>
			if bit {
				qch <- qubit.NewMinus()
			}else{
				qch <- qubit.NewPlus()
			}

		}else{  // 0 -> encoding by the standard basis
			// 1 -> |1>
			// 0 -> |0>
			if bit {
				qch <- qubit.NewOne()
			}else{
				qch <- qubit.NewZero()
			}
		}
	}
}

func (alice *Alice) Key() []bool {
  return alice.key
}