package main

import (
	"fmt"
	"time"
)

func main() {
	log("hello world")
	log("new log")
}

func log(s string) {
	var now = time.Now().Format(time.RFC822)
	fmt.Print("[main][", now, "]: ")
	fmt.Println(s)
}
