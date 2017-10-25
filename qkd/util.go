package qkd

import (
	"math/rand"
	"log"
	"time"
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
//
// The first returned value is the new key.
// This is substituted to the 'key' variable like append function to slice.
// The second returned value is consumed bits when appended. This contains discarded bits.
func AppendMatchingBits(key, bits, matches []bool, max int) ([]bool, int) {
	for i, match := range matches {
		if match {
			key = append(key, bits[i])
			if len(key) == max {
				return key, i
			}
		}
	}
	return key, len(bits)
}

func EstablishKey(kc KeyContainer, ch Channel, done chan<- struct{}){
	if alice, ok := kc.(Alice); ok {
		alice.EstablishKey(ch.OnAlice())
		done <- struct{}{}

	} else if bob, ok := kc.(Bob); ok {
		bob.EstablishKey(ch.OnBob())
		done <- struct{}{}

	}else{
		log.Panicf("KeyContainer must be qkd.Alice or qkd.Bob: %T", kc)
	}
}

func EstablishKeys(alice Alice, bob Bob) (aliceKey, bobKey Key) {
	rand.Seed(time.Now().UnixNano())

	ch := NewChannel()
	defer ch.Close()
	done := make(chan struct{}, 2)

	go EstablishKey(alice, ch, done)

	go EstablishKey(bob, ch, done)

	<-done
	<-done

	return alice.Key(), bob.Key()
}

func EstablishKeysWithEavesdropping(alice Alice, bob Bob, eve Eve) (aliceKey, bobKey, eveKey Key) {
	rand.Seed(time.Now().UnixNano())

	ch := NewInsecureChannel()
	defer ch.Close()
	done := make(chan struct{}, 2)

	go eve.Eavesdrop(ch.Internal())
	go EstablishKey(alice, ch, done)
	go EstablishKey(bob, ch, done)

	<-done
	<-done
	eve.Stop()

	if e, ok := eve.(KeyContainer); ok {
		return alice.Key(), bob.Key(), e.Key()
	} else {
		return alice.Key(), bob.Key(),nil
	}
}