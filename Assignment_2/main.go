package main

func main() {
	s := Square{
		sideLength: 4,
	}

	t := Triangle{
		height: 2,
		base:   6,
	}

	printArea(s)
	printArea(t)
}
