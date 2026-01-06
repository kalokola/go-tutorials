package main

import (
	"example.com/greetings"
	"fmt"
)

func main() {
	message := greetings.Hello("Brightius")
	fmt.Println(message)
}