package main

import (
	"log"
	"math/rand"
)

const pool = 1000

func generateRandInt(intChan chan int) {
	intChan <- rand.Intn(pool)
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go generateRandInt(intChan)
	num := <-intChan
	log.Println(num)
}
