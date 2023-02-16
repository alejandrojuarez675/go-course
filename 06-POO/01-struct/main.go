package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) Greeting() {
	fmt.Printf("Hi! my name is %s. I'm %d\n", p.name, p.age)
}

func main() {
	p1 := Person{"ale", 28}
	p1.name = "alejandro"
	fmt.Println("p1", p1)
	p1.Greeting()

	p2 := Person{
		name: "ale",
		age:  28,
	}
	fmt.Println("p2", p2)
	p2.Greeting()
}
