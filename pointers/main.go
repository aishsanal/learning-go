package main

import "log"

func main() {
	var colour string = "Green"
	log.Println("colour is set to", colour)
	changeColor(&colour)
	log.Println("After function call, colour is set to", colour)
}

func changeColor(color *string) {
	var red string = "Red"
	*color = red
}
