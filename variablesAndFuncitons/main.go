package main

import "fmt"

func main() {
	fmt.Println("Hello, world!!")

	var whatToSay string
	var number int

	whatToSay = "I started learning go"
	fmt.Println(whatToSay)

	number = 9
	fmt.Println("number is set to ", number)

	whatIsSaidFirst, whatIsSaidSecond := saySomething()
	fmt.Println("The function returned", whatIsSaidFirst, whatIsSaidSecond, "!!")
}

func saySomething() (string, string) {
	return "something", "else"
}
