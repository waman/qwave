package nket

import (
	"github.com/waman/qwave/system"
	"math"
	"math/cmplx"
	. "github.com/waman/qwave/system/nqubits"
)

type NState struct {
	MetricVector
}

func New(qbtCount int, isNormalized bool, cs ...complex128) *NState{
	return NewDense(qbtCount, cs, isNormalized)
}

func NewBase(qbtCount int, i int) *NState {
  return &NState{&BaseMetricVector{Dim(qbtCount), i}}
}

func NewDense(qbtCount int, cs []complex128, isNormalized bool) *NState {
	cs_ := make([]complex128, Dim(qbtCount))

	if isNormalized {
		copy(cs_, cs)
	} else {
		nf := complex(1.0/math.Sqrt(Norm2(cs)), 0)
		for i, n := 0, len(cs); i < n; i++ {
			cs_[i] = cs[i]*nf
		}
	}

	return &NState{&DenseMetricVector{cs: cs_}}
}

func NewSparse(qbtCount int, cMap map[int]complex128, isNormalized bool) *NState {
  cMap_ := make(map[int]complex128)

  if isNormalized {
  	for i, c := range cMap {
			cMap[i] = c
		}

	} else {
		nf := complex(1.0/math.Sqrt(Norm2Map(cMap)), 0)
		for i, c := range cMap{
			cMap_[i] = c*nf
		}
	}

	return &NState{&SparseMetricVector{Dim(qbtCount), cMap_}}
}

// |<x|y>|
func (x NState) Amplitude(y NState) float64 {
	return cmplx.Abs(x.Prod(y))
}

// |<x|y>|^2
func (x NState) Probability(y NState) float64 {
	amp := x.Amplitude(y)
	return amp*amp
}

// |<x|y>| == 1 (|<x|y>| >= 1-delta)
func (x NState) EqualState(y NState, delta float64) bool {
	if x == y { return true }
	return x.Amplitude(y) >= 1-delta
}

func (x NState) IsOrthogonalTo(y NState, delta float64) bool {
	return x.Amplitude(y) <= delta
}

func (s NState) String() string {
	return system.ToString(s.Coefficients()...)
}