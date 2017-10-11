// basis package contains utility factory functions of qubit.Basis:
//
//   import github.com/waman/qwave/qubit/basis
//
//   b := basis.Standard()
package basis

import "github.com/waman/qwave/qubit"

func New(first, second *qubit.State, delta float64) *qubit.Basis {
	return qubit.NewBasis(first, second, delta)
}

func Standard()  *qubit.Basis { return qubit.BasisStandard() }
func Hadamard()  *qubit.Basis { return qubit.BasisHadamard() }
func Imaginary() *qubit.Basis { return qubit.BasisImaginary() }