package log

import (
	"fmt"
	"time"
)

func log(s string) {
	var now = time.Now().Format(time.RFC822)
	fmt.Print("[main][", now, "]: ")
	fmt.Println(s)
}
