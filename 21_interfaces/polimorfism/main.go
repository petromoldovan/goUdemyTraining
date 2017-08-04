package main

import (
	"fmt"
	"math"
)

type Square struct {
	side 	float64
}

type Circle struct {
	radius float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

func (c Circle) area() float64  {
	return c.radius * c.radius * math.Pi
}

//every struct that has method area() will implement Shape interface from now on
type Shape interface {
	area() float64
}


func getInfo(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := &Square{2}
	c := &Circle{4}

	getInfo(s)
	getInfo(c)
}

