package main

import "fmt"

func main() {
	//returning func
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

	//callback
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(numbers)
	fmt.Println("Sum of even numbers:", sumEven(sum(), numbers...))
	fmt.Println("Sum of odd numbers:", sumOdd(sum(), numbers...))
	fmt.Println("------------------------")
	fmt.Println()

	//closure

	//variable with global scope
	fmt.Println(a)
	plusOne()
	fmt.Println(a)
	plusOne()
	fmt.Println(a)
	fmt.Println()

	//variable with local scope
	{
		b := 24
		fmt.Println(b)
		fmt.Println()
	}
	//this will cause error because b is variable in scope {} above
	// fmt.Println(b)

	c := incrementor()
	fmt.Printf("%T\n", c)
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println()

	d := incrementor()()
	fmt.Println(d)
	fmt.Println(d)
	fmt.Println(d)
	fmt.Println("-----------------------")
	fmt.Println()
	//the different between variable c and d is

	//c is variable with value type "func() int" wich has var x in its local scope incrementor
	//Value of x always increment because each time the program calling "c()",
	//it just calling the function inside function incrementor

	//in other hand, d is variable with value type "int" and the value always the same
	//It`s because d always call method incrementor that makes value of x always reset
	//because incrementor evaluate from first line method until end line of method incrementor

	//recursion
	fmt.Print(factorial(4))
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

//callback
func sumEven(f func(numbers ...int) int, input ...int) int {
	var temp []int
	for _, v := range input {
		if v%2 == 0 {
			temp = append(temp, v)
		}
	}

	return f(temp...)
}

func sumOdd(f func(numbers ...int) int, input ...int) int {
	var arrOdd []int
	for _, v := range input {
		if v%2 != 0 {
			arrOdd = append(arrOdd, v)
		}
	}

	return f(arrOdd...)
}

//closure
var a int

func plusOne() {
	a++
}

func incrementor() func() int {
	var result int

	return func() int {
		result++
		return result
	}
}

//recursion
func factorial(input int) int {
	if input == 1 {
		return 1
	}

	return input * factorial(input-1)
}
