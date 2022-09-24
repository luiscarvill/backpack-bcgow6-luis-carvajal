package main

import (
	"fmt"
)

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	//keys := reflect.ValueOf(employees).MapKeys()
	fmt.Println("Benjamin: ", employees["Benjamin"])
	fmt.Println("=========================")
	for key, value := range employees {
		switch {
		case value >= 21:
			fmt.Printf("Nombre: %v - Edad: %v\n", key, value)
		default:
			continue
		}
	}
	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println("=========================\n")
	fmt.Println(employees)

}
