package main

import "fmt"

//Task 1
//Using the code from the previous exercise,
//At the package level scope, assign the following values to the three variables
//for x assign 42
var x int = 42

//for y assign “James Bond”
var y string = "James Bond"

//for z assign true
var z bool = true

func main() {
	//Task 2
	//in func main
	//use fmt.Sprintf to print all of the VALUES to one single string.
	//ASSIGN the returned value of TYPE string using the short declaration
	//operator to a VARIABLE with the IDENTIFIER “s”
	//print out the value stored by variable “s”
	s := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println(s)
}
