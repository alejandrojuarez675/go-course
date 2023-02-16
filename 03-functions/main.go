package main

import (
	"fmt"
	"time"
)

func main() {
	log("hello world")
	log("new log")

	isTurnedOn := true
	log("formatted log is %v", isTurnedOn)

	// lambdas func
	isBefore20230304 := func() bool {
		now := time.Now()
		return now.Before(time.Date(2023, time.March, 4, 0, 0, 0, 0, time.Local))
	}
	log("is before 4/3/23 %v", isBefore20230304())

	// reference passage value on functions
	value1 := "car"
	log("value: %v, memory space: %v", value1, &value1)
	value1 = modifyValueWithValuePassagedParam(value1)
	log("value: %v, memory space: %v", value1, &value1)

	value2 := "duck"
	log("value: %v, memory space: %v", value2, &value2)
	modifyValueWithReferencePassagedParam(&value2)
	log("value: %v, memory space: %v", value2, &value2)

}

func modifyValueWithValuePassagedParam(s string) string {
	return "modify very large to force change of memory space. "
}

func modifyValueWithReferencePassagedParam(s *string) {
	*s = "modify"
}

func log(format string, values ...any) {
	var now = time.Now().Format(time.RFC822)
	msg := "[main][" + now + "]: " + format + "\n"
	fmt.Printf(msg, values...)
}
