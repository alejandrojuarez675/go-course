package maths

import "testing"

func TestSumShouldBeOk(t *testing.T) {
	result := Sum(5, 5)

	if result != 10 {
		t.Errorf("Expected 10, but have %d", result)
	}
}

func TestDivideByZeroShouldReturnError(t *testing.T) {
	result, err := Divide(32, 0)

	if result != 0 || err == nil {
		t.Errorf("Expected error")
	}
}

func TestDivideShouldBeOk(t *testing.T) {
	result, err := Divide(32, 4)

	if result != 8.0 || err != nil {
		t.Errorf("Expected 8.0, but have %f", result)
	}
}
