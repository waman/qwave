package nop

import (
	"log"
	"fmt"
)

func Id(n int) Matrix {
	return &identity{n}
}

type identity struct {
	n int  // dimension
}

func (id *identity) Dim() int {
	return id.n
}

func (id *identity) At(i, j int) complex128 {
	if i < 0 || id.n <= i || j < 0 || id.n <= j {
		log.Panicf("The index must be in [0, %d): (%d, %d) appears", id.n, i, j)
	}
	if i == j {
		return 1
	} else {
		return 0
	}
}

func (id *identity) Equals(v Matrix) bool {
	if id.n != v.Dim() {
		return false
	}

	for i := 0; i < id.n; i++ {
		for j := 0; j < id.n; j ++ {
			if i == j {
				if v.At(i, j) != 1 { return false }
			} else {
				if v.At(i, j) != 0 { return false }
			}
		}
	}
	return true
}

func (id *identity) Dense() Matrix {
	return id.CloneMutable()
}

func (id *identity) Sparse() Matrix {
	n := id.n
	cMap := make(map[int]map[int]complex128)
	for i := 0; i < n; i++ {
		m := make(map[int]complex128)
		m[i] = 1
		cMap[i] = m
	}
	return NewSparse(n, cMap, false)
}

func (id *identity) CloneMutable() MutableMatrix {
	cs := make([][]complex128, id.n)
	for i := range cs {
		cs[i] = make([]complex128, id.n)
		cs[i][i] = 1
	}
	return NewDense(cs, false)
}

func (id *identity) Apply(v []complex128) []complex128 {
	return v
}

func (id *identity) ApplyToMap(v map[int]complex128) map[int]complex128 {
	return v
}

func (id *identity) String() string {
	return fmt.Sprintf("I_%d", id.n)
}
