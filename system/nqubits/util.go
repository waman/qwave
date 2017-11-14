package nqubits

import "math/cmplx"

func Dim(qubitCount int) int {
	return 1 << uint(qubitCount)
}

func Norm2(cs []complex128) float64 {
	norm2 := 0.0
	for _, c := range cs {
		abs2 := cmplx.Abs(c)
		norm2 += abs2*abs2
	}
	return norm2
}

func Norm2Map(cMap map[int]complex128) float64 {
	norm2 := 0.0
	for _, c := range cMap {
		abs2 := cmplx.Abs(c)
		norm2 += abs2*abs2
	}
	return norm2
}