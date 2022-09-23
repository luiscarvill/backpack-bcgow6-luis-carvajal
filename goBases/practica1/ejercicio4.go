package main

import (
	"fmt"
)

func main() {
	var apellido string = "Gomez"
	//var edad int = "35" // Corregido, variable definida como int contiene valor string
	var edad int = 35

	boolean := "false"
	//var sueldo string = 45857.90 // Corregido,  variable strig contiene valor numerico
	//opcion 2 var sueldo float32 = 45857.90
	var sueldo string = "45857.90"

	var nombre string = "Julián"

	//print para prueba de variables
	fmt.Println(apellido, edad, boolean, sueldo, nombre)

}
