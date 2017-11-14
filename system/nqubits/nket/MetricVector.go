package nket

import (
	"log"
	"math/cmplx"
)

type MetricVector interface{
	Dim() int
	Coefficients() []complex128
	CoefficientMap() map[int]complex128
	Get(i int) complex128

	Prod(y MetricVector) complex128
}

//***** BaseMetricVector (uni-component MetricVector) *****
type BaseMetricVector struct {
	n int // dimension of state vector (2^qubitCount)
	i int // index of non-zero component
}

func (s *BaseMetricVector) Dim() int {
	return s.n
}

func (s *BaseMetricVector) Coefficients() []complex128 {
	cs := make([]complex128, s.n)
	cs[s.i] = 1
	return cs
}

func (s *BaseMetricVector) CoefficientMap() map[int]complex128 {
	cMap := make(map[int]complex128)
	cMap[s.i] = 1
	return cMap
}

func (s *BaseMetricVector) Get(i int) complex128 {
	if i == s.i {
		return 1
	} else if 0 <= i && i < s.n {
		return 0
	} else {
		log.Panicf("Index must be in [0, %d)", s.n)
		return cmplx.NaN()
	}
}

func (s *BaseMetricVector) Prod(y MetricVector) complex128 {
	return y.Get(s.i)
}

//***** DenseMetricVector (slice-base implementation) *****
type DenseMetricVector struct {
	cs []complex128
}

func (s *DenseMetricVector) Dim() int {
	return len(s.cs)
}

func (s *DenseMetricVector) Coefficients() []complex128 {
	cs := make([]complex128, len(s.cs))
	copy(cs, s.cs)
	return cs
}

func (s *DenseMetricVector) CoefficientMap() map[int]complex128 {
	cMap := make(map[int]complex128)
	for i, c := range s.cs {
		if c != 0 {
			cMap[i] = c
		}
	}
	return cMap
}

func (s *DenseMetricVector) Get(i int) complex128 {
	return s.cs[i]
}

// <x|y>
func (s *DenseMetricVector) Prod(y MetricVector) complex128 {
	result := 0i
	for i, c := range s.cs {
		result += cmplx.Conj(c)*y.Get(i)
	}
	return result
}

//***** SparseMetricVector (uni-component MetricVector) *****
type SparseMetricVector struct {
	n int  // dimension of vector (2^qubitCount)
	cMap map[int]complex128
}

func (s *SparseMetricVector) Dim() int {
	return s.n
}

func (s *SparseMetricVector) Coefficients() []complex128 {
	cs := make([]complex128, s.n)
	for i, c := range s.cMap {
		cs[i] = c
	}
	return cs
}

func (s *SparseMetricVector) CoefficientMap() map[int]complex128 {
	cMap := make(map[int]complex128)
	for i, c := range s.cMap {
		cMap[i] = c
	}
	return cMap
}

func (s *SparseMetricVector) Get(i int) complex128 {
	if 0 <= i && i < s.n {
		return s.cMap[i]
	} else {
		log.Panicf("Index must be in [0, %d)", s.n)
		return cmplx.NaN()
	}
}

func (s *SparseMetricVector) Prod(y MetricVector) complex128 {
	result := 0i
	for i, c := range s.cMap {
		result += cmplx.Conj(c)*y.Get(i)
	}
	return result
}