package state

import (
	"testing"
)

const delta float64 = 1e-10

func TestEqualState(t *testing.T){
	var tests = []struct {
		first, second *State
		want  bool
	}{
		{Zero()  , New(1, 0), true},
		{One()   , New(0, 1), true},
		{Plus()  , New(1, 1), true},
		{Minus() , New(1, -1), true},
		{PlusI() , New(1, 1i), true},
		{MinusI(), New(1, -1i), true},
		{Zero(), New(0, 1), false},
	}
	for _, test := range tests {
		if got := test.first.EqualState(test.second, delta); got != test.want {
			t.Errorf("(%v).Equal(%v) = %v", test.first, test.second, got)
		}
	}
}

func TestIsOrthogonalTo(t *testing.T){
	var tests = []struct {
		first, second *State
		want  bool
	}{
		{Zero(), One(), true},
		{Plus(), Minus(), true},
		{PlusI(), MinusI(), true},
		{Zero(), Plus(), false},
		{Zero(), MinusI(), false},
	}
	for _, test := range tests {
		if got := test.first.IsOrthogonalTo(test.second, delta); got != test.want {
			t.Errorf("(%v).IsOrthogonalTo(%v) = %v", test.first, test.second, got)
		}
	}
}

func TestStringMethod(t *testing.T){
	var tests = []struct {
		state *State
		want  string
	}{
		{Zero(), "|0>"},
		{One(), "|1>"},
		{Plus(), "|+>"},
		{Minus(), "|->"},
		{PlusI(), "|+i>"},
		{MinusI(), "|-i>"},
		{New(3, 4), "0.6|0> + 0.7999999999999999|1>"},
	}
	for i, test := range tests {
		if got := test.state.String(); got != test.want {
			t.Errorf("%d番目：(%q).String() != %s", i, got, test.want)
		}
	}
}
