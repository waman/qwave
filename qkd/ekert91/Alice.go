package ekert91

import (
	"github.com/waman/qwave/system/qubit"
	"github.com/waman/qwave/qkd"
	"github.com/waman/qwave/system/qubit/ket"
	"github.com/waman/qwave/system/qubits2"
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
    firsts, seconds := createEntangledQubits(qkd.ProperBitCount)
		ch.Qch() <- seconds

		bases := qkd.NewRandomBits(qkd.ProperBitCount)
    bits := observeQubits(firsts, bases)

		matches := matchBases(bases, <- ch.FromBob())
		ch.ToBob() <- matches
		alice.key = qkd.AppendMatchingBits(alice.key, bits, matches, alice.n)
	}
}

func createEntangledQubits(n int) (firsts, seconds []qubit.Qubit) {
	firsts = make([]qubit.Qubit, n)
	seconds = make([]qubit.Qubit, n)
	for i := 0; i < n; i++ {
		firsts[i], seconds[i] = qubits2.NewPhiPlus().Split()
	}
	return
}

func observeQubits(qbts []qubit.Qubit, bases []bool) []bool {
	n := len(qbts)
	bits := make([]bool, n)
	for i := 0; i < n; i++ {
		if bases[i] {
			bits[i] = qbts[i].ObserveInHadamardBasis() == ket.Minus()
		} else {
			bits[i] = qbts[i].ObserveInStandardBasis() == ket.One()
		}
	}
	return bits
}

func matchBases(bases, bobsBases []bool) []bool {
	matches := make([]bool, len(bases))
	for i, basis := range bases {
		matches[i] = basis == bobsBases[i]
	}
	return matches
}