package main

import "fmt"

func zero(x *int) {
	*x = 0
}

func main() {
	x := 22
	zero(&x)
	fmt.Println(x)  //prints 0 because we are changing value assigned to certain memory address that we pass to zero func.
}
