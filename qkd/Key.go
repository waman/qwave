package qkd

import (
	"bytes"
	"log"
)

type Key []bool

func (key Key) String() string {
	var buf bytes.Buffer
	for _, b := range key {
		if b {
			buf.WriteString("1")
		}else{
			buf.WriteString("0")
		}
	}
	return buf.String()
}

func (key1 Key) Equals(key2 Key) bool {
	return key1.String() == key2.String()
}

func (key1 Key) Concordance(key2 Key) int {
	n := len(key1)
	if n2 := len(key2); n2 != n {
		log.Panicf("Two keys must have the same length: %d, %d", n, n2)
	}

	matched := 0
	for i, bit1 := range key1 {
		if bit1 == key2[i] { matched++ }
	}
	return matched
}

func (key1 Key) ConcordanceRate(key2 Key) float32 {
	return float32(key1.Concordance(key2))/float32(len(key1))
}

type KeyContainer interface {
	Key() Key
}