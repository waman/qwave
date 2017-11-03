package qkd

import (
	"math/rand"
)

// ProperBitCount == 31 or 63 (PC-bit dependent)
const ProperBitCount = (32 << (^uint(0) >> 63)) - 1

func NewRandomBits(n int) []bool {
	bs := make([]bool, n)
	index := 0
	for {
		var max int
		var isLast bool
		if m := n-index; m <= ProperBitCount {
			max = m
			isLast = true
		}else{
			max = ProperBitCount
			isLast = false
		}

		r := rand.Int()
		for i := 0; i < max; i++ {
			bs[index] = r & 1 == 1
			r >>= 1
			index++
		}

		if isLast { break }
	}
	return bs
}

// AppendMatchingBits function append the i-th bit of 'bits' to 'key'
// if the i-th bit of 'matches' is true.
// If the length of key reaches to 'max', the control is returned.
func AppendMatchingBits(key, bits, matches []bool, max int) []bool {
	for i, match := range matches {
		if match {
			key = append(key, bits[i])
			if len(key) == max { return key }
		}
	}
	return key
}

// AppendMatchingBitsLogged function is simlar to AppendMatchingBits function
// but returned value.
//
// The first returned value is the new key.
// This is substituted to the 'key' variable like append function to slice.
// The second returned value is consumed bits when appended. This contains discarded bits.
func AppendMatchingBitsLogged(key, bits, matches []bool, max int) ([]bool, int) {
	for i, match := range matches {
		if match {
			key = append(key, bits[i])
			if len(key) == max { return key, i }
		}
	}
	return key, len(bits)
}

func EstablishKeys(alice Alice, bob Bob) (aliceKey, bobKey Key) {
	ch := NewChannel()
	defer ch.Close()
	done := make(chan struct{}, 2)

	go func(){
		alice.EstablishKey(ch.OnAlice())
		done <- struct{}{}
	}()

	go func(){
		bob.EstablishKey(ch.OnBob())
		done <- struct{}{}
	}()

	<-done
	<-done

	aliceKey = alice.Key()
	bobKey = bob.Key()
	return
}

func EstablishKeysWithEavesdropping(
	  alice Alice, bob Bob, eve Eve) (aliceKey, bobKey, eveKey Key) {

	ch := NewInsecureChannel()
	defer ch.Close()
	done := make(chan struct{}, 2)

	go eve.Eavesdrop(ch.Internal())

	go func(){
		alice.EstablishKey(ch.OnAlice())
		done <- struct{}{}
	}()

	go func(){
		bob.EstablishKey(ch.OnBob())
		done <- struct{}{}
	}()

	<-done
	<-done
	eve.Stop()

	aliceKey = alice.Key()
	bobKey = bob.Key()
	if e, ok := eve.(KeyContainer); ok { eveKey = e.Key() }
	return
}