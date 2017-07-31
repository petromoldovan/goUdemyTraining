package main

import (
	"fmt"
)

//NOTE: undelying data structure under map is HASH TABLE

func main() {
	myMap := make(map[string]string)

	myMap["student1"] = "Max"
	myMap["student2"] = "Timmy"
	myMap["student3"] = "Jessy"

	fmt.Println(myMap)
	fmt.Println(myMap["student2"])
	fmt.Printf("%T \n", myMap)


	delete(myMap, "student2")
	fmt.Println(myMap)

	fmt.Println(len(myMap))
}

