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

func divide(x, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New("cannot divide by zero")
	} else {
		return x / y, nil
	}
}
