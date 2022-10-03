package main

import (
	"errors"
	"fmt"
)

const (
	catA  float64 = 3_000
	percA float64 = 0.5
	catB  float64 = 1_500
	percB float64 = 0.2
	catC  float64 = 1_000
	hour  float64 = 60
)

func convMinToHours(value float64) float64 {
	return value / hour
}

func calcSalary(minutes float64, category string) (float64, error) {
	if minutes < 0 {
		return -1, errors.New("Minutos debe ser un valor positivo")
	}
	if category != "A" && category != "B" && category != "C" {
		return -1, errors.New("Categoria no permitida (A, B, C)")
	}
	hours := convMinToHours(minutes)

	switch category {
	case "A":
		return (catA * hours) + (catA*hours)*percA, nil
	case "B":
		return (catB * hours) + (catB*hours)*percB, nil
	default:
		return catC * hours, nil

	}
}

func main() {
	categA, err := calcSalary(120, "c")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(categA)
	}

}
