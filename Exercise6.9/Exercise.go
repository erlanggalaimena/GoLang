package main

import "fmt"

/*
	TASK
	A “callback” is when we pass a func into a func as an argument.
	For this exercise, pass a func into a func as an argument
*/

func main() {
	arrNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sumNumbers := foo(sum, arrNumbers)
	temp := sum

	fmt.Println(sumNumbers)
	fmt.Printf("The type of 'sum' is %T\n", temp)
}

func sum(numbers ...int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}

	return total
}

func foo(callback func(numbers ...int) int, input []int) int {
	return callback(input...)
}
