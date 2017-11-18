package qubit

import (
	"github.com/waman/qwave/system/qubit/ket"
	. "github.com/waman/qwave/system"
)

type Matrix2 interface {
	Apply(x, y complex128) (complex128, complex128)
	ApplyToState(s *ket.State) *ket.State
}

type defaultMatrix2 struct {
	a, b,
	c, d complex128
}

func (u *defaultMatrix2) Apply(x, y complex128) (complex128, complex128) {
	return u.a*x + u.b*y,
	       u.c*x + u.d*y
}

func (u *defaultMatrix2) ApplyToState(s *ket.State) *ket.State {
	newA, newB := u.Apply(s.Coefficients())
	return ket.New(newA, newB, true)
}

var(
	x = &defaultMatrix2{0, 1, 1, 0}
	y = &defaultMatrix2{0, 1, -1, 0}
	z = &defaultMatrix2{1, 0, 0, -1}
	h = &defaultMatrix2{S2I(), S2I(), S2I(), -S2I()}
)

func I(qbt Qubit) Qubit {
	return qbt
}

func X(qbt Qubit) Qubit {
	qbt.Apply(x)
	return qbt
}

func Y(qbt Qubit) Qubit {
	qbt.Apply(y)
	return qbt
}

func Z(qbt Qubit) Qubit {
	qbt.Apply(z)
	return qbt
}

func H(qbt Qubit) Qubit {
	qbt.Apply(h)
	return qbt
}
