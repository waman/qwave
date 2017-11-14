package system

import (
	"bytes"
	"fmt"
	"math"
)

var s2i = complex(1/math.Sqrt(2), 0)

// 1/√2 (sqrt 2 inverse)
func S2I() complex128 {
	return s2i
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
