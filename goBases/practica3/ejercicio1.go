package main

import (
	"errors"
	"fmt"
)

/*
Ejercicio 1 - Impuestos de salario
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo,
para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo y
si gana más de $150.000 se le descontará además un 10%.
*/
const (
	firstTax  float64 = 0.17
	secondTax float64 = 0.10
)

func calcDiscount(value float64, tax float64) float64 {
	return value * tax
}

// funcion calcula el impuesto correspondiente en base al salario
func salaryTaxes(value float64) (float64, error) {
	if value <= 0 {
		return value, errors.New("El valor no puede ser negativo")

	}
	switch {
	case value > 100_000:
		return value - calcDiscount(value, firstTax) - calcDiscount(value, secondTax), nil
	case value > 50_000:
		return value - calcDiscount(value, firstTax), nil
	default:
		return value, nil
	}
}
func main() {
	//TODO pendiente formatear el resultado para que imprima
	//fmt.Printf(salaryTaxes(40000))
	salary, err := salaryTaxes(60_000)
	if err != nil {
		panic(err)
	}
	fmt.Println(salary)
	fmt.Println()
	salary1, err1 := salaryTaxes(400_000)
	if err1 != nil {
		panic(err1)
	}
	fmt.Println(salary1)
	fmt.Println()
	salary2, err2 := salaryTaxes(-4000)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(salary2)

}
