package bb84

import (
	"github.com/waman/qwave/qkd"
	"fmt"
	"log"
	"time"
	"math/rand"
)

func ExampleBB84Protocol(){
	rand.Seed(time.Now().UnixNano())

	n := 50
	ch := qkd.NewChannel()
	defer ch.Close()
	done := make(chan struct{}, 2)

	alice := NewAlice(n)
	go qkd.EstablishKey(alice, ch, done)

	bob := NewBob(n)
	go qkd.EstablishKey(bob, ch, done)

	<-done
	<-done

	aliceKey, bobKey := alice.Key(), bob.Key()
	log.Printf("Alice's key: %s", aliceKey)
	log.Printf("Bob's key  : %s", bobKey)
	fmt.Println(aliceKey.String() == bobKey.String())
	// Output:
	// true
}

func ExampleBB84ProtocolWithEve(){
	rand.Seed(time.Now().UnixNano())

	n := 50
	ch := qkd.NewInsecureChannel()
	defer ch.Close()
	done := make(chan struct{}, 2)

	eve := NewEve()
	go eve.Eavesdrop(ch.Internal())

	alice := NewAlice(n)
	go qkd.EstablishKey(alice, ch, done)

	bob := NewBob(n)
	go qkd.EstablishKey(bob, ch, done)

	<-done
	<-done
	eve.Stop()

	aliceKey, bobKey := alice.Key(), bob.Key()
	log.Printf("Alice's key: %s", aliceKey)
	log.Printf("Bob's key  : %s", bobKey)
	fmt.Println(aliceKey.String() == bobKey.String())
	// Output:
	// false
}

func ExampleExercise2_9(){
	rand.Seed(time.Now().UnixNano())

	const nTry = 10000
	for nKey := 6; nKey < 11; nKey++ {
		matched := 0
		for i := 0; i < nTry; i++ {
			ch := qkd.NewInsecureChannel()
			done := make(chan struct{}, 2)

			eve := NewEve()
			go eve.Eavesdrop(ch.Internal())

			alice := NewAlice(nKey)
			go qkd.EstablishKey(alice, ch, done)

			bob := NewBob(nKey)
			go qkd.EstablishKey(bob, ch, done)

			<-done
			<-done
			eve.Stop()
			ch.Close()

			if alice.Key().String() != bob.Key().String() {
				matched++
			}
		}
		log.Printf("%d: %.3f", nKey, float32(matched)/float32(nTry))
	}
	fmt.Println("Done.")
	// Output:
	// Done.
}