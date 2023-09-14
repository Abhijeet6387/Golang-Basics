package main

import (
	"fmt"
	"math"
)

// Interfaces are used to bind structs based on methods
// Here both square and circle structs had area() and circumf() methods, thus can be combined in an interface
type shape interface {
	area() float64
	circumf() float64
}

// struct sqaure
type square struct {
	length float64
}

// struct circle
type circle struct {
	radius float64
}

// for square - methods
func (s square) area() float64 {
	return s.length * s.length
}
func (s square) circumf() float64 {
	return 4 * s.length
}

// for circle - methods
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

// for printing info
func printInfo(sh shape) {
	fmt.Printf("Area of %T is : %0.2f \n", sh, sh.area())
	fmt.Printf("Circumference of %T is : %0.2f \n", sh, sh.circumf())
}