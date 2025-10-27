package main

import "log"

type Animal interface {
	Says() string
	NoOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

func main() {
	dog := Dog{Name: "Johny", Breed: "Minpin"}
	log.Println(dog.Name, dog.Breed, dog.NoOfLegs(), dog.Says())
}

func (dog *Dog) NoOfLegs() int {
	return 4
}

func (dog *Dog) Says() string {
	return "Boww"
}
