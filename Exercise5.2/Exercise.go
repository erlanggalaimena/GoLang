package main

import "fmt"

type person struct {
	firstName              string
	lastName               string
	favoriteIceCreamFlavor []string
}

func main() {
	//TASK
	//Take the code from the previous exercise, then store the values of type
	//person in a map with the key of last name. Access each value in the map.
	//Print out the values, ranging over the slice.

	p1 := person{
		firstName:              "Erlangga",
		lastName:               "Laimena",
		favoriteIceCreamFlavor: []string{"vanilla"},
	}

	p2 := person{
		firstName:              "Irvi",
		lastName:               "Ramadhani",
		favoriteIceCreamFlavor: []string{"vanilla", "chocolate"},
	}

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for k, v := range m {
		fmt.Printf("index %s:\n", k)
		fmt.Printf("\t%s\n", v.firstName)
		fmt.Printf("\t%s\n", v.lastName)

		for _, v2 := range v.favoriteIceCreamFlavor {
			fmt.Printf("\t%s\n", v2)
		}
	}
}
