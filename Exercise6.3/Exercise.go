package main

import "fmt"

/*
    TASK
	Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
*/

func main() {
	first := first()
	later := later()

	defer fmt.Println(first)
	fmt.Println(later)
}

func first() string {
	return "called first"
}

func later() string {
	return "called later"
}
