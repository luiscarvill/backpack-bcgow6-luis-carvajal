package calculadora

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResta(t *testing.T) {
	a, b, c, d := 10, 5, 12, 0
	esperado := 5
	result, err := Restar(a, b)
	fmt.Println("================== Test1 error nil ==================")
	assert.Nil(t, err)

	fmt.Printf("============= Test2 result %d - %d ==============\n", a, b)
	assert.Equal(t, esperado, result, "Resultado (%d) es distinto al esperado (%d)", result, esperado)

	fmt.Printf("============= Test3 BadSolution %d - %d ==============\n", c, d)
	_, err1 := Restar(c, d)
	assert.NotNil(t, err1)

	fmt.Printf("============= Test4 BadSolution %d - %d ==============\n", d, c)
	_, err2 := Restar(d, c)
	assert.NotNil(t, err2)
}
