package main

import "fmt"

/*
	TASK
	Build and use an anonymous func
*/

func main() {
	x := func(input string) string {
		return fmt.Sprintf("My name is %s", input)
	}("Erlangga Laimena")

	y := func(numbers ...int) int {
		total := 0
		for _, v := range numbers {
			total += v
		}
		return total
	}

	arrNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(x)
	fmt.Println(y(arrNumbers...))
}
