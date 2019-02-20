package main

import "fmt"

//Task 1
//Building on the code from the previous example
//At the package level scope, using the “var” keyword, create a VARIABLE with the IDENTIFIER “y”.
//The variable should be of the UNDERLYING TYPE of your custom TYPE “x”
type erlangga int

var x erlangga
var y int

func main() {
	//Task 2
	//in func main
	//this should already be done
	//print out the value of the variable “x”
	//print out the type of the variable “x”
	//assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
	//print out the value of the variable “x”
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 24
	fmt.Println(x)

	//Task 3
	//now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
	//then use the “=” operator to ASSIGN that value to “y”
	//print out the value stored in “y”
	//print out the type of “y”
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
