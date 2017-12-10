package nket

import (
	"github.com/waman/qwave/system"
	"math/cmplx"
)

type NState interface {
	MetricVector

	QubitCount() int
  Amplitude(w NState) float64
  Probability(y NState) float64
  EqualState(y NState, delta float64) bool
  IsOrthogonalTo(y NState, delta float64) bool

  String() string
}

func New(qbtCount int, isNormalized bool, cs ...complex128) NState{
	return NewDense(qbtCount, cs, isNormalized)
}

func NewBase(qbtCount int, i int) NState {
  return NewWith(qbtCount, &BaseMetricVector{system.Dim(qbtCount), i})
}

func NewDense(qbtCount int, cs []complex128, isNormalized bool) NState {
	if !isNormalized { system.Normalize(cs)}
	return NewWith(qbtCount, &DenseMetricVector{system.CreateCopy(cs)})
}

func NewSparse(qbtCount int, cMap map[int]complex128, isNormalized bool) NState {
	if !isNormalized { system.NormalizeMap(cMap) }
	v := &SparseMetricVector{ system.Dim(qbtCount), system.CreateCopyMap(cMap)}
	return NewWith(qbtCount, v)
}

func NewWith(qbtCount int, v MetricVector) NState {
	return &defaultNState{qbtCount, v}
}

type defaultNState struct {
	qubitCount int
	MetricVector
}

func (x *defaultNState) QubitCount() int {
	return x.qubitCount
}

// |<x|y>|
func (x *defaultNState) Amplitude(w NState) float64 {
	return cmplx.Abs(x.InnerProduct(w))
}

// |<x|y>|^2
func (x *defaultNState) Probability(y NState) float64 {
	amp := x.Amplitude(y)
	return amp*amp
}

// |<x|y>| == 1 (|<x|y>| >= 1-delta)
func (x *defaultNState) EqualState(y NState, delta float64) bool {
	if x == y { return true }
	return x.Amplitude(y) >= 1-delta
}

func (x *defaultNState) IsOrthogonalTo(y NState, delta float64) bool {
	return x.Amplitude(y) <= delta
}

func (x *defaultNState) String() string {
	return system.ToString(x.Coefficients()...)
}