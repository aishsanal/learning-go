package main

import "log"

type User struct {
	firstName string
	lastName  string
}

func main() {
	myMap := make(map[string]User)

	me := User{
		firstName: "Aimusan",
		lastName:  "PP",
	}

	myMap["me"] = me

	log.Println("FirstName:", myMap["me"].firstName, "lastName:", myMap["me"].lastName)
}