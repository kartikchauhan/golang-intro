package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {
	irrfan := Person{
		firstName: "Irrfan",
		lastName:  "Khan",
		contactInfo: contactInfo{
			email: "irrfan_khan@gmail.com",
			zip:   110011,
		},
	}

	irrfanPointer := &irrfan
	irrfanPointer.updateFirstName("Sir Irrfan")
	irrfan.print()
}

func (p *Person) updateFirstName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p Person) print() {
	fmt.Printf("%+v\n", p)
}
