package qkd

import (
	"math/rand"
	"github.com/waman/qwave/system/qubit"
	"log"
)

//func NewRandomBitSimply(n int) []bool {
//	bs := make([]bool, n)
//	for i := 0; i < n; i++ {
//		bs[i] = rand.Intn(2) == 1  // [0, 2) => 0 or 1
//	}
//	return bs
//}

const pcBitMinus1 = (32 << (^uint(0) >> 63)) - 1 // == 31 or 63

func NewRandomBit(n int) []bool {
	bs := make([]bool, n)
	index := 0
	for {
		var max int
		var isLast bool
		if m := n-index; m <= pcBitMinus1 {
			max = m
			isLast = true
		}else{
			max = pcBitMinus1
			isLast = false
		}

		r := rand.Int()
		for i := 0; i < max; i++ {
			bs[index] = r & 1 == 1
			r >>= 1
			index++
		}

		if isLast { break }
	}
	return bs
}

func EstablishKey(kc KeyContainer, ch Channel, done chan<- struct{}){
	if alice, ok := kc.(Alice); ok {
		alice.EstablishKey(ch.OnAlice())
		done <- struct{}{}

	} else if bob, ok := kc.(Bob); ok {
		bob.EstablishKey(ch.OnBob())
		done <- struct{}{}

	}else{
		log.Panicf("KeyContainer must be qkd.Alice or qkd.Bob: %T", kc)
	}
}

func EavseDrop(eve Eve, ch *InsecureChannel, done chan<- struct{}){
	eve.Eavsedrop(ch.Internal())
	done <- struct{}{}
}

func ManipulateQch(
	  from ChannelOnBob, to ChannelOnAlice, f func(qubits []qubit.Qubit)[]qubit.Qubit){
	qubits := <- from.Qch()
	to.Qch() <- f(qubits)
}

func ManipulateAlicesMessage(
	  from ChannelOnBob, to ChannelOnAlice, f func(qubits []bool)[]bool){
	bits := <- from.FromAlice()
	to.ToBob() <- f(bits)
}

func ManipulateBobsMessage(
	  from ChannelOnAlice, to ChannelOnBob, f func(qubits []bool)[]bool){
	bits := <- from.FromBob()
	to.ToAlice() <- f(bits)
}