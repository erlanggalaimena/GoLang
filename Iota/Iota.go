package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
)

const (
	d = iota + 1
	e = iota
	f = iota
)

const (
	g = iota + 1
	h
	i
)

const (
	j = iota
	k
	l
	m = iota
	n
	o
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println("------------------------")

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println("------------------------")

	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println("------------------------")

	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(o)
}
