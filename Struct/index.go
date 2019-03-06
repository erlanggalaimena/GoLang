package main

import "fmt"

type person struct {
	first_name string
	last_name  string
	age        int
}

func main() {
	p1 := person{
		first_name: "Erlangga",
		last_name:  "laimena",
		age:        23,
	}

	p2 := person{
		first_name: "irvi",
		last_name:  "ramadhani",
		age:        24,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println("------------------------")

	fmt.Println(p1.first_name)
	fmt.Println(p2.last_name)
	fmt.Println(p1.age)
}
