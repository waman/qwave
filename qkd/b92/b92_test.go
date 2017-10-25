package b92

import (
	"github.com/waman/qwave/qkd"
	"fmt"
	"log"
	"time"
	"math/rand"
)

func ExampleB92Protocol(){
	n := 50
	aliceKey, bobKey := qkd.EstablishKeys(NewAlice(n), NewBob(n))

	log.Printf("Alice's key: %s", aliceKey)
	log.Printf("Bob's key  : %s", bobKey)
	fmt.Println(aliceKey.Equals(bobKey))
	fmt.Println(aliceKey.ConcordanceRate(bobKey))
	// Output:
	// true
	// 1
}

type LoggingAlice struct {
	Alice
	Consumed int
}

func (alice *LoggingAlice) EstablishKey(ch qkd.ChannelOnAlice){
	for len(alice.key) < alice.n {
		bits  := qkd.NewRandomBits(qkd.ProperBitCount)
		ch.Qch() <- encode(bits)

		matches := <- ch.FromBob()
		newKey, m := qkd.AppendMatchingBits(alice.key, bits, matches, alice.n)
		alice.key = newKey
		alice.Consumed += m
	}
}

func ExampleB92ProtocolWithLoggingAlice(){
	n := 50
	alice := &LoggingAlice{*NewAlice(n), 0}
	aliceKey, bobKey := qkd.EstablishKeys(alice, NewBob(n))

	log.Printf("Accepted Bit Rate: %f", float32(n)/float32(alice.Consumed))
	fmt.Println(aliceKey.Equals(bobKey))
	fmt.Println(aliceKey.ConcordanceRate(bobKey))
	// Output:
	// true
	// 1
}

func ExampleB92ProtocolWithEve(){
	n := 50
	aliceKey, bobKey, _ := qkd.EstablishKeysWithEavesdropping(
		NewAlice(n), NewBob(n), qkd.NewObservingEve())

	log.Printf("Alice's key: %s", aliceKey)
	log.Printf("Bob's key  : %s", bobKey)
	log.Printf("Concordance rate: %f", aliceKey.ConcordanceRate(bobKey))
	fmt.Println(aliceKey.Equals(bobKey))
	// Output:
	// false
}

func ExampleSuccessRateOfEavesdropping(){
	rand.Seed(time.Now().UnixNano())

	const nTry = 10000
	for nKey := 6; nKey < 11; nKey++ {
		matched := 0
		for i := 0; i < nTry; i++ {
			aliceKey, bobKey, _ := qkd.EstablishKeysWithEavesdropping(
				NewAlice(nKey), NewBob(nKey), qkd.NewObservingEve())

			if aliceKey.Equals(bobKey) { matched++ }
		}
		log.Printf("%d: %.3f", nKey, float32(matched)/float32(nTry))
	}
	fmt.Println("Done.")
	// Output:
	// Done.
}