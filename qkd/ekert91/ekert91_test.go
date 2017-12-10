package ekert91

import (
	"github.com/waman/qwave/qkd"
	"fmt"
)

func ExampleEkert91Protocol(){
	n := 10000
	aliceKey, bobKey := qkd.EstablishKeys(NewAlice(n), NewBob(n))

	// fmt.Printf("Alice's key: %s...\n", aliceKey[:20])
	// fmt.Printf("Bob's key  : %s...\n", bobKey[:20])
	fmt.Println(aliceKey.Equals(bobKey))
	fmt.Println(aliceKey.ConcordanceRate(bobKey))
	// Output:
	// true
	// 1
}

func ExampleEkert91ProtocolWithEve(){
	n := 100000
	aliceKey, bobKey, _ := qkd.EstablishKeysWithEavesdropping(
		NewAlice(n), NewBob(n), qkd.NewObservingEve())

	// fmt.Printf("Alice's key: %s...\n", aliceKey[:20])
	// fmt.Printf("Bob's key  : %s...\n", bobKey[:20])
	fmt.Printf("Concordance rate: %.2f\n", aliceKey.ConcordanceRate(bobKey))
	fmt.Println(aliceKey.Equals(bobKey))
	// Output:
	// Concordance rate: 0.75
	// false
}
