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
}
