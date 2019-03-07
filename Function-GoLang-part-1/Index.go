package main

import "fmt"

func main() {
	foo()
	bar("Erlangga")

	x := woo("Laimena")
	fmt.Println(x)

	witcher, status := isWitcher("Geralt")
	witcher2, status2 := isWitcher("Yennefer")

	fmt.Println(witcher, status)
	fmt.Println(witcher2, status2)
	fmt.Println("------------------------")

	sum := sumNumbers(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(sum)
	fmt.Println("------------------------")

	//unfurling slice
	xi := []int{1, 3, 5, 7, 9}
	sum2 := sumNumbers(xi...)
	fmt.Println(sum2)
	fmt.Println("------------------------")

	//passing zero paramater to function that has parameter
	fmt.Println(sumNumbers())
	fmt.Println("------------------------")

	//another way to pass parameter into variadic function
	sum3 := sumNumbers(2, 4, 6, 8)
	fmt.Println(sum3)
	fmt.Println("------------------------")

	//Anonymous function
	func() {
		fmt.Println("Called from Anonymous Function")
	}()

	func(input int) {
		fmt.Println(input, "was called from Anonymous Function")
	}(10)
	fmt.Println("------------------------")

	//func expression
	a := func() {
		fmt.Println("This value is generate by function expression")
	}

	b := func(input int) {
		fmt.Println(input, "was called from function expression")
	}

	a()
	b(24)
}

func foo() {
	fmt.Println("Hello from foo")
}

func bar(s string) {
	fmt.Println("Hello,", s)
}

//function with return type
func woo(s string) string {
	return s
}

//function with multiple return value
func isWitcher(name string) (string, bool) {
	result1 := name
	result2 := false
	if name == "Geralt" {
		result2 = true
	}

	return result1, result2
}

//function with "variadic parameter"
func sumNumbers(input ...int) int {
	result := 0
	for _, v := range input {
		result += v
	}

	return result
}
