package basis

import (
	"fmt"
	"log"
	"github.com/waman/qwave/system/qubit/ket"
)

var(
	Standard  = newBasis(ket.Zero, ket.One)
	Hadamard  = newBasis(ket.Plus, ket.Minus)
	Imaginary = newBasis(ket.PlusI, ket.MinusI)
)

func newBasis(first, second *ket.State) *Basis {
	return &Basis{first, second}
}

func New(first, second *ket.State, delta float64) *Basis {
	if !first.IsOrthogonalTo(second, delta) {
		log.Panicf("Two kets of a Basis must be orthogonal to each other: {%v, %v}", first, second)
	}
	return newBasis(first, second)
}

type Basis struct{
	first, second *ket.State
}

func (basis *Basis) First() *ket.State {
	return basis.first
}

func (basis *Basis) Second() *ket.State {
	return basis.second
}

// States() returns a copy of ket array.
func (basis *Basis) States() [2]*ket.State {
	return [2]*ket.State{ basis.first, basis.second}
}

func (basis *Basis) String() string {
	return fmt.Sprintf("{%v, %v}", basis.first, basis.second)
}
