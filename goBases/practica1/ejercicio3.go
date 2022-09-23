package main

import (
	"fmt"
)

func main() {
	//var 1nombre string //corregida, no puede empezar con numero la definicion de la variable
	var unoNombre string
	var apellido string
	//var int edad // corregida, tipo de variable debe ir despues del nombre
	var edad int
	//1apellido := 6 //corregida, no puede empezar con numero la definicion de la variable
	unoApellido := 6
	var licencia_de_conducir = true
	//var estatura de la persona int // corregida la definicion de la variable no puede llevar espacios
	var estaturaDeLaPersona int
	cantidadDeHijos := 2

	//print para prueba de variables
	fmt.Println(unoNombre, apellido, edad, unoApellido, licencia_de_conducir, estaturaDeLaPersona, cantidadDeHijos)

}
