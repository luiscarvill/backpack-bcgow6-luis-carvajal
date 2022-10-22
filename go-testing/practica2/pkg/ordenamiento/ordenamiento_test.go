package ordenamiento

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	esperado := []int{1, 2, 3, 4, 5}
	esperado2 := []int{1, 1, 2, 2, 2, 3, 4, 5}
	esperado3 := []int{-2, -1, 1, 2, 2, 3, 4, 5}
	esperado4 := []int{0, 0, 0, 0, 0}

	test1, _ := OrderSlice(1, 3, 2, 5, 4)

	fmt.Println("================== Test slice ==================")
	fmt.Println("================== Test1 slice diferente al esperado ==================")
	assert.Equal(t, esperado, test1, "Resultado diferente al esperado test1")

	fmt.Println("================== Test2 slice error valores esperados  ==================")
	_, test2 := OrderSlice()
	assert.Equal(t, errors.New("Debe definin minimo 2 valores para ordenar"), test2)
	test3, _ := OrderSlice(1, 3, 1, 2, 2, 5, 2, 4)

	fmt.Println("================== Test3 slice numeros repetidos ==================")
	assert.Equal(t, esperado2, test3, "Resultado diferente al esperado test3")
	test4, _ := OrderSlice(1, 3, -1, 2, -2, 5, 2, 4)

	fmt.Println("================== Test4 slice numeros negativos ==================")
	assert.Equal(t, esperado3, test4, "Resultado diferente al esperado test4")
	test5, err1 := OrderSlice(0, 0, 0, 0, 0)

	fmt.Println("================== Test5 slice mismo numero ==================")
	assert.Equal(t, esperado4, test5, "Resultado diferente al esperado test5")
	assert.Nil(t, err1)
}
