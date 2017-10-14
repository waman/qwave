package main

import (
	"fmt"
	"github.com/waman/qwave/qkd"
	"github.com/waman/qwave/qkd/bb84"
)

func main(){
	n := 40
	ch := qkd.NewChannel()
	done := make(chan struct{}, 2)

	alice := bb84.NewAlice(n)
	go alice.EstablishKey(ch.OnAlice(), done)

	bob := bb84.NewBob(n)
	go bob.EstablishKey(ch.OnBob(), done)

	<-done
	<-done
	ch.Close()

	fmt.Printf("Alice's key: %s\n", alice.Key().String())
	fmt.Printf("Bob's key  : %s\n", bob.Key().String())
	fmt.Println(alice.Key().String() == bob.Key().String())
}
