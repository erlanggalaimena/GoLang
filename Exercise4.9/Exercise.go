package main

import "fmt"

func main() {
	//TASK
	//Using the code from the previous example, add a record to your map. Now
	//print the map out using the “range” loop

	x := map[string][]string{
		"bond_james":      {`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": {`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`},
		"no_dr":           {`no_dr`, `Being evil`, `Ice cream`, `Sunsets`},
	}

	x["Laimena_Erlangga"] = []string{"Laimena_Erlangga", "Spirit is everything", "Martabak", "Programmer"}

	for k, v := range x {
		fmt.Println("Values of index", k)

		for i, v2 := range v {
			fmt.Printf("\t%d: %v\n", i, v2)
		}
	}

}
