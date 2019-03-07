package main

import "fmt"

//TASK
//Create a new type: vehicle.
//The underlying type is a struct.
//The fields:
//doors
//color

type vehicle struct {
	doors int
	color string
}

//Create two new types: truck & sedan.
//The underlying type of each of these new types is a struct.
//Embed the “vehicle” type in both truck & sedan.
//Give truck the field “fourWheel” which will be set to bool.
//Give sedan the field “luxury” which will be set to bool. solution

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	//Using the vehicle, truck, and sedan structs:
	//using a composite literal, create a value of type truck and assign values to the fields;
	//using a composite literal, create a value of type sedan and assign values to the fields.

	v1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "yellow",
		},
		fourWheel: true,
	}

	v2 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: true,
	}

	//Print out each of these values.
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println("------------------------")

	//Print out a single field from each of these values.
	fmt.Println("truck")
	fmt.Printf("\t%d\n", v1.doors)
	fmt.Printf("\t%s\n", v1.color)
	fmt.Printf("\t%t\n", v1.fourWheel)

	fmt.Println("sedan")
	fmt.Printf("\t%d\n", v2.doors)
	fmt.Printf("\t%s\n", v2.color)
	fmt.Printf("\t%t\n", v2.luxury)
}
