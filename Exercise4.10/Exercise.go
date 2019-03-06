package main

import "fmt"

func main() {
	//TASK
	//Using the code from the previous example, delete a record from your map.
	//Now print the map out using the “range” loop
	x := map[string][]string{
		"bond_james":      {`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": {`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`},
		"no_dr":           {`no_dr`, `Being evil`, `Ice cream`, `Sunsets`},
	}

	x["Laimena_Erlangga"] = []string{"Laimena_Erlangga", "Spirit is everything", "Martabak", "Programmer"}

	delete(x, "no_dr")

	for k, v := range x {
		fmt.Println("The values of key", k)
		for i, v2 := range v {
			fmt.Printf("\t%d: %v\n", i, v2)
		}
	}
}
