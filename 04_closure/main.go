package main

import "fmt"

func main() {
	x := 42
	{
		fmt.Println(x)
		y := "works inside closure"
		fmt.Println(y)
	}
	//fmt.Println(y)
}