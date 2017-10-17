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
	go func() {
		alice.EstablishKey(ch.OnAlice())
		done <- struct{}{}
	}()

	bob := NewBob(n)
	go func(){
		bob.EstablishKey(ch.OnBob())
		done <- struct{}{}
	}()

	<-done
	<-done

	log.Printf("Alice's key: %s\n", alice.Key().String())
	log.Printf("Bob's key  : %s\n", bob.Key().String())
	fmt.Println(alice.Key().String() == bob.Key().String())
	// Output:
	// true
}

func ExampleBB84ProtocolWithEve(){
	n := 40
	ch := qkd.NewInsecureChannel()
	defer ch.Close()
	done := make(chan struct{}, 2)

	eve := NewEve(n)
	go func(){
		eve.Eavsedrop(ch.Internal())
		done <- struct{}{}
	}()

	alice := NewAlice(n)
	go func(){
		alice.EstablishKey(ch.OnAlice())
		done <- struct{}{}
	}()

	bob := NewBob(n)
	go func(){
		bob.EstablishKey(ch.OnBob())
		done <- struct{}{}
	}()

	<-done
	<-done
	<-done

	fmt.Printf("Alice's key: %s\n", alice.Key().String())
	fmt.Printf("Bob's key  : %s\n", bob.Key().String())
	fmt.Println(alice.Key().String() == bob.Key().String())
	// Output:
	// false
}
