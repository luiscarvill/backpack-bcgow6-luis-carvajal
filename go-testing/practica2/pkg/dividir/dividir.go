package dividir

import "errors"

func Dividir(num, den int) (int, error) {
	if den == 0 {
		return 0, errors.New("No es posible dividir por 0")
	}
	return num / den, nil
}
