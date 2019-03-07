package main

import "fmt"

func main() {
	//defer make the function  deferred and call it in the end
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
