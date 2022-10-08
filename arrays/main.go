package main

import "fmt"

func main() {
	grades := [...]int{97, 98, 99} // undefined size

	var students [3]string

	students[0] = "Raman"
	students[1] = "Raghav"
	students[2] = "Raman Raghav"

	fmt.Printf("%v\n", grades)

	fmt.Printf("%v\n", students)

	fmt.Printf("length %v\n", len(students))

	//////////////////////////////////////////////////////////

	var classes [3]int = [3]int{5, 6, 7}

	fmt.Printf("%v\n", classes)

	////////////////////////////////////////////////////////////

	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}

	var identityMatrix2 [3][3]int = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}

	fmt.Printf("identityMatrix %v\n", identityMatrix)

	fmt.Printf("identityMatrix2 %v\n", identityMatrix2)

	/////////////////////////////////////////////////////////////////

	a := [...]int{1, 2, 3}
	b := a

	b[1] = 5 // New array doesn't point to the original array's location but gets copied entirely.

	fmt.Printf("%v\n%v\n", a, b)

	c := &a
	c[1] = 5

	fmt.Println(a)
	fmt.Println(c)
}
