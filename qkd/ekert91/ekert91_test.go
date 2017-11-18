package ekert91

import (
	"github.com/waman/qwave/qkd"
	"log"
	"fmt"
)

func ExampleEkert91Protocol(){
	n := 10000
	aliceKey, bobKey := qkd.EstablishKeys(NewAlice(n), NewBob(n))

	log.Printf("Alice's key: %s", aliceKey)
	log.Printf("Bob's key  : %s", bobKey)
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

	//log.Printf("Alice's key: %s", aliceKey)
	//log.Printf("Bob's key  : %s", bobKey)
	log.Printf("Concordance rate: %f", aliceKey.ConcordanceRate(bobKey))
	fmt.Println(aliceKey.Equals(bobKey))
	// Output:
	// false
}
