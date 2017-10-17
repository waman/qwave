package qkd

type Alice interface {
	Key() Key
	EstablishKey(ch ChannelOnAlice)
}