package main

import "fmt"

func main() {
	//TASK
	//Print every rune code point of the uppercase alphabet three times. Your output should look like this:
	//65
	//U+0041 'A'
	//U+0041 'A'
	//U+0041 'A'
	//66
	//U+0042 'B'
	//U+0042 'B'
	//U+0042 'B'
	//… through the rest of the alphabet characters

	for i := 65; i <= 122; i++ {
		if (i < 91) || (i > 96) {
			fmt.Println(i)
			for j := 0; j < 3; j++ {
				fmt.Printf("\t%#U\n", i)
			}
		}
	}
}
