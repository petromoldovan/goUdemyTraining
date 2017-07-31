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


	for key, value := range myMap {
		fmt.Println(key, "---", value)
	}

}

