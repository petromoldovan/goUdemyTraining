package main

import "fmt"

func main(){

	for i:=0; i < 300 ;i++ {
		if i % 2 == 0 && i >= 50 {
			fmt.Println([]byte(string(i)))
		}
	}
}
