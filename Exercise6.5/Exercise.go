package main

import (
	"fmt"
	"math"
)

/*
	create a type SQUARE
	create a type CIRCLE

	attach a method to each that calculates AREA and returns it
		circle area= π r 2
		square area = L * W

	create a type SHAPE that defines an interface as anything that has the AREA method

	create a value of type square
	create a value of type circle

	use func info to print the area of square
	use func info to print the area of circle

*/

func main() {
	object1 := square{
		length: 4,
		height: 4,
	}

	object2 := circle{
		radius: 49,
	}

	info(object1)
	info(object2)
}

type shape interface {
	area() float64
}

type square struct {
	length, height float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	result := s.length * s.height
	return result
}

func (c circle) area() float64 {
	result := math.Pi * math.Pow(c.radius, 2)
	return result
}

func info(sh shape) {
	fmt.Println(sh.area())
}
