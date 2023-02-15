package main

import "fmt"

func main() {
	var name1 = "Ale1" // declare dynamic var and assign value

	var name2 string // declare var
	name2 = "Ale2"   // assign value

	name3 := "Ale3" // declare dynamic var and assign value

	fmt.Println(name1)
	fmt.Println(name2)
	fmt.Println(name3)

	// initializing multiple variables at same line
	var name, lastname = "ale", "juarez"
	fmt.Println(name, lastname)

	name4, lastname4 := "ale", "juarez"
	fmt.Println(name4, lastname4)

	// constants
	const Course = "Go course"
	fmt.Println(Course)
}
