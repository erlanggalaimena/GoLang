package main

import "fmt"

//Task 1
//Create your own type. Have the underlying type be an int.
type erlangga int

//Task 2
//Create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword
var x erlangga

func main() {
	//Task 3
	//in func main
	//print out the value of the variable “x”
	fmt.Println(x)
	//print out the type of the variable “x”
	fmt.Printf("%T\n", x)
	//assign 42 to the VARIABLE “x” using the “=” OPERATOR
	x = 42
	//print out the value of the variable “x”
	fmt.Println(x)
}
