package main

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	First 	string
	Last 		string
	Age 		int
}

func main() {
	var p1 Person
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)


	bs := []byte(`{"First":"Sam","Last":"Michels","Age":22}`)
	//here take string of bytes and insert its fields into struct
	json.Unmarshal(bs, &p1)

	fmt.Println("------------------")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
}
