package main

import "log"

var string1 string = "cat"

func main() {
	ifElse()
	switchCase()
}

func ifElse() {
	if string1 == "dog" {
		log.Println("It is a dog")
	} else if string1 == "cat" {
		log.Println("It is a cat")
	} else {
		log.Println("It is something else")
	}
}

func switchCase() {
	switch string1 {
		case "dog":
			log.Println("It is a dog")
		case "cat":
			log.Println("It is a cat")
		default:
			log.Println("It is something else")
	}
}