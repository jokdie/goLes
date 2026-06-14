package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (r Circle) Area() float64 {
	return math.Pi * r.Radius * r.Radius
}

func PrintArea(s Shape) {
	fmt.Printf("%.2f\n", s.Area())
}

func main() {
	r := Rectangle{
		Width:  10,
		Height: 5,
	}

	c := Circle{
		Radius: 3,
	}

	PrintArea(r)
	PrintArea(c)
}
