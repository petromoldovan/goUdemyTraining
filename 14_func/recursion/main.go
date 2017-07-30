package main

import "fmt"

func factorial(x int) int {
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)    //returns 4 * 3 * 2 * 1 * 1
}

func main() {
	fmt.Println(factorial(4))
}

