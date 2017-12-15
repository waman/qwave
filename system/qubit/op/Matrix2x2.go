package op

import (
	"fmt"
	"github.com/waman/qwave/system"
	"log"
	"math/cmplx"
)

type Matrix2x2 interface {
	Dim() int
	At(i, j int) complex128
	Components() (a, b, c, d complex128)
	A() complex128
	B() complex128
	C() complex128
	D() complex128

	Equals(v Matrix2x2) bool
	String() string

	Apply(x, y complex128) (complex128, complex128)
	CloneMutable() MutableMatrix2x2
}

func NewMatrix2x2(a, b, c, d complex128) Matrix2x2 {
	return &matrix2x2{a, b, c, d}
}

type matrix2x2 struct {
	a, b,
	c, d complex128
}

func (u *matrix2x2) Dim() int { return 2 }

func (u *matrix2x2) At(i, j int) complex128 {
	switch i {
	case 0:
		switch j {
		case 0: return u.a
		case 1: return u.b
		}

	case 1:
		switch j {
		case 0: return u.c
		case 1: return u.d
		}
	}
	log.Panicf("The indices must be in [0, 2): (%d, %d) appears.", i, j)
	return cmplx.NaN()
}

func (u *matrix2x2) Components() (a, b, c, d complex128) {
	a = u.a
	b = u.b
	c = u.c
	d = u.d
	return
}

func (u *matrix2x2) A() complex128 { return u.a }
func (u *matrix2x2) B() complex128 { return u.b }
func (u *matrix2x2) C() complex128 { return u.c }
func (u *matrix2x2) D() complex128 { return u.d }

func (u *matrix2x2) Equals(v Matrix2x2) bool {
	return u.a == v.A() && u.b == v.B() &&
		     u.c == v.C() && u.d == v.D()
}

func (u *matrix2x2) Apply(x, y complex128) (complex128, complex128) {
	return u.a*x + u.b*y,
	       u.c*x + u.d*y
}

func (u *matrix2x2) String() string {
	return fmt.Sprintf("[[%v %v] [%v %v]]", u.a, u.b, u.c, u.d)
}

func (u *matrix2x2) CloneMutable() MutableMatrix2x2 {
	return New(u.a, u.b, u.c, u.d)
}

var (
	id = NewMatrix2x2(1, 0, 0, 1)
	x = NewMatrix2x2(0, 1, 1, 0)
	y = NewMatrix2x2(0, 1, -1, 0)
	z = NewMatrix2x2(1, 0, 0, -1)
	h = NewMatrix2x2(system.S2I(), system.S2I(), system.S2I(), -system.S2I())
)

func Id() Matrix2x2 { return id }

func PauliX() Matrix2x2 { return x }
func PauliY() Matrix2x2 { return y }
func PauliZ() Matrix2x2 { return z }

func PauliMatrices() (Matrix2x2, Matrix2x2, Matrix2x2) {
	return x, y, z
}

func Hadamard() Matrix2x2 { return h }