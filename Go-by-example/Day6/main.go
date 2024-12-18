package main

import (
	"fmt"
	"math"
)

//Interfaces

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

//Enums

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "Error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("Unknown state: %s", s))
	}
}

func main() {
	// Interfaces
	r := rect{5, 10}
	c := circle{10}
	measure(r)
	measure(c)

	//Enums

	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)

}
