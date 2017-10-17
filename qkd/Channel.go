package qkd

import (
	"github.com/waman/qwave/system/qubit"
)

type Channel interface {
	OnAlice() ChannelOnAlice
	OnBob()   ChannelOnBob
	Close()
}

type channel struct {
	qch chan []qubit.Qubit
	aliceToBob chan []bool
	bobToAlice chan []bool
}

func NewChannel() Channel {
	qch := make(chan []qubit.Qubit)
	aliceToBob := make(chan []bool)
	bobToAlice := make(chan []bool)
	return &channel{qch, aliceToBob, bobToAlice}
}

func (ch *channel) OnAlice() ChannelOnAlice {
	return &chOnAlice{ch.qch, ch.aliceToBob, ch.bobToAlice}
}

func (ch *channel) OnBob() ChannelOnBob {
	return &chOnBob{ch.qch, ch.bobToAlice, ch.aliceToBob}
}

func (ch *channel) Close(){
	close(ch.qch)
	close(ch.aliceToBob)
	close(ch.bobToAlice)
}

//***** OnAlice *****
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
