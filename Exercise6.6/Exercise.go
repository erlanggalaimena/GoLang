package main

import "fmt"

/*
	TASK
	Build and use an anonymous func
*/

func main() {
	func(input string) {
		fmt.Println("Some input", input)
	}("Erlangga")
}
