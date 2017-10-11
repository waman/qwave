package qubit_test

import (
	"testing"
	"github.com/waman/qwave/qubit/state"
	"github.com/waman/qwave/qubit"
)

const delta float64 = 1e-10

func TestEqualState(t *testing.T){
	var tests = []struct {
		first, second *qubit.State
		want  bool
	}{
		{state.Zero()  , state.New(1, 0), true},
		{state.One()   , state.New(0, 1), true},
		{state.Plus()  , state.New(1, 1), true},
		{state.Minus() , state.New(1, -1), true},
		{state.PlusI() , state.New(1, 1i), true},
		{state.MinusI(), state.New(1, -1i), true},
		{state.Zero(), state.New(0, 1), false},
	}
	for _, test := range tests {
		if got := test.first.EqualState(test.second, delta); got != test.want {
			t.Errorf("(%v).Equalstate.(%v) = %v", test.first, test.second, got)
		}
	}
}

func TestIsOrthogonalTo(t *testing.T){
	var tests = []struct {
		first, second *qubit.State
		want  bool
	}{
		{state.Zero(), state.One(), true},
		{state.Plus(), state.Minus(), true},
		{state.PlusI(), state.MinusI(), true},
		{state.Zero(), state.Plus(), false},
		{state.Zero(), state.MinusI(), false},
	}
	for _, test := range tests {
		if got := test.first.IsOrthogonalTo(test.second, delta); got != test.want {
			t.Errorf("(%v).IsOrthogonalTo(%v) = %v", test.first, test.second, got)
		}
	}
}

func TestStringMethod(t *testing.T){
	var tests = []struct {
		state *qubit.State
		want  string
	}{
		{state.Zero(), "|0>"},
		{state.One(), "|1>"},
		{state.Plus(), "|+>"},
		{state.Minus(), "|->"},
		{state.PlusI(), "|+i>"},
		{state.MinusI(), "|-i>"},
		{state.New(3, 4), "0.6|0> + 0.7999999999999999|1>"},
	}
	for i, test := range tests {
		if got := test.state.String(); got != test.want {
			t.Errorf("%d番目：(%q).String() != %s", i, got, test.want)
		}
	}
}
