package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type secretAgent struct {
	person
	active bool
}

func (p person) getFullName() {
	fmt.Println("My name is", p.firstName, p.lastName)
}

func main() {
	p1 := secretAgent{
		person: person{
			firstName: "James",
			lastName:  "Bond",
		},
		active: true,
	}

	p2 := person{
		firstName: "Cloud",
		lastName:  "Strife",
	}

	fmt.Println(p1)
	p1.getFullName()
	fmt.Println("------------------------")

	fmt.Println(p2)
	p2.getFullName()
}
