package main

import (
	"fmt"
)

func main() {
	type person struct {
		ID 							int
		FirstName				string
		Surname 				string
	}

	p1 := person{1, "Sam", "Jackson"}

	fmt.Println(p1)
	fmt.Printf("%T \n", p1)
	fmt.Println(p1.Surname)
}
