package nket

import (
	"math/cmplx"
)

func InnerProduct(x []complex128, y MetricVector) complex128 {
	result := 0i
	for i, c := range x {
		result += cmplx.Conj(c)*y.Get(i)
	}
	return result
}
