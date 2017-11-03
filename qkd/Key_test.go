package qkd

import (
	"testing"
	"log"
)

func TestKey_ConcordanceRate(t *testing.T) {
	key0 := Key([]bool{true, true, true, true})
	key1 := Key([]bool{true, false, false})

	log.Print(key0.ConcordanceRate(key1))  // 0.5

	var tests = []struct {
		key0, key1 Key
		want  float32
	}{
		{[]bool{true, true}, []bool{false, false}, 0.0},
		{[]bool{true, true, true}, []bool{true, true, true}, 1.0},
		{[]bool{true, true, true, true}, []bool{true, false, false, true}, 0.5},
	}
	for _, test := range tests {
		if got := test.key0.ConcordanceRate(test.key1); got != test.want {
			t.Errorf("(%v).ConcodanceRate(%v) = %v != %v", test.key0, test.key1, got, test.want)
		}
	}
}