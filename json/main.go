package main

import (
	"encoding/json"
	"log"
)

type F1 struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsGood    bool   `json:"isGood"`
}

func main() {
	myJson := `
	[
		{
			"first_name": "suttu",
			"last_name": "wifi",
			"isGood": false
		}, 
		{
			"first_name": "mittu",
			"last_name": "tharaav",
			"isGood": true
		}
	]
	`
	var unmarshalled []F1
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Reading json failed")
	} else {
		log.Printf("Unmarshalled : %v", unmarshalled)
	}

	newJson, err := json.MarshalIndent(unmarshalled, "", "        ")
	if err != nil {
		log.Println("Writing json failed")
	} else {
		log.Println(string(newJson))
	}
}
