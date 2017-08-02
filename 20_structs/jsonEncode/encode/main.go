package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	First 				string
	Last 					string
	Age 					int
	notExported 	int
}


//encode and decode for situations when smth is coming in. Marshal/unmarshal is for internal opertaions
func main() {
	p1 := Person{"James", "Jefferson", 22, 1}
	json.NewEncoder(os.Stdout).Encode(p1)
}
