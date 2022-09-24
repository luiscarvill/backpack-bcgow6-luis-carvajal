package main

import (
	"fmt"
)

func main() {
	var palabra string = "Luis Carvajal"
	fmt.Printf("%v\n", len(palabra))
	fmt.Println("===================")
	var s []string
	for i := 0; i < len(palabra); i++ {
		fmt.Printf("%c\n", palabra[i])
		//prueba
		s = append(s, string(palabra[i]))
	}

	fmt.Printf("Array: %v\n", palabra)
}
