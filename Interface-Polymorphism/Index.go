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

type human interface {
	speak()
	getFullName() string
}

func (s secretAgent) speak() {
	fmt.Println("My Name is", s.firstName, s.lastName)
}

func (p person) speak() {
	fmt.Println("My Name is", p.firstName, p.lastName)
}

func (s secretAgent) getFullName() string {
	return fmt.Sprintf("%s %s", s.firstName, s.lastName)
}

func (p person) getFullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

func identification(h human) {
	switch h.(type) {
	case person:
		fmt.Println(h.(person).getFullName(), "is human species")
	case secretAgent:
		fmt.Println(h.(secretAgent).getFullName(), "is human species")
	}
}

func main() {
	sa1 := secretAgent{
		person: person{
			firstName: "James",
			lastName:  "Bond",
		},
		active: true,
	}

	sa2 := secretAgent{
		person: person{
			firstName: "Ethan",
			lastName:  "Hunt",
		},
		active: true,
	}

	p1 := person{
		firstName: "Erlangga",
		lastName:  "Laimena",
	}

	fmt.Println(sa1)
	fmt.Println(sa2)
	fmt.Println(p1)
	fmt.Println("------------------------")
	fmt.Println()

	sa1.speak()
	sa2.speak()
	p1.speak()
	fmt.Println("------------------------")
	fmt.Println()

	fmt.Println(sa1.getFullName())
	fmt.Println(sa2.getFullName())
	fmt.Println(p1.getFullName())
	fmt.Println("------------------------")
	fmt.Println()

	identification(sa1)
	identification(sa2)
	identification(p1)
}
