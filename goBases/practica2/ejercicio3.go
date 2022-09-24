package main

import (
	"fmt"
)

func main() {
	var mes int = 1
	meses := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}
	mesT := mes - 1
	switch {
	case mesT >= 0 && mesT <= 11:
		fmt.Printf("%v\n", meses[mesT])
	default:
		fmt.Printf("Valor ingresado invalido, seleccione una opcion valida 1 - 12")
	}
}
