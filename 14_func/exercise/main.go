package main

import "fmt"

func findMax(list ...int) int {
	fmt.Printf("%T \n", list)

	var max int
	for _, v := range list {
		if v > max {
			max = v
		}
	}

	return max
}

func main(){
	res := findMax(1,231,4,5,7,338,44)

	fmt.Println(res)
}
