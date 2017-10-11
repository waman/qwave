// state package contains utility factory functions of qubit.State:
//
//   import github.com/waman/qwave/qubit/state
//
//   s := state.Zero()
package state

import "github.com/waman/qwave/qubit"

func New(a, b complex128) *qubit.State {
	return qubit.NewState(a, b)
}

func Zero()   *qubit.State { return qubit.StateZero() }
func One()    *qubit.State { return qubit.StateOne() }
func Plus()   *qubit.State { return qubit.StatePlus() }
func Minus()  *qubit.State { return qubit.StateMinus() }
func PlusI()  *qubit.State { return qubit.StatePlusI() }
func MinusI() *qubit.State { return qubit.StateMinusI() }
