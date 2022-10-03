package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("./texto.txt")
	//fmt.Println(string(data))
	defer func() {
		fmt.Println("Ejecución finalizada")
	}()
	if data == nil {

		panic("el archivo indicado no fue encontrado o está dañado")
	}
	fmt.Println(string(data))

}
