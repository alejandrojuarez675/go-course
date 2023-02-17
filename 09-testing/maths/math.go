package maths

import "errors"

func Sum(a, b int) int {
	return a + b
}

func Divide(a, b int) (float32, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide for zero")
	}
	return float32(a) / float32(b), nil
}
