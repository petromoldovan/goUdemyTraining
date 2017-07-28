package main

import "fmt"

func switcher(arg string) {
	switch arg {
		case "mike":
		fmt.Println("mike")
		case "sam":
		fmt.Println("%v", arg )
		default:
		fmt.Println("default" )
	}
}

func main() {
	switcher("sam")
}
