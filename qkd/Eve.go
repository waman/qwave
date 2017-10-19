package qkd

type Eve interface {
	Key() Key
	Eavsedrop(in *InternalOfChannel)
}
