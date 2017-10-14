package basis

import (
	"fmt"
	"log"
	"github.com/waman/qwave/qubit/state"
)

var(
	standard  *Basis = newBasis(state.Zero(), state.One())
	hadamard  *Basis = newBasis(state.Plus(), state.Minus())
	imaginary *Basis = newBasis(state.PlusI(), state.MinusI())
)

func Standard()  *Basis { return standard }
func Hadamard()  *Basis { return hadamard }
func Imaginary() *Basis { return imaginary}

func newBasis(first, second *state.State) *Basis {
	return &Basis{first, second}
}

func New(first, second *state.State, delta float64) *Basis {
	if !first.IsOrthogonalTo(second, delta) {
		log.Panicf("Two states of a Basis must be orthogonal to each other: {%v, %v}", first, second)
	}
	return newBasis(first, second)
}

type Basis struct{
	first, second *state.State
}

func (basis *Basis) First() *state.State {
	return basis.first
}

func (basis *Basis) Second() *state.State {
	return basis.second
}

// States() returns a copy of state array.
func (basis *Basis) States() [2]*state.State {
	return [2]*state.State{ basis.first, basis.second}
}

func (basis *Basis) String() string {
	return fmt.Sprintf("{%v, %v}", basis.first, basis.second)
}
