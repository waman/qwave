package bb84

import (
	"github.com/waman/qwave/qkd"
	"fmt"
	"log"
)

func ExampleBB84Protocol(){
	n := 40
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
	n := 40
	ch := qkd.NewInsecureChannel()
	defer ch.Close()
	done := make(chan struct{}, 3)

	eve := NewEve(n)
	go qkd.EavseDrop(eve, ch, done)

	alice := NewAlice(n)
	go qkd.EstablishKey(alice, ch, done)

	bob := NewBob(n)
	go qkd.EstablishKey(bob, ch, done)

	<-done
	<-done
	<-done

	aliceKey, bobKey := alice.Key(), bob.Key()
	log.Printf("Alice's key: %s", aliceKey)
	log.Printf("Bob's key  : %s", bobKey)
	fmt.Println(aliceKey.String() == bobKey.String())
	// Output:
	// false
}