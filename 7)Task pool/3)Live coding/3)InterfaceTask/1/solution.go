package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}
type Rectangle struct {
	Width  float64
	Height float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	circle := Circle{10}
	rectangle := Rectangle{6, 8}

	var shape1 Shape = &circle
	var shape2 Shape = &rectangle

	fmt.Println(shape1.Area())
	fmt.Println(shape2.Area())
}
