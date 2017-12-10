package basis2

import (
	"github.com/waman/qwave/system/qubits2/ket2"
	"github.com/waman/qwave/system/nqubits/nbasis"
	"github.com/waman/qwave/system/nqubits/nket"
)

var(
	standard = New(
		ket2.Zero(), ket2.One(),
		ket2.Two(), ket2.Three())

	bell = New(
		ket2.PhiPlus(), ket2.PhiMinus(),
		ket2.PsiPlus(), ket2.PsiMinus())
)

func New(a, b, c, d nket.NState) nbasis.NBasis {
	return nbasis.New(2, a, b, c, d)
}

func Standard() nbasis.NBasis { return standard }
func Bell() nbasis.NBasis { return bell }