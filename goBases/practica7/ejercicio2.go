package main

import (
	"errors"
	"fmt"
	"os"
)

/*

Ejercicio 2 - Impuestos de salario #2
------------------------------------------------------------------------------
Haz lo mismo que en el ejercicio anterior pero reformulando el código para que,
 en reemplazo de “Error()”,  se implemente “errors.New()”.
-------------------------------------------------------------------------------

(Ejercicio 1 - Impuestos de salario #1
En tu función “main”, define una variable llamada “salary” y
asignarle un valor de tipo “int”.
Crea un error personalizado con un struct que implemente “Error()”
con el mensaje “error: el salario ingresado no alcanza el mínimo
imponible" y lánzalo en caso de que “salary” sea menor a 150.000.
Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.)
*/
/*
type custError struct {
	status int
	msg    string
}

func (e custError) Error() string {
	return fmt.Sprintf("%d -.- %v", e.status, e.msg)
}
func customError(salary int) (int, error) {
	if salary < 150_000 {

		return 400, &custError{
			400,
			"error: el salario ingresado no alcanza el mínimo imponible",
		}
	}
	return salary, nil

}*/

func main() {
	var salary int = 140_000
	//_, erro := customError(salary)
	if salary < 150_000 {
		fmt.Println(errors.New("error: el salario ingresado no alcanza el mínimo imponible"))
		os.Exit(1)
	}
	fmt.Println("Debe pagar impuesto")
}
