package basis

import (
	"fmt"
	"log"
	"github.com/waman/qwave/system/qubit/ket"
)

var(
	standard  *Basis = newBasis(ket.Zero(), ket.One())
	hadamard  *Basis = newBasis(ket.Plus(), ket.Minus())
	imaginary *Basis = newBasis(ket.PlusI(), ket.MinusI())
)

func Standard()  *Basis { return standard }
func Hadamard()  *Basis { return hadamard }
func Imaginary() *Basis { return imaginary}

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
