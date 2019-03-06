package main

import "fmt"

func main() {
	//TASK
	//Create a slice of a slice of string ([][]string). Store the following
	//data in the multi-dimensional slice:
	//"James", "Bond", "Shaken, not stirred"
	//"Miss", "Moneypenny", "Helloooooo, James."

	y := []string{"James", "Bond", "Shaken, not stirred"}
	z := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	fmt.Println(y)
	fmt.Println(z)

	x := [][]string{y, z}
	fmt.Println(x)

	//Range over the records, then range over the data in each record.
	for i, valueX := range x {
		fmt.Println("Record : ", i)

		for _, value := range valueX {
			fmt.Print(value, " ")
		}
		fmt.Println()
	}
}
