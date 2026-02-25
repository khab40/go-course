package main

import "fmt"

type square struct {
	sideLength float64
}

type triangle struct {
	base   float64
	height float64
}

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func main() {
	t := triangle{2, 2}
	s := square{2}

	printArea(t)
	printArea(s)
}
