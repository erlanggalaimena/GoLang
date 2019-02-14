package main

import "fmt"

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
}
