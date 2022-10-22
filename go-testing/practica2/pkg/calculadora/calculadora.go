package calculadora

import "errors"

// Restar : substract two numbers
func Restar(num1, num2 int) (int, error) {
	if num1 == 0 || num2 == 0 {
		return 0, errors.New("Los numeros a restar deben ser diferentes de 0")
	}
	return num1 - num2, nil
}
