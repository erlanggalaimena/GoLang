package main

import "fmt"

func main() {
	fmt.Println("Normal Print")
	fmt.Println("Hello everyone")
	fmt.Println("-----------------")
	fmt.Println()

	fmt.Println("Method Call")
	foo()
	fmt.Println("-----------------")
	fmt.Println()

	fmt.Println("Print Iteration")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("------------------")
	fmt.Println()

	fmt.Println("Print Iteration With Just Even Number")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("I`m Called From Method Foo")
}
