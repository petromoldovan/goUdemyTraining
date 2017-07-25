package main

import "fmt"

func zero(x int) {
	x = 0
}

func main() {
	x := 22
	zero(x)
	fmt.Println(x)  //prints 22. The value of x is not changed because we make a copy of this var that we change inside zero func.
					//see example with-pointer where i pass memory address inside of zero func
}
