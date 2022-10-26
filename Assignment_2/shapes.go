package main

import "fmt"

type Shape interface {
	getArea() float64
}

type Square struct {
	sideLength float64
}

type Triangle struct {
	height float64
	base   float64
}

func (t Triangle) getArea() float64 {
	return (0.5 * t.height * t.base)
}

func (sq Square) getArea() float64 {
	return (sq.sideLength * sq.sideLength)
}

func printArea(s Shape) {
	fmt.Println(s.getArea())
}
