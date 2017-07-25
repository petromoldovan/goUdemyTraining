package main

import "fmt"

func main() {
	a := 43
	var b *int = &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 11
	fmt.Println(a)
}
