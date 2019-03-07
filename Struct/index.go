package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	p1 := person{
		firstName: "Erlangga",
		lastName:  "laimena",
		age:       23,
	}

	p2 := person{
		firstName: "irvi",
		lastName:  "ramadhani",
		age:       24,
	}

	sa1 := secretAgent{
		person: p1,
		ltk:    true,
	}

	sa2 := secretAgent{
		person: p2,
		ltk:    false,
	}

	fmt.Println(sa1.firstName, sa1.lastName, sa1.age, sa1.ltk)
	fmt.Println(sa2.firstName, sa2.lastName, sa2.age, sa2.ltk)
	fmt.Println("------------------------")

	//Anonymous struck
	pa := struct {
		gender    string
		firstName string
		lastName  string
	}{
		gender:    "male",
		firstName: "Erlangga",
		lastName:  "Laimena",
	}

	fmt.Println(pa.gender, pa.firstName, pa.lastName)
}
