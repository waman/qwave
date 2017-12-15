package nop

import (
	"fmt"
	"testing"
	"github.com/ToQoz/gopwt/assert"
)

func ExampleNew(){
  mat := New(2, 0, 1, 1, 0)
  fmt.Println(mat)
  // Output:
  // [[(0+0i) (1+0i)] [(1+0i) (0+0i)]]
}

func TestNew(t *testing.T){
	var tests = []struct {
		mat Matrix
		want string
	}{
		{New(2, 1, 1), "[[(1+0i) (1+0i)] [(0+0i) (0+0i)]]"},
		{New(2, 0, 1, 1, 0), "[[(0+0i) (1+0i)] [(1+0i) (0+0i)]]"},
	}
	for _, test := range tests {
		assert.OK(t, test.mat.String() == test.want)
	}
}