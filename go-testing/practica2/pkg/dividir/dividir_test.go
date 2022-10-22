package dividir

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	a, b, c, d := 1, 2, 0, 3

	result1, err1 := Dividir(a, b)
	fmt.Println("================== Test division ==================")
	fmt.Println("================== Test1 division igual ==================")
	assert.Equal(t, c, result1)
	assert.Nil(t, err1)
	result2, _ := Dividir(b, a)
	assert.Equal(t, b, result2)
	fmt.Println("================== Test2 error division 0  ==================")
	result2, err2 := Dividir(d, c)
	assert.ErrorContains(t, err2, "No es posible dividir por 0")
	fmt.Println("================== Test3 error return 0  ==================")
	assert.Equal(t, c, result2)
	fmt.Println("================== Test4 error return 0  ==================")
	result3, _ := Dividir(b, b)
	assert.Equal(t, a, result3)

}
