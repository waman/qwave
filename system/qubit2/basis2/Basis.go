package basis2

import (
	"github.com/waman/qwave/system/qubit2/ket2"
)

var(
	standard = &Basis{&[4]*ket2.State{
		ket2.Zero(),
		ket2.One(),
		ket2.Two(),
		ket2.Three()}}

	bell = &Basis{&[4]*ket2.State{
		ket2.PhiPlus(),
		ket2.PhiMinus(),
		ket2.PsiPlus(),
		ket2.PsiMinus()}}
)

func Standard() *Basis { return standard }
func Bell() *Basis { return bell }

type Basis struct {
	states *[4]*ket2.State
}

func (basis *Basis) Get(i int) *ket2.State {
	return basis.states[i]
}

func (basis *Basis) States() [4]*ket2.State {
	return [4]*ket2.State{basis.states[0], basis.states[1], basis.states[2], basis.states[3]}
}