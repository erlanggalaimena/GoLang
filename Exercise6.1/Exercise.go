package main

import "fmt"

//TASK
//create a func with the identifier foo that returns an int
//create a func with the identifier bar that returns an int and a string
//call both funcs
//print out their results

func main() {
	a := foo(24)
	b, c := bar(18, "erlangga")

	fmt.Println(a)
	fmt.Println(b, c)
}

func foo(number int) int {
	return number
}

func bar(number int, input string) (int, string) {
	return number, input
}
