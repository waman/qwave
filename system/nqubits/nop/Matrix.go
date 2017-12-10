package nop

type Matrix interface {
	Dim() int
	At(i, j int) complex128

	Equals(v Matrix) bool
	String() string

	CloneMutable() MutableMatrix

	Apply(v []complex128) []complex128
	ApplyToMap(v map[int]complex128) map[int]complex128

	Dense() Matrix
	Sparse() Matrix
}

type MutableMatrix interface {
	Matrix
	Set(i, j int, c complex128)
	Mul(x, y Matrix) Matrix
}

func New(n int, cs ...complex128) Matrix {
	cs_ := make([][]complex128, n)
	for i := 0; i < n; i++ {
		cs_[i] = make([]complex128, n)
		if rest := len(cs); rest >= n {
			copy(cs_[i], cs[:n])
			cs = cs[n:]
		} else if rest > 0 {
			copy(cs_[i], cs)
			cs = nil
		}
	}
	return NewDense(cs_, false)
}

func equalMatrices(u, v Matrix) bool {

	if u == v { return true }
	n := v.Dim()
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if u.At(i, j) != v.At(i, j) {
				return false
			}
		}
	}
	return true
}

func Prod(m Matrix, ms ...Matrix) Matrix {
	if len(ms) == 0 { return m }

	var result = m.CloneMutable()
	for _, m := range ms {
		result.Mul(result, m)
	}

	return result
}