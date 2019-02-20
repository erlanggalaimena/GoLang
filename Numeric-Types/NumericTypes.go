package main

import "fmt"

var a int
var b float64

//int and uint type
var c int8 = -128
var d uint8 = 255

func main() {
	//short declaration
	x := 42
	y := 42.34534

	fmt.Println(x)
	fmt.Println(y)

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	//variable declaration
	a = 24
	b = 24.25435

	fmt.Println(a)
	fmt.Println(b)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	//difference between int and uint
	//int contains positive number and negative
	//uint just contain positive number
	fmt.Println(c)
	fmt.Println(d)

	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)
}
