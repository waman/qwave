package qkd

type Eve interface {
	Key() Key
	Eavsedrop(chWithAlice ChannelOnBob, chWithBob ChannelOnAlice)
}
