package qkd

import "github.com/waman/qwave/system/qubit"

type InsecureChannel struct {
	withAlice, withBob Channel
}

func NewInsecureChannel() *InsecureChannel {
	return &InsecureChannel{NewChannel(), NewChannel()}
}

func (ich *InsecureChannel) OnAlice() ChannelOnAlice {
	return ich.withAlice.OnAlice()
}

func (ich *InsecureChannel) OnBob() ChannelOnBob {
	return ich.withBob.OnBob()
}

func (ich *InsecureChannel) Close() {
	ich.withAlice.Close()
	ich.withBob.Close()
}

func (ich *InsecureChannel) Internal() *InternalOfChannel {
	return &InternalOfChannel{ich.withAlice.OnBob(), ich.withBob.OnAlice()}
}

type InternalOfChannel struct {
	withAlice ChannelOnBob
	withBob   ChannelOnAlice
}

func (in *InternalOfChannel) WithAlice() ChannelOnBob {
	return in.withAlice
}

func (in *InternalOfChannel) WithBob() ChannelOnAlice {
	return in.withBob
}

func (in *InternalOfChannel) QubitsFromAlice() <-chan []qubit.Qubit {
	return in.withAlice.Qch()
}

func (in *InternalOfChannel) QubitsToBob() chan<- []qubit.Qubit {
	return in.withBob.Qch()
}

func (in *InternalOfChannel) FromAlice() <-chan []bool {
	return in.withAlice.FromAlice()
}

func (in *InternalOfChannel) ToAlice() chan<- []bool {
	return in.withAlice.ToAlice()
}

func (in *InternalOfChannel) FromBob() <-chan []bool {
	return in.withBob.FromBob()
}

func (in *InternalOfChannel) ToBob() chan<- []bool {
	return in.withBob.ToBob()
}

func (in *InternalOfChannel) ForwardQubits(){
	in.withBob.Qch() <- <- in.withAlice.Qch()
}

func (in *InternalOfChannel) ForwardAtoB(){
	in.withBob.ToBob() <- <- in.withAlice.FromAlice()
}

func (in *InternalOfChannel) ForwardBtoA(){
	in.withAlice.ToAlice() <- <- in.withBob.FromBob()
}

func (in *InternalOfChannel) EavsedropQubits(f func([]qubit.Qubit)){
	qbts := <- in.withAlice.Qch()
	f(qbts)
	in.withBob.Qch() <- qbts
}

func (in *InternalOfChannel) EavsedropAtoB(f func([]bool)){
	bits := <- in.withAlice.FromAlice()
	f(bits)
	in.withBob.ToBob() <- bits
}

func (in *InternalOfChannel) EavsedropBtoA(f func([]bool)){
	bits := <- in.withBob.FromBob()
	f(bits)
	in.withAlice.ToAlice() <- bits
}


func (in *InternalOfChannel) ManipulateQubits(f func([]qubit.Qubit)[]qubit.Qubit){
	qbts := <- in.withAlice.Qch()
	in.withBob.Qch() <- f(qbts)
}

func (in *InternalOfChannel) ManipulateAtoB(f func([]bool)[]bool){
	bits := <- in.withAlice.FromAlice()
	in.withBob.ToBob() <- f(bits)
}

func (in *InternalOfChannel) ManipulateBtoA(f func([]bool)[]bool){
	bits := <- in.withBob.FromBob()
	in.withAlice.ToAlice() <- f(bits)
}

