package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) Greeting() {
	fmt.Printf("Hi! my name is %s. I'm %d\n", p.name, p.age)
}

// Hierarch
type Employee struct {
	Person
	sector string
}

func (e *Employee) Work() {
	fmt.Printf("%s is working\n", e.name)
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

	e1 := Employee{
		sector: "TG",
	}
	e1.name = "raul"
	e1.age = 32
	fmt.Println("e1", e1)

	e1.Greeting()
	e1.Work()
}
