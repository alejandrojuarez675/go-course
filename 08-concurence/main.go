package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	serverToCheck := []string{
		"https://beta.biolibre.ar",
		"https://beta.biolibre.br",
		"https://beta.biolibre.co",
		"https://alpha.biolibre.ar",
		"https://alpha.biolibre.br",
		"https://alpha.biolibre.co",
		"https://omega.biolibre.ar",
		"https://omega.biolibre.br",
		"https://omega.biolibre.co",
		"https://gamma.biolibre.ar",
		"https://gamma.biolibre.br",
		"https://gamma.biolibre.co",
		"https://biolibre.ar",
		"https://biolibre.br",
		"https://biolibre.co",
	}

	checkServersSync(serverToCheck)
	checkServersAsync(serverToCheck)

	// higher order function
	checkServerSyncInline := func(data []string) {
		for _, url := range data {
			checkServer(url)
		}
	}
	timeChecker("sync check server", serverToCheck, checkServerSyncInline)
}

func timeChecker(title string, data []string, callback func(data []string)) {
	fmt.Printf("------- Init %s -------\n", title)
	initTime := time.Now()

	callback(data)

	diffTime := time.Since(initTime)
	fmt.Println(title, "takes", diffTime)
}

func checkServersSync(serverToCheck []string) {
	fmt.Println("------- Init sync check server -------")
	initTime := time.Now()

	for _, url := range serverToCheck {
		checkServer(url)
	}

	diffTime := time.Since(initTime)
	fmt.Println("Sync method take", diffTime)
}

func checkServer(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "is not online")
	} else {
		fmt.Println(url, "is online")
	}
}

func checkServersAsync(serverToCheck []string) {
	fmt.Println("------- Init async check server -------")
	initTime := time.Now()

	channel := make(chan string)
	for _, url := range serverToCheck {
		go checkServerAsync(url, channel)
	}

	for i := 0; i < len(serverToCheck); i++ {
		fmt.Println(<-channel)
	}

	diffTime := time.Since(initTime)
	fmt.Println("Async method take", diffTime)
}

func checkServerAsync(url string, channel chan string) {
	_, err := http.Get(url)
	if err != nil {
		channel <- url + "is not online"
	} else {
		channel <- url + "is online"
	}
}
