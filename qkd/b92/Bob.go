package b92

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
		bits := qkd.NewRandomBits(len(qbts))

		resultBits := decode(qbts, bits)
		ch.ToAlice() <- resultBits

		bob.key = qkd.AppendMatchingBits(bob.key, bits, resultBits, bob.n)
	}
}

func decode(qbts []qubit.Qubit, bits []bool) []bool {
	resultBits := make([]bool, len(qbts))
	for i, qbt := range qbts {
		if bits[i] {  // observe in the standard basis
			resultBits[i] = qbt.ObserveInStandardBasis() == ket.One()
		} else {  // observe in the Hadamard basis
			resultBits[i] = qbt.ObserveInHadamardBasis() == ket.Minus()
		}
	}
	return resultBits
}