package qkd

import (
	"github.com/waman/qwave/qubit"
)

type Channel struct {
	qch chan []qubit.Qubit
	aliceToBob chan []bool
	bobToAlice chan []bool
}

func NewChannel() *Channel {
	qch := make(chan []qubit.Qubit)
	aliceToBob := make(chan []bool)
	bobToAlice := make(chan []bool)
	return &Channel{qch, aliceToBob, bobToAlice}
}

func (ch *Channel) Close(){
	close(ch.qch)
	close(ch.aliceToBob)
	close(ch.bobToAlice)
}

//***** OnAlice *****
func (ch *Channel) OnAlice() ChannelOnAlice {
	return &chOnAlice{ch.qch, ch.aliceToBob, ch.bobToAlice}
}

type ChannelOnAlice interface {
	Qch() chan<- []qubit.Qubit
	ToBob() chan<- []bool
	FromBob() <-chan []bool
}

type chOnAlice struct {
	qch chan<- []qubit.Qubit
	toBob chan<- []bool
	fromBob <-chan []bool
}

func (ch *chOnAlice) Qch() chan<- []qubit.Qubit {
	return ch.qch
}

func (ch *chOnAlice) ToBob() chan<- []bool {
	return ch.toBob
}

func (ch *chOnAlice) FromBob() <-chan []bool {
	return ch.fromBob
}

//***** OnBob *****
type ChannelOnBob interface {
	Qch() <-chan []qubit.Qubit
	ToAlice() chan<- []bool
	FromAlice() <-chan []bool
}

func (ch *Channel) OnBob() ChannelOnBob {
	return &chOnBob{ch.qch, ch.bobToAlice, ch.aliceToBob}
}

type chOnBob struct {
	qch <-chan []qubit.Qubit
	toAlice chan<- []bool
	fromAlice <-chan []bool
}

func (ch *chOnBob) Qch() <-chan []qubit.Qubit {
	return ch.qch
}

func (ch *chOnBob) ToAlice() chan<- []bool {
	return ch.toAlice
}

func (ch *chOnBob) FromAlice() <-chan []bool {
	return ch.fromAlice
}
