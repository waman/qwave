package qubits2

import (
	"github.com/waman/qwave/system/qubits2/basis2"
	"fmt"
	"github.com/waman/qwave/system/qubits2/ket2"
	"testing"
	"github.com/ToQoz/gopwt/assert"
)

func ExampleKet2New(){
	w2 := ket2.New(1, 1, 1, 1, false)
  fmt.Println(w2)
  // Output:
  // 0.5|00> + 0.5|01> + 0.5|10> + 0.5|11>
}

func Example(){
	std := basis2.Standard()
	fmt.Println(std)

	bell := basis2.Bell()
	fmt.Println(bell)
	// Output:
	// {|00>, |01>, |10>, |11>}
	// {|Φ+>, |Φ->, |Ψ+>, |Ψ->}
}

func TestInstanceEquality(t *testing.T){
	qbts0 := NewPhiPlus()
	qbts1 := NewPhiPlus()

	assert.OK(t, qbts0 != qbts1)
}