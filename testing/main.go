package main

import (
	"errors"
	"log"
)

func main() {
	result, error := divide(4, 0)
	if error != nil {
		log.Println("Divison failed")
		return
	}
	log.Println("Result of division: ", result)
}

func divide(x, y int) (int, error) {
	var result int

	if y == 0 {
		return result, errors.New("cannot divide by zero")
	} else {
		return x / y, nil
	}
}
