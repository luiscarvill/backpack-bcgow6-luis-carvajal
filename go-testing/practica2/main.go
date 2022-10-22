package main

import (
	"fmt"

	"github.com/backpack-bcgow6-luis-carvajal/go_testing/practica2/pkg/calculadora"
	"github.com/backpack-bcgow6-luis-carvajal/go_testing/practica2/pkg/ordenamiento"
)

func main() {
	a, b := 10, 5
	//sum, _ := operaciones.Add(a, b)
	//fmt.Println(sum)
	resta, _ := calculadora.Restar(a, b)
	fmt.Println(resta)
	// test practica 1
	orderSlice, _ := ordenamiento.OrderSlice(3, 2, 2, 1, -1, 4)
	orderSlice2, _ := ordenamiento.OrderSlice()
	fmt.Println(orderSlice)
	fmt.Println(orderSlice2)

}
