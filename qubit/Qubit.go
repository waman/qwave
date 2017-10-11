package qubit

import (
	"math/rand"
)

type Qubit interface{
	Observe(basis *Basis) *State
}

type defaultQubit struct {
	state *State
}

func (qubit *defaultQubit) Observe(basis *Basis) *State {
  prob := qubit.state.Probability(basis.First())

  var nextState *State
  if p := rand.Float64(); p < prob {
  	nextState = basis.First()
	}else{
		nextState = basis.Second()
	}

	qubit.state = nextState
	return nextState
}

func New(a, b complex128) Qubit {
  return NewWithState(NewState(a, b))
}

func NewWithState(state *State) Qubit {
	var qubit Qubit = &defaultQubit{state}
	return qubit
}

func NewZero()   Qubit { return NewWithState(StateZero()) }
func NewOne()    Qubit { return NewWithState(StateOne()) }
func NewPlus()   Qubit { return NewWithState(StatePlus()) }
func NewMinus()  Qubit { return NewWithState(StateMinus()) }
func NewPlusI()  Qubit { return NewWithState(StatePlusI()) }
func NewMinusI() Qubit { return NewWithState(StateMinusI()) }