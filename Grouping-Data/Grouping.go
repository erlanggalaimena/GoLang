package main

import "fmt"

func main() {
	//Different between Array and Slice is that array has a fixed length and Slice has a dynamic length

	//Array
	var x [5]int
	fmt.Println(x)

	x[3] = 24
	fmt.Println(x)

	//print length of array
	fmt.Println(len(x))
	fmt.Println("------------------------")

	//SLICE allows you to group together values of same type
	//composite literals
	a := []int{1, 3, 5, 7, 9}

	//Print all value of variable "a" using "for range"
	fmt.Println("using 'for range'")
	for i, v := range a {
		fmt.Println(i, "=>", v)
	}

	//Print all value of variable "a" using for
	fmt.Println("using 'for'")
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], " ")
	}
	fmt.Println()
	fmt.Println("------------------------")

	//Slicing a SLICE
	b := a[:]
	fmt.Println(b)
	c := a[1:]
	fmt.Println(c)
	d := a[2:5]
	fmt.Println(d)

	//APPEND slice "a" with some values
	a = append(a, 11, 13, 15)
	fmt.Println(a)

	//APPEND two slices
	y := []int{1, 2, 3, 4, 5}
	z := []int{6, 7, 8, 9, 10}
	fmt.Println(append(y, z...))
	fmt.Println("------------------------")

	//Deleting from slice
	g := []int{234, 456, 678, 987}
	h := []int{125, 250, 500, 1000}
	i := append(g, h...)
	fmt.Println(i)
	i = append(i[:2], i[4:]...)
	fmt.Println(i)
	fmt.Println("------------------------")

	//Make Slice
	j := make([]int, 10, 12)
	fmt.Println(j)
	fmt.Println(len(j))
	fmt.Println(cap(j))
	fmt.Println()

	j[0] = 42
	j[9] = 999
	fmt.Println(j)
	fmt.Println(len(j))
	fmt.Println(cap(j))
	fmt.Println()

	j = append(j, 1000)
	fmt.Println(j)
	fmt.Println(len(j))
	fmt.Println(cap(j))
	fmt.Println()

	j = append(j, 1001)
	fmt.Println(j)
	fmt.Println(len(j))
	fmt.Println(cap(j))
	fmt.Println()

	j = append(j, 1002)
	fmt.Println(j)
	fmt.Println(len(j))
	fmt.Println(cap(j))
	fmt.Println("------------------------")

	//Multidimensional Slice
	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	mp := []string{"Miss", "Penny", "Strawberry", "Hazzelnut"}

	fmt.Println(jb)
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)
	fmt.Println("------------------------")

	//Map
	m := map[string]int{
		"James Bond": 32,
		"Miss Penny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["James Bond"])
	fmt.Println(m["james bond"])

	//'comma ok' idiom
	v, ok := m["james bond"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["James Bond"]; ok {
		fmt.Println(v)
	}

	//Add new element to map
	m["Erlangga"] = 24

	for k, v := range m {
		fmt.Printf("%s : %d\n", k, v)
	}
	fmt.Println()

	//deleting from map
	delete(m, "James Bond")
	for k, v := range m {
		fmt.Printf("%s : %d\n", k, v)
	}
}
