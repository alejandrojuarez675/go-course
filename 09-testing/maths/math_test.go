package maths

import "testing"

func TestSumShouldBeOk(t *testing.T) {
	testingData := []struct {
		a, b           int
		expectedResult int
	}{
		{1, 2, 3},
		{12, 2, 14},
		{412, 12, 424},
		{123124, 12323, 135447},
		{123, 2, 125},
	}

	for _, data := range testingData {
		result := Sum(data.a, data.b)

		if result != data.expectedResult {
			t.Errorf("Expected %d, have %d", data.expectedResult, result)
		}
	}
}

func TestDivideByZeroShouldReturnError(t *testing.T) {
	result, err := Divide(32, 0)

	if result != 0 || err == nil {
		t.Errorf("Expected error")
	}
}

func TestDivideShouldBeOk(t *testing.T) {
	testingData := []struct {
		a, b           int
		expectedResult float32
	}{
		{1, 2, 0.5},
		{12, 2, 6},
		{412, 412, 1.0},
		{123124, 1231240, 0.1},
		{123, 2, 61.5},
	}

	for _, data := range testingData {
		result, err := Divide(data.a, data.b)

		if result != data.expectedResult || err != nil {
			t.Errorf("Expected %f, but have %f", data.expectedResult, result)
		}
	}

}
