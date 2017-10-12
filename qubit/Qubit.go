package qubit

import (
	"math/rand"
	"github.com/waman/qwave/qubit/state"
	"github.com/waman/qwave/qubit/basis"
)

type Qubit interface{
	Observe(basis *basis.Basis) *state.State
}

type defaultQubit struct {
	state *state.State
}

func (qubit *defaultQubit) Observe(basis *basis.Basis) *state.State {
  prob := qubit.state.Probability(basis.First())

  var nextState *state.State
  if p := rand.Float64(); p < prob {
  	nextState = basis.First()
	}else{
		nextState = basis.Second()
	}

	qubit.state = nextState
	return nextState
}

func New(a, b complex128) Qubit {
  return NewWithState(state.New(a, b))
}

func NewWithState(state *state.State) Qubit {
	var qubit Qubit = &defaultQubit{state}
	return qubit
}

func NewZero()   Qubit { return NewWithState(state.Zero()) }
func NewOne()    Qubit { return NewWithState(state.One()) }
func NewPlus()   Qubit { return NewWithState(state.Plus()) }
func NewMinus()  Qubit { return NewWithState(state.Minus()) }
func NewPlusI()  Qubit { return NewWithState(state.PlusI()) }
func NewMinusI() Qubit { return NewWithState(state.MinusI()) }