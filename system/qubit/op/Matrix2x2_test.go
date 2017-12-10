package op_test

import (
	"testing"
	. "github.com/waman/qwave/system/qubit/op"
)

func TestProd(t *testing.T){
	result := Prod(PauliX(), PauliX())

	if !result.Equals(Id()) {
		t.Errorf("InnerProduct()")
	}
}