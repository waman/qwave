package nop

import (
	"fmt"
	"log"
	"github.com/waman/qwave/system"
)

// Note that the returned object is not immutable (like big.Float etc.).
func NewDense(cs [][]complex128, doCopy bool) MutableMatrix {
	if doCopy {
		n := len(cs)
		_cs := make([][]complex128, n)
		for i := range cs {
			if len(cs[i]) != n { log.Panicf("???") }
			_cs[i] = make([]complex128, n)
			copy(_cs[i], cs[i])
		}
		return &denseMatrix{_cs}

	} else {
		return &denseMatrix{cs}
	}
}

type denseMatrix struct {
	cs [][]complex128
}

func (u *denseMatrix) Dim() int {
	return len(u.cs)
}

func (u *denseMatrix) At(i, j int) complex128 {
	return u.cs[i][j]
}

func (u *denseMatrix) Equals(v Matrix) bool {
	return equalMatrices(u, v)
}

func (u *denseMatrix) CloneMutable() MutableMatrix {
	return NewDense(u.cs, true)
}

func (u *denseMatrix) Apply(v []complex128) []complex128 {
	n := u.Dim()
	result := make([]complex128, n)

	for i := 0; i < n; i++ {
		x := 0i
		for j := 0; j < n; j++ {
			x += u.cs[i][j]*v[j]
		}
		result[i] = x
	}

	return result
}

func (u *denseMatrix) ApplyToMap(v map[int]complex128) map[int]complex128 {
	result := make(map[int]complex128)

	for i, n := 0, u.Dim(); i < n; i++ {
		x := 0i
		for j, c := range v {
			x += u.cs[i][j]*c
		}
		if x != 0 {	result[i] = x	}
	}

	return result
}

func (u *denseMatrix) Set(i, j int, c complex128){
	u.cs[i][j] = c
}

func (z *denseMatrix) Mul(x, y Matrix) Matrix {
	n := x.Dim()
	temp := make([]complex128, n)

	var copyToTemp func(i int)
	if xx, ok := x.(*denseMatrix); ok {
		copyToTemp = func(i int){
			copy(temp, xx.cs[i])
		}
	} else {
		copyToTemp = func(i int){
			for j := 0; j < n; j++ {
				temp[j] = x.At(i, j)
			}
		}
	}

	for i := 0; i < n; i++ {
		copyToTemp(i)
		for j := 0; j < n; j++ {
			s := 0i
			for k, temp_k := range temp {
				s += temp_k * y.At(k, j)
			}
			z.cs[i][j] = s
		}
	}
	return z
}

func (u *denseMatrix) Dense() Matrix {
	return u
}

func (u *denseMatrix) Sparse() Matrix {
	sparse := make(map[int]map[int]complex128)
	for i, c_i := range u.cs {
		sparse[i] = system.SliceToMap(c_i)
	}
	return NewSparse(u.Dim(), sparse, false)
}

func (u *denseMatrix) String() string {
	return fmt.Sprintf("%v", u.cs)
}