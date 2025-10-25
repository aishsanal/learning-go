package main

import (
	"log"
	"sort"
)

func main() {
	var mySlice []int

	mySlice = append(mySlice, 3)
	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 1)
	log.Println(mySlice)
	
	sort.Ints(mySlice)
	log.Println(mySlice)

	mySlice1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	log.Println(mySlice1[0:3])
}