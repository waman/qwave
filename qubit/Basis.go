package qubit

import (
	"fmt"
	"log"
)

var(
	standard  *Basis = newBasis(zero, one)
	hadamard  *Basis = newBasis(plus, minus)
	imaginary *Basis = newBasis(plusI, minusI)
)

func BasisStandard()  *Basis { return standard }
func BasisHadamard()  *Basis { return hadamard }
func BasisImaginary() *Basis { return imaginary}

func newBasis(first, second *State) *Basis {
	return &Basis{first, second}
}

func NewBasis(first, second *State, delta float64) *Basis {
	if !first.IsOrthogonalTo(second, delta) {
		log.Panicf("Two states of a Basis must be orthogonal to each other: {%v, %v}", first, second)
	}
	return newBasis(first, second)
}

type Basis struct{
	first, second *State
}

func (basis *Basis) First() *State {
	return basis.first
}

func (basis *Basis) Second() *State {
	return basis.second
}

// Stats() returns a copy of state array.
func (basis *Basis) States() [2]*State {
	return [2]*State{ basis.first, basis.second}
}

func (basis *Basis) String() string {
	return fmt.Sprint("{%v, %v}", basis.first, basis.second)
}
