package qkd

type Eve interface {
	Eavesdrop(in *InternalOfChannel)
	Stop()
}
