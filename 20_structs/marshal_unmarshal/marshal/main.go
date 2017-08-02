package main

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	First 			string	`json:"first_name"`
	Last 				string
	Age 				int			`json:"user_age"`
	notExported int
}

func main() {
	p1 := &Person{"Sam", "Michels", 22, 1}

	//Marschal transforms input into slice of bytes
	bs, _ := json.Marshal(p1)

	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
}
