package qkd

type InsecureChannel interface {
	Channel
	Internal() (chWithAlice ChannelOnBob, chWithBob ChannelOnAlice)
}

func NewInsecureChannel() InsecureChannel {
	return &insecureChannel{NewChannel(), NewChannel()}
}

type insecureChannel struct {
	withAlice, withBob Channel
}

func (ch *insecureChannel) OnAlice() ChannelOnAlice {
	return ch.withAlice.OnAlice()
}

func (ch *insecureChannel) OnBob() ChannelOnBob {
	return ch.withBob.OnBob()
}

func (ch *insecureChannel) Close() {
	ch.withAlice.Close()
	ch.withBob.Close()
}

func (ch *insecureChannel) Internal() (chWithAlice ChannelOnBob, chWithBob ChannelOnAlice) {
	return ch.withAlice.OnBob(), ch.withBob.OnAlice()
}
