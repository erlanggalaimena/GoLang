package main

import "fmt"

func main() {
	a := "angga"
	b := "irvi"

	switch {
	case (a == b):
		fmt.Println("Satu Entitas")
	case (a != b):
		fmt.Println("Dua Entitas")
	}
	fmt.Println("------------------------")

	//"falltrough" on switch
	switch {
	case (a == b):
		fmt.Println("Satu Entitas")
	case (a != b):
		fmt.Println("Dua Entitas")
		fallthrough
	case ((1 + 1) == 2):
		fmt.Println("Penjumlahan yang benar")
		fallthrough
	case ("a" != "b"):
		fmt.Println("Huruf yang tidak sama")
		fallthrough
	case (2 == 3):
		fmt.Println("Pernyataan yang ngaco")
		fallthrough
	case (1 != 2):
		fmt.Println("Angka yang tidak sama")
	}
	fmt.Println("-------------------------")

	//default on switch
	switch {
	case (a == b):
		fmt.Println("Satu Entitas")
	default:
		fmt.Println("Default was called")
	}
	fmt.Println("-------------------------")

	//Switch on value
	x := 18
	switch x {
	case 21:
		fmt.Println("Lebih besar dari 20")
	case 20:
		fmt.Println("Sama dengan 20")
	default:
		fmt.Println("Lebih kecil dari 20")
	}
	fmt.Println("-------------------------")

	//Switch on multiple matches
	x = 22
	switch x {
	case 21, 22, 23:
		fmt.Println("Lebih besar dari 20")
	case 20:
		fmt.Println("Sama dengan 20")
	default:
		fmt.Println("Lebih kecil dari 20")
	}
}
