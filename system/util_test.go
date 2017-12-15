package system

import (
	"testing"
	"github.com/ToQoz/gopwt/assert"
)

func TestStringMethod(t *testing.T){
	var tests = []struct {
		cs []complex128
		want  string
	}{
		{[]complex128{2, 3, 4, 5}, "2|00> + 3|01> + 4|10> + 5|11>"},
		{[]complex128{2, -3, 4i, -5i, 6, -7, 8i, -9i},
		  "2|000> - 3|001> + 4i|010> - 5i|011> + 6|100> - 7|101> + 8i|110> - 9i|111>"},
	}
	for _, test := range tests {
		assert.OK(t, ToString(test.cs...) == test.want)
	}
}