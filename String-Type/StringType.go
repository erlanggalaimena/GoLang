package main

import "fmt"

func main() {
	s := "Hello World"

	fmt.Println(s)
	fmt.Printf("%T\n", s)
	fmt.Println()

	for i := 0; i < len(s); i++ {
		if string(s[i]) != " " {
			fmt.Printf("%s ", string(s[i]))
		}
	}
}
