package main

import "fmt"

func main() {
	a := []int{1, 2, 3} // empty [] means slice

	b := a // slice are naturally of reference type
	b[1] = 5

	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%v \n", len(a))
	fmt.Printf("%v \n", cap(a)) // capacity

	////////////////////////////////////////////////////////////

	c := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	d := c[:]
	e := c[3:] // index 3 is inclusive
	f := c[:6] // index 6 is exclusive
	g := c[3:6]
	g[1] = 100

	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)

	////////////////////////////////////////////////////////////////
 
}
