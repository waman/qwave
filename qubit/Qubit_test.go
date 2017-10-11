package qubit_test

import (
	"testing"
	"github.com/waman/qwave/qubit/basis"
	"github.com/waman/qwave/qubit"
)

func TestObserve(t *testing.T){
	x := qubit.NewZero()
	s0 := x.Observe(basis.Standard())
	s1 := x.Observe(basis.Standard())
	if !s0.EqualState(s1, 1e-10) {
		t.Errorf("Observe")
	}
}