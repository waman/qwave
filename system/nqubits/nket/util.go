package nket

import (
	"math/cmplx"
	"github.com/waman/qwave/system"
	"github.com/waman/qwave/system/nqubits/nop"
	"log"
)

func checkMatchingDimensions(x, y MetricVector){
	if x.Dim() != y.Dim() {
		log.Panicf("Two MetricVectors have different dimensions: %d, %d", x.Dim(), y.Dim())
	}
}

func InnerProduct(x []complex128, y system.ComplexContainer) complex128 {
	result := 0i
	for i, c := range x {
		result += cmplx.Conj(c)*y.At(i)
	}
	return result
}

// This method DO NOT check matching dimensions of x and y.
func InnerProductMap(x map[int]complex128, y system.ComplexContainer) complex128 {
	result := 0i
	for i, c := range x {
		result += cmplx.Conj(c)*y.At(i)
	}
	return result
}

func OuterProduct(x []complex128, y system.ComplexContainer) nop.Matrix {
	n := len(x)
	cs := make([][]complex128, n)
	for i := 0; i < n; i++ {
		cs[i] = make([]complex128, n)
		for j := 0; j < n; j++ {
			cs[i][j] = cmplx.Conj(x[i]) * y.At(j)
		}
	}
	return nop.NewDense(cs, false)
}

// This method DO NOT check matching dimensions of x and y.
func OuterProductMap(n int, x map[int]complex128, y system.ComplexContainer) nop.Matrix {
	cMap := make(map[int]map[int]complex128)
	for i, c_i := range x {
		_c_i := make(map[int]complex128)
		cc_i := cmplx.Conj(c_i)
		for j := 0; j < n; j++ {
			_c_i[j] = cc_i * y.At(j)
		}
		if len(_c_i) != 0 { cMap[i] = _c_i }
	}
	return nop.NewSparse(n, cMap, false)
}