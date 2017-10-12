package bb84

import (
	"github.com/waman/qwave/qubit"
)

type Channel struct {
	qch chan qubit.Qubit
	aliceToBob chan bool
	bobToAlice chan bool
}

func NewChannel(n int) *Channel {
	qch := make(chan qubit.Qubit, n)
	aliceToBob := make(chan bool, n)
	bobToAlice := make(chan bool, n)
	return &Channel{qch, aliceToBob, bobToAlice}
}

//***** OnAlice *****
func (ch *Channel) OnAlice() ChannelOnAlice {
	return &chOnAlice{ch.qch, ch.aliceToBob, ch.bobToAlice}
}

type ChannelOnAlice interface {
	Qch() chan<- qubit.Qubit
	ToBob() chan<- bool
	FromBob() <-chan bool
}

type chOnAlice struct {
	qch chan<- qubit.Qubit
	toBob chan<- bool
	fromBom <-chan bool
}

func (ch *chOnAlice) Qch() chan<- qubit.Qubit {
	return ch.qch
}

func (ch *chOnAlice) ToBob() chan<- bool {
	return ch.toBob
}

func (ch *chOnAlice) FromBob() <-chan bool {
	return ch.fromBom
}

//***** OnBob *****
type ChannelOnBob interface {
	Qch() <-chan qubit.Qubit
	ToAlice() chan<- bool
	FromAlice() <-chan bool
}

func (ch *Channel) OnBob() ChannelOnBob {
	return &chOnBob{ch.qch, ch.aliceToBob, ch.bobToAlice}
}

type chOnBob struct {
	qch <-chan qubit.Qubit
	toAlice chan<- bool
	fromAlice <-chan bool
}

func (ch *chOnBob) Qch() <-chan qubit.Qubit {
	return ch.qch
}

func (ch *chOnBob) ToAlice() chan<- bool {
	return ch.toAlice
}

func (ch *chOnBob) FromAlice() <-chan bool {
	return ch.fromAlice
}
