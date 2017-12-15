package qubit

import (
	"testing"
	"github.com/waman/qwave/system/qubit/basis"
	"github.com/waman/qwave/system/qubit/ket"
	"github.com/waman/qwave/system/qubit/op"
	"github.com/ToQoz/gopwt/assert"
)

func TestEquality(t *testing.T){
	q1 := NewOne()
	q2 := NewOne()
	assert.OK(t, q1 != q2)
}

func TestObserve(t *testing.T){
	qbt := NewPlus()
	std := basis.Standard()
	result := qbt.Observe(std)
	assert.OK(t, result == std.First() || result == std.Second())
}

func TestObserveInTheSameBasis(t *testing.T){
	qbt := NewPlus()
	std := basis.Standard()
	result := qbt.Observe(std)
	for i := 0; i < 100; i++ {
		assert.OK(t, qbt.Observe(std) == result)
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


func TestU(t *testing.T){
	qbt := NewZero()
	X := op.PauliX()
	U(X, X, X)(qbt)

	if qbt.ObserveInStandardBasis() != ket.One() {
		t.Errorf("U()")
	}
}
