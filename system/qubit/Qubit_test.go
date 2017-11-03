package qubit

import (
	"testing"
	"github.com/waman/qwave/system/qubit/basis"
)

func TestEquality(t *testing.T){
	q1 := NewOne()
	q2 := NewOne()
	if q1 == q2 {
		t.Errorf("== operator")
	}
}

func TestObserve(t *testing.T){
	qbt := NewPlus()
	std := basis.Standard()
	result := qbt.Observe(std)
	if !(result == std.First() ||	result == std.Second()){
		t.Errorf("Observe method")
	}
}

func TestObserveInTheSameBasis(t *testing.T){
	qbt := NewPlus()
	std := basis.Standard()
	result := qbt.Observe(std)
	for i := 0; i < 100; i++ {
		if qbt.Observe(std) != result {
			t.Errorf("Observe method")
		}
	}
}

//func ExampleObserve(){
//	result := make(map[*ket.State]int)
//	for i := 0; i < 10000; i++ {
//		s := NewPlus().Observe(basis.Standard())
//		if s == ket.Zero() {
//			result[ket.Zero()]++
//		}else{
//			result[ket.One()]++
//		}
//	}
//	fmt.Println(result)
//	// Output:
//	// map[|0>:5000 |1>:5000]
//}
