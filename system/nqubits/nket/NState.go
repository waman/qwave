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
	v MetricVector
}

func (x *defaultNState) QubitCount() int {
	return x.qubitCount
}

func (x *defaultNState) Dim() int {
	return x.v.Dim()
}

func (x *defaultNState) Coefficients() []complex128 {
	return x.v.Coefficients()
}

func (x *defaultNState) CoefficientMap() map[int]complex128 {
	return x.v.CoefficientMap()
}

func (x *defaultNState) Get(i int) complex128 {
	return x.v.Get(i)
}

func (x *defaultNState) Prod(y MetricVector) complex128 {
	return x.v.Prod(y)
}

// |<x|y>|
func (x *defaultNState) Amplitude(w NState) float64 {
	return cmplx.Abs(x.Prod(w))
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
	return system.ToString(x.v.Coefficients()...)
}