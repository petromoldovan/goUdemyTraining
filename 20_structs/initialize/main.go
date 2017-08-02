package main

import (
	"fmt"
)

type Address struct {
	Street					string
	Housenumber			int
	City						string
	Zip 						int
}

type Hobby struct {
	Name						[]string
	Number			 		int
}

type User struct {
	ID 							int
	FirstName				string
	Surname 				string
	Address 				Address
	Hobby						Hobby
}

func (u User)	getFullName() string {
	return u.FirstName + " " + u.Surname
}

func (u User) getHobbies() []string {
	return u.Hobby.Name
}

func main() {

	user1 := User{
		ID: 1,
		FirstName: "Pit",
		Surname: "Bad",
		Address: Address {
			Street: "Adresleber",
			Housenumber: 12,
			City: "Berlin",
			Zip: 12344,
		},
		Hobby: Hobby {
			Name: []string{"Running", "Swimming"},
			Number: 2,
		},
	}

	fmt.Println(user1)
	fmt.Println(user1.getFullName())
	fmt.Println(user1.getHobbies())


	userHobbies := user1.getHobbies()
	fmt.Println(userHobbies[0])

	for idx, v := range userHobbies {
		fmt.Println(idx, "-- Hobby is --", v)
	}
}
