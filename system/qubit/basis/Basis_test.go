package basis

import (
	"github.com/waman/qwave/system/qubit/ket"
	"fmt"
)

func ExampleBasisNew(){
	delta := 1e-10
	b := New(ket.New(3, 4, false), ket.New(4, -3, false), delta)
	fmt.Println(b)
	// Output:
	// {0.6|0> + 0.8|1>, 0.8|0> - 0.6|1>}
}

func ExampleBasisMethods(){
	b := Standard()
	fmt.Println(b.First() == ket.Zero())
	fmt.Println(b.Second() == ket.One())
	// Output:
	// true
	// true
}
