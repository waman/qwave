package system

import (
	"bytes"
	"fmt"
	"math"
	"math/cmplx"
)

var s2i = complex(1/math.Sqrt(2), 0)

// 1/√2 (sqrt 2 inverse)
func S2I() complex128 {
	return s2i
}

func Dim(qbtCount int) int {
	return 1 << uint(qbtCount)
}

func SliceToMap(cs []complex128) map[int]complex128 {
	cMap := make(map[int]complex128)
	for i, c := range cs {
		if c != 0 {
			cMap[i] = c
		}
	}
	return cMap
}

func CreateCopy(cs []complex128) []complex128 {
	cs_ := make([]complex128, len(cs))
	copy(cs_, cs)
	return cs_
}

func CreateCopyMap(cMap map[int]complex128) map[int]complex128 {
	cMap_ := make(map[int]complex128)
	for i, c := range cMap {
		cMap_[i] = c
	}
	return cMap_
}

func Normalize(cs []complex128){
	nf := complex(1.0/math.Sqrt(Norm2(cs)), 0)
	for i, c := range cs {
		cs[i] = c*nf
	}
}

func NormalizeMap(cMap map[int]complex128){
	nf := complex(1.0/math.Sqrt(Norm2Map(cMap)), 0)
	for i, c := range cMap {
		cMap[i] = c*nf
	}
}

func Norm2(cs []complex128) float64 {
	norm2 := 0.0
	for _, c := range cs {
		abs2 := cmplx.Abs(c)
		norm2 += abs2*abs2
	}
	return norm2
}

func Norm2Map(cMap map[int]complex128) float64 {
	norm2 := 0.0
	for _, c := range cMap {
		abs2 := cmplx.Abs(c)
		norm2 += abs2*abs2
	}
	return norm2
}

func ToString(cs ...complex128) string {
	var buf bytes.Buffer

	n := len(cs)
	d := 0  // n = 2^d
	for m := n; m != 1; m /= 2 { d++ }

	i := 0
	for _, c := range cs {
		if c != 0 { break }
		i++
	}

	c := cs[i] // Note c != 0
	if re, im := real(c), imag(c); im == 0 {
		if re == -1 {
			buf.WriteString("- ")
		}else if re != 1 {
			fmt.Fprint(&buf, re)
		}
		// if re == 1, do nothing
	} else if re == 0 {
		if im == 1 {
			buf.WriteString("i")
		} else if im == -1 {
			buf.WriteString("- i")
		}else {
			fmt.Fprint(&buf, im, "i")
		}
	} else {
		fmt.Fprint(&buf, c)
	}
	fmt.Fprintf(&buf, "|%0*b>", d, i)
	i++

	for n := len(cs); i < n; i++ {
		c := cs[i]
		if c == 0 { continue }

		if re, im := real(c), imag(c); im == 0 {
			if re > 0 {
				fmt.Fprint(&buf, " + ", re)
			} else if re < 0 {
				fmt.Fprint(&buf, " - ", -re)
			}
		} else if re == 0 {
			if im > 0 {
				fmt.Fprint(&buf, " + ", im, "i")
			} else if im < 0 {
				fmt.Fprint(&buf, " - ", -im, "i")
			}
		} else {
			fmt.Fprint(&buf, " + ", c)
		}

		fmt.Fprintf(&buf, "|%0*b>", d, i)
	}

	return buf.String()
}
