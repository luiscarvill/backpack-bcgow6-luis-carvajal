package main

import (
	"fmt"
	"os"
)

/*

Ejercicio 1 - Impuestos de salario #1

En tu función “main”, define una variable llamada “salary” y
asignarle un valor de tipo “int”.
Crea un error personalizado con un struct que implemente “Error()”
con el mensaje “error: el salario ingresado no alcanza el mínimo
imponible" y lánzalo en caso de que “salary” sea menor a 150.000.
Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.

*/

type custError struct {
	msg string
}

func (e *custError) Error() string {
	return e.msg
}
func customError(salary int) error {
	if salary < 150_000 {
		return &custError{
			"error: el salario ingresado no alcanza el mínimo imponible",
		}
	}
	return nil
}

func main() {
	var salary int = 140_000
	erro := customError(salary)
	if erro != nil {
		fmt.Println(erro)
		os.Exit(1)
	}
	fmt.Println("Debe pagar impuesto")
}
