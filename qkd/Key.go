package qkd

import "bytes"

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

type KeyContainer interface {
	Key() Key
}