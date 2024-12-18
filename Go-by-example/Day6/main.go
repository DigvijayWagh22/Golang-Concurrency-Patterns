package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	peri() float64
}

type circle struct {
	radius float64
}

type rect struct {
	height, width float64
}

func (r rect) area() float64 {
	return r.height * r.width
}
func (r rect) peri() float64 {
	return 2 * (r.height + r.width)
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) peri() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("area ", g.area())
	fmt.Println("perimeter ", g.peri())
}

func main() {
	r := rect{5, 10}
	c := circle{10}
	measure(r)
	measure(c)
}
