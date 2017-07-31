package main

import "fmt"


func firstWayToInitializeSlice() {
	//when slice initialized in this way, the capacity equals to number of values automatically
	slice := []int{1,2,3,4,5,6,}

	fmt.Println("------------------")
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(slice)
	fmt.Println("------------------")
}

func seconWayToInitializeSlice() {
	//Try to always make slices with make
	//when slice initialized in this way, the capacity is the last param in make func.
	slice := make([]int, 2, 6)   // <= make slice of ints with length of 2 and capacity of 6

	slice[1] = 50 //this works because length of underlying array is 2, so it has index of 1
	//slice[2] = 122   <-- this will no work becase undelying array does not have index 2. USE APPEND to added elems to slice in this case!

	slice = append(slice, 123)

	fmt.Println("------------------")
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(slice)
	fmt.Println("------------------")
}

func main() {
	firstWayToInitializeSlice()
	seconWayToInitializeSlice()
}

//on capacity in general:
//when initial capacity is 3 and we append the 4th value to the slice, then go will increase the capacity of
//underlying array by 2 times to capacity of 6.


//NOTE:
//can also initialise with make like this:
//slice := make([]int, 4)   <-- make sice with length and capacity of 4