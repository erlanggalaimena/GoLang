package main

import "fmt"

/*
	Create a func which returns a func
	assign the returned func to a variable
	call the returned func

*/

func main() {
	x := sum()
	fmt.Printf("%T\n", x)
	fmt.Println(x(1, 2, 3, 4, 5, 6, 7, 8, 9))
}

func sum() func(numbers ...int) int {
	return func(input ...int) int {
		total := 0
		for _, v := range input {
			total += v
		}
		return total
	}
}
