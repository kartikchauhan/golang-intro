package main

import (
	"fmt"
	"strconv"
)

var m int = 25 // only full declaration allowed for package level variables

var (
	doctorName string = "Deepti" // variables can be grouped together with a single declaration
	doctorCity string = "Prayagraj"
	doctorAge  int    = 27
)

// Three levels of visibility of variables:
var p int = 35 // 1. lowercase variables at package level are not exposed to the outside world. Any file in the same can access the variable

var P int = 35 // 2. UPPERCASE variables at package level are exposed outside the package

func main() {
	var i int // 3. Block scope variable
	i = 10

	var j float32 = 15

	k := 20.

	fmt.Printf("i = %v %T \n", i, i)
	fmt.Printf("j = %v %T \n", j, j) // %v -> value, %T -> Type
	fmt.Printf("k = %v %T \n", k, k)

	/////////////////////////////////////////////////////////////////////

	fmt.Printf("m = %v %T \n", m, m)

	var m int = 30 // variable declaration is allowed when its scope is different. This is called shadowing.

	// m := 35 // variable redeclaration not allowed in the same scope

	fmt.Printf("m = %v %T \n", m, m)

	//////////////////////////////////////////////////////////////////////

	var URL string = "www.google.com" // naming convention - acronyms should be written in UPPERCASE. Don't use '-' or '_' in variable names.

	fmt.Printf("URL = %v %T \n", URL, URL)

	//////////////////////////////////////////////////////////////////////

	var q int = 40
	fmt.Printf("q = %v %T \n", q, q)

	var r float32
	r = float32(q)

	fmt.Printf("r = %v %T \n", r, r)

	///////////////////////////////////////////////////////////////////////

	// int to string conversio

	var a int = 42

	fmt.Printf("a = %v %T \n", a, a)

	var b string
	b = string(a) // string is a set of bytes. So ASCII value 42 points to character '*'.

	fmt.Printf("b = %v %T \n", b, b)

	b = strconv.Itoa(a)

	fmt.Printf("b = %v %T \n", b, b)

}
