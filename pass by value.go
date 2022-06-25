package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func value(p Person) {}

func pointer(p *Person) {
	p.firstName = "KANHAIYA"
	p.lastName = "KUMAR"
}

func main() {
	person := Person{
		firstName: "Kanhaiya",
		lastName:  "Kumar",
	}
	person1 := Person{}

	fmt.Println(person)

	pointer(&person1)

	fmt.Println(person1)
}
