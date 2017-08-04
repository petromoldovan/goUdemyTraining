package main

import "fmt"

type Square struct {
	side 	float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

//every struct that has method area() will implement Shape interface from now on
type Shape interface {
	area() float64
}

//here we can pass any shape interface. See next polimorfism
func getInfo(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := &Square{2}

	getInfo(s)
}

