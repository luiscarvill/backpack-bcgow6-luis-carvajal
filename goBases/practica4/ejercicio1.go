package main

import "fmt"

/*
Ejercicio 1 - Registro de estudiantes
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle
*/
type Alumno struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    string
}
type Alumnos struct {
	Alumn []Alumno
}

func (a *Alumnos) Read() ([]Alumno, error) {
	return a.Alumn, nil
}

func main() {
	alm := Alumno{"Luis", "Carvajal", "1221414", "20/09/2022"}
	alms := &Alumnos{Alumn: []Alumno{alm}}
	alms.Read()
	fmt.Println(alms.Read())
}
