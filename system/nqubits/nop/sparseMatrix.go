package nop

import (
	"log"
	"github.com/waman/qwave/system"
)

// Note that the returned object is not immutable (like big.Float etc.).
func NewSparse(n int, cMap map[int]map[int]complex128, doCopy bool) MutableMatrix {
	if cMap == nil {
		return &sparseMatrix{n, make(map[int]map[int]complex128)}
	}

	if doCopy {
		_cMap := make(map[int]map[int]complex128)
		for i, c_i := range cMap {
			_c_i := make(map[int]complex128)
			for j, c_ij := range c_i {
				_c_i[j] = c_ij
			}
			_cMap[i] = _c_i
		}
		return &sparseMatrix{n, _cMap}

	} else {
		return &sparseMatrix{n, cMap}
	}
}

type sparseMatrix struct {
	n int
	cMap map[int]map[int]complex128
}

func (u *sparseMatrix) Dim() int {
	return u.n
}

func (u *sparseMatrix) At(i, j int) complex128 {
	if i < 0 || u.n <= i { log.Panicf("???")}
	if j < 0 || u.n <= j { log.Panicf("???") }
	m := u.cMap[i]
	if m == nil {
		return 0
	} else {
		return m[j]
	}
}

func (u *sparseMatrix) Equals(v Matrix) bool {
	return equalMatrices(u, v)
}

func (u *sparseMatrix) CloneMutable() MutableMatrix {
	return NewSparse(u.n, u.cMap, true)
}

func (u *sparseMatrix) Apply(v []complex128) []complex128 {
	n := u.Dim()
	result := make([]complex128, n)

	for i := 0; i < n; i++ {
		if u_i := u.cMap[i]; u_i != nil {
			x := 0i
			for j, u_ij := range u_i {
				x += u_ij *v[j]
			}
			result[i] = x
		}
	}

	return result
}

func (u *sparseMatrix) ApplyToMap(v map[int]complex128) map[int]complex128 {
	n := u.Dim()
	cs := make(map[int]complex128)

	for i := 0; i < n; i++ {
		if c_i := u.cMap[i]; c_i != nil {
			x := 0i
			for j, c_ij := range c_i {
				x += c_ij *v[j]
			}
			cs[i] = x
		}
	}

	return cs
}

func (u *sparseMatrix) Set(i, j int, c complex128){
	c_i := u.cMap[i]

	if c != 0 {
		// put entry
		if c_i := u.cMap[i]; c_i != nil {
			c_i[j] = c

		} else {
			c_i = make(map[int]complex128)
			c_i[j] = c
			u.cMap[i] = c_i

		}
	} else {
		// delete entry
		if c_i != nil {
			delete(c_i, j)
			if len(c_i) == 0 {
				delete(u.cMap, i)
			}
		}
	}
}

func (z *sparseMatrix) Mul(x, y Matrix) Matrix {
	n := x.Dim()
	temp := make(map[int]complex128)


	for i := 0; i < n; i++ {
		// copy i-th row to the temp
		for j := 0; j < n; j++ {
			if c_ij := x.At(i, j); c_ij != 0 {
				temp[j] = c_ij
			} else {
				delete(temp, j)
			}
		}

		for j := 0; j < n; j++ {
			s := 0i
			for k, temp_k := range temp {
				s += temp_k * y.At(k, j)
			}
			z.Set(i, j, 0)
		}
	}
	return z
}

func (u *sparseMatrix) Dense() Matrix {
	n := u.Dim()
	dense := make([][]complex128, n)
	for i, m := range u.cMap {
		dense[i] = system.MapToSlice(n, m)
	}
	return NewDense(dense, false)
}

func (u *sparseMatrix) Sparse() Matrix {
	return u
}

func (u *sparseMatrix) String() string {
	return u.Dense().String()
}