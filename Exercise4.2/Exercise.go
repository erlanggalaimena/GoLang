package main

import "fmt"

func main() {
	//TASK
	//Using a COMPOSITE LITERAL:
	//create a SLICE of TYPE int
	//assign 10 VALUES
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//Range over the slice and print the values out.
	for i, v := range x {
		fmt.Printf("%d => %d\n", i, v)
	}

	//Using format printing
	//print out the TYPE of the slice
	fmt.Printf("%T\n", x)

}
