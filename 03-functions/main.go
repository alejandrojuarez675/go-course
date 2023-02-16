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
}

func log(format string, values ...any) {
	var now = time.Now().Format(time.RFC822)
	msg := "[main][" + now + "]: " + format + "\n"
	fmt.Printf(msg, values...)
}
