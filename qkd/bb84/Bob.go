package bb84

import (
	"github.com/waman/qwave/system/qubit/ket"
	"github.com/waman/qwave/qkd"
	"github.com/waman/qwave/system/qubit"
)

func NewBob(n int) *Bob {
	return &Bob{n, nil}
}

type Bob struct{
	n int
	key qkd.Key
}

func (bob *Bob) Key() qkd.Key {
	return bob.key
}

func (bob *Bob) EstablishKey(ch qkd.ChannelOnBob){
	for len(bob.key) < bob.n {
		qbts := <-ch.Qch()
		bases := qkd.NewRandomBits(len(qbts))
		bits := decode(qbts, bases)
		ch.ToAlice() <- bases

		matches := <- ch.FromAlice()
		bob.key = qkd.AppendMatchingBits(bob.key, bits, matches, bob.n)
	}
}

func decode(qbts []qubit.Qubit, bases []bool) []bool {
	bits := make([]bool, len(qbts))
	for i, qbt := range qbts {
		if bases[i] {  // 1 -> observing by the Hadamard basis
			// |-> -> 1
			// |+> -> 0
			bits[i] = qbt.ObserveInHadamardBasis() == ket.Minus()
		}else{  // 0 -> observing by the standard basis
			// |1> -> 1
			// |0> -> 0
			bits[i] = qbt.ObserveInStandardBasis() == ket.One()
		}
	}
	return bits
}