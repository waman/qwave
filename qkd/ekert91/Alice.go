package ekert91

import (
	"github.com/waman/qwave/system/qubit"
	"github.com/waman/qwave/qkd"
	"github.com/waman/qwave/system/qubit2"
	"github.com/waman/qwave/system/qubit/ket"
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
    q2s := createEntangledQubits(qkd.ProperBitCount)
    firsts, seconds := separateQubits(q2s)
		ch.Qch() <- seconds

		bases := qkd.NewRandomBits(qkd.ProperBitCount)
    bits := observeQubits(firsts, bases)

		matches := matchBases(bases, <- ch.FromBob())
		ch.ToBob() <- matches
		alice.key = qkd.AppendMatchingBits(alice.key, bits, matches, alice.n)
	}
}

func createEntangledQubits(n int) []*qubit2.Qubit2 {
	q2s := make([]*qubit2.Qubit2, qkd.ProperBitCount)
	for i, n := 0, len(q2s); i < n; i++ {
		q2s[i] = qubit2.NewPhiPlus()
	}
	return q2s
}

func separateQubits(q2s []*qubit2.Qubit2) (firsts, seconds []qubit.Qubit) {
	firsts = make([]qubit.Qubit, len(q2s))
	seconds = make([]qubit.Qubit, len(q2s))
	for i, q2 := range q2s {
		firsts[i] = q2.Qubit(0)
		seconds[i] = q2.Qubit(1)
	}
	return
}

func observeQubits(qbts []qubit.Qubit, bases []bool) []bool {
	bits := make([]bool, len(qbts))
	for i, n := 0, len(qbts); i < n; i++ {
		if bases[i] {
			bits[i] = qbts[i].ObserveInHadamardBasis() == ket.Minus()
		} else {
			bits[i] = qbts[i].ObserveInStandardBasis() == ket.One()
		}
	}
	return bits
}

func matchBases(bases, bobsBases []bool) []bool {
	var match = make([]bool, len(bases))
	for i, basis := range bases {
		match[i] = basis == bobsBases[i]
	}
	return match
}