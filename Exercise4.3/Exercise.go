package main

import "fmt"

func main() {
	//TASK
	// use SLICING to create the following new slices which are then printed:
	//[2 3 4 5 6]
	//[7 8 9 10 11]
	//[4 5 6 7 8]
	//[3 4 5 6 7]
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	fmt.Println(x)
	fmt.Println(x[1:6])
	fmt.Println(x[6:11])
	fmt.Println(x[3:8])
	fmt.Println(x[2:7])
}
