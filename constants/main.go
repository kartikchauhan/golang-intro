package main

import "fmt"

const (
	a = iota // iota is a special counter we can use when we're creating enumerated constants. It's value always starts with 0.
	b
	c
)

const (
	a2 = iota
)

const (
	a3 = "a"
	b3
	c3
)

const (
	_ = iota + 5 // operations with primitive types are allowed during constant's declaration. Function expressions are not allowed as the value needs to determined during compile time.
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

func main() {
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", b, b)
	fmt.Printf("%v %T\n", c, c)
	fmt.Printf("%v %T\n", a2, a2)
	fmt.Printf("%v %T\n", a3, a3)
	fmt.Printf("%v %T\n", b3, b3)
	fmt.Printf("%v %T\n", c3, c3)

	var specialistType int

	fmt.Println("specialistType == catSpecialist", specialistType == catSpecialist)
}
