package main

import "fmt"

func main() {
	//Task
	//Create a program that uses a switch statement with the switch expression
	//specified as a variable of TYPE string with the IDENTIFIER “favSport”.
	favSport := "football"
	switch favSport {
	case "basketball":
		fmt.Println("Group 1")
	case "football":
		fmt.Println("Group 2")
	case "volley ball":
		fmt.Println("Group 3")
	}
}
