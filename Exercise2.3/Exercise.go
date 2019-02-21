package main

import "fmt"

//TASK
//Create TYPED and UNTYPED constants. Print the values of the constants.

//Untyped
const (
	a = 24
	b = 24
)

//Typed
const (
	c string  = "Erlangga"
	d float32 = 2.17
)

func main() {
	fmt.Println(a, b)
	fmt.Println(c, d)
}
