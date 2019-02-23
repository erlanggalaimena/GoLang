package main

import "fmt"

func main() {
	//Task
	//Building on the previous hands-on exercise, create a program that uses “else if” and “else”.
	x := 24
	if x > 30 {
		fmt.Println("x lebih besar dari 30")
	} else if x > 25 {
		fmt.Println("x lebih besar dari 25")
	} else {
		fmt.Println(x)
	}
}
