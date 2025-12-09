package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}
type rect2 struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect2) area() float64 {
	return r.width * r.height
}
func (r rect2) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func detectType(g geometry) {
	switch v := g.(type) {
	case rect2:
		fmt.Println("This is a rectangle", v)
	case circle:
		fmt.Println("This is a circle", v)
	default:
		fmt.Println("Unknown geometry type")
	}
}

func interfaceTest() {
	// r := rect2{width: 3, height: 4}
	c := circle{radius: 5}

	// measure(r)
	// measure(c)

	// detectType(r)
	// detectType(c)

	// Check if rect2 implements geometry correctly
	var _ geometry = rect2{} // Compile-time check - will fail if rect2 doesn't implement geometry

	// // Runtime check - assign to interface and verify
	var g geometry = c
	fmt.Printf("circle implements geometry: %T implements all methods\n", g)

	// You can also use type assertion on the interface variable
	if x, ok := g.(rect2); ok {
		fmt.Printf("Successfully asserted back to rect2: %+v\n", x)
	} else {
		fmt.Printf("Not the right type.")
	}
}

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

func (ss ServerState) String() string {
	var stateName = map[ServerState]string{
		StateIdle:      "idle",
		StateConnected: "connected",
		StateError:     "error",
		StateRetrying:  "retrying",
	}
	return stateName[ss]
}

func enumTest() {
	fmt.Println(StateIdle)
}
