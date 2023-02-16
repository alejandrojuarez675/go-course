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
}

func log(format string, values ...any) {
	var now = time.Now().Format(time.RFC822)
	msg := "[main][" + now + "]: " + format + "\n"
	fmt.Printf(msg, values...)
}
