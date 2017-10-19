package bb84

import (
	"github.com/waman/qwave/qkd"
	"github.com/waman/qwave/system/qubit/basis"
)

type Eve struct {
	n int
	key qkd.Key
}

func NewEve(n int) *Eve {
  return &Eve{n, nil}
}

func (eve *Eve) Key() qkd.Key {
	return eve.key
}

func (eve *Eve) Eavsedrop(in *qkd.InternalOfChannel) {
	nKey := 0
	for nKey < eve.n {
		select {
		case qubits := <- in.QchFromAlice():
			for _, qbt := range qubits {
				qbt.Observe(basis.Standard)
			}
			in.QchToBob() <- qubits

		case bits := <- in.FromAlice():
			in.ToBob() <- bits

		case bits := <- in.FromBob():
			if len(bits) > 0 {
				for _, match := range bits {
					if match { nKey++ }
				}
			}
			in.ToAlice() <- bits
		}
	}

	//for nKey < eve.n {
	//	in.EavsedropQubits(func(qbts []qubit.Qubit){
	//		for _, qbt := range qbts {
	//			qbt.Observe(basis.Standard)
	//	  }
	//	})
	//
	//	in.ForwardBtoA() // nil
	//	in.ForwardAtoB() // bases
	//
	//	in.EavsedropBtoA(func(matches []bool){
	//		for _, match := range matches {
	//			if match { nKey++ }
	//		}
	//	})
	//}
}