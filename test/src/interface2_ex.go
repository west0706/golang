package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}
type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	r := Rect{10., 20.}
	r2 := Rect{10., 20.}
	c := Circle{10}

	showArea(r, r2, c)
}

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.area()
		// b := s.perimeter()
		fmt.Println(a)
	}
}
