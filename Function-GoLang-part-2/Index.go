package main

import "fmt"

func main() {
	x := foo()()
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println()

	y := bar()(24)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println()

	tmpSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	z := sum()(tmpSlice...)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println("------------------------")
	fmt.Println()
}

//returning func
func foo() func() string {
	return func() string {
		return "My name is Erlangga Laimena"
	}
}

func bar() func(input int) int {
	return func(input int) int {
		return input
	}
}

func sum() func(input ...int) int {
	return func(x ...int) int {
		result := 0
		for _, v := range x {
			result += v
		}
		return result
	}
}
