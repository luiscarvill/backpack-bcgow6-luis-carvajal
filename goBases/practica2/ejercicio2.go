package main

import (
	"fmt"
)

var (
	edad       int     = 20
	empleados  bool    = false
	antiguedad int     = 0
	salario    float64 = 100000
)

func main() {

	switch {
	case edad < 22:
		fmt.Println("El cliente no cuple la edad minima mayor a 22")
		fallthrough
	case empleados == false:
		fmt.Println("El cliente no cumple el requerimiento de ser empleado")
		fallthrough
	case antiguedad < 1:
		fmt.Println("El cliente no tiene el tiempo laboral requerido")
	case salario >= 100000:
		fmt.Println("El cliente cumple los requerimientos y no se le cobra intereses")
	default:
		fmt.Println("El cliente cumple los requerimientos para un prestamo")
	}
}
