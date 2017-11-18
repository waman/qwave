package nqubits

import (
	"github.com/waman/qwave/system/nqubits/nbasis"
	"github.com/waman/qwave/system/nqubits/nket"
	"github.com/waman/qwave/system/qubit/basis"
	"github.com/waman/qwave/system/qubit/ket"
	"github.com/waman/qwave/system/qubit"
	"math/rand"
	"math"
	"math/cmplx"
	"github.com/waman/qwave/system"
	"sync"
)

func NewDense(qbtCount int, cs []complex128, isNormalized bool) NQubits {
	if !isNormalized { system.Normalize(cs) }

	var cs_ []complex128
	if dim := system.Dim(qbtCount); len(cs) != dim {
		cs_ = make([]complex128, dim)
		copy(cs_, cs)
	} else {
		cs_ = system.CreateCopy(cs)
	}

	var mu sync.Mutex
	return &denseNQubits{qbtCount, mu, cs_}
}

type denseNQubits struct {
	qubitCount int
	mu sync.Mutex
	cs []complex128
}

func (nqbts *denseNQubits) QubitCount() int {
	return nqbts.qubitCount
}

func (nqbts *denseNQubits) Dim() int {
	return len(nqbts.cs)
}

func (nqbts *denseNQubits) nstate() nket.NState {
	return nket.NewWith(nqbts.qubitCount, &mutableDenseMetricVector{nqbts.cs})
}

func (nqbts *denseNQubits) Observe(b nbasis.NBasis) nket.NState {
	nqbts.mu.Lock()
	defer nqbts.mu.Unlock()

	r := rand.Float64()
	p := 0.0
	current := nqbts.nstate()

	for i, n := 0, nqbts.Dim()-1; i < n; i++ {
		s := b.Get(i)
		p += s.Probability(current)
		if r < p {
			nqbts.cs = s.Coefficients()
			return s
		}
	}

	s := b.Get(nqbts.Dim()-1)
	nqbts.cs = s.Coefficients()
	return s
}

func (nqbts *denseNQubits) ObserveInStandardBasis() nket.NState {
	return nqbts.Observe(nbasis.Standard(nqbts.qubitCount))
}

func (nqbts *denseNQubits) ObserveQubit(i int, b *basis.Basis) *ket.State {
	nqbts.mu.Lock()
	defer nqbts.mu.Unlock()

	panic("implement me")
}

func (nqbts *denseNQubits) ObserveQubitInStandardBasis(i int) *ket.State {
	nqbts.mu.Lock()
	defer nqbts.mu.Unlock()

	mask := 1 << uint(nqbts.qubitCount-i-1)

	p := 0.0
	for j, c := range nqbts.cs {
		if j & mask == 0 {
			amp := cmplx.Abs(c)
			p += amp*amp
		}
	}

	if r := rand.Float64(); r < p {
		nf := complex(1.0/math.Sqrt(p), 0)
		for j, c := range nqbts.cs {
			if j & mask == 0 {
				nqbts.cs[j] = c*nf
			} else {
				nqbts.cs[j] = 0
			}
		}
		return ket.Zero()

	} else {
		nf := complex(1.0/math.Sqrt(1-p), 0)
		for j, c := range nqbts.cs {
			if j & mask == 0 {
				nqbts.cs[j] = 0
			} else {
				nqbts.cs[j] = c*nf
			}
		}
		return ket.One()
	}
}

func (nqbts *denseNQubits) ObserveQubitInHadamardBasis(i int) *ket.State {
	nqbts.mu.Lock()
	defer nqbts.mu.Unlock()

	mask := 1 << uint(nqbts.qubitCount-i-1)
	p := 0.0
	for j, c := range nqbts.cs {
		if j & mask == 0 {
			amp := cmplx.Abs(c + nqbts.cs[j+mask])
			p += amp*amp/2.0
		}
	}

	if r := rand.Float64(); r < p {
		nf := complex(1.0/(math.Sqrt(p)*2.0), 0)
		for j, c := range nqbts.cs {
			if j & mask == 0 {
				x := (c + nqbts.cs[j+mask])*nf
				nqbts.cs[j] = x
				nqbts.cs[j+mask] = x
			}
		}
		return ket.Plus()

	} else {
		nf := complex(1.0/(math.Sqrt(1-p)*2.0), 0)
		for j, c := range nqbts.cs {
			if j & mask == 0 {
				x := (c - nqbts.cs[j+mask])*nf
				nqbts.cs[j] = x
				nqbts.cs[j+mask] = -x
			}
		}
		return ket.Minus()
	}
}

func (nqbts *denseNQubits) Qubit(i int) qubit.Qubit {
	return &qubitImpl{nqbts, i}
}

func (nqbts *denseNQubits) ApplyToQubit(i int, u qubit.Matrix2) {
	panic("implement me")
}

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

func (m *mutableDenseMetricVector) Get(i int) complex128 {
	return m.cs[i]
}

func (m *mutableDenseMetricVector) Prod(y nket.MetricVector) complex128 {
	return nket.InnerProduct(m.cs, y)
}
