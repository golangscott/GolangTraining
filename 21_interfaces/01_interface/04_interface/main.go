package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	side float64
}

type trangle struct {
	a float64
	b float64
	c float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.side * s.side
}

func (t trangle) area() float64 {
	return (t.a + t.b + t.c)/2
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

// a new method which takes the INTERFACE TYPE shape
func totalArea(shapes ...shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func main() {
	s := square{10}
	c := circle{5}
	t := trangle{3,4,5}
	info(s)
	info(c)
	info(t)
	fmt.Println("Total Area: ", totalArea(c, s, t))
}
