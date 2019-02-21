package main

import "fmt"

//TASK
//Using iota, create 4 constants for the NEXT 4 years. Print the constant values.
const (
	a = 2019 + iota
	b = 2019 + iota
	c = 2019 + iota
	d = 2019 + iota
)

func main() {
	fmt.Print(a, b, c, d)
}
