package main

import (
	"fmt"
)

func main() {
	myMap := map[int]string {
		0: "string",
		1: "bool",
		2: "number",
	}

	fmt.Println(myMap)

	//here we initialize two vars v, exists and work with them withing if scope
	if v, exists := myMap[2]; exists {
		fmt.Println("val: ", v)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("Value does not exists")
		fmt.Println("exists: ", exists)
	}

}

