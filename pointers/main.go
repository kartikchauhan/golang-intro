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

	// irrfanPointer := &irrfan
	// irrfanPointer.updateFirstName("Sir Irrfan")

	// we can just pass the original identifier and go will deduce the *type* if the receiver is of type "pointer to *type*"
	irrfan.updateFirstName("Sir Irrfan")
	irrfan.print()
}

func (p *Person) updateFirstName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p Person) print() {
	fmt.Printf("%+v\n", p)
}
