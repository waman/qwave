package bb84

import (
	"github.com/waman/qwave/qubit/basis"
	"github.com/waman/qwave/qubit/state"
	"github.com/waman/qwave/qkd"
	"github.com/waman/qwave/qubit"
)

type Bob struct{
	n int
	key qkd.Key
}

func NewBob(n int) *Bob {
  return &Bob{n, nil}
}

func (bob *Bob) EstablishKey(ch qkd.ChannelOnBob, done chan<- struct{}){
	for len(bob.key) < bob.n {
		bases := qkd.NewRandomBit(bob.n)

		bits := observeQubits(bases, ch.Qch())
		ch.ToAlice() <- nil

		matches := matchBases(bases, ch.FromAlice())
		ch.ToAlice() <- matches
		bob.key = AppendMatchingBit(bob.key, bits, matches, bob.n)
	}
	done <- struct{}{}
}

func observeQubits(bases []bool, ch <-chan []qubit.Qubit) []bool {
	bits := make([]bool, len(bases))
	for i, qbt := range <- ch {
		if bases[i] {  // 1 -> observing by the Hadamard basis
			// |-> -> 1
			// |+> -> 0
			bits[i] = qbt.Observe(basis.Hadamard()) == state.Minus()
		}else{  // 0 -> observing by the standard basis
			// |1> -> 1
			// |0> -> 0
			bits[i] = qbt.Observe(basis.Standard()) == state.One()
		}
	}
	return bits
}

func matchBases(bases []bool, ch <-chan []bool) []bool {
	var match = make([]bool, len(bases))
	i := 0
	for _, trueBasis := range <- ch {
		match[i] = trueBasis == bases[i]
    i++
	}
	return match
}

func (bob *Bob) Key() qkd.Key {
	return bob.key
}