package main

import "fmt"

func main() {
	//Task
	//Create a for loop using this syntax
	//for { }
	i := 1
	for {
		fmt.Println(i)

		if i > 7 {
			break
		}

		i++
	}
}
