package qkd

type Bob interface {
	Key() Key
	EstablishKey(ch ChannelOnBob)
}