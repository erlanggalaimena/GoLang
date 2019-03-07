package main

import "fmt"

func main() {
	//TASK
	//Create and use an anonymous struct.
	x := struct {
		firstName string
		lastName  string
		age       int
	}{
		firstName: "Erlangga",
		lastName:  "Laimena",
		age:       23,
	}

	fmt.Println(x)
}
