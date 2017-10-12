package bb84

import (
	"fmt"
)

func ExampleBB84Protocol(){
	n := 20
	ch := NewChannel(n)
	done := make(chan struct{})

	alice := NewAlice(n)
	go alice.EstablishKey(ch.OnAlice(), done)

	bob := NewBob(n)
	go bob.EstablishKey(ch.OnBob())

	<-done
	fmt.Printf("Alice's key: %s", alice.key)
	fmt.Printf("Bob's key  : %s", bob.key)
}
