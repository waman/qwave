package qubit

import (
	"math/rand"
	"github.com/waman/qwave/system/qubit/ket"
	"github.com/waman/qwave/system/qubit/basis"
	"sync"
	"github.com/waman/qwave/system/qubit/op"
)

type Qubit interface{
	Observe(b *basis.Basis)  *ket.State
	ObserveInStandardBasis() *ket.State
	ObserveInHadamardBasis() *ket.State

	Apply(u op.Matrix2x2)
}

type defaultQubit struct {
	mu    *sync.Mutex
	state *ket.State
}

func (qbt *defaultQubit) Observe(b *basis.Basis) *ket.State {
	qbt.mu.Lock()
	defer qbt.mu.Unlock()

  p := qbt.state.Probability(b.First())

  var newState *ket.State
  if r := rand.Float64(); r < p {
  	newState = b.First()
	}else{
		newState = b.Second()
	}

	qbt.state = newState
	return newState
}

func (qbt *defaultQubit) ObserveInStandardBasis() *ket.State {
	return qbt.Observe(basis.Standard())
}

func (qbt *defaultQubit) ObserveInHadamardBasis() *ket.State {
	return qbt.Observe(basis.Hadamard())
}

func (qbt *defaultQubit) Apply(u op.Matrix2x2){
	newA, newB := u.Apply(qbt.state.Coefficients())
	qbt.state = ket.New(newA, newB, true)  // TODO
}

// New function return the new Qubit object.
func New(a, b complex128, isNormalized bool) Qubit {
  return NewWith(ket.New(a, b, isNormalized))
}

func NewWith(state *ket.State) Qubit {
	var mu sync.Mutex
	return &defaultQubit{&mu, state}
}

func NewZero()   Qubit { return NewWith(ket.Zero()) }
func NewOne()    Qubit { return NewWith(ket.One()) }
func NewPlus()   Qubit { return NewWith(ket.Plus()) }
func NewMinus()  Qubit { return NewWith(ket.Minus()) }
func NewPlusI()  Qubit { return NewWith(ket.PlusI()) }
func NewMinusI() Qubit { return NewWith(ket.MinusI()) }

func U(m op.Matrix2x2, ms ...op.Matrix2x2) func(qbt Qubit) Qubit {
	mtrx := op.Prod(m, ms...)
	return func(qbt Qubit) Qubit {
		qbt.Apply(mtrx)
		return qbt
	}
}