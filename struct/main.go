package main

import (
	"log"
	"struct/structsSub"
)

func main() {
	user := structsSub.User{
		FirstName: "Aimu",
		LastName:  "Nvdu",
		Age:       26,
	}

	log.Println("FirstName :", user.FirstName, "LastName :", user.LastName, "Age:", user.Age)
	log.Println("Printing using receiver and function")
	log.Println("FirstName: ", user.PrintFirstName())
}
