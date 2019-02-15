package main

import "fmt"

//var declare at package level
var city = "Tangerang"

//assign var with data type
//variable example is assign with int data type
//because example is assign with just data type, the value of example is 0(default)
var example int

//example of variable assign with data type and value
var example2 int = 100

func main() {
	//First declaration of new variable must use ":=" operator
	x := 24
	fmt.Println(x)

	//Then you can change the value of variable x with operator "="
	//Because the variable x already define using ":=" operator
	x = 42
	fmt.Println(x)

	//In the operator ":=", ":" is meaning assign and "=" is meaning declare

	y := 100 + 24
	fmt.Println(y)

	z := "Erlangga Laimena"
	fmt.Println(z)

	//var can be declare both on package or function level
	//var declare at function level
	var country = "Indonesia"
	fmt.Println(country)

	//call var that declare in package level
	fmt.Println(city)
	fmt.Println(example)
	fmt.Println(example2)

	//print type of variable
	fmt.Printf("%s %T\n", "Data-Type of variable example is : ", example)
}
