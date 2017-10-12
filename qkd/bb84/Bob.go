package bb84

import (
	"github.com/waman/qwave/qubit/basis"
	"github.com/waman/qwave/qubit/state"
)

type Bob struct{
	n int
	key *[]bool
}

func NewBob(n int,) *Bob {
  return &Bob{n, nil}
}

func (bob *Bob) EstablishKey(ch ChannelOnBob){
  bases := NewRandomBit(bob.n)
  bits  := make([]bool, bob.n)
  i := 0
  for qb := range ch.Qch() {
  	if bases[i] {  // 1 -> observing by the Hadamard basis
  	  // |-> -> 1
  	  // |+> -> 0
  		bits[i] = qb.Observe(basis.Hadamard()) == state.Minus()
		}else{  // 0 -> observing by the standard basis
			// |1> -> 1
			// |0> -> 0
			bits[i] = qb.Observe(basis.Standard()) == state.One()
		}
		i++
	}

	ch.ToAlice() <- true
}

func (bob *Bob) Key() *[]bool {
	return bob.key
}