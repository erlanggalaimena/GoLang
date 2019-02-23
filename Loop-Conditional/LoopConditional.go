package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
	fmt.Println("------------------------")

	//another implementation of "for"
	x := 0
	for x < 5 {
		fmt.Println(x)
		x++
	}
	fmt.Println("------------------------")

	x = 1
	for {
		fmt.Println(x)
		if x > 7 {
			break
		}
		x++
	}
	fmt.Println("------------------------")

	//implementation continue and break
	y := 0
	for {
		y++

		if y > 10 {
			break
		}

		if y%2 != 0 {
			continue
		}

		fmt.Println(y)
	}
}
