package main

import (
	"fmt"
)

func main() {
	var (
		temperatura float32 = 25
		humedad     int     = 69
		presion     float32 = 1017
	)
	fmt.Printf("Temperatura: %v C \n Humedad: %v %%\n Presión: %v hPa \n", temperatura, humedad, presion)
}
