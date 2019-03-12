package main

import (
	"fmt"
	"strings"
)

/*
	TASK
	Closure is when we have “enclosed” the scope of a variable in some code block.
	For this hands-on exercise, create a func which “encloses” the scope of a variable:
*/

func main() {
	function1 := func(input string) string {
		if strings.EqualFold("Erlangga", input) {
			return "Erlangga Laimena"
		} else if strings.EqualFold("Irvi", input) {
			return "Irvi Ramadhani"
		} else {
			return "Not Identified"
		}
	}

	fmt.Println(function1("erlangga"))
	fmt.Println(function1("iRvI"))
	fmt.Println()

	x := increament()
	fmt.Println("The value of variable x:")
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println()

	y := increament()
	fmt.Println("The value of variable y:")
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())

}

func increament() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
