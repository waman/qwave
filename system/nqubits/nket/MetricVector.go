package nket

import (
	"log"
	"math/cmplx"
	"github.com/waman/qwave/system"
	"github.com/waman/qwave/system/nqubits/nop"
)

type MetricVector interface{
	Dim() int
	Coefficients() []complex128
	CoefficientMap() map[int]complex128
	At(i int) complex128

	InnerProduct(y MetricVector) complex128
	OuterProduct(y MetricVector) nop.Matrix
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

func (s *BaseMetricVector) At(i int) complex128 {
	if i == s.i {
		return 1
	} else if 0 <= i && i < s.n {
		return 0
	} else {
		log.Panicf("Index must be in [0, %d)", s.n)
		return cmplx.NaN()
	}
}

func (x *BaseMetricVector) InnerProduct(y MetricVector) complex128 {
	checkMatchingDimensions(x, y)
	return y.At(x.i)
}

func (s *BaseMetricVector) OuterProduct(y MetricVector) nop.Matrix {
	cMap := make(map[int]map[int]complex128)
	cMap[s.i] = y.CoefficientMap()
	return nop.NewSparse(s.n, cMap, false)
}

//***** DenseMetricVector (slice-base implementation) *****
type DenseMetricVector struct {
	cs []complex128
}

func (s *DenseMetricVector) Dim() int {
	return len(s.cs)
}

func (s *DenseMetricVector) Coefficients() []complex128 {
	return system.CreateCopy(s.cs)
}

func (s *DenseMetricVector) CoefficientMap() map[int]complex128 {
	return system.SliceToMap(s.cs)
}

func (s *DenseMetricVector) At(i int) complex128 {
	return s.cs[i]
}

// <x|y>
func (x *DenseMetricVector) InnerProduct(y MetricVector) complex128 {
	checkMatchingDimensions(x, y)
	return InnerProduct(x.cs, y)
}

func (s *DenseMetricVector) OuterProduct(y MetricVector) nop.Matrix {
	return OuterProduct(s.cs, y)
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
	return system.CreateCopyMap(s.cMap)
}

func (s *SparseMetricVector) At(i int) complex128 {
	if 0 <= i && i < s.n {
		return s.cMap[i]
	} else {
		log.Panicf("Index must be in [0, %d)", s.n)
		return cmplx.NaN()
	}
}

func (x *SparseMetricVector) InnerProduct(y MetricVector) complex128 {
	checkMatchingDimensions(x, y)
	return InnerProductMap(x.cMap, y)
}

func (x *SparseMetricVector) OuterProduct(y MetricVector) nop.Matrix {
	checkMatchingDimensions(x, y)
	return OuterProductMap(x.Dim(), x.cMap, y)
}