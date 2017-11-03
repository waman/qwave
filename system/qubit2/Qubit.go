package qubit2

import (
	"github.com/waman/qwave/system/qubit2/basis2"
	"github.com/waman/qwave/system/qubit2/ket2"
	"math/rand"
	"sync"
	"github.com/waman/qwave/system/qubit/ket"
	"github.com/waman/qwave/system/qubit/basis"
	"github.com/waman/qwave/system/qubit"
	"log"
	"math/cmplx"
	"math"
)

type Qubit2 struct{
	mu    sync.Mutex
	state *ket2.State
}

func New(a, b, c, d complex128, isNormalized bool) *Qubit2 {
	return NewWith(ket2.New(a, b, c, d, isNormalized))
}

func NewWith(s *ket2.State) *Qubit2 {
	var mu sync.Mutex
	return &Qubit2{mu, s}
}

func NewZero() *Qubit2 {
	return NewWith(ket2.Zero())
}

func NewOne() *Qubit2 {
	return NewWith(ket2.One())
}

func NewTwo() *Qubit2 {
	return NewWith(ket2.Two())
}

func NewThree() *Qubit2 {
	return NewWith(ket2.Three())
}

func NewPhiPlus() *Qubit2 {
	return NewWith(ket2.PhiPlus())
}

func NewPhiMinus() *Qubit2 {
	return NewWith(ket2.PhiMinus())
}

func NewPsiPlus() *Qubit2 {
	return NewWith(ket2.PsiPlus())
}

func NewPsiMinus() *Qubit2 {
	return NewWith(ket2.PsiMinus())
}

func (qbt2 *Qubit2) Observe(basis *basis2.Basis) *ket2.State {
	r := rand.Float64()

	qbt2.mu.Lock()
	defer qbt2.mu.Unlock()

	p := 0.0
	for i := 0; i < 3; i++ {
		s := basis.Get(i)
		p += qbt2.state.Probability(s)
		if r < p {
			qbt2.state = s
			return s
		}
	}

	s := basis.Get(3)
	qbt2.state = s
	return s
}

func (qbt2 *Qubit2) ObserveQubit(i int, b *basis.Basis) *ket.State {
  if b == basis.Standard() {
  	return qbt2.ObserveQubitInStandardBasis(i)
	} else if b == basis.Hadamard() {
		return qbt2.ObserveQubitInHadamardBasis(i)
	}
	panic("implement me.")
	return nil
}

func (qbt2 *Qubit2) ObserveQubitInStandardBasis(i int) *ket.State {
	r := rand.Float64()

	qbt2.mu.Lock()
	defer qbt2.mu.Unlock()

	var c0, c1, c2, c3 complex128
  if i == 0 {
  	c0, c1, c2, c3 = qbt2.state.Coefficients()
	} else if i == 1 {
		c0, c2, c1, c3 = qbt2.state.Coefficients()
	} else {
		log.Panicf("Index is out of bounds [0, 2): %d", i)
	}

	c0Abs, c1Abs := cmplx.Abs(c0), cmplx.Abs(c1)
	p := c0Abs*c0Abs + c1Abs*c1Abs  // = |c0|^2 + |c1|^2
	if r < p {
		norm := complex(math.Sqrt(p), 0)
		qbt2.state = ket2.New(c0/norm, c1/norm, 0, 0, true)
		return ket.Zero()

	} else {
		norm := complex(math.Sqrt(1-p), 0)
		qbt2.state = ket2.New(0, 0, c2/norm, c3/norm, true)
		return ket.One()
	}
}

func (qbt2 *Qubit2) ObserveQubitInHadamardBasis(i int) *ket.State {
	r := rand.Float64()

	qbt2.mu.Lock()
	defer qbt2.mu.Unlock()

	var c0, c1, c2, c3 complex128
	if i == 0 {
		c0, c1, c2, c3 = qbt2.state.Coefficients()
	} else if i == 1 {
		c0, c2, c1, c3 = qbt2.state.Coefficients()
	} else {
		log.Panicf("Index is out of bounds [0, 2): %d", i)
	}

	c02Abs, c13Abs := cmplx.Abs(c0+c2), cmplx.Abs(c1+c3)
	p := (c02Abs*c02Abs + c13Abs*c13Abs)/2
	if r < p {
		norm := complex(2*math.Sqrt(p), 0)
		c02, c13 := (c0+c2)/norm, (c1+c3)/norm
		qbt2.state = ket2.New(c02, c13, c02, c13, true)
		return ket.Plus()

	} else {
		norm := complex(2*math.Sqrt(1-p), 0)
		c02, c13 := (c0-c2)/norm, (c1-c3)/norm
		qbt2.state = ket2.New(c02, c13, -c02, -c13, true)
		return ket.Minus()
	}
}

func (qbt2 *Qubit2) Qubit(i int) qubit.Qubit {
	return &qubitImpl{qbt2, i}
}

type qubitImpl struct {
	parent *Qubit2
	i      int
}

func (qbt *qubitImpl) Observe(basis *basis.Basis) *ket.State {
	return qbt.parent.ObserveQubit(qbt.i, basis)
}

func (qbt *qubitImpl) ObserveInStandardBasis() *ket.State {
	return qbt.parent.ObserveQubitInStandardBasis(qbt.i)
}

func (qbt *qubitImpl) ObserveInHadamardBasis() *ket.State {
	return qbt.parent.ObserveQubitInHadamardBasis(qbt.i)
}


