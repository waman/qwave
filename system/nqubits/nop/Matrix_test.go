package nop

import (
	"fmt"
	"testing"
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
	for i, test := range tests {
		if got := test.mat.String(); got != test.want {
			t.Errorf("nop.New() %d番目：%s != %s",	i, got, test.want)
		}
	}
}