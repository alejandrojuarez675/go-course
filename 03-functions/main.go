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
	fmt.Print("[main][", now, "]: ")
	fmt.Printf(format, values...)
	fmt.Print("\n")
}
