package main

import "fmt"

func main() {
	//TASK
	//start with this slice
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	//append to that slice 52
	//print out the slice
	x = append(x, 52)
	fmt.Println(x)

	//in ONE STATEMENT append to that slice these values
	//53
	//54
	//55
	//print out the slice
	x = append(x, 53, 54, 55)
	fmt.Println(x)

	//append to the slice this slice
	y := []int{56, 57, 58, 59, 60}

	z := append(x, y...)
	fmt.Println(z)
}
