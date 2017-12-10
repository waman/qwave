package nqubits

import (
	"github.com/waman/qwave/system"
	"github.com/waman/qwave/system/nqubits/nbasis"
	"github.com/waman/qwave/system/nqubits/nket"
	"github.com/waman/qwave/system/nqubits/nop"
	"github.com/waman/qwave/system/qubit"
	"github.com/waman/qwave/system/qubit/basis"
	"github.com/waman/qwave/system/qubit/ket"
	"github.com/waman/qwave/system/qubit/op"
	"log"
	"math"
	"math/cmplx"
	"math/rand"
	"sync"
)

func NewDense(qbtCount int, cs []complex128, isNormalized bool) MultiQubitSystem {
	if !isNormalized { system.Normalize(cs)	}

	var _cs []complex128
	if dim := system.Dim(qbtCount); len(cs) != dim {
		_cs = make([]complex128, dim)
		copy(_cs, cs)
	} else {
		_cs = system.CreateCopy(cs)
	}

	var mu sync.Mutex
  return &denseMultiQubitSystem{qbtCount, &mu, _cs}
}

type denseMultiQubitSystem struct {
	qubitCount int
	mu         *sync.Mutex
	cs         []complex128
}

func (nqbts *denseMultiQubitSystem) QubitCount() int {
	return nqbts.qubitCount
}

func (nqbts *denseMultiQubitSystem) Dim() int {
	return len(nqbts.cs)
}

func (nqbts *denseMultiQubitSystem) nstate() nket.NState {
	return nket.NewWith(nqbts.qubitCount, &mutableDenseMetricVector{nqbts.cs})
}

func (nqbts *denseMultiQubitSystem) Observe(b nbasis.NBasis) nket.NState {
	if b.QubitCount() != nqbts.qubitCount {
		log.Panicf("The number of qubits is not match: MultiQubitSystem %d, NBasis %d",
			nqbts.qubitCount, b.QubitCount())
	}

	nqbts.mu.Lock()
	defer nqbts.mu.Unlock()

	r := rand.Float64()
	p := 0.0
	current := nqbts.nstate()

	for i, n := 0, nqbts.Dim()-1; i < n; i++ {
		s := b.At(i)
		p += s.Probability(current)
		if r < p {
			nqbts.cs = s.Coefficients()
			return s
		}
	}

	s := b.At(nqbts.Dim()-1)
	nqbts.cs = s.Coefficients()
	return s
}

func (nqbts *denseMultiQubitSystem) ObserveInStandardBasis() nket.NState {
	return nqbts.Observe(nbasis.Standard(nqbts.qubitCount))
}

func (nqbts *denseMultiQubitSystem) iterateQubit(i int, f func(j, k int)) {
	mask := 1 << uint(nqbts.qubitCount-i-1)
	for j, n := 0, nqbts.Dim(); j < n; j++ {
		if j&mask == 0 {
			f(j, j+mask)
		}
	}
}

func (nqbts *denseMultiQubitSystem) ObserveQubit(i int, b *basis.Basis) *ket.State {
	nqbts.mu.Lock()
	defer nqbts.mu.Unlock()

	cs := nqbts.cs

	s := b.First()
	m := s.OuterProduct(s)

  p := 0.0
	nqbts.iterateQubit(i, func(j, k int){
		x, y := m.Apply(cs[j], cs[k])
		xAbs, yAbs := cmplx.Abs(x), cmplx.Abs(y)
    p += xAbs*xAbs + yAbs*yAbs
	})

	var nf complex128  // normaization factor
	if r := rand.Float64(); r < p {
    nf = complex(1.0/math.Sqrt(p), 0)
	} else {
		s = b.Second()
		m = s.OuterProduct(s)
		nf = complex(1.0/math.Sqrt(1-p), 0)
	}

	nqbts.iterateQubit(i, func(j, k int){
		x, y := m.Apply(cs[j], cs[k])
		cs[j], cs[k] = x*nf, y*nf
	})
	return s
}

func (nqbts *denseMultiQubitSystem) ObserveQubitInStandardBasis(i int) *ket.State {
	nqbts.mu.Lock()
	defer nqbts.mu.Unlock()

	cs := nqbts.cs
	p := 0.0
	nqbts.iterateQubit(i, func(j, k int) {
		amp := cmplx.Abs(cs[j])
		p += amp * amp
	})

	if r := rand.Float64(); r < p {
		nf := complex(1.0/math.Sqrt(p), 0)
		nqbts.iterateQubit(i, func(j, k int) {
			cs[j] = cs[j] * nf
			cs[k] = 0
		})
		return ket.Zero()

	} else {
		nf := complex(1.0/math.Sqrt(1-p), 0)
		nqbts.iterateQubit(i, func(j, k int) {
			cs[j] = 0
			cs[k] = cs[k] * nf
		})
		return ket.One()
	}
}

func (nqbts *denseMultiQubitSystem) ObserveQubitInHadamardBasis(i int) *ket.State {
	nqbts.mu.Lock()
	defer nqbts.mu.Unlock()

	cs := nqbts.cs
	p := 0.0
	nqbts.iterateQubit(i, func(j, k int) {
		amp := cmplx.Abs(cs[j] + cs[k])
		p += amp * amp / 2.0
	})

	if r := rand.Float64(); r < p {
		nf := complex(1.0/(math.Sqrt(p)*2.0), 0)
		nqbts.iterateQubit(i, func(j, k int) {
			x := (cs[j] + cs[k]) * nf
			cs[j] = x
			cs[k] = x

		})
		return ket.Plus()

	} else {
		nf := complex(1.0/(math.Sqrt(1-p)*2.0), 0)
		nqbts.iterateQubit(i, func(j, k int) {
			x := (cs[j] - cs[k]) * nf
			cs[j] = x
			cs[k] = -x
		})
		return ket.Minus()
	}
}

func (nqbts *denseMultiQubitSystem) Qubit(i int) qubit.Qubit {
	return &qubitImpl{nqbts, i}
}

func (nqbts *denseMultiQubitSystem) Apply(u nop.Matrix) {
	nqbts.mu.Lock()
	defer nqbts.mu.Unlock()

	nqbts.cs = u.Apply(nqbts.cs)
}

func (nqbts *denseMultiQubitSystem) ApplyToQubit(i int, u op.Matrix2x2) {
	nqbts.mu.Lock()
	defer nqbts.mu.Unlock()

	cs := nqbts.cs
	nqbts.iterateQubit(i, func(j, k int) {
		cs[j], cs[k] = u.Apply(cs[j], cs[k])
	})
}

//***** mutableDenseMetricVector *****
type mutableDenseMetricVector struct {
	cs []complex128
}

func (m *mutableDenseMetricVector) Dim() int {
	return len(m.cs)
}

// Note: This method DO NOT copy the slice
func (m *mutableDenseMetricVector) Coefficients() []complex128 {
	return m.cs
}

func (m *mutableDenseMetricVector) CoefficientMap() map[int]complex128 {
	return system.SliceToMap(m.cs)
}

func (m *mutableDenseMetricVector) At(i int) complex128 {
	return m.cs[i]
}

func (m *mutableDenseMetricVector) InnerProduct(y nket.MetricVector) complex128 {
	return nket.InnerProduct(m.cs, y)
}

func (m *mutableDenseMetricVector) OuterProduct(y nket.MetricVector) nop.Matrix {
	return nket.OuterProduct(m.cs, y)
}
