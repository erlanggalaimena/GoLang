package main

import "fmt"

//TASK
//Create your own type “person” which will have an underlying type of
//“struct” so that it can store the following data:
//first name
//last name
//favorite ice cream flavors

type person struct {
	firstName        string
	lastName         string
	favoriteIceCream []string
}

func main() {
	//Create two VALUES of TYPE person. Print out the values, ranging over the
	//elements in the slice which stores the favorite flavors.
	p1 := person{
		firstName:        "Erlangga",
		lastName:         "Laimena",
		favoriteIceCream: []string{"chocolate", "Vanilla"},
	}

	p2 := person{
		firstName:        "Irvi",
		lastName:         "Ramadhani",
		favoriteIceCream: []string{"chocolate", "Vanilla", "Gellato"},
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println("------------------------")

	fmt.Printf("%s %s really loved ice cream with flavors:\n", p1.firstName, p1.lastName)
	for _, v := range p1.favoriteIceCream {
		fmt.Println(v)
	}

	fmt.Printf("%s %s really loved ice cream with flavors:\n", p2.firstName, p2.lastName)
	for _, v := range p2.favoriteIceCream {
		fmt.Println(v)
	}
}
