package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("dividing 4/0")
	result1, errDivision1 := divide(4, 0)
	if errDivision1 == nil {
		fmt.Println("4/0=", result1)
	} else {
		fmt.Println(errDivision1.Error())
	}

	fmt.Println("dividing 4/2")
	if result2, errDivision2 := divide(4, 2); errDivision2 == nil {
		fmt.Println("4/2=", result2)
	} else {
		fmt.Println(errDivision2.Error())
	}
}

func divide(i1, i2 int) (int, error) {
	if i2 == 0 {
		errorMsg := "It can't possible divide by zero -> " + strconv.Itoa(i1) + "/" + strconv.Itoa(i2)
		return 0, errors.New(errorMsg)
	}

	return i1 / i2, nil
}
