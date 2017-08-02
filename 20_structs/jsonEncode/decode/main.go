package main

import (
	"encoding/json"
	"strings"
	"fmt"
)

type Person struct {
	First 				string
	Last 					string
	Age 					int
	notExported 	int
}


//encode and decode for situations when smth is coming in. Marshal/unmarshal is for internal opertaions
func main() {
	var p1 Person

	rdr := strings.NewReader(`{"First":"Sam","Last":"Michels","Age":22}`)
	json.NewDecoder(rdr).Decode(&p1)

	fmt.Println(p1)
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
}
