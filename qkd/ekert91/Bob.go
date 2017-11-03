package ekert91

import (
	"github.com/waman/qwave/qkd"
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
		bits := observeQubits(qbts, bases)

		ch.ToAlice() <- bases

		matches := <- ch.FromAlice()
		bob.key = qkd.AppendMatchingBits(bob.key, bits, matches, bob.n)
	}
}