package qubit

import (
	"math/rand"
	"github.com/waman/qwave/qubit/state"
	"github.com/waman/qwave/qubit/basis"
	"sync"
)

type Qubit interface{
	State() *state.State
	Observe(basis *basis.Basis) *state.State
}

func (dq *defaultQubit) State() *state.State {
	return dq.state
}

type defaultQubit struct {
	mu    sync.Mutex
	state *state.State
}

func (qbt *defaultQubit) Observe(basis *basis.Basis) *state.State {
	qbt.mu.Lock()
	defer qbt.mu.Unlock()

  prob := qbt.state.Probability(basis.First())

  var nextState *state.State
  if p := rand.Float64(); p < prob {
  	nextState = basis.First()
	}else{
		nextState = basis.Second()
	}

	qbt.state = nextState
	return nextState
}

func New(a, b complex128) Qubit {
  return NewWithState(state.New(a, b))
}

func NewWithState(state *state.State) Qubit {
	var mu sync.Mutex
	var qubit Qubit = &defaultQubit{mu, state}
	return qubit
}

func NewZero()   Qubit { return NewWithState(state.Zero()) }
func NewOne()    Qubit { return NewWithState(state.One()) }
func NewPlus()   Qubit { return NewWithState(state.Plus()) }
func NewMinus()  Qubit { return NewWithState(state.Minus()) }
func NewPlusI()  Qubit { return NewWithState(state.PlusI()) }
func NewMinusI() Qubit { return NewWithState(state.MinusI()) }