package main

import (
	"errors"
	"fmt"
)

func sumValues(numbers []int32) (int32, error) {
	var result int32
	for num := range numbers {

		if numbers[num] < 0 {
			//fmt.Println(numbers[num] < 0, numbers[num])
			return numbers[num], errors.New("El valor ingresado no es entero positivo")
		}
		result += numbers[num]
	}
	return result, nil
}

func averageCalc(values ...int32) (int32, error) {
	total, errorR := sumValues(values)
	if errorR != nil {
		return total, errorR
	}
	return total / int32(len(values)), nil
}
func main() {
	av1, err1 := averageCalc(12, 12, 12, 12)
	if err1 != nil {
		panic(err1)
	}
	fmt.Println(av1)
	av2, err2 := averageCalc(10, 5, 10, 5)
	if err2 != nil {
		panic(err1)
	}
	fmt.Println(av2)
	av3, err3 := averageCalc(10, 5, -10, -5)
	if err3 != nil {
		panic(err3)
	}
	fmt.Println(av3)
}
