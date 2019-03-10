package main

import "fmt"

/*
	Create a user defined struct with
	the identifier “person”
	the fields:
		first
		last
		age

	attach a method to type person with
	the identifier “speak”
	the method should have the person say their name and age

	create a value of type person

	call the method from the value of type person
*/

func main() {
	p1 := person{
		first: "Erlangga",
		last:  "Laimena",
		age:   23,
	}

	p2 := person{
		first: "Cloud",
		last:  "Strife",
		age:   25,
	}

	p1.speak()
	p2.speak()
}

type person struct {
	first, last string
	age         int
}

func (p person) speak() {
	fmt.Println("My name is", p.first, p.last)
	fmt.Println("My age is", p.age)
	fmt.Println()
}
