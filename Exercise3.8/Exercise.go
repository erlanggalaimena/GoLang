package main

import "fmt"

func main() {
	//Task
	//Create a program that uses a switch statement with no switch expression specified.
	x := 24
	switch {
	case x > 100:
		fmt.Println("Group A")
	case x > 75:
		fmt.Println("Group B")
	case x > 50:
		fmt.Println("Group C")
	case x > 25:
		fmt.Println("Group D")
	default:
		fmt.Println("Group E")
	}
}
