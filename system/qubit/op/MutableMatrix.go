package op

import (
	"log"
	"math/cmplx"
	"fmt"
)

type MutableMatrix2x2 interface {
	Matrix2x2
	Mul(x, y Matrix2x2) Matrix2x2
}

func New(a, b, c, d complex128) MutableMatrix2x2 {
	return &mutableMatrix2x2{a, b, c, d}
}

type mutableMatrix2x2 struct {
	a, b,
	c, d complex128
}

func (u *mutableMatrix2x2) Dim() int { return 2 }

func (u *mutableMatrix2x2) At(i, j int) complex128 {
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

func (u *mutableMatrix2x2) Components() (a, b, c, d complex128) {
	a = u.a
	b = u.b
	c = u.c
	d = u.d
	return
}

func (u *mutableMatrix2x2) A() complex128 { return u.a }
func (u *mutableMatrix2x2) B() complex128 { return u.b }
func (u *mutableMatrix2x2) C() complex128 { return u.c }
func (u *mutableMatrix2x2) D() complex128 { return u.d }

func (u *mutableMatrix2x2) Equals(v Matrix2x2) bool {
	return u.a == v.A() && u.b == v.B() &&
		     u.c == v.C() && u.d == v.D()
}

func (u *mutableMatrix2x2) String() string {
	return fmt.Sprintf("[[%v %v] [%v %v]]", u.a, u.b, u.c, u.d)
}

func (u *mutableMatrix2x2) Apply(x, y complex128) (complex128, complex128) {
	return u.a*x + u.b*y,
		     u.c*x + u.d*y
}

func (u *mutableMatrix2x2) CloneMutable() MutableMatrix2x2 {
	return New(u.a, u.b, u.c, u.d)
}

func (z *mutableMatrix2x2) Mul(x, y Matrix2x2) Matrix2x2 {
	z.a, z.b,
	z.c, z.d =
		x.A()*y.A() + x.B()*y.C(), x.A()*y.B() + x.B()*y.D(),
		x.C()*y.A() + x.D()*y.C(), x.C()*y.B() + x.D()*y.D()

	return z
}

func Prod(m Matrix2x2, ms ...Matrix2x2) Matrix2x2 {
	if len(ms) == 0 { return m }

	var result = m.CloneMutable()
	for _, m := range ms {
		result.Mul(result, m)
	}

	return result
}
