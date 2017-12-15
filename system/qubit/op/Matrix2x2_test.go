package op_test

import (
	"testing"
	. "github.com/waman/qwave/system/qubit/op"
	"github.com/ToQoz/gopwt/assert"
	"log"
)

func TestProd(t *testing.T){
	result := Prod(PauliX(), PauliX())
  log.Println(result)
	assert.OK(t, result.Equals(Id()))
}