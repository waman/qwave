package qubit

import (
	"testing"
)

func TestStringMethod(t *testing.T){
	var tests = []struct {
		state *State
		want  string
	}{
		{Zero, "|0>"},
		{One, "|1>"},
		{Plus, "0.7071067811865476|0> + 0.7071067811865476|1>"},
		{Minus, "0.7071067811865476|0> - 0.7071067811865476|1>"},
		{PlusI, "0.7071067811865476|0> + 0.7071067811865476i|1>"},
		{MinusI, "0.7071067811865476|0> - 0.7071067811865476i|1>"},
	}
	for i, test := range tests {
		if got := test.state.String(); got != test.want {
			t.Errorf("%d番目：(%q).String() = %s", i, test.state, got)
		}
	}
}
