package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func main() {
	t := triangle{height: 5, base: 5}
	s := square{sideLength: 10}
	printArea(t)
	printArea(s)
}

func printArea(s shape) {
	fmt.Println("The area of the shape is", s.getArea())
}
