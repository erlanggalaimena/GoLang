package main

import "fmt"

const (
	_ = iota //'_' -> blank identifier
	b = 1 << (iota * 10)
	c = 1 << (iota * 10)
	d = 1 << (iota * 10)
)

func main() {
	x := 2
	fmt.Printf("%d\t\t%b\n", x, x)

	y := 7
	fmt.Printf("%d\t\t%b\n", y, y)

	//shifting one
	z := y << 1
	fmt.Printf("%d\t\t%b\n", z, z)
	z = y >> 1
	fmt.Printf("%d\t\t%b\n", z, z)
	fmt.Println("------------------------")
	fmt.Println()

	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
	fmt.Println("------------------------")
	fmt.Println()

	fmt.Printf("%d\t\t\t%b\n", b, b)
	fmt.Printf("%d\t\t\t%b\n", c, c)
	fmt.Printf("%d\t\t%b\n", d, d)
}
