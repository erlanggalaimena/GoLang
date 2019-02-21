package main

import "fmt"

//TASK
//Write a program that :
//assigns an int to a variable
//prints that int in decimal, binary, and hex
//shifts the bits of that int over 1 position to the left, and assigns that to a variable
//prints that variable in decimal, binary, and hex

func main() {
	x := 18
	fmt.Println(x)
	fmt.Printf("Decimal : %d\n", x)
	fmt.Printf("Binary : %b\n", x)
	fmt.Printf("Hexadecimal : %#x\n", x)
	fmt.Println("------------------------")

	y := x << 1
	fmt.Println(y)
	fmt.Printf("Decimal : %d\n", y)
	fmt.Printf("Binary : %b\n", y)
	fmt.Printf("Hexadecimal : %#x\n", y)
}
