package system

import (
	"bytes"
	"fmt"
	"math"
	"math/cmplx"
)

type ComplexContainer interface {
	At(i int) complex128
}

var s2i = complex(1/math.Sqrt(2), 0)

// 1/√2 (sqrt 2 inverse)
func S2I() complex128 {
	return s2i
}

func Dim(qbtCount int) int {
	return 1 << uint(qbtCount)
}

func SliceToMap(cs []complex128) map[int]complex128 {
	var cMap map[int]complex128
	for i, c_i := range cs {
		if c_i != 0 {
			if cMap == nil {
				cMap = make(map[int]complex128)
			}
			cMap[i] = c_i
		}
	}
	return cMap
}

func MapToSlice(n int, cMap map[int]complex128) []complex128 {
	cs := make([]complex128, n)
	for i, c_i := range cMap {
		cs[i] = c_i
	}
	return cs
}

func CreateCopy(cs []complex128) []complex128 {
	_cs := make([]complex128, len(cs))
	copy(_cs, cs)
	return _cs
}

func CreateCopyMap(cMap map[int]complex128) map[int]complex128 {
	_cMap := make(map[int]complex128)
	for i, c_i := range cMap {
		_cMap[i] = c_i
	}
	return _cMap
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
	if len(cs) == 2 {
		a, b := cs[0], cs[1]
		if b == 0 { return "|0>"}
		if a == 0 { return "|1>"}
		if b ==  a { return "|+>"}
		if b == -a { return "|->"}
		if b ==  a*1i { return "|+i>"}
		if b == -a*1i { return "|-i>"}

	} else if len(cs) == 4 {
		a, b, c, d := cs[0], cs[1], cs[2], cs[3]
		if b == 0 && c == 0 && d == 0 { return "|00>"}
		if a == 0 && c == 0 && d == 0 { return "|01>"}
		if a == 0 && b == 0 && d == 0 { return "|10>"}
		if a == 0 && b == 0 && c == 0 { return "|11>"}

		if b == 0 && c == 0 && a ==  d { return "|Φ+>"}
		if b == 0 && c == 0 && a == -d { return "|Φ->"}
		if a == 0 && d == 0 && b ==  c { return "|Ψ+>"}
		if a == 0 && d == 0 && b == -c { return "|Ψ->"}
	}

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
