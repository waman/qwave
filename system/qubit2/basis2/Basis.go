package basis2

import (
	"github.com/waman/qwave/system/qubit2/ket2"
	"github.com/waman/qwave/system/nqubits/nbasis"
)

var(
	standard = nbasis.New(2,
		ket2.Zero(), ket2.One(),
		ket2.Two(), ket2.Three())

	bell = nbasis.New(2,
		ket2.PhiPlus(), ket2.PhiMinus(),
		ket2.PsiPlus(), ket2.PsiMinus())
)

func Standard() nbasis.NBasis { return standard }
func Bell() nbasis.NBasis { return bell }