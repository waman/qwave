package qubit

import (
	"math/rand"
	"github.com/waman/qwave/system/qubit/ket"
	"github.com/waman/qwave/system/qubit/basis"
	"sync"
)

type Qubit interface{
	Observe(basis *basis.Basis) *ket.State
	ObserveInStandardBasis() *ket.State
	ObserveInHadamardBasis() *ket.State
}

type defaultQubit struct {
	mu    sync.Mutex
	state *ket.State
}

func (qbt *defaultQubit) Observe(basis *basis.Basis) *ket.State {
	qbt.mu.Lock()
	defer qbt.mu.Unlock()

  prob := qbt.state.Probability(basis.First())

  var nextState *ket.State
  if p := rand.Float64(); p < prob {
  	nextState = basis.First()
	}else{
		nextState = basis.Second()
	}

	qbt.state = nextState
	return nextState
}

func (qbt *defaultQubit) ObserveInStandardBasis() *ket.State {
	return qbt.Observe(basis.Standard())
}

func (qbt *defaultQubit) ObserveInHadamardBasis() *ket.State {
	return qbt.Observe(basis.Hadamard())
}

func New(a, b complex128, isNormalized bool) Qubit {
  return NewWith(ket.New(a, b, isNormalized))
}

func NewWith(state *ket.State) Qubit {
	var mu sync.Mutex
	return &defaultQubit{mu, state}
}

func NewZero()   Qubit { return NewWith(ket.Zero()) }
func NewOne()    Qubit { return NewWith(ket.One()) }
func NewPlus()   Qubit { return NewWith(ket.Plus()) }
func NewMinus()  Qubit { return NewWith(ket.Minus()) }
func NewPlusI()  Qubit { return NewWith(ket.PlusI()) }
func NewMinusI() Qubit { return NewWith(ket.MinusI()) }