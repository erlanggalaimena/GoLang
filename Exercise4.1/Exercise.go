package main

import "fmt"

func main() {
	//TASK
	//Using a COMPOSITE LITERAL:
	//create an ARRAY which holds 5 VALUES of TYPE int
	//assign VALUES to each index position.
	x := [5]int{1, 2, 3, 4, 5}

	//Range over the array and print the values out.
	//Using format printing
	//print out the TYPE of the array
	for i, v := range x {
		fmt.Printf("%d => %d\n", i, v)
	}
	fmt.Printf("%T\n", x)
}
